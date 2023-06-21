package system

import (
	"context"
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/apply"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/general"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
	"strconv"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSystemConfig
//@description: 读取配置文件
//@return: conf config.Server, err error

type SystemConfigService struct{}

func (systemConfigService *SystemConfigService) GetSystemConfig() (conf config.Server, err error) {
	return global.GVA_CONFIG, nil
}

func (systemConfigService *SystemConfigService) GetParam() (conf config.Param, err error) {
	return global.GVA_CONFIG.Param, nil
}

// @description   set system config,
//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetSystemConfig
//@description: 设置配置文件
//@param: system model.System
//@return: err error

func (systemConfigService *SystemConfigService) SetSystemConfig(system system.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	err = global.GVA_VP.WriteConfig()
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetServerInfo
//@description: 获取服务器信息
//@return: server *utils.Server, err error

func (systemConfigService *SystemConfigService) GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.GVA_LOG.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Ram, err = utils.InitRAM(); err != nil {
		global.GVA_LOG.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.GVA_LOG.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}

func LoadConfigParamToCache() {
	var data []system.SysParam
	err := global.GVA_DB.Model(&system.SysParam{}).Where("param_key LIKE 'app.%'").Select("param_key,param_value").Find(&data).Error
	if err != nil {
		global.GVA_LOG.Error("加载数据库SysParam失败!", zap.Error(err))
		return
	}
	var ids = make([]uint64, 0, 4)
	var mapApply = make(map[string]interface{})
	for _, obj := range data {
		bytes, _ := json.Marshal(obj) // Marshal函数转成的是 byte 数组

		if obj.ParamKey == utils.MiniProgramId {
			var data general.MiniProgramBaseInfo
			miniProgramId, _ := strconv.ParseInt(obj.ParamValue, 10, 64)
			err1 := global.GVA_DB.Raw("SELECT t.name,t.icon,t.company_name,t.program_id,p.version,p.code,p.packet_address,t.hidden FROM hk_mini_program t LEFT JOIN hk_mini_program_packet p ON  t.cur_packet_id = p.id WHERE t.id = ?", miniProgramId).First(&data).Error
			if err1 != nil {
				global.GVA_LOG.Error("查找 miniProgramId 失败", zap.Error(err1))
			} else {
				str, _ := json.Marshal(data)
				global.GVA_REDIS.HSet(context.Background(), utils.ConfigParamKey, utils.MiniProgramKey, str)
			}
		} else if obj.ParamKey == utils.SysParamActivityManagerApplyId ||
			obj.ParamKey == utils.SysParamCircleManagerApplyId ||
			obj.ParamKey == utils.SysParamWalletBillApplyId ||
			obj.ParamKey == utils.SysParamGoldBillApplyId ||
			obj.ParamKey == utils.SysParamGoldShopApplyId {
			id, _ := strconv.ParseUint(obj.ParamValue, 10, 64)
			ids = append(ids, id)
			mapApply[obj.ParamKey] = id
		} else {
			global.GVA_REDIS.HSet(context.Background(), utils.ConfigParamKey, obj.ParamKey, bytes)
		}
	}
	var applyList []apply.Apply
	err = global.GVA_DB.Model(&apply.Apply{}).Where("id in ?", ids).Find(&applyList).Error
	if err == nil {
		var miniIds = make([]uint64, 0, 4)
		for _, obj := range applyList {
			if obj.Type == utils.ApplyTypeMimiProgram {
				miniIds = append(miniIds, obj.MiniProgramId)
			}
		}
		if len(miniIds) > 0 {
			var miniProgram []apply.MiniProgramBaseInfo
			err1 := global.GVA_DB.Raw("SELECT t.id,t.name,t.icon,t.company_name,t.program_id,p.version,p.code,p.packet_address,t.hidden FROM hk_mini_program t LEFT JOIN hk_mini_program_packet p ON  t.cur_packet_id = p.id WHERE t.id in ?", miniIds).Find(&miniProgram).Error
			if err1 == nil && len(miniProgram) > 0 {
				for index, obj := range applyList {
					if obj.Type == utils.ApplyTypeMimiProgram {
						for _, program := range miniProgram {
							if obj.MiniProgramId == program.Id {
								applyList[index].ProgramId = program.ProgramId
								applyList[index].MiniProgram = &program
								break
							}
						}
					}
				}
			}
		}
		if len(applyList) > 0 {
			for key, applyId := range mapApply {
				for _, obj := range applyList {
					if obj.ID == applyId {
						bytes, _ := json.Marshal(obj)
						global.GVA_REDIS.HSet(context.Background(), utils.ConfigParamKey, key, bytes)
						break
					}
				}
			}
		}
	}
}
