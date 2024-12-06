package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"saas/pkg/db/ent/schema/mixins"
)

type Bootcamp struct {
	ent.Schema
}

func (Bootcamp) Fields() []ent.Field {
	return []ent.Field{

		field.String("name").Comment("训练营名称").Optional(),
		field.Int64("sign_number").Comment("报名人数").Optional(),
		field.Time("sign_start_at").Comment("报名开始时间").Optional(),
		field.Time("sign_end_at").Comment("报名结束时间").Optional(),
		field.Int64("number").Comment("训练营人数").Optional(),
		field.Time("start_at").Comment("训练营开始时间").Optional(),
		field.Time("end_at").Comment("训练营结束时间").Optional(),
		field.String("pic").Comment("训练营图片").Optional(),
		field.String("sponsor").Comment("主办方").Optional(),
		field.Float("fee").Comment("费用").Optional(),
		field.Int64("is_fee").Default(1).Comment("是否有费用 1 无 2 有").Optional(),

		field.Int64("is_cancel").Comment("是否支持取消报名 0支持 1不支持").Default(0).Optional(),
		field.Int64("cancel_time").Comment("取消时间").Default(0).Optional(),
		field.String("detail").Comment("详情").Optional(),
		field.Text("sign_fields").Comment("报名信息").Optional(),

		field.Int64("condition").
			Default(1).
			Optional().
			Comment("状态[1:未报名;2:报名中;3:未开始;4:进行中;5:结束]"),
	}
}

func (Bootcamp) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Bootcamp) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bootcamp_participants", BootcampParticipant.Type),
	}
}

func (Bootcamp) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "sign_start_at", "sign_end_at", "start_at", "end_at").
			Unique(),
	}
}

func (Bootcamp) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "bootcamp"},
		entsql.WithComments(true),
	}
}
