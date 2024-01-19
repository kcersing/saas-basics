package main

import (
	"context"
	user "saas_basics/kitex_gen/cwg/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// AdminLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) AdminLogin(ctx context.Context, req *user.AdminLoginRequest) (resp *user.AdminLoginResponse, err error) {
	// TODO: Your code here...

	resp = new(user.AdminLoginResponse)

	return
}
