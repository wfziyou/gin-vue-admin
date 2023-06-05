package request

import (
	jwt "github.com/golang-jwt/jwt/v4"
)

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	ID          uint64
	UID         string
	Platform    string
	AuthorityId uint64
}
type BaseClaimsEx struct {
	ID          uint64
	IDStr       string
	Platform    string
	AuthorityId uint64
}
type CustomClaimsEx struct {
	BaseClaimsEx
	BufferTime int64
	jwt.RegisteredClaims
}

func (obj *CustomClaims) GetUserId() (userId uint64) {
	return obj.BaseClaims.ID
}
