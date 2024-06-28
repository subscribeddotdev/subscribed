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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Application is an object representing the database table.
type Application struct {
	ID            string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	EnvironmentID string    `boil:"environment_id" json:"environment_id" toml:"environment_id" yaml:"environment_id"`
	Name          string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedAt     time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *applicationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L applicationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ApplicationColumns = struct {
	ID            string
	EnvironmentID string
	Name          string
	CreatedAt     string
}{
	ID:            "id",
	EnvironmentID: "environment_id",
	Name:          "name",
	CreatedAt:     "created_at",
}

var ApplicationTableColumns = struct {
	ID            string
	EnvironmentID string
	Name          string
	CreatedAt     string
}{
	ID:            "applications.id",
	EnvironmentID: "applications.environment_id",
	Name:          "applications.name",
	CreatedAt:     "applications.created_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod   { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT ILIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ApplicationWhere = struct {
	ID            whereHelperstring
	EnvironmentID whereHelperstring
	Name          whereHelperstring
	CreatedAt     whereHelpertime_Time
}{
	ID:            whereHelperstring{field: "\"applications\".\"id\""},
	EnvironmentID: whereHelperstring{field: "\"applications\".\"environment_id\""},
	Name:          whereHelperstring{field: "\"applications\".\"name\""},
	CreatedAt:     whereHelpertime_Time{field: "\"applications\".\"created_at\""},
}

// ApplicationRels is where relationship names are stored.
var ApplicationRels = struct {
	Environment string
	Endpoints   string
}{
	Environment: "Environment",
	Endpoints:   "Endpoints",
}

// applicationR is where relationships are stored.
type applicationR struct {
	Environment *Environment  `boil:"Environment" json:"Environment" toml:"Environment" yaml:"Environment"`
	Endpoints   EndpointSlice `boil:"Endpoints" json:"Endpoints" toml:"Endpoints" yaml:"Endpoints"`
}

// NewStruct creates a new relationship struct
func (*applicationR) NewStruct() *applicationR {
	return &applicationR{}
}

func (r *applicationR) GetEnvironment() *Environment {
	if r == nil {
		return nil
	}
	return r.Environment
}

func (r *applicationR) GetEndpoints() EndpointSlice {
	if r == nil {
		return nil
	}
	return r.Endpoints
}

// applicationL is where Load methods for each relationship are stored.
type applicationL struct{}

var (
	applicationAllColumns            = []string{"id", "environment_id", "name", "created_at"}
	applicationColumnsWithoutDefault = []string{"id", "environment_id", "name", "created_at"}
	applicationColumnsWithDefault    = []string{}
	applicationPrimaryKeyColumns     = []string{"id"}
	applicationGeneratedColumns      = []string{}
)

type (
	// ApplicationSlice is an alias for a slice of pointers to Application.
	// This should almost always be used instead of []Application.
	ApplicationSlice []*Application
	// ApplicationHook is the signature for custom Application hook methods
	ApplicationHook func(context.Context, boil.ContextExecutor, *Application) error

	applicationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	applicationType                 = reflect.TypeOf(&Application{})
	applicationMapping              = queries.MakeStructMapping(applicationType)
	applicationPrimaryKeyMapping, _ = queries.BindMapping(applicationType, applicationMapping, applicationPrimaryKeyColumns)
	applicationInsertCacheMut       sync.RWMutex
	applicationInsertCache          = make(map[string]insertCache)
	applicationUpdateCacheMut       sync.RWMutex
	applicationUpdateCache          = make(map[string]updateCache)
	applicationUpsertCacheMut       sync.RWMutex
	applicationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/SentAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var applicationAfterSelectHooks []ApplicationHook

var applicationBeforeInsertHooks []ApplicationHook
var applicationAfterInsertHooks []ApplicationHook

var applicationBeforeUpdateHooks []ApplicationHook
var applicationAfterUpdateHooks []ApplicationHook

var applicationBeforeDeleteHooks []ApplicationHook
var applicationAfterDeleteHooks []ApplicationHook

var applicationBeforeUpsertHooks []ApplicationHook
var applicationAfterUpsertHooks []ApplicationHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Application) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range applicationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Application) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range applicationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Application) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range applicationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Application) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range applicationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Application) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range applicationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Application) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range applicationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Application) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range applicationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Application) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range applicationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Application) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range applicationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddApplicationHook registers your hook function for all future operations.
func AddApplicationHook(hookPoint boil.HookPoint, applicationHook ApplicationHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		applicationAfterSelectHooks = append(applicationAfterSelectHooks, applicationHook)
	case boil.BeforeInsertHook:
		applicationBeforeInsertHooks = append(applicationBeforeInsertHooks, applicationHook)
	case boil.AfterInsertHook:
		applicationAfterInsertHooks = append(applicationAfterInsertHooks, applicationHook)
	case boil.BeforeUpdateHook:
		applicationBeforeUpdateHooks = append(applicationBeforeUpdateHooks, applicationHook)
	case boil.AfterUpdateHook:
		applicationAfterUpdateHooks = append(applicationAfterUpdateHooks, applicationHook)
	case boil.BeforeDeleteHook:
		applicationBeforeDeleteHooks = append(applicationBeforeDeleteHooks, applicationHook)
	case boil.AfterDeleteHook:
		applicationAfterDeleteHooks = append(applicationAfterDeleteHooks, applicationHook)
	case boil.BeforeUpsertHook:
		applicationBeforeUpsertHooks = append(applicationBeforeUpsertHooks, applicationHook)
	case boil.AfterUpsertHook:
		applicationAfterUpsertHooks = append(applicationAfterUpsertHooks, applicationHook)
	}
}

