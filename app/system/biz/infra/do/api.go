package do

import (
	"rpc_gen/kitex_gen/base"
	"rpc_gen/kitex_gen/system/auth"
)

type Api interface {
	Create(req auth.ApiInfo) error
	Update(req auth.ApiInfo) error
	Delete(id int64) error
	List(req auth.ApiPageReq) (resp []*auth.ApiInfo, total int, err error)
	ApiTree(req auth.ApiPageReq) (resp []*base.Tree, total int, err error)
}
