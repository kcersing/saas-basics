package do

import (
	"time"
)

type User interface {
	Create(req CreateOrUpdateUserReq) error
	Update(req CreateOrUpdateUserReq) error
	ChangePassword(userID int64, oldPassword, newPassword string) error
	UserInfo(id int64) (userInfo *UserInfo, err error)
	List(req UserListReq) (userList []*UserInfo, total int, err error)
	UpdateUserStatus(id int64, status int64) error
	DeleteUser(id int64) error
	UpdateProfile(req UpdateUserProfileReq) error
}

type CreateOrUpdateUserReq struct {
	ID       int64  `json:"id,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	RoleID   int64  `json:"roleID,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	Email    string `json:"email,omitempty"`
	Status   int64  `json:"status,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Nickname string `json:"nickname,omitempty"`
}

type UserInfo struct {
	ID          int64     `json:"id,omitempty"`
	Status      int8      `json:"status,omitempty"`
	Username    string    `json:"username,omitempty"`
	Password    string    `json:"password,omitempty"`
	Nickname    string    `json:"nickname,omitempty"`
	SideMode    string    `json:"sideMode,omitempty"`
	BaseColor   string    `json:"baseColor,omitempty"`
	ActiveColor string    `json:"activeColor,omitempty"`
	RoleID      int64     `json:"roleID,omitempty"`
	Mobile      string    `json:"mobile,omitempty"`
	Email       string    `json:"email,omitempty"`
	Wecom       string    `json:"wecom,omitempty"`
	Avatar      string    `json:"avatar,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	RoleName    string    `json:"roleName,omitempty"`
	RoleValue   string    `json:"roleValue,omitempty"`
}

type UserListReq struct {
	Page     int64  `json:"page,omitempty"`
	PageSize int64  `json:"pageSize,omitempty"`
	Username string `json:"username,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Email    string `json:"email,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	RoleID   int64  `json:"roleID,omitempty"`
}

type UpdateUserProfileReq struct {
	ID       int64  `json:"id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Email    string `json:"email,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}
