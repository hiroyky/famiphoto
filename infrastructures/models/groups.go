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

// Group is an object representing the database table.
type Group struct {
	GroupID int    `boil:"group_id" json:"group_id" toml:"group_id" yaml:"group_id"`
	Name    string `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *groupR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L groupL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GroupColumns = struct {
	GroupID string
	Name    string
}{
	GroupID: "group_id",
	Name:    "name",
}

var GroupTableColumns = struct {
	GroupID string
	Name    string
}{
	GroupID: "groups.group_id",
	Name:    "groups.name",
}

// Generated where

var GroupWhere = struct {
	GroupID whereHelperint
	Name    whereHelperstring
}{
	GroupID: whereHelperint{field: "`groups`.`group_id`"},
	Name:    whereHelperstring{field: "`groups`.`name`"},
}

// GroupRels is where relationship names are stored.
var GroupRels = struct {
	GroupUsers string
	Photos     string
}{
	GroupUsers: "GroupUsers",
	Photos:     "Photos",
}

// groupR is where relationships are stored.
type groupR struct {
	GroupUsers GroupUserSlice `boil:"GroupUsers" json:"GroupUsers" toml:"GroupUsers" yaml:"GroupUsers"`
	Photos     PhotoSlice     `boil:"Photos" json:"Photos" toml:"Photos" yaml:"Photos"`
}

// NewStruct creates a new relationship struct
func (*groupR) NewStruct() *groupR {
	return &groupR{}
}

// groupL is where Load methods for each relationship are stored.
type groupL struct{}

var (
	groupAllColumns            = []string{"group_id", "name"}
	groupColumnsWithoutDefault = []string{"name"}
	groupColumnsWithDefault    = []string{"group_id"}
	groupPrimaryKeyColumns     = []string{"group_id"}
	groupGeneratedColumns      = []string{}
)

type (
	// GroupSlice is an alias for a slice of pointers to Group.
	// This should almost always be used instead of []Group.
	GroupSlice []*Group
	// GroupHook is the signature for custom Group hook methods
	GroupHook func(context.Context, boil.ContextExecutor, *Group) error

	groupQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	groupType                 = reflect.TypeOf(&Group{})
	groupMapping              = queries.MakeStructMapping(groupType)
	groupPrimaryKeyMapping, _ = queries.BindMapping(groupType, groupMapping, groupPrimaryKeyColumns)
	groupInsertCacheMut       sync.RWMutex
	groupInsertCache          = make(map[string]insertCache)
	groupUpdateCacheMut       sync.RWMutex
	groupUpdateCache          = make(map[string]updateCache)
	groupUpsertCacheMut       sync.RWMutex
	groupUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var groupAfterSelectHooks []GroupHook

var groupBeforeInsertHooks []GroupHook
var groupAfterInsertHooks []GroupHook

var groupBeforeUpdateHooks []GroupHook
var groupAfterUpdateHooks []GroupHook

var groupBeforeDeleteHooks []GroupHook
var groupAfterDeleteHooks []GroupHook

var groupBeforeUpsertHooks []GroupHook
var groupAfterUpsertHooks []GroupHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Group) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Group) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Group) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Group) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Group) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Group) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Group) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Group) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Group) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGroupHook registers your hook function for all future operations.
func AddGroupHook(hookPoint boil.HookPoint, groupHook GroupHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		groupAfterSelectHooks = append(groupAfterSelectHooks, groupHook)
	case boil.BeforeInsertHook:
		groupBeforeInsertHooks = append(groupBeforeInsertHooks, groupHook)
	case boil.AfterInsertHook:
		groupAfterInsertHooks = append(groupAfterInsertHooks, groupHook)
	case boil.BeforeUpdateHook:
		groupBeforeUpdateHooks = append(groupBeforeUpdateHooks, groupHook)
	case boil.AfterUpdateHook:
		groupAfterUpdateHooks = append(groupAfterUpdateHooks, groupHook)
	case boil.BeforeDeleteHook:
		groupBeforeDeleteHooks = append(groupBeforeDeleteHooks, groupHook)
	case boil.AfterDeleteHook:
		groupAfterDeleteHooks = append(groupAfterDeleteHooks, groupHook)
	case boil.BeforeUpsertHook:
		groupBeforeUpsertHooks = append(groupBeforeUpsertHooks, groupHook)
	case boil.AfterUpsertHook:
		groupAfterUpsertHooks = append(groupAfterUpsertHooks, groupHook)
	}
}

// One returns a single group record from the query.
func (q groupQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Group, error) {
	o := &Group{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for groups")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Group records from the query.
func (q groupQuery) All(ctx context.Context, exec boil.ContextExecutor) (GroupSlice, error) {
	var o []*Group

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Group slice")
	}

	if len(groupAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Group records in the query.
func (q groupQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count groups rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q groupQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if groups exists")
	}

	return count > 0, nil
}

// GroupUsers retrieves all the group_user's GroupUsers with an executor.
func (o *Group) GroupUsers(mods ...qm.QueryMod) groupUserQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`group_users`.`group_id`=?", o.GroupID),
	)

	query := GroupUsers(queryMods...)
	queries.SetFrom(query.Query, "`group_users`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`group_users`.*"})
	}

	return query
}

// Photos retrieves all the photo's Photos with an executor.
func (o *Group) Photos(mods ...qm.QueryMod) photoQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`photos`.`group_id`=?", o.GroupID),
	)

	query := Photos(queryMods...)
	queries.SetFrom(query.Query, "`photos`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`photos`.*"})
	}

	return query
}

// LoadGroupUsers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (groupL) LoadGroupUsers(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGroup interface{}, mods queries.Applicator) error {
	var slice []*Group
	var object *Group

	if singular {
		object = maybeGroup.(*Group)
	} else {
		slice = *maybeGroup.(*[]*Group)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &groupR{}
		}
		args = append(args, object.GroupID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &groupR{}
			}

			for _, a := range args {
				if a == obj.GroupID {
					continue Outer
				}
			}

			args = append(args, obj.GroupID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`group_users`),
		qm.WhereIn(`group_users.group_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load group_users")
	}

	var resultSlice []*GroupUser
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice group_users")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on group_users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for group_users")
	}

	if len(groupUserAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.GroupUsers = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &groupUserR{}
			}
			foreign.R.Group = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.GroupID == foreign.GroupID {
				local.R.GroupUsers = append(local.R.GroupUsers, foreign)
				if foreign.R == nil {
					foreign.R = &groupUserR{}
				}
				foreign.R.Group = local
				break
			}
		}
	}

	return nil
}

