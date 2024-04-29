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
	ID       int64  `json:"id,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	Email    string `json:"email,omitempty"`
	Status   int64  `json:"status,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Name     string `json:"Name,omitempty"`
	Wecom    string `json:"wecom,omitempty"`
	Gender   int64  `json:"gender,omitempty"`
	Age      string `json:"age,omitempty"`
	Birthday string `json:"birthday,omitempty"`
}

type CreateOrUpdateMemberPrivateReq struct {
	MemberId         int64  `json:"member_id,omitempty"`
	IdentityCard     string `json:"identity_card,omitempty"`
	FaceIdentityCard string `json:"face_identity_card,omitempty"`
	BackIdentityCard string `json:"back_identity_card,omitempty"`
	FacePic          string `json:"face_pic,omitempty"`
	FaceEigenvalue   string `json:"face_eigenvalue,omitempty"`
}

type MemberInfo struct {
	ID        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Condition int64  `json:"condition,omitempty"`
	Status    int64  `json:"status,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	Mobile    string `json:"mobile,omitempty"`
	Email     string `json:"email,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Gender    string `json:"gender,omitempty"`
	Age       string `json:"age,omitempty"`

	//微信号
	Wecom string `json:"wecom,omitempty"`
	//身份证号
	IdentityCard string `json:"identity_card,omitempty"`
	//正面
	FaceIdentityCard string `json:"face_identity_card,omitempty"`
	//反面
	BackIdentityCard string `json:"back_identity_card,omitempty"`
	//人脸照片
	FacePic string `json:"face_pic,omitempty"`
	//特征值
	FaceEigenvalue string `json:"face_eigenvalue,omitempty"`
	//人脸更新时间
	FacePicUpdatedTime string `json:"face_pic_updated_time,omitempty"`
	//消费总金额
	MoneySum float64 `json:"money_sum,omitempty"`
	//首次的产品
	ProductId int64 `json:"product_id,omitempty"`
	//首次消费场馆
	ProductVenue int64 `json:"product_venue,omitempty"`
	//进馆总次数
	EntrySum int64 `json:"entry_sum,omitempty"`
	//最后一次进馆时间
	EntryLastTime int64 `json:"entry_last_time,omitempty"`
	//进馆最后期限时间
	EntryDeadlineTime int64 `json:"entry_deadline_time,omitempty"`
	//最后一次上课时间
	ClassLastTime int64 `json:"class_last_time,omitempty"`
	//关联员工
	RelationUid int64 `json:"relation_uid,omitempty"`
	//关联会员
	RelationMid int64 `json:"relation_mid,omitempty"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type MemberListReq struct {
	Page     int64  `json:"page,omitempty"`
	PageSize int64  `json:"pageSize,omitempty"`
	Name     string `json:"name,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
}
