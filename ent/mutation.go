// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/wenerme/ent-demo/ent/pet"
	"github.com/wenerme/ent-demo/ent/predicate"
	"github.com/wenerme/ent-demo/ent/user"
	"github.com/wenerme/ent-demo/models"
	"github.com/xtgo/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypePet  = "Pet"
	TypeUser = "User"
)

// PetMutation represents an operation that mutates the Pet nodes in the graph.
type PetMutation struct {
	config
	op                Op
	typ               string
	id                *models.ID
	uid               **uuid.UUID
	ownerID           *models.ID
	ownerType         *string
	name              *string
	clearedFields     map[string]struct{}
	owningUser        *models.ID
	clearedowningUser bool
	done              bool
	oldValue          func(context.Context) (*Pet, error)
	predicates        []predicate.Pet
}

var _ ent.Mutation = (*PetMutation)(nil)

// petOption allows management of the mutation configuration using functional options.
type petOption func(*PetMutation)

// newPetMutation creates new mutation for the Pet entity.
func newPetMutation(c config, op Op, opts ...petOption) *PetMutation {
	m := &PetMutation{
		config:        c,
		op:            op,
		typ:           TypePet,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPetID sets the ID field of the mutation.
func withPetID(id models.ID) petOption {
	return func(m *PetMutation) {
		var (
			err   error
			once  sync.Once
			value *Pet
		)
		m.oldValue = func(ctx context.Context) (*Pet, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Pet.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPet sets the old Pet of the mutation.
func withPet(node *Pet) petOption {
	return func(m *PetMutation) {
		m.oldValue = func(context.Context) (*Pet, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PetMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PetMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Pet entities.
func (m *PetMutation) SetID(id models.ID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *PetMutation) ID() (id models.ID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetUID sets the "uid" field.
func (m *PetMutation) SetUID(u *uuid.UUID) {
	m.uid = &u
}

// UID returns the value of the "uid" field in the mutation.
func (m *PetMutation) UID() (r *uuid.UUID, exists bool) {
	v := m.uid
	if v == nil {
		return
	}
	return *v, true
}

// OldUID returns the old "uid" field's value of the Pet entity.
// If the Pet object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PetMutation) OldUID(ctx context.Context) (v *uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUID: %w", err)
	}
	return oldValue.UID, nil
}

// ClearUID clears the value of the "uid" field.
func (m *PetMutation) ClearUID() {
	m.uid = nil
	m.clearedFields[pet.FieldUID] = struct{}{}
}

// UIDCleared returns if the "uid" field was cleared in this mutation.
func (m *PetMutation) UIDCleared() bool {
	_, ok := m.clearedFields[pet.FieldUID]
	return ok
}

// ResetUID resets all changes to the "uid" field.
func (m *PetMutation) ResetUID() {
	m.uid = nil
	delete(m.clearedFields, pet.FieldUID)
}

// SetOwnerID sets the "ownerID" field.
func (m *PetMutation) SetOwnerID(value models.ID) {
	m.ownerID = &value
}

// OwnerID returns the value of the "ownerID" field in the mutation.
func (m *PetMutation) OwnerID() (r models.ID, exists bool) {
	v := m.ownerID
	if v == nil {
		return
	}
	return *v, true
}

// OldOwnerID returns the old "ownerID" field's value of the Pet entity.
// If the Pet object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PetMutation) OldOwnerID(ctx context.Context) (v *models.ID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldOwnerID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldOwnerID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOwnerID: %w", err)
	}
	return oldValue.OwnerID, nil
}

// ClearOwnerID clears the value of the "ownerID" field.
func (m *PetMutation) ClearOwnerID() {
	m.ownerID = nil
	m.clearedFields[pet.FieldOwnerID] = struct{}{}
}

// OwnerIDCleared returns if the "ownerID" field was cleared in this mutation.
func (m *PetMutation) OwnerIDCleared() bool {
	_, ok := m.clearedFields[pet.FieldOwnerID]
	return ok
}

// ResetOwnerID resets all changes to the "ownerID" field.
func (m *PetMutation) ResetOwnerID() {
	m.ownerID = nil
	delete(m.clearedFields, pet.FieldOwnerID)
}

// SetOwnerType sets the "ownerType" field.
func (m *PetMutation) SetOwnerType(s string) {
	m.ownerType = &s
}

// OwnerType returns the value of the "ownerType" field in the mutation.
func (m *PetMutation) OwnerType() (r string, exists bool) {
	v := m.ownerType
	if v == nil {
		return
	}
	return *v, true
}

// OldOwnerType returns the old "ownerType" field's value of the Pet entity.
// If the Pet object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PetMutation) OldOwnerType(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldOwnerType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldOwnerType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOwnerType: %w", err)
	}
	return oldValue.OwnerType, nil
}

// ClearOwnerType clears the value of the "ownerType" field.
func (m *PetMutation) ClearOwnerType() {
	m.ownerType = nil
	m.clearedFields[pet.FieldOwnerType] = struct{}{}
}

// OwnerTypeCleared returns if the "ownerType" field was cleared in this mutation.
func (m *PetMutation) OwnerTypeCleared() bool {
	_, ok := m.clearedFields[pet.FieldOwnerType]
	return ok
}

// ResetOwnerType resets all changes to the "ownerType" field.
func (m *PetMutation) ResetOwnerType() {
	m.ownerType = nil
	delete(m.clearedFields, pet.FieldOwnerType)
}

// SetOwningUserID sets the "owningUserID" field.
func (m *PetMutation) SetOwningUserID(value models.ID) {
	m.owningUser = &value
}

// OwningUserID returns the value of the "owningUserID" field in the mutation.
func (m *PetMutation) OwningUserID() (r models.ID, exists bool) {
	v := m.owningUser
	if v == nil {
		return
	}
	return *v, true
}

// OldOwningUserID returns the old "owningUserID" field's value of the Pet entity.
// If the Pet object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PetMutation) OldOwningUserID(ctx context.Context) (v *models.ID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldOwningUserID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldOwningUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOwningUserID: %w", err)
	}
	return oldValue.OwningUserID, nil
}

// ClearOwningUserID clears the value of the "owningUserID" field.
func (m *PetMutation) ClearOwningUserID() {
	m.owningUser = nil
	m.clearedFields[pet.FieldOwningUserID] = struct{}{}
}

// OwningUserIDCleared returns if the "owningUserID" field was cleared in this mutation.
func (m *PetMutation) OwningUserIDCleared() bool {
	_, ok := m.clearedFields[pet.FieldOwningUserID]
	return ok
}

// ResetOwningUserID resets all changes to the "owningUserID" field.
func (m *PetMutation) ResetOwningUserID() {
	m.owningUser = nil
	delete(m.clearedFields, pet.FieldOwningUserID)
}

// SetName sets the "name" field.
func (m *PetMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *PetMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Pet entity.
// If the Pet object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PetMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *PetMutation) ResetName() {
	m.name = nil
}

// ClearOwningUser clears the "owningUser" edge to the User entity.
func (m *PetMutation) ClearOwningUser() {
	m.clearedowningUser = true
}

// OwningUserCleared reports if the "owningUser" edge to the User entity was cleared.
func (m *PetMutation) OwningUserCleared() bool {
	return m.OwningUserIDCleared() || m.clearedowningUser
}

// OwningUserIDs returns the "owningUser" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// OwningUserID instead. It exists only for internal usage by the builders.
func (m *PetMutation) OwningUserIDs() (ids []models.ID) {
	if id := m.owningUser; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetOwningUser resets all changes to the "owningUser" edge.
func (m *PetMutation) ResetOwningUser() {
	m.owningUser = nil
	m.clearedowningUser = false
}

// Op returns the operation name.
func (m *PetMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Pet).
func (m *PetMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PetMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.uid != nil {
		fields = append(fields, pet.FieldUID)
	}
	if m.ownerID != nil {
		fields = append(fields, pet.FieldOwnerID)
	}
	if m.ownerType != nil {
		fields = append(fields, pet.FieldOwnerType)
	}
	if m.owningUser != nil {
		fields = append(fields, pet.FieldOwningUserID)
	}
	if m.name != nil {
		fields = append(fields, pet.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PetMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case pet.FieldUID:
		return m.UID()
	case pet.FieldOwnerID:
		return m.OwnerID()
	case pet.FieldOwnerType:
		return m.OwnerType()
	case pet.FieldOwningUserID:
		return m.OwningUserID()
	case pet.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PetMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case pet.FieldUID:
		return m.OldUID(ctx)
	case pet.FieldOwnerID:
		return m.OldOwnerID(ctx)
	case pet.FieldOwnerType:
		return m.OldOwnerType(ctx)
	case pet.FieldOwningUserID:
		return m.OldOwningUserID(ctx)
	case pet.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Pet field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PetMutation) SetField(name string, value ent.Value) error {
	switch name {
	case pet.FieldUID:
		v, ok := value.(*uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUID(v)
		return nil
	case pet.FieldOwnerID:
		v, ok := value.(models.ID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOwnerID(v)
		return nil
	case pet.FieldOwnerType:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOwnerType(v)
		return nil
	case pet.FieldOwningUserID:
		v, ok := value.(models.ID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOwningUserID(v)
		return nil
	case pet.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Pet field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PetMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PetMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PetMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Pet numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PetMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(pet.FieldUID) {
		fields = append(fields, pet.FieldUID)
	}
	if m.FieldCleared(pet.FieldOwnerID) {
		fields = append(fields, pet.FieldOwnerID)
	}
	if m.FieldCleared(pet.FieldOwnerType) {
		fields = append(fields, pet.FieldOwnerType)
	}
	if m.FieldCleared(pet.FieldOwningUserID) {
		fields = append(fields, pet.FieldOwningUserID)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PetMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *PetMutation) ClearField(name string) error {
	switch name {
	case pet.FieldUID:
		m.ClearUID()
		return nil
	case pet.FieldOwnerID:
		m.ClearOwnerID()
		return nil
	case pet.FieldOwnerType:
		m.ClearOwnerType()
		return nil
	case pet.FieldOwningUserID:
		m.ClearOwningUserID()
		return nil
	}
	return fmt.Errorf("unknown Pet nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PetMutation) ResetField(name string) error {
	switch name {
	case pet.FieldUID:
		m.ResetUID()
		return nil
	case pet.FieldOwnerID:
		m.ResetOwnerID()
		return nil
	case pet.FieldOwnerType:
		m.ResetOwnerType()
		return nil
	case pet.FieldOwningUserID:
		m.ResetOwningUserID()
		return nil
	case pet.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Pet field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PetMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.owningUser != nil {
		edges = append(edges, pet.EdgeOwningUser)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PetMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case pet.EdgeOwningUser:
		if id := m.owningUser; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PetMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PetMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PetMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedowningUser {
		edges = append(edges, pet.EdgeOwningUser)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PetMutation) EdgeCleared(name string) bool {
	switch name {
	case pet.EdgeOwningUser:
		return m.clearedowningUser
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PetMutation) ClearEdge(name string) error {
	switch name {
	case pet.EdgeOwningUser:
		m.ClearOwningUser()
		return nil
	}
	return fmt.Errorf("unknown Pet unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PetMutation) ResetEdge(name string) error {
	switch name {
	case pet.EdgeOwningUser:
		m.ResetOwningUser()
		return nil
	}
	return fmt.Errorf("unknown Pet edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *models.ID
	uid           **uuid.UUID
	name          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id models.ID) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id models.ID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *UserMutation) ID() (id models.ID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetUID sets the "uid" field.
func (m *UserMutation) SetUID(u *uuid.UUID) {
	m.uid = &u
}

// UID returns the value of the "uid" field in the mutation.
func (m *UserMutation) UID() (r *uuid.UUID, exists bool) {
	v := m.uid
	if v == nil {
		return
	}
	return *v, true
}

// OldUID returns the old "uid" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUID(ctx context.Context) (v *uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUID: %w", err)
	}
	return oldValue.UID, nil
}

// ClearUID clears the value of the "uid" field.
func (m *UserMutation) ClearUID() {
	m.uid = nil
	m.clearedFields[user.FieldUID] = struct{}{}
}

// UIDCleared returns if the "uid" field was cleared in this mutation.
func (m *UserMutation) UIDCleared() bool {
	_, ok := m.clearedFields[user.FieldUID]
	return ok
}

// ResetUID resets all changes to the "uid" field.
func (m *UserMutation) ResetUID() {
	m.uid = nil
	delete(m.clearedFields, user.FieldUID)
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.uid != nil {
		fields = append(fields, user.FieldUID)
	}
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldUID:
		return m.UID()
	case user.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldUID:
		return m.OldUID(ctx)
	case user.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldUID:
		v, ok := value.(*uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUID(v)
		return nil
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldUID) {
		fields = append(fields, user.FieldUID)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldUID:
		m.ClearUID()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldUID:
		m.ResetUID()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
