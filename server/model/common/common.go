package common

import "time"

type CacheCaptcha struct {
	Code     string    `json:"code"`     // 验证码
	Overtime time.Time `json:"overtime"` //过期时间
}
