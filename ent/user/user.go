// Code generated by entc, DO NOT EDIT.

package user

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUID holds the string denoting the uid field in the database.
	FieldUID = "uid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldBirth holds the string denoting the birth field in the database.
	FieldBirth = "birth"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUID,
	FieldName,
	FieldBirth,
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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/wenerme/ent-demo/ent/runtime"
//
var (
	Hooks [1]ent.Hook
	// DefaultUID holds the default value on creation for the "uid" field.
	DefaultUID func() uuid.UUID
)
