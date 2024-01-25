package mysql

import "saas/pkg/db/ent"

type AdminManager struct {
	salt string
	db   *ent.Client
}

func NewAdminManager(db *ent.Client, salt string) *AdminManager {
	return &AdminManager{
		db:   db,
		salt: salt,
	}
}
