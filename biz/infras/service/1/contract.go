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
	"saas/pkg/db/ent/contract"
	"saas/pkg/db/ent/predicate"
	"strconv"
	"time"
)

type Contract struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (c Contract) Info(id int64) (info *do.ContractInfo, err error) {
	inter, exist := c.cache.Get("contractInfo" + strconv.Itoa(int(id)))
	if exist {
		if v, ok := inter.(*do.ContractInfo); ok {
			return v, nil
		}
	}
	one, err := c.db.Contract.Query().Where(contract.IDEQ(id)).First(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "get contract failed")
		return info, err
	}

	err = copier.Copy(&info, &one)
	if err != nil {
		err = errors.Wrap(err, "copy contract info failed")
		return info, err
	}

	c.cache.SetWithTTL("contractInfo"+strconv.Itoa(int(info.ID)), &info, 1, 1*time.Hour)
	return
}

func (c Contract) Create(req *do.ContractInfo) error {
	_, err := c.db.Contract.Create().
		SetName(req.Name).
		SetContent(req.Content).
		Save(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "create Contract failed")
		return err
	}
	return nil
}

func (c Contract) Update(req *do.ContractInfo) error {
	_, err := c.db.Contract.Update().
		Where(contract.IDEQ(req.ID)).
		SetName(req.Name).
		SetContent(req.Content).
		Save(c.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Contract failed")
		return err
	}

	return nil
}

func (c Contract) UpdateStatus(ID int64, status int64) error {
	_, err := c.db.Contract.Update().Where(contract.IDEQ(ID)).SetStatus(status).Save(c.ctx)
	return err
}

func (c Contract) List(req *do.ContractListReq) (list []*do.ContractInfo, total int, err error) {
	var predicates []predicate.Contract

	if req.Name != "" {
		predicates = append(predicates, contract.NameEQ(req.Name))
	}

	all, err := c.db.Contract.
		Query().
		Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Contract list failed")
		return list, total, err
	}

	err = copier.Copy(&list, &all)
	if err != nil {
		err = errors.Wrap(err, "copy Contract info failed")
		return list, 0, err
	}

	for i, v := range all {
		list[i].CreatedAt = v.CreatedAt.Format(time.DateTime)
	}
	total, _ = c.db.Contract.Query().Where(predicates...).Count(c.ctx)
	return
}

func NewContract(ctx context.Context, c *app.RequestContext) do.Contract {
	return &Contract{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
