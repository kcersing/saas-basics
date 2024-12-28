package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	contract2 "saas/biz/dal/db/ent/contract"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/contract"
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

func (c Contract) Delete(ID int64) error {
	_, err := c.db.Contract.Update().Where(contract2.IDEQ(ID)).SetDelete(1).Save(c.ctx)
	return err
}

func NewContract(ctx context.Context, c *app.RequestContext) do.Contract {
	return &Contract{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}

func (c Contract) Info(id int64) (info *contract.ContractInfo, err error) {
	inter, exist := c.cache.Get("contractInfo" + strconv.Itoa(int(id)))
	if exist {
		if v, ok := inter.(*contract.ContractInfo); ok {
			return v, nil
		}
	}
	one, err := c.db.Contract.Query().Where(contract2.IDEQ(id)).First(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "get contract failed")
		return info, err
	}
	info = entContractInfo(one)

	c.cache.SetWithTTL("contractInfo"+strconv.Itoa(int(info.ID)), &info, 1, 1*time.Hour)
	return
}

func (c Contract) Create(req *contract.CreateOrUpdateContractReq) error {
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

func (c Contract) Update(req *contract.CreateOrUpdateContractReq) error {
	_, err := c.db.Contract.Update().
		Where(contract2.IDEQ(req.ID)).
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
	_, err := c.db.Contract.Update().Where(contract2.IDEQ(ID)).SetStatus(status).Save(c.ctx)
	return err
}

func (c Contract) List(req *contract.ContractListReq) (list []*contract.ContractInfo, total int, err error) {
	var predicates []predicate.Contract

	if req.Name != "" {
		predicates = append(predicates, contract2.NameEQ(req.Name))
	}
	predicates = append(predicates, contract2.Delete(0))
	all, err := c.db.Contract.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Contract list failed")
		return list, total, err
	}

	for _, v := range all {
		list = append(list, entContractInfo(v))
	}
	total, _ = c.db.Contract.Query().Where(predicates...).Count(c.ctx)
	return
}

func entContractInfo(v *ent.Contract) *contract.ContractInfo {

	return &contract.ContractInfo{
		ID:        v.ID,
		Name:      v.Name,
		Content:   v.Content,
		Status:    v.Status,
		CreatedAt: v.CreatedAt.Format(time.DateTime),
	}
}
