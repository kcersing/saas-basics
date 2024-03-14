package do

type MemberProduct interface {
	CreateProperty(req PropertyInfo) error
	UpdateProperty(req PropertyInfo) error
	DeleteProperty(id int64) error
	PropertyList(req ProductListReq) (resp []*PropertyInfo, total int, err error)

	Create(req ProductInfo) error
	Update(req ProductInfo) error
	Delete(id int64) error
	List(req ProductListReq) (resp []*ProductInfo, total int, err error)
	UpdateStatus(ID int64, status int8) error
	InfoByID(ID int64) (roleInfo *ProductInfo, err error)
}

type MemberPropertyInfo struct {
	ID        int64   `json:"id,omitempty"`
	MemberId  int64   `json:"member_id,omitempty"`
	ProductId int64   `json:"productId,omitempty"`
	SpuName   string  `json:"spuName,omitempty"`
	SpuPrice  float64 `json:"spuPrice,omitempty"`
	Duration  int64   `json:"duration,omitempty"`
	Length    int64   `json:"length,omitempty"`
	Count     int64   `json:"count,omitempty"`
	Type      string  `json:"type,omitempty"`
	Data      string  `json:"data,omitempty"`
	Status    int64   `json:"status,omitempty"`
}

type MemberProductInfo struct {
	ID          int64          `json:"id,omitempty"`
	MemberId    int64          `json:"member_id,omitempty"`
	Name        string         `json:"name,omitempty"`
	Pic         string         `json:"pic,omitempty"`
	Description string         `json:"description,omitempty"`
	Property    []PropertyInfo `json:"property,omitempty"`
	Price       float64        `json:"price,omitempty"`
	Stock       int64          `json:"stock,omitempty"`
	Status      int64          `json:"status,omitempty"`
}
type MemberProductListReq struct {
	Page     int64  `json:"page,omitempty"`
	PageSize int64  `json:"pageSize,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	Name     string `json:"name,omitempty"`
	SpuName  string `json:"spuName,omitempty"`
	Status   int64  `json:"status,omitempty"`
}