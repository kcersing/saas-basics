// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/pkg/db/ent/logs"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Logs is the model entity for the Logs schema.
type Logs struct {
	config `json:"-"`
	// ID of the ent.
	// primary key
	ID int64 `json:"id,omitempty"`
	// created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// last update time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// type of log | 日志类型
	Type string `json:"type,omitempty"`
	// method of log | 日志请求方法
	Method string `json:"method,omitempty"`
	// api of log | 日志请求api
	API string `json:"api,omitempty"`
	// success of log | 日志请求是否成功
	Success bool `json:"success,omitempty"`
	// content of request log | 日志请求内容
	ReqContent string `json:"req_content,omitempty"`
	// content of response log | 日志返回内容
	RespContent string `json:"resp_content,omitempty"`
	// ip of log | 日志IP
	IP string `json:"ip,omitempty"`
	// user_agent of log | 日志用户客户端
	UserAgent string `json:"user_agent,omitempty"`
	// operator of log | 日志操作者
	Operator string `json:"operator,omitempty"`
	// time of log(millisecond) | 日志时间(毫秒)
	Time         int `json:"time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Logs) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case logs.FieldSuccess:
			values[i] = new(sql.NullBool)
		case logs.FieldID, logs.FieldTime:
			values[i] = new(sql.NullInt64)
		case logs.FieldType, logs.FieldMethod, logs.FieldAPI, logs.FieldReqContent, logs.FieldRespContent, logs.FieldIP, logs.FieldUserAgent, logs.FieldOperator:
			values[i] = new(sql.NullString)
		case logs.FieldCreatedAt, logs.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Logs fields.
func (l *Logs) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case logs.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int64(value.Int64)
		case logs.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				l.CreatedAt = value.Time
			}
		case logs.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				l.UpdatedAt = value.Time
			}
		case logs.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				l.Type = value.String
			}
		case logs.FieldMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				l.Method = value.String
			}
		case logs.FieldAPI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field api", values[i])
			} else if value.Valid {
				l.API = value.String
			}
		case logs.FieldSuccess:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field success", values[i])
			} else if value.Valid {
				l.Success = value.Bool
			}
		case logs.FieldReqContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field req_content", values[i])
			} else if value.Valid {
				l.ReqContent = value.String
			}
		case logs.FieldRespContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resp_content", values[i])
			} else if value.Valid {
				l.RespContent = value.String
			}
		case logs.FieldIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip", values[i])
			} else if value.Valid {
				l.IP = value.String
			}
		case logs.FieldUserAgent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_agent", values[i])
			} else if value.Valid {
				l.UserAgent = value.String
			}
		case logs.FieldOperator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operator", values[i])
			} else if value.Valid {
				l.Operator = value.String
			}
		case logs.FieldTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field time", values[i])
			} else if value.Valid {
				l.Time = int(value.Int64)
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Logs.
// This includes values selected through modifiers, order, etc.
func (l *Logs) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// Update returns a builder for updating this Logs.
// Note that you need to call Logs.Unwrap() before calling this method if this Logs
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Logs) Update() *LogsUpdateOne {
	return NewLogsClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Logs entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Logs) Unwrap() *Logs {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Logs is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Logs) String() string {
	var builder strings.Builder
	builder.WriteString("Logs(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("created_at=")
	builder.WriteString(l.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(l.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(l.Type)
	builder.WriteString(", ")
	builder.WriteString("method=")
	builder.WriteString(l.Method)
	builder.WriteString(", ")
	builder.WriteString("api=")
	builder.WriteString(l.API)
	builder.WriteString(", ")
	builder.WriteString("success=")
	builder.WriteString(fmt.Sprintf("%v", l.Success))
	builder.WriteString(", ")
	builder.WriteString("req_content=")
	builder.WriteString(l.ReqContent)
	builder.WriteString(", ")
	builder.WriteString("resp_content=")
	builder.WriteString(l.RespContent)
	builder.WriteString(", ")
	builder.WriteString("ip=")
	builder.WriteString(l.IP)
	builder.WriteString(", ")
	builder.WriteString("user_agent=")
	builder.WriteString(l.UserAgent)
	builder.WriteString(", ")
	builder.WriteString("operator=")
	builder.WriteString(l.Operator)
	builder.WriteString(", ")
	builder.WriteString("time=")
	builder.WriteString(fmt.Sprintf("%v", l.Time))
	builder.WriteByte(')')
	return builder.String()
}

// LogsSlice is a parsable slice of Logs.
type LogsSlice []*Logs
