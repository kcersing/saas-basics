package do

import "saas/idl_gen/model/community"

type Community interface {
	CreateCommunity(req community.CommunityInfo) error
	UpdateCommunity(req community.CommunityInfo) error
	DeleteCommunity(id int64) error
	CommunityList(req community.CommunityListReq) (resp []*community.CommunityInfo, total int, err error)
	CommunityInfo(id int64) (resp *community.CommunityInfo, err error)
	UpdateCommunityStatus(ID int64, status int64) error
	UpdateCommunityShow(ID int64, status int64) error
	DelCommunity(ID int64) error

	CreateParticipant(req community.ParticipantInfo) error
	UpdateParticipant(req community.ParticipantInfo) error
	DeleteParticipant(id int64) error
	ParticipantInfo(id int64) (resp *community.ParticipantInfo, err error)
	ParticipantList(req community.ParticipantListReq) (resp []*community.ParticipantInfo, total int, err error)
	UpdateParticipantStatus(ID int64, status int64) error
	ParticipantListExport(req community.ParticipantListReq) (string, error)
}
