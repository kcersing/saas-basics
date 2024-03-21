package do

type MemberProduct interface {
	Create(req MemberProductInfo) error
	Update(req MemberProductInfo) error
	Delete(id int64) error
	List(req MemberProductListReq) (resp []*MemberProductInfo, total int, err error)
	UpdateStatus(ID int64, status int64) error
	InfoByID(ID int64) (roleInfo *MemberProductInfo, err error)
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
	Status   int64  `json:"status,omitempty"`
}
type MemberPropertyInfo struct {
	ID        int64   `json:"id,omitempty"`
	MemberId  int64   `json:"member_id,omitempty"`
	ProductId int64   `json:"productId,omitempty"`
	Name      string  `json:"name,omitempty"`
	Price     float64 `json:"price,omitempty"`
	Duration  int64   `json:"duration,omitempty"`
	Length    int64   `json:"length,omitempty"`
	Count     int64   `json:"count,omitempty"`
	Type      string  `json:"type,omitempty"`
	Data      string  `json:"data,omitempty"`
	Status    int64   `json:"status,omitempty"`
}
