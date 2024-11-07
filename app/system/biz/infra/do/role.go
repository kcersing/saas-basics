package do

import (
	"rpc_gen/kitex_gen/base"
	"rpc_gen/kitex_gen/system/role"
)

type Role interface {
	Create(req *role.RoleInfo) error
	Update(req *role.RoleInfo) error
	Delete(id int64) error
	RoleInfoByID(ID int64) (roleInfo *role.RoleInfo, err error)
	List(req *base.PageInfoReq) (roleInfoList []*role.RoleInfo, total int, err error)
	UpdateStatus(ID int64, status int64) error
}
