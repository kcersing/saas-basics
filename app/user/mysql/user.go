package mysql

import (
	"context"
	"errors"
	"saas/kitex_gen/cwg/user"
	"saas/pkg/db/ent"
	user2 "saas/pkg/db/ent/user"
	"saas/pkg/errno"
	"saas/pkg/md5"
)

type UserManager struct {
	salt string
	db   *ent.Client
}

func NewUserManager(db *ent.Client, salt string) *UserManager {
	return &UserManager{
		db:   db,
		salt: salt,
	}
}

func (m *UserManager) CreateUser(req *user.AddUserRequest) (*ent.User, error) {
	if _, err := m.GetUserByMobile(req.Mobile); err == nil {
		return nil, errno.RecordAlreadyExist
	} else if !errors.Is(err, errno.RecordNotFound) {
		return nil, err
	}
	var cryPassword string
	if req.Password == "" {
		cryPassword = md5.Md5Crypt("s123456", m.salt)
	} else {
		cryPassword = md5.Md5Crypt(req.Password, m.salt)
	}
	u, err := m.db.User.Create().
		SetGender(req.Gender).
		SetMobile(req.Mobile).
		SetUsername(req.Username).
		SetPassword(cryPassword).
		SetAge(req.Age).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (m *UserManager) GetUserByMobile(mobile string) (*ent.User, error) {
	u, err := m.db.User.Query().Where(user2.Mobile(mobile)).First(context.Background())
	if err != nil {
		return nil, errno.RecordNotFound
	}
	return u, nil
}

func (m *UserManager) GetUserByAccountId(aid string) (*ent.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *UserManager) GetUserList() ([]*ent.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *UserManager) DeleteUser(aid string) error {
	//TODO implement me
	panic("implement me")
}
