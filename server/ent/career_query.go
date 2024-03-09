// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"metrograma/ent/career"
	"metrograma/ent/predicate"
	"metrograma/ent/subject"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CareerQuery is the builder for querying Career entities.
type CareerQuery struct {
	config
	ctx          *QueryContext
	order        []career.OrderOption
	inters       []Interceptor
	predicates   []predicate.Career
	withSubjects *SubjectQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CareerQuery builder.
func (cq *CareerQuery) Where(ps ...predicate.Career) *CareerQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CareerQuery) Limit(limit int) *CareerQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CareerQuery) Offset(offset int) *CareerQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CareerQuery) Unique(unique bool) *CareerQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CareerQuery) Order(o ...career.OrderOption) *CareerQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QuerySubjects chains the current query on the "subjects" edge.
func (cq *CareerQuery) QuerySubjects() *SubjectQuery {
	query := (&SubjectClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(career.Table, career.FieldID, selector),
			sqlgraph.To(subject.Table, subject.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, career.SubjectsTable, career.SubjectsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Career entity from the query.
// Returns a *NotFoundError when no Career was found.
func (cq *CareerQuery) First(ctx context.Context) (*Career, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{career.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CareerQuery) FirstX(ctx context.Context) *Career {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Career ID from the query.
// Returns a *NotFoundError when no Career ID was found.
func (cq *CareerQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{career.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CareerQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Career entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Career entity is found.
// Returns a *NotFoundError when no Career entities are found.
func (cq *CareerQuery) Only(ctx context.Context) (*Career, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{career.Label}
	default:
		return nil, &NotSingularError{career.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CareerQuery) OnlyX(ctx context.Context) *Career {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Career ID in the query.
// Returns a *NotSingularError when more than one Career ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CareerQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{career.Label}
	default:
		err = &NotSingularError{career.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CareerQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Careers.
func (cq *CareerQuery) All(ctx context.Context) ([]*Career, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Career, *CareerQuery]()
	return withInterceptors[[]*Career](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CareerQuery) AllX(ctx context.Context) []*Career {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Career IDs.
func (cq *CareerQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(career.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CareerQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CareerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CareerQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CareerQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CareerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CareerQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CareerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CareerQuery) Clone() *CareerQuery {
	if cq == nil {
		return nil
	}
	return &CareerQuery{
		config:       cq.config,
		ctx:          cq.ctx.Clone(),
		order:        append([]career.OrderOption{}, cq.order...),
		inters:       append([]Interceptor{}, cq.inters...),
		predicates:   append([]predicate.Career{}, cq.predicates...),
		withSubjects: cq.withSubjects.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithSubjects tells the query-builder to eager-load the nodes that are connected to
// the "subjects" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CareerQuery) WithSubjects(opts ...func(*SubjectQuery)) *CareerQuery {
	query := (&SubjectClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withSubjects = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Career.Query().
//		GroupBy(career.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CareerQuery) GroupBy(field string, fields ...string) *CareerGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CareerGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = career.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Career.Query().
//		Select(career.FieldName).
//		Scan(ctx, &v)
func (cq *CareerQuery) Select(fields ...string) *CareerSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CareerSelect{CareerQuery: cq}
	sbuild.label = career.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CareerSelect configured with the given aggregations.
func (cq *CareerQuery) Aggregate(fns ...AggregateFunc) *CareerSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CareerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !career.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CareerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Career, error) {
	var (
		nodes       = []*Career{}
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withSubjects != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Career).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Career{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withSubjects; query != nil {
		if err := cq.loadSubjects(ctx, query, nodes,
			func(n *Career) { n.Edges.Subjects = []*Subject{} },
			func(n *Career, e *Subject) { n.Edges.Subjects = append(n.Edges.Subjects, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CareerQuery) loadSubjects(ctx context.Context, query *SubjectQuery, nodes []*Career, init func(*Career), assign func(*Career, *Subject)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Career)
	nids := make(map[uuid.UUID]map[*Career]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(career.SubjectsTable)
		s.Join(joinT).On(s.C(subject.FieldID), joinT.C(career.SubjectsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(career.SubjectsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(career.SubjectsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Career]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Subject](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "subjects" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (cq *CareerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CareerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(career.Table, career.Columns, sqlgraph.NewFieldSpec(career.FieldID, field.TypeUUID))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, career.FieldID)
		for i := range fields {
			if fields[i] != career.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CareerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(career.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = career.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CareerGroupBy is the group-by builder for Career entities.
type CareerGroupBy struct {
	selector
	build *CareerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CareerGroupBy) Aggregate(fns ...AggregateFunc) *CareerGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CareerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CareerQuery, *CareerGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CareerGroupBy) sqlScan(ctx context.Context, root *CareerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CareerSelect is the builder for selecting fields of Career entities.
type CareerSelect struct {
	*CareerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CareerSelect) Aggregate(fns ...AggregateFunc) *CareerSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CareerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CareerQuery, *CareerSelect](ctx, cs.CareerQuery, cs, cs.inters, v)
}

func (cs *CareerSelect) sqlScan(ctx context.Context, root *CareerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}