package service

import (
	"context"
	"rpc_gen/kitex_gen/base"
	"user/biz/dal/mysql/ent/user"

	"user/biz/dal/cache"
	"user/biz/dal/mysql"
	"user/biz/dal/mysql/ent"
	"user/biz/dal/mysql/ent/predicate"
	"user/biz/infra/do"

	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
)

type Sys struct {
	ctx   context.Context
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func NewSys(ctx context.Context) do.Sys {
	return &Sys{
		ctx:   ctx,
		salt:  "",
		db:    mysql.DB,
		cache: cache.Cache,
	}
}
func (s Sys) StaffList(req base.ListReq) (list []base.SysList, total int64, err error) {
	var predicates []predicate.User

	if *req.Name != "" {
		predicates = append(predicates, user.Nickname(*req.Name))
	}

	lists, err := s.db.User.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		list = append(list, base.SysList{
			Id:   v.ID,
			Name: v.Nickname,
		})
	}
	total = int64(len(list))
	return
}
