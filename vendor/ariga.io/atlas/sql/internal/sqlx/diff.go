// Copyright 2021-present The Atlas Authors. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sqlx

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"ariga.io/atlas/sql/schema"
)

type (
	// A Diff provides a generic schema.Differ for diffing schema elements.
	//
	// The DiffDriver is required for supporting database/dialect specific
	// diff capabilities, like diffing custom types or attributes.
	Diff struct {
		DiffDriver
	}

	// A DiffDriver wraps all required methods for diffing elements that may
	// have database-specific diff logic. See sql/schema/mysql/diff.go for an
	// implementation example.
	DiffDriver interface {
		// SchemaAttrDiff returns a changeset for migrating schema attributes
		// from one state to the other. For example, changing schema collation.
		SchemaAttrDiff(from, to *schema.Schema) []schema.Change

		// TableAttrDiff returns a changeset for migrating table attributes from
		// one state to the other. For example, dropping or adding a `CHECK` constraint.
		TableAttrDiff(from, to *schema.Table) ([]schema.Change, error)

		// ColumnChange returns the schema changes (if any) for migrating one column to the other.
		ColumnChange(from, to *schema.Column) (schema.ChangeKind, error)

		// IndexAttrChanged reports if the index attributes were changed.
		// For example, an index type or predicate (for partial indexes).
		IndexAttrChanged(from, to []schema.Attr) bool

		// IndexPartAttrChanged reports if the index-part attributes were
		// changed. For example, an index-part collation.
		IndexPartAttrChanged(from, to *schema.IndexPart) bool

		// IsGeneratedIndexName reports if the index name was generated by the database
		// for unnamed INDEX or UNIQUE constraints. In such cases, the differ will look
		// for unnamed schema.Indexes on the desired state, before tagging the index as
		// a candidate for deletion.
		IsGeneratedIndexName(*schema.Table, *schema.Index) bool

		// ReferenceChanged reports if the foreign key referential action was
		// changed. For example, action was changed from RESTRICT to CASCADE.
		ReferenceChanged(from, to schema.ReferenceOption) bool
	}

	// A Normalizer wraps the Normalize method for normalizing the from and to tables before
	// running diffing. The "from" usually represents the inspected database state (current),
	// and the second represents the desired state.
	//
	// If the DiffDriver implements the Normalizer interface, TableDiff normalizes its table
	// inputs before starting the diff process.
	Normalizer interface {
		Normalize(from, to *schema.Table)
	}
)

// RealmDiff implements the schema.Differ for Realm objects and returns a list of changes
// that need to be applied in order to move a database from the current state to the desired.
func (d *Diff) RealmDiff(from, to *schema.Realm) ([]schema.Change, error) {
	var changes []schema.Change
	// Drop or modify schema.
	for _, s1 := range from.Schemas {
		s2, ok := to.Schema(s1.Name)
		if !ok {
			changes = append(changes, &schema.DropSchema{S: s1})
			continue
		}
		change, err := d.SchemaDiff(s1, s2)
		if err != nil {
			return nil, err
		}
		changes = append(changes, change...)
	}
	// Add schemas.
	for _, s1 := range to.Schemas {
		if _, ok := from.Schema(s1.Name); ok {
			continue
		}
		changes = append(changes, &schema.AddSchema{S: s1})
		for _, t := range s1.Tables {
			changes = append(changes, &schema.AddTable{T: t})
		}
	}
	return changes, nil
}

// SchemaDiff implements the schema.Differ interface and returns a list of
// changes that need to be applied in order to move from one state to the other.
func (d *Diff) SchemaDiff(from, to *schema.Schema) ([]schema.Change, error) {
	if from.Name != to.Name {
		return nil, fmt.Errorf("mismatched schema names: %q != %q", from.Name, to.Name)
	}
	var changes []schema.Change
	// Drop or modify attributes (collations, charset, etc).
	if change := d.SchemaAttrDiff(from, to); len(change) > 0 {
		changes = append(changes, &schema.ModifySchema{
			S:       to,
			Changes: change,
		})
	}

	// Drop or modify tables.
	for _, t1 := range from.Tables {
		t2, ok := to.Table(t1.Name)
		if !ok {
			changes = append(changes, &schema.DropTable{T: t1})
			continue
		}
		change, err := d.TableDiff(t1, t2)
		if err != nil {
			return nil, err
		}
		if len(change) > 0 {
			changes = append(changes, &schema.ModifyTable{
				T:       t2,
				Changes: change,
			})
		}
	}
	// Add tables.
	for _, t1 := range to.Tables {
		if _, ok := from.Table(t1.Name); !ok {
			changes = append(changes, &schema.AddTable{T: t1})
		}
	}
	return changes, nil
}

