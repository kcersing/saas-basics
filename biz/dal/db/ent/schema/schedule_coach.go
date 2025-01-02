/*
 * @Author: kcersing wt4@live.cn
 * @Date: 2024-06-01 23:43:32
 * @LastEditors: kcersing wt4@live.cn
 * @LastEditTime: 2024-06-02 01:49:02
 * @FilePath: \saas\biz\dal\db\ent\schema\schedule_coach.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package schema

import (
	"entgo.io/ent/schema/index"
	"saas/biz/dal/db/ent/schema/mixins"

	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ScheduleCoach struct {
	ent.Schema
}

func (ScheduleCoach) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("venue_id").Comment("场馆id").Optional(),
		field.Int64("coach_id").Comment("教练ID").Optional(),
		field.Int64("place_id").Comment("场地ID").Optional(),
		field.Int64("schedule_id").Comment("课程ID").Optional(),
		field.Int64("product_id").Comment("课程").Optional(),
		field.String("schedule_name").Comment("课程名称").Optional(),
		field.String("type").Comment("类型").Optional(),
		field.Time("start_time").Default(time.Now).Comment("开始时间").Optional(),
		field.Time("end_time").Default(time.Now).Comment("结束时间").Optional(),
		field.Time("sign_start_time").Default(time.Now).Comment("上课签到时间").Optional(),
		field.Time("sign_end_time").Default(time.Now).Comment("下课签到时间").Optional(),

		field.String("coach_name").Comment("教练名称").Optional(),
	}
}

func (ScheduleCoach) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (ScheduleCoach) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("schedule", Schedule.Type).Ref("coachs").Unique().Field("schedule_id"),
	}
}

func (ScheduleCoach) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("venue_id"),
		index.Fields("coach_id"),
		index.Fields("schedule_id"),
	}
}

func (ScheduleCoach) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "schedule_coach", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
