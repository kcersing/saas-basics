package do

import (
	"saas/idl_gen/model/product"
)

type Product interface {
	CreateProduct(req product.CreateOrUpdateProductReq) error
	UpdateProduct(req product.CreateOrUpdateProductReq) error
	DeleteProduct(id int64) error
	UpdateProductStatus(ID int64, status int64) error
	ProductList(req *product.ProductListReq) (resp []*product.ProductInfo, total int, err error)
	ProductInfo(id int64) (resp *product.ProductInfo, err error)
	ProductListExport(req *product.ProductListReq) (resp string, err error)
}
