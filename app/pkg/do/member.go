package do

import "time"

type Member interface {
	Create(req MemberInfo) error
	Update(req MemberInfo) error
	Delete(id int64) error
	List(req MemberListReq) (resp []*MemberInfo, total int, err error)
	Info(ID int64) (memberInfo *MemberInfo, err error)
	ChangePassword(ID int64, oldPassword, newPassword string) error
	UpdateStatus(ID int64, status int8) error
}

type CreateOrUpdateMemberReq struct {
	ID       int64  `json:"id,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	Email    string `json:"email,omitempty"`
	Status   int64  `json:"status,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Nickname string `json:"nickname,omitempty"`
}
type MemberInfo struct {
	ID        int64     `json:"id,omitempty"`
	Status    int8      `json:"status,omitempty"`
	Username  string    `json:"username,omitempty"`
	Password  string    `json:"password,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Mobile    string    `json:"mobile,omitempty"`
	Email     string    `json:"email,omitempty"`
	Avatar    string    `json:"avatar,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type MemberListReq struct {
	Page     int64  `json:"page,omitempty"`
	PageSize int64  `json:"pageSize,omitempty"`
	Username string `json:"username,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
}
