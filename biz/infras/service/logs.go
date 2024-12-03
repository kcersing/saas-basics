package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/auth"

	"saas/pkg/db/ent"
	"saas/pkg/db/ent/logs"
	"saas/pkg/db/ent/predicate"
	"time"
)

type Logs struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func NewLogs(ctx context.Context, c *app.RequestContext) do.Logs {
	return &Logs{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}

func (l Logs) Create(logsReq *auth.LogsInfo) error {
	err := l.db.Logs.Create().
		SetType(logsReq.Type).
		SetMethod(logsReq.Method).
		SetAPI(logsReq.Api).
		SetSuccess(logsReq.Success).
		SetReqContent(logsReq.ReqContent).
		SetRespContent(logsReq.RespContent).
		SetIP(logsReq.Ip).
		SetUserAgent(logsReq.UserAgent).
		SetOperator(logsReq.Operator).
		SetTime(int(logsReq.Time)).
		Exec(l.ctx)
	if err != nil {
		err = errors.Wrap(err, "create logs failed")
		return err
	}
	return nil
}

func (l Logs) List(req *auth.LogsListReq) (list []*auth.LogsInfo, total int64, err error) {
	var predicates []predicate.Logs
	if *req.Type != "" {
		predicates = append(predicates, logs.TypeEQ(*req.Type))
	}
	if *req.Method != "" {
		predicates = append(predicates, logs.MethodEQ(*req.Method))
	}
	if *req.API != "" {
		predicates = append(predicates, logs.APIContains(*req.API))
	}
	if *req.Operators != "" {
		predicates = append(predicates, logs.OperatorContains(*req.Operators))
	}
	//if req.Success != nil {
	//	predicates = append(predicates, logs.SuccessEQ(req.Success))
	//}
	logsData, err := l.db.Logs.Query().Where(predicates...).
		Offset(int((*req.Page - 1) * *req.PageSize)).
		Limit(int(*req.PageSize)).
		Order(ent.Desc(logs.FieldCreatedAt)).All(l.ctx)
	if err != nil {
		return nil, 0, errors.Wrap(err, "query logsData list failed")
	}
	for _, v := range logsData {
		list = append(list, &auth.LogsInfo{
			Type:        v.Type,
			Method:      v.Method,
			Api:         v.API,
			Success:     v.Success,
			ReqContent:  v.ReqContent,
			RespContent: v.RespContent,
			Ip:          v.IP,
			UserAgent:   v.UserAgent,
			Operator:    v.Operator,
			Time:        int64(v.Time),
			CreatedAt:   v.CreatedAt.Format(time.DateTime),
			UpdatedAt:   v.UpdatedAt.Format(time.DateTime),
		})
	}
	t, err := l.db.Logs.Query().Where(predicates...).Count(l.ctx)
	total = int64(t)
	if err != nil {
		return nil, 0, errors.Wrap(err, "query logsData count failed")
	}
	return
}

func (l Logs) DeleteAll(ids []int64) error {
	_, err := l.db.Logs.Delete().Exec(l.ctx)
	if err != nil {
		return errors.Wrap(err, "delete logsData failed")
	}
	return nil
}
