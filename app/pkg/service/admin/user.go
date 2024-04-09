package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/predicate"
	"saas/pkg/db/ent/user"
	"saas/pkg/encrypt"
	"strconv"
	"time"
)

type User struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (u User) Create(req do.CreateOrUpdateUserReq) error {
	password, _ := encrypt.Crypt(req.Password)
	_, err := u.db.User.Create().
		SetAvatar(req.Avatar).
		SetRoleID(req.RoleID).
		SetMobile(req.Mobile).
		SetEmail(req.Email).
		SetStatus(req.Status).
		SetUsername(req.Username).
		SetPassword(password).
		SetNickname(req.Nickname).
		Save(u.ctx)
	if err != nil {
		err = errors.Wrap(err, "create user failed")
		return err
	}

	return nil
}

func (u User) Update(req do.CreateOrUpdateUserReq) error {
	password, _ := encrypt.Crypt(req.Password)
	_, err := u.db.User.Update().
		Where(user.IDEQ(req.ID)).
		SetAvatar(req.Avatar).
		SetRoleID(req.RoleID).
		SetMobile(req.Mobile).
		SetEmail(req.Email).
		SetStatus(req.Status).
		SetUsername(req.Username).
		SetPassword(password).
		SetNickname(req.Nickname).
		Save(u.ctx)
	if err != nil {
		err = errors.Wrap(err, "update user failed")
		return err
	}

	return nil
}

func (u User) ChangePassword(userID int64, oldPassword, newPassword string) error {
	// get user info
	targetUser, err := u.db.User.Query().Where(user.IDEQ(userID)).First(u.ctx)
	if err != nil {
		return errors.Wrap(err, "targetUser not found")
	}
	// check old password
	if ok := encrypt.VerifyPassword(oldPassword, targetUser.Password); !ok {
		err = errors.New("wrong old password")
		return err
	}
	// update password
	password, _ := encrypt.Crypt(newPassword)
	_, err = u.db.User.Update().Where(user.IDEQ(userID)).SetPassword(password).Save(u.ctx)

	return err
}

func (u User) UserInfo(id int64) (userInfo *do.UserInfo, err error) {
	userInfo = new(do.UserInfo)

	//errChan := make(chan error, 7)
	//defer close(errChan)
	//var wg sync.WaitGroup
	//wg.Add(7)
	//go func() {
	//
	//	//userInterface, exist := u.cache.Get("userInfo" + strconv.Itoa(int(id)))
	//	//if exist {
	//	//	if u, ok := userInterface.(*do.UserInfo); ok {
	//	//		errChan <- err
	//	//	}
	//	//}
	//
	//	wg.Done()
	//}()
	//wg.Wait()
	//select {
	//case result := <-errChan:
	//	return &do.UserInfo{}, result
	//default:
	//}

	userInterface, exist := u.cache.Get("userInfo" + strconv.Itoa(int(id)))
	if exist {
		if u, ok := userInterface.(*do.UserInfo); ok {
			return u, nil
		}
	}
	userEnt, err := u.db.User.Query().Where(user.IDEQ(id)).First(u.ctx)
	if err != nil {
		err = errors.Wrap(err, "get user failed")
		return userInfo, err
	}
	err = copier.Copy(&userInfo, &userEnt)
	if err != nil {
		err = errors.Wrap(err, "copy user info failed")
		return userInfo, err
	}

	roleInterface, exist := u.cache.Get("roleData" + strconv.Itoa(int(userInfo.RoleID)))
	if exist {
		if role, ok := roleInterface.(*ent.Role); ok {
			userInfo.RoleName = role.Name
			userInfo.RoleValue = role.Value
		}
	}
	u.cache.SetWithTTL("userInfo"+strconv.Itoa(int(userInfo.ID)), &userInfo, 1, 72*time.Hour)
	return
}

func (u User) List(req do.UserListReq) (userList []*do.UserInfo, total int, err error) {
	var predicates []predicate.User
	if req.Mobile != "" {
		predicates = append(predicates, user.MobileEQ(req.Mobile))
	}

	if req.Username != "" {
		predicates = append(predicates, user.UsernameContains(req.Username))
	}

	if req.Email != "" {
		predicates = append(predicates, user.EmailEQ(req.Email))
	}

	if req.Nickname != "" {
		predicates = append(predicates, user.NicknameContains(req.Nickname))
	}

	if req.RoleID != 0 {
		predicates = append(predicates, user.RoleIDEQ(req.RoleID))
	}

	users, err := u.db.User.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(u.ctx)
	if err != nil {
		err = errors.Wrap(err, "get user list failed")
		return userList, total, err
	}
	// copy to UserInfo struct
	err = copier.Copy(&userList, &users)
	if err != nil {
		err = errors.Wrap(err, "copy user info failed")
		return userList, 0, err
	}
	total, _ = u.db.User.Query().Where(predicates...).Count(u.ctx)
	return
}

func (u User) UpdateUserStatus(id int64, status int64) error {
	_, err := u.db.User.Update().Where(user.IDEQ(id)).SetStatus(status).Save(u.ctx)
	return err
}

func (u User) DeleteUser(id int64) error {
	_, err := u.db.User.Delete().Where(user.IDEQ(id)).Exec(u.ctx)
	return err
}

func (u User) UpdateProfile(req do.UpdateUserProfileReq) error {
	_, err := u.db.User.Update().
		Where(user.IDEQ(req.ID)).
		SetNickname(req.Nickname).
		SetEmail(req.Email).
		SetMobile(req.Mobile).
		SetAvatar(req.Avatar).
		Save(u.ctx)
	return err
}

func NewUser(ctx context.Context, c *app.RequestContext) do.User {
	return &User{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
