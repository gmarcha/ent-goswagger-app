// Code generated by entc, DO NOT EDIT.

package event

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the event type in the database.
	Label = "event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldTutorsRequired holds the string denoting the tutorsrequired field in the database.
	FieldTutorsRequired = "tutors_required"
	// FieldWalletsReward holds the string denoting the walletsreward field in the database.
	FieldWalletsReward = "wallets_reward"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldStartAt holds the string denoting the startat field in the database.
	FieldStartAt = "start_at"
	// FieldEndAt holds the string denoting the endat field in the database.
	FieldEndAt = "end_at"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeCategory holds the string denoting the category edge name in mutations.
	EdgeCategory = "category"
	// Table holds the table name of the event in the database.
	Table = "events"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "event_users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// CategoryTable is the table that holds the category relation/edge.
	CategoryTable = "events"
	// CategoryInverseTable is the table name for the EventType entity.
	// It exists in this package in order to avoid circular dependency with the "eventtype" package.
	CategoryInverseTable = "event_types"
	// CategoryColumn is the table column denoting the category relation/edge.
	CategoryColumn = "event_type_events"
)

// Columns holds all SQL columns for event fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldTutorsRequired,
	FieldWalletsReward,
	FieldCreatedAt,
	FieldStartAt,
	FieldEndAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "events"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"event_type_events",
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"event_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// TutorsRequiredValidator is a validator for the "tutorsRequired" field. It is called by the builders before save.
	TutorsRequiredValidator func(int64) error
	// WalletsRewardValidator is a validator for the "walletsReward" field. It is called by the builders before save.
	WalletsRewardValidator func(int64) error
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
