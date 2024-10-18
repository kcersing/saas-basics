package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/api"
	"saas/pkg/db/ent/predicate"
	"time"
)

type Api struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (a Api) ApiTree(req do.ListApiReq) (resp []*do.Tree, total int, err error) {

	inter, exist := a.cache.Get("ApiTree")
	if exist {
		if v, ok := inter.([]*do.Tree); ok {
			return v, len(v), nil
		}
	}

	var predicates []predicate.API

	apis, err := a.db.API.Query().All(a.ctx)
	if err != nil {
		err = errors.Wrap(err, "get api list failed")
		return resp, total, err
	}
	apiGroups, err := a.db.API.Query().GroupBy(api.FieldAPIGroup).Strings(a.ctx)

	for _, apiGroup := range apiGroups {
		g := &do.Tree{
			Title: apiGroup,
		}
		for _, v := range apis {
			if v.APIGroup == g.Title {
				g.Children = append(g.Children, &do.Tree{
					Title: v.Title,
					Value: map[string]string{"path": v.Path, "method": v.Method},
					Key:   v.ID, // map[string]string{"path": v.Path, "method": v.Method},
				})
			}
		}

		resp = append(resp, g)
	}

	total, _ = a.db.API.Query().Where(predicates...).Count(a.ctx)

	a.cache.SetWithTTL("ApiTree", &resp, 1, 30*time.Hour)
	return
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

	err = copier.Copy(&resp, &apis)
	if err != nil {
		err = errors.Wrap(err, "copy Api failed")
		return resp, 0, err
	}
	for i, v := range apis {
		resp[i].CreatedAt = v.CreatedAt.Format(time.DateTime)
		resp[i].UpdatedAt = v.UpdatedAt.Format(time.DateTime)

	}

	total, _ = a.db.API.Query().Where(predicates...).Count(a.ctx)
	return resp, total, nil
}

func NewApi(ctx context.Context, c *app.RequestContext) do.Api {
	return &Api{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
