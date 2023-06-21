package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/apply"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type ConfigParam struct {
	ActivityManagerApply *apply.Apply      `json:"activityManagerApply" yaml:"activityManagerApply"` //活动管理访问应用
	CircleManagerApply   *apply.Apply      `json:"circleManagerApply" yaml:"circleManagerApply"`     //圈子管理访问应用
	WalletBillApply      *apply.Apply      `json:"walletBillApply" yaml:"walletBillApply"`           //钱包账单访问应用
	GoldBillApply        *apply.Apply      `json:"goldBillApply" yaml:"goldBillApply"`               //金币账单访问应用
	GoldShopApply        *apply.Apply      `json:"goldShopApply" yaml:"goldShopApply"`               //金币商城访问应用
	ParamList            []system.SysParam `json:"paramList" yaml:"paramList"`                       //参数参数
}

type ResponseCheckAppUpdate struct {
	Version     AppVersion `json:"version"`      //最新版本
	ForceUpdate int        `json:"forceUpdate" ` //强制更新：0否 1是
}
