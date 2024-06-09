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

// BookCopy is an object representing the database table.
type BookCopy struct {
	ID           string `boil:"id" json:"id" toml:"id" yaml:"id"`
	BookID       string `boil:"book_id" json:"book_id" toml:"book_id" yaml:"book_id"`
	Availability string `boil:"availability" json:"availability" toml:"availability" yaml:"availability"`

	R *bookCopyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bookCopyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BookCopyColumns = struct {
	ID           string
	BookID       string
	Availability string
}{
	ID:           "id",
	BookID:       "book_id",
	Availability: "availability",
}

var BookCopyTableColumns = struct {
	ID           string
	BookID       string
	Availability string
}{
	ID:           "book_copies.id",
	BookID:       "book_copies.book_id",
	Availability: "book_copies.availability",
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

var BookCopyWhere = struct {
	ID           whereHelperstring
	BookID       whereHelperstring
	Availability whereHelperstring
}{
	ID:           whereHelperstring{field: "\"book_copies\".\"id\""},
	BookID:       whereHelperstring{field: "\"book_copies\".\"book_id\""},
	Availability: whereHelperstring{field: "\"book_copies\".\"availability\""},
}

// BookCopyRels is where relationship names are stored.
var BookCopyRels = struct {
	Book  string
	Loans string
}{
	Book:  "Book",
	Loans: "Loans",
}

// bookCopyR is where relationships are stored.
type bookCopyR struct {
	Book  *Book     `boil:"Book" json:"Book" toml:"Book" yaml:"Book"`
	Loans LoanSlice `boil:"Loans" json:"Loans" toml:"Loans" yaml:"Loans"`
}

// NewStruct creates a new relationship struct
func (*bookCopyR) NewStruct() *bookCopyR {
	return &bookCopyR{}
}

func (r *bookCopyR) GetBook() *Book {
	if r == nil {
		return nil
	}
	return r.Book
}

func (r *bookCopyR) GetLoans() LoanSlice {
	if r == nil {
		return nil
	}
	return r.Loans
}

// bookCopyL is where Load methods for each relationship are stored.
type bookCopyL struct{}

var (
	bookCopyAllColumns            = []string{"id", "book_id", "availability"}
	bookCopyColumnsWithoutDefault = []string{"id", "book_id"}
	bookCopyColumnsWithDefault    = []string{"availability"}
	bookCopyPrimaryKeyColumns     = []string{"id"}
	bookCopyGeneratedColumns      = []string{}
)

type (
	// BookCopySlice is an alias for a slice of pointers to BookCopy.
	// This should almost always be used instead of []BookCopy.
	BookCopySlice []*BookCopy
	// BookCopyHook is the signature for custom BookCopy hook methods
	BookCopyHook func(context.Context, boil.ContextExecutor, *BookCopy) error

	bookCopyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bookCopyType                 = reflect.TypeOf(&BookCopy{})
	bookCopyMapping              = queries.MakeStructMapping(bookCopyType)
	bookCopyPrimaryKeyMapping, _ = queries.BindMapping(bookCopyType, bookCopyMapping, bookCopyPrimaryKeyColumns)
	bookCopyInsertCacheMut       sync.RWMutex
	bookCopyInsertCache          = make(map[string]insertCache)
	bookCopyUpdateCacheMut       sync.RWMutex
	bookCopyUpdateCache          = make(map[string]updateCache)
	bookCopyUpsertCacheMut       sync.RWMutex
	bookCopyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var bookCopyAfterSelectHooks []BookCopyHook

var bookCopyBeforeInsertHooks []BookCopyHook
var bookCopyAfterInsertHooks []BookCopyHook

var bookCopyBeforeUpdateHooks []BookCopyHook
var bookCopyAfterUpdateHooks []BookCopyHook

var bookCopyBeforeDeleteHooks []BookCopyHook
var bookCopyAfterDeleteHooks []BookCopyHook

var bookCopyBeforeUpsertHooks []BookCopyHook
var bookCopyAfterUpsertHooks []BookCopyHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BookCopy) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookCopyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BookCopy) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookCopyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BookCopy) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookCopyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BookCopy) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookCopyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BookCopy) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookCopyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BookCopy) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookCopyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BookCopy) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookCopyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BookCopy) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookCopyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BookCopy) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookCopyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBookCopyHook registers your hook function for all future operations.
func AddBookCopyHook(hookPoint boil.HookPoint, bookCopyHook BookCopyHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		bookCopyAfterSelectHooks = append(bookCopyAfterSelectHooks, bookCopyHook)
	case boil.BeforeInsertHook:
		bookCopyBeforeInsertHooks = append(bookCopyBeforeInsertHooks, bookCopyHook)
	case boil.AfterInsertHook:
		bookCopyAfterInsertHooks = append(bookCopyAfterInsertHooks, bookCopyHook)
	case boil.BeforeUpdateHook:
		bookCopyBeforeUpdateHooks = append(bookCopyBeforeUpdateHooks, bookCopyHook)
	case boil.AfterUpdateHook:
		bookCopyAfterUpdateHooks = append(bookCopyAfterUpdateHooks, bookCopyHook)
	case boil.BeforeDeleteHook:
		bookCopyBeforeDeleteHooks = append(bookCopyBeforeDeleteHooks, bookCopyHook)
	case boil.AfterDeleteHook:
		bookCopyAfterDeleteHooks = append(bookCopyAfterDeleteHooks, bookCopyHook)
	case boil.BeforeUpsertHook:
		bookCopyBeforeUpsertHooks = append(bookCopyBeforeUpsertHooks, bookCopyHook)
	case boil.AfterUpsertHook:
		bookCopyAfterUpsertHooks = append(bookCopyAfterUpsertHooks, bookCopyHook)
	}
}

