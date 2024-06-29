// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Endpoint is an object representing the database table.
type Endpoint struct {
	ID            string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	ApplicationID string      `boil:"application_id" json:"application_id" toml:"application_id" yaml:"application_id"`
	URL           string      `boil:"url" json:"url" toml:"url" yaml:"url"`
	Description   null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	SigningSecret string      `boil:"signing_secret" json:"signing_secret" toml:"signing_secret" yaml:"signing_secret"`
	CreatedAt     time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt     time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *endpointR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L endpointL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EndpointColumns = struct {
	ID            string
	ApplicationID string
	URL           string
	Description   string
	SigningSecret string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	ApplicationID: "application_id",
	URL:           "url",
	Description:   "description",
	SigningSecret: "signing_secret",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

var EndpointTableColumns = struct {
	ID            string
	ApplicationID string
	URL           string
	Description   string
	SigningSecret string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "endpoints.id",
	ApplicationID: "endpoints.application_id",
	URL:           "endpoints.url",
	Description:   "endpoints.description",
	SigningSecret: "endpoints.signing_secret",
	CreatedAt:     "endpoints.created_at",
	UpdatedAt:     "endpoints.updated_at",
}

// Generated where

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_String) LIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" LIKE ?", x)
}
func (w whereHelpernull_String) NLIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" NOT LIKE ?", x)
}
func (w whereHelpernull_String) ILIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" ILIKE ?", x)
}
func (w whereHelpernull_String) NILIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" NOT ILIKE ?", x)
}
func (w whereHelpernull_String) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_String) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var EndpointWhere = struct {
	ID            whereHelperstring
	ApplicationID whereHelperstring
	URL           whereHelperstring
	Description   whereHelpernull_String
	SigningSecret whereHelperstring
	CreatedAt     whereHelpertime_Time
	UpdatedAt     whereHelpertime_Time
}{
	ID:            whereHelperstring{field: "\"endpoints\".\"id\""},
	ApplicationID: whereHelperstring{field: "\"endpoints\".\"application_id\""},
	URL:           whereHelperstring{field: "\"endpoints\".\"url\""},
	Description:   whereHelpernull_String{field: "\"endpoints\".\"description\""},
	SigningSecret: whereHelperstring{field: "\"endpoints\".\"signing_secret\""},
	CreatedAt:     whereHelpertime_Time{field: "\"endpoints\".\"created_at\""},
	UpdatedAt:     whereHelpertime_Time{field: "\"endpoints\".\"updated_at\""},
}

// EndpointRels is where relationship names are stored.
var EndpointRels = struct {
	Application string
}{
	Application: "Application",
}

// endpointR is where relationships are stored.
type endpointR struct {
	Application *Application `boil:"Application" json:"Application" toml:"Application" yaml:"Application"`
}

// NewStruct creates a new relationship struct
func (*endpointR) NewStruct() *endpointR {
	return &endpointR{}
}

func (r *endpointR) GetApplication() *Application {
	if r == nil {
		return nil
	}
	return r.Application
}

// endpointL is where Load methods for each relationship are stored.
type endpointL struct{}

var (
	endpointAllColumns            = []string{"id", "application_id", "url", "description", "signing_secret", "created_at", "updated_at"}
	endpointColumnsWithoutDefault = []string{"id", "application_id", "url", "signing_secret"}
	endpointColumnsWithDefault    = []string{"description", "created_at", "updated_at"}
	endpointPrimaryKeyColumns     = []string{"id"}
	endpointGeneratedColumns      = []string{}
)