// TableDiff implements the schema.TableDiffer interface and returns a list of
// changes that need to be applied in order to move from one state to the other.
func (d *Diff) TableDiff(from, to *schema.Table) ([]schema.Change, error) {
	if from.Name != to.Name {
		return nil, fmt.Errorf("mismatched table names: %q != %q", from.Name, to.Name)
	}
	// Normalizing tables before starting the diff process.
	if n, ok := d.DiffDriver.(Normalizer); ok {
		n.Normalize(from, to)
	}
	var changes []schema.Change
	if from.Name != to.Name {
		return nil, fmt.Errorf("mismatched table names: %q != %q", from.Name, to.Name)
	}
	// PK modification is not supported.
	if pk1, pk2 := from.PrimaryKey, to.PrimaryKey; (pk1 != nil) != (pk2 != nil) || (pk1 != nil) && d.pkChange(pk1, pk2) != schema.NoChange {
		return nil, fmt.Errorf("changing %q table primary key is not supported", to.Name)
	}

	// Drop or modify attributes (collations, checks, etc).
	change, err := d.TableAttrDiff(from, to)
	if err != nil {
		return nil, err
	}
	changes = append(changes, change...)

	// Drop or modify columns.
	for _, c1 := range from.Columns {
		c2, ok := to.Column(c1.Name)
		if !ok {
			changes = append(changes, &schema.DropColumn{C: c1})
			continue
		}
		change, err := d.ColumnChange(c1, c2)
		if err != nil {
			return nil, err
		}
		if change != schema.NoChange {
			changes = append(changes, &schema.ModifyColumn{
				From:   c1,
				To:     c2,
				Change: change,
			})
		}
	}
	// Add columns.
	for _, c1 := range to.Columns {
		if _, ok := from.Column(c1.Name); !ok {
			changes = append(changes, &schema.AddColumn{C: c1})
		}
	}

	// Index changes.
	changes = append(changes, d.indexDiff(from, to)...)

	// Drop or modify foreign-keys.
	for _, fk1 := range from.ForeignKeys {
		fk2, ok := to.ForeignKey(fk1.Symbol)
		if !ok {
			changes = append(changes, &schema.DropForeignKey{F: fk1})
			continue
		}
		if change := d.fkChange(fk1, fk2); change != schema.NoChange {
			changes = append(changes, &schema.ModifyForeignKey{
				From:   fk1,
				To:     fk2,
				Change: change,
			})
		}
	}
	// Add foreign-keys.
	for _, fk1 := range to.ForeignKeys {
		if _, ok := from.ForeignKey(fk1.Symbol); !ok {
			changes = append(changes, &schema.AddForeignKey{F: fk1})
		}
	}
	return changes, nil
}

// indexDiff returns the schema changes (if any) for migrating table
// indexes from current state to the desired state.
func (d *Diff) indexDiff(from, to *schema.Table) []schema.Change {
	var (
		changes []schema.Change
		exists  = make(map[*schema.Index]bool)
	)
	// Drop or modify indexes.
	for _, idx1 := range from.Indexes {
		idx2, ok := to.Index(idx1.Name)
		// Found directly.
		if ok {
			if change := d.indexChange(idx1, idx2); change != schema.NoChange {
				changes = append(changes, &schema.ModifyIndex{
					From:   idx1,
					To:     idx2,
					Change: change,
				})
			}
			exists[idx2] = true
			continue
		}
		// Found indirectly.
		if d.IsGeneratedIndexName(from, idx1) {
			if idx2, ok := d.similarUnnamedIndex(to, idx1); ok {
				exists[idx2] = true
				continue
			}
		}
		// Not found.
		changes = append(changes, &schema.DropIndex{I: idx1})
	}
	// Add indexes.
	for _, idx := range to.Indexes {
		if exists[idx] {
			continue
		}
		if _, ok := from.Index(idx.Name); !ok {
			changes = append(changes, &schema.AddIndex{I: idx})
		}
	}
	return changes
}

// pkChange returns the schema changes (if any) for migrating one primary key to the other.
func (d *Diff) pkChange(from, to *schema.Index) schema.ChangeKind {
	change := d.indexChange(from, to)
	return change & ^schema.ChangeUnique
}

