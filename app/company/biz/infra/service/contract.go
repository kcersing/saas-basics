package service

import (
	"company/biz/dal/cache"
	"company/biz/dal/mysql"
	"company/biz/dal/mysql/ent"
	"rpc_gen/kitex_gen/company/contract"

	contract2 "company/biz/dal/mysql/ent/contract"
	"company/biz/dal/mysql/ent/predicate"
	"company/biz/infra/do"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

type Contract struct {
	ctx   context.Context
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func NewContract(ctx context.Context, c *app.RequestContext) do.Contract {
	return &Contract{
		ctx:   ctx,
		salt:  "",
		db:    mysql.DB,
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

	err = copier.Copy(&info, &one)
	if err != nil {
		err = errors.Wrap(err, "copy contract info failed")
		return info, err
	}

	c.cache.SetWithTTL("contractInfo"+strconv.Itoa(int(info.Id)), &info, 1, 1*time.Hour)
	return
}

func (c Contract) Create(req *contract.ContractInfo) error {
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

func (c Contract) Update(req *contract.ContractInfo) error {
	_, err := c.db.Contract.Update().
		Where(contract2.IDEQ(req.Id)).
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

	if *req.Name != "" {
		predicates = append(predicates, contract2.NameEQ(*req.Name))
	}

	all, err := c.db.Contract.
		Query().
		Where(predicates...).
		Offset(int(*req.Page-1) * int(*req.PageSize)).
		Limit(int(*req.PageSize)).All(c.ctx)
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
