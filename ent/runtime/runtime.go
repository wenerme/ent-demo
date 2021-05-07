// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"github.com/google/uuid"
	"github.com/wenerme/ent-demo/ent/pet"
	"github.com/wenerme/ent-demo/ent/schema"
	"github.com/wenerme/ent-demo/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	petMixin := schema.Pet{}.Mixin()
	petMixinHooks0 := petMixin[0].Hooks()
	pet.Hooks[0] = petMixinHooks0[0]
	petMixinFields0 := petMixin[0].Fields()
	_ = petMixinFields0
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescUID is the schema descriptor for uid field.
	petDescUID := petMixinFields0[1].Descriptor()
	// pet.DefaultUID holds the default value on creation for the uid field.
	pet.DefaultUID = petDescUID.Default.(func() *uuid.UUID)
	userMixin := schema.User{}.Mixin()
	userMixinHooks0 := userMixin[0].Hooks()
	user.Hooks[0] = userMixinHooks0[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUID is the schema descriptor for uid field.
	userDescUID := userMixinFields0[1].Descriptor()
	// user.DefaultUID holds the default value on creation for the uid field.
	user.DefaultUID = userDescUID.Default.(func() *uuid.UUID)
}

const (
	Version = "v0.8.1-0.20210506074326-8837b53115bd"            // Version of ent codegen.
	Sum     = "h1:3a7gkaQ6/x/RJZhhDenbErKlg4bDmWqG9MZiRlNObts=" // Sum of ent codegen.
)