// One returns a single bookCopy record from the query.
func (q bookCopyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BookCopy, error) {
	o := &BookCopy{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for book_copies")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all BookCopy records from the query.
func (q bookCopyQuery) All(ctx context.Context, exec boil.ContextExecutor) (BookCopySlice, error) {
	var o []*BookCopy

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BookCopy slice")
	}

	if len(bookCopyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all BookCopy records in the query.
func (q bookCopyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count book_copies rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q bookCopyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if book_copies exists")
	}

	return count > 0, nil
}

// Book pointed to by the foreign key.
func (o *BookCopy) Book(mods ...qm.QueryMod) bookQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.BookID),
	}

	queryMods = append(queryMods, mods...)

	return Books(queryMods...)
}

// Loans retrieves all the loan's Loans with an executor.
func (o *BookCopy) Loans(mods ...qm.QueryMod) loanQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"loans\".\"book_copy_id\"=?", o.ID),
	)

	return Loans(queryMods...)
}

// LoadBook allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (bookCopyL) LoadBook(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBookCopy interface{}, mods queries.Applicator) error {
	var slice []*BookCopy
	var object *BookCopy

	if singular {
		var ok bool
		object, ok = maybeBookCopy.(*BookCopy)
		if !ok {
			object = new(BookCopy)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeBookCopy)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeBookCopy))
			}
		}
	} else {
		s, ok := maybeBookCopy.(*[]*BookCopy)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeBookCopy)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeBookCopy))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &bookCopyR{}
		}
		args = append(args, object.BookID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &bookCopyR{}
			}

			for _, a := range args {
				if a == obj.BookID {
					continue Outer
				}
			}

			args = append(args, obj.BookID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`books`),
		qm.WhereIn(`books.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Book")
	}

	var resultSlice []*Book
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Book")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for books")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for books")
	}

	if len(bookAfterSelectHooks) != 0 {
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
		object.R.Book = foreign
		if foreign.R == nil {
			foreign.R = &bookR{}
		}
		foreign.R.BookCopies = append(foreign.R.BookCopies, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.BookID == foreign.ID {
				local.R.Book = foreign
				if foreign.R == nil {
					foreign.R = &bookR{}
				}
				foreign.R.BookCopies = append(foreign.R.BookCopies, local)
				break
			}
		}
	}

	return nil
}

// LoadLoans allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (bookCopyL) LoadLoans(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBookCopy interface{}, mods queries.Applicator) error {
	var slice []*BookCopy
	var object *BookCopy

	if singular {
		var ok bool
		object, ok = maybeBookCopy.(*BookCopy)
		if !ok {
			object = new(BookCopy)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeBookCopy)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeBookCopy))
			}
		}
	} else {
		s, ok := maybeBookCopy.(*[]*BookCopy)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeBookCopy)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeBookCopy))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &bookCopyR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &bookCopyR{}
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
		qm.From(`loans`),
		qm.WhereIn(`loans.book_copy_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load loans")
	}

	var resultSlice []*Loan
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice loans")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on loans")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for loans")
	}

	if len(loanAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Loans = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &loanR{}
			}
			foreign.R.BookCopy = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.BookCopyID {
				local.R.Loans = append(local.R.Loans, foreign)
				if foreign.R == nil {
					foreign.R = &loanR{}
				}
				foreign.R.BookCopy = local
				break
			}
		}
	}

	return nil
}

