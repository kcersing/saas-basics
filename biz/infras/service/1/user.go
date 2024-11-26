package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/admin/pkg/minio"
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

func (u User) SetDefaultVenue(id, venueId int64) error {
	_, err := u.db.User.Update().
		Where(user.IDEQ(id)).
		SetDefaultVenueID(venueId).
		Save(u.ctx)

	if err != nil {
		err = errors.Wrap(err, "update DefaultVenue  ID   failed")
		return err
	}

	return nil
}

func (u User) SetRole(id, roleID int64) error {
	_, err := u.db.User.Update().
		Where(user.IDEQ(id)).
		SetRoleID(roleID).
		Save(u.ctx)

	if err != nil {
		err = errors.Wrap(err, "update user role failed")
		return err
	}

	return nil
}

func (u User) Info(id int64) (info *do.UserInfo, err error) {
	info = new(do.UserInfo)
	userInterface, exist := u.cache.Get("userInfo" + strconv.Itoa(int(id)))
	if exist {
		if u, ok := userInterface.(*do.UserInfo); ok {
			return u, nil
		}
	}
	userEnt, err := u.db.User.Query().Where(user.IDEQ(id)).First(u.ctx)
	if err != nil {
		err = errors.Wrap(err, "get user failed")
		return info, err
	}
	err = copier.Copy(&info, &userEnt)
	if err != nil {
		err = errors.Wrap(err, "copy user info failed")
		return info, err
	}

	roleInterface, exist := u.cache.Get("roleData" + strconv.Itoa(int(info.RoleID)))
	if exist {
		if role, ok := roleInterface.(*ent.Role); ok {
			info.RoleName = role.Name
			info.RoleValue = role.Value
		}
	}
	info.Avatar = minio.URLconvert(u.ctx, u.c, info.Avatar)
	if userEnt.Gender == 0 {
		info.Gender = "女性"
	} else if userEnt.Gender == 1 {
		info.Gender = "男性"
	} else {
		info.Gender = "保密"
	}
	if !userEnt.Birthday.IsZero() {
		info.Age = int64(time.Now().Sub(userEnt.Birthday).Hours() / 24 / 365)
	}
	info.Birthday = userEnt.Birthday.Format(time.DateOnly)
	u.cache.SetWithTTL("userInfo"+strconv.Itoa(int(info.ID)), &info, 1, 1*time.Hour)
	return
}

func (u User) Create(req do.CreateOrUpdateUserReq) error {
	password, _ := encrypt.Crypt(req.Password)
	var gender int64
	if req.Gender == "女性" {
		gender = 0
	} else if req.Gender == "男性" {
		gender = 1
	} else {
		gender = 2
	}
	parsedTime, _ := time.Parse(time.DateOnly, req.Birthday)

	tx, err := u.db.Tx(u.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}
	noe, err := tx.User.Create().
		SetAvatar(req.Avatar).
		SetMobile(req.Mobile).
		SetEmail(req.Email).
		SetStatus(req.Status).
		SetUsername(req.Mobile).
		SetPassword(password).
		SetNickname(req.Name).
		SetBirthday(parsedTime).
		SetGender(gender).
		SetWecom(req.Wecom).
		Save(u.ctx)

	if err != nil {
		err = rollback(tx, errors.Wrap(err, "create user failed"))
		return err
	}
	_, err = tx.Face.Create().
		SetUserFaces(noe).
		Save(u.ctx)

	if err != nil {
		err = rollback(tx, errors.Wrap(err, "create Face failed"))
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (u User) Update(req do.CreateOrUpdateUserReq) error {

	var gender int64
	if req.Gender == "女性" {
		gender = 0
	} else if req.Gender == "男性" {
		gender = 1
	} else {
		gender = 2
	}
	parsedTime, _ := time.Parse(time.DateOnly, req.Birthday)

	password, _ := encrypt.Crypt(req.Password)
	_, err := u.db.User.Update().
		Where(user.IDEQ(req.ID)).
		SetAvatar(req.Avatar).
		SetMobile(req.Mobile).
		SetEmail(req.Email).
		SetStatus(req.Status).
		SetUsername(req.Mobile).
		SetPassword(password).
		SetNickname(req.Name).
		SetBirthday(parsedTime).
		SetGender(gender).
		SetStatus(1).
		SetWecom(req.Wecom).
		Save(u.ctx)

	if err != nil {
		err = errors.Wrap(err, "update user failed")
		return err
	}

	return nil
}

func (u User) ChangePassword(userID int64, newPassword string) error {
	////get user info
	//targetUser, err := u.db.User.Query().Where(user.IDEQ(userID)).First(u.ctx)
	//if err != nil {
	//	return errors.Wrap(err, "targetUser not found")
	//}
	//// check old password
	//if ok := encrypt.VerifyPassword(oldPassword, targetUser.Password); !ok {
	//	err = errors.New("wrong old password")
	//	return err
	//}
	// update password
	password, _ := encrypt.Crypt(newPassword)
	_, err := u.db.User.Update().Where(user.IDEQ(userID)).SetPassword(password).Save(u.ctx)

	return err
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

	for _, v := range userList {
		v.Avatar = minio.URLconvert(u.ctx, u.c, v.Avatar)
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
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
