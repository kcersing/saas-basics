// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/pkg/db/ent/token"
	"saas/pkg/db/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
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
	// mobile number | 手机号
	Mobile string `json:"mobile,omitempty"`
	// name | 姓名
	Name string `json:"name,omitempty"`
	// 性别 | [0:女性;1:男性;3:保密]
	Gender int64 `json:"gender,omitempty"`
	// user's login name | 登录名
	Username string `json:"username,omitempty"`
	// password | 密码
	Password string `json:"password,omitempty"`
	// functions | 职能
	Functions string `json:"functions,omitempty"`
	// job time | [1:全职;2:兼职;]
	JobTime int64 `json:"job_time,omitempty"`
	// role id | 角色ID
	RoleID int64 `json:"role_id,omitempty"`
	// 登陆后默认场馆ID
	DefaultVenueID int64 `json:"default_venue_id,omitempty"`
	// avatar | 头像路径
	Avatar string `json:"avatar,omitempty"`
	// 详情
	Detail string `json:"detail,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Token holds the value of the token edge.
	Token *Token `json:"token,omitempty"`
	// Tag holds the value of the tag edge.
	Tag []*DictionaryDetail `json:"tag,omitempty"`
	// CreatedOrders holds the value of the created_orders edge.
	CreatedOrders []*Order `json:"created_orders,omitempty"`
	// UserEntry holds the value of the user_entry edge.
	UserEntry []*EntryLogs `json:"user_entry,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// TokenOrErr returns the Token value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) TokenOrErr() (*Token, error) {
	if e.loadedTypes[0] {
		if e.Token == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: token.Label}
		}
		return e.Token, nil
	}
	return nil, &NotLoadedError{edge: "token"}
}

// TagOrErr returns the Tag value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TagOrErr() ([]*DictionaryDetail, error) {
	if e.loadedTypes[1] {
		return e.Tag, nil
	}
	return nil, &NotLoadedError{edge: "tag"}
}

// CreatedOrdersOrErr returns the CreatedOrders value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CreatedOrdersOrErr() ([]*Order, error) {
	if e.loadedTypes[2] {
		return e.CreatedOrders, nil
	}
	return nil, &NotLoadedError{edge: "created_orders"}
}

// UserEntryOrErr returns the UserEntry value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserEntryOrErr() ([]*EntryLogs, error) {
	if e.loadedTypes[3] {
		return e.UserEntry, nil
	}
	return nil, &NotLoadedError{edge: "user_entry"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldDelete, user.FieldCreatedID, user.FieldStatus, user.FieldGender, user.FieldJobTime, user.FieldRoleID, user.FieldDefaultVenueID:
			values[i] = new(sql.NullInt64)
		case user.FieldMobile, user.FieldName, user.FieldUsername, user.FieldPassword, user.FieldFunctions, user.FieldAvatar, user.FieldDetail:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int64(value.Int64)
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				u.Delete = value.Int64
			}
		case user.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				u.CreatedID = value.Int64
			}
		case user.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				u.Status = value.Int64
			}
		case user.FieldMobile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mobile", values[i])
			} else if value.Valid {
				u.Mobile = value.String
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldGender:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				u.Gender = value.Int64
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldFunctions:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field functions", values[i])
			} else if value.Valid {
				u.Functions = value.String
			}
		case user.FieldJobTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field job_time", values[i])
			} else if value.Valid {
				u.JobTime = value.Int64
			}
		case user.FieldRoleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value.Valid {
				u.RoleID = value.Int64
			}
		case user.FieldDefaultVenueID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field default_venue_id", values[i])
			} else if value.Valid {
				u.DefaultVenueID = value.Int64
			}
		case user.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				u.Avatar = value.String
			}
		case user.FieldDetail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field detail", values[i])
			} else if value.Valid {
				u.Detail = value.String
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryToken queries the "token" edge of the User entity.
func (u *User) QueryToken() *TokenQuery {
	return NewUserClient(u.config).QueryToken(u)
}

// QueryTag queries the "tag" edge of the User entity.
func (u *User) QueryTag() *DictionaryDetailQuery {
	return NewUserClient(u.config).QueryTag(u)
}

// QueryCreatedOrders queries the "created_orders" edge of the User entity.
func (u *User) QueryCreatedOrders() *OrderQuery {
	return NewUserClient(u.config).QueryCreatedOrders(u)
}

// QueryUserEntry queries the "user_entry" edge of the User entity.
func (u *User) QueryUserEntry() *EntryLogsQuery {
	return NewUserClient(u.config).QueryUserEntry(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", u.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", u.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", u.Status))
	builder.WriteString(", ")
	builder.WriteString("mobile=")
	builder.WriteString(u.Mobile)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(fmt.Sprintf("%v", u.Gender))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteString(", ")
	builder.WriteString("functions=")
	builder.WriteString(u.Functions)
	builder.WriteString(", ")
	builder.WriteString("job_time=")
	builder.WriteString(fmt.Sprintf("%v", u.JobTime))
	builder.WriteString(", ")
	builder.WriteString("role_id=")
	builder.WriteString(fmt.Sprintf("%v", u.RoleID))
	builder.WriteString(", ")
	builder.WriteString("default_venue_id=")
	builder.WriteString(fmt.Sprintf("%v", u.DefaultVenueID))
	builder.WriteString(", ")
	builder.WriteString("avatar=")
	builder.WriteString(u.Avatar)
	builder.WriteString(", ")
	builder.WriteString("detail=")
	builder.WriteString(u.Detail)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
