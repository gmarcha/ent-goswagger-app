// Code generated by entc, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Login applies equality check predicate on the "login" field. It's identical to LoginEQ.
func Login(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogin), v))
	})
}

// FirstName applies equality check predicate on the "firstName" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFirstName), v))
	})
}

// LastName applies equality check predicate on the "lastName" field. It's identical to LastNameEQ.
func LastName(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastName), v))
	})
}

// DisplayName applies equality check predicate on the "displayName" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisplayName), v))
	})
}

// ImagePath applies equality check predicate on the "imagePath" field. It's identical to ImagePathEQ.
func ImagePath(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImagePath), v))
	})
}

// LoginEQ applies the EQ predicate on the "login" field.
func LoginEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogin), v))
	})
}

// LoginNEQ applies the NEQ predicate on the "login" field.
func LoginNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLogin), v))
	})
}

// LoginIn applies the In predicate on the "login" field.
func LoginIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLogin), v...))
	})
}

// LoginNotIn applies the NotIn predicate on the "login" field.
func LoginNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLogin), v...))
	})
}

// LoginGT applies the GT predicate on the "login" field.
func LoginGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLogin), v))
	})
}

// LoginGTE applies the GTE predicate on the "login" field.
func LoginGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLogin), v))
	})
}

// LoginLT applies the LT predicate on the "login" field.
func LoginLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLogin), v))
	})
}

// LoginLTE applies the LTE predicate on the "login" field.
func LoginLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLogin), v))
	})
}

// LoginContains applies the Contains predicate on the "login" field.
func LoginContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLogin), v))
	})
}

// LoginHasPrefix applies the HasPrefix predicate on the "login" field.
func LoginHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLogin), v))
	})
}

// LoginHasSuffix applies the HasSuffix predicate on the "login" field.
func LoginHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLogin), v))
	})
}

// LoginEqualFold applies the EqualFold predicate on the "login" field.
func LoginEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLogin), v))
	})
}

// LoginContainsFold applies the ContainsFold predicate on the "login" field.
func LoginContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLogin), v))
	})
}

// FirstNameEQ applies the EQ predicate on the "firstName" field.
func FirstNameEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFirstName), v))
	})
}

// FirstNameNEQ applies the NEQ predicate on the "firstName" field.
func FirstNameNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFirstName), v))
	})
}

// FirstNameIn applies the In predicate on the "firstName" field.
func FirstNameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFirstName), v...))
	})
}

// FirstNameNotIn applies the NotIn predicate on the "firstName" field.
func FirstNameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFirstName), v...))
	})
}

// FirstNameGT applies the GT predicate on the "firstName" field.
func FirstNameGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFirstName), v))
	})
}

// FirstNameGTE applies the GTE predicate on the "firstName" field.
func FirstNameGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFirstName), v))
	})
}

// FirstNameLT applies the LT predicate on the "firstName" field.
func FirstNameLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFirstName), v))
	})
}

// FirstNameLTE applies the LTE predicate on the "firstName" field.
func FirstNameLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFirstName), v))
	})
}

// FirstNameContains applies the Contains predicate on the "firstName" field.
func FirstNameContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFirstName), v))
	})
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "firstName" field.
func FirstNameHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFirstName), v))
	})
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "firstName" field.
func FirstNameHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFirstName), v))
	})
}

// FirstNameIsNil applies the IsNil predicate on the "firstName" field.
func FirstNameIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFirstName)))
	})
}

// FirstNameNotNil applies the NotNil predicate on the "firstName" field.
func FirstNameNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFirstName)))
	})
}

// FirstNameEqualFold applies the EqualFold predicate on the "firstName" field.
func FirstNameEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFirstName), v))
	})
}

// FirstNameContainsFold applies the ContainsFold predicate on the "firstName" field.
func FirstNameContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFirstName), v))
	})
}

// LastNameEQ applies the EQ predicate on the "lastName" field.
func LastNameEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastName), v))
	})
}

// LastNameNEQ applies the NEQ predicate on the "lastName" field.
func LastNameNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastName), v))
	})
}

// LastNameIn applies the In predicate on the "lastName" field.
func LastNameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastName), v...))
	})
}

// LastNameNotIn applies the NotIn predicate on the "lastName" field.
func LastNameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastName), v...))
	})
}

// LastNameGT applies the GT predicate on the "lastName" field.
func LastNameGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastName), v))
	})
}

// LastNameGTE applies the GTE predicate on the "lastName" field.
func LastNameGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastName), v))
	})
}

// LastNameLT applies the LT predicate on the "lastName" field.
func LastNameLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastName), v))
	})
}

// LastNameLTE applies the LTE predicate on the "lastName" field.
func LastNameLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastName), v))
	})
}

// LastNameContains applies the Contains predicate on the "lastName" field.
func LastNameContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLastName), v))
	})
}

// LastNameHasPrefix applies the HasPrefix predicate on the "lastName" field.
func LastNameHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLastName), v))
	})
}

