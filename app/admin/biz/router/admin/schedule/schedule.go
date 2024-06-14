// Code generated by hertz generator. DO NOT EDIT.

package schedule

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	schedule "saas/app/admin/biz/handler/admin/schedule"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_admin := _api.Group("/admin", _adminMw()...)
			{
				_schedule := _admin.Group("/schedule", _scheduleMw()...)
				_schedule.POST("/create", append(_createscheduleMw(), schedule.CreateSchedule)...)
				_schedule.POST("/date-list", append(_datelistscheduleMw(), schedule.DateListSchedule)...)
				_schedule.GET("/info", append(_getschedulebyidMw(), schedule.GetScheduleById)...)
				_schedule.POST("/list", append(_listscheduleMw(), schedule.ListSchedule)...)
				_schedule.POST("/schedule-member-list", append(_getschedulememberlistMw(), schedule.GetScheduleMemberList)...)
				_schedule.POST("/schedule-member-subscribe", append(_schedulemembersubscribeMw(), schedule.ScheduleMemberSubscribe)...)
				_schedule.POST("/status", append(_updatestatusMw(), schedule.UpdateStatus)...)
				_schedule.POST("/update", append(_updatescheduleMw(), schedule.UpdateSchedule)...)
			}
		}
	}
}
