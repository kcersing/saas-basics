package admin

import (
	"context"
	"fmt"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/pkg/db/ent/user"

	"github.com/cloudwego/hertz/pkg/app"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/token"
	"time"
)

type Token struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (t Token) Create(req *do.TokenInfo) error {
	//TODO implement me
	expiredAt, _ := time.ParseInLocation("2006-01-02 15:04:05", req.ExpiredAt, time.Local)
	if expiredAt.Sub(time.Now()).Seconds() < 5 {
		return errors.New("expired time must be greater than now, more than 5s")
	}
	tokenExist, err := t.db.Token.Query().Where(token.UserID(req.UserID)).Only(t.ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return errors.Wrap(err, "create Token failed")
		}
	}
	if tokenExist != nil {
		req.ID = tokenExist.ID
		return t.Update(req)
	}

	userInfo, err := t.db.User.Query().Where(user.IDEQ(req.UserID)).Only(t.ctx)
	if err != nil {
		return errors.Wrap(err, "get userinfo failed")
	}
	_, err = t.db.Token.Create().
		SetOwner(userInfo).
		SetUserID(req.UserID).
		SetToken(req.Token).
		SetSource(req.Source).
		SetExpiredAt(expiredAt).
		Save(t.ctx)
	if err != nil {
		return errors.Wrap(err, "create Token failed")
	}
	t.cache.SetWithTTL(fmt.Sprintf("token_%d", req.UserID), req.UserID, 0, expiredAt.Sub(time.Now()))
	t.cache.Wait()
	return nil
}

func (t Token) Update(req *do.TokenInfo) error {
	//TODO implement me
	expiredAt, _ := time.ParseInLocation("2006-01-02 15:04:05", req.ExpiredAt, time.Local)
	if expiredAt.Sub(time.Now()).Seconds() < 5 {
		return errors.New("expired time must be greater than now, more than 5s")
	}
	_, err := t.db.Token.UpdateOneID(req.ID).
		SetUserID(req.UserID).
		SetToken(req.Token).
		SetSource(req.Source).
		SetUpdatedAt(time.Now()).
		SetExpiredAt(expiredAt).
		Save(t.ctx)
	if err != nil {
		return errors.Wrap(err, "update Token failed")
	}

	t.cache.SetWithTTL(fmt.Sprintf("token_%d", req.UserID), req.UserID, 1, expiredAt.Sub(time.Now()))
	return nil
}

func (t Token) IsExistByUserID(userID uint64) bool {
	//TODO implement me
	_, exist := t.cache.Get(fmt.Sprintf("token_%d", userID))
	if exist {
		return true
	}
	exist, _ = t.db.Token.Query().Where(token.UserID(userID)).Exist(t.ctx)
	return exist
}

func (t Token) Delete(userID uint64) error {
	_, err := t.db.Token.Delete().Where(token.UserID(userID)).Exec(t.ctx)
	if err != nil {
		return errors.Wrap(err, "delete Token failed")
	}
	t.cache.Del(fmt.Sprintf("token_%d", userID))
	return nil
}

func (t Token) List(req *do.TokenListReq) (res []*do.TokenInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func NewToken(ctx context.Context, c *app.RequestContext) do.Token {
	return &Token{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
