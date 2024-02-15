// Code generated by ent, DO NOT EDIT.

package menu

import (
	"saas/pkg/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldID, id))
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v int) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldParentID, v))
}

// RouteName applies equality check predicate on the "route_name" field. It's identical to RouteNameEQ.
func RouteName(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldRouteName, v))
}

// RoutePath applies equality check predicate on the "route_path" field. It's identical to RoutePathEQ.
func RoutePath(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldRoutePath, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldStatus, v))
}

// MenuName applies equality check predicate on the "menu_name" field. It's identical to MenuNameEQ.
func MenuName(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldMenuName, v))
}

// MenuType applies equality check predicate on the "menu_type" field. It's identical to MenuTypeEQ.
func MenuType(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldMenuType, v))
}

// IconType applies equality check predicate on the "icon_type" field. It's identical to IconTypeEQ.
func IconType(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldIconType, v))
}

// Icon applies equality check predicate on the "icon" field. It's identical to IconEQ.
func Icon(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldIcon, v))
}

// I18nKey applies equality check predicate on the "i18n_key" field. It's identical to I18nKeyEQ.
func I18nKey(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldI18nKey, v))
}

// Level applies equality check predicate on the "level" field. It's identical to LevelEQ.
func Level(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldLevel, v))
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v int) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldParentID, v))
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v int) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldParentID, v))
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...int) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldParentID, vs...))
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...int) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldParentID, vs...))
}

// ParentIDIsNil applies the IsNil predicate on the "parent_id" field.
func ParentIDIsNil() predicate.Menu {
	return predicate.Menu(sql.FieldIsNull(FieldParentID))
}

// ParentIDNotNil applies the NotNil predicate on the "parent_id" field.
func ParentIDNotNil() predicate.Menu {
	return predicate.Menu(sql.FieldNotNull(FieldParentID))
}

// RouteNameEQ applies the EQ predicate on the "route_name" field.
func RouteNameEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldRouteName, v))
}

// RouteNameNEQ applies the NEQ predicate on the "route_name" field.
func RouteNameNEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldRouteName, v))
}

// RouteNameIn applies the In predicate on the "route_name" field.
func RouteNameIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldRouteName, vs...))
}

// RouteNameNotIn applies the NotIn predicate on the "route_name" field.
func RouteNameNotIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldRouteName, vs...))
}

// RouteNameGT applies the GT predicate on the "route_name" field.
func RouteNameGT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldRouteName, v))
}

// RouteNameGTE applies the GTE predicate on the "route_name" field.
func RouteNameGTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldRouteName, v))
}

// RouteNameLT applies the LT predicate on the "route_name" field.
func RouteNameLT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldRouteName, v))
}

// RouteNameLTE applies the LTE predicate on the "route_name" field.
func RouteNameLTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldRouteName, v))
}

// RouteNameContains applies the Contains predicate on the "route_name" field.
func RouteNameContains(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContains(FieldRouteName, v))
}

// RouteNameHasPrefix applies the HasPrefix predicate on the "route_name" field.
func RouteNameHasPrefix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasPrefix(FieldRouteName, v))
}

// RouteNameHasSuffix applies the HasSuffix predicate on the "route_name" field.
func RouteNameHasSuffix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasSuffix(FieldRouteName, v))
}

// RouteNameIsNil applies the IsNil predicate on the "route_name" field.
func RouteNameIsNil() predicate.Menu {
	return predicate.Menu(sql.FieldIsNull(FieldRouteName))
}

// RouteNameNotNil applies the NotNil predicate on the "route_name" field.
func RouteNameNotNil() predicate.Menu {
	return predicate.Menu(sql.FieldNotNull(FieldRouteName))
}

// RouteNameEqualFold applies the EqualFold predicate on the "route_name" field.
func RouteNameEqualFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEqualFold(FieldRouteName, v))
}

// RouteNameContainsFold applies the ContainsFold predicate on the "route_name" field.
func RouteNameContainsFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContainsFold(FieldRouteName, v))
}

// RoutePathEQ applies the EQ predicate on the "route_path" field.
func RoutePathEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldRoutePath, v))
}

// RoutePathNEQ applies the NEQ predicate on the "route_path" field.
func RoutePathNEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldRoutePath, v))
}

// RoutePathIn applies the In predicate on the "route_path" field.
func RoutePathIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldRoutePath, vs...))
}

// RoutePathNotIn applies the NotIn predicate on the "route_path" field.
func RoutePathNotIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldRoutePath, vs...))
}

// RoutePathGT applies the GT predicate on the "route_path" field.
func RoutePathGT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldRoutePath, v))
}

// RoutePathGTE applies the GTE predicate on the "route_path" field.
func RoutePathGTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldRoutePath, v))
}

// RoutePathLT applies the LT predicate on the "route_path" field.
func RoutePathLT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldRoutePath, v))
}

