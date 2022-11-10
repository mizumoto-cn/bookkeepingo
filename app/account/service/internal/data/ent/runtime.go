// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/mizumoto-cn/bookkeepingo/app/account/service/internal/data/ent/account"
	"github.com/mizumoto-cn/bookkeepingo/app/account/service/internal/data/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescCreatedAt is the schema descriptor for created_at field.
	accountDescCreatedAt := accountFields[3].Descriptor()
	// account.DefaultCreatedAt holds the default value on creation for the created_at field.
	account.DefaultCreatedAt = accountDescCreatedAt.Default.(func() time.Time)
	// accountDescUpdatedAt is the schema descriptor for updated_at field.
	accountDescUpdatedAt := accountFields[4].Descriptor()
	// account.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	account.DefaultUpdatedAt = accountDescUpdatedAt.Default.(func() time.Time)
	// accountDescDeleted is the schema descriptor for deleted field.
	accountDescDeleted := accountFields[7].Descriptor()
	// account.DefaultDeleted holds the default value on creation for the deleted field.
	account.DefaultDeleted = accountDescDeleted.Default.(bool)
}
