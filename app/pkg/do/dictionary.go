package do

type Dictionary interface {
	Create(req *DictionaryInfo) error
	Update(req *DictionaryInfo) error
	Delete(id int64) error
	List(req *DictListReq) (list []*DictionaryInfo, total int, err error)
	CreateDetail(req *DictionaryDetail) error
	UpdateDetail(req *DictionaryDetail) error
	DeleteDetail(id int64) error
	DetailListByDict(req *DetailListReq) (list []*DictionaryDetail, total int64, err error)
	DetailByDictNameAndKey(dictName, key string) (detail *DictionaryDetail, err error)
}

type DictionaryInfo struct {
	ID          int64  `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Name        string `json:"name,omitempty"`
	Status      int64  `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}

type DictListReq struct {
	Title    string `json:"title,omitempty"`
	Name     string `json:"name,omitempty"`
	Page     int64  `json:"page,omitempty"`
	PageSize int64  `json:"pageSize,omitempty"`
}

type DetailListReq struct {
	Name         string `json:"name,omitempty"`
	DictionaryId int64  `json:"dictionaryId,omitempty"`
}

type DictionaryDetail struct {
	ID        int64  `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Key       string `json:"key,omitempty"`
	Value     string `json:"value,omitempty"`
	Status    int64  `json:"status,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	ParentID  int64  `json:"parentID,omitempty"`
}
