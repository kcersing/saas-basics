package service

// This file is auto-generated, don't edit it. Thanks.

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/biz/dal/db/ent"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/dal/db/ent/venuesms"
	"saas/biz/dal/db/ent/venuesmslog"
	"saas/idl_gen/model/sms"
	"time"
)

type Sms struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (s Sms) Buy(req sms.SmsBuyReq) (resp string, err error) {
	//TODO implement me
	return
}

func (s Sms) Info(id int64) (resp *sms.SmsInfo, err error) {

	first, err := s.db.VenueSms.Query().Where(venuesms.IDEQ(id)).First(s.ctx)
	if first != nil {
		resp = &sms.SmsInfo{
			NoticeCount: first.NoticeCount,
			UsedNotice:  first.UsedNotice,
		}
	}
	return
}

func (s Sms) SendList(req sms.SmsSendListReq) (resp []*sms.SmsSend, total int, err error) {
	var predicates []predicate.VenueSmsLog

	if req.VenueId == 0 {
		err = errors.Wrap(err, "场馆ID不能为空")
		return resp, total, err
	}
	if req.Mobile != "" {
		predicates = append(predicates, venuesmslog.MobileEQ(req.Mobile))
	}
	predicates = append(predicates, venuesmslog.VenueID(req.VenueId))
	all, err := s.db.VenueSmsLog.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Venue Sms log list failed")
		return resp, total, err
	}

	for _, l := range all {
		log := sms.SmsSend{
			CreatedAt:  l.CreatedAt.Format(time.DateTime),
			Status:     l.Status,
			Mobile:     l.Mobile,
			Code:       l.Code,
			VenueId:    l.VenueID,
			BizId:      l.BizID,
			NotifyType: l.NotifyType,
			Content:    l.Content,
			Templates:  l.Template,
		}
		resp = append(resp, &log)
	}

	total, _ = s.db.VenueSmsLog.Query().Where(predicates...).Count(s.ctx)
	return
}