type (
	// EndpointSlice is an alias for a slice of pointers to Endpoint.
	// This should almost always be used instead of []Endpoint.
	EndpointSlice []*Endpoint
	// EndpointHook is the signature for custom Endpoint hook methods
	EndpointHook func(context.Context, boil.ContextExecutor, *Endpoint) error

	endpointQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	endpointType                 = reflect.TypeOf(&Endpoint{})
	endpointMapping              = queries.MakeStructMapping(endpointType)
	endpointPrimaryKeyMapping, _ = queries.BindMapping(endpointType, endpointMapping, endpointPrimaryKeyColumns)
	endpointInsertCacheMut       sync.RWMutex
	endpointInsertCache          = make(map[string]insertCache)
	endpointUpdateCacheMut       sync.RWMutex
	endpointUpdateCache          = make(map[string]updateCache)
	endpointUpsertCacheMut       sync.RWMutex
	endpointUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var endpointAfterSelectHooks []EndpointHook

var endpointBeforeInsertHooks []EndpointHook
var endpointAfterInsertHooks []EndpointHook

var endpointBeforeUpdateHooks []EndpointHook
var endpointAfterUpdateHooks []EndpointHook

var endpointBeforeDeleteHooks []EndpointHook
var endpointAfterDeleteHooks []EndpointHook

var endpointBeforeUpsertHooks []EndpointHook
var endpointAfterUpsertHooks []EndpointHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Endpoint) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range endpointAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Endpoint) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range endpointBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Endpoint) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range endpointAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Endpoint) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range endpointBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Endpoint) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range endpointAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Endpoint) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range endpointBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Endpoint) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range endpointAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Endpoint) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range endpointBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Endpoint) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range endpointAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEndpointHook registers your hook function for all future operations.
func AddEndpointHook(hookPoint boil.HookPoint, endpointHook EndpointHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		endpointAfterSelectHooks = append(endpointAfterSelectHooks, endpointHook)
	case boil.BeforeInsertHook:
		endpointBeforeInsertHooks = append(endpointBeforeInsertHooks, endpointHook)
	case boil.AfterInsertHook:
		endpointAfterInsertHooks = append(endpointAfterInsertHooks, endpointHook)
	case boil.BeforeUpdateHook:
		endpointBeforeUpdateHooks = append(endpointBeforeUpdateHooks, endpointHook)
	case boil.AfterUpdateHook:
		endpointAfterUpdateHooks = append(endpointAfterUpdateHooks, endpointHook)
	case boil.BeforeDeleteHook:
		endpointBeforeDeleteHooks = append(endpointBeforeDeleteHooks, endpointHook)
	case boil.AfterDeleteHook:
		endpointAfterDeleteHooks = append(endpointAfterDeleteHooks, endpointHook)
	case boil.BeforeUpsertHook:
		endpointBeforeUpsertHooks = append(endpointBeforeUpsertHooks, endpointHook)
	case boil.AfterUpsertHook:
		endpointAfterUpsertHooks = append(endpointAfterUpsertHooks, endpointHook)
	}
}

// One returns a single endpoint record from the query.
func (q endpointQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Endpoint, error) {
	o := &Endpoint{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for endpoints")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Endpoint records from the query.
func (q endpointQuery) All(ctx context.Context, exec boil.ContextExecutor) (EndpointSlice, error) {
	var o []*Endpoint

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Endpoint slice")
	}

	if len(endpointAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Endpoint records in the query.
func (q endpointQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count endpoints rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q endpointQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if endpoints exists")
	}

	return count > 0, nil
}

// Application pointed to by the foreign key.
func (o *Endpoint) Application(mods ...qm.QueryMod) applicationQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ApplicationID),
	}

	queryMods = append(queryMods, mods...)

	return Applications(queryMods...)
}

