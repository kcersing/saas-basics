package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/user"
	"strconv"
	"time"
)

type User struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (u User) Create(req do.CreateOrUpdateUserReq) error {
	//TODO implement me
	panic("implement me")
}

func (u User) Update(req do.CreateOrUpdateUserReq) error {
	//TODO implement me
	panic("implement me")
}

func (u User) ChangePassword(userID uint64, oldPassword, newPassword string) error {
	//TODO implement me
	panic("implement me")
}

func (u User) UserInfo(id uint64) (userInfo *do.UserInfo, err error) {
	//TODO implement me
	userInfo = new(do.UserInfo)
	userInterface, exist := u.cache.Get("userInfo" + strconv.Itoa(int(id)))
	if exist {
		if u, ok := userInterface.(*do.UserInfo); ok {
			return u, nil
		}
	}
	userEnt, err := u.db.User.Query().Where(user.IDEQ(id)).First(u.ctx)
	if err != nil {
		err = errors.Wrap(err, "get user failed")
		return userInfo, err
	}
	hlog.Info(userEnt)
	err = copier.Copy(&userInfo, &userEnt)
	if err != nil {
		err = errors.Wrap(err, "copy user info failed")
		return userInfo, err
	}

	roleInterface, exist := u.cache.Get("roleData" + strconv.Itoa(int(userInfo.RoleID)))
	if exist {
		if role, ok := roleInterface.(*ent.Role); ok {
			userInfo.RoleName = role.Name
			userInfo.RoleValue = role.Value
		}
	}
	u.cache.SetWithTTL("userInfo"+strconv.Itoa(int(userInfo.ID)), &userInfo, 1, 72*time.Hour)
	return
}

func (u User) List(req do.UserListReq) (userList []*do.UserInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (u User) UpdateUserStatus(id uint64, status uint64) error {
	//TODO implement me
	panic("implement me")
}

func (u User) DeleteUser(id uint64) error {
	//TODO implement me
	panic("implement me")
}

func (u User) UpdateProfile(req do.UpdateUserProfileReq) error {
	//TODO implement me
	panic("implement me")
}

func NewUser(ctx context.Context, c *app.RequestContext) do.User {
	return &User{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