// LoadPhotos allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (groupL) LoadPhotos(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGroup interface{}, mods queries.Applicator) error {
	var slice []*Group
	var object *Group

	if singular {
		object = maybeGroup.(*Group)
	} else {
		slice = *maybeGroup.(*[]*Group)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &groupR{}
		}
		args = append(args, object.GroupID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &groupR{}
			}

			for _, a := range args {
				if a == obj.GroupID {
					continue Outer
				}
			}

			args = append(args, obj.GroupID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`photos`),
		qm.WhereIn(`photos.group_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load photos")
	}

	var resultSlice []*Photo
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice photos")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on photos")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for photos")
	}

	if len(photoAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Photos = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &photoR{}
			}
			foreign.R.Group = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.GroupID == foreign.GroupID {
				local.R.Photos = append(local.R.Photos, foreign)
				if foreign.R == nil {
					foreign.R = &photoR{}
				}
				foreign.R.Group = local
				break
			}
		}
	}

	return nil
}

// AddGroupUsers adds the given related objects to the existing relationships
// of the group, optionally inserting them as new records.
// Appends related to o.R.GroupUsers.
// Sets related.R.Group appropriately.
func (o *Group) AddGroupUsers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*GroupUser) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.GroupID = o.GroupID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `group_users` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"group_id"}),
				strmangle.WhereClause("`", "`", 0, groupUserPrimaryKeyColumns),
			)
			values := []interface{}{o.GroupID, rel.GroupID, rel.UserID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.GroupID = o.GroupID
		}
	}

	if o.R == nil {
		o.R = &groupR{
			GroupUsers: related,
		}
	} else {
		o.R.GroupUsers = append(o.R.GroupUsers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &groupUserR{
				Group: o,
			}
		} else {
			rel.R.Group = o
		}
	}
	return nil
}

// AddPhotos adds the given related objects to the existing relationships
// of the group, optionally inserting them as new records.
// Appends related to o.R.Photos.
// Sets related.R.Group appropriately.
func (o *Group) AddPhotos(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Photo) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.GroupID = o.GroupID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `photos` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"group_id"}),
				strmangle.WhereClause("`", "`", 0, photoPrimaryKeyColumns),
			)
			values := []interface{}{o.GroupID, rel.PhotoID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.GroupID = o.GroupID
		}
	}

	if o.R == nil {
		o.R = &groupR{
			Photos: related,
		}
	} else {
		o.R.Photos = append(o.R.Photos, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &photoR{
				Group: o,
			}
		} else {
			rel.R.Group = o
		}
	}
	return nil
}

// Groups retrieves all the records using an executor.
func Groups(mods ...qm.QueryMod) groupQuery {
	mods = append(mods, qm.From("`groups`"))
	return groupQuery{NewQuery(mods...)}
}

// FindGroup retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGroup(ctx context.Context, exec boil.ContextExecutor, groupID int, selectCols ...string) (*Group, error) {
	groupObj := &Group{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `groups` where `group_id`=?", sel,
	)

	q := queries.Raw(query, groupID)

	err := q.Bind(ctx, exec, groupObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from groups")
	}

	if err = groupObj.doAfterSelectHooks(ctx, exec); err != nil {
		return groupObj, err
	}

	return groupObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Group) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no groups provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(groupColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	groupInsertCacheMut.RLock()
	cache, cached := groupInsertCache[key]
	groupInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			groupAllColumns,
			groupColumnsWithDefault,
			groupColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(groupType, groupMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(groupType, groupMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `groups` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `groups` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `groups` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, groupPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into groups")
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

	o.GroupID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == groupMapping["group_id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.GroupID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for groups")
	}

CacheNoHooks:
	if !cached {
		groupInsertCacheMut.Lock()
		groupInsertCache[key] = cache
		groupInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Group.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Group) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	groupUpdateCacheMut.RLock()
	cache, cached := groupUpdateCache[key]
	groupUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			groupAllColumns,
			groupPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update groups, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `groups` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, groupPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(groupType, groupMapping, append(wl, groupPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update groups row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for groups")
	}

	if !cached {
		groupUpdateCacheMut.Lock()
		groupUpdateCache[key] = cache
		groupUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q groupQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for groups")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for groups")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GroupSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `groups` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in group slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all group")
	}
	return rowsAff, nil
}

var mySQLGroupUniqueColumns = []string{
	"group_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Group) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no groups provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(groupColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLGroupUniqueColumns, o)

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

	groupUpsertCacheMut.RLock()
	cache, cached := groupUpsertCache[key]
	groupUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			groupAllColumns,
			groupColumnsWithDefault,
			groupColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			groupAllColumns,
			groupPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert groups, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`groups`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `groups` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(groupType, groupMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(groupType, groupMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for groups")
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

	o.GroupID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == groupMapping["group_id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(groupType, groupMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for groups")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for groups")
	}

CacheNoHooks:
	if !cached {
		groupUpsertCacheMut.Lock()
		groupUpsertCache[key] = cache
		groupUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Group record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Group) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Group provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), groupPrimaryKeyMapping)
	sql := "DELETE FROM `groups` WHERE `group_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from groups")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for groups")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q groupQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no groupQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from groups")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for groups")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GroupSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(groupBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `groups` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from group slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for groups")
	}

	if len(groupAfterDeleteHooks) != 0 {
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
func (o *Group) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGroup(ctx, exec, o.GroupID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GroupSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GroupSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `groups`.* FROM `groups` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in GroupSlice")
	}

	*o = slice

	return nil
}

// GroupExists checks if the Group row exists.
func GroupExists(ctx context.Context, exec boil.ContextExecutor, groupID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `groups` where `group_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, groupID)
	}
	row := exec.QueryRowContext(ctx, sql, groupID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if groups exists")
	}

	return exists, nil
}
