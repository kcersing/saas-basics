package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/contest"
	"saas/pkg/db/ent"
)

type Contest struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (c Contest) CreateContest(req contest.ContestInfo) error {
	//TODO implement me
	panic("implement me")
}

func (c Contest) UpdateContest(req contest.ContestInfo) error {
	//TODO implement me
	panic("implement me")
}

func (c Contest) DeleteContest(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (c Contest) ContestList(req contest.ContestListReq) (resp []*contest.ContestInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (c Contest) ContestInfo(id int64) (resp *contest.ContestInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (c Contest) UpdateContestStatus(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func (c Contest) CreateParticipant(req contest.ParticipantInfo) error {
	//TODO implement me
	panic("implement me")
}

func (c Contest) UpdateParticipant(req contest.ParticipantInfo) error {
	//TODO implement me
	panic("implement me")
}

func (c Contest) DeleteParticipant(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (c Contest) ParticipantInfo(id int64) (resp *contest.ParticipantInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (c Contest) ParticipantList(req contest.ParticipantListReq) (resp []*contest.ParticipantInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (c Contest) UpdateParticipantStatus(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func NewContest(ctx context.Context, c *app.RequestContext) do.Contest {
	return &Contest{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
