// Code generated by entc, DO NOT EDIT.

package eventtype

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the eventtype type in the database.
	Label = "event_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldColor holds the string denoting the color field in the database.
	FieldColor = "color"
	// EdgeEvents holds the string denoting the events edge name in mutations.
	EdgeEvents = "events"
	// Table holds the table name of the eventtype in the database.
	Table = "event_types"
	// EventsTable is the table that holds the events relation/edge.
	EventsTable = "events"
	// EventsInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventsInverseTable = "events"
	// EventsColumn is the table column denoting the events relation/edge.
	EventsColumn = "event_type_events"
)

// Columns holds all SQL columns for eventtype fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldColor,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
