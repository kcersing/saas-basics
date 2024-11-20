package do

type Member interface {
	Create(req CreateOrUpdateMemberReq) error
	Update(req CreateOrUpdateMemberReq) error
	Delete(id int64) error
	List(req MemberListReq) (resp []*MemberInfo, total int, err error)
	Info(id int64) (info *MemberInfo, err error)
	ChangePassword(ID int64, oldPassword, newPassword string) error
	UpdateStatus(ID int64, status int64) error
	Search(option string, value string) (memberInfo *MemberInfo, err error)
}

type CreateOrUpdateMemberPrivateReq struct {
	MemberId         int64  `json:"member_id"`
	IdentityCard     string `json:"identity_card"`
	FaceIdentityCard string `json:"face_identity_card"`
	BackIdentityCard string `json:"back_identity_card"`
	FacePic          string `json:"face_pic"`
	FaceEigenvalue   string `json:"face_eigenvalue"`
}
