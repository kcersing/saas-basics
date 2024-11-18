// Code generated by Kitex v0.9.1. DO NOT EDIT.

package userservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	base "rpc_gen/kitex_gen/base"
	user "rpc_gen/kitex_gen/user"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Register": kitex.NewMethodInfo(
		registerHandler,
		newUserServiceRegisterArgs,
		newUserServiceRegisterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UserPermCode": kitex.NewMethodInfo(
		userPermCodeHandler,
		newUserServiceUserPermCodeArgs,
		newUserServiceUserPermCodeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ChangePassword": kitex.NewMethodInfo(
		changePasswordHandler,
		newUserServiceChangePasswordArgs,
		newUserServiceChangePasswordResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CreateUser": kitex.NewMethodInfo(
		createUserHandler,
		newUserServiceCreateUserArgs,
		newUserServiceCreateUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateUser": kitex.NewMethodInfo(
		updateUserHandler,
		newUserServiceUpdateUserArgs,
		newUserServiceUpdateUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UserInfo": kitex.NewMethodInfo(
		userInfoHandler,
		newUserServiceUserInfoArgs,
		newUserServiceUserInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UserList": kitex.NewMethodInfo(
		userListHandler,
		newUserServiceUserListArgs,
		newUserServiceUserListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeleteUser": kitex.NewMethodInfo(
		deleteUserHandler,
		newUserServiceDeleteUserArgs,
		newUserServiceDeleteUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateProfile": kitex.NewMethodInfo(
		updateProfileHandler,
		newUserServiceUpdateProfileArgs,
		newUserServiceUpdateProfileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UserProfile": kitex.NewMethodInfo(
		userProfileHandler,
		newUserServiceUserProfileArgs,
		newUserServiceUserProfileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateUserStatus": kitex.NewMethodInfo(
		updateUserStatusHandler,
		newUserServiceUpdateUserStatusArgs,
		newUserServiceUpdateUserStatusResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"SetUserRole": kitex.NewMethodInfo(
		setUserRoleHandler,
		newUserServiceSetUserRoleArgs,
		newUserServiceSetUserRoleResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"SetDefaultVenue": kitex.NewMethodInfo(
		setDefaultVenueHandler,
		newUserServiceSetDefaultVenueArgs,
		newUserServiceSetDefaultVenueResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	userServiceServiceInfo                = NewServiceInfo()
	userServiceServiceInfoForClient       = NewServiceInfoForClient()
	userServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceRegisterArgs)
	realResult := result.(*user.UserServiceRegisterResult)
	success, err := handler.(user.UserService).Register(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceRegisterArgs() interface{} {
	return user.NewUserServiceRegisterArgs()
}

func newUserServiceRegisterResult() interface{} {
	return user.NewUserServiceRegisterResult()
}

func userPermCodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserPermCodeArgs)
	realResult := result.(*user.UserServiceUserPermCodeResult)
	success, err := handler.(user.UserService).UserPermCode(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserPermCodeArgs() interface{} {
	return user.NewUserServiceUserPermCodeArgs()
}

func newUserServiceUserPermCodeResult() interface{} {
	return user.NewUserServiceUserPermCodeResult()
}

func changePasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceChangePasswordArgs)
	realResult := result.(*user.UserServiceChangePasswordResult)
	success, err := handler.(user.UserService).ChangePassword(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceChangePasswordArgs() interface{} {
	return user.NewUserServiceChangePasswordArgs()
}

func newUserServiceChangePasswordResult() interface{} {
	return user.NewUserServiceChangePasswordResult()
}

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceCreateUserArgs)
	realResult := result.(*user.UserServiceCreateUserResult)
	success, err := handler.(user.UserService).CreateUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceCreateUserArgs() interface{} {
	return user.NewUserServiceCreateUserArgs()
}

func newUserServiceCreateUserResult() interface{} {
	return user.NewUserServiceCreateUserResult()
}

func updateUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdateUserArgs)
	realResult := result.(*user.UserServiceUpdateUserResult)
	success, err := handler.(user.UserService).UpdateUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUpdateUserArgs() interface{} {
	return user.NewUserServiceUpdateUserArgs()
}

func newUserServiceUpdateUserResult() interface{} {
	return user.NewUserServiceUpdateUserResult()
}

func userInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserInfoArgs)
	realResult := result.(*user.UserServiceUserInfoResult)
	success, err := handler.(user.UserService).UserInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserInfoArgs() interface{} {
	return user.NewUserServiceUserInfoArgs()
}

func newUserServiceUserInfoResult() interface{} {
	return user.NewUserServiceUserInfoResult()
}

func userListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserListArgs)
	realResult := result.(*user.UserServiceUserListResult)
	success, err := handler.(user.UserService).UserList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserListArgs() interface{} {
	return user.NewUserServiceUserListArgs()
}

func newUserServiceUserListResult() interface{} {
	return user.NewUserServiceUserListResult()
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

func updateProfileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdateProfileArgs)
	realResult := result.(*user.UserServiceUpdateProfileResult)
	success, err := handler.(user.UserService).UpdateProfile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUpdateProfileArgs() interface{} {
	return user.NewUserServiceUpdateProfileArgs()
}

