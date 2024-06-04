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
		err = errors.Wrap(err, "create Venue failed")
		return err
	}

	return nil
}

func (c Contract) UpdateStatus(ID int64, status int64) error {
	_, err := c.db.Contract.Update().Where(contract.IDEQ(ID)).SetStatus(status).Save(c.ctx)
	return err
}

func (c Contract) ByID(id int64) (*do.ContractInfo, error) {
	contractInterface, ok := c.cache.Get("contractData" + strconv.Itoa(int(id)))
	if ok {
		if l, ok := contractInterface.(*ent.Contract); ok {
			return &do.ContractInfo{
				ID:      l.ID,
				Name:    l.Name,
				Content: l.Content,
			}, nil
		}
	}
	// get Content from db
	contentEnt, err := c.db.Contract.Query().Where(contract.IDEQ(id)).Only(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Contract failed")
		return nil, err
	}
	// set Content to cache
	c.cache.SetWithTTL("contractData"+strconv.Itoa(int(id)), contentEnt, 1, 24*time.Hour)
	// convert to contentInfo
	return &do.ContractInfo{
		ID:      contentEnt.ID,
		Name:    contentEnt.Name,
		Content: contentEnt.Content,
	}, err
}

func (c Contract) List(req *do.ContractListReq) (list []*do.ContractInfo, total int, err error) {
	var predicates []predicate.Contract

	if req.Name != "" {
		predicates = append(predicates, contract.NameEQ(req.Name))
	}

	all, err := c.db.Contract.Query().Where(predicates...).
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
	total, _ = c.db.Contract.Query().Where(predicates...).Count(c.ctx)
	return
}

func NewContract(ctx context.Context, c *app.RequestContext) do.Contract {
	return &Contract{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
