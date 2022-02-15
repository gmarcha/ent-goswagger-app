// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gamarcha/ent-goswagger-app/ent/event"
	"github.com/google/uuid"
)

// Event is the model entity for the Event schema.
type Event struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// TutorsRequired holds the value of the "tutorsRequired" field.
	TutorsRequired int64 `json:"tutorsRequired,omitempty"`
	// TutorsSubscribed holds the value of the "tutorsSubscribed" field.
	TutorsSubscribed int64 `json:"tutorsSubscribed,omitempty"`
	// WalletsRewards holds the value of the "walletsRewards" field.
	WalletsRewards int64 `json:"walletsRewards,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// StartAt holds the value of the "startAt" field.
	StartAt time.Time `json:"startAt,omitempty"`
	// EndAt holds the value of the "endAt" field.
	EndAt time.Time `json:"endAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EventQuery when eager-loading is set.
	Edges EventEdges `json:"edges"`
}

// EventEdges holds the relations/edges for other nodes in the graph.
type EventEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e EventEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Event) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case event.FieldTutorsRequired, event.FieldTutorsSubscribed, event.FieldWalletsRewards:
			values[i] = new(sql.NullInt64)
		case event.FieldName, event.FieldDescription:
			values[i] = new(sql.NullString)
		case event.FieldCreatedAt, event.FieldStartAt, event.FieldEndAt:
			values[i] = new(sql.NullTime)
		case event.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Event", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Event fields.
func (e *Event) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case event.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				e.ID = *value
			}
		case event.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		case event.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				e.Description = value.String
			}
		case event.FieldTutorsRequired:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tutorsRequired", values[i])
			} else if value.Valid {
				e.TutorsRequired = value.Int64
			}
		case event.FieldTutorsSubscribed:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tutorsSubscribed", values[i])
			} else if value.Valid {
				e.TutorsSubscribed = value.Int64
			}
		case event.FieldWalletsRewards:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field walletsRewards", values[i])
			} else if value.Valid {
				e.WalletsRewards = value.Int64
			}
		case event.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Time
			}
		case event.FieldStartAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field startAt", values[i])
			} else if value.Valid {
				e.StartAt = value.Time
			}
		case event.FieldEndAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field endAt", values[i])
			} else if value.Valid {
				e.EndAt = value.Time
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the Event entity.
func (e *Event) QueryUsers() *UserQuery {
	return (&EventClient{config: e.config}).QueryUsers(e)
}

// Update returns a builder for updating this Event.
// Note that you need to call Event.Unwrap() before calling this method if this Event
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Event) Update() *EventUpdateOne {
	return (&EventClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Event entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Event) Unwrap() *Event {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Event is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Event) String() string {
	var builder strings.Builder
	builder.WriteString("Event(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteString(", name=")
	builder.WriteString(e.Name)
	builder.WriteString(", description=")
	builder.WriteString(e.Description)
	builder.WriteString(", tutorsRequired=")
	builder.WriteString(fmt.Sprintf("%v", e.TutorsRequired))
	builder.WriteString(", tutorsSubscribed=")
	builder.WriteString(fmt.Sprintf("%v", e.TutorsSubscribed))
	builder.WriteString(", walletsRewards=")
	builder.WriteString(fmt.Sprintf("%v", e.WalletsRewards))
	builder.WriteString(", createdAt=")
	builder.WriteString(e.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", startAt=")
	builder.WriteString(e.StartAt.Format(time.ANSIC))
	builder.WriteString(", endAt=")
	builder.WriteString(e.EndAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Events is a parsable slice of Event.
type Events []*Event

func (e Events) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
