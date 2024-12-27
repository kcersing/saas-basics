package do

import (
	"saas/idl_gen/model/user"
)

type User interface {
	Create(req user.CreateOrUpdateUserReq) error
	Update(req user.CreateOrUpdateUserReq) error
	Info(id int64) (info *user.UserInfo, err error)
	List(req user.UserListReq) (userList []*user.UserInfo, total int, err error)

	ChangePassword(userID int64, newPassword string) error
	UpdateUserStatus(id int64, status int64) error
	DeleteUser(id int64) error
	SetRole(id int64, roleID []int64) error
	SetDefaultVenue(id, venueId int64) error
}