func newUserServiceUpdateProfileResult() interface{} {
	return user.NewUserServiceUpdateProfileResult()
}

func userProfileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserProfileArgs)
	realResult := result.(*user.UserServiceUserProfileResult)
	success, err := handler.(user.UserService).UserProfile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserProfileArgs() interface{} {
	return user.NewUserServiceUserProfileArgs()
}

func newUserServiceUserProfileResult() interface{} {
	return user.NewUserServiceUserProfileResult()
}

func updateUserStatusHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdateUserStatusArgs)
	realResult := result.(*user.UserServiceUpdateUserStatusResult)
	success, err := handler.(user.UserService).UpdateUserStatus(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUpdateUserStatusArgs() interface{} {
	return user.NewUserServiceUpdateUserStatusArgs()
}

func newUserServiceUpdateUserStatusResult() interface{} {
	return user.NewUserServiceUpdateUserStatusResult()
}

func setUserRoleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceSetUserRoleArgs)
	realResult := result.(*user.UserServiceSetUserRoleResult)
	success, err := handler.(user.UserService).SetUserRole(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceSetUserRoleArgs() interface{} {
	return user.NewUserServiceSetUserRoleArgs()
}

func newUserServiceSetUserRoleResult() interface{} {
	return user.NewUserServiceSetUserRoleResult()
}

func setDefaultVenueHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceSetDefaultVenueArgs)
	realResult := result.(*user.UserServiceSetDefaultVenueResult)
	success, err := handler.(user.UserService).SetDefaultVenue(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceSetDefaultVenueArgs() interface{} {
	return user.NewUserServiceSetDefaultVenueArgs()
}

func newUserServiceSetDefaultVenueResult() interface{} {
	return user.NewUserServiceSetDefaultVenueResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, req *user.RegisterReq) (r *base.NilResponse, err error) {
	var _args user.UserServiceRegisterArgs
	_args.Req = req
	var _result user.UserServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserPermCode(ctx context.Context, req *base.Empty) (r *base.NilResponse, err error) {
	var _args user.UserServiceUserPermCodeArgs
	_args.Req = req
	var _result user.UserServiceUserPermCodeResult
	if err = p.c.Call(ctx, "UserPermCode", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChangePassword(ctx context.Context, req *user.ChangePasswordReq) (r *base.NilResponse, err error) {
	var _args user.UserServiceChangePasswordArgs
	_args.Req = req
	var _result user.UserServiceChangePasswordResult
	if err = p.c.Call(ctx, "ChangePassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateUser(ctx context.Context, req *user.CreateOrUpdateUserReq) (r *base.NilResponse, err error) {
	var _args user.UserServiceCreateUserArgs
	_args.Req = req
	var _result user.UserServiceCreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUser(ctx context.Context, req *user.CreateOrUpdateUserReq) (r *base.NilResponse, err error) {
	var _args user.UserServiceUpdateUserArgs
	_args.Req = req
	var _result user.UserServiceUpdateUserResult
	if err = p.c.Call(ctx, "UpdateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserInfo(ctx context.Context, req *base.Empty) (r *base.NilResponse, err error) {
	var _args user.UserServiceUserInfoArgs
	_args.Req = req
	var _result user.UserServiceUserInfoResult
	if err = p.c.Call(ctx, "UserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserList(ctx context.Context, req *user.UserListReq) (r *base.NilResponse, err error) {
	var _args user.UserServiceUserListArgs
	_args.Req = req
	var _result user.UserServiceUserListResult
	if err = p.c.Call(ctx, "UserList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteUser(ctx context.Context, req *base.IDReq) (r *base.NilResponse, err error) {
	var _args user.UserServiceDeleteUserArgs
	_args.Req = req
	var _result user.UserServiceDeleteUserResult
	if err = p.c.Call(ctx, "DeleteUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateProfile(ctx context.Context, req *user.ProfileReq) (r *base.NilResponse, err error) {
	var _args user.UserServiceUpdateProfileArgs
	_args.Req = req
	var _result user.UserServiceUpdateProfileResult
	if err = p.c.Call(ctx, "UpdateProfile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserProfile(ctx context.Context, req *base.Empty) (r *base.NilResponse, err error) {
	var _args user.UserServiceUserProfileArgs
	_args.Req = req
	var _result user.UserServiceUserProfileResult
	if err = p.c.Call(ctx, "UserProfile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUserStatus(ctx context.Context, req *base.StatusCodeReq) (r *base.NilResponse, err error) {
	var _args user.UserServiceUpdateUserStatusArgs
	_args.Req = req
	var _result user.UserServiceUpdateUserStatusResult
	if err = p.c.Call(ctx, "UpdateUserStatus", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SetUserRole(ctx context.Context, req *user.SetUserRole) (r *base.NilResponse, err error) {
	var _args user.UserServiceSetUserRoleArgs
	_args.Req = req
	var _result user.UserServiceSetUserRoleResult
	if err = p.c.Call(ctx, "SetUserRole", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SetDefaultVenue(ctx context.Context, req *user.SetDefaultVenueReq) (r *base.NilResponse, err error) {
	var _args user.UserServiceSetDefaultVenueArgs
	_args.Req = req
	var _result user.UserServiceSetDefaultVenueResult
	if err = p.c.Call(ctx, "SetDefaultVenue", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
