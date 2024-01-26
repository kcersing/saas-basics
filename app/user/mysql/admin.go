package mysql

import "saas/pkg/db/ent"

type AdminManager struct {
	salt string
	db   *ent.Client
}

func (a AdminManager) GetAdminByAccountId(aid string) error {
	//TODO implement me
	panic("implement me")
}
func (a AdminManager) GetAdminByName(name string) error {
	//TODO implement me
	panic("implement me")
}

func (a AdminManager) UpdateAdminPassword(aid string, password string) error {
	//TODO implement me
	panic("implement me")
}

func NewAdminManager(db *ent.Client, salt string) *AdminManager {
	return &AdminManager{
		db:   db,
		salt: salt,
	}
}
