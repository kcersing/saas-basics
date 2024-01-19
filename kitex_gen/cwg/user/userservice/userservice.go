// Code generated by Kitex v0.8.0. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	user "saas_basics/kitex_gen/cwg/user"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"AdminLogin": kitex.NewMethodInfo(adminLoginHandler, newUserServiceAdminLoginArgs, newUserServiceAdminLoginResult, false),
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

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
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
