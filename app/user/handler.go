package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/paseto"
	user "saas/kitex_gen/cwg/user"
	"saas/pkg/db/ent"
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
	CreateUser(req *user.AddUserRequest) (*ent.User, error)
	GetUserByMobile(mobile string) (*ent.User, error)

	GetUserByAccountId(aid string) (*ent.User, error)
	GetUserList() ([]*ent.User, error)

	DeleteUser(aid string) error
}
type AdminMysqlManager interface {
	GetAdminByAccountId(aid string) error
	GetAdminByName(name string) error
	UpdateAdminPassword(aid string, password string) error
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

	resp = new(user.AddUserResponse)
	u, err := s.UserMysqlManager.CreateUser(req)
	if err != nil {
		//if errors.Is(err, errno.RecordAlreadyExist) {
		//	klog.Error("add user error", err)
		//	resp.BaseResp = utils.BuildBaseResp(errno.RecordAlreadyExist)
		//	return resp, nil
		//}
		//klog.Error("add user error", err)
		//resp.BaseResp = utils.BuildBaseResp(errno.UserSrvErr)
		//return resp, nil
	}
	klog.Info(u, err)
	//resp.BaseResp = utils.BuildBaseResp(nil)
	resp.BaseResp.StatusCode = 20000
	resp.BaseResp.StatusMsg = "aaaaaa"
	return resp, nil
}

// DeleteUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (resp *user.DeleteUserResponse, err error) {
	// TODO: Your code here...
	return
}
