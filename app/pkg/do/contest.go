package do

type Contest interface {
	CreateContest(req ContestInfo) error
	UpdateContest(req ContestInfo) error
	DeleteContest(id int64) error
	ContestList(req ContestListReq) (resp []*ContestInfo, total int, err error)

	CreateParticipant(req ParticipantInfo) error
	UpdateParticipant(req ParticipantInfo) error
	DeleteParticipant(id int64) error
	ParticipantList(req ParticipantListReq) (resp []*ParticipantInfo, total int, err error)
}

type ParticipantInfo struct {
	ContestId int64  `json:"contestId"`
	Name      string `json:"Name"`
	Mobile    string `json:"mobile"`
	Fields    string `json:"fields"`
}
type ParticipantListReq struct {
}
type ContestInfo struct {
	Id          int64   `json:"id"`
	Name        string  `json:"Name"`
	SignNumber  int64   `json:"signNumber"`
	SignStartAt string  `json:"signStartAt"`
	SignEndAt   string  `json:"signEndAt"`
	Number      int64   `json:"number"`
	StartAt     string  `json:"startAt"`
	EndAt       string  `json:"endAt"`
	Pic         string  `json:"pic"`
	Sponsor     string  `json:"sponsor"`
	Fee         float64 `json:"fee"`
	IsCancel    int64   `json:"isCancel"`
	CancelTime  int64   `json:"cancelTime"`
	Detail      string  `json:"detail"`
	SignFields  string  `json:"signFields"`
	CreatedAt   int64   `json:"createdAt"`
	UpdatedAt   int64   `json:"updatedAt"`
}
type ContestListReq struct {
}