// LoadApplication allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (endpointL) LoadApplication(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEndpoint interface{}, mods queries.Applicator) error {
	var slice []*Endpoint
	var object *Endpoint

	if singular {
		var ok bool
		object, ok = maybeEndpoint.(*Endpoint)
		if !ok {
			object = new(Endpoint)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeEndpoint)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeEndpoint))
			}
		}
	} else {
		s, ok := maybeEndpoint.(*[]*Endpoint)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeEndpoint)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeEndpoint))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &endpointR{}
		}
		args = append(args, object.ApplicationID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &endpointR{}
			}

			for _, a := range args {
				if a == obj.ApplicationID {
					continue Outer
				}
			}

			args = append(args, obj.ApplicationID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`applications`),
		qm.WhereIn(`applications.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Application")
	}

	var resultSlice []*Application
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Application")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for applications")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for applications")
	}

	if len(applicationAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Application = foreign
		if foreign.R == nil {
			foreign.R = &applicationR{}
		}
		foreign.R.Endpoints = append(foreign.R.Endpoints, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ApplicationID == foreign.ID {
				local.R.Application = foreign
				if foreign.R == nil {
					foreign.R = &applicationR{}
				}
				foreign.R.Endpoints = append(foreign.R.Endpoints, local)
				break
			}
		}
	}

	return nil
}

// SetApplication of the endpoint to the related item.
// Sets o.R.Application to related.
// Adds o to related.R.Endpoints.
func (o *Endpoint) SetApplication(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Application) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"endpoints\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"application_id"}),
		strmangle.WhereClause("\"", "\"", 2, endpointPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ApplicationID = related.ID
	if o.R == nil {
		o.R = &endpointR{
			Application: related,
		}
	} else {
		o.R.Application = related
	}

	if related.R == nil {
		related.R = &applicationR{
			Endpoints: EndpointSlice{o},
		}
	} else {
		related.R.Endpoints = append(related.R.Endpoints, o)
	}

	return nil
}

// Endpoints retrieves all the records using an executor.
func Endpoints(mods ...qm.QueryMod) endpointQuery {
	mods = append(mods, qm.From("\"endpoints\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"endpoints\".*"})
	}

	return endpointQuery{q}
}

// FindEndpoint retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEndpoint(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Endpoint, error) {
	endpointObj := &Endpoint{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"endpoints\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, endpointObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from endpoints")
	}

	if err = endpointObj.doAfterSelectHooks(ctx, exec); err != nil {
		return endpointObj, err
	}

	return endpointObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Endpoint) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no endpoints provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(endpointColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	endpointInsertCacheMut.RLock()
	cache, cached := endpointInsertCache[key]
	endpointInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			endpointAllColumns,
			endpointColumnsWithDefault,
			endpointColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(endpointType, endpointMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(endpointType, endpointMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"endpoints\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"endpoints\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into endpoints")
	}

	if !cached {
		endpointInsertCacheMut.Lock()
		endpointInsertCache[key] = cache
		endpointInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Endpoint.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Endpoint) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	endpointUpdateCacheMut.RLock()
	cache, cached := endpointUpdateCache[key]
	endpointUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			endpointAllColumns,
			endpointPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update endpoints, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"endpoints\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, endpointPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(endpointType, endpointMapping, append(wl, endpointPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update endpoints row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for endpoints")
	}

	if !cached {
		endpointUpdateCacheMut.Lock()
		endpointUpdateCache[key] = cache
		endpointUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q endpointQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for endpoints")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for endpoints")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EndpointSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), endpointPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"endpoints\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, endpointPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in endpoint slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all endpoint")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Endpoint) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no endpoints provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(endpointColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	endpointUpsertCacheMut.RLock()
	cache, cached := endpointUpsertCache[key]
	endpointUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			endpointAllColumns,
			endpointColumnsWithDefault,
			endpointColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			endpointAllColumns,
			endpointPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert endpoints, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(endpointPrimaryKeyColumns))
			copy(conflict, endpointPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"endpoints\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(endpointType, endpointMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(endpointType, endpointMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert endpoints")
	}

	if !cached {
		endpointUpsertCacheMut.Lock()
		endpointUpsertCache[key] = cache
		endpointUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Endpoint record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Endpoint) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Endpoint provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), endpointPrimaryKeyMapping)
	sql := "DELETE FROM \"endpoints\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from endpoints")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for endpoints")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q endpointQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no endpointQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from endpoints")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for endpoints")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EndpointSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(endpointBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), endpointPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"endpoints\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, endpointPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from endpoint slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for endpoints")
	}

	if len(endpointAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Endpoint) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEndpoint(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EndpointSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EndpointSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), endpointPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"endpoints\".* FROM \"endpoints\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, endpointPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EndpointSlice")
	}

	*o = slice

	return nil
}

// EndpointExists checks if the Endpoint row exists.
func EndpointExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"endpoints\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if endpoints exists")
	}

	return exists, nil
}

// Exists checks if the Endpoint row exists.
func (o *Endpoint) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return EndpointExists(ctx, exec, o.ID)
}
