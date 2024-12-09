package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	user2 "saas/biz/dal/db/ent/user"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/user"
	"strconv"
	"time"
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
	//if ok := encrypt.VerifyPassword(password, only.Password); !ok {
	//	err = errors.New("wrong password")
	//	return nil, err
	//}

	res = new(user.LoginResp)
	res.Username = username
	res.UserId = only.ID
	res.RoleId = only.RoleID
	res.RoleName, res.RoleValue, err = l.getRoleInfo(only.RoleID)

	return res, err
}

func (l *Login) getRoleInfo(roleID int64) (roleName, roleValue string, err error) {
	v, exist := l.cache.Get("roleData" + strconv.Itoa(int(roleID)))
	if exist {
		roleName = v.(*ent.Role).Name
		roleValue = v.(*ent.Role).Value
		return
	}
	roleData, err := l.db.Role.Query().All(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			err = errors.Wrap(err, "fail to find any roles")
			return "", "", err
		}
		hlog.Error(err, "database error")
		return "", "", err
	}

	for _, v := range roleData {
		l.cache.SetWithTTL("roleData"+strconv.Itoa(int(v.ID)), v, 1, 1*time.Hour)
		if v.ID == roleID {
			roleName = v.Name
			roleValue = v.Value
		}
	}
	if roleName == "" || roleValue == "" {
		err = errors.New("fail to find any role info")
	}
	return roleName, roleValue, err

}
