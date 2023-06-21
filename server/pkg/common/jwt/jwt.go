package jwt

import (
	"context"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/common/constant"
	db "github.com/flipped-aurora/gin-vue-admin/server/pkg/common/db"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/common/utils"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GVA_CONFIG.JWT.SigningKey),
	}
}
func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(global.GVA_CONFIG.JWT.SigningKey), nil
	}
}

func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	bf, _ := utils.ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	ep, _ := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	now := time.Now()
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"GVA"},            // 受众
			NotBefore: jwt.NewNumericDate(now.Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(now.Add(ep)),    // 过期时间 7天  配置文件
			Issuer:    global.GVA_CONFIG.JWT.Issuer,       // 签名的发行者
		},
	}
	return claims
}
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) CreateTokenEx(userId uint64, platformID int, authorityId uint64) (string, int64, error) {
	userIdStr := strconv.FormatUint(userId, 10)
	claims := j.CreateClaims(request.BaseClaims{
		ID:          userId,
		UID:         userIdStr,
		Platform:    constant.PlatformIDToName(platformID),
		AuthorityId: authorityId,
	})
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.SigningKey)
	if err != nil {
		return "", 0, err
	}
	//remove Invalid token
	m, err := db.Redis.GetTokenMapByUidPid(userIdStr, constant.PlatformIDToName(platformID))
	if err != nil && err != redis.Nil {
		return "", 0, err
	}

	var deleteTokenKey []string
	for k, v := range m {
		_, err = GetClaimFromToken(k)
		if err != nil || v != constant.NormalToken {
			deleteTokenKey = append(deleteTokenKey, k)
		}
		deleteTokenKey = append(deleteTokenKey, k)
	}
	if len(deleteTokenKey) != 0 {
		err = db.Redis.DeleteTokenByUidPid(userIdStr, platformID, deleteTokenKey)
		if err != nil {
			return "", 0, err
		}
	}
	err = db.Redis.AddTokenFlag(userIdStr, platformID, tokenString, constant.NormalToken)
	if err != nil {
		return "", 0, err
	}
	return tokenString, claims.ExpiresAt.Time.Unix(), err
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	v, err, _ := global.GVA_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (claims *request.CustomClaims, err error) {
	claims, err = GetClaimFromToken(tokenString)
	if err != nil {
		return nil, utils.Wrap(err, "")
	}

	m, err := db.Redis.GetTokenMapByUidPid(claims.UID, claims.Platform)
	if err != nil {
		global.GVA_LOG.Error("get token from redis err", zap.Error(err))
		return nil, utils.Wrap(constant.ErrTokenInvalid, "get token from redis err")
	}
	if m == nil {
		global.GVA_LOG.Error("get token from redis err, not in redis,m is nil. tokenString:" + tokenString)
		return nil, utils.Wrap(constant.ErrTokenInvalid, "get token from redis err")
	}

	if v, ok := m[tokenString]; ok {
		switch v {
		case constant.NormalToken:
			global.GVA_LOG.Debug("this is normal return. tokenString:" + tokenString)
			return claims, nil
		case constant.KickedToken:
			global.GVA_LOG.Error("this token has been kicked by other same terminal ")
			return nil, utils.Wrap(constant.ErrTokenKicked, "this token has been kicked by other same terminal ")
		default:
			return nil, utils.Wrap(constant.ErrTokenInvalid, "")
		}
	}
	global.GVA_LOG.Error("redis token map not find")
	return nil, utils.Wrap(constant.ErrTokenInvalid, "redis token map not find")
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func (j *JWT) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.GVA_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func (j *JWT) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
	// err := global.GVA_DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	// isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	// return !isNotFound
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: redisJWT string, err error

func (j *JWT) GetRedisJWT(userId uint64) (redisJWT string, err error) {
	redisJWT, err = global.GVA_REDIS.Get(context.Background(), strconv.FormatUint(userId, 10)).Result()
	return redisJWT, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

func (j *JWT) SetRedisJWT(jwt string, userId uint64) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.GVA_REDIS.Set(context.Background(), strconv.FormatUint(userId, 10), jwt, timer).Err()
	return err
}

func LoadAll() {
	var data []string
	err := global.GVA_DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.GVA_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
func GetClaimFromToken(tokensString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokensString, &request.CustomClaims{}, secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, utils.Wrap(constant.ErrTokenMalformed, "")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, utils.Wrap(constant.ErrTokenExpired, "")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, utils.Wrap(constant.ErrTokenNotValidYet, "")
			} else {
				return nil, utils.Wrap(constant.ErrTokenUnknown, "")
			}
		} else {
			return nil, utils.Wrap(constant.ErrTokenNotValidYet, "")
		}
	} else {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, utils.Wrap(constant.ErrTokenNotValidYet, "")
	}
}
