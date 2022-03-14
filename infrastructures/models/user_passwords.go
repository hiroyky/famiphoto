// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// UserPassword is an object representing the database table.
type UserPassword struct {
	UserID         string    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Password       string    `boil:"password" json:"password" toml:"password" yaml:"password"`
	LastModifiedAt time.Time `boil:"last_modified_at" json:"last_modified_at" toml:"last_modified_at" yaml:"last_modified_at"`
	IsInitialized  int8      `boil:"is_initialized" json:"is_initialized" toml:"is_initialized" yaml:"is_initialized"`
	CreatedAt      time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt      time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *userPasswordR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userPasswordL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserPasswordColumns = struct {
	UserID         string
	Password       string
	LastModifiedAt string
	IsInitialized  string
	CreatedAt      string
	UpdatedAt      string
}{
	UserID:         "user_id",
	Password:       "password",
	LastModifiedAt: "last_modified_at",
	IsInitialized:  "is_initialized",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

var UserPasswordTableColumns = struct {
	UserID         string
	Password       string
	LastModifiedAt string
	IsInitialized  string
	CreatedAt      string
	UpdatedAt      string
}{
	UserID:         "user_passwords.user_id",
	Password:       "user_passwords.password",
	LastModifiedAt: "user_passwords.last_modified_at",
	IsInitialized:  "user_passwords.is_initialized",
	CreatedAt:      "user_passwords.created_at",
	UpdatedAt:      "user_passwords.updated_at",
}

// Generated where

type whereHelperint8 struct{ field string }

func (w whereHelperint8) EQ(x int8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint8) NEQ(x int8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint8) LT(x int8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint8) LTE(x int8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint8) GT(x int8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint8) GTE(x int8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint8) IN(slice []int8) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint8) NIN(slice []int8) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var UserPasswordWhere = struct {
	UserID         whereHelperstring
	Password       whereHelperstring
	LastModifiedAt whereHelpertime_Time
	IsInitialized  whereHelperint8
	CreatedAt      whereHelpertime_Time
	UpdatedAt      whereHelpertime_Time
}{
	UserID:         whereHelperstring{field: "`user_passwords`.`user_id`"},
	Password:       whereHelperstring{field: "`user_passwords`.`password`"},
	LastModifiedAt: whereHelpertime_Time{field: "`user_passwords`.`last_modified_at`"},
	IsInitialized:  whereHelperint8{field: "`user_passwords`.`is_initialized`"},
	CreatedAt:      whereHelpertime_Time{field: "`user_passwords`.`created_at`"},
	UpdatedAt:      whereHelpertime_Time{field: "`user_passwords`.`updated_at`"},
}

// UserPasswordRels is where relationship names are stored.
var UserPasswordRels = struct {
	User string
}{
	User: "User",
}

// userPasswordR is where relationships are stored.
type userPasswordR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*userPasswordR) NewStruct() *userPasswordR {
	return &userPasswordR{}
}

// userPasswordL is where Load methods for each relationship are stored.
type userPasswordL struct{}

var (
	userPasswordAllColumns            = []string{"user_id", "password", "last_modified_at", "is_initialized", "created_at", "updated_at"}
	userPasswordColumnsWithoutDefault = []string{"user_id", "password", "last_modified_at", "is_initialized"}
	userPasswordColumnsWithDefault    = []string{"created_at", "updated_at"}
	userPasswordPrimaryKeyColumns     = []string{"user_id"}
	userPasswordGeneratedColumns      = []string{}
)

