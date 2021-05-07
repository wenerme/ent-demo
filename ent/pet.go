// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/wenerme/ent-demo/ent/pet"
	"github.com/wenerme/ent-demo/ent/user"
	"github.com/wenerme/ent-demo/models"
)

// Pet is the model entity for the Pet schema.
type Pet struct {
	config `json:"-"`
	// ID of the ent.
	ID models.ID `json:"id,omitempty"`
	// UID holds the value of the "uid" field.
	UID *uuid.UUID `json:"uid,omitempty"`
	// OwnerID holds the value of the "ownerID" field.
	OwnerID *models.ID `json:"ownerID,omitempty"`
	// OwnerType holds the value of the "ownerType" field.
	OwnerType *string `json:"ownerType,omitempty"`
	// OwningUserID holds the value of the "owningUserID" field.
	OwningUserID *models.ID `json:"owningUserID,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PetQuery when eager-loading is set.
	Edges PetEdges `json:"edges"`
}

// PetEdges holds the relations/edges for other nodes in the graph.
type PetEdges struct {
	// OwningUser holds the value of the owningUser edge.
	OwningUser *User `json:"owningUser,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwningUserOrErr returns the OwningUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PetEdges) OwningUserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.OwningUser == nil {
			// The edge owningUser was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.OwningUser, nil
	}
	return nil, &NotLoadedError{edge: "owningUser"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pet) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case pet.FieldOwnerID, pet.FieldOwningUserID:
			values[i] = &sql.NullScanner{S: new(models.ID)}
		case pet.FieldUID:
			values[i] = new(*uuid.UUID)
		case pet.FieldID:
			values[i] = new(models.ID)
		case pet.FieldOwnerType, pet.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Pet", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pet fields.
func (pe *Pet) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pet.FieldID:
			if value, ok := values[i].(*models.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pe.ID = *value
			}
		case pet.FieldUID:
			if value, ok := values[i].(**uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uid", values[i])
			} else if value != nil {
				pe.UID = *value
			}
		case pet.FieldOwnerID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field ownerID", values[i])
			} else if value.Valid {
				pe.OwnerID = new(models.ID)
				*pe.OwnerID = *value.S.(*models.ID)
			}
		case pet.FieldOwnerType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ownerType", values[i])
			} else if value.Valid {
				pe.OwnerType = new(string)
				*pe.OwnerType = value.String
			}
		case pet.FieldOwningUserID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field owningUserID", values[i])
			} else if value.Valid {
				pe.OwningUserID = new(models.ID)
				*pe.OwningUserID = *value.S.(*models.ID)
			}
		case pet.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pe.Name = value.String
			}
		}
	}
	return nil
}

// QueryOwningUser queries the "owningUser" edge of the Pet entity.
func (pe *Pet) QueryOwningUser() *UserQuery {
	return (&PetClient{config: pe.config}).QueryOwningUser(pe)
}

// Update returns a builder for updating this Pet.
// Note that you need to call Pet.Unwrap() before calling this method if this Pet
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Pet) Update() *PetUpdateOne {
	return (&PetClient{config: pe.config}).UpdateOne(pe)
}

// Unwrap unwraps the Pet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Pet) Unwrap() *Pet {
	tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pet is not a transactional entity")
	}
	pe.config.driver = tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Pet) String() string {
	var builder strings.Builder
	builder.WriteString("Pet(")
	builder.WriteString(fmt.Sprintf("id=%v", pe.ID))
	builder.WriteString(", uid=")
	builder.WriteString(fmt.Sprintf("%v", pe.UID))
	if v := pe.OwnerID; v != nil {
		builder.WriteString(", ownerID=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := pe.OwnerType; v != nil {
		builder.WriteString(", ownerType=")
		builder.WriteString(*v)
	}
	if v := pe.OwningUserID; v != nil {
		builder.WriteString(", owningUserID=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", name=")
	builder.WriteString(pe.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Pets is a parsable slice of Pet.
type Pets []*Pet

func (pe Pets) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}
