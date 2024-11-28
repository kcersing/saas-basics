package do

import "saas/idl_gen/model/user"

type Login interface {
	Login(username, password string) (res *user.LoginResp, err error)
}
