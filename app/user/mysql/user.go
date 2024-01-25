package mysql

import "saas/pkg/db/ent"

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
