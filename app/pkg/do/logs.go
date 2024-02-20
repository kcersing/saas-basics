package do

type Logs interface {
	Create(logsReq *LogsInfo) error
	List(req *LogsListReq) (list []*LogsInfo, total int, err error)
	DeleteAll() error
}

type LogsInfo struct {
	Type        string
	Method      string
	Api         string
	Success     bool
	ReqContent  string
	RespContent string
	Ip          string
	UserAgent   string
	Operator    string
	Time        uint32
	CreatedAt   string
	UpdatedAt   string
}

type LogsListReq struct {
	Page     uint64
	limit    uint64
	Type     string
	Method   string
	Api      string
	Success  *bool
	Operator string
}
