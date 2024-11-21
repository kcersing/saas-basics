package do

import "rpc_gen/kitex_gen/user"

type Token interface {
	Create(req *user.TokenInfo) error
	Update(req *user.TokenInfo) error
	IsExistByUserID(userID int64) bool
	Delete(userID int64) error
	List(req *user.TokenListReq) (res []*user.TokenInfo, total int64, err error)
}
