package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
)

type User struct {
	ctx  context.Context
	c    *app.RequestContext
	salt string
	db   *ent.Client
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
	panic("implement me")
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
		ctx:  ctx,
		c:    c,
		salt: config.GlobalServerConfig.MysqlInfo.Salt,
		db:   infras.DB,
	}
}
