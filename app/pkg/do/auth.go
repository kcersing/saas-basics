package do

type Auth interface {
	UpdateApiAuth(roleIDStr string, infos []*ApiAuthInfo) error
	ApiAuth(roleIDStr string) (infos []*ApiAuthInfo, err error)
	UpdateMenuAuth(roleID int64, menuIDs []int64) error
	MenuAuth(roleID int64) (menuIDs []int64, err error)
}

type ApiAuthInfo struct {
	Path   string `json:"path,omitempty"`
	Method string `json:"method,omitempty"`
}
