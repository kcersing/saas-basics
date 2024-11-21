package service

import (
	"context"
	"fmt"
	"rpc_gen/kitex_gen/user"
	"time"
	"user/biz/dal/cache"
	"user/biz/dal/mysql"
	"user/biz/dal/mysql/ent"
	"user/biz/dal/mysql/ent/predicate"
	"user/biz/dal/mysql/ent/token"
	user2 "user/biz/dal/mysql/ent/user"
	"user/biz/infra/do"

	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
)

type Token struct {
	ctx   context.Context
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func NewToken(ctx context.Context) do.Token {
	return &Token{
		ctx:   ctx,
		salt:  "",
		db:    mysql.DB,
		cache: cache.Cache,
	}
}

func (t Token) Create(req *user.TokenInfo) error {
	expiredAt, _ := time.ParseInLocation(time.DateTime, req.ExpiredAt, time.Local)
	if expiredAt.Sub(time.Now()).Seconds() < 5 {
		return errors.New("expired time must be greater than now, more than 5s")
	}
	tokenExist, err := t.db.Token.Query().Where(token.UserID(req.UserId)).Only(t.ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return errors.Wrap(err, "create Token failed")
		}
	}
	if tokenExist != nil {
		req.Id = tokenExist.ID
		return t.Update(req)
	}

	userInfo, err := t.db.User.Query().Where(user2.IDEQ(req.UserId)).Only(t.ctx)
	if err != nil {
		return errors.Wrap(err, "get userinfo failed")
	}
	_, err = t.db.Token.Create().
		SetOwner(userInfo).
		SetUserID(req.UserId).
		SetToken(req.Token).
		SetSource(req.Source).
		SetExpiredAt(expiredAt).
		Save(t.ctx)
	if err != nil {
		return errors.Wrap(err, "create Token failed")
	}
	t.cache.SetWithTTL(fmt.Sprintf("token_%d", req.UserId), req.UserId, 0, expiredAt.Sub(time.Now()))
	t.cache.Wait()
	return nil
}

func (t Token) Update(req *user.TokenInfo) error {
	//TODO implement me
	expiredAt, _ := time.ParseInLocation(time.DateTime, req.ExpiredAt, time.Local)
	if expiredAt.Sub(time.Now()).Seconds() < 5 {
		return errors.New("expired time must be greater than now, more than 5s")
	}
	_, err := t.db.Token.UpdateOneID(req.Id).
		SetUserID(req.UserId).
		SetToken(req.Token).
		SetSource(req.Source).
		SetUpdatedAt(time.Now()).
		SetExpiredAt(expiredAt).
		Save(t.ctx)
	if err != nil {
		return errors.Wrap(err, "update Token failed")
	}

	t.cache.SetWithTTL(fmt.Sprintf("token_%d", req.UserId), req.UserId, 1, expiredAt.Sub(time.Now()))
	return nil
}

func (t Token) IsExistByUserID(userID int64) bool {
	_, exist := t.cache.Get(fmt.Sprintf("token_%d", userID))
	if exist {
		return true
	}
	exist, _ = t.db.Token.Query().Where(token.UserID(userID)).Exist(t.ctx)
	return exist
}

func (t Token) Delete(userID int64) error {
	_, err := t.db.Token.Delete().Where(token.UserID(userID)).Exec(t.ctx)
	if err != nil {
		return errors.Wrap(err, "delete Token failed")
	}
	t.cache.Del(fmt.Sprintf("token_%d", userID))
	return nil
}

func (t Token) List(req *user.TokenListReq) (res []*user.TokenInfo, total int64, err error) {
	// list token with user info
	var userPredicates = []predicate.User{user2.HasToken()}
	if req.Username != "" {
		userPredicates = append(userPredicates, user2.UsernameContainsFold(req.Username))
	}
	if req.UserId != 0 {
		userPredicates = append(userPredicates, user2.IDEQ(req.UserId))
	}
	UserTokens, err := t.db.User.Query().Where(userPredicates...).
		WithToken(func(q *ent.TokenQuery) {
			// get token all fields default, or use q.Select() to get some fields
		}).Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(t.ctx)
	if err != nil {
		return res, total, errors.Wrap(err, "get User - Token list failed")
	}

	for _, userEnt := range UserTokens {
		res = append(res, &user.TokenInfo{
			Id:        userEnt.Edges.Token.ID,
			UserId:    userEnt.ID,
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
	tl, _ := t.db.User.Query().Where(userPredicates...).Count(t.ctx)
	total = int64(tl)
	return
}
