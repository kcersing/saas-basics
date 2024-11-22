// Code generated by ent, DO NOT EDIT.

package ent

import (
	"company/biz/dal/mysql/ent/contract"
	"company/biz/dal/mysql/ent/entrylogs"
	"company/biz/dal/mysql/ent/face"
	"company/biz/dal/mysql/ent/member"
	"company/biz/dal/mysql/ent/messages"
	"company/biz/dal/mysql/ent/schema"
	"company/biz/dal/mysql/ent/user"
	"company/biz/dal/mysql/ent/venue"
	"company/biz/dal/mysql/ent/venueplace"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	contractMixin := schema.Contract{}.Mixin()
	contractMixinFields0 := contractMixin[0].Fields()
	_ = contractMixinFields0
	contractMixinFields1 := contractMixin[1].Fields()
	_ = contractMixinFields1
	contractFields := schema.Contract{}.Fields()
	_ = contractFields
	// contractDescCreatedAt is the schema descriptor for created_at field.
	contractDescCreatedAt := contractMixinFields0[1].Descriptor()
	// contract.DefaultCreatedAt holds the default value on creation for the created_at field.
	contract.DefaultCreatedAt = contractDescCreatedAt.Default.(func() time.Time)
	// contractDescUpdatedAt is the schema descriptor for updated_at field.
	contractDescUpdatedAt := contractMixinFields0[2].Descriptor()
	// contract.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	contract.DefaultUpdatedAt = contractDescUpdatedAt.Default.(func() time.Time)
	// contract.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	contract.UpdateDefaultUpdatedAt = contractDescUpdatedAt.UpdateDefault.(func() time.Time)
	// contractDescStatus is the schema descriptor for status field.
	contractDescStatus := contractMixinFields1[0].Descriptor()
	// contract.DefaultStatus holds the default value on creation for the status field.
	contract.DefaultStatus = contractDescStatus.Default.(int64)
	entrylogsMixin := schema.EntryLogs{}.Mixin()
	entrylogsMixinFields0 := entrylogsMixin[0].Fields()
	_ = entrylogsMixinFields0
	entrylogsFields := schema.EntryLogs{}.Fields()
	_ = entrylogsFields
	// entrylogsDescCreatedAt is the schema descriptor for created_at field.
	entrylogsDescCreatedAt := entrylogsMixinFields0[1].Descriptor()
	// entrylogs.DefaultCreatedAt holds the default value on creation for the created_at field.
	entrylogs.DefaultCreatedAt = entrylogsDescCreatedAt.Default.(func() time.Time)
	// entrylogsDescUpdatedAt is the schema descriptor for updated_at field.
	entrylogsDescUpdatedAt := entrylogsMixinFields0[2].Descriptor()
	// entrylogs.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	entrylogs.DefaultUpdatedAt = entrylogsDescUpdatedAt.Default.(func() time.Time)
	// entrylogs.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	entrylogs.UpdateDefaultUpdatedAt = entrylogsDescUpdatedAt.UpdateDefault.(func() time.Time)
	// entrylogsDescUserID is the schema descriptor for user_id field.
	entrylogsDescUserID := entrylogsFields[1].Descriptor()
	// entrylogs.DefaultUserID holds the default value on creation for the user_id field.
	entrylogs.DefaultUserID = entrylogsDescUserID.Default.(int64)
	// entrylogsDescMemberID is the schema descriptor for member_id field.
	entrylogsDescMemberID := entrylogsFields[2].Descriptor()
	// entrylogs.DefaultMemberID holds the default value on creation for the member_id field.
	entrylogs.DefaultMemberID = entrylogsDescMemberID.Default.(int64)
	faceMixin := schema.Face{}.Mixin()
	faceMixinFields0 := faceMixin[0].Fields()
	_ = faceMixinFields0
	faceFields := schema.Face{}.Fields()
	_ = faceFields
	// faceDescCreatedAt is the schema descriptor for created_at field.
	faceDescCreatedAt := faceMixinFields0[1].Descriptor()
	// face.DefaultCreatedAt holds the default value on creation for the created_at field.
	face.DefaultCreatedAt = faceDescCreatedAt.Default.(func() time.Time)
	// faceDescUpdatedAt is the schema descriptor for updated_at field.
	faceDescUpdatedAt := faceMixinFields0[2].Descriptor()
	// face.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	face.DefaultUpdatedAt = faceDescUpdatedAt.Default.(func() time.Time)
	// face.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	face.UpdateDefaultUpdatedAt = faceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// faceDescFaceIdentityCard is the schema descriptor for face_identity_card field.
	faceDescFaceIdentityCard := faceFields[3].Descriptor()
	// face.DefaultFaceIdentityCard holds the default value on creation for the face_identity_card field.
	face.DefaultFaceIdentityCard = faceDescFaceIdentityCard.Default.(string)
	// faceDescBackIdentityCard is the schema descriptor for back_identity_card field.
	faceDescBackIdentityCard := faceFields[4].Descriptor()
	// face.DefaultBackIdentityCard holds the default value on creation for the back_identity_card field.
	face.DefaultBackIdentityCard = faceDescBackIdentityCard.Default.(string)
	// faceDescFacePic is the schema descriptor for face_pic field.
	faceDescFacePic := faceFields[5].Descriptor()
	// face.DefaultFacePic holds the default value on creation for the face_pic field.
	face.DefaultFacePic = faceDescFacePic.Default.(string)
	// faceDescFaceEigenvalue is the schema descriptor for face_eigenvalue field.
	faceDescFaceEigenvalue := faceFields[6].Descriptor()
	// face.DefaultFaceEigenvalue holds the default value on creation for the face_eigenvalue field.
	face.DefaultFaceEigenvalue = faceDescFaceEigenvalue.Default.(string)
	// faceDescFacePicUpdatedTime is the schema descriptor for face_pic_updated_time field.
	faceDescFacePicUpdatedTime := faceFields[7].Descriptor()
	// face.DefaultFacePicUpdatedTime holds the default value on creation for the face_pic_updated_time field.
	face.DefaultFacePicUpdatedTime = faceDescFacePicUpdatedTime.Default.(func() time.Time)
	memberMixin := schema.Member{}.Mixin()
	memberMixinFields0 := memberMixin[0].Fields()
	_ = memberMixinFields0
	memberMixinFields1 := memberMixin[1].Fields()
	_ = memberMixinFields1
	memberFields := schema.Member{}.Fields()
	_ = memberFields
	// memberDescCreatedAt is the schema descriptor for created_at field.
	memberDescCreatedAt := memberMixinFields0[1].Descriptor()
	// member.DefaultCreatedAt holds the default value on creation for the created_at field.
	member.DefaultCreatedAt = memberDescCreatedAt.Default.(func() time.Time)
	// memberDescUpdatedAt is the schema descriptor for updated_at field.
	memberDescUpdatedAt := memberMixinFields0[2].Descriptor()
	// member.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	member.DefaultUpdatedAt = memberDescUpdatedAt.Default.(func() time.Time)
	// member.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	member.UpdateDefaultUpdatedAt = memberDescUpdatedAt.UpdateDefault.(func() time.Time)
	// memberDescStatus is the schema descriptor for status field.
	memberDescStatus := memberMixinFields1[0].Descriptor()
	// member.DefaultStatus holds the default value on creation for the status field.
	member.DefaultStatus = memberDescStatus.Default.(int64)
	// memberDescAvatar is the schema descriptor for avatar field.
	memberDescAvatar := memberFields[4].Descriptor()
	// member.DefaultAvatar holds the default value on creation for the avatar field.
	member.DefaultAvatar = memberDescAvatar.Default.(string)
	// memberDescCondition is the schema descriptor for condition field.
	memberDescCondition := memberFields[5].Descriptor()
	// member.DefaultCondition holds the default value on creation for the condition field.
	member.DefaultCondition = memberDescCondition.Default.(int64)
	messagesMixin := schema.Messages{}.Mixin()
	messagesMixinFields0 := messagesMixin[0].Fields()
	_ = messagesMixinFields0
	messagesFields := schema.Messages{}.Fields()
	_ = messagesFields
	// messagesDescCreatedAt is the schema descriptor for created_at field.
	messagesDescCreatedAt := messagesMixinFields0[1].Descriptor()
	// messages.DefaultCreatedAt holds the default value on creation for the created_at field.
	messages.DefaultCreatedAt = messagesDescCreatedAt.Default.(func() time.Time)
	// messagesDescUpdatedAt is the schema descriptor for updated_at field.
	messagesDescUpdatedAt := messagesMixinFields0[2].Descriptor()
	// messages.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	messages.DefaultUpdatedAt = messagesDescUpdatedAt.Default.(func() time.Time)
	// messages.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	messages.UpdateDefaultUpdatedAt = messagesDescUpdatedAt.UpdateDefault.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields1 := userMixin[1].Fields()
	_ = userMixinFields1
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescStatus is the schema descriptor for status field.
	userDescStatus := userMixinFields1[0].Descriptor()
	// user.DefaultStatus holds the default value on creation for the status field.
	user.DefaultStatus = userDescStatus.Default.(int64)
	// userDescSideMode is the schema descriptor for side_mode field.
	userDescSideMode := userFields[3].Descriptor()
	// user.DefaultSideMode holds the default value on creation for the side_mode field.
	user.DefaultSideMode = userDescSideMode.Default.(string)
	// userDescBaseColor is the schema descriptor for base_color field.
	userDescBaseColor := userFields[4].Descriptor()
	// user.DefaultBaseColor holds the default value on creation for the base_color field.
	user.DefaultBaseColor = userDescBaseColor.Default.(string)
	// userDescActiveColor is the schema descriptor for active_color field.
	userDescActiveColor := userFields[5].Descriptor()
	// user.DefaultActiveColor holds the default value on creation for the active_color field.
	user.DefaultActiveColor = userDescActiveColor.Default.(string)
	// userDescRoleID is the schema descriptor for role_id field.
	userDescRoleID := userFields[6].Descriptor()
	// user.DefaultRoleID holds the default value on creation for the role_id field.
	user.DefaultRoleID = userDescRoleID.Default.(int64)
	// userDescGender is the schema descriptor for gender field.
	userDescGender := userFields[14].Descriptor()
	// user.DefaultGender holds the default value on creation for the gender field.
	user.DefaultGender = userDescGender.Default.(int64)
	venueMixin := schema.Venue{}.Mixin()
	venueMixinFields0 := venueMixin[0].Fields()
	_ = venueMixinFields0
	venueMixinFields1 := venueMixin[1].Fields()
	_ = venueMixinFields1
	venueFields := schema.Venue{}.Fields()
	_ = venueFields
	// venueDescCreatedAt is the schema descriptor for created_at field.
	venueDescCreatedAt := venueMixinFields0[1].Descriptor()
	// venue.DefaultCreatedAt holds the default value on creation for the created_at field.
	venue.DefaultCreatedAt = venueDescCreatedAt.Default.(func() time.Time)
	// venueDescUpdatedAt is the schema descriptor for updated_at field.
	venueDescUpdatedAt := venueMixinFields0[2].Descriptor()
	// venue.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	venue.DefaultUpdatedAt = venueDescUpdatedAt.Default.(func() time.Time)
	// venue.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	venue.UpdateDefaultUpdatedAt = venueDescUpdatedAt.UpdateDefault.(func() time.Time)
	// venueDescStatus is the schema descriptor for status field.
	venueDescStatus := venueMixinFields1[0].Descriptor()
	// venue.DefaultStatus holds the default value on creation for the status field.
	venue.DefaultStatus = venueDescStatus.Default.(int64)
	venueplaceMixin := schema.VenuePlace{}.Mixin()
	venueplaceMixinFields0 := venueplaceMixin[0].Fields()
	_ = venueplaceMixinFields0
	venueplaceMixinFields1 := venueplaceMixin[1].Fields()
	_ = venueplaceMixinFields1
	venueplaceFields := schema.VenuePlace{}.Fields()
	_ = venueplaceFields
	// venueplaceDescCreatedAt is the schema descriptor for created_at field.
	venueplaceDescCreatedAt := venueplaceMixinFields0[1].Descriptor()
	// venueplace.DefaultCreatedAt holds the default value on creation for the created_at field.
	venueplace.DefaultCreatedAt = venueplaceDescCreatedAt.Default.(func() time.Time)
	// venueplaceDescUpdatedAt is the schema descriptor for updated_at field.
	venueplaceDescUpdatedAt := venueplaceMixinFields0[2].Descriptor()
	// venueplace.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	venueplace.DefaultUpdatedAt = venueplaceDescUpdatedAt.Default.(func() time.Time)
	// venueplace.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	venueplace.UpdateDefaultUpdatedAt = venueplaceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// venueplaceDescStatus is the schema descriptor for status field.
	venueplaceDescStatus := venueplaceMixinFields1[0].Descriptor()
	// venueplace.DefaultStatus holds the default value on creation for the status field.
	venueplace.DefaultStatus = venueplaceDescStatus.Default.(int64)
}
