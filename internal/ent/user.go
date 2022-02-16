// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/gamarcha/ent-goswagger-app/internal/ent/user"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Login holds the value of the "login" field.
	Login string `json:"login,omitempty"`
	// FirstName holds the value of the "firstName" field.
	FirstName string `json:"firstName,omitempty"`
	// LastName holds the value of the "lastName" field.
	LastName string `json:"lastName,omitempty"`
	// HoursDone holds the value of the "hoursDone" field.
	HoursDone int64 `json:"hoursDone,omitempty"`
	// TutorScope holds the value of the "tutorScope" field.
	TutorScope bool `json:"tutorScope,omitempty"`
	// AdminScope holds the value of the "adminScope" field.
	AdminScope bool `json:"adminScope,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Events holds the value of the events edge.
	Events []*Event `json:"events,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EventsOrErr returns the Events value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) EventsOrErr() ([]*Event, error) {
	if e.loadedTypes[0] {
		return e.Events, nil
	}
	return nil, &NotLoadedError{edge: "events"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldTutorScope, user.FieldAdminScope:
			values[i] = new(sql.NullBool)
		case user.FieldHoursDone:
			values[i] = new(sql.NullInt64)
		case user.FieldLogin, user.FieldFirstName, user.FieldLastName:
			values[i] = new(sql.NullString)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldLogin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field login", values[i])
			} else if value.Valid {
				u.Login = value.String
			}
		case user.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field firstName", values[i])
			} else if value.Valid {
				u.FirstName = value.String
			}
		case user.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lastName", values[i])
			} else if value.Valid {
				u.LastName = value.String
			}
		case user.FieldHoursDone:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field hoursDone", values[i])
			} else if value.Valid {
				u.HoursDone = value.Int64
			}
		case user.FieldTutorScope:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field tutorScope", values[i])
			} else if value.Valid {
				u.TutorScope = value.Bool
			}
		case user.FieldAdminScope:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field adminScope", values[i])
			} else if value.Valid {
				u.AdminScope = value.Bool
			}
		}
	}
	return nil
}

// QueryEvents queries the "events" edge of the User entity.
func (u *User) QueryEvents() *EventQuery {
	return (&UserClient{config: u.config}).QueryEvents(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", login=")
	builder.WriteString(u.Login)
	builder.WriteString(", firstName=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", lastName=")
	builder.WriteString(u.LastName)
	builder.WriteString(", hoursDone=")
	builder.WriteString(fmt.Sprintf("%v", u.HoursDone))
	builder.WriteString(", tutorScope=")
	builder.WriteString(fmt.Sprintf("%v", u.TutorScope))
	builder.WriteString(", adminScope=")
	builder.WriteString(fmt.Sprintf("%v", u.AdminScope))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}