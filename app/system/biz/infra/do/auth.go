package do

import (
	"rpc_gen/kitex_gen/system/auth"
)

type Auth interface {
	ApiAuth(roleIDStr string) (infos []*auth.ApiAuthInfo, err error)
	MenuAuth(roleID int64) (menuIDs []int64, err error)

	UpdateApiAuth(roleIDStr string, apis []int64) error
	UpdateMenuAuth(roleID int64, menuIDs []int64) error
}