// RoutePathLTE applies the LTE predicate on the "route_path" field.
func RoutePathLTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldRoutePath, v))
}

// RoutePathContains applies the Contains predicate on the "route_path" field.
func RoutePathContains(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContains(FieldRoutePath, v))
}

// RoutePathHasPrefix applies the HasPrefix predicate on the "route_path" field.
func RoutePathHasPrefix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasPrefix(FieldRoutePath, v))
}

// RoutePathHasSuffix applies the HasSuffix predicate on the "route_path" field.
func RoutePathHasSuffix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasSuffix(FieldRoutePath, v))
}

// RoutePathIsNil applies the IsNil predicate on the "route_path" field.
func RoutePathIsNil() predicate.Menu {
	return predicate.Menu(sql.FieldIsNull(FieldRoutePath))
}

// RoutePathNotNil applies the NotNil predicate on the "route_path" field.
func RoutePathNotNil() predicate.Menu {
	return predicate.Menu(sql.FieldNotNull(FieldRoutePath))
}

// RoutePathEqualFold applies the EqualFold predicate on the "route_path" field.
func RoutePathEqualFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEqualFold(FieldRoutePath, v))
}

// RoutePathContainsFold applies the ContainsFold predicate on the "route_path" field.
func RoutePathContainsFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContainsFold(FieldRoutePath, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Menu {
	return predicate.Menu(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Menu {
	return predicate.Menu(sql.FieldNotNull(FieldStatus))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContainsFold(FieldStatus, v))
}

// MenuNameEQ applies the EQ predicate on the "menu_name" field.
func MenuNameEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldMenuName, v))
}

// MenuNameNEQ applies the NEQ predicate on the "menu_name" field.
func MenuNameNEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldMenuName, v))
}

// MenuNameIn applies the In predicate on the "menu_name" field.
func MenuNameIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldMenuName, vs...))
}

// MenuNameNotIn applies the NotIn predicate on the "menu_name" field.
func MenuNameNotIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldMenuName, vs...))
}

// MenuNameGT applies the GT predicate on the "menu_name" field.
func MenuNameGT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldMenuName, v))
}

// MenuNameGTE applies the GTE predicate on the "menu_name" field.
func MenuNameGTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldMenuName, v))
}

// MenuNameLT applies the LT predicate on the "menu_name" field.
func MenuNameLT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldMenuName, v))
}

// MenuNameLTE applies the LTE predicate on the "menu_name" field.
func MenuNameLTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldMenuName, v))
}

// MenuNameContains applies the Contains predicate on the "menu_name" field.
func MenuNameContains(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContains(FieldMenuName, v))
}

// MenuNameHasPrefix applies the HasPrefix predicate on the "menu_name" field.
func MenuNameHasPrefix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasPrefix(FieldMenuName, v))
}

// MenuNameHasSuffix applies the HasSuffix predicate on the "menu_name" field.
func MenuNameHasSuffix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasSuffix(FieldMenuName, v))
}

// MenuNameIsNil applies the IsNil predicate on the "menu_name" field.
func MenuNameIsNil() predicate.Menu {
	return predicate.Menu(sql.FieldIsNull(FieldMenuName))
}

// MenuNameNotNil applies the NotNil predicate on the "menu_name" field.
func MenuNameNotNil() predicate.Menu {
	return predicate.Menu(sql.FieldNotNull(FieldMenuName))
}

// MenuNameEqualFold applies the EqualFold predicate on the "menu_name" field.
func MenuNameEqualFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEqualFold(FieldMenuName, v))
}

// MenuNameContainsFold applies the ContainsFold predicate on the "menu_name" field.
func MenuNameContainsFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContainsFold(FieldMenuName, v))
}

// MenuTypeEQ applies the EQ predicate on the "menu_type" field.
func MenuTypeEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldMenuType, v))
}

// MenuTypeNEQ applies the NEQ predicate on the "menu_type" field.
func MenuTypeNEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldMenuType, v))
}

// MenuTypeIn applies the In predicate on the "menu_type" field.
func MenuTypeIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldMenuType, vs...))
}

// MenuTypeNotIn applies the NotIn predicate on the "menu_type" field.
func MenuTypeNotIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldMenuType, vs...))
}

// MenuTypeGT applies the GT predicate on the "menu_type" field.
func MenuTypeGT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldMenuType, v))
}

// MenuTypeGTE applies the GTE predicate on the "menu_type" field.
func MenuTypeGTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldMenuType, v))
}

// MenuTypeLT applies the LT predicate on the "menu_type" field.
func MenuTypeLT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldMenuType, v))
}

// MenuTypeLTE applies the LTE predicate on the "menu_type" field.
func MenuTypeLTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldMenuType, v))
}

