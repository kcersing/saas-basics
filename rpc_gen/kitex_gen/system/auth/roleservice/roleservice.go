// Code generated by Kitex v0.9.1. DO NOT EDIT.

package roleservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CreateRole": kitex.NewMethodInfo(
		createRoleHandler,
		newRoleServiceCreateRoleArgs,
		newRoleServiceCreateRoleResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateRole": kitex.NewMethodInfo(
		updateRoleHandler,
		newRoleServiceUpdateRoleArgs,
		newRoleServiceUpdateRoleResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeleteRole": kitex.NewMethodInfo(
		deleteRoleHandler,
		newRoleServiceDeleteRoleArgs,
		newRoleServiceDeleteRoleResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"RoleByID": kitex.NewMethodInfo(
		roleByIDHandler,
		newRoleServiceRoleByIDArgs,
		newRoleServiceRoleByIDResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"RoleList": kitex.NewMethodInfo(
		roleListHandler,
		newRoleServiceRoleListArgs,
		newRoleServiceRoleListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateRoleStatus": kitex.NewMethodInfo(
		updateRoleStatusHandler,
		newRoleServiceUpdateRoleStatusArgs,
		newRoleServiceUpdateRoleStatusResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CreateAuthority": kitex.NewMethodInfo(
		createAuthorityHandler,
		newRoleServiceCreateAuthorityArgs,
		newRoleServiceCreateAuthorityResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateApiAuthority": kitex.NewMethodInfo(
		updateApiAuthorityHandler,
		newRoleServiceUpdateApiAuthorityArgs,
		newRoleServiceUpdateApiAuthorityResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ApiAuthority": kitex.NewMethodInfo(
		apiAuthorityHandler,
		newRoleServiceApiAuthorityArgs,
		newRoleServiceApiAuthorityResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CreateMenuAuthority": kitex.NewMethodInfo(
		createMenuAuthorityHandler,
		newRoleServiceCreateMenuAuthorityArgs,
		newRoleServiceCreateMenuAuthorityResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateMenuAuthority": kitex.NewMethodInfo(
		updateMenuAuthorityHandler,
		newRoleServiceUpdateMenuAuthorityArgs,
		newRoleServiceUpdateMenuAuthorityResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"MenuAuthority": kitex.NewMethodInfo(
		menuAuthorityHandler,
		newRoleServiceMenuAuthorityArgs,
		newRoleServiceMenuAuthorityResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CreateApi": kitex.NewMethodInfo(
		createApiHandler,
		newRoleServiceCreateApiArgs,
		newRoleServiceCreateApiResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateApi": kitex.NewMethodInfo(
		updateApiHandler,
		newRoleServiceUpdateApiArgs,
		newRoleServiceUpdateApiResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeleteApi": kitex.NewMethodInfo(
		deleteApiHandler,
		newRoleServiceDeleteApiArgs,
		newRoleServiceDeleteApiResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ApiList": kitex.NewMethodInfo(
		apiListHandler,
		newRoleServiceApiListArgs,
		newRoleServiceApiListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ApiTree": kitex.NewMethodInfo(
		apiTreeHandler,
		newRoleServiceApiTreeArgs,
		newRoleServiceApiTreeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	roleServiceServiceInfo                = NewServiceInfo()
	roleServiceServiceInfoForClient       = NewServiceInfoForClient()
	roleServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return roleServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return roleServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return roleServiceServiceInfoForClient
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
	serviceName := "RoleService"
	handlerType := (*auth.RoleService)(nil)
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
		"PackageName": "auth",
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

func createRoleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.RoleServiceCreateRoleArgs)
	realResult := result.(*auth.RoleServiceCreateRoleResult)
	success, err := handler.(auth.RoleService).CreateRole(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoleServiceCreateRoleArgs() interface{} {
	return auth.NewRoleServiceCreateRoleArgs()
}

func newRoleServiceCreateRoleResult() interface{} {
	return auth.NewRoleServiceCreateRoleResult()
}

func updateRoleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.RoleServiceUpdateRoleArgs)
	realResult := result.(*auth.RoleServiceUpdateRoleResult)
	success, err := handler.(auth.RoleService).UpdateRole(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoleServiceUpdateRoleArgs() interface{} {
	return auth.NewRoleServiceUpdateRoleArgs()
}

func newRoleServiceUpdateRoleResult() interface{} {
	return auth.NewRoleServiceUpdateRoleResult()
}

func deleteRoleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.RoleServiceDeleteRoleArgs)
	realResult := result.(*auth.RoleServiceDeleteRoleResult)
	success, err := handler.(auth.RoleService).DeleteRole(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoleServiceDeleteRoleArgs() interface{} {
	return auth.NewRoleServiceDeleteRoleArgs()
}

func newRoleServiceDeleteRoleResult() interface{} {
	return auth.NewRoleServiceDeleteRoleResult()
}

func roleByIDHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.RoleServiceRoleByIDArgs)
	realResult := result.(*auth.RoleServiceRoleByIDResult)
	success, err := handler.(auth.RoleService).RoleByID(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoleServiceRoleByIDArgs() interface{} {
	return auth.NewRoleServiceRoleByIDArgs()
}

func newRoleServiceRoleByIDResult() interface{} {
	return auth.NewRoleServiceRoleByIDResult()
}

func roleListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.RoleServiceRoleListArgs)
	realResult := result.(*auth.RoleServiceRoleListResult)
	success, err := handler.(auth.RoleService).RoleList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoleServiceRoleListArgs() interface{} {
	return auth.NewRoleServiceRoleListArgs()
}

func newRoleServiceRoleListResult() interface{} {
	return auth.NewRoleServiceRoleListResult()
}

func updateRoleStatusHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.RoleServiceUpdateRoleStatusArgs)
	realResult := result.(*auth.RoleServiceUpdateRoleStatusResult)
	success, err := handler.(auth.RoleService).UpdateRoleStatus(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoleServiceUpdateRoleStatusArgs() interface{} {
	return auth.NewRoleServiceUpdateRoleStatusArgs()
}

func newRoleServiceUpdateRoleStatusResult() interface{} {
	return auth.NewRoleServiceUpdateRoleStatusResult()
}

func createAuthorityHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.RoleServiceCreateAuthorityArgs)
	realResult := result.(*auth.RoleServiceCreateAuthorityResult)
	success, err := handler.(auth.RoleService).CreateAuthority(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoleServiceCreateAuthorityArgs() interface{} {
	return auth.NewRoleServiceCreateAuthorityArgs()
}

func newRoleServiceCreateAuthorityResult() interface{} {
	return auth.NewRoleServiceCreateAuthorityResult()
}

func updateApiAuthorityHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.RoleServiceUpdateApiAuthorityArgs)
	realResult := result.(*auth.RoleServiceUpdateApiAuthorityResult)
	success, err := handler.(auth.RoleService).UpdateApiAuthority(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoleServiceUpdateApiAuthorityArgs() interface{} {
	return auth.NewRoleServiceUpdateApiAuthorityArgs()
}

func newRoleServiceUpdateApiAuthorityResult() interface{} {
	return auth.NewRoleServiceUpdateApiAuthorityResult()
}

func apiAuthorityHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.RoleServiceApiAuthorityArgs)
	realResult := result.(*auth.RoleServiceApiAuthorityResult)
	success, err := handler.(auth.RoleService).ApiAuthority(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoleServiceApiAuthorityArgs() interface{} {
	return auth.NewRoleServiceApiAuthorityArgs()
}

func newRoleServiceApiAuthorityResult() interface{} {
	return auth.NewRoleServiceApiAuthorityResult()
}

func createMenuAuthorityHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.RoleServiceCreateMenuAuthorityArgs)
	realResult := result.(*auth.RoleServiceCreateMenuAuthorityResult)
	success, err := handler.(auth.RoleService).CreateMenuAuthority(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoleServiceCreateMenuAuthorityArgs() interface{} {
	return auth.NewRoleServiceCreateMenuAuthorityArgs()
}

func newRoleServiceCreateMenuAuthorityResult() interface{} {
	return auth.NewRoleServiceCreateMenuAuthorityResult()
}

func updateMenuAuthorityHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.RoleServiceUpdateMenuAuthorityArgs)
	realResult := result.(*auth.RoleServiceUpdateMenuAuthorityResult)
	success, err := handler.(auth.RoleService).UpdateMenuAuthority(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoleServiceUpdateMenuAuthorityArgs() interface{} {
	return auth.NewRoleServiceUpdateMenuAuthorityArgs()
}

func newRoleServiceUpdateMenuAuthorityResult() interface{} {
	return auth.NewRoleServiceUpdateMenuAuthorityResult()
}

func menuAuthorityHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.RoleServiceMenuAuthorityArgs)
	realResult := result.(*auth.RoleServiceMenuAuthorityResult)
	success, err := handler.(auth.RoleService).MenuAuthority(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoleServiceMenuAuthorityArgs() interface{} {
	return auth.NewRoleServiceMenuAuthorityArgs()
}

func newRoleServiceMenuAuthorityResult() interface{} {
	return auth.NewRoleServiceMenuAuthorityResult()
}

func createApiHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.RoleServiceCreateApiArgs)
	realResult := result.(*auth.RoleServiceCreateApiResult)
	success, err := handler.(auth.RoleService).CreateApi(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoleServiceCreateApiArgs() interface{} {
	return auth.NewRoleServiceCreateApiArgs()
}

func newRoleServiceCreateApiResult() interface{} {
	return auth.NewRoleServiceCreateApiResult()
}

func updateApiHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.RoleServiceUpdateApiArgs)
	realResult := result.(*auth.RoleServiceUpdateApiResult)
	success, err := handler.(auth.RoleService).UpdateApi(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoleServiceUpdateApiArgs() interface{} {
	return auth.NewRoleServiceUpdateApiArgs()
}

func newRoleServiceUpdateApiResult() interface{} {
	return auth.NewRoleServiceUpdateApiResult()
}

func deleteApiHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.RoleServiceDeleteApiArgs)
	realResult := result.(*auth.RoleServiceDeleteApiResult)
	success, err := handler.(auth.RoleService).DeleteApi(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoleServiceDeleteApiArgs() interface{} {
	return auth.NewRoleServiceDeleteApiArgs()
}

func newRoleServiceDeleteApiResult() interface{} {
	return auth.NewRoleServiceDeleteApiResult()
}

func apiListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.RoleServiceApiListArgs)
	realResult := result.(*auth.RoleServiceApiListResult)
	success, err := handler.(auth.RoleService).ApiList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoleServiceApiListArgs() interface{} {
	return auth.NewRoleServiceApiListArgs()
}

func newRoleServiceApiListResult() interface{} {
	return auth.NewRoleServiceApiListResult()
}

func apiTreeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.RoleServiceApiTreeArgs)
	realResult := result.(*auth.RoleServiceApiTreeResult)
	success, err := handler.(auth.RoleService).ApiTree(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoleServiceApiTreeArgs() interface{} {
	return auth.NewRoleServiceApiTreeArgs()
}

func newRoleServiceApiTreeResult() interface{} {
	return auth.NewRoleServiceApiTreeResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateRole(ctx context.Context, req *auth.RoleInfo) (r *base.NilResponse, err error) {
	var _args auth.RoleServiceCreateRoleArgs
	_args.Req = req
	var _result auth.RoleServiceCreateRoleResult
	if err = p.c.Call(ctx, "CreateRole", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateRole(ctx context.Context, req *auth.RoleInfo) (r *base.NilResponse, err error) {
	var _args auth.RoleServiceUpdateRoleArgs
	_args.Req = req
	var _result auth.RoleServiceUpdateRoleResult
	if err = p.c.Call(ctx, "UpdateRole", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteRole(ctx context.Context, req *base.IDReq) (r *base.NilResponse, err error) {
	var _args auth.RoleServiceDeleteRoleArgs
	_args.Req = req
	var _result auth.RoleServiceDeleteRoleResult
	if err = p.c.Call(ctx, "DeleteRole", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RoleByID(ctx context.Context, req *base.IDReq) (r *auth.RoleInfo, err error) {
	var _args auth.RoleServiceRoleByIDArgs
	_args.Req = req
	var _result auth.RoleServiceRoleByIDResult
	if err = p.c.Call(ctx, "RoleByID", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RoleList(ctx context.Context, req *base.PageInfoReq) (r []*auth.RoleInfo, err error) {
	var _args auth.RoleServiceRoleListArgs
	_args.Req = req
	var _result auth.RoleServiceRoleListResult
	if err = p.c.Call(ctx, "RoleList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateRoleStatus(ctx context.Context, req *base.StatusCodeReq) (r *base.NilResponse, err error) {
	var _args auth.RoleServiceUpdateRoleStatusArgs
	_args.Req = req
	var _result auth.RoleServiceUpdateRoleStatusResult
	if err = p.c.Call(ctx, "UpdateRoleStatus", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateAuthority(ctx context.Context, req *auth.CreateOrUpdateApiAuthorityReq) (r *base.NilResponse, err error) {
	var _args auth.RoleServiceCreateAuthorityArgs
	_args.Req = req
	var _result auth.RoleServiceCreateAuthorityResult
	if err = p.c.Call(ctx, "CreateAuthority", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateApiAuthority(ctx context.Context, req *auth.CreateOrUpdateApiAuthorityReq) (r *base.NilResponse, err error) {
	var _args auth.RoleServiceUpdateApiAuthorityArgs
	_args.Req = req
	var _result auth.RoleServiceUpdateApiAuthorityResult
	if err = p.c.Call(ctx, "UpdateApiAuthority", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ApiAuthority(ctx context.Context, req *base.IDReq) (r *base.NilResponse, err error) {
	var _args auth.RoleServiceApiAuthorityArgs
	_args.Req = req
	var _result auth.RoleServiceApiAuthorityResult
	if err = p.c.Call(ctx, "ApiAuthority", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateMenuAuthority(ctx context.Context, req *auth.MenuAuthorityInfoReq) (r *base.NilResponse, err error) {
	var _args auth.RoleServiceCreateMenuAuthorityArgs
	_args.Req = req
	var _result auth.RoleServiceCreateMenuAuthorityResult
	if err = p.c.Call(ctx, "CreateMenuAuthority", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateMenuAuthority(ctx context.Context, req *auth.MenuAuthorityInfoReq) (r *base.NilResponse, err error) {
	var _args auth.RoleServiceUpdateMenuAuthorityArgs
	_args.Req = req
	var _result auth.RoleServiceUpdateMenuAuthorityResult
	if err = p.c.Call(ctx, "UpdateMenuAuthority", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MenuAuthority(ctx context.Context, req *base.IDReq) (r *base.NilResponse, err error) {
	var _args auth.RoleServiceMenuAuthorityArgs
	_args.Req = req
	var _result auth.RoleServiceMenuAuthorityResult
	if err = p.c.Call(ctx, "MenuAuthority", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateApi(ctx context.Context, req *auth.ApiInfo) (r *base.NilResponse, err error) {
	var _args auth.RoleServiceCreateApiArgs
	_args.Req = req
	var _result auth.RoleServiceCreateApiResult
	if err = p.c.Call(ctx, "CreateApi", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateApi(ctx context.Context, req *auth.ApiInfo) (r *base.NilResponse, err error) {
	var _args auth.RoleServiceUpdateApiArgs
	_args.Req = req
	var _result auth.RoleServiceUpdateApiResult
	if err = p.c.Call(ctx, "UpdateApi", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteApi(ctx context.Context, req *base.IDReq) (r *base.NilResponse, err error) {
	var _args auth.RoleServiceDeleteApiArgs
	_args.Req = req
	var _result auth.RoleServiceDeleteApiResult
	if err = p.c.Call(ctx, "DeleteApi", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ApiList(ctx context.Context, req *auth.ApiPageReq) (r []*auth.ApiInfo, err error) {
	var _args auth.RoleServiceApiListArgs
	_args.Req = req
	var _result auth.RoleServiceApiListResult
	if err = p.c.Call(ctx, "ApiList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ApiTree(ctx context.Context, req *auth.ApiPageReq) (r []*base.Tree, err error) {
	var _args auth.RoleServiceApiTreeArgs
	_args.Req = req
	var _result auth.RoleServiceApiTreeResult
	if err = p.c.Call(ctx, "ApiTree", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
