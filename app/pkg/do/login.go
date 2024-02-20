package do

type Login interface {
	Login(username, password string) (res *LoginResp, err error)
	LoginByOAuth(provider, credential string) (res *LoginResp, err error)
}

type LoginResp struct {
	UserID    uint64
	Username  string
	RoleName  string
	RoleValue string
	RoleID    uint64
}
