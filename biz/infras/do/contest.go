package do

import "saas/idl_gen/model/contest"

type Contest interface {
	CreateContest(req contest.ContestInfo) error
	UpdateContest(req contest.ContestInfo) error
	DeleteContest(id int64) error
	ContestList(req contest.ContestListReq) (resp []*contest.ContestInfo, total int, err error)

	CreateParticipant(req contest.ParticipantInfo) error
	UpdateParticipant(req contest.ParticipantInfo) error
	DeleteParticipant(id int64) error
	ParticipantList(req contest.ParticipantListReq) (resp []*contest.ParticipantInfo, total int, err error)
}
