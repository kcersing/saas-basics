package redis

import (
	"strconv"
)

const (
	tokenKey       = ":token"
	memberTokenKey = ":memberToken"
)

type (
	Token struct{}
)

func (t Token) addToken(id int64, token string) {
	add(rdToken, strconv.FormatInt(id, 10)+tokenKey, token)
}
func (t Token) addMemberToken(id int64, token string) {
	add(memberToken, strconv.FormatInt(id, 10)+memberTokenKey, token)
}
func (t Token) delToken(id int64, token string) {
	del(rdToken, strconv.FormatInt(id, 10)+tokenKey, token)
}
func (t Token) delMemberToken(id int64, token string) {
	del(memberToken, strconv.FormatInt(id, 10)+memberTokenKey, token)
}
func (t Token) checkToken(id int64) bool {
	return check(rdToken, strconv.FormatInt(id, 10)+tokenKey)
}
func (t Token) checkMemberToken(id int64) bool {
	return check(memberToken, strconv.FormatInt(id, 10)+memberTokenKey)
}
func (t Token) ExistToken(id int64, token string) bool {
	return exist(rdToken, strconv.FormatInt(id, 10)+tokenKey, token)
}
func (t Token) ExistMemberToken(id int64, token string) bool {
	return exist(memberToken, strconv.FormatInt(id, 10)+memberTokenKey, token)
}
func (t Token) getToken(id int64) []interface{} {
	return get(rdToken, strconv.FormatInt(id, 10)+tokenKey)
}
func (t Token) getMemberToken(id int64) []interface{} {
	return get(memberToken, strconv.FormatInt(id, 10)+memberTokenKey)
}
