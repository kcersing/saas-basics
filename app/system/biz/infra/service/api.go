package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"rpc_gen/kitex_gen/base"
	"rpc_gen/kitex_gen/system/menu"
	"strconv"
	"system/biz/dal/cache"
	"system/biz/dal/mysql"
	"system/biz/dal/mysql/ent"
	"system/biz/dal/mysql/ent/api"
	"system/biz/dal/mysql/ent/predicate"
	"system/biz/infra/do"

	"time"
)

type Api struct {
	ctx   context.Context
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func NewApi(ctx context.Context) do.Api {
	return &Api{
		ctx:   ctx,
		salt:  "",
		db:    mysql.DB,
		cache: cache.Cache,
	}
}

func (a Api) ApiTree(req menu.ApiPageReq) (resp []*base.Tree, total int64, err error) {

	inter, exist := a.cache.Get("ApiTree")
	if exist {
		if v, ok := inter.([]*base.Tree); ok {
			return v, int64(len(v)), nil
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
		g := &base.Tree{
			Title: apiGroup,
		}
		for _, v := range apis {
			if v.APIGroup == g.Title {
				dataType, _ := json.Marshal(map[string]string{"path": v.Path, "method": v.Method})
				dataString := string(dataType)
				g.Children = append(g.Children, &base.Tree{
					Title: v.Title,
					Value: dataString,
					Key:   strconv.FormatInt(v.ID, 10),
				})
			}
		}

		resp = append(resp, g)
	}

	t, _ := a.db.API.Query().Where(predicates...).Count(a.ctx)
	total = int64(t)
	a.cache.SetWithTTL("ApiTree", &resp, 1, 30*time.Hour)
	return
}

func (a Api) Create(req menu.ApiInfo) error {
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

func (a Api) Update(req menu.ApiInfo) error {
	_, err := a.db.API.UpdateOneID(req.Id).
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

func (a Api) List(req menu.ApiPageReq) (resp []*menu.ApiInfo, total int64, err error) {
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

	t, _ := a.db.API.Query().Where(predicates...).Count(a.ctx)
	total = int64(t)
	return resp, int64(total), nil
}
