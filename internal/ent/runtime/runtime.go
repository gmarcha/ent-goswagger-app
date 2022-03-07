// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/gmarcha/ent-goswagger-app/internal/ent/event"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/eventtype"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/role"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/schema"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	eventHooks := schema.Event{}.Hooks()
	event.Hooks[0] = eventHooks[0]
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescName is the schema descriptor for name field.
	eventDescName := eventFields[1].Descriptor()
	// event.NameValidator is a validator for the "name" field. It is called by the builders before save.
	event.NameValidator = eventDescName.Validators[0].(func(string) error)
	// eventDescTutorsRequired is the schema descriptor for tutorsRequired field.
	eventDescTutorsRequired := eventFields[3].Descriptor()
	// event.TutorsRequiredValidator is a validator for the "tutorsRequired" field. It is called by the builders before save.
	event.TutorsRequiredValidator = eventDescTutorsRequired.Validators[0].(func(int64) error)
	// eventDescWalletsReward is the schema descriptor for walletsReward field.
	eventDescWalletsReward := eventFields[4].Descriptor()
	// event.WalletsRewardValidator is a validator for the "walletsReward" field. It is called by the builders before save.
	event.WalletsRewardValidator = eventDescWalletsReward.Validators[0].(func(int64) error)
	// eventDescCreatedAt is the schema descriptor for createdAt field.
	eventDescCreatedAt := eventFields[5].Descriptor()
	// event.DefaultCreatedAt holds the default value on creation for the createdAt field.
	event.DefaultCreatedAt = eventDescCreatedAt.Default.(func() time.Time)
	// eventDescID is the schema descriptor for id field.
	eventDescID := eventFields[0].Descriptor()
	// event.DefaultID holds the default value on creation for the id field.
	event.DefaultID = eventDescID.Default.(func() uuid.UUID)
	eventtypeFields := schema.EventType{}.Fields()
	_ = eventtypeFields
	// eventtypeDescID is the schema descriptor for id field.
	eventtypeDescID := eventtypeFields[0].Descriptor()
	// eventtype.DefaultID holds the default value on creation for the id field.
	eventtype.DefaultID = eventtypeDescID.Default.(func() uuid.UUID)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescEvent is the schema descriptor for event field.
	roleDescEvent := roleFields[2].Descriptor()
	// role.DefaultEvent holds the default value on creation for the event field.
	role.DefaultEvent = roleDescEvent.Default.(string)
	// roleDescEventWrite is the schema descriptor for event_write field.
	roleDescEventWrite := roleFields[3].Descriptor()
	// role.DefaultEventWrite holds the default value on creation for the event_write field.
	role.DefaultEventWrite = roleDescEventWrite.Default.(string)
	// roleDescUser is the schema descriptor for user field.
	roleDescUser := roleFields[4].Descriptor()
	// role.DefaultUser holds the default value on creation for the user field.
	role.DefaultUser = roleDescUser.Default.(string)
	// roleDescUserSubscription is the schema descriptor for user_subscription field.
	roleDescUserSubscription := roleFields[5].Descriptor()
	// role.DefaultUserSubscription holds the default value on creation for the user_subscription field.
	role.DefaultUserSubscription = roleDescUserSubscription.Default.(string)
	// roleDescUserWrite is the schema descriptor for user_write field.
	roleDescUserWrite := roleFields[6].Descriptor()
	// role.DefaultUserWrite holds the default value on creation for the user_write field.
	role.DefaultUserWrite = roleDescUserWrite.Default.(string)
	// roleDescID is the schema descriptor for id field.
	roleDescID := roleFields[0].Descriptor()
	// role.DefaultID holds the default value on creation for the id field.
	role.DefaultID = roleDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescLogin is the schema descriptor for login field.
	userDescLogin := userFields[1].Descriptor()
	// user.LoginValidator is a validator for the "login" field. It is called by the builders before save.
	user.LoginValidator = userDescLogin.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.10.1"                                         // Version of ent codegen.
	Sum     = "h1:dM5h4Zk6yHGIgw4dCqVzGw3nWgpGYJiV4/kyHEF6PFo=" // Sum of ent codegen.
)
