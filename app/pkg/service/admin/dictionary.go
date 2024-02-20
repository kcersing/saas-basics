package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
)

type Dictionary struct {
	ctx  context.Context
	c    *app.RequestContext
	salt string
	db   *ent.Client
}

func (d Dictionary) Create(req *do.DictionaryInfo) error {
	//TODO implement me
	panic("implement me")
}

func (d Dictionary) Update(req *do.DictionaryInfo) error {
	//TODO implement me
	panic("implement me")
}

func (d Dictionary) Delete(id uint64) error {
	//TODO implement me
	panic("implement me")
}

func (d Dictionary) List(req *do.DictListReq) (list []*do.DictionaryInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (d Dictionary) CreateDetail(req *do.DictionaryDetail) error {
	//TODO implement me
	panic("implement me")
}

func (d Dictionary) UpdateDetail(req *do.DictionaryDetail) error {
	//TODO implement me
	panic("implement me")
}

func (d Dictionary) DeleteDetail(id uint64) error {
	//TODO implement me
	panic("implement me")
}

func (d Dictionary) DetailListByDictName(dictName string) (list []*do.DictionaryDetail, total uint64, err error) {
	//TODO implement me
	panic("implement me")
}

func (d Dictionary) DetailByDictNameAndKey(dictName, key string) (detail *do.DictionaryDetail, err error) {
	//TODO implement me
	panic("implement me")
}

func NewDictionary(ctx context.Context, c *app.RequestContext) do.Dictionary {
	return &Dictionary{
		ctx:  ctx,
		c:    c,
		salt: config.GlobalServerConfig.MysqlInfo.Salt,
		db:   infras.DB,
	}
}
