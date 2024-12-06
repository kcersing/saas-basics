package do

import "saas/idl_gen/model/banner"

type Banner interface {
	Create(req banner.BannerInfo) error
	Update(req banner.BannerInfo) error
	Delete(id int64) error
	List() (resp []*banner.BannerInfo, total int64, err error)
	UpdateShow(ID int64, status int64) error
}
