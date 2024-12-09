package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/dal/db/ent/role"
	user2 "saas/biz/dal/db/ent/user"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/dictionary"
	"saas/idl_gen/model/user"
	"saas/pkg/encrypt"
	"strings"
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
	//userInterface, exist := u.cache.Get("userInfo" + strconv.Itoa(int(id)))
	//if exist {
	//	if u, ok := userInterface.(*user.UserInfo); ok {
	//		return u, nil
	//	}
	//}
	userEnt, err := u.db.User.Query().Where(user2.IDEQ(id)).First(u.ctx)
	if err != nil {
		err = errors.Wrap(err, "get user failed")
		return info, err
	}
	info = u.entUserInfo(*userEnt)
	//if !userEnt.Birthday.IsZero() {
	//	info.Age = int64(time.Now().Sub(userEnt.Birthday).Hours() / 24 / 365)
	//}
	//info.Birthday = userEnt.Birthday.Format(time.DateTime)
	//u.cache.SetWithTTL("userInfo"+strconv.Itoa(int(info.Id)), &info, 1, 1*time.Hour)
	return
}

func (u User) Create(req user.CreateOrUpdateUserReq) error {

	ok, _ := u.db.User.Query().Where(user2.Username(req.Username)).Exist(u.ctx)
	if ok {
		return errors.New("账号重复！")
	}
	ok, _ = u.db.User.Query().Where(user2.Mobile(req.Mobile)).Exist(u.ctx)
	if ok {
		return errors.New("手机号重复！")
	}

	password, _ := encrypt.Crypt(req.Password)
	var gender int64
	if req.Gender == "女性" {
		gender = 0
	} else if req.Gender == "男性" {
		gender = 1
	} else {
		gender = 2
	}
	//parsedTime, _ := time.Parse(time.DateTime, req.Birthday)
	functions := strings.Join(req.Functions, ",")

	tx, err := u.db.Tx(u.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}
	one, err := tx.User.Create().
		SetAvatar(req.Avatar).
		SetMobile(req.Mobile).
		SetJobTime(req.JobTime).
		SetStatus(req.Status).
		SetUsername(req.Username).
		SetPassword(password).
		SetName(req.Name).
		SetFunctions(functions).
		SetGender(gender).
		SetRoleID(req.RoleId).
		SetDetail(req.Detail).
		//SetTags().
		Save(u.ctx)

	if err != nil {
		err = rollback(tx, errors.Wrap(err, "create user failed"))
		return err
	}
	//	err = tx.Role.UpdateOneID(roleID).AddMenuIDs(menuIDs...).Exec(a.ctx)
	tx.User.UpdateOneID(one.ID).AddTagIDs(req.UserTags...).Exec(u.ctx)
	//_, err = tx.Face.Create().
	//	SetUserFaces(noe).
	//	Save(u.ctx)

	//if err != nil {
	//	err = rollback(tx, errors.Wrap(err, "create Face failed"))
	//	return err
	//}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (u User) Update(req user.CreateOrUpdateUserReq) error {

	var gender int64
	if req.Gender == "女性" {
		gender = 0
	} else if req.Gender == "男性" {
		gender = 1
	} else {
		gender = 2
	}
	//parsedTime, _ := time.Parse(time.DateTime, req.Birthday)

	//password, _ := encrypt.Crypt(req.Password)

	functions := strings.Join(req.Functions, ",")
	//*req.UserTags...
	_, err := u.db.User.Update().
		Where(user2.IDEQ(req.ID)).
		SetAvatar(req.Avatar).
		SetMobile(req.Mobile).
		SetJobTime(req.JobTime).
		SetStatus(req.Status).
		SetUsername(req.Username).
		SetName(req.Name).
		SetFunctions(functions).
		SetGender(gender).
		SetRoleID(req.RoleId).
		SetDetail(req.Detail).

		//SetTags().
		Save(u.ctx)

	if err != nil {
		err = errors.Wrap(err, "update user failed")
		return err
	}
	u.db.User.UpdateOneID(req.ID).AddTagIDs(req.UserTags...).Exec(u.ctx)
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
	if req.Mobile != "" {
		predicates = append(predicates, user2.MobileEQ(req.Mobile))
	}

	if req.JobTime != 0 {
		predicates = append(predicates, user2.JobTime(req.JobTime))
	}
	if req.Mobile != "" {
		predicates = append(predicates, user2.Mobile(req.Mobile))
	}
	if req.Name != "" {
		predicates = append(predicates, user2.NameContains(req.Name))
	}
	if req.Status != 0 {
		predicates = append(predicates, user2.Status(req.Status))
	}

	if req.RoleId != 0 {
		predicates = append(predicates, user2.RoleIDEQ(req.RoleId))
	}
	predicates = append(predicates, user2.Delete(0))
	users, err := u.db.User.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(u.ctx)
	if err != nil {
		err = errors.Wrap(err, "get user list failed")
		return userList, total, err
	}
	// copy to UserInfo struct

	for _, v := range users {
		mr := u.entUserInfo(*v)
		userList = append(userList, mr)
	}
	total, _ = u.db.User.Query().Where(predicates...).Count(u.ctx)
	return
}

func (u User) entUserInfo(userEnt ent.User) (info *user.UserInfo) {
	info = new(user.UserInfo)

	info.Id = userEnt.ID
	info.Status = userEnt.Status
	info.Username = userEnt.Username

	info.Name = userEnt.Name
	info.RoleId = userEnt.RoleID
	info.Mobile = userEnt.Mobile
	info.CreatedAt = userEnt.CreatedAt.Format(time.DateTime)
	info.UpdatedAt = userEnt.UpdatedAt.Format(time.DateTime)

	roleInfo, _ := u.db.Role.Query().Where(role.IDEQ(info.RoleId)).First(u.ctx)
	if roleInfo != nil {
		info.RoleName = roleInfo.Name
		info.RoleValue = roleInfo.Value
	}

	if userEnt.Gender == 0 {
		info.Gender = "女性"
	} else if userEnt.Gender == 1 {
		info.Gender = "男性"
	} else {
		info.Gender = "保密"
	}

	info.Functions = strings.Split(userEnt.Functions, ",")

	var gender string
	if userEnt.Gender == 0 {
		gender = "女性"
	} else if userEnt.Gender == 1 {
		gender = "男性"
	} else {
		gender = "未知"
	}

	info.Gender = gender
	info.Detail = userEnt.Detail
	info.JobTime = &userEnt.JobTime
	info.DefaultVenueId = userEnt.DefaultVenueID

	Tags, _ := userEnt.QueryTag().All(u.ctx)
	if len(Tags) > 0 {
		for _, d := range Tags {
			dd := dictionary.DictionaryDetail{
				ID:     d.ID,
				Title:  d.Title,
				Key:    d.Key,
				Value:  d.Value,
				Status: d.Status,
			}
			info.UserTags = append(info.UserTags, &dd)
		}
	}

	return info
}

func (u User) UpdateUserStatus(id int64, status int64) error {
	_, err := u.db.User.Update().Where(user2.IDEQ(id)).SetStatus(status).Save(u.ctx)
	return err
}

func (u User) DeleteUser(id int64) error {
	_, err := u.db.User.Delete().Where(user2.IDEQ(id)).Exec(u.ctx)
	return err
}
