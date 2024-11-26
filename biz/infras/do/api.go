package do

import (
	"saas/idl_gen/model/base"
	"saas/idl_gen/model/menu"
)

type Api interface {
	Create(req menu.ApiInfo) error
	Update(req menu.ApiInfo) error
	Delete(id int64) error
	List(req menu.ApiPageReq) (resp []*menu.ApiInfo, total int64, err error)
	ApiTree(req menu.ApiPageReq) (resp []*base.Tree, total int64, err error)
}
