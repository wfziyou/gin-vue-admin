package db

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/common/constant"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/common/utils"
	"github.com/go-redis/redis/v8"
)

const (
	accountTempCode               = "ACCOUNT_TEMP_CODE"
	resetPwdTempCode              = "RESET_PWD_TEMP_CODE"
	userIncrSeq                   = "REDIS_USER_INCR_SEQ:" // user incr seq
	appleDeviceToken              = "DEVICE_TOKEN"
	userMinSeq                    = "REDIS_USER_MIN_SEQ:"
	uidPidToken                   = "UID_PID_TOKEN_STATUS_:"
	conversationReceiveMessageOpt = "CON_RECV_MSG_OPT:"
	getuiToken                    = "GETUI_TOKEN"
	getuiTaskID                   = "GETUI_TASK_ID"
	messageCache                  = "MESSAGE_CACHE:"
	SignalCache                   = "SIGNAL_CACHE:"
	SignalListCache               = "SIGNAL_LIST_CACHE:"
	GlobalMsgRecvOpt              = "GLOBAL_MSG_RECV_OPT"
	FcmToken                      = "FCM_TOKEN:"
	groupUserMinSeq               = "GROUP_USER_MIN_SEQ:"
	groupMaxSeq                   = "GROUP_MAX_SEQ:"
	groupMinSeq                   = "GROUP_MIN_SEQ:"
	sendMsgFailedFlag             = "SEND_MSG_FAILED_FLAG:"
	userBadgeUnreadCountSum       = "USER_BADGE_UNREAD_COUNT_SUM:"
	exTypeKeyLocker               = "EX_LOCK:"

	//temp
	superGroupUserNotRecvOfflineMsgOptTemp = "SG_RECV_MSG_OPT_TEMP:"
)

var Redis *RedisModel

type RedisModel struct {
	Client *redis.Client
}

func (d *RedisModel) AddTokenFlag(userID string, platformID int, token string, flag int) error {
	key := uidPidToken + userID + ":" + constant.PlatformIDToName(platformID)
	//global.GVA_LOG.Debug("add token key is " + key)
	return d.Client.HSet(context.Background(), key, token, flag).Err()
}
func (d *RedisModel) GetTokenMapByUidPid(userID, platformID string) (map[string]int, error) {
	key := uidPidToken + userID + ":" + platformID
	//global.GVA_LOG.Debug("get token key is " + key)
	m, err := d.Client.HGetAll(context.Background(), key).Result()
	mm := make(map[string]int)
	for k, v := range m {
		mm[k] = utils.StringToInt(v)
	}
	return mm, err
}
func (d *RedisModel) SetTokenMapByUidPid(userID string, platformID int, m map[string]int) error {
	key := uidPidToken + userID + ":" + constant.PlatformIDToName(platformID)
	mm := make(map[string]interface{})
	for k, v := range m {
		mm[k] = v
	}
	return d.Client.HSet(context.Background(), key, mm).Err()
}
func (d *RedisModel) DeleteTokenByUidPid(userID string, platformID int, fields []string) error {
	key := uidPidToken + userID + ":" + constant.PlatformIDToName(platformID)
	return d.Client.HDel(context.Background(), key, fields...).Err()
}
