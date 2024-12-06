package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/token"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/predicate"
	token2 "saas/pkg/db/ent/token"
	entuser "saas/pkg/db/ent/user"
	"time"
)

type Token struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func NewToken(ctx context.Context, c *app.RequestContext) do.Token {
	return &Token{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}

func (t Token) Create(req *token.TokenInfo) error {
	expiredAt, _ := time.ParseInLocation(time.DateTime, req.ExpiredAt, time.Local)
	if expiredAt.Sub(time.Now()).Seconds() < 5 {
		return errors.New("expired time must be greater than now, more than 5s")
	}
	tokenExist, err := t.db.Token.Query().Where(token2.UserID(req.UserID)).Only(t.ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return errors.Wrap(err, "create Token failed")
		}
	}
	if tokenExist != nil {
		req.ID = tokenExist.ID
		return t.Update(req)
	}

	userInfo, err := t.db.User.Query().Where(entuser.IDEQ(req.UserID)).Only(t.ctx)
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

func (t Token) Update(req *token.TokenInfo) error {
	//TODO implement me
	expiredAt, _ := time.ParseInLocation(time.DateTime, req.ExpiredAt, time.Local)
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

func (t Token) IsExistByUserID(userID int64) bool {
	_, exist := t.cache.Get(fmt.Sprintf("token_%d", userID))
	if exist {
		return true
	}
	exist, _ = t.db.Token.Query().Where(token2.UserID(userID)).Exist(t.ctx)
	return exist
}

func (t Token) Delete(userID int64) error {
	_, err := t.db.Token.Delete().Where(token2.UserID(userID)).Exec(t.ctx)
	if err != nil {
		return errors.Wrap(err, "delete Token failed")
	}
	t.cache.Del(fmt.Sprintf("token_%d", userID))
	return nil
}

func (t Token) List(req *token.TokenListReq) (res []*token.TokenInfo, total int, err error) {
	// list token with user info
	var userPredicates = []predicate.User{entuser.HasToken()}
	if req.Username != "" {
		userPredicates = append(userPredicates, entuser.UsernameContainsFold(req.Username))
	}
	if req.UserId != 0 {
		userPredicates = append(userPredicates, entuser.IDEQ(req.UserId))
	}
	userPredicates = append(userPredicates, entuser.Delete(0))
	UserTokens, err := t.db.User.Query().Where(userPredicates...).
		WithToken(func(q *ent.TokenQuery) {
			// get token all fields default, or use q.Select() to get some fields
		}).Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(t.ctx)
	if err != nil {
		return res, total, errors.Wrap(err, "get User - Token list failed")
	}

	for _, userEnt := range UserTokens {
		res = append(res, &token.TokenInfo{
			ID:        userEnt.Edges.Token.ID,
			UserID:    userEnt.ID,
			Username:  userEnt.Username,
			Token:     userEnt.Edges.Token.Token,
			Source:    userEnt.Edges.Token.Source,
			CreatedAt: userEnt.CreatedAt.Format(time.DateTime),
			UpdatedAt: userEnt.UpdatedAt.Format(time.DateTime),
			ExpiredAt: userEnt.Edges.Token.ExpiredAt.Format(time.DateTime),
		})

		// delete expired token from db
		if userEnt.Edges.Token.ExpiredAt.Sub(time.Now()).Seconds() < 0 {
			_ = t.Delete(userEnt.ID)
		}
	}
	total, _ = t.db.User.Query().Where(userPredicates...).Count(t.ctx)
	return
}
