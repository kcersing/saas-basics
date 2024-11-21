package main

import (
	"context"
	schedule "rpc_gen/kitex_gen/admin/schedule"
	base "rpc_gen/kitex_gen/base"
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

// ListSchedule implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) ListSchedule(ctx context.Context, req *schedule.ListScheduleReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewListScheduleService(ctx).Run(req)

	return resp, err
}

// DateListSchedule implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) DateListSchedule(ctx context.Context, req *schedule.ListScheduleReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDateListScheduleService(ctx).Run(req)

	return resp, err
}

// GetScheduleById implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) GetScheduleById(ctx context.Context, req *base.IDReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewGetScheduleByIdService(ctx).Run(req)

	return resp, err
}

// GetScheduleMemberList implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) GetScheduleMemberList(ctx context.Context, req *schedule.ScheduleMemberReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewGetScheduleMemberListService(ctx).Run(req)

	return resp, err
}

// SearchSubscribeByMember implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) SearchSubscribeByMember(ctx context.Context, req *schedule.SearchSubscribeByMemberReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewSearchSubscribeByMemberService(ctx).Run(req)

	return resp, err
}

// MemberSubscribe implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) MemberSubscribe(ctx context.Context, req *schedule.MemberSubscribeReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewMemberSubscribeService(ctx).Run(req)

	return resp, err
}

// UpdateMemberStatus implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) UpdateMemberStatus(ctx context.Context, req *base.StatusCodeReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateMemberStatusService(ctx).Run(req)

	return resp, err
}

// GetScheduleCoachList implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) GetScheduleCoachList(ctx context.Context, req *schedule.ScheduleMemberReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewGetScheduleCoachListService(ctx).Run(req)

	return resp, err
}

// UpdateCoachStatus implements the ScheduleServiceImpl interface.
func (s *ScheduleServiceImpl) UpdateCoachStatus(ctx context.Context, req *base.StatusCodeReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateCoachStatusService(ctx).Run(req)

	return resp, err
}
