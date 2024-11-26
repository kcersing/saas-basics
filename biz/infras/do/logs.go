package do

type Logs interface {
	Create(logsReq *LogsInfo) error
	List(req *LogsListReq) (list []*LogsInfo, total int, err error)
	DeleteAll() error
}

type LogsInfo struct {
	Type        string `json:"type"`
	Method      string `json:"method"`
	Api         string `json:"api"`
	Success     bool   `json:"success"`
	ReqContent  string `json:"reqContent"`
	RespContent string `json:"respContent"`
	Ip          string `json:"ip"`
	UserAgent   string `json:"userAgent"`
	Operator    string `json:"operator"`
	Time        int32  `json:"time"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type LogsListReq struct {
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
	Type     string `json:"type"`
	Method   string `json:"method"`
	Api      string `json:"api"`
	Success  *bool  `json:"success"`
	Operator string `json:"operator"`
}
