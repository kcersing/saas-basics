package do

type Dictionary interface {
	Create(req *DictionaryInfo) error
	Update(req *DictionaryInfo) error
	Delete(id uint64) error
	List(req *DictListReq) (list []*DictionaryInfo, total int, err error)
	CreateDetail(req *DictionaryDetail) error
	UpdateDetail(req *DictionaryDetail) error
	DeleteDetail(id uint64) error
	DetailListByDictName(dictName string) (list []*DictionaryDetail, total uint64, err error)
	DetailByDictNameAndKey(dictName, key string) (detail *DictionaryDetail, err error)
}

type DictionaryInfo struct {
	ID          uint64
	Title       string
	Name        string
	Status      uint64
	Description string
	CreatedAt   string
	UpdatedAt   string
}

type DictListReq struct {
	Title    string
	Name     string
	Page     uint64
	PageSize uint64
}

type DictionaryDetail struct {
	ID        uint64
	Title     string
	Key       string
	Value     string
	Status    uint64
	CreatedAt string
	UpdatedAt string
	ParentID  uint64
}