// One returns a single application record from the query.
func (q applicationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Application, error) {
	o := &Application{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for applications")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Application records from the query.
func (q applicationQuery) All(ctx context.Context, exec boil.ContextExecutor) (ApplicationSlice, error) {
	var o []*Application

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Application slice")
	}

	if len(applicationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Application records in the query.
func (q applicationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count applications rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q applicationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if applications exists")
	}

	return count > 0, nil
}

// Environment pointed to by the foreign key.
func (o *Application) Environment(mods ...qm.QueryMod) environmentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.EnvironmentID),
	}

	queryMods = append(queryMods, mods...)

	return Environments(queryMods...)
}

// Endpoints retrieves all the endpoint's Endpoints with an executor.
func (o *Application) Endpoints(mods ...qm.QueryMod) endpointQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"endpoints\".\"application_id\"=?", o.ID),
	)

	return Endpoints(queryMods...)
}

// LoadEnvironment allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (applicationL) LoadEnvironment(ctx context.Context, e boil.ContextExecutor, singular bool, maybeApplication interface{}, mods queries.Applicator) error {
	var slice []*Application
	var object *Application

	if singular {
		var ok bool
		object, ok = maybeApplication.(*Application)
		if !ok {
			object = new(Application)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeApplication)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeApplication))
			}
		}
	} else {
		s, ok := maybeApplication.(*[]*Application)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeApplication)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeApplication))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &applicationR{}
		}
		args = append(args, object.EnvironmentID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &applicationR{}
			}

			for _, a := range args {
				if a == obj.EnvironmentID {
					continue Outer
				}
			}

			args = append(args, obj.EnvironmentID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`environments`),
		qm.WhereIn(`environments.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Environment")
	}

	var resultSlice []*Environment
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Environment")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for environments")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for environments")
	}

	if len(environmentAfterSelectHooks) != 0 {
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
		object.R.Environment = foreign
		if foreign.R == nil {
			foreign.R = &environmentR{}
		}
		foreign.R.Applications = append(foreign.R.Applications, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.EnvironmentID == foreign.ID {
				local.R.Environment = foreign
				if foreign.R == nil {
					foreign.R = &environmentR{}
				}
				foreign.R.Applications = append(foreign.R.Applications, local)
				break
			}
		}
	}

	return nil
}

// LoadEndpoints allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (applicationL) LoadEndpoints(ctx context.Context, e boil.ContextExecutor, singular bool, maybeApplication interface{}, mods queries.Applicator) error {
	var slice []*Application
	var object *Application

	if singular {
		var ok bool
		object, ok = maybeApplication.(*Application)
		if !ok {
			object = new(Application)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeApplication)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeApplication))
			}
		}
	} else {
		s, ok := maybeApplication.(*[]*Application)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeApplication)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeApplication))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &applicationR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &applicationR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`endpoints`),
		qm.WhereIn(`endpoints.application_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load endpoints")
	}

	var resultSlice []*Endpoint
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice endpoints")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on endpoints")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for endpoints")
	}

	if len(endpointAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Endpoints = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &endpointR{}
			}
			foreign.R.Application = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ApplicationID {
				local.R.Endpoints = append(local.R.Endpoints, foreign)
				if foreign.R == nil {
					foreign.R = &endpointR{}
				}
				foreign.R.Application = local
				break
			}
		}
	}

	return nil
}

