package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	member2 "saas/biz/dal/db/ent/member"
	memberproduct2 "saas/biz/dal/db/ent/memberproduct"

	"saas/biz/dal/db/ent/predicate"
	memberProduct "saas/idl_gen/model/member/memberProduct"
	"saas/pkg/enums"
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
	first, err := m.db.MemberProduct.Query().Where(memberproduct2.IDEQ(id)).First(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get member product failed")
		return nil, err
	}

	info = m.entMemberProductInfo(first)

	return
}
func (m MemberProduct) entMemberProductInfo(req *ent.MemberProduct) (info *memberProduct.MemberProductInfo) {
	var courses, lessons []*memberProduct.MemberProductCourses
	if req.Type == enums.Lessons {
		lessonsAll, _ := req.QueryMemberLessons().All(m.ctx)
		if len(lessonsAll) > 0 {
			for _, ve := range lessonsAll {
				lessons = append(lessons, &memberProduct.MemberProductCourses{
					ID:              ve.ID,
					Type:            ve.Type,
					Name:            ve.Name,
					CoursesId:       ve.CoursesID,
					MemberProductId: ve.MemberProductID,
				},
				)
			}
		}
	} else {
		if req.Type != enums.Card {
			coursesAll, _ := req.QueryMemberCourses().All(m.ctx)
			if len(coursesAll) > 0 {
				for _, ve := range coursesAll {
					courses = append(courses, &memberProduct.MemberProductCourses{
						ID:              ve.ID,
						Type:            ve.Type,
						Name:            ve.Name,
						Number:          req.Number,
						NumberSurplus:   req.NumberSurplus,
						CoursesId:       ve.CoursesID,
						MemberProductId: ve.MemberProductID,
					},
					)
				}
			}
		}
	}

	return &memberProduct.MemberProductInfo{
		ID:            req.ID,
		Name:          req.Name,
		Price:         req.Price,
		Fee:           req.Fee,
		Status:        req.Status,
		Duration:      req.Duration,
		Length:        req.Length,
		Type:          req.Type,
		Deadline:      req.Deadline,
		Number:        req.Number,
		NumberSurplus: req.NumberSurplus,
		ValidityAt:    req.ValidityAt.Format(time.DateTime),
		CancelAt:      req.CancelAt.Format(time.DateTime),
		CreatedAt:     req.CreatedAt.Format(time.DateTime),
		UpdatedAt:     req.UpdatedAt.Format(time.DateTime),
		OrderId:       req.OrderID,
		VenueId:       req.VenueID,
		ProductId:     req.ProductID,
		MemberId:      req.MemberID,
		SubType:       req.SubType,
		Sn:            req.Sn,
		Courses:       courses,
		Lessons:       lessons,
		IsCourse:      req.IsCourse,
	}
}

func (m MemberProduct) MemberProductList(req memberProduct.MemberProductListReq) (resp []*memberProduct.MemberProductInfo, total int, err error) {
	var predicates []predicate.MemberProduct

	if req.MemberId == 0 {
		return resp, total, errors.New("会员ID不能为空")
	}

	if req.Name != "" {
		predicates = append(predicates, memberproduct2.NameEQ(req.Name))
	}

	if len(req.Status) > 0 {
		predicates = append(predicates, memberproduct2.StatusIn(req.Status...))
	}
	if req.Type != "" {
		predicates = append(predicates, memberproduct2.TypeEQ(req.Type))
	}
	if req.SubType != "" {
		predicates = append(predicates, memberproduct2.SubTypeEQ(req.SubType))
	}
	if req.VenueId > 0 {
		predicates = append(predicates, memberproduct2.VenueIDEQ(req.VenueId))
	}
	predicates = append(predicates, memberproduct2.MemberIDEQ(req.MemberId))
	predicates = append(predicates, memberproduct2.Delete(0))
	predicates = append(predicates, memberproduct2.StatusNEQ(0))

	lists, err := m.db.MemberProduct.Query().Where(predicates...).
		Order(ent.Desc(member2.FieldID)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Member Product list failed")
		return resp, total, err
	}

	for _, ve := range lists {
		resp = append(resp, m.entMemberProductInfo(ve))
	}

	total, _ = m.db.MemberProduct.Query().Where(predicates...).Count(m.ctx)
	return
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
