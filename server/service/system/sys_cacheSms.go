package system

import (
	"context"
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"time"
)

type CacheSmsService struct{}

//@function: GetCacheSms
//@description: 从redis取jwt
//@return: cacheCaptcha common.CacheCaptcha, err error

func (cacheSmsService *CacheSmsService) GetCacheSms(key string) (cacheCaptcha common.CacheCaptcha, err error) {
	cacheCaptchaStr, err := global.GVA_REDIS.Get(context.Background(), key).Bytes()
	json.Unmarshal(cacheCaptchaStr, &cacheCaptcha)
	return cacheCaptcha, err
}

//@function: SetCacheSms
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

func (cacheSmsService *CacheSmsService) SetCacheSms(key string, code string) (err error) {
	// 过期时间
	dr, err := utils.ParseDuration(global.GVA_CONFIG.CacheCaptcha.ExpiresTime)
	if err != nil {
		return err
	}
	// 发短信间隔时间
	sms, err := utils.ParseDuration(global.GVA_CONFIG.CacheCaptcha.SmsTime)
	if err != nil {
		return err
	}
	jsonBytes, err := json.Marshal(common.CacheCaptcha{
		Code:     code,
		Overtime: time.Now().Add(sms),
	})
	timer := dr
	err = global.GVA_REDIS.Set(context.Background(), key, jsonBytes, timer).Err()
	return err
}
