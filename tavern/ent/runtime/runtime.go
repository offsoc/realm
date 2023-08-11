// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/kcarretto/realm/tavern/ent/beacon"
	"github.com/kcarretto/realm/tavern/ent/file"
	"github.com/kcarretto/realm/tavern/ent/quest"
	"github.com/kcarretto/realm/tavern/ent/schema"
	"github.com/kcarretto/realm/tavern/ent/tag"
	"github.com/kcarretto/realm/tavern/ent/task"
	"github.com/kcarretto/realm/tavern/ent/tome"
	"github.com/kcarretto/realm/tavern/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	beaconFields := schema.Beacon{}.Fields()
	_ = beaconFields
	// beaconDescName is the schema descriptor for name field.
	beaconDescName := beaconFields[0].Descriptor()
	// beacon.DefaultName holds the default value on creation for the name field.
	beacon.DefaultName = beaconDescName.Default.(func() string)
	// beacon.NameValidator is a validator for the "name" field. It is called by the builders before save.
	beacon.NameValidator = beaconDescName.Validators[0].(func(string) error)
	// beaconDescPrincipal is the schema descriptor for principal field.
	beaconDescPrincipal := beaconFields[1].Descriptor()
	// beacon.PrincipalValidator is a validator for the "principal" field. It is called by the builders before save.
	beacon.PrincipalValidator = beaconDescPrincipal.Validators[0].(func(string) error)
	// beaconDescHostname is the schema descriptor for hostname field.
	beaconDescHostname := beaconFields[2].Descriptor()
	// beacon.HostnameValidator is a validator for the "hostname" field. It is called by the builders before save.
	beacon.HostnameValidator = beaconDescHostname.Validators[0].(func(string) error)
	// beaconDescIdentifier is the schema descriptor for identifier field.
	beaconDescIdentifier := beaconFields[3].Descriptor()
	// beacon.DefaultIdentifier holds the default value on creation for the identifier field.
	beacon.DefaultIdentifier = beaconDescIdentifier.Default.(func() string)
	// beacon.IdentifierValidator is a validator for the "identifier" field. It is called by the builders before save.
	beacon.IdentifierValidator = beaconDescIdentifier.Validators[0].(func(string) error)
	// beaconDescAgentIdentifier is the schema descriptor for agent_identifier field.
	beaconDescAgentIdentifier := beaconFields[4].Descriptor()
	// beacon.AgentIdentifierValidator is a validator for the "agent_identifier" field. It is called by the builders before save.
	beacon.AgentIdentifierValidator = beaconDescAgentIdentifier.Validators[0].(func(string) error)
	// beaconDescHostIdentifier is the schema descriptor for host_identifier field.
	beaconDescHostIdentifier := beaconFields[5].Descriptor()
	// beacon.HostIdentifierValidator is a validator for the "host_identifier" field. It is called by the builders before save.
	beacon.HostIdentifierValidator = beaconDescHostIdentifier.Validators[0].(func(string) error)
	fileMixin := schema.File{}.Mixin()
	fileHooks := schema.File{}.Hooks()
	file.Hooks[0] = fileHooks[0]
	fileMixinFields0 := fileMixin[0].Fields()
	_ = fileMixinFields0
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescCreatedAt is the schema descriptor for created_at field.
	fileDescCreatedAt := fileMixinFields0[0].Descriptor()
	// file.DefaultCreatedAt holds the default value on creation for the created_at field.
	file.DefaultCreatedAt = fileDescCreatedAt.Default.(func() time.Time)
	// fileDescLastModifiedAt is the schema descriptor for last_modified_at field.
	fileDescLastModifiedAt := fileMixinFields0[1].Descriptor()
	// file.DefaultLastModifiedAt holds the default value on creation for the last_modified_at field.
	file.DefaultLastModifiedAt = fileDescLastModifiedAt.Default.(func() time.Time)
	// file.UpdateDefaultLastModifiedAt holds the default value on update for the last_modified_at field.
	file.UpdateDefaultLastModifiedAt = fileDescLastModifiedAt.UpdateDefault.(func() time.Time)
	// fileDescName is the schema descriptor for name field.
	fileDescName := fileFields[0].Descriptor()
	// file.NameValidator is a validator for the "name" field. It is called by the builders before save.
	file.NameValidator = fileDescName.Validators[0].(func(string) error)
	// fileDescSize is the schema descriptor for size field.
	fileDescSize := fileFields[1].Descriptor()
	// file.DefaultSize holds the default value on creation for the size field.
	file.DefaultSize = fileDescSize.Default.(int)
	// file.SizeValidator is a validator for the "size" field. It is called by the builders before save.
	file.SizeValidator = fileDescSize.Validators[0].(func(int) error)
	// fileDescHash is the schema descriptor for hash field.
	fileDescHash := fileFields[2].Descriptor()
	// file.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	file.HashValidator = fileDescHash.Validators[0].(func(string) error)
	questMixin := schema.Quest{}.Mixin()
	questMixinFields0 := questMixin[0].Fields()
	_ = questMixinFields0
	questFields := schema.Quest{}.Fields()
	_ = questFields
	// questDescCreatedAt is the schema descriptor for created_at field.
	questDescCreatedAt := questMixinFields0[0].Descriptor()
	// quest.DefaultCreatedAt holds the default value on creation for the created_at field.
	quest.DefaultCreatedAt = questDescCreatedAt.Default.(func() time.Time)
	// questDescLastModifiedAt is the schema descriptor for last_modified_at field.
	questDescLastModifiedAt := questMixinFields0[1].Descriptor()
	// quest.DefaultLastModifiedAt holds the default value on creation for the last_modified_at field.
	quest.DefaultLastModifiedAt = questDescLastModifiedAt.Default.(func() time.Time)
	// quest.UpdateDefaultLastModifiedAt holds the default value on update for the last_modified_at field.
	quest.UpdateDefaultLastModifiedAt = questDescLastModifiedAt.UpdateDefault.(func() time.Time)
	// questDescName is the schema descriptor for name field.
	questDescName := questFields[0].Descriptor()
	// quest.NameValidator is a validator for the "name" field. It is called by the builders before save.
	quest.NameValidator = questDescName.Validators[0].(func(string) error)
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescName is the schema descriptor for name field.
	tagDescName := tagFields[0].Descriptor()
	// tag.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tag.NameValidator = tagDescName.Validators[0].(func(string) error)
	taskMixin := schema.Task{}.Mixin()
	taskMixinFields0 := taskMixin[0].Fields()
	_ = taskMixinFields0
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescCreatedAt is the schema descriptor for created_at field.
	taskDescCreatedAt := taskMixinFields0[0].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the created_at field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(func() time.Time)
	// taskDescLastModifiedAt is the schema descriptor for last_modified_at field.
	taskDescLastModifiedAt := taskMixinFields0[1].Descriptor()
	// task.DefaultLastModifiedAt holds the default value on creation for the last_modified_at field.
	task.DefaultLastModifiedAt = taskDescLastModifiedAt.Default.(func() time.Time)
	// task.UpdateDefaultLastModifiedAt holds the default value on update for the last_modified_at field.
	task.UpdateDefaultLastModifiedAt = taskDescLastModifiedAt.UpdateDefault.(func() time.Time)
	tomeMixin := schema.Tome{}.Mixin()
	tomeHooks := schema.Tome{}.Hooks()
	tome.Hooks[0] = tomeHooks[0]
	tomeMixinFields0 := tomeMixin[0].Fields()
	_ = tomeMixinFields0
	tomeFields := schema.Tome{}.Fields()
	_ = tomeFields
	// tomeDescCreatedAt is the schema descriptor for created_at field.
	tomeDescCreatedAt := tomeMixinFields0[0].Descriptor()
	// tome.DefaultCreatedAt holds the default value on creation for the created_at field.
	tome.DefaultCreatedAt = tomeDescCreatedAt.Default.(func() time.Time)
	// tomeDescLastModifiedAt is the schema descriptor for last_modified_at field.
	tomeDescLastModifiedAt := tomeMixinFields0[1].Descriptor()
	// tome.DefaultLastModifiedAt holds the default value on creation for the last_modified_at field.
	tome.DefaultLastModifiedAt = tomeDescLastModifiedAt.Default.(func() time.Time)
	// tome.UpdateDefaultLastModifiedAt holds the default value on update for the last_modified_at field.
	tome.UpdateDefaultLastModifiedAt = tomeDescLastModifiedAt.UpdateDefault.(func() time.Time)
	// tomeDescName is the schema descriptor for name field.
	tomeDescName := tomeFields[0].Descriptor()
	// tome.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tome.NameValidator = tomeDescName.Validators[0].(func(string) error)
	// tomeDescHash is the schema descriptor for hash field.
	tomeDescHash := tomeFields[3].Descriptor()
	// tome.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	tome.HashValidator = tomeDescHash.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = func() func(string) error {
		validators := userDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescSessionToken is the schema descriptor for session_token field.
	userDescSessionToken := userFields[3].Descriptor()
	// user.DefaultSessionToken holds the default value on creation for the session_token field.
	user.DefaultSessionToken = userDescSessionToken.Default.(func() string)
	// user.SessionTokenValidator is a validator for the "session_token" field. It is called by the builders before save.
	user.SessionTokenValidator = userDescSessionToken.Validators[0].(func(string) error)
	// userDescIsActivated is the schema descriptor for is_activated field.
	userDescIsActivated := userFields[4].Descriptor()
	// user.DefaultIsActivated holds the default value on creation for the is_activated field.
	user.DefaultIsActivated = userDescIsActivated.Default.(bool)
	// userDescIsAdmin is the schema descriptor for is_admin field.
	userDescIsAdmin := userFields[5].Descriptor()
	// user.DefaultIsAdmin holds the default value on creation for the is_admin field.
	user.DefaultIsAdmin = userDescIsAdmin.Default.(bool)
}

const (
	Version = "v0.11.9"                                         // Version of ent codegen.
	Sum     = "h1:dbbCkAiPVTRBIJwoZctiSYjB7zxQIBOzVSU5H9VYIQI=" // Sum of ent codegen.
)
