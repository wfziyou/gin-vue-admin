package request

import (
	jwt "github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint64
	Username    string
	NickName    string
	AuthorityId uint64
}

func (obj *CustomClaims) GetUserId() (userId uint64) {
	return obj.BaseClaims.ID
}
