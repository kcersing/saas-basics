package do

type Auth interface {
	UpdateApiAuth(roleIDStr string, infos []*ApiAuthInfo) error
	ApiAuth(roleIDStr string) (infos []*ApiAuthInfo, err error)
	UpdateMenuAuth(roleID uint64, menuIDs []uint64) error
	MenuAuth(roleID uint64) (menuIDs []uint64, err error)
}

type ApiAuthInfo struct {
	Path   string
	Method string
}
