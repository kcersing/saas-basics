package do

type EntryLogs interface {
	Create(logsReq *EntryLogsInfo) error
	List(req *EntryLogsListReq) (list []*EntryLogsInfo, total int, err error)
}

type CreateCreateLogs struct {
	MemberProductId  int64 `json:"member_product_id"`
	MemberPropertyId int64 `json:"member_property_id"`
	EntryTime        int64 `json:"entry_time"`
	LeavingTime      int64 `json:"leaving_time"`
	MemberId         int64 `json:"member_id"`
	UserId           int64 `json:"user_id"`
	VenueId          int64 `json:"venue_id"`
}

type EntryLogsInfo struct {
	Id               string `json:"id"`
	MemberProductId  int64  `json:"member_product_id"`
	MemberPropertyId int64  `json:"member_property_id"`
	EntryTime        int64  `json:"entry_time"`
	LeavingTime      int64  `json:"leaving_time"`
	MemberId         int64  `json:"member_id"`
	UserId           int64  `json:"user_id"`
	VenueId          int64  `json:"venue_id"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`

	VenueName          string `json:"venue_name"`
	UserName           string `json:"user_name"`
	MemberName         string `json:"member_name"`
	MemberProductName  string `json:"member_product_name"`
	MemberPropertyName string `json:"member_property_name"`
}

type EntryLogsListReq struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`

	MemberProductId  int64  `json:"member_product_id"`
	MemberPropertyId int64  `json:"member_property_id"`
	EntryTime        string `json:"entry_time"`
	LeavingTime      string `json:"leaving_time"`
	MemberId         int64  `json:"member_id"`
	UserId           int64  `json:"user_id"`
	VenueId          int64  `json:"venue_id"`
}