// indexChange returns the schema changes (if any) for migrating one index to the other.
func (d *Diff) indexChange(from, to *schema.Index) schema.ChangeKind {
	var change schema.ChangeKind
	if from.Unique != to.Unique {
		change |= schema.ChangeUnique
	}
	if d.IndexAttrChanged(from.Attrs, to.Attrs) {
		change |= schema.ChangeAttr
	}
	change |= d.partsChange(from.Parts, to.Parts)
	change |= CommentChange(from.Attrs, to.Attrs)
	return change
}

func (d *Diff) partsChange(from, to []*schema.IndexPart) schema.ChangeKind {
	if len(from) != len(to) {
		return schema.ChangeParts
	}
	sort.Slice(to, func(i, j int) bool { return to[i].SeqNo < to[j].SeqNo })
	sort.Slice(from, func(i, j int) bool { return from[i].SeqNo < from[j].SeqNo })
	for i := range from {
		switch {
		case from[i].Desc != to[i].Desc || d.IndexPartAttrChanged(from[i], to[i]):
			return schema.ChangeParts
		case from[i].C != nil && to[i].C != nil:
			if from[i].C.Name != to[i].C.Name {
				return schema.ChangeParts
			}
		case from[i].X != nil && to[i].X != nil:
			if from[i].X.(*schema.RawExpr).X != to[i].X.(*schema.RawExpr).X {
				return schema.ChangeParts
			}
		default: // (C1 != nil) != (C2 != nil) || (X1 != nil) != (X2 != nil).
			return schema.ChangeParts
		}
	}
	return schema.NoChange
}

// fkChange returns the schema changes (if any) for migrating one index to the other.
func (d *Diff) fkChange(from, to *schema.ForeignKey) schema.ChangeKind {
	var change schema.ChangeKind
	switch {
	case from.Table.Name != to.Table.Name:
		change |= schema.ChangeRefTable | schema.ChangeRefColumn
	case len(from.RefColumns) != len(to.RefColumns):
		change |= schema.ChangeRefColumn
	default:
		for i := range from.RefColumns {
			if from.RefColumns[i].Name != to.RefColumns[i].Name {
				change |= schema.ChangeRefColumn
			}
		}
	}
	switch {
	case len(from.Columns) != len(to.Columns):
		change |= schema.ChangeColumn
	default:
		for i := range from.Columns {
			if from.Columns[i].Name != to.Columns[i].Name {
				change |= schema.ChangeColumn
			}
		}
	}
	if d.ReferenceChanged(from.OnUpdate, to.OnUpdate) {
		change |= schema.ChangeUpdateAction
	}
	if d.ReferenceChanged(from.OnDelete, to.OnDelete) {
		change |= schema.ChangeDeleteAction
	}
	return change
}

// similarUnnamedIndex searches for an unnamed index with the same index-parts in the table.
func (d *Diff) similarUnnamedIndex(t *schema.Table, idx1 *schema.Index) (*schema.Index, bool) {
	for _, idx2 := range t.Indexes {
		if idx2.Name != "" || len(idx2.Parts) != len(idx1.Parts) || idx2.Unique != idx1.Unique {
			continue
		}
		if d.partsChange(idx1.Parts, idx2.Parts) == schema.NoChange {
			return idx2, true
		}
	}
	return nil, false
}

// CommentChange reports if the element comment was changed.
func CommentChange(from, to []schema.Attr) schema.ChangeKind {
	var c1, c2 schema.Comment
	if Has(from, &c1) != Has(to, &c2) || c1.Text != c2.Text {
		return schema.ChangeComment
	}
	return schema.NoChange
}

var (
	attrsType   = reflect.TypeOf(([]schema.Attr)(nil))
	clausesType = reflect.TypeOf(([]schema.Clause)(nil))
	exprsType   = reflect.TypeOf(([]schema.Expr)(nil))
)

// Has finds the first element in the elements list that
// matches target, and if so, sets target to that attribute
// value and returns true.
func Has(elements, target interface{}) bool {
	ev := reflect.ValueOf(elements)
	if t := ev.Type(); t != attrsType && t != clausesType && t != exprsType {
		panic(fmt.Sprintf("unexpected elements type: %T", elements))
	}
	tv := reflect.ValueOf(target)
	if tv.Kind() != reflect.Ptr || tv.IsNil() {
		panic("target must be a non-nil pointer")
	}
	for i := 0; i < ev.Len(); i++ {
		idx := ev.Index(i)
		if idx.IsNil() {
			continue
		}
		if e := idx.Elem(); e.Type().AssignableTo(tv.Type()) {
			tv.Elem().Set(e.Elem())
			return true
		}
	}
	return false
}

