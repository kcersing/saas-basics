package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	memberProduct "saas/idl_gen/model/member/memberProduct"
	"time"

	"saas/biz/infras/do"
	"saas/config"
)

type MemberProduct struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (m MemberProduct) MemberProductInfo(id int64) (info *memberProduct.MemberProductInfo, err error) {
	//TODO implement me
	panic("implement me")
}
func (m MemberProduct) entMemberProductInfo(req *ent.MemberProduct) (info *memberProduct.MemberProductInfo) {

	var courses, lessons []*memberProduct.MemberProductCourses
	coursesAll, _ := req.QueryMemberCourses().All(m.ctx)
	if len(coursesAll) > 0 {
		for _, ve := range coursesAll {
			courses = append(courses, &memberProduct.MemberProductCourses{
				ID:              ve.ID,
				Type:            ve.Type,
				Name:            ve.Name,
				Number:          ve.Number,
				CoursesId:       ve.CoursesID,
				MemberProductId: ve.MemberProductID,
			},
			)

		}
	}
	lessonsAll, _ := req.QueryMemberLessons().All(m.ctx)
	if len(lessonsAll) > 0 {
		for _, ve := range lessonsAll {
			lessons = append(lessons, &memberProduct.MemberProductCourses{
				ID:              ve.ID,
				Type:            ve.Type,
				Name:            ve.Name,
				Number:          ve.Number,
				CoursesId:       ve.CoursesID,
				MemberProductId: ve.MemberProductID,
			},
			)

		}
	}

	return &memberProduct.MemberProductInfo{
		ID:           req.ID,
		Name:         req.Name,
		Price:        req.Price,
		Fee:          req.Fee,
		Status:       req.Status,
		Duration:     req.Duration,
		Length:       req.Length,
		Type:         req.Type,
		Deadline:     req.Deadline,
		Count:        req.Count,
		CountSurplus: req.CountSurplus,
		ValidityAt:   req.ValidityAt.Format(time.DateTime),
		CancelAt:     req.CancelAt.Format(time.DateTime),
		CreatedAt:    req.CreatedAt.Format(time.DateTime),
		UpdatedAt:    req.UpdatedAt.Format(time.DateTime),
		OrderId:      req.OrderID,
		VenueId:      req.VenueID,
		ProductId:    req.ProductID,
		MemberId:     req.MemberID,
		SubType:      req.SubType,
		Sn:           req.Sn,
		Courses:      courses,
		Lessons:      lessons,
	}
}

func (m MemberProduct) MemberProductList(req memberProduct.MemberProductListReq) (resp []*memberProduct.MemberProductInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func NewMemberProduct(ctx context.Context, c *app.RequestContext) do.MemberProduct {
	return &MemberProduct{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
