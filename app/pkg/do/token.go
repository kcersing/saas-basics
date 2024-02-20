package do

type Token interface {
	Create(req *TokenInfo) error
	Update(req *TokenInfo) error
	IsExistByUserID(userID uint64) bool
	Delete(userID uint64) error
	List(req *TokenListReq) (res []*TokenInfo, total int, err error)
}

type TokenInfo struct {
	ID        uint64
	CreatedAt string
	UpdatedAt string
	UserID    uint64
	UserName  string
	Token     string
	Source    string
	ExpiredAt string
}

type TokenListReq struct {
	Page     uint64
	PageSize uint64
	Username string
	UserID   uint64
}