// MenuTypeContains applies the Contains predicate on the "menu_type" field.
func MenuTypeContains(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContains(FieldMenuType, v))
}

// MenuTypeHasPrefix applies the HasPrefix predicate on the "menu_type" field.
func MenuTypeHasPrefix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasPrefix(FieldMenuType, v))
}

// MenuTypeHasSuffix applies the HasSuffix predicate on the "menu_type" field.
func MenuTypeHasSuffix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasSuffix(FieldMenuType, v))
}

// MenuTypeIsNil applies the IsNil predicate on the "menu_type" field.
func MenuTypeIsNil() predicate.Menu {
	return predicate.Menu(sql.FieldIsNull(FieldMenuType))
}

// MenuTypeNotNil applies the NotNil predicate on the "menu_type" field.
func MenuTypeNotNil() predicate.Menu {
	return predicate.Menu(sql.FieldNotNull(FieldMenuType))
}

// MenuTypeEqualFold applies the EqualFold predicate on the "menu_type" field.
func MenuTypeEqualFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEqualFold(FieldMenuType, v))
}

// MenuTypeContainsFold applies the ContainsFold predicate on the "menu_type" field.
func MenuTypeContainsFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContainsFold(FieldMenuType, v))
}

// IconTypeEQ applies the EQ predicate on the "icon_type" field.
func IconTypeEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldIconType, v))
}

// IconTypeNEQ applies the NEQ predicate on the "icon_type" field.
func IconTypeNEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldIconType, v))
}

// IconTypeIn applies the In predicate on the "icon_type" field.
func IconTypeIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldIconType, vs...))
}

// IconTypeNotIn applies the NotIn predicate on the "icon_type" field.
func IconTypeNotIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldIconType, vs...))
}

// IconTypeGT applies the GT predicate on the "icon_type" field.
func IconTypeGT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldIconType, v))
}

// IconTypeGTE applies the GTE predicate on the "icon_type" field.
func IconTypeGTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldIconType, v))
}

// IconTypeLT applies the LT predicate on the "icon_type" field.
func IconTypeLT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldIconType, v))
}

// IconTypeLTE applies the LTE predicate on the "icon_type" field.
func IconTypeLTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldIconType, v))
}

// IconTypeContains applies the Contains predicate on the "icon_type" field.
func IconTypeContains(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContains(FieldIconType, v))
}

// IconTypeHasPrefix applies the HasPrefix predicate on the "icon_type" field.
func IconTypeHasPrefix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasPrefix(FieldIconType, v))
}

// IconTypeHasSuffix applies the HasSuffix predicate on the "icon_type" field.
func IconTypeHasSuffix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasSuffix(FieldIconType, v))
}

// IconTypeIsNil applies the IsNil predicate on the "icon_type" field.
func IconTypeIsNil() predicate.Menu {
	return predicate.Menu(sql.FieldIsNull(FieldIconType))
}

// IconTypeNotNil applies the NotNil predicate on the "icon_type" field.
func IconTypeNotNil() predicate.Menu {
	return predicate.Menu(sql.FieldNotNull(FieldIconType))
}

// IconTypeEqualFold applies the EqualFold predicate on the "icon_type" field.
func IconTypeEqualFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEqualFold(FieldIconType, v))
}

// IconTypeContainsFold applies the ContainsFold predicate on the "icon_type" field.
func IconTypeContainsFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContainsFold(FieldIconType, v))
}

// IconEQ applies the EQ predicate on the "icon" field.
func IconEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldIcon, v))
}

// IconNEQ applies the NEQ predicate on the "icon" field.
func IconNEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldIcon, v))
}

// IconIn applies the In predicate on the "icon" field.
func IconIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldIcon, vs...))
}

// IconNotIn applies the NotIn predicate on the "icon" field.
func IconNotIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldIcon, vs...))
}

// IconGT applies the GT predicate on the "icon" field.
func IconGT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldIcon, v))
}

// IconGTE applies the GTE predicate on the "icon" field.
func IconGTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldIcon, v))
}

// IconLT applies the LT predicate on the "icon" field.
func IconLT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldIcon, v))
}

// IconLTE applies the LTE predicate on the "icon" field.
func IconLTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldIcon, v))
}

// IconContains applies the Contains predicate on the "icon" field.
func IconContains(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContains(FieldIcon, v))
}

// IconHasPrefix applies the HasPrefix predicate on the "icon" field.
func IconHasPrefix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasPrefix(FieldIcon, v))
}

// IconHasSuffix applies the HasSuffix predicate on the "icon" field.
func IconHasSuffix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasSuffix(FieldIcon, v))
}

// IconIsNil applies the IsNil predicate on the "icon" field.
func IconIsNil() predicate.Menu {
	return predicate.Menu(sql.FieldIsNull(FieldIcon))
}

// IconNotNil applies the NotNil predicate on the "icon" field.
func IconNotNil() predicate.Menu {
	return predicate.Menu(sql.FieldNotNull(FieldIcon))
}

