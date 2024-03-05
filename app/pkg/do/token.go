package do

type Token interface {
	Create(req *TokenInfo) error
	Update(req *TokenInfo) error
	IsExistByUserID(userID int64) bool
	Delete(userID int64) error
	List(req *TokenListReq) (res []*TokenInfo, total int, err error)
}

type TokenInfo struct {
	ID        int64  `json:"id,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	UserID    int64  `json:"userID,omitempty"`
	UserName  string `json:"userName,omitempty"`
	Token     string `json:"token,omitempty"`
	Source    string `json:"source,omitempty"`
	ExpiredAt string `json:"expiredAt,omitempty"`
}

type TokenListReq struct {
	Page     int64  `json:"page,omitempty"`
	PageSize int64  `json:"pageSize,omitempty"`
	Username string `json:"username,omitempty"`
	UserID   int64  `json:"userID,omitempty"`
}
