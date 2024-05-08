package do

type Token interface {
	Create(req *TokenInfo) error
	Update(req *TokenInfo) error
	IsExistByUserID(userID int64) bool
	Delete(userID int64) error
	List(req *TokenListReq) (res []*TokenInfo, total int, err error)
}

type TokenInfo struct {
	ID        int64  `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	UserID    int64  `json:"userID"`
	UserName  string `json:"userName"`
	Token     string `json:"token"`
	Source    string `json:"source"`
	ExpiredAt string `json:"expiredAt"`
}

type TokenListReq struct {
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
	Username string `json:"username"`
	UserID   int64  `json:"userID"`
}