// UnsupportedTypeError describes an unsupported type error.
type UnsupportedTypeError struct {
	schema.Type
}

func (e UnsupportedTypeError) Error() string {
	return fmt.Sprintf("unsupported type %T", e.Type)
}

// CommentDiff computes the comment diff between the 2 attribute list.
// Note that, the implementation relies on the fact that both PostgreSQL
// and MySQL treat empty comment as "no comment" and a way to clear comments.
func CommentDiff(from, to []schema.Attr) schema.Change {
	var fromC, toC schema.Comment
	switch fromHas, toHas := Has(from, &fromC), Has(to, &toC); {
	case !fromHas && !toHas:
	case !fromHas && toC.Text != "":
		return &schema.AddAttr{
			A: &toC,
		}
	case !toHas:
		// In MySQL, there is no way to DROP a comment. Instead, setting it to empty ('')
		// will remove it from INFORMATION_SCHEMA. We use the same approach in PostgreSQL,
		// because comments can be dropped either by setting them to NULL or empty string.
		// See: postgres/backend/commands/comment.c#CreateComments.
		return &schema.ModifyAttr{
			From: &fromC,
			To:   &toC,
		}
	default:
		v1, err1 := Unquote(fromC.Text)
		v2, err2 := Unquote(toC.Text)
		if err1 == nil && err2 == nil && v1 != v2 {
			return &schema.ModifyAttr{
				From: &fromC,
				To:   &toC,
			}
		}
	}
	return nil
}

// CheckDiff computes the change diff between the 2 tables. A compare
// function is provided to check if a Check object was modified.
func CheckDiff(from, to *schema.Table, compare ...func(c1, c2 *schema.Check) bool) []schema.Change {
	var changes []schema.Change
	// Drop or modify checks.
	for _, c1 := range checks(from.Attrs) {
		switch c2, ok := similarCheck(to.Attrs, c1); {
		case !ok:
			changes = append(changes, &schema.DropCheck{
				C: c1,
			})
		case len(compare) == 1 && !compare[0](c1, c2):
			changes = append(changes, &schema.ModifyCheck{
				From: c1,
				To:   c2,
			})
		}
	}
	// Add checks.
	for _, c1 := range checks(to.Attrs) {
		if _, ok := similarCheck(from.Attrs, c1); !ok {
			changes = append(changes, &schema.AddCheck{
				C: c1,
			})
		}
	}
	return changes
}

// checks extracts all constraints from table attributes.
func checks(attr []schema.Attr) (checks []*schema.Check) {
	for i := range attr {
		if c, ok := attr[i].(*schema.Check); ok {
			checks = append(checks, c)
		}
	}
	return checks
}

// similarCheck returns a CHECK by its constraints name or expression.
func similarCheck(attrs []schema.Attr, c *schema.Check) (*schema.Check, bool) {
	var byName, byExpr *schema.Check
	for i := 0; i < len(attrs) && (byName == nil || byExpr == nil); i++ {
		check, ok := attrs[i].(*schema.Check)
		if !ok {
			continue
		}
		if check.Name != "" && check.Name == c.Name {
			byName = check
		}
		if check.Expr == c.Expr {
			byExpr = check
		}
	}
	// Give precedence to constraint name.
	if byName != nil {
		return byName, true
	}
	if byExpr != nil {
		return byExpr, true
	}
	return nil, false
}

// Unquote single or double quotes.
func Unquote(s string) (string, error) {
	switch {
	case IsQuoted(s, '"'):
		return strconv.Unquote(s)
	case IsQuoted(s, '\''):
		return strings.ReplaceAll(s[1:len(s)-1], "''", "'"), nil
	default:
		return s, nil
	}
}

// SingleQuote quotes the given string with single quote.
func SingleQuote(s string) (string, error) {
	switch {
	case IsQuoted(s, '\''):
		return s, nil
	case IsQuoted(s, '"'):
		v, err := strconv.Unquote(s)
		if err != nil {
			return "", err
		}
		s = v
		fallthrough
	default:
		return "'" + strings.ReplaceAll(s, "'", "''") + "'", nil
	}
}
