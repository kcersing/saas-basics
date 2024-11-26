package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/dictionary"
	"saas/pkg/db/ent/dictionarydetail"
	"saas/pkg/db/ent/predicate"
	"time"
)

type Dictionary struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (d Dictionary) Create(req *do.DictionaryInfo) error {
	// whether dictionary name exists
	dictionaryExist, _ := d.db.Dictionary.Query().Where(dictionary.Name(req.Name)).Exist(d.ctx)
	if dictionaryExist {
		return errors.New("dictionary name already exists")
	}
	// create dictionary
	_, err := d.db.Dictionary.Create().
		SetTitle(req.Title).
		SetName(req.Name).
		SetStatus(req.Status).
		SetDescription(req.Description).
		Save(d.ctx)
	if err != nil {
		return errors.Wrap(err, "create Dictionary failed")
	}
	return nil
}

func (d Dictionary) Update(req *do.DictionaryInfo) error {
	// whether dictionary is exists
	dictionaryExist, _ := d.db.Dictionary.Query().Where(dictionary.ID(req.ID)).Exist(d.ctx)
	if !dictionaryExist {
		return errors.New("The dictionary try to update is not exists")
	}
	// update dictionary
	_, err := d.db.Dictionary.UpdateOneID(req.ID).
		SetTitle(req.Title).
		SetName(req.Name).
		SetStatus(req.Status).
		SetDescription(req.Description).
		Save(d.ctx)
	if err != nil {
		return errors.Wrap(err, "update Dictionary failed")
	}
	return nil
}

func (d Dictionary) Delete(id int64) error {
	// whether dictionary is exists
	dict, err := d.db.Dictionary.Query().Where(dictionary.ID(id)).Only(d.ctx)
	if err != nil {
		return errors.Wrap(err, "query Dictionary failed")
	}
	if dict == nil {
		return errors.New(fmt.Sprintf("The dictionary(id=%d) try to delete is not exists", id))
	}
	// whether dictionary has detail
	// query dictionary detail
	details, err := d.db.DictionaryDetail.Query().
		Where(dictionarydetail.HasDictionaryWith(dictionary.NameEQ(dict.Name))).
		// union query to get the fields of the associated table
		WithDictionary(func(q *ent.DictionaryQuery) {
			// get all fields default, or use q.Select() to get some fields
		}).All(d.ctx)
	if err != nil {
		return errors.Wrap(err, "query DictionaryDetail failed")
	}
	if len(details) > 0 {
		return errors.New("dictionary has detail, please delete detail first")
	}
	// delete dictionary
	err = d.db.Dictionary.DeleteOneID(id).Exec(d.ctx)
	if err != nil {
		return errors.Wrap(err, "delete Dictionary failed")
	}
	return nil
}

func (d Dictionary) List(req *do.DictListReq) (list []*do.DictionaryInfo, total int, err error) {
	// query dictionary
	var predicates []predicate.Dictionary
	if req.Title != "" {
		predicates = append(predicates, dictionary.TitleContains(req.Title))
	}
	if req.Name != "" {
		predicates = append(predicates, dictionary.NameContains(req.Name))
	}
	dictionaries, err := d.db.Dictionary.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(d.ctx)
	if err != nil {
		return nil, 0, errors.Wrap(err, "query Dictionary list failed")
	}

	// format result
	for _, dict := range dictionaries {
		list = append(list, &do.DictionaryInfo{
			ID:          dict.ID,
			Title:       dict.Title,
			Name:        dict.Name,
			Status:      dict.Status,
			Description: dict.Description,
			CreatedAt:   dict.CreatedAt.Format(time.DateTime),
			UpdatedAt:   dict.UpdatedAt.Format(time.DateTime),
		})
	}
	total, _ = d.db.Dictionary.Query().Where(predicates...).Count(d.ctx)
	return
}

func (d Dictionary) CreateDetail(req *do.DictionaryDetail) error {
	// whether dictionary detail is exists
	exist, err := d.db.DictionaryDetail.Query().
		Where(dictionarydetail.Key(req.Key)).
		Where(dictionarydetail.Value(req.Value)).
		Where(dictionarydetail.HasDictionaryWith(dictionary.ID(req.ParentID))).
		Exist(d.ctx)
	if err != nil {
		return errors.Wrap(err, "query DictionaryDetail exist failed")
	}
	if exist {
		return errors.New("dictionary detail already exists")
	}

	// find dictionary by id
	dict, err := d.db.Dictionary.Query().Where(dictionary.ID(req.ParentID)).Only(d.ctx)
	if err != nil {
		return errors.Wrap(err, "query Dictionary failed")
	}
	if dict == nil {
		return errors.New(fmt.Sprintf("dictionary not found, please check dictionary id, %d", req.ParentID))
	}

	// create dictionary detail
	_, err = d.db.DictionaryDetail.Create().
		SetDictionary(dict). // set parent dictionary
		SetTitle(req.Title).
		SetKey(req.Key).
		SetValue(req.Value).
		SetStatus(req.Status).
		Save(d.ctx)
	if err != nil {
		return errors.Wrap(err, "create DictionaryDetail failed")
	}
	return nil
}

