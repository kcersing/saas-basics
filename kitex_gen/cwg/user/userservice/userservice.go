// Code generated by Kitex v0.8.0. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	user "saas/kitex_gen/cwg/user"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Login":               kitex.NewMethodInfo(loginHandler, newUserServiceLoginArgs, newUserServiceLoginResult, false),
		"AdminLogin":          kitex.NewMethodInfo(adminLoginHandler, newUserServiceAdminLoginArgs, newUserServiceAdminLoginResult, false),
		"ChangeAdminPassword": kitex.NewMethodInfo(changeAdminPasswordHandler, newUserServiceChangeAdminPasswordArgs, newUserServiceChangeAdminPasswordResult, false),
		"GetUser":             kitex.NewMethodInfo(getUserHandler, newUserServiceGetUserArgs, newUserServiceGetUserResult, false),
		"AddUser":             kitex.NewMethodInfo(addUserHandler, newUserServiceAddUserArgs, newUserServiceAddUserResult, false),
		"DeleteUser":          kitex.NewMethodInfo(deleteUserHandler, newUserServiceDeleteUserArgs, newUserServiceDeleteUserResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "user",
		"ServiceFilePath": `..\..\idl\rpc\user.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceLoginArgs)
	realResult := result.(*user.UserServiceLoginResult)
	success, err := handler.(user.UserService).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceLoginArgs() interface{} {
	return user.NewUserServiceLoginArgs()
}

func newUserServiceLoginResult() interface{} {
	return user.NewUserServiceLoginResult()
}

func adminLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceAdminLoginArgs)
	realResult := result.(*user.UserServiceAdminLoginResult)
	success, err := handler.(user.UserService).AdminLogin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceAdminLoginArgs() interface{} {
	return user.NewUserServiceAdminLoginArgs()
}

func newUserServiceAdminLoginResult() interface{} {
	return user.NewUserServiceAdminLoginResult()
}

func changeAdminPasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceChangeAdminPasswordArgs)
	realResult := result.(*user.UserServiceChangeAdminPasswordResult)
	success, err := handler.(user.UserService).ChangeAdminPassword(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceChangeAdminPasswordArgs() interface{} {
	return user.NewUserServiceChangeAdminPasswordArgs()
}

func newUserServiceChangeAdminPasswordResult() interface{} {
	return user.NewUserServiceChangeAdminPasswordResult()
}

func getUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetUserArgs)
	realResult := result.(*user.UserServiceGetUserResult)
	success, err := handler.(user.UserService).GetUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetUserArgs() interface{} {
	return user.NewUserServiceGetUserArgs()
}

func newUserServiceGetUserResult() interface{} {
	return user.NewUserServiceGetUserResult()
}

func addUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceAddUserArgs)
	realResult := result.(*user.UserServiceAddUserResult)
	success, err := handler.(user.UserService).AddUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceAddUserArgs() interface{} {
	return user.NewUserServiceAddUserArgs()
}

func newUserServiceAddUserResult() interface{} {
	return user.NewUserServiceAddUserResult()
}

func deleteUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceDeleteUserArgs)
	realResult := result.(*user.UserServiceDeleteUserResult)
	success, err := handler.(user.UserService).DeleteUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceDeleteUserArgs() interface{} {
	return user.NewUserServiceDeleteUserArgs()
}

func newUserServiceDeleteUserResult() interface{} {
	return user.NewUserServiceDeleteUserResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Login(ctx context.Context, req *user.LoginRequest) (r *user.LoginResponse, err error) {
	var _args user.UserServiceLoginArgs
	_args.Req = req
	var _result user.UserServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AdminLogin(ctx context.Context, req *user.AdminLoginRequest) (r *user.AdminLoginResponse, err error) {
	var _args user.UserServiceAdminLoginArgs
	_args.Req = req
	var _result user.UserServiceAdminLoginResult
	if err = p.c.Call(ctx, "AdminLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChangeAdminPassword(ctx context.Context, req *user.ChangeAdminPasswordRequest) (r *user.ChangeAdminPasswordResponse, err error) {
	var _args user.UserServiceChangeAdminPasswordArgs
	_args.Req = req
	var _result user.UserServiceChangeAdminPasswordResult
	if err = p.c.Call(ctx, "ChangeAdminPassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUser(ctx context.Context, req *user.GetUserRequest) (r *user.GetUserInfoResponse, err error) {
	var _args user.UserServiceGetUserArgs
	_args.Req = req
	var _result user.UserServiceGetUserResult
	if err = p.c.Call(ctx, "GetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddUser(ctx context.Context, req *user.AddUserRequest) (r *user.AddUserResponse, err error) {
	var _args user.UserServiceAddUserArgs
	_args.Req = req
	var _result user.UserServiceAddUserResult
	if err = p.c.Call(ctx, "AddUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (r *user.DeleteUserResponse, err error) {
	var _args user.UserServiceDeleteUserArgs
	_args.Req = req
	var _result user.UserServiceDeleteUserResult
	if err = p.c.Call(ctx, "DeleteUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
