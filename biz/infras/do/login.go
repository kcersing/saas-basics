package do

type Login interface {
	Login(username, password string) (res *LoginResp, err error)
}

type LoginResp struct {
	UserID    int64  `json:"userID"`
	Username  string `json:"username"`
	RoleName  string `json:"roleName"`
	RoleValue string `json:"roleValue"`
	RoleID    int64  `json:"roleID"`
}
