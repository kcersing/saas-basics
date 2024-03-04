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
	ID          uint64 `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Name        string `json:"name,omitempty"`
	Status      uint64 `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}

type DictListReq struct {
	Title    string `json:"title,omitempty"`
	Name     string `json:"name,omitempty"`
	Page     uint64 `json:"page,omitempty"`
	PageSize uint64 `json:"pageSize,omitempty"`
}

type DictionaryDetail struct {
	ID        uint64 `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Key       string `json:"key,omitempty"`
	Value     string `json:"value,omitempty"`
	Status    uint64 `json:"status,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	ParentID  uint64 `json:"parentID,omitempty"`
}