// LastNameHasSuffix applies the HasSuffix predicate on the "lastName" field.
func LastNameHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLastName), v))
	})
}

// LastNameIsNil applies the IsNil predicate on the "lastName" field.
func LastNameIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLastName)))
	})
}

// LastNameNotNil applies the NotNil predicate on the "lastName" field.
func LastNameNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLastName)))
	})
}

// LastNameEqualFold applies the EqualFold predicate on the "lastName" field.
func LastNameEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLastName), v))
	})
}

// LastNameContainsFold applies the ContainsFold predicate on the "lastName" field.
func LastNameContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLastName), v))
	})
}

// DisplayNameEQ applies the EQ predicate on the "displayName" field.
func DisplayNameEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisplayName), v))
	})
}

// DisplayNameNEQ applies the NEQ predicate on the "displayName" field.
func DisplayNameNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDisplayName), v))
	})
}

// DisplayNameIn applies the In predicate on the "displayName" field.
func DisplayNameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDisplayName), v...))
	})
}

// DisplayNameNotIn applies the NotIn predicate on the "displayName" field.
func DisplayNameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDisplayName), v...))
	})
}

// DisplayNameGT applies the GT predicate on the "displayName" field.
func DisplayNameGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDisplayName), v))
	})
}

// DisplayNameGTE applies the GTE predicate on the "displayName" field.
func DisplayNameGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDisplayName), v))
	})
}

// DisplayNameLT applies the LT predicate on the "displayName" field.
func DisplayNameLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDisplayName), v))
	})
}

// DisplayNameLTE applies the LTE predicate on the "displayName" field.
func DisplayNameLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDisplayName), v))
	})
}

// DisplayNameContains applies the Contains predicate on the "displayName" field.
func DisplayNameContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDisplayName), v))
	})
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "displayName" field.
func DisplayNameHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDisplayName), v))
	})
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "displayName" field.
func DisplayNameHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDisplayName), v))
	})
}

// DisplayNameIsNil applies the IsNil predicate on the "displayName" field.
func DisplayNameIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDisplayName)))
	})
}

// DisplayNameNotNil applies the NotNil predicate on the "displayName" field.
func DisplayNameNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDisplayName)))
	})
}

// DisplayNameEqualFold applies the EqualFold predicate on the "displayName" field.
func DisplayNameEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDisplayName), v))
	})
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "displayName" field.
func DisplayNameContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDisplayName), v))
	})
}

// ImagePathEQ applies the EQ predicate on the "imagePath" field.
func ImagePathEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImagePath), v))
	})
}

// ImagePathNEQ applies the NEQ predicate on the "imagePath" field.
func ImagePathNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldImagePath), v))
	})
}

// ImagePathIn applies the In predicate on the "imagePath" field.
func ImagePathIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldImagePath), v...))
	})
}

// ImagePathNotIn applies the NotIn predicate on the "imagePath" field.
func ImagePathNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldImagePath), v...))
	})
}

// ImagePathGT applies the GT predicate on the "imagePath" field.
func ImagePathGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldImagePath), v))
	})
}

// ImagePathGTE applies the GTE predicate on the "imagePath" field.
func ImagePathGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldImagePath), v))
	})
}

// ImagePathLT applies the LT predicate on the "imagePath" field.
func ImagePathLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldImagePath), v))
	})
}

// ImagePathLTE applies the LTE predicate on the "imagePath" field.
func ImagePathLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldImagePath), v))
	})
}

// ImagePathContains applies the Contains predicate on the "imagePath" field.
func ImagePathContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldImagePath), v))
	})
}

// ImagePathHasPrefix applies the HasPrefix predicate on the "imagePath" field.
func ImagePathHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldImagePath), v))
	})
}

// ImagePathHasSuffix applies the HasSuffix predicate on the "imagePath" field.
func ImagePathHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldImagePath), v))
	})
}

// ImagePathIsNil applies the IsNil predicate on the "imagePath" field.
func ImagePathIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldImagePath)))
	})
}

// ImagePathNotNil applies the NotNil predicate on the "imagePath" field.
func ImagePathNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldImagePath)))
	})
}

// ImagePathEqualFold applies the EqualFold predicate on the "imagePath" field.
func ImagePathEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldImagePath), v))
	})
}

// ImagePathContainsFold applies the ContainsFold predicate on the "imagePath" field.
func ImagePathContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldImagePath), v))
	})
}

// HasRoles applies the HasEdge predicate on the "roles" edge.
func HasRoles() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RolesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RolesTable, RolesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRolesWith applies the HasEdge predicate on the "roles" edge with a given conditions (other predicates).
func HasRolesWith(preds ...predicate.Role) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RolesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RolesTable, RolesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEvents applies the HasEdge predicate on the "events" edge.
func HasEvents() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, EventsTable, EventsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventsWith applies the HasEdge predicate on the "events" edge with a given conditions (other predicates).
func HasEventsWith(preds ...predicate.Event) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, EventsTable, EventsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
