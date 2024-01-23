// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"realm.pub/tavern/internal/ent/host"
	"realm.pub/tavern/internal/ent/hostfile"
	"realm.pub/tavern/internal/ent/predicate"
	"realm.pub/tavern/internal/ent/task"
)

// HostFileQuery is the builder for querying HostFile entities.
type HostFileQuery struct {
	config
	ctx        *QueryContext
	order      []hostfile.OrderOption
	inters     []Interceptor
	predicates []predicate.HostFile
	withHost   *HostQuery
	withTask   *TaskQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*HostFile) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HostFileQuery builder.
func (hfq *HostFileQuery) Where(ps ...predicate.HostFile) *HostFileQuery {
	hfq.predicates = append(hfq.predicates, ps...)
	return hfq
}

// Limit the number of records to be returned by this query.
func (hfq *HostFileQuery) Limit(limit int) *HostFileQuery {
	hfq.ctx.Limit = &limit
	return hfq
}

// Offset to start from.
func (hfq *HostFileQuery) Offset(offset int) *HostFileQuery {
	hfq.ctx.Offset = &offset
	return hfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (hfq *HostFileQuery) Unique(unique bool) *HostFileQuery {
	hfq.ctx.Unique = &unique
	return hfq
}

// Order specifies how the records should be ordered.
func (hfq *HostFileQuery) Order(o ...hostfile.OrderOption) *HostFileQuery {
	hfq.order = append(hfq.order, o...)
	return hfq
}

// QueryHost chains the current query on the "host" edge.
func (hfq *HostFileQuery) QueryHost() *HostQuery {
	query := (&HostClient{config: hfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(hostfile.Table, hostfile.FieldID, selector),
			sqlgraph.To(host.Table, host.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, hostfile.HostTable, hostfile.HostColumn),
		)
		fromU = sqlgraph.SetNeighbors(hfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTask chains the current query on the "task" edge.
func (hfq *HostFileQuery) QueryTask() *TaskQuery {
	query := (&TaskClient{config: hfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(hostfile.Table, hostfile.FieldID, selector),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, hostfile.TaskTable, hostfile.TaskColumn),
		)
		fromU = sqlgraph.SetNeighbors(hfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first HostFile entity from the query.
// Returns a *NotFoundError when no HostFile was found.
func (hfq *HostFileQuery) First(ctx context.Context) (*HostFile, error) {
	nodes, err := hfq.Limit(1).All(setContextOp(ctx, hfq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{hostfile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hfq *HostFileQuery) FirstX(ctx context.Context) *HostFile {
	node, err := hfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first HostFile ID from the query.
// Returns a *NotFoundError when no HostFile ID was found.
func (hfq *HostFileQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hfq.Limit(1).IDs(setContextOp(ctx, hfq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{hostfile.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (hfq *HostFileQuery) FirstIDX(ctx context.Context) int {
	id, err := hfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single HostFile entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one HostFile entity is found.
// Returns a *NotFoundError when no HostFile entities are found.
func (hfq *HostFileQuery) Only(ctx context.Context) (*HostFile, error) {
	nodes, err := hfq.Limit(2).All(setContextOp(ctx, hfq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{hostfile.Label}
	default:
		return nil, &NotSingularError{hostfile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hfq *HostFileQuery) OnlyX(ctx context.Context) *HostFile {
	node, err := hfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only HostFile ID in the query.
// Returns a *NotSingularError when more than one HostFile ID is found.
// Returns a *NotFoundError when no entities are found.
func (hfq *HostFileQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hfq.Limit(2).IDs(setContextOp(ctx, hfq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{hostfile.Label}
	default:
		err = &NotSingularError{hostfile.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hfq *HostFileQuery) OnlyIDX(ctx context.Context) int {
	id, err := hfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of HostFiles.
func (hfq *HostFileQuery) All(ctx context.Context) ([]*HostFile, error) {
	ctx = setContextOp(ctx, hfq.ctx, "All")
	if err := hfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*HostFile, *HostFileQuery]()
	return withInterceptors[[]*HostFile](ctx, hfq, qr, hfq.inters)
}

// AllX is like All, but panics if an error occurs.
func (hfq *HostFileQuery) AllX(ctx context.Context) []*HostFile {
	nodes, err := hfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of HostFile IDs.
func (hfq *HostFileQuery) IDs(ctx context.Context) (ids []int, err error) {
	if hfq.ctx.Unique == nil && hfq.path != nil {
		hfq.Unique(true)
	}
	ctx = setContextOp(ctx, hfq.ctx, "IDs")
	if err = hfq.Select(hostfile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hfq *HostFileQuery) IDsX(ctx context.Context) []int {
	ids, err := hfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hfq *HostFileQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, hfq.ctx, "Count")
	if err := hfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, hfq, querierCount[*HostFileQuery](), hfq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (hfq *HostFileQuery) CountX(ctx context.Context) int {
	count, err := hfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hfq *HostFileQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, hfq.ctx, "Exist")
	switch _, err := hfq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (hfq *HostFileQuery) ExistX(ctx context.Context) bool {
	exist, err := hfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HostFileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hfq *HostFileQuery) Clone() *HostFileQuery {
	if hfq == nil {
		return nil
	}
	return &HostFileQuery{
		config:     hfq.config,
		ctx:        hfq.ctx.Clone(),
		order:      append([]hostfile.OrderOption{}, hfq.order...),
		inters:     append([]Interceptor{}, hfq.inters...),
		predicates: append([]predicate.HostFile{}, hfq.predicates...),
		withHost:   hfq.withHost.Clone(),
		withTask:   hfq.withTask.Clone(),
		// clone intermediate query.
		sql:  hfq.sql.Clone(),
		path: hfq.path,
	}
}

// WithHost tells the query-builder to eager-load the nodes that are connected to
// the "host" edge. The optional arguments are used to configure the query builder of the edge.
func (hfq *HostFileQuery) WithHost(opts ...func(*HostQuery)) *HostFileQuery {
	query := (&HostClient{config: hfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hfq.withHost = query
	return hfq
}

// WithTask tells the query-builder to eager-load the nodes that are connected to
// the "task" edge. The optional arguments are used to configure the query builder of the edge.
func (hfq *HostFileQuery) WithTask(opts ...func(*TaskQuery)) *HostFileQuery {
	query := (&TaskClient{config: hfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hfq.withTask = query
	return hfq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.HostFile.Query().
//		GroupBy(hostfile.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (hfq *HostFileQuery) GroupBy(field string, fields ...string) *HostFileGroupBy {
	hfq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &HostFileGroupBy{build: hfq}
	grbuild.flds = &hfq.ctx.Fields
	grbuild.label = hostfile.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.HostFile.Query().
//		Select(hostfile.FieldCreatedAt).
//		Scan(ctx, &v)
func (hfq *HostFileQuery) Select(fields ...string) *HostFileSelect {
	hfq.ctx.Fields = append(hfq.ctx.Fields, fields...)
	sbuild := &HostFileSelect{HostFileQuery: hfq}
	sbuild.label = hostfile.Label
	sbuild.flds, sbuild.scan = &hfq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a HostFileSelect configured with the given aggregations.
func (hfq *HostFileQuery) Aggregate(fns ...AggregateFunc) *HostFileSelect {
	return hfq.Select().Aggregate(fns...)
}

func (hfq *HostFileQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range hfq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, hfq); err != nil {
				return err
			}
		}
	}
	for _, f := range hfq.ctx.Fields {
		if !hostfile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if hfq.path != nil {
		prev, err := hfq.path(ctx)
		if err != nil {
			return err
		}
		hfq.sql = prev
	}
	return nil
}

func (hfq *HostFileQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*HostFile, error) {
	var (
		nodes       = []*HostFile{}
		withFKs     = hfq.withFKs
		_spec       = hfq.querySpec()
		loadedTypes = [2]bool{
			hfq.withHost != nil,
			hfq.withTask != nil,
		}
	)
	if hfq.withHost != nil || hfq.withTask != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, hostfile.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*HostFile).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &HostFile{config: hfq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(hfq.modifiers) > 0 {
		_spec.Modifiers = hfq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, hfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := hfq.withHost; query != nil {
		if err := hfq.loadHost(ctx, query, nodes, nil,
			func(n *HostFile, e *Host) { n.Edges.Host = e }); err != nil {
			return nil, err
		}
	}
	if query := hfq.withTask; query != nil {
		if err := hfq.loadTask(ctx, query, nodes, nil,
			func(n *HostFile, e *Task) { n.Edges.Task = e }); err != nil {
			return nil, err
		}
	}
	for i := range hfq.loadTotal {
		if err := hfq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (hfq *HostFileQuery) loadHost(ctx context.Context, query *HostQuery, nodes []*HostFile, init func(*HostFile), assign func(*HostFile, *Host)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*HostFile)
	for i := range nodes {
		if nodes[i].host_file_host == nil {
			continue
		}
		fk := *nodes[i].host_file_host
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(host.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "host_file_host" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (hfq *HostFileQuery) loadTask(ctx context.Context, query *TaskQuery, nodes []*HostFile, init func(*HostFile), assign func(*HostFile, *Task)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*HostFile)
	for i := range nodes {
		if nodes[i].task_reported_files == nil {
			continue
		}
		fk := *nodes[i].task_reported_files
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(task.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "task_reported_files" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (hfq *HostFileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hfq.querySpec()
	if len(hfq.modifiers) > 0 {
		_spec.Modifiers = hfq.modifiers
	}
	_spec.Node.Columns = hfq.ctx.Fields
	if len(hfq.ctx.Fields) > 0 {
		_spec.Unique = hfq.ctx.Unique != nil && *hfq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, hfq.driver, _spec)
}

func (hfq *HostFileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(hostfile.Table, hostfile.Columns, sqlgraph.NewFieldSpec(hostfile.FieldID, field.TypeInt))
	_spec.From = hfq.sql
	if unique := hfq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if hfq.path != nil {
		_spec.Unique = true
	}
	if fields := hfq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hostfile.FieldID)
		for i := range fields {
			if fields[i] != hostfile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := hfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hfq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hfq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (hfq *HostFileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(hfq.driver.Dialect())
	t1 := builder.Table(hostfile.Table)
	columns := hfq.ctx.Fields
	if len(columns) == 0 {
		columns = hostfile.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if hfq.sql != nil {
		selector = hfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if hfq.ctx.Unique != nil && *hfq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range hfq.predicates {
		p(selector)
	}
	for _, p := range hfq.order {
		p(selector)
	}
	if offset := hfq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hfq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// HostFileGroupBy is the group-by builder for HostFile entities.
type HostFileGroupBy struct {
	selector
	build *HostFileQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hfgb *HostFileGroupBy) Aggregate(fns ...AggregateFunc) *HostFileGroupBy {
	hfgb.fns = append(hfgb.fns, fns...)
	return hfgb
}

// Scan applies the selector query and scans the result into the given value.
func (hfgb *HostFileGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hfgb.build.ctx, "GroupBy")
	if err := hfgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HostFileQuery, *HostFileGroupBy](ctx, hfgb.build, hfgb, hfgb.build.inters, v)
}

func (hfgb *HostFileGroupBy) sqlScan(ctx context.Context, root *HostFileQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(hfgb.fns))
	for _, fn := range hfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*hfgb.flds)+len(hfgb.fns))
		for _, f := range *hfgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*hfgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hfgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// HostFileSelect is the builder for selecting fields of HostFile entities.
type HostFileSelect struct {
	*HostFileQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (hfs *HostFileSelect) Aggregate(fns ...AggregateFunc) *HostFileSelect {
	hfs.fns = append(hfs.fns, fns...)
	return hfs
}

// Scan applies the selector query and scans the result into the given value.
func (hfs *HostFileSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hfs.ctx, "Select")
	if err := hfs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HostFileQuery, *HostFileSelect](ctx, hfs.HostFileQuery, hfs, hfs.inters, v)
}

func (hfs *HostFileSelect) sqlScan(ctx context.Context, root *HostFileQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(hfs.fns))
	for _, fn := range hfs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*hfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
