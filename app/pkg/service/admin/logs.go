package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/pkg/errors"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
)

type Logs struct {
	ctx  context.Context
	c    *app.RequestContext
	salt string
	db   *ent.Client
}

func NewLogs(ctx context.Context, c *app.RequestContext) do.Logs {
	return &Logs{
		ctx:  ctx,
		c:    c,
		salt: config.GlobalServerConfig.MysqlInfo.Salt,
		db:   infras.DB,
	}
}

func (l Logs) Create(logsReq *do.LogsInfo) error {
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

func (l Logs) List(req *do.LogsListReq) (list []*do.LogsInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (l Logs) DeleteAll() error {
	_, err := l.db.Logs.Delete().Exec(l.ctx)
	if err != nil {
		return errors.Wrap(err, "delete logsData failed")
	}
	return nil
}
