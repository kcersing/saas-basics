package do

type Logs interface {
	Create(logsReq *LogsInfo) error
	List(req *LogsListReq) (list []*LogsInfo, total int, err error)
	DeleteAll() error
}

type LogsInfo struct {
	Type        string `json:"type,omitempty"`
	Method      string `json:"method,omitempty"`
	Api         string `json:"api,omitempty"`
	Success     bool   `json:"success,omitempty"`
	ReqContent  string `json:"reqContent,omitempty"`
	RespContent string `json:"respContent,omitempty"`
	Ip          string `json:"ip,omitempty"`
	UserAgent   string `json:"userAgent,omitempty"`
	Operator    string `json:"operator,omitempty"`
	Time        int32  `json:"time,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}

type LogsListReq struct {
	Page     int64  `json:"page,omitempty"`
	Limit    int64  `json:"limit,omitempty"`
	Type     string `json:"type,omitempty"`
	Method   string `json:"method,omitempty"`
	Api      string `json:"api,omitempty"`
	Success  *bool  `json:"success,omitempty"`
	Operator string `json:"operator,omitempty"`
}
