package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	user2 "saas/biz/dal/db/ent/user"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/user"
	"saas/pkg/encrypt"
)

type Login struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func NewLogin(ctx context.Context, c *app.RequestContext) do.Login {
	return &Login{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}

func (l *Login) Login(username, password string) (res *user.LoginResp, err error) {
	//TODO implement me
	only, err := l.db.User.Query().
		Where(
			user2.UsernameEQ(username),
			user2.Status(1),
		).Only(l.ctx)
	if err != nil {
		return nil, err
	}
	if only == nil {
		err = errors.New("login user not exist")
		return nil, err
	}
	//VerifyPassword
	if ok := encrypt.VerifyPassword(password, only.Password); !ok {
		err = errors.New("wrong password")
		return nil, err
	}

	info, err := NewUser(l.ctx, l.c).Info(only.ID)
	if err != nil {
		return nil, err
	}

	res = new(user.LoginResp)
	res.Username = info.Username
	res.UserId = info.ID
	res.UserRole = info.UserRole
	res.UserRoleIds = info.UserRoleIds
	return res, err
}
