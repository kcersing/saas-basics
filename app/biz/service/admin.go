package service

import "github.com/hertz-contrib/paseto"

type AdminServiceImpl struct {
	TokenGenerator
}

// TokenGenerator creates token.
type TokenGenerator interface {
	CreateToken(claims *paseto.StandardClaims) (token string, err error)
}

type UserManager struct {
	salt string
	db   *gorm.DB
}