func (d Dictionary) UpdateDetail(req *do.DictionaryDetail) error {
	// query dictionary detail
	detail, err := d.db.DictionaryDetail.Query().
		Where(dictionarydetail.ID(req.ID)).
		// union query to get the fields of the associated table
		WithDictionary(func(q *ent.DictionaryQuery) {
			// get all fields default, or use q.Select() to get some fields
		}).Only(d.ctx)
	if err != nil {
		return errors.Wrap(err, "query DictionaryDetail failed")
	}
	// update dictionary detail
	_, err = d.db.DictionaryDetail.UpdateOneID(req.ID).
		SetTitle(req.Title).
		SetKey(req.Key).
		SetValue(req.Value).
		SetStatus(req.Status).
		Save(d.ctx)
	if err != nil {
		return errors.Wrap(err, "update DictionaryDetail failed")
	}
	// delete dictionary detail from cache
	d.cache.Del(fmt.Sprintf("Dictionary%s-key%s", detail.Edges.Dictionary.Name, detail.Key))
	d.cache.Del(fmt.Sprintf("Dictionary%s-value%s", detail.Edges.Dictionary.Name, detail.Value))
	return nil
}

func (d Dictionary) DeleteDetail(id int64) error {
	// query dictionary detail
	detail, err := d.db.DictionaryDetail.Query().
		Where(dictionarydetail.ID(id)).
		// union query to get the fields of the associated table
		WithDictionary(func(q *ent.DictionaryQuery) {
			// get all fields default, or use q.Select() to get some fields
		}).Only(d.ctx)
	if err != nil {
		return errors.Wrap(err, "query DictionaryDetail failed")
	}
	// delete dictionary detail
	err = d.db.DictionaryDetail.DeleteOneID(id).Exec(d.ctx)
	if err != nil {
		return errors.Wrap(err, "delete DictionaryDetail failed")
	}
	// delete dictionary detail from cache
	d.cache.Del(fmt.Sprintf("Dictionary%s-key%s", detail.Edges.Dictionary.Name, detail.Key))
	d.cache.Del(fmt.Sprintf("Dictionary%s-value%s", detail.Edges.Dictionary.Name, detail.Value))
	return nil
}

func (d Dictionary) DetailListByDict(req *do.DetailListReq) (list []*do.DictionaryDetail, total int64, err error) {

	var predicates []predicate.DictionaryDetail
	if req.Name != "" {
		predicates = append(predicates, dictionarydetail.HasDictionaryWith(dictionary.NameEQ(req.Name)))
	}

	if req.DictionaryId != 0 {
		predicates = append(predicates, dictionarydetail.HasDictionaryWith(dictionary.IDEQ(req.DictionaryId)))
	}

	// query dictionary detail
	details, err := d.db.DictionaryDetail.Query().
		Where(predicates...).
		//Where(dictionarydetail.HasDictionaryWith(dictionary.NameEQ(dictName))).
		// union query to get the fields of the associated table
		WithDictionary(func(q *ent.DictionaryQuery) {
			// get all fields default, or use q.Select() to get some fields
		}).All(d.ctx)
	if err != nil {
		return nil, 0, errors.Wrap(err, "query DictionaryDetail list failed")
	}

	// format result
	for _, detail := range details {
		list = append(list, &do.DictionaryDetail{
			ID:        detail.ID,
			Title:     detail.Title,
			Key:       detail.Key,
			Value:     detail.Value,
			Status:    detail.Status,
			CreatedAt: detail.CreatedAt.Format(time.DateTime),
			UpdatedAt: detail.UpdatedAt.Format(time.DateTime),
			ParentID:  detail.Edges.Dictionary.ID,
		})
	}
	total = int64(len(list))
	return
}

func (d Dictionary) DetailByDictNameAndKey(dictName, key string) (detail *do.DictionaryDetail, err error) {
	// query dictionary detail from cache
	v, found := d.cache.Get(fmt.Sprintf("Dictionary%s-key%s", dictName, key))
	if found {
		return v.(*do.DictionaryDetail), nil
	}
	// query dictionary detail from database
	dictDetail, err := d.db.DictionaryDetail.Query().
		Where(dictionarydetail.HasDictionaryWith(dictionary.NameEQ(dictName))).
		Where(dictionarydetail.Key(key)).First(d.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "query DictionaryDetail list failed")
	}

	// format result
	detail = new(do.DictionaryDetail)
	detail.ID = dictDetail.ID
	detail.Title = dictDetail.Title
	detail.Key = dictDetail.Key
	detail.Value = dictDetail.Value
	detail.Status = dictDetail.Status
	detail.CreatedAt = dictDetail.CreatedAt.Format(time.DateTime)
	detail.UpdatedAt = dictDetail.UpdatedAt.Format(time.DateTime)

	// set cache
	d.cache.SetWithTTL(fmt.Sprintf("Dictionary%s-key%s", dictName, key), detail, 1, 1*time.Hour)
	return detail, nil
}

func NewDictionary(ctx context.Context, c *app.RequestContext) do.Dictionary {
	return &Dictionary{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