// SetBook of the bookCopy to the related item.
// Sets o.R.Book to related.
// Adds o to related.R.BookCopies.
func (o *BookCopy) SetBook(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Book) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"book_copies\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"book_id"}),
		strmangle.WhereClause("\"", "\"", 2, bookCopyPrimaryKeyColumns),
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

	o.BookID = related.ID
	if o.R == nil {
		o.R = &bookCopyR{
			Book: related,
		}
	} else {
		o.R.Book = related
	}

	if related.R == nil {
		related.R = &bookR{
			BookCopies: BookCopySlice{o},
		}
	} else {
		related.R.BookCopies = append(related.R.BookCopies, o)
	}

	return nil
}

// AddLoans adds the given related objects to the existing relationships
// of the book_copy, optionally inserting them as new records.
// Appends related to o.R.Loans.
// Sets related.R.BookCopy appropriately.
func (o *BookCopy) AddLoans(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Loan) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.BookCopyID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"loans\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"book_copy_id"}),
				strmangle.WhereClause("\"", "\"", 2, loanPrimaryKeyColumns),
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

			rel.BookCopyID = o.ID
		}
	}

	if o.R == nil {
		o.R = &bookCopyR{
			Loans: related,
		}
	} else {
		o.R.Loans = append(o.R.Loans, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &loanR{
				BookCopy: o,
			}
		} else {
			rel.R.BookCopy = o
		}
	}
	return nil
}

// BookCopies retrieves all the records using an executor.
func BookCopies(mods ...qm.QueryMod) bookCopyQuery {
	mods = append(mods, qm.From("\"book_copies\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"book_copies\".*"})
	}

	return bookCopyQuery{q}
}

// FindBookCopy retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBookCopy(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*BookCopy, error) {
	bookCopyObj := &BookCopy{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"book_copies\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, bookCopyObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from book_copies")
	}

	if err = bookCopyObj.doAfterSelectHooks(ctx, exec); err != nil {
		return bookCopyObj, err
	}

	return bookCopyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BookCopy) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no book_copies provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bookCopyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bookCopyInsertCacheMut.RLock()
	cache, cached := bookCopyInsertCache[key]
	bookCopyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bookCopyAllColumns,
			bookCopyColumnsWithDefault,
			bookCopyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(bookCopyType, bookCopyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bookCopyType, bookCopyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"book_copies\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"book_copies\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into book_copies")
	}

	if !cached {
		bookCopyInsertCacheMut.Lock()
		bookCopyInsertCache[key] = cache
		bookCopyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the BookCopy.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BookCopy) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	bookCopyUpdateCacheMut.RLock()
	cache, cached := bookCopyUpdateCache[key]
	bookCopyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bookCopyAllColumns,
			bookCopyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update book_copies, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"book_copies\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, bookCopyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bookCopyType, bookCopyMapping, append(wl, bookCopyPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update book_copies row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for book_copies")
	}

	if !cached {
		bookCopyUpdateCacheMut.Lock()
		bookCopyUpdateCache[key] = cache
		bookCopyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q bookCopyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for book_copies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for book_copies")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BookCopySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookCopyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"book_copies\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, bookCopyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in bookCopy slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all bookCopy")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BookCopy) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no book_copies provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bookCopyColumnsWithDefault, o)

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

	bookCopyUpsertCacheMut.RLock()
	cache, cached := bookCopyUpsertCache[key]
	bookCopyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bookCopyAllColumns,
			bookCopyColumnsWithDefault,
			bookCopyColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			bookCopyAllColumns,
			bookCopyPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert book_copies, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(bookCopyPrimaryKeyColumns))
			copy(conflict, bookCopyPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"book_copies\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(bookCopyType, bookCopyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bookCopyType, bookCopyMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert book_copies")
	}

	if !cached {
		bookCopyUpsertCacheMut.Lock()
		bookCopyUpsertCache[key] = cache
		bookCopyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single BookCopy record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BookCopy) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BookCopy provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bookCopyPrimaryKeyMapping)
	sql := "DELETE FROM \"book_copies\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from book_copies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for book_copies")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q bookCopyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no bookCopyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from book_copies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for book_copies")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BookCopySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(bookCopyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookCopyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"book_copies\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bookCopyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bookCopy slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for book_copies")
	}

	if len(bookCopyAfterDeleteHooks) != 0 {
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
func (o *BookCopy) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBookCopy(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BookCopySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BookCopySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookCopyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"book_copies\".* FROM \"book_copies\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bookCopyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BookCopySlice")
	}

	*o = slice

	return nil
}

// BookCopyExists checks if the BookCopy row exists.
func BookCopyExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"book_copies\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if book_copies exists")
	}

	return exists, nil
}

// Exists checks if the BookCopy row exists.
func (o *BookCopy) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return BookCopyExists(ctx, exec, o.ID)
}
