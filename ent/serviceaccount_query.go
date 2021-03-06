// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wenerme/ent-demo/ent/predicate"
	"github.com/wenerme/ent-demo/ent/serviceaccount"
)

// ServiceAccountQuery is the builder for querying ServiceAccount entities.
type ServiceAccountQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ServiceAccount
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ServiceAccountQuery builder.
func (saq *ServiceAccountQuery) Where(ps ...predicate.ServiceAccount) *ServiceAccountQuery {
	saq.predicates = append(saq.predicates, ps...)
	return saq
}

// Limit adds a limit step to the query.
func (saq *ServiceAccountQuery) Limit(limit int) *ServiceAccountQuery {
	saq.limit = &limit
	return saq
}

// Offset adds an offset step to the query.
func (saq *ServiceAccountQuery) Offset(offset int) *ServiceAccountQuery {
	saq.offset = &offset
	return saq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (saq *ServiceAccountQuery) Unique(unique bool) *ServiceAccountQuery {
	saq.unique = &unique
	return saq
}

// Order adds an order step to the query.
func (saq *ServiceAccountQuery) Order(o ...OrderFunc) *ServiceAccountQuery {
	saq.order = append(saq.order, o...)
	return saq
}

// First returns the first ServiceAccount entity from the query.
// Returns a *NotFoundError when no ServiceAccount was found.
func (saq *ServiceAccountQuery) First(ctx context.Context) (*ServiceAccount, error) {
	nodes, err := saq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{serviceaccount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (saq *ServiceAccountQuery) FirstX(ctx context.Context) *ServiceAccount {
	node, err := saq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ServiceAccount ID from the query.
// Returns a *NotFoundError when no ServiceAccount ID was found.
func (saq *ServiceAccountQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = saq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{serviceaccount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (saq *ServiceAccountQuery) FirstIDX(ctx context.Context) string {
	id, err := saq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ServiceAccount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ServiceAccount entity is found.
// Returns a *NotFoundError when no ServiceAccount entities are found.
func (saq *ServiceAccountQuery) Only(ctx context.Context) (*ServiceAccount, error) {
	nodes, err := saq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{serviceaccount.Label}
	default:
		return nil, &NotSingularError{serviceaccount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (saq *ServiceAccountQuery) OnlyX(ctx context.Context) *ServiceAccount {
	node, err := saq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ServiceAccount ID in the query.
// Returns a *NotSingularError when more than one ServiceAccount ID is found.
// Returns a *NotFoundError when no entities are found.
func (saq *ServiceAccountQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = saq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{serviceaccount.Label}
	default:
		err = &NotSingularError{serviceaccount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (saq *ServiceAccountQuery) OnlyIDX(ctx context.Context) string {
	id, err := saq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ServiceAccounts.
func (saq *ServiceAccountQuery) All(ctx context.Context) ([]*ServiceAccount, error) {
	if err := saq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return saq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (saq *ServiceAccountQuery) AllX(ctx context.Context) []*ServiceAccount {
	nodes, err := saq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ServiceAccount IDs.
func (saq *ServiceAccountQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := saq.Select(serviceaccount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (saq *ServiceAccountQuery) IDsX(ctx context.Context) []string {
	ids, err := saq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (saq *ServiceAccountQuery) Count(ctx context.Context) (int, error) {
	if err := saq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return saq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (saq *ServiceAccountQuery) CountX(ctx context.Context) int {
	count, err := saq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (saq *ServiceAccountQuery) Exist(ctx context.Context) (bool, error) {
	if err := saq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return saq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (saq *ServiceAccountQuery) ExistX(ctx context.Context) bool {
	exist, err := saq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ServiceAccountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (saq *ServiceAccountQuery) Clone() *ServiceAccountQuery {
	if saq == nil {
		return nil
	}
	return &ServiceAccountQuery{
		config:     saq.config,
		limit:      saq.limit,
		offset:     saq.offset,
		order:      append([]OrderFunc{}, saq.order...),
		predicates: append([]predicate.ServiceAccount{}, saq.predicates...),
		// clone intermediate query.
		sql:    saq.sql.Clone(),
		path:   saq.path,
		unique: saq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Sid int `json:"sid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ServiceAccount.Query().
//		GroupBy(serviceaccount.FieldSid).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (saq *ServiceAccountQuery) GroupBy(field string, fields ...string) *ServiceAccountGroupBy {
	grbuild := &ServiceAccountGroupBy{config: saq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := saq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return saq.sqlQuery(ctx), nil
	}
	grbuild.label = serviceaccount.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Sid int `json:"sid,omitempty"`
//	}
//
//	client.ServiceAccount.Query().
//		Select(serviceaccount.FieldSid).
//		Scan(ctx, &v)
//
func (saq *ServiceAccountQuery) Select(fields ...string) *ServiceAccountSelect {
	saq.fields = append(saq.fields, fields...)
	selbuild := &ServiceAccountSelect{ServiceAccountQuery: saq}
	selbuild.label = serviceaccount.Label
	selbuild.flds, selbuild.scan = &saq.fields, selbuild.Scan
	return selbuild
}

func (saq *ServiceAccountQuery) prepareQuery(ctx context.Context) error {
	for _, f := range saq.fields {
		if !serviceaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if saq.path != nil {
		prev, err := saq.path(ctx)
		if err != nil {
			return err
		}
		saq.sql = prev
	}
	return nil
}

func (saq *ServiceAccountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ServiceAccount, error) {
	var (
		nodes = []*ServiceAccount{}
		_spec = saq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*ServiceAccount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &ServiceAccount{config: saq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, saq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (saq *ServiceAccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := saq.querySpec()
	_spec.Node.Columns = saq.fields
	if len(saq.fields) > 0 {
		_spec.Unique = saq.unique != nil && *saq.unique
	}
	return sqlgraph.CountNodes(ctx, saq.driver, _spec)
}

func (saq *ServiceAccountQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := saq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (saq *ServiceAccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   serviceaccount.Table,
			Columns: serviceaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: serviceaccount.FieldID,
			},
		},
		From:   saq.sql,
		Unique: true,
	}
	if unique := saq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := saq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, serviceaccount.FieldID)
		for i := range fields {
			if fields[i] != serviceaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := saq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := saq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := saq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := saq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (saq *ServiceAccountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(saq.driver.Dialect())
	t1 := builder.Table(serviceaccount.Table)
	columns := saq.fields
	if len(columns) == 0 {
		columns = serviceaccount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if saq.sql != nil {
		selector = saq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if saq.unique != nil && *saq.unique {
		selector.Distinct()
	}
	for _, p := range saq.predicates {
		p(selector)
	}
	for _, p := range saq.order {
		p(selector)
	}
	if offset := saq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := saq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ServiceAccountGroupBy is the group-by builder for ServiceAccount entities.
type ServiceAccountGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sagb *ServiceAccountGroupBy) Aggregate(fns ...AggregateFunc) *ServiceAccountGroupBy {
	sagb.fns = append(sagb.fns, fns...)
	return sagb
}

// Scan applies the group-by query and scans the result into the given value.
func (sagb *ServiceAccountGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sagb.path(ctx)
	if err != nil {
		return err
	}
	sagb.sql = query
	return sagb.sqlScan(ctx, v)
}

func (sagb *ServiceAccountGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sagb.fields {
		if !serviceaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sagb *ServiceAccountGroupBy) sqlQuery() *sql.Selector {
	selector := sagb.sql.Select()
	aggregation := make([]string, 0, len(sagb.fns))
	for _, fn := range sagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sagb.fields)+len(sagb.fns))
		for _, f := range sagb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sagb.fields...)...)
}

// ServiceAccountSelect is the builder for selecting fields of ServiceAccount entities.
type ServiceAccountSelect struct {
	*ServiceAccountQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sas *ServiceAccountSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sas.prepareQuery(ctx); err != nil {
		return err
	}
	sas.sql = sas.ServiceAccountQuery.sqlQuery(ctx)
	return sas.sqlScan(ctx, v)
}

func (sas *ServiceAccountSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sas.sql.Query()
	if err := sas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
