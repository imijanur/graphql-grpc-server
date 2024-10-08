// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// UserContact is an object representing the database table.
type UserContact struct {
	ID        int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	FirstName string `boil:"first_name" json:"first_name" toml:"first_name" yaml:"first_name"`
	LastName  string `boil:"last_name" json:"last_name" toml:"last_name" yaml:"last_name"`
	Phone     string `boil:"phone" json:"phone" toml:"phone" yaml:"phone"`
	UserID    int    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`

	R *userContactR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userContactL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserContactColumns = struct {
	ID        string
	FirstName string
	LastName  string
	Phone     string
	UserID    string
}{
	ID:        "id",
	FirstName: "first_name",
	LastName:  "last_name",
	Phone:     "phone",
	UserID:    "user_id",
}

var UserContactTableColumns = struct {
	ID        string
	FirstName string
	LastName  string
	Phone     string
	UserID    string
}{
	ID:        "user_contact.id",
	FirstName: "user_contact.first_name",
	LastName:  "user_contact.last_name",
	Phone:     "user_contact.phone",
	UserID:    "user_contact.user_id",
}

// Generated where

var UserContactWhere = struct {
	ID        whereHelperint
	FirstName whereHelperstring
	LastName  whereHelperstring
	Phone     whereHelperstring
	UserID    whereHelperint
}{
	ID:        whereHelperint{field: "`user_contact`.`id`"},
	FirstName: whereHelperstring{field: "`user_contact`.`first_name`"},
	LastName:  whereHelperstring{field: "`user_contact`.`last_name`"},
	Phone:     whereHelperstring{field: "`user_contact`.`phone`"},
	UserID:    whereHelperint{field: "`user_contact`.`user_id`"},
}

// UserContactRels is where relationship names are stored.
var UserContactRels = struct {
	User string
}{
	User: "User",
}

// userContactR is where relationships are stored.
type userContactR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*userContactR) NewStruct() *userContactR {
	return &userContactR{}
}

func (r *userContactR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// userContactL is where Load methods for each relationship are stored.
type userContactL struct{}

var (
	userContactAllColumns            = []string{"id", "first_name", "last_name", "phone", "user_id"}
	userContactColumnsWithoutDefault = []string{"first_name", "last_name", "phone", "user_id"}
	userContactColumnsWithDefault    = []string{"id"}
	userContactPrimaryKeyColumns     = []string{"id"}
	userContactGeneratedColumns      = []string{}
)

type (
	// UserContactSlice is an alias for a slice of pointers to UserContact.
	// This should almost always be used instead of []UserContact.
	UserContactSlice []*UserContact
	// UserContactHook is the signature for custom UserContact hook methods
	UserContactHook func(context.Context, boil.ContextExecutor, *UserContact) error

	userContactQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userContactType                 = reflect.TypeOf(&UserContact{})
	userContactMapping              = queries.MakeStructMapping(userContactType)
	userContactPrimaryKeyMapping, _ = queries.BindMapping(userContactType, userContactMapping, userContactPrimaryKeyColumns)
	userContactInsertCacheMut       sync.RWMutex
	userContactInsertCache          = make(map[string]insertCache)
	userContactUpdateCacheMut       sync.RWMutex
	userContactUpdateCache          = make(map[string]updateCache)
	userContactUpsertCacheMut       sync.RWMutex
	userContactUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userContactAfterSelectMu sync.Mutex
var userContactAfterSelectHooks []UserContactHook

var userContactBeforeInsertMu sync.Mutex
var userContactBeforeInsertHooks []UserContactHook
var userContactAfterInsertMu sync.Mutex
var userContactAfterInsertHooks []UserContactHook

var userContactBeforeUpdateMu sync.Mutex
var userContactBeforeUpdateHooks []UserContactHook
var userContactAfterUpdateMu sync.Mutex
var userContactAfterUpdateHooks []UserContactHook

var userContactBeforeDeleteMu sync.Mutex
var userContactBeforeDeleteHooks []UserContactHook
var userContactAfterDeleteMu sync.Mutex
var userContactAfterDeleteHooks []UserContactHook

var userContactBeforeUpsertMu sync.Mutex
var userContactBeforeUpsertHooks []UserContactHook
var userContactAfterUpsertMu sync.Mutex
var userContactAfterUpsertHooks []UserContactHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserContact) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userContactAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserContact) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userContactBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserContact) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userContactAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserContact) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userContactBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserContact) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userContactAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserContact) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userContactBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserContact) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userContactAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserContact) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userContactBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserContact) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userContactAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserContactHook registers your hook function for all future operations.
func AddUserContactHook(hookPoint boil.HookPoint, userContactHook UserContactHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		userContactAfterSelectMu.Lock()
		userContactAfterSelectHooks = append(userContactAfterSelectHooks, userContactHook)
		userContactAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		userContactBeforeInsertMu.Lock()
		userContactBeforeInsertHooks = append(userContactBeforeInsertHooks, userContactHook)
		userContactBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		userContactAfterInsertMu.Lock()
		userContactAfterInsertHooks = append(userContactAfterInsertHooks, userContactHook)
		userContactAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		userContactBeforeUpdateMu.Lock()
		userContactBeforeUpdateHooks = append(userContactBeforeUpdateHooks, userContactHook)
		userContactBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		userContactAfterUpdateMu.Lock()
		userContactAfterUpdateHooks = append(userContactAfterUpdateHooks, userContactHook)
		userContactAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		userContactBeforeDeleteMu.Lock()
		userContactBeforeDeleteHooks = append(userContactBeforeDeleteHooks, userContactHook)
		userContactBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		userContactAfterDeleteMu.Lock()
		userContactAfterDeleteHooks = append(userContactAfterDeleteHooks, userContactHook)
		userContactAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		userContactBeforeUpsertMu.Lock()
		userContactBeforeUpsertHooks = append(userContactBeforeUpsertHooks, userContactHook)
		userContactBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		userContactAfterUpsertMu.Lock()
		userContactAfterUpsertHooks = append(userContactAfterUpsertHooks, userContactHook)
		userContactAfterUpsertMu.Unlock()
	}
}

// One returns a single userContact record from the query.
func (q userContactQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserContact, error) {
	o := &UserContact{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user_contact")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserContact records from the query.
func (q userContactQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserContactSlice, error) {
	var o []*UserContact

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UserContact slice")
	}

	if len(userContactAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserContact records in the query.
func (q userContactQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user_contact rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userContactQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user_contact exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *UserContact) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userContactL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserContact interface{}, mods queries.Applicator) error {
	var slice []*UserContact
	var object *UserContact

	if singular {
		var ok bool
		object, ok = maybeUserContact.(*UserContact)
		if !ok {
			object = new(UserContact)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUserContact)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUserContact))
			}
		}
	} else {
		s, ok := maybeUserContact.(*[]*UserContact)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUserContact)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUserContact))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &userContactR{}
		}
		args[object.UserID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userContactR{}
			}

			args[obj.UserID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
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
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.UserContacts = append(foreign.R.UserContacts, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserContacts = append(foreign.R.UserContacts, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the userContact to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserContacts.
func (o *UserContact) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `user_contact` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, userContactPrimaryKeyColumns),
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

	o.UserID = related.ID
	if o.R == nil {
		o.R = &userContactR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			UserContacts: UserContactSlice{o},
		}
	} else {
		related.R.UserContacts = append(related.R.UserContacts, o)
	}

	return nil
}

// UserContacts retrieves all the records using an executor.
func UserContacts(mods ...qm.QueryMod) userContactQuery {
	mods = append(mods, qm.From("`user_contact`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`user_contact`.*"})
	}

	return userContactQuery{q}
}

// FindUserContact retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserContact(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*UserContact, error) {
	userContactObj := &UserContact{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `user_contact` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, userContactObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user_contact")
	}

	if err = userContactObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userContactObj, err
	}

	return userContactObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserContact) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_contact provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userContactColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userContactInsertCacheMut.RLock()
	cache, cached := userContactInsertCache[key]
	userContactInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userContactAllColumns,
			userContactColumnsWithDefault,
			userContactColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userContactType, userContactMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userContactType, userContactMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `user_contact` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `user_contact` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `user_contact` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, userContactPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into user_contact")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == userContactMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for user_contact")
	}

CacheNoHooks:
	if !cached {
		userContactInsertCacheMut.Lock()
		userContactInsertCache[key] = cache
		userContactInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserContact.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserContact) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userContactUpdateCacheMut.RLock()
	cache, cached := userContactUpdateCache[key]
	userContactUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userContactAllColumns,
			userContactPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user_contact, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `user_contact` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, userContactPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userContactType, userContactMapping, append(wl, userContactPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update user_contact row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for user_contact")
	}

	if !cached {
		userContactUpdateCacheMut.Lock()
		userContactUpdateCache[key] = cache
		userContactUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userContactQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for user_contact")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for user_contact")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserContactSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userContactPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `user_contact` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userContactPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in userContact slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all userContact")
	}
	return rowsAff, nil
}

var mySQLUserContactUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserContact) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_contact provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userContactColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLUserContactUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	userContactUpsertCacheMut.RLock()
	cache, cached := userContactUpsertCache[key]
	userContactUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			userContactAllColumns,
			userContactColumnsWithDefault,
			userContactColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			userContactAllColumns,
			userContactPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert user_contact, could not build update column list")
		}

		ret := strmangle.SetComplement(userContactAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`user_contact`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `user_contact` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(userContactType, userContactMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userContactType, userContactMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for user_contact")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == userContactMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(userContactType, userContactMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for user_contact")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for user_contact")
	}

CacheNoHooks:
	if !cached {
		userContactUpsertCacheMut.Lock()
		userContactUpsertCache[key] = cache
		userContactUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserContact record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserContact) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserContact provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userContactPrimaryKeyMapping)
	sql := "DELETE FROM `user_contact` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from user_contact")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for user_contact")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userContactQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userContactQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user_contact")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_contact")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserContactSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userContactBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userContactPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `user_contact` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userContactPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from userContact slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_contact")
	}

	if len(userContactAfterDeleteHooks) != 0 {
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
func (o *UserContact) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserContact(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserContactSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserContactSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userContactPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `user_contact`.* FROM `user_contact` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userContactPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserContactSlice")
	}

	*o = slice

	return nil
}

// UserContactExists checks if the UserContact row exists.
func UserContactExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `user_contact` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user_contact exists")
	}

	return exists, nil
}

// Exists checks if the UserContact row exists.
func (o *UserContact) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return UserContactExists(ctx, exec, o.ID)
}
