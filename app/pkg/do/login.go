package do

type Login interface {
	Login(username, password string) (res *LoginResp, err error)
}

type LoginResp struct {
	UserID    int64  `json:"userID,omitempty"`
	Username  string `json:"username,omitempty"`
	RoleName  string `json:"roleName,omitempty"`
	RoleValue string `json:"roleValue,omitempty"`
	RoleID    int64  `json:"roleID,omitempty"`
}
