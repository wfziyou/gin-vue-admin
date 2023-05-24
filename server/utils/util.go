package utils

import (
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"strings"
)

const LatestVersion = "latestVersion"     //当前版本
const RequiredVersion = "requiredVersion" // 最小版本
const MiniProgramId = "miniProgramId"     //小程序id
const ConfigParamKey = "configParam"
const MiniProgramKey = "miniProgram"
const BrowsingUserNum = 5

//HomePageTopNewsNum 首页置顶资讯数
const HomePageTopNewsNum = 3

//CircleHotApplyNum 圈子热门应用数
const CircleHotApplyNum = 6

const PageSizeMax = 20

// 类型：0 测试，1注册，2修改密码，3绑定电话，4忘记密码，5绑定银行
const (
	//VerificationLogin 0登录
	VerificationLogin = 0
	//VerificationRegister 1注册
	VerificationRegister = 1
	//VerificationChangePwd 2修改密码
	VerificationChangePwd = 2
	//VerificationBindTel 3绑定电话
	VerificationBindTel = 3
	//VerificationResetPwd 4忘记密码
	VerificationResetPwd = 4
	//VerificationBindBank 5绑定银行
	VerificationBindBank = 5
)

// 类型：0绑定邮箱
const (
	//EmailCodeBindEmail 0绑定邮箱
	EmailCodeBindEmail = 0
)

//
const (
	QuestionStatusOpen  = 0
	QuestionStatusClose = 1
)

//付费货币：1人民、2积分、3金币
const (
	CurrencyGold  = 1
	CurrencyMoney = 2
	CurrencyScore
)

func UuidTo32String(u uuid.UUID) string {
	buf := make([]byte, 32)
	hex.Encode(buf[0:8], u[0:4])
	hex.Encode(buf[8:12], u[4:6])
	hex.Encode(buf[12:16], u[6:8])
	hex.Encode(buf[16:20], u[8:10])
	hex.Encode(buf[20:], u[10:])
	return string(buf)
}
func SplitToInt32List(str string, sep string) (i32List []int32) {
	if str == "" {
		return
	}
	strList := strings.Split(str, sep)
	if len(strList) == 0 {
		return
	}
	for _, item := range strList {
		if item == "" {
			continue
		}
		val, err := strconv.ParseInt(item, 10, 32)
		if err != nil {
			// logs.CtxError(ctx, "ParseInt fail, err=%v, str=%v, sep=%v", err, str, sep) // 此处打印出log报错信息
			continue
		}
		i32List = append(i32List, int32(val))
	}
	return
}
func SplitToUint64List(str string, sep string) (i32List []uint64) {
	if str == "" {
		return
	}
	strList := strings.Split(str, sep)
	if len(strList) == 0 {
		return
	}
	for _, item := range strList {
		if item == "" {
			continue
		}
		val, err := strconv.ParseUint(item, 10, 64)
		if err != nil {
			// logs.CtxError(ctx, "ParseInt fail, err=%v, str=%v, sep=%v", err, str, sep) // 此处打印出log报错信息
			continue
		}
		i32List = append(i32List, val)
	}
	return
}

// 类型转换
func InterfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}

func IdsToArr(data string) (_ids []uint64) {
	var slice = strings.Split(data, ",")
	length := len(slice)
	ids := make([]uint64, 0, length)
	for i := 0; i < length; i++ {
		id, _ := strconv.ParseUint(slice[i], 10, 64)
		ids = append(ids, id)
	}
	return ids
}
