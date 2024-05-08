package do

import "time"

type Member interface {
	Create(req MemberInfo) error
	Update(req MemberInfo) error
	Delete(id int64) error
	List(req MemberListReq) (resp []*MemberInfo, total int, err error)
	Info(ID int64) (memberInfo *MemberInfo, err error)
	ChangePassword(ID int64, oldPassword, newPassword string) error
	UpdateStatus(ID int64, status int64) error
}

type CreateOrUpdateMemberReq struct {
	ID       int64  `json:"id"`
	Avatar   string `json:"avatar"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Status   int64  `json:"status"`
	Nickname string `json:"nickname"`
	Name     string `json:"Name"`
	Wecom    string `json:"wecom"`
	Gender   int64  `json:"gender"`
	Age      string `json:"age"`
	Birthday string `json:"birthday"`
}

type CreateOrUpdateMemberPrivateReq struct {
	MemberId         int64  `json:"member_id"`
	IdentityCard     string `json:"identity_card"`
	FaceIdentityCard string `json:"face_identity_card"`
	BackIdentityCard string `json:"back_identity_card"`
	FacePic          string `json:"face_pic"`
	FaceEigenvalue   string `json:"face_eigenvalue"`
}

type MemberInfo struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Condition int64  `json:"condition"`
	Status    int64  `json:"status"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Nickname  string `json:"nickname"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Gender    string `json:"gender"`
	Age       string `json:"age"`

	//微信号
	Wecom string `json:"wecom"`
	//身份证号
	IdentityCard string `json:"identity_card"`
	//正面
	FaceIdentityCard string `json:"face_identity_card"`
	//反面
	BackIdentityCard string `json:"back_identity_card"`
	//人脸照片
	FacePic string `json:"face_pic"`
	//特征值
	FaceEigenvalue string `json:"face_eigenvalue"`
	//人脸更新时间
	FacePicUpdatedTime string `json:"face_pic_updated_time"`
	//消费总金额
	MoneySum float64 `json:"money_sum"`
	//首次的产品
	ProductId int64 `json:"product_id"`
	//首次消费场馆
	ProductVenue int64 `json:"product_venue"`
	//进馆总次数
	EntrySum int64 `json:"entry_sum"`
	//最后一次进馆时间
	EntryLastTime int64 `json:"entry_last_time"`
	//进馆最后期限时间
	EntryDeadlineTime int64 `json:"entry_deadline_time"`
	//最后一次上课时间
	ClassLastTime int64 `json:"class_last_time"`
	//关联员工
	RelationUid int64 `json:"relation_uid"`
	//关联会员
	RelationMid int64 `json:"relation_mid"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type MemberListReq struct {
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
}
