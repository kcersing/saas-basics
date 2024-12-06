package do

import "saas/idl_gen/model/contest"

type Contest interface {
	CreateContest(req contest.ContestInfo) error
	UpdateContest(req contest.ContestInfo) error
	DeleteContest(id int64) error
	ContestList(req contest.ContestListReq) (resp []*contest.ContestInfo, total int, err error)
	ContestInfo(id int64) (resp *contest.ContestInfo, err error)
	UpdateContestStatus(ID int64, status int64) error
	UpdateContestShow(ID int64, status int64) error

	CreateParticipant(req contest.ParticipantInfo) error
	UpdateParticipant(req contest.ParticipantInfo) error
	DeleteParticipant(id int64) error
	ParticipantInfo(id int64) (resp *contest.ParticipantInfo, err error)
	ParticipantList(req contest.ParticipantListReq) (resp []*contest.ParticipantInfo, total int, err error)
	UpdateParticipantStatus(ID int64, status int64) error
}
