// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wenerme/ent-demo/ent/pet"
	"github.com/wenerme/ent-demo/ent/user"
	"github.com/wenerme/ent-demo/models"
)

// PetCreate is the builder for creating a Pet entity.
type PetCreate struct {
	config
	mutation *PetMutation
	hooks    []Hook
}

// SetOwnerID sets the "ownerID" field.
func (pc *PetCreate) SetOwnerID(m *models.ID) *PetCreate {
	pc.mutation.SetOwnerID(m)
	return pc
}

// SetOwnerType sets the "ownerType" field.
func (pc *PetCreate) SetOwnerType(s string) *PetCreate {
	pc.mutation.SetOwnerType(s)
	return pc
}

// SetNillableOwnerType sets the "ownerType" field if the given value is not nil.
func (pc *PetCreate) SetNillableOwnerType(s *string) *PetCreate {
	if s != nil {
		pc.SetOwnerType(*s)
	}
	return pc
}

// SetOwningUserID sets the "owningUserID" field.
func (pc *PetCreate) SetOwningUserID(m *models.ID) *PetCreate {
	pc.mutation.SetOwningUserID(m)
	return pc
}

// SetName sets the "name" field.
func (pc *PetCreate) SetName(s string) *PetCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetID sets the "id" field.
func (pc *PetCreate) SetID(m models.ID) *PetCreate {
	pc.mutation.SetID(m)
	return pc
}

// SetOwningUser sets the "owningUser" edge to the User entity.
func (pc *PetCreate) SetOwningUser(u *User) *PetCreate {
	return pc.SetOwningUserID(u.ID)
}

// Mutation returns the PetMutation object of the builder.
func (pc *PetCreate) Mutation() *PetMutation {
	return pc.mutation
}

// Save creates the Pet in the database.
func (pc *PetCreate) Save(ctx context.Context) (*Pet, error) {
	var (
		err  error
		node *Pet
	)
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PetCreate) SaveX(ctx context.Context) *Pet {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (pc *PetCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	return nil
}

func (pc *PetCreate) sqlSave(ctx context.Context) (*Pet, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (pc *PetCreate) createSpec() (*Pet, *sqlgraph.CreateSpec) {
	var (
		_node = &Pet{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: pet.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: pet.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.OwnerID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pet.FieldOwnerID,
		})
		_node.OwnerID = value
	}
	if value, ok := pc.mutation.OwnerType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pet.FieldOwnerType,
		})
		_node.OwnerType = &value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pet.FieldName,
		})
		_node.Name = value
	}
	if nodes := pc.mutation.OwningUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   pet.OwningUserTable,
			Columns: []string{pet.OwningUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwningUserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PetCreateBulk is the builder for creating many Pet entities in bulk.
type PetCreateBulk struct {
	config
	builders []*PetCreate
}

// Save creates the Pet entities in the database.
func (pcb *PetCreateBulk) Save(ctx context.Context) ([]*Pet, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Pet, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PetMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PetCreateBulk) SaveX(ctx context.Context) []*Pet {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
