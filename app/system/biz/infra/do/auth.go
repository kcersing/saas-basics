package do

import (
	"rpc_gen/kitex_gen/system/auth"
)

type Auth interface {
	UpdateApiAuth(roleIDStr string, apis []int64) error
	ApiAuth(roleIDStr string) (infos []*auth.ApiAuthorityInfo, err error)
	UpdateMenuAuth(roleID int64, menuIDs []int64) error
	MenuAuth(roleID int64) (menuIDs []int64, err error)
}
