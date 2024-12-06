// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/pkg/db/ent/banner"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Banner is the model entity for the Banner schema.
type Banner struct {
	config `json:"-"`
	// ID of the ent.
	// primary key
	ID int64 `json:"id,omitempty"`
	// created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// last update time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// last delete
	Delete int64 `json:"delete,omitempty"`
	// created
	CreatedID int64 `json:"created_id,omitempty"`
	// 状态[1:正常,2:禁用]
	Status int64 `json:"status,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 图片
	Pic string `json:"pic,omitempty"`
	// 跳转链接
	Link string `json:"link,omitempty"`
	// 是否展示 1 展示 2 不展示
	IsShow       int64 `json:"is_show,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Banner) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case banner.FieldID, banner.FieldDelete, banner.FieldCreatedID, banner.FieldStatus, banner.FieldIsShow:
			values[i] = new(sql.NullInt64)
		case banner.FieldName, banner.FieldPic, banner.FieldLink:
			values[i] = new(sql.NullString)
		case banner.FieldCreatedAt, banner.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Banner fields.
func (b *Banner) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case banner.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int64(value.Int64)
		case banner.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case banner.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case banner.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				b.Delete = value.Int64
			}
		case banner.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				b.CreatedID = value.Int64
			}
		case banner.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				b.Status = value.Int64
			}
		case banner.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				b.Name = value.String
			}
		case banner.FieldPic:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pic", values[i])
			} else if value.Valid {
				b.Pic = value.String
			}
		case banner.FieldLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field link", values[i])
			} else if value.Valid {
				b.Link = value.String
			}
		case banner.FieldIsShow:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_show", values[i])
			} else if value.Valid {
				b.IsShow = value.Int64
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Banner.
// This includes values selected through modifiers, order, etc.
func (b *Banner) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// Update returns a builder for updating this Banner.
// Note that you need to call Banner.Unwrap() before calling this method if this Banner
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Banner) Update() *BannerUpdateOne {
	return NewBannerClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Banner entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Banner) Unwrap() *Banner {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Banner is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Banner) String() string {
	var builder strings.Builder
	builder.WriteString("Banner(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", b.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", b.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", b.Status))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(b.Name)
	builder.WriteString(", ")
	builder.WriteString("pic=")
	builder.WriteString(b.Pic)
	builder.WriteString(", ")
	builder.WriteString("link=")
	builder.WriteString(b.Link)
	builder.WriteString(", ")
	builder.WriteString("is_show=")
	builder.WriteString(fmt.Sprintf("%v", b.IsShow))
	builder.WriteByte(')')
	return builder.String()
}

// Banners is a parsable slice of Banner.
type Banners []*Banner
