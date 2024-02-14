package do

import (
	"context"
	"github.com/hertz-contrib/paseto"
	"saas/app/biz/model/admin"
)

type Admin interface {
	Login(ctx context.Context, req admin.AdminLoginRequest) error

	GetAdminByAccountId(aid string) error
	GetAdminByName(name string) error
	UpdateAdminPassword(aid string, password string) error
}

type Token interface {
	CreateToken(claims *paseto.StandardClaims) (token string, err error)
}
