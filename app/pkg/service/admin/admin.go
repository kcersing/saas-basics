package admin

//import (
//	"context"
//	"github.com/cloudwego/hertz/pkg/app"
//	"saas/app/biz/do"
//	"saas/app/biz/model/admin/admin"
//	"saas/app/config"
//	"saas/pkg/db"
//	"saas/pkg/db/ent"
//)
//
//type Admin struct {
//	ctx  context.Context
//	c    *app.RequestContext
//	salt string
//	db   *ent.Client
//}
//
//func (a Admin) GetAdminByAccountId(aid string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (a Admin) GetAdminByName(name string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (a Admin) UpdateAdminPassword(aid string, password string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (a Admin) Login(ctx context.Context, req admin.AdminLoginRequest) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//// NewAdmin create admin service
//func NewAdmin(ctx context.Context, c *app.RequestContext) do.Admin {
//	return &Admin{
//		ctx:  ctx,
//		c:    c,
//		salt: config.GlobalServerConfig.MysqlInfo.Salt,
//		db:   db.InitDB(config.GlobalServerConfig.MysqlInfo.Host),
//	}
//}
