// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"saas/biz/dal/db/ent/product"
	"saas/idl_gen/model/base"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Product is the model entity for the Product schema.
type Product struct {
	config `json:"-"`
	// ID of the ent.
	// primary key
	ID int64 `json:"id,omitempty"`
	// created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// last update time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// last delete  1:已删除
	Delete int64 `json:"delete,omitempty"`
	// created
	CreatedID int64 `json:"created_id,omitempty"`
	// 状态[1:正常,2:禁用]
	Status int64 `json:"status,omitempty"`
	// 类型
	Type string `json:"type,omitempty"`
	// 商品名
	Name string `json:"name,omitempty"`
	// 库存
	Stock int64 `json:"stock,omitempty"`
	// 激活期限
	Deadline int64 `json:"deadline,omitempty"`
	// 有效期
	Duration int64 `json:"duration,omitempty"`
	// 课程课时
	Length int64 `json:"length,omitempty"`
	// 售价
	Price float64 `json:"price,omitempty"`
	// 次数
	Times int64 `json:"times,omitempty"`
	// 团课预约 1支持2不支持
	IsLessons int64 `json:"is_lessons,omitempty"`
	// 售卖信息[售价等]
	Sales []*base.Sales `json:"sales,omitempty"`
	// 销售方式 1会员端
	IsSales int64 `json:"is_sales,omitempty"`
	// 开始售卖时间
	SignSalesAt time.Time `json:"sign_sales_at,omitempty"`
	// 结束售卖时间
	EndSalesAt time.Time `json:"end_sales_at,omitempty"`
	// 主图
	Pic string `json:"pic,omitempty"`
	// 详情
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductQuery when eager-loading is set.
	Edges        ProductEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProductEdges holds the relations/edges for other nodes in the graph.
type ProductEdges struct {
	// Tag holds the value of the tag edge.
	Tag []*DictionaryDetail `json:"tag,omitempty"`
	// Contracts holds the value of the contracts edge.
	Contracts []*Contract `json:"contracts,omitempty"`
	// Products holds the value of the products edge.
	Products []*Product `json:"products,omitempty"`
	// Lessons holds the value of the lessons edge.
	Lessons []*Product `json:"lessons,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// TagOrErr returns the Tag value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) TagOrErr() ([]*DictionaryDetail, error) {
	if e.loadedTypes[0] {
		return e.Tag, nil
	}
	return nil, &NotLoadedError{edge: "tag"}
}

// ContractsOrErr returns the Contracts value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) ContractsOrErr() ([]*Contract, error) {
	if e.loadedTypes[1] {
		return e.Contracts, nil
	}
	return nil, &NotLoadedError{edge: "contracts"}
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) ProductsOrErr() ([]*Product, error) {
	if e.loadedTypes[2] {
		return e.Products, nil
	}
	return nil, &NotLoadedError{edge: "products"}
}

// LessonsOrErr returns the Lessons value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) LessonsOrErr() ([]*Product, error) {
	if e.loadedTypes[3] {
		return e.Lessons, nil
	}
	return nil, &NotLoadedError{edge: "lessons"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Product) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case product.FieldSales:
			values[i] = new([]byte)
		case product.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case product.FieldID, product.FieldDelete, product.FieldCreatedID, product.FieldStatus, product.FieldStock, product.FieldDeadline, product.FieldDuration, product.FieldLength, product.FieldTimes, product.FieldIsLessons, product.FieldIsSales:
			values[i] = new(sql.NullInt64)
		case product.FieldType, product.FieldName, product.FieldPic, product.FieldDescription:
			values[i] = new(sql.NullString)
		case product.FieldCreatedAt, product.FieldUpdatedAt, product.FieldSignSalesAt, product.FieldEndSalesAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Product fields.
func (pr *Product) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case product.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int64(value.Int64)
		case product.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case product.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case product.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				pr.Delete = value.Int64
			}
		case product.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				pr.CreatedID = value.Int64
			}
		case product.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pr.Status = value.Int64
			}
		case product.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				pr.Type = value.String
			}
		case product.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case product.FieldStock:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field stock", values[i])
			} else if value.Valid {
				pr.Stock = value.Int64
			}
		case product.FieldDeadline:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deadline", values[i])
			} else if value.Valid {
				pr.Deadline = value.Int64
			}
		case product.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				pr.Duration = value.Int64
			}
		case product.FieldLength:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field length", values[i])
			} else if value.Valid {
				pr.Length = value.Int64
			}
		case product.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				pr.Price = value.Float64
			}
		case product.FieldTimes:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field times", values[i])
			} else if value.Valid {
				pr.Times = value.Int64
			}
		case product.FieldIsLessons:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_lessons", values[i])
			} else if value.Valid {
				pr.IsLessons = value.Int64
			}
		case product.FieldSales:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field sales", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.Sales); err != nil {
					return fmt.Errorf("unmarshal field sales: %w", err)
				}
			}
		case product.FieldIsSales:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_sales", values[i])
			} else if value.Valid {
				pr.IsSales = value.Int64
			}
		case product.FieldSignSalesAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field sign_sales_at", values[i])
			} else if value.Valid {
				pr.SignSalesAt = value.Time
			}
		case product.FieldEndSalesAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_sales_at", values[i])
			} else if value.Valid {
				pr.EndSalesAt = value.Time
			}
		case product.FieldPic:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pic", values[i])
			} else if value.Valid {
				pr.Pic = value.String
			}
		case product.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pr.Description = value.String
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Product.
// This includes values selected through modifiers, order, etc.
func (pr *Product) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryTag queries the "tag" edge of the Product entity.
func (pr *Product) QueryTag() *DictionaryDetailQuery {
	return NewProductClient(pr.config).QueryTag(pr)
}

// QueryContracts queries the "contracts" edge of the Product entity.
func (pr *Product) QueryContracts() *ContractQuery {
	return NewProductClient(pr.config).QueryContracts(pr)
}

// QueryProducts queries the "products" edge of the Product entity.
func (pr *Product) QueryProducts() *ProductQuery {
	return NewProductClient(pr.config).QueryProducts(pr)
}

// QueryLessons queries the "lessons" edge of the Product entity.
func (pr *Product) QueryLessons() *ProductQuery {
	return NewProductClient(pr.config).QueryLessons(pr)
}

// Update returns a builder for updating this Product.
// Note that you need to call Product.Unwrap() before calling this method if this Product
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Product) Update() *ProductUpdateOne {
	return NewProductClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Product entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Product) Unwrap() *Product {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Product is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Product) String() string {
	var builder strings.Builder
	builder.WriteString("Product(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", pr.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pr.Status))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(pr.Type)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("stock=")
	builder.WriteString(fmt.Sprintf("%v", pr.Stock))
	builder.WriteString(", ")
	builder.WriteString("deadline=")
	builder.WriteString(fmt.Sprintf("%v", pr.Deadline))
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", pr.Duration))
	builder.WriteString(", ")
	builder.WriteString("length=")
	builder.WriteString(fmt.Sprintf("%v", pr.Length))
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", pr.Price))
	builder.WriteString(", ")
	builder.WriteString("times=")
	builder.WriteString(fmt.Sprintf("%v", pr.Times))
	builder.WriteString(", ")
	builder.WriteString("is_lessons=")
	builder.WriteString(fmt.Sprintf("%v", pr.IsLessons))
	builder.WriteString(", ")
	builder.WriteString("sales=")
	builder.WriteString(fmt.Sprintf("%v", pr.Sales))
	builder.WriteString(", ")
	builder.WriteString("is_sales=")
	builder.WriteString(fmt.Sprintf("%v", pr.IsSales))
	builder.WriteString(", ")
	builder.WriteString("sign_sales_at=")
	builder.WriteString(pr.SignSalesAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_sales_at=")
	builder.WriteString(pr.EndSalesAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("pic=")
	builder.WriteString(pr.Pic)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pr.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Products is a parsable slice of Product.
type Products []*Product
