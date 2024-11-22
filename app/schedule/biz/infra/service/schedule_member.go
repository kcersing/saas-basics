package service

import (
	"common/utils/minio"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"rpc_gen/kitex_gen/schedule"
	"schedule/biz/dal/mysql/ent/predicate"
	schedule2 "schedule/biz/dal/mysql/ent/schedule"
	"schedule/biz/dal/mysql/ent/schedulemember"
	"time"
)

func (c Schedule) CreateMember(req schedule.CreateMemberReq) error {

	sc, err := c.db.Schedule.Query().
		Where(schedule2.ID(*req.Schedule)).
		First(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "未查询到该课程")
		return err
	}

	if (sc.NumSurplus - int64(len(req.MemberProductPropertyId))) < 0 {
		err = errors.Wrap(errors.New("超出约课人数"+fmt.Sprintf("%+v", sc.NumSurplus-int64(len(req.MemberProductPropertyId)))), "未查询到该课程")
		return err
	}
	for _, v := range req.MemberProductPropertyId {

		mpp, err := c.db.MemberProductProperty.Query().
			Where(memberproductproperty.ID(v)).
			First(c.ctx)

		if err != nil {
			err = errors.Wrap(err, "未查询到该会员属性")
			return err
		}
		m, err := c.db.Member.Query().
			Where(member.ID(mpp.MemberID)).
			First(c.ctx)
		if err != nil {
			err = errors.Wrap(err, "未查询到该会员")
			return err
		}
		mp, err := c.db.MemberProduct.Query().
			Where(memberproduct.ID(mpp.MemberProductID)).
			First(c.ctx)
		if err != nil {
			err = errors.Wrap(err, "未查询到该会员产品")
			return err
		}

		_, err = c.db.ScheduleMember.Create().
			SetSchedule(sc).
			SetScheduleName(sc.Name).
			SetMemberName(m.Name).
			SetStatus(1).
			SetStartTime(sc.StartTime).
			SetEndTime(sc.EndTime).
			SetMemberID(m.ID).
			SetType(sc.Type).
			SetVenueID(sc.VenueID).
			SetMemberProductID(mp.ID).
			SetMemberProductName(mp.Name).
			SetMemberProductPropertyID(mpp.ID).
			SetMemberProductPropertyName(mpp.Name).
			Save(c.ctx)
		if err != nil {
			return err
		}

	}

	return nil

}

func (c Schedule) UpdateMember(req schedule.ScheduleMemberInfo) error {
	//TODO implement me
	panic("implement me")
}

func (c Schedule) DeleteMember(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (c Schedule) MemberList(req schedule.ScheduleMemberListReq) (resp []*schedule.ScheduleMemberInfo, total int64, err error) {

	var predicates []predicate.ScheduleMember
	if *req.Member > 0 {
		predicates = append(predicates, schedulemember.MemberID(*req.Member))
	}
	if *req.Schedule > 0 {
		predicates = append(predicates, schedulemember.ScheduleID(*req.Schedule))
	}

	lists, err := c.db.Debug().ScheduleMember.Query().Where(predicates...).
		Offset(int(*req.Page-1) * int(*req.PageSize)).
		Limit(int(*req.PageSize)).All(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Schedule Member list failed")
		return resp, total, err
	}

	err = copier.Copy(&resp, &lists)
	if err != nil {
		err = errors.Wrap(err, "copy Schedule Member info failed")
		return resp, 0, err
	}
	for i, v := range lists {
		vInfo, err := NewVenue(c.ctx, c.c).VenueInfo(v.VenueID)
		if err == nil {
			resp[i].VenueName = vInfo.Name
		}
		m, err := NewMember(c.ctx, c.c).Info(v.MemberID)
		if err == nil {
			resp[i].Mobile = m.Mobile
			resp[i].MemberName = m.Name
		}

		md, err := c.db.MemberDetails.Query().Where(memberdetails.MemberID(v.MemberID)).First(c.ctx)
		if err != nil {
			err = errors.Wrap(err, "未找到会员详情")
			return resp, 0, err
		} else {
			if md.Gender == 0 {
				resp[i].Gender = "女性"
			} else if md.Gender == 1 {
				resp[i].Gender = "男性"
			} else {
				resp[i].Gender = "保密"
			}
			if !md.Birthday.IsZero() {
				resp[i].Birthday = int64(time.Now().Sub(md.Birthday).Hours() / 24 / 365)

			}
		}
		resp[i].CreatedAt = v.CreatedAt.Format(time.DateTime)

	}

	total, _ = c.db.ScheduleMember.Query().Where(predicates...).Count(c.ctx)
	return

}

func (c Schedule) UpdateMemberStatus(ID int64, status int64) error {
	_, err := c.db.ScheduleMember.Update().Where(schedulemember.ID(ID)).SetStatus(status).Save(c.ctx)
	return err
}

func (c Schedule) MemberInfo(ID int64) (info *schedule.ScheduleMemberInfo, err error) {

	m, err := c.db.ScheduleMember.Query().Where(schedulemember.ID(ID)).First(c.ctx)

	if err != nil {
		err = errors.Wrap(err, "get Schedule Member failed")
		return info, err
	}

	err = copier.Copy(&info, &m)
	if err != nil {
		err = errors.Wrap(err, "copy Schedule Member info failed")
		return info, err
	}

	return info, nil

}

func (s Schedule) SearchSubscribeByMember(req schedule.SearchSubscribeByMemberReq) (list []*schedule.SubscribeByMember, total int64, err error) {

	m, err := s.db.Member.Query().Where(member.Mobile(req.Mobile)).First(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, 0, err
	}

	mpp, err := s.db.MemberProductProperty.
		Query().
		Where(
			memberproductproperty.PropertyID(req.PropertyId),
			memberproductproperty.HasVenuesWith(venue.ID(req.Venue)),
		).
		All(s.ctx)

	if err != nil {
		return nil, 0, err
	}
	for _, v := range mpp {
		mp, _ := v.QueryOwner().First(s.ctx)
		list = append(list, &schedule.SubscribeByMember{
			Avatar:                    minio.URLconvert(s.ctx, s.c, m.Avatar),
			Mobile:                    m.Mobile,
			MemberID:                  m.ID,
			MemberName:                m.Name,
			MemberProductID:           v.MemberProductID,
			MemberProductPropertyId:   v.ID,
			MemberProductName:         mp.Name,
			MemberProductPropertyName: v.Name,
		})
	}

	total = int64(len(list))
	return
}
