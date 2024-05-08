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
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Name        string `json:"name"`
	Status      int64  `json:"status"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type DictListReq struct {
	Title    string `json:"title"`
	Name     string `json:"name"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
}

type DetailListReq struct {
	Name         string `json:"name"`
	DictionaryId int64  `json:"dictionaryId"`
}

type DictionaryDetail struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	Status    int64  `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	ParentID  int64  `json:"parentID"`
}
