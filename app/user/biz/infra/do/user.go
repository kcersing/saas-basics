package do

import "rpc_gen/kitex_gen/user"

type User interface {
	Create(req *user.CreateOrUpdateUserReq) error
	Update(req *user.CreateOrUpdateUserReq) error
	ChangePassword(userID int64, newPassword string) error
	Info(id int64) (info *user.UserInfo, err error)
	List(req *user.UserListReq) (userList []*user.UserInfo, total int64, err error)
	UpdateUserStatus(id int64, status int64) error
	DeleteUser(id int64) error
	SetRole(id, roleID int64) error
	SetDefaultVenue(id, venueId int64) error
}
