// Code generated by ent, DO NOT EDIT.

package ent

import (
	"saas/db/ent/schema"
	"saas/db/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescGender is the schema descriptor for gender field.
	userDescGender := userFields[3].Descriptor()
	// user.DefaultGender holds the default value on creation for the gender field.
	user.DefaultGender = userDescGender.Default.(int)
}