type (
	// UserPasswordSlice is an alias for a slice of pointers to UserPassword.
	// This should almost always be used instead of []UserPassword.
	UserPasswordSlice []*UserPassword
	// UserPasswordHook is the signature for custom UserPassword hook methods
	UserPasswordHook func(context.Context, boil.ContextExecutor, *UserPassword) error

	userPasswordQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userPasswordType                 = reflect.TypeOf(&UserPassword{})
	userPasswordMapping              = queries.MakeStructMapping(userPasswordType)
	userPasswordPrimaryKeyMapping, _ = queries.BindMapping(userPasswordType, userPasswordMapping, userPasswordPrimaryKeyColumns)
	userPasswordInsertCacheMut       sync.RWMutex
	userPasswordInsertCache          = make(map[string]insertCache)
	userPasswordUpdateCacheMut       sync.RWMutex
	userPasswordUpdateCache          = make(map[string]updateCache)
	userPasswordUpsertCacheMut       sync.RWMutex
	userPasswordUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userPasswordAfterSelectHooks []UserPasswordHook

var userPasswordBeforeInsertHooks []UserPasswordHook
var userPasswordAfterInsertHooks []UserPasswordHook

var userPasswordBeforeUpdateHooks []UserPasswordHook
var userPasswordAfterUpdateHooks []UserPasswordHook

var userPasswordBeforeDeleteHooks []UserPasswordHook
var userPasswordAfterDeleteHooks []UserPasswordHook

var userPasswordBeforeUpsertHooks []UserPasswordHook
var userPasswordAfterUpsertHooks []UserPasswordHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserPassword) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userPasswordAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserPassword) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userPasswordBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserPassword) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userPasswordAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserPassword) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userPasswordBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserPassword) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userPasswordAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserPassword) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userPasswordBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserPassword) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userPasswordAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserPassword) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userPasswordBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserPassword) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userPasswordAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserPasswordHook registers your hook function for all future operations.
func AddUserPasswordHook(hookPoint boil.HookPoint, userPasswordHook UserPasswordHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		userPasswordAfterSelectHooks = append(userPasswordAfterSelectHooks, userPasswordHook)
	case boil.BeforeInsertHook:
		userPasswordBeforeInsertHooks = append(userPasswordBeforeInsertHooks, userPasswordHook)
	case boil.AfterInsertHook:
		userPasswordAfterInsertHooks = append(userPasswordAfterInsertHooks, userPasswordHook)
	case boil.BeforeUpdateHook:
		userPasswordBeforeUpdateHooks = append(userPasswordBeforeUpdateHooks, userPasswordHook)
	case boil.AfterUpdateHook:
		userPasswordAfterUpdateHooks = append(userPasswordAfterUpdateHooks, userPasswordHook)
	case boil.BeforeDeleteHook:
		userPasswordBeforeDeleteHooks = append(userPasswordBeforeDeleteHooks, userPasswordHook)
	case boil.AfterDeleteHook:
		userPasswordAfterDeleteHooks = append(userPasswordAfterDeleteHooks, userPasswordHook)
	case boil.BeforeUpsertHook:
		userPasswordBeforeUpsertHooks = append(userPasswordBeforeUpsertHooks, userPasswordHook)
	case boil.AfterUpsertHook:
		userPasswordAfterUpsertHooks = append(userPasswordAfterUpsertHooks, userPasswordHook)
	}
}

// One returns a single userPassword record from the query.
func (q userPasswordQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserPassword, error) {
	o := &UserPassword{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user_passwords")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserPassword records from the query.
func (q userPasswordQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserPasswordSlice, error) {
	var o []*UserPassword

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UserPassword slice")
	}

	if len(userPasswordAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserPassword records in the query.
func (q userPasswordQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user_passwords rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userPasswordQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user_passwords exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *UserPassword) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`user_id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "`users`")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userPasswordL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserPassword interface{}, mods queries.Applicator) error {
	var slice []*UserPassword
	var object *UserPassword

	if singular {
		object = maybeUserPassword.(*UserPassword)
	} else {
		slice = *maybeUserPassword.(*[]*UserPassword)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userPasswordR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userPasswordR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.user_id in ?`, args...),
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

	if len(userPasswordAfterSelectHooks) != 0 {
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
		foreign.R.UserPassword = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.UserID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserPassword = local
				break
			}
		}
	}

	return nil
}

// SetUser of the userPassword to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserPassword.
func (o *UserPassword) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `user_passwords` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, userPasswordPrimaryKeyColumns),
	)
	values := []interface{}{related.UserID, o.UserID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.UserID
	if o.R == nil {
		o.R = &userPasswordR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			UserPassword: o,
		}
	} else {
		related.R.UserPassword = o
	}

	return nil
}

