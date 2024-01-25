package main

import (
	"context"
	"github.com/hertz-contrib/paseto"
	user "saas/kitex_gen/cwg/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	OpenIDResolver
	EncryptManager
	AdminMysqlManager
	UserMysqlManager
	TokenGenerator
}

// TokenGenerator creates token.
type TokenGenerator interface {
	CreateToken(claims *paseto.StandardClaims) (token string, err error)
}

// OpenIDResolver resolves an authorization code
type OpenIDResolver interface {
	Resolve(code string) string
}

// EncryptManager  encrypt password
type EncryptManager interface {
	EncryptPassword(code string) string
}
type UserMysqlManager interface {
	CreateUser()
}
type AdminMysqlManager interface {
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	return
}

// AdminLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) AdminLogin(ctx context.Context, req *user.AdminLoginRequest) (resp *user.AdminLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// ChangeAdminPassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangeAdminPassword(ctx context.Context, req *user.ChangeAdminPasswordRequest) (resp *user.ChangeAdminPasswordResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.GetUserRequest) (resp *user.GetUserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// AddUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddUser(ctx context.Context, req *user.AddUserRequest) (resp *user.AddUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.AddUserResponse)
	_, err = s.UserMysqlManager.CreateUser()

	if err != nil {
		//if err == errno.NewErrNo()
	}

	return
}

// DeleteUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (resp *user.DeleteUserResponse, err error) {
	// TODO: Your code here...
	return
}
