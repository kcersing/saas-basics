package schedule

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	schedule "rpc_gen/kitex_gen/schedule"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"rpc_gen/kitex_gen/schedule/scheduleservice"
)

type RPCClient interface {
	KitexClient() scheduleservice.Client
	Service() string
	CreateSchedule(ctx context.Context, req *schedule.CreateOrUpdateScheduleReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateSchedule(ctx context.Context, req *schedule.CreateOrUpdateScheduleReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	ScheduleList(ctx context.Context, req *schedule.ScheduleListReq, callOptions ...callopt.Option) (r *schedule.ScheduleListResp, err error)
	ScheduleDateList(ctx context.Context, req *schedule.ScheduleListReq, callOptions ...callopt.Option) (r *schedule.ScheduleDateListResp, err error)
	ScheduleInfo(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *schedule.ScheduleInfo, err error)
	MemberList(ctx context.Context, req *schedule.ScheduleMemberListReq, callOptions ...callopt.Option) (r *schedule.ScheduleMemberListResp, err error)
	CreateMember(ctx context.Context, req *schedule.CreateMemberReq, callOptions ...callopt.Option) (r *schedule.ScheduleMemberListResp, err error)
	UpdateMemberStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	SearchSubscribeByMember(ctx context.Context, req *schedule.SearchSubscribeByMemberReq, callOptions ...callopt.Option) (r *schedule.SearchSubscribeByMemberResp, err error)
	CoachList(ctx context.Context, req *schedule.ScheduleCoachListReq, callOptions ...callopt.Option) (r *schedule.ScheduleCoachListResp, err error)
	UpdateCoachStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := scheduleservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient scheduleservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() scheduleservice.Client {
	return c.kitexClient
}

func (c *clientImpl) CreateSchedule(ctx context.Context, req *schedule.CreateOrUpdateScheduleReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.CreateSchedule(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateSchedule(ctx context.Context, req *schedule.CreateOrUpdateScheduleReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateSchedule(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateStatus(ctx, req, callOptions...)
}

func (c *clientImpl) ScheduleList(ctx context.Context, req *schedule.ScheduleListReq, callOptions ...callopt.Option) (r *schedule.ScheduleListResp, err error) {
	return c.kitexClient.ScheduleList(ctx, req, callOptions...)
}

func (c *clientImpl) ScheduleDateList(ctx context.Context, req *schedule.ScheduleListReq, callOptions ...callopt.Option) (r *schedule.ScheduleDateListResp, err error) {
	return c.kitexClient.ScheduleDateList(ctx, req, callOptions...)
}

func (c *clientImpl) ScheduleInfo(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *schedule.ScheduleInfo, err error) {
	return c.kitexClient.ScheduleInfo(ctx, req, callOptions...)
}

func (c *clientImpl) MemberList(ctx context.Context, req *schedule.ScheduleMemberListReq, callOptions ...callopt.Option) (r *schedule.ScheduleMemberListResp, err error) {
	return c.kitexClient.MemberList(ctx, req, callOptions...)
}

func (c *clientImpl) CreateMember(ctx context.Context, req *schedule.CreateMemberReq, callOptions ...callopt.Option) (r *schedule.ScheduleMemberListResp, err error) {
	return c.kitexClient.CreateMember(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateMemberStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateMemberStatus(ctx, req, callOptions...)
}

func (c *clientImpl) SearchSubscribeByMember(ctx context.Context, req *schedule.SearchSubscribeByMemberReq, callOptions ...callopt.Option) (r *schedule.SearchSubscribeByMemberResp, err error) {
	return c.kitexClient.SearchSubscribeByMember(ctx, req, callOptions...)
}

func (c *clientImpl) CoachList(ctx context.Context, req *schedule.ScheduleCoachListReq, callOptions ...callopt.Option) (r *schedule.ScheduleCoachListResp, err error) {
	return c.kitexClient.CoachList(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateCoachStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateCoachStatus(ctx, req, callOptions...)
}