// IconEqualFold applies the EqualFold predicate on the "icon" field.
func IconEqualFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEqualFold(FieldIcon, v))
}

// IconContainsFold applies the ContainsFold predicate on the "icon" field.
func IconContainsFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContainsFold(FieldIcon, v))
}

// I18nKeyEQ applies the EQ predicate on the "i18n_key" field.
func I18nKeyEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldI18nKey, v))
}

// I18nKeyNEQ applies the NEQ predicate on the "i18n_key" field.
func I18nKeyNEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldI18nKey, v))
}

// I18nKeyIn applies the In predicate on the "i18n_key" field.
func I18nKeyIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldI18nKey, vs...))
}

// I18nKeyNotIn applies the NotIn predicate on the "i18n_key" field.
func I18nKeyNotIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldI18nKey, vs...))
}

// I18nKeyGT applies the GT predicate on the "i18n_key" field.
func I18nKeyGT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldI18nKey, v))
}

// I18nKeyGTE applies the GTE predicate on the "i18n_key" field.
func I18nKeyGTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldI18nKey, v))
}

// I18nKeyLT applies the LT predicate on the "i18n_key" field.
func I18nKeyLT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldI18nKey, v))
}

// I18nKeyLTE applies the LTE predicate on the "i18n_key" field.
func I18nKeyLTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldI18nKey, v))
}

// I18nKeyContains applies the Contains predicate on the "i18n_key" field.
func I18nKeyContains(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContains(FieldI18nKey, v))
}

// I18nKeyHasPrefix applies the HasPrefix predicate on the "i18n_key" field.
func I18nKeyHasPrefix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasPrefix(FieldI18nKey, v))
}

// I18nKeyHasSuffix applies the HasSuffix predicate on the "i18n_key" field.
func I18nKeyHasSuffix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasSuffix(FieldI18nKey, v))
}

// I18nKeyIsNil applies the IsNil predicate on the "i18n_key" field.
func I18nKeyIsNil() predicate.Menu {
	return predicate.Menu(sql.FieldIsNull(FieldI18nKey))
}

// I18nKeyNotNil applies the NotNil predicate on the "i18n_key" field.
func I18nKeyNotNil() predicate.Menu {
	return predicate.Menu(sql.FieldNotNull(FieldI18nKey))
}

// I18nKeyEqualFold applies the EqualFold predicate on the "i18n_key" field.
func I18nKeyEqualFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEqualFold(FieldI18nKey, v))
}

// I18nKeyContainsFold applies the ContainsFold predicate on the "i18n_key" field.
func I18nKeyContainsFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContainsFold(FieldI18nKey, v))
}

// LevelEQ applies the EQ predicate on the "level" field.
func LevelEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldLevel, v))
}

// LevelNEQ applies the NEQ predicate on the "level" field.
func LevelNEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldLevel, v))
}

// LevelIn applies the In predicate on the "level" field.
func LevelIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldLevel, vs...))
}

// LevelNotIn applies the NotIn predicate on the "level" field.
func LevelNotIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldLevel, vs...))
}

// LevelGT applies the GT predicate on the "level" field.
func LevelGT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldLevel, v))
}

// LevelGTE applies the GTE predicate on the "level" field.
func LevelGTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldLevel, v))
}

// LevelLT applies the LT predicate on the "level" field.
func LevelLT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldLevel, v))
}

// LevelLTE applies the LTE predicate on the "level" field.
func LevelLTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldLevel, v))
}

// LevelContains applies the Contains predicate on the "level" field.
func LevelContains(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContains(FieldLevel, v))
}

// LevelHasPrefix applies the HasPrefix predicate on the "level" field.
func LevelHasPrefix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasPrefix(FieldLevel, v))
}

// LevelHasSuffix applies the HasSuffix predicate on the "level" field.
func LevelHasSuffix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasSuffix(FieldLevel, v))
}

// LevelIsNil applies the IsNil predicate on the "level" field.
func LevelIsNil() predicate.Menu {
	return predicate.Menu(sql.FieldIsNull(FieldLevel))
}

// LevelNotNil applies the NotNil predicate on the "level" field.
func LevelNotNil() predicate.Menu {
	return predicate.Menu(sql.FieldNotNull(FieldLevel))
}

// LevelEqualFold applies the EqualFold predicate on the "level" field.
func LevelEqualFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEqualFold(FieldLevel, v))
}

// LevelContainsFold applies the ContainsFold predicate on the "level" field.
func LevelContainsFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContainsFold(FieldLevel, v))
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Menu) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Menu) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Menu) predicate.Menu {
	return predicate.Menu(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Menu) predicate.Menu {
	return predicate.Menu(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Menu) predicate.Menu {
	return predicate.Menu(sql.NotPredicates(p))
}