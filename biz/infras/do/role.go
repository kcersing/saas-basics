package do

import (
	"saas/idl_gen/model/auth"
	"saas/idl_gen/model/base"
)

type Role interface {
	Create(req *auth.RoleInfo) error
	Update(req *auth.RoleInfo) error
	Delete(id int64) error
	RoleInfoByID(ID int64) (roleInfo *auth.RoleInfo, err error)
	List(req *base.PageInfoReq) (roleInfoList []*auth.RoleInfo, total int64, err error)
	UpdateStatus(ID int64, status int64) error
}
