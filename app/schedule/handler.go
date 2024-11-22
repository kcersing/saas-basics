package main

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	schedule "rpc_gen/kitex_gen/schedule"
	"schedule/biz/service"
)

// ScheduleServiceImpl implements the last service interface defined in the IDL.
type ScheduleServiceImpl struct{}

// CreateSchedule implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) CreateSchedule(ctx context.Context, req *schedule.CreateOrUpdateScheduleReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewCreateScheduleService(ctx).Run(req)

	return resp, err
}

// UpdateSchedule implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) UpdateSchedule(ctx context.Context, req *schedule.CreateOrUpdateScheduleReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateScheduleService(ctx).Run(req)

	return resp, err
}

// UpdateStatus implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) UpdateStatus(ctx context.Context, req *base.StatusCodeReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateStatusService(ctx).Run(req)

	return resp, err
}

// ScheduleListResp implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) ScheduleListResp(ctx context.Context, req *schedule.ScheduleListReq) (resp *schedule.ScheduleListResp, err error) {
	resp, err = service.NewScheduleListRespService(ctx).Run(req)

	return resp, err
}

// ScheduleDateList implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) ScheduleDateList(ctx context.Context, req *schedule.ScheduleListReq) (resp *schedule.ScheduleListResp, err error) {
	resp, err = service.NewScheduleDateListService(ctx).Run(req)

	return resp, err
}

// ScheduleInfo implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) ScheduleInfo(ctx context.Context, req *base.IDReq) (resp *schedule.ScheduleInfo, err error) {
	resp, err = service.NewScheduleInfoService(ctx).Run(req)

	return resp, err
}

// MemberList implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) MemberList(ctx context.Context, req *schedule.ScheduleMemberListReq) (resp *schedule.ScheduleMemberListResp, err error) {
	resp, err = service.NewMemberListService(ctx).Run(req)

	return resp, err
}

// CreateMember implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) CreateMember(ctx context.Context, req *schedule.CreateMemberReq) (resp *schedule.ScheduleMemberListResp, err error) {
	resp, err = service.NewCreateMemberService(ctx).Run(req)

	return resp, err
}

// UpdateMemberStatus implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) UpdateMemberStatus(ctx context.Context, req *base.StatusCodeReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateMemberStatusService(ctx).Run(req)

	return resp, err
}

// SearchSubscribeByMember implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) SearchSubscribeByMember(ctx context.Context, req *schedule.SearchSubscribeByMemberReq) (resp *schedule.ScheduleMemberInfo, err error) {
	resp, err = service.NewSearchSubscribeByMemberService(ctx).Run(req)

	return resp, err
}

// CoachList implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) CoachList(ctx context.Context, req *schedule.ScheduleCoachListReq) (resp *schedule.ScheduleCoachListResp, err error) {
	resp, err = service.NewCoachListService(ctx).Run(req)

	return resp, err
}

// UpdateCoachStatus implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) UpdateCoachStatus(ctx context.Context, req *base.StatusCodeReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateCoachStatusService(ctx).Run(req)

	return resp, err
}

// ScheduleList implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) ScheduleList(ctx context.Context, req *schedule.ScheduleListReq) (resp *schedule.ScheduleListResp, err error) {
	resp, err = service.NewScheduleListService(ctx).Run(req)

	return resp, err
}