// UserPasswords retrieves all the records using an executor.
func UserPasswords(mods ...qm.QueryMod) userPasswordQuery {
	mods = append(mods, qm.From("`user_passwords`"))
	return userPasswordQuery{NewQuery(mods...)}
}

// FindUserPassword retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserPassword(ctx context.Context, exec boil.ContextExecutor, userID string, selectCols ...string) (*UserPassword, error) {
	userPasswordObj := &UserPassword{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `user_passwords` where `user_id`=?", sel,
	)

	q := queries.Raw(query, userID)

	err := q.Bind(ctx, exec, userPasswordObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user_passwords")
	}

	if err = userPasswordObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userPasswordObj, err
	}

	return userPasswordObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserPassword) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_passwords provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(userPasswordColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userPasswordInsertCacheMut.RLock()
	cache, cached := userPasswordInsertCache[key]
	userPasswordInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userPasswordAllColumns,
			userPasswordColumnsWithDefault,
			userPasswordColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userPasswordType, userPasswordMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userPasswordType, userPasswordMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `user_passwords` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `user_passwords` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `user_passwords` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, userPasswordPrimaryKeyColumns))
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into user_passwords")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.UserID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for user_passwords")
	}

CacheNoHooks:
	if !cached {
		userPasswordInsertCacheMut.Lock()
		userPasswordInsertCache[key] = cache
		userPasswordInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserPassword.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserPassword) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userPasswordUpdateCacheMut.RLock()
	cache, cached := userPasswordUpdateCache[key]
	userPasswordUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userPasswordAllColumns,
			userPasswordPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user_passwords, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `user_passwords` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, userPasswordPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userPasswordType, userPasswordMapping, append(wl, userPasswordPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update user_passwords row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for user_passwords")
	}

	if !cached {
		userPasswordUpdateCacheMut.Lock()
		userPasswordUpdateCache[key] = cache
		userPasswordUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userPasswordQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for user_passwords")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for user_passwords")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserPasswordSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userPasswordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `user_passwords` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userPasswordPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in userPassword slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all userPassword")
	}
	return rowsAff, nil
}

var mySQLUserPasswordUniqueColumns = []string{
	"user_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserPassword) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_passwords provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(userPasswordColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLUserPasswordUniqueColumns, o)

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

	userPasswordUpsertCacheMut.RLock()
	cache, cached := userPasswordUpsertCache[key]
	userPasswordUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userPasswordAllColumns,
			userPasswordColumnsWithDefault,
			userPasswordColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			userPasswordAllColumns,
			userPasswordPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert user_passwords, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`user_passwords`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `user_passwords` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(userPasswordType, userPasswordMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userPasswordType, userPasswordMapping, ret)
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for user_passwords")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(userPasswordType, userPasswordMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for user_passwords")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for user_passwords")
	}

CacheNoHooks:
	if !cached {
		userPasswordUpsertCacheMut.Lock()
		userPasswordUpsertCache[key] = cache
		userPasswordUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserPassword record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserPassword) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserPassword provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userPasswordPrimaryKeyMapping)
	sql := "DELETE FROM `user_passwords` WHERE `user_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from user_passwords")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for user_passwords")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userPasswordQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userPasswordQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user_passwords")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_passwords")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserPasswordSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userPasswordBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userPasswordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `user_passwords` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userPasswordPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from userPassword slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_passwords")
	}

	if len(userPasswordAfterDeleteHooks) != 0 {
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
func (o *UserPassword) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserPassword(ctx, exec, o.UserID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserPasswordSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserPasswordSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userPasswordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `user_passwords`.* FROM `user_passwords` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userPasswordPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserPasswordSlice")
	}

	*o = slice

	return nil
}

// UserPasswordExists checks if the UserPassword row exists.
func UserPasswordExists(ctx context.Context, exec boil.ContextExecutor, userID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `user_passwords` where `user_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, userID)
	}
	row := exec.QueryRowContext(ctx, sql, userID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user_passwords exists")
	}

	return exists, nil
}
