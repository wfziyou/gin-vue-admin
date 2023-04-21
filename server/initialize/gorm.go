package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/apply"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/general"
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"

	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(

		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},
		system.SysChatGptOption{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},

		community.ForumPosts{},
		community.Circle{},
		community.CircleAddRequest{},
		community.CircleClassify{},
		community.CircleRelation{},
		community.CircleRequest{},
		community.CircleUser{},
		community.CommentThumbsUp{},
		community.ForbiddenSpeak{},
		community.ForbiddenSpeakDuration{},
		community.ForbiddenSpeakReason{},
		community.ForumComment{},
		community.ForumTag{},
		community.ForumTagGroup{},
		community.ForumThumbsUp{},
		community.ForumTopic{},
		community.ForumTopicGroup{},
		community.ForumTopicPostsMapping{},
		community.Report{},
		community.UserCircleApply{},
		community.ReportReason{},
		community.UserExtend{},
		community.User{},
		community.HkChannel{},
		community.FocusUser{},
		community.ActivityUser{},
		community.ActivityAddRequest{},
		community.RecordBrowsingUserHomepage{},
		community.RecordBrowsingCircleHomepage{},
		community.HkThirdPlatform{},
		community.HkGoldBill{},
		community.HkUserBill{},
		community.HkUserRecharge{},
		community.HkUserExtract{},
		community.HkOrder{},
		
		apply.Apply{},
		apply.CircleApply{},
		apply.CircleApplyGroup{},
		apply.PlatApply{},
		apply.PlatApplyGroup{},

		general.MiniProgram{},
		general.MiniProgramPacket{},
		general.Protocol{},
		general.BugReport{},
		general.UserCollect{},
		general.UserBrowsingHistory{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
