package service

import (
	"saas/idl_gen/model/community"
)

func (c Community) CreateParticipant(req community.ParticipantInfo) error {
	//TODO implement me
	panic("implement me")
}

func (c Community) UpdateParticipant(req community.ParticipantInfo) error {
	//TODO implement me
	panic("implement me")
}

func (c Community) DeleteParticipant(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (c Community) ParticipantInfo(id int64) (resp *community.ParticipantInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (c Community) ParticipantList(req community.ParticipantListReq) (resp []*community.ParticipantInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (c Community) UpdateParticipantStatus(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func (c Community) ParticipantListExport(req community.ParticipantListReq) (string, error) {
	//TODO implement me
	panic("implement me")
}
