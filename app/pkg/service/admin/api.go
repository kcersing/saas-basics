package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/api"
	"saas/pkg/db/ent/predicate"
)

type Api struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (a Api) Create(req do.ApiInfo) error {
	_, err := a.db.API.Create().
		SetPath(req.Path).
		SetDescription(req.Description).
		SetAPIGroup(req.Group).
		SetMethod(req.Method).
		Save(a.ctx)
	if err != nil {
		err = errors.Wrap(err, "create Api failed")
		return err
	}
	return nil
}

func (a Api) Update(req do.ApiInfo) error {
	_, err := a.db.API.UpdateOneID(req.ID).
		SetPath(req.Path).
		SetDescription(req.Description).
		SetAPIGroup(req.Group).
		SetMethod(req.Method).
		Save(a.ctx)
	if err != nil {
		err = errors.Wrap(err, "update Api failed")
		return err
	}
	return nil
}

func (a Api) Delete(id int64) error {
	err := a.db.API.DeleteOneID(id).Exec(a.ctx)
	return err
}

func (a Api) List(req do.ListApiReq) (resp []*do.ApiInfo, total int, err error) {
	var predicates []predicate.API
	if req.Path != "" {
		predicates = append(predicates, api.PathContains(req.Path))
	}
	if req.Description != "" {
		predicates = append(predicates, api.DescriptionContains(req.Description))
	}
	if req.Method != "" {
		predicates = append(predicates, api.MethodContains(req.Method))
	}
	if req.Group != "" {
		predicates = append(predicates, api.APIGroupContains(req.Group))
	}

	apis, err := a.db.API.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(a.ctx)
	if err != nil {
		err = errors.Wrap(err, "get api list failed")
		return resp, total, err
	}
	for _, a := range apis {
		resp = append(resp, &do.ApiInfo{
			ID:          a.ID,
			CreatedAt:   a.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   a.UpdatedAt.Format("2006-01-02 15:04:05"),
			Path:        a.Path,
			Description: a.Description,
			Group:       a.APIGroup,
			Method:      a.Method,
		})
	}
	total, _ = a.db.API.Query().Where(predicates...).Count(a.ctx)
	return resp, total, nil
}

func NewApi(ctx context.Context, c *app.RequestContext) do.Api {
	return &Api{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
