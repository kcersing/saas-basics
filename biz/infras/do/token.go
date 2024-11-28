package do

import "saas/idl_gen/model/token"

type Token interface {
	Create(req *token.TokenInfo) error
	Update(req *token.TokenInfo) error
	IsExistByUserID(userID int64) bool
	Delete(userID int64) error
	List(req *token.TokenListReq) (res []*token.TokenInfo, total int, err error)
}
