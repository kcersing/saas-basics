package service

import (
	"context"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"rpc_gen/kitex_gen/system/auth"
	"system/biz/dal/mysql/ent"
	"system/biz/dal/mysql/ent/logs"
	"system/biz/dal/mysql/ent/predicate"
	"system/biz/infra/do"

	"system/biz/dal/cache"
	"system/biz/dal/mysql"
	"time"
)

type Logs struct {
	ctx   context.Context
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func NewLogs(ctx context.Context) do.Logs {
	return &Logs{
		ctx:   ctx,
		salt:  "",
		db:    mysql.DB,
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
		SetTime(logsReq.Time).
		Exec(l.ctx)
	if err != nil {
		err = errors.Wrap(err, "create logs failed")
		return err
	}
	return nil
}

func (l Logs) List(req *auth.LogsListReq) (list []*auth.LogsInfo, total int64, err error) {
	var predicates []predicate.Logs
	if req.Type != "" {
		predicates = append(predicates, logs.TypeEQ(req.Type))
	}
	if req.Method != "" {
		predicates = append(predicates, logs.MethodEQ(req.Method))
	}
	if req.Api != "" {
		predicates = append(predicates, logs.APIContains(req.Api))
	}
	//if req.Operator != "" {
	//	predicates = append(predicates, logs.OperatorContains(req.Operator))
	//}
	//if req.Success != nil {
	//	predicates = append(predicates, logs.SuccessEQ(*req.Success))
	//}
	logsData, err := l.db.Logs.Query().Where(predicates...).
		Offset(int((req.Page - 1) * req.PageSize)).
		Limit(int(req.PageSize)).
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
			Time:        v.Time,
			CreatedAt:   v.CreatedAt.Format(time.DateTime),
			UpdatedAt:   v.UpdatedAt.Format(time.DateTime),
		})
	}
	t, err := l.db.Logs.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return nil, 0, errors.Wrap(err, "query logsData count failed")
	}
	total = int64(t)
	return
}

func (l Logs) DeleteAll(ids []int64) error {
	_, err := l.db.Logs.Delete().Where(logs.IDIn(ids...)).Exec(l.ctx)
	if err != nil {
		return errors.Wrap(err, "delete logsData failed")
	}
	return nil
}
