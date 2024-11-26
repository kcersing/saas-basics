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

	ProductSearch(members []int64) (info *ProductInfo, err error)
	PropertySearch(memberProducts []int64) (info *PropertyInfo, err error)
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
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	Password string `json:"password"`
	CreateId int64  `json:"createId"`
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
	Nickname  string `json:"nickname"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Gender    string `json:"gender"`
	Age       int64  `json:"age"`
	Birthday  string `json:"birthday"`
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
	ProductId   int64 `json:"product_id"`
	ProductName int64 `json:"product_name"`
	//首次消费场馆
	ProductVenue     int64 `json:"product_venue"`
	ProductVenueName int64 `json:"product_venue_name"`
	//进馆总次数
	EntrySum int64 `json:"entry_sum"`
	//最后一次进馆时间
	EntryLastTime string `json:"entry_last_time"`
	//进馆最后期限时间
	EntryDeadlineTime string `json:"entry_deadline_time"`
	//最后一次上课时间
	ClassLastTime string `json:"class_last_time"`
	//关联员工
	RelationUid   int64 `json:"relation_uid"`
	RelationUname int64 `json:"relation_uname"`
	//关联会员
	RelationMid   int64 `json:"relation_mid"`
	RelationMname int64 `json:"relation_mname"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type MemberListReq struct {
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
}