// SetEnvironment of the application to the related item.
// Sets o.R.Environment to related.
// Adds o to related.R.Applications.
func (o *Application) SetEnvironment(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Environment) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"applications\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"environment_id"}),
		strmangle.WhereClause("\"", "\"", 2, applicationPrimaryKeyColumns),
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

	o.EnvironmentID = related.ID
	if o.R == nil {
		o.R = &applicationR{
			Environment: related,
		}
	} else {
		o.R.Environment = related
	}

	if related.R == nil {
		related.R = &environmentR{
			Applications: ApplicationSlice{o},
		}
	} else {
		related.R.Applications = append(related.R.Applications, o)
	}

	return nil
}

// AddEndpoints adds the given related objects to the existing relationships
// of the application, optionally inserting them as new records.
// Appends related to o.R.Endpoints.
// Sets related.R.Application appropriately.
func (o *Application) AddEndpoints(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Endpoint) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ApplicationID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"endpoints\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"application_id"}),
				strmangle.WhereClause("\"", "\"", 2, endpointPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ApplicationID = o.ID
		}
	}

	if o.R == nil {
		o.R = &applicationR{
			Endpoints: related,
		}
	} else {
		o.R.Endpoints = append(o.R.Endpoints, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &endpointR{
				Application: o,
			}
		} else {
			rel.R.Application = o
		}
	}
	return nil
}

// Applications retrieves all the records using an executor.
func Applications(mods ...qm.QueryMod) applicationQuery {
	mods = append(mods, qm.From("\"applications\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"applications\".*"})
	}

	return applicationQuery{q}
}

// FindApplication retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindApplication(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Application, error) {
	applicationObj := &Application{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"applications\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, applicationObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from applications")
	}

	if err = applicationObj.doAfterSelectHooks(ctx, exec); err != nil {
		return applicationObj, err
	}

	return applicationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Application) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no applications provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(applicationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	applicationInsertCacheMut.RLock()
	cache, cached := applicationInsertCache[key]
	applicationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			applicationAllColumns,
			applicationColumnsWithDefault,
			applicationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(applicationType, applicationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(applicationType, applicationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"applications\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"applications\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into applications")
	}

	if !cached {
		applicationInsertCacheMut.Lock()
		applicationInsertCache[key] = cache
		applicationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Application.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Application) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	applicationUpdateCacheMut.RLock()
	cache, cached := applicationUpdateCache[key]
	applicationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			applicationAllColumns,
			applicationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update applications, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"applications\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, applicationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(applicationType, applicationMapping, append(wl, applicationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update applications row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for applications")
	}

	if !cached {
		applicationUpdateCacheMut.Lock()
		applicationUpdateCache[key] = cache
		applicationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q applicationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for applications")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for applications")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ApplicationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), applicationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"applications\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, applicationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in application slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all application")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Application) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no applications provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(applicationColumnsWithDefault, o)

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

	applicationUpsertCacheMut.RLock()
	cache, cached := applicationUpsertCache[key]
	applicationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			applicationAllColumns,
			applicationColumnsWithDefault,
			applicationColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			applicationAllColumns,
			applicationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert applications, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(applicationPrimaryKeyColumns))
			copy(conflict, applicationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"applications\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(applicationType, applicationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(applicationType, applicationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert applications")
	}

	if !cached {
		applicationUpsertCacheMut.Lock()
		applicationUpsertCache[key] = cache
		applicationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Application record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Application) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Application provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), applicationPrimaryKeyMapping)
	sql := "DELETE FROM \"applications\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from applications")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for applications")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q applicationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no applicationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from applications")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for applications")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ApplicationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(applicationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), applicationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"applications\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, applicationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from application slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for applications")
	}

	if len(applicationAfterDeleteHooks) != 0 {
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
func (o *Application) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindApplication(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ApplicationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ApplicationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), applicationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"applications\".* FROM \"applications\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, applicationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ApplicationSlice")
	}

	*o = slice

	return nil
}

// ApplicationExists checks if the Application row exists.
func ApplicationExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"applications\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if applications exists")
	}

	return exists, nil
}

// Exists checks if the Application row exists.
func (o *Application) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ApplicationExists(ctx, exec, o.ID)
}
