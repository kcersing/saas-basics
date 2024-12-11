package do

import "saas/idl_gen/model/bootcamp"

type Bootcamp interface {
	CreateBootcamp(req bootcamp.BootcampInfo) error
	UpdateBootcamp(req bootcamp.BootcampInfo) error
	DeleteBootcamp(id int64) error
	BootcampList(req bootcamp.BootcampListReq) (resp []*bootcamp.BootcampInfo, total int, err error)
	BootcampInfo(id int64) (resp *bootcamp.BootcampInfo, err error)
	UpdateBootcampStatus(ID int64, status int64) error
	UpdateBootcampShow(ID int64, status int64) error
	DelBootcamp(ID int64) error

	CreateParticipant(req bootcamp.ParticipantInfo) error
	UpdateParticipant(req bootcamp.ParticipantInfo) error
	DeleteParticipant(id int64) error
	ParticipantInfo(id int64) (resp *bootcamp.ParticipantInfo, err error)
	ParticipantList(req bootcamp.ParticipantListReq) (resp []*bootcamp.ParticipantInfo, total int, err error)
	UpdateParticipantStatus(ID int64, status int64) error
	ParticipantListListExport(req bootcamp.ParticipantListReq) (string, error)
}
