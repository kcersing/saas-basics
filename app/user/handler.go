package main

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	user "rpc_gen/kitex_gen/user"
	"user/biz/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// UserPermCode implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserPermCode(ctx context.Context, req *base.Empty) (resp *base.NilResponse, err error) {
	resp, err = service.NewUserPermCodeService(ctx).Run(req)

	return resp, err
}

// ChangePassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangePassword(ctx context.Context, req *user.ChangePasswordReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewChangePasswordService(ctx).Run(req)

	return resp, err
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateOrUpdateUserReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewCreateUserService(ctx).Run(req)

	return resp, err
}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *user.CreateOrUpdateUserReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateUserService(ctx).Run(req)

	return resp, err
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *base.Empty) (resp *base.NilResponse, err error) {
	resp, err = service.NewUserInfoService(ctx).Run(req)

	return resp, err
}

// UserList implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserList(ctx context.Context, req *user.UserListReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUserListService(ctx).Run(req)

	return resp, err
}

// DeleteUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *base.IDReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteUserService(ctx).Run(req)

	return resp, err
}

// UpdateProfile implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateProfile(ctx context.Context, req *user.ProfileReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateProfileService(ctx).Run(req)

	return resp, err
}

// UserProfile implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserProfile(ctx context.Context, req *base.Empty) (resp *base.NilResponse, err error) {
	resp, err = service.NewUserProfileService(ctx).Run(req)

	return resp, err
}

// UpdateUserStatus implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUserStatus(ctx context.Context, req *base.StatusCodeReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateUserStatusService(ctx).Run(req)

	return resp, err
}

// SetUserRole implements the UserServiceImpl interface.
func (s *UserServiceImpl) SetUserRole(ctx context.Context, req *user.SetUserRole) (resp *base.NilResponse, err error) {
	resp, err = service.NewSetUserRoleService(ctx).Run(req)

	return resp, err
}

// SetDefaultVenue implements the UserServiceImpl interface.
func (s *UserServiceImpl) SetDefaultVenue(ctx context.Context, req *user.SetDefaultVenueReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewSetDefaultVenueService(ctx).Run(req)

	return resp, err
}
