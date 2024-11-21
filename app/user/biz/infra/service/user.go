package service

import (
	"common/utils"
	"common/utils/minio"
	"context"
	user2 "user/biz/dal/mysql/ent/user"

	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"rpc_gen/kitex_gen/user"
	"strconv"
	"time"
	"user/biz/dal/cache"
	"user/biz/dal/mysql"
	"user/biz/dal/mysql/ent"
	"user/biz/dal/mysql/ent/predicate"
	"user/biz/infra/do"
)

type User struct {
	ctx   context.Context
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func NewUser(ctx context.Context) do.User {
	return &User{
		ctx:   ctx,
		salt:  "",
		db:    mysql.DB,
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
	if !userEnt.Birthday.IsZero() {
		info.Age = int64(time.Now().Sub(userEnt.Birthday).Hours() / 24 / 365)
	}
	info.Birthday = userEnt.Birthday.Format(time.DateOnly)
	u.cache.SetWithTTL("userInfo"+strconv.Itoa(int(info.Id)), &info, 1, 1*time.Hour)
	return
}

func (u User) Create(req *user.CreateOrUpdateUserReq) error {
	password, _ := utils.Crypt(*req.Password)
	var gender int64
	if *req.Gender == "女性" {
		gender = 0
	} else if *req.Gender == "男性" {
		gender = 1
	} else {
		gender = 2
	}
	parsedTime, _ := time.Parse(time.DateOnly, *req.Birthday)

	tx, err := u.db.Tx(u.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}
	noe, err := tx.User.Create().
		SetAvatar(*req.Avatar).
		SetMobile(*req.Mobile).
		SetEmail(*req.Email).
		SetStatus(*req.Status).
		SetUsername(*req.Mobile).
		SetPassword(password).
		SetNickname(*req.Name).
		SetBirthday(parsedTime).
		SetGender(gender).
		SetWecom(*req.Wecom).
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

func (u User) Update(req *user.CreateOrUpdateUserReq) error {

	var gender int64
	if *req.Gender == "女性" {
		gender = 0
	} else if *req.Gender == "男性" {
		gender = 1
	} else {
		gender = 2
	}
	parsedTime, _ := time.Parse(time.DateOnly, *req.Birthday)

	password, _ := utils.Crypt(*req.Password)
	_, err := u.db.User.Update().
		Where(user2.IDEQ(*req.Id)).
		SetAvatar(*req.Avatar).
		SetMobile(*req.Mobile).
		SetEmail(*req.Email).
		SetStatus(*req.Status).
		SetUsername(*req.Mobile).
		SetPassword(password).
		SetNickname(*req.Name).
		SetBirthday(parsedTime).
		SetGender(gender).
		SetStatus(1).
		SetWecom(*req.Wecom).
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
	password, _ := utils.Crypt(newPassword)
	_, err := u.db.User.Update().Where(user2.IDEQ(userID)).SetPassword(password).Save(u.ctx)

	return err
}

func (u User) List(req *user.UserListReq) (userList []*user.UserInfo, total int64, err error) {
	var predicates []predicate.User
	if *req.Mobile != "" {
		predicates = append(predicates, user2.MobileEQ(*req.Mobile))
	}

	if *req.Username != "" {
		predicates = append(predicates, user2.UsernameContains(*req.Username))
	}

	if *req.Email != "" {
		predicates = append(predicates, user2.EmailEQ(*req.Email))
	}

	if *req.Nickname != "" {
		predicates = append(predicates, user2.NicknameContains(*req.Nickname))
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

	tl, _ := u.db.User.Query().Where(predicates...).Count(u.ctx)
	total = int64(tl)
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
