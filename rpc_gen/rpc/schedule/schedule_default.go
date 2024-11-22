package schedule

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	base "rpc_gen/kitex_gen/base"
	schedule "rpc_gen/kitex_gen/schedule"
)

func CreateSchedule(ctx context.Context, req *schedule.CreateOrUpdateScheduleReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.CreateSchedule(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateSchedule call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateSchedule(ctx context.Context, req *schedule.CreateOrUpdateScheduleReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateSchedule(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateSchedule call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateStatus(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateStatus call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ScheduleList(ctx context.Context, req *schedule.ScheduleListReq, callOptions ...callopt.Option) (resp *schedule.ScheduleListResp, err error) {
	resp, err = defaultClient.ScheduleList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ScheduleList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ScheduleDateList(ctx context.Context, req *schedule.ScheduleListReq, callOptions ...callopt.Option) (resp *schedule.ScheduleDateListResp, err error) {
	resp, err = defaultClient.ScheduleDateList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ScheduleDateList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ScheduleInfo(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp *schedule.ScheduleInfo, err error) {
	resp, err = defaultClient.ScheduleInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ScheduleInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MemberList(ctx context.Context, req *schedule.ScheduleMemberListReq, callOptions ...callopt.Option) (resp *schedule.ScheduleMemberListResp, err error) {
	resp, err = defaultClient.MemberList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MemberList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CreateMember(ctx context.Context, req *schedule.CreateMemberReq, callOptions ...callopt.Option) (resp *schedule.ScheduleMemberListResp, err error) {
	resp, err = defaultClient.CreateMember(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateMember call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateMemberStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateMemberStatus(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateMemberStatus call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SearchSubscribeByMember(ctx context.Context, req *schedule.SearchSubscribeByMemberReq, callOptions ...callopt.Option) (resp *schedule.SearchSubscribeByMemberResp, err error) {
	resp, err = defaultClient.SearchSubscribeByMember(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SearchSubscribeByMember call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CoachList(ctx context.Context, req *schedule.ScheduleCoachListReq, callOptions ...callopt.Option) (resp *schedule.ScheduleCoachListResp, err error) {
	resp, err = defaultClient.CoachList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CoachList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateCoachStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateCoachStatus(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateCoachStatus call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
