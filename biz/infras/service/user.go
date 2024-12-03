package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/user"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/predicate"
	user2 "saas/pkg/db/ent/user"
	"saas/pkg/encrypt"
	"saas/pkg/minio"
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

func NewUser(ctx context.Context, c *app.RequestContext) do.User {

	return &User{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}

func (u User) SetDefaultVenue(id, venueId int64) error {
	_, err := u.db.User.Update().
		Where(user2.IDEQ(id)).
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
		Where(user2.IDEQ(id)).
		SetRoleID(roleID).
		Save(u.ctx)

	if err != nil {
		err = errors.Wrap(err, "update user role failed")
		return err
	}

	return nil
}

func (u User) Info(id int64) (info *user.UserInfo, err error) {
	info = new(user.UserInfo)
	userInterface, exist := u.cache.Get("userInfo" + strconv.Itoa(int(id)))
	if exist {
		if u, ok := userInterface.(*user.UserInfo); ok {
			return u, nil
		}
	}
	userEnt, err := u.db.User.Query().Where(user2.IDEQ(id)).First(u.ctx)
	if err != nil {
		err = errors.Wrap(err, "get user failed")
		return info, err
	}
	err = copier.Copy(&info, &userEnt)
	if err != nil {
		err = errors.Wrap(err, "copy user info failed")
		return info, err
	}

	roleInterface, exist := u.cache.Get("roleData" + strconv.Itoa(int(info.RoleId)))
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
	//if !userEnt.Birthday.IsZero() {
	//	info.Age = int64(time.Now().Sub(userEnt.Birthday).Hours() / 24 / 365)
	//}
	//info.Birthday = userEnt.Birthday.Format(time.DateOnly)
	u.cache.SetWithTTL("userInfo"+strconv.Itoa(int(info.Id)), &info, 1, 1*time.Hour)
	return
}

func (u User) Create(req user.CreateOrUpdateUserReq) error {
	password, _ := encrypt.Crypt(*req.Password)
	var gender int64
	if *req.Gender == "女性" {
		gender = 0
	} else if *req.Gender == "男性" {
		gender = 1
	} else {
		gender = 2
	}
	//parsedTime, _ := time.Parse(time.DateOnly, req.Birthday)

	tx, err := u.db.Tx(u.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}
	_, err = tx.User.Create().
		SetAvatar(*req.Avatar).
		SetMobile(*req.Mobile).
		SetJobTime(*req.JobTime).
		SetStatus(*req.Status).
		SetUsername(*req.Username).
		SetPassword(password).
		SetName(*req.Name).
		//SetFunctions(*req.Functions).
		SetGender(gender).
		SetRoleID(*req.RoleId).
		SetDetail(*req.Detail).
		Save(u.ctx)

	if err != nil {
		err = rollback(tx, errors.Wrap(err, "create user failed"))
		return err
	}
	//_, err = tx.Face.Create().
	//	SetUserFaces(noe).
	//	Save(u.ctx)

	if err != nil {
		err = rollback(tx, errors.Wrap(err, "create Face failed"))
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (u User) Update(req user.CreateOrUpdateUserReq) error {

	var gender int64
	if *req.Gender == "女性" {
		gender = 0
	} else if *req.Gender == "男性" {
		gender = 1
	} else {
		gender = 2
	}
	//parsedTime, _ := time.Parse(time.DateOnly, req.Birthday)

	password, _ := encrypt.Crypt(*req.Password)
	_, err := u.db.User.Update().
		Where(user2.IDEQ(*req.ID)).
		SetAvatar(*req.Avatar).
		SetMobile(*req.Mobile).
		SetJobTime(*req.JobTime).
		SetStatus(*req.Status).
		SetUsername(*req.Username).
		SetPassword(password).
		SetName(*req.Name).
		//SetFunctions(*req.Functions).
		SetGender(gender).
		SetRoleID(*req.RoleId).
		SetDetail(*req.Detail).
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
	_, err := u.db.User.Update().Where(user2.IDEQ(userID)).SetPassword(password).Save(u.ctx)

	return err
}

func (u User) List(req user.UserListReq) (userList []*user.UserInfo, total int, err error) {
	var predicates []predicate.User
	if *req.Mobile != "" {
		predicates = append(predicates, user2.MobileEQ(*req.Mobile))
	}
	//	Name     *string `thrift:"name,4,optional" form:"name" json:"name" query:"name"`
	//	JobTime  *int64  `thrift:"jobTime,5,optional" form:"jobTime" json:"jobTime" query:"jobTime"`
	//	Mobile   *string `thrift:"mobile,6,optional" form:"mobile" json:"mobile" query:"mobile"`
	//	RoleId   *int64  `thrift:"roleId,7,optional" form:"roleId" json:"roleId" query:"roleId"`
	if *req.JobTime != 0 {
		predicates = append(predicates, user2.JobTime(*req.JobTime))
	}
	if *req.Mobile != "" {
		predicates = append(predicates, user2.Mobile(*req.Mobile))
	}

	if *req.Name != "" {
		predicates = append(predicates, user2.NameContains(*req.Name))
	}

	if *req.RoleId != 0 {
		predicates = append(predicates, user2.RoleIDEQ(*req.RoleId))
	}

	users, err := u.db.User.Query().Where(predicates...).
		Offset(int(*req.Page-1) * int(*req.PageSize)).
		Limit(int(*req.PageSize)).All(u.ctx)
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
	_, err := u.db.User.Update().Where(user2.IDEQ(id)).SetStatus(status).Save(u.ctx)
	return err
}

func (u User) DeleteUser(id int64) error {
	_, err := u.db.User.Delete().Where(user2.IDEQ(id)).Exec(u.ctx)
	return err
}
