package do

import (
	"time"
)

type User interface {
	Create(req CreateOrUpdateUserReq) error
	Update(req CreateOrUpdateUserReq) error
	ChangePassword(userID int64, newPassword string) error
	Info(id int64) (info *UserInfo, err error)
	List(req UserListReq) (userList []*UserInfo, total int, err error)
	UpdateUserStatus(id int64, status int64) error
	DeleteUser(id int64) error
	UpdateProfile(req UpdateUserProfileReq) error
	SetRole(id, roleID int64) error

	SetDefaultVenue(id, venueId int64) error
}

type CreateOrUpdateUserReq struct {
	ID       int64  `json:"id"`
	Avatar   string `json:"avatar"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Status   int64  `json:"status"`
	Name     string `json:"Name"`
	Wecom    string `json:"wecom"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	Password string `json:"password"`
	CreateId int64  `json:"createId"`
}

type UserInfo struct {
	ID               int64     `json:"id"`
	Status           int8      `json:"status"`
	Username         string    `json:"username"`
	Password         string    `json:"password"`
	Nickname         string    `json:"nickname"`
	SideMode         string    `json:"sideMode"`
	BaseColor        string    `json:"baseColor"`
	ActiveColor      string    `json:"activeColor"`
	RoleID           int64     `json:"roleID"`
	Mobile           string    `json:"mobile"`
	Email            string    `json:"email"`
	Wecom            string    `json:"wecom"`
	Avatar           string    `json:"avatar"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	RoleName         string    `json:"roleName"`
	RoleValue        string    `json:"roleValue"`
	Gender           string    `json:"gender"`
	Age              int64     `json:"age"`
	Birthday         string    `json:"birthday"`
	Job              string    `json:"job"`
	JobName          string    `json:"jobName"`
	Organization     string    `json:"organization"`
	OrganizationName string    `json:"organizationName"`
	DefaultVenueId   int64     `json:"defaultVenueId"`
}

type UserListReq struct {
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	RoleID   int64  `json:"roleID"`
}

type UpdateUserProfileReq struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Avatar   string `json:"avatar"`
}
