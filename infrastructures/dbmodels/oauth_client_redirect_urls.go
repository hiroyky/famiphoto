// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

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

// OauthClientRedirectURL is an object representing the database table.
type OauthClientRedirectURL struct {
	OauthClientID            string    `boil:"oauth_client_id" json:"oauth_client_id" toml:"oauth_client_id" yaml:"oauth_client_id"`
	RedirectURL              string    `boil:"redirect_url" json:"redirect_url" toml:"redirect_url" yaml:"redirect_url"`
	OauthClientRedirectURLID int       `boil:"oauth_client_redirect_url_id" json:"oauth_client_redirect_url_id" toml:"oauth_client_redirect_url_id" yaml:"oauth_client_redirect_url_id"`
	CreatedAt                time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt                time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *oauthClientRedirectURLR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L oauthClientRedirectURLL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OauthClientRedirectURLColumns = struct {
	OauthClientID            string
	RedirectURL              string
	OauthClientRedirectURLID string
	CreatedAt                string
	UpdatedAt                string
}{
	OauthClientID:            "oauth_client_id",
	RedirectURL:              "redirect_url",
	OauthClientRedirectURLID: "oauth_client_redirect_url_id",
	CreatedAt:                "created_at",
	UpdatedAt:                "updated_at",
}

var OauthClientRedirectURLTableColumns = struct {
	OauthClientID            string
	RedirectURL              string
	OauthClientRedirectURLID string
	CreatedAt                string
	UpdatedAt                string
}{
	OauthClientID:            "oauth_client_redirect_urls.oauth_client_id",
	RedirectURL:              "oauth_client_redirect_urls.redirect_url",
	OauthClientRedirectURLID: "oauth_client_redirect_urls.oauth_client_redirect_url_id",
	CreatedAt:                "oauth_client_redirect_urls.created_at",
	UpdatedAt:                "oauth_client_redirect_urls.updated_at",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var OauthClientRedirectURLWhere = struct {
	OauthClientID            whereHelperstring
	RedirectURL              whereHelperstring
	OauthClientRedirectURLID whereHelperint
	CreatedAt                whereHelpertime_Time
	UpdatedAt                whereHelpertime_Time
}{
	OauthClientID:            whereHelperstring{field: "`oauth_client_redirect_urls`.`oauth_client_id`"},
	RedirectURL:              whereHelperstring{field: "`oauth_client_redirect_urls`.`redirect_url`"},
	OauthClientRedirectURLID: whereHelperint{field: "`oauth_client_redirect_urls`.`oauth_client_redirect_url_id`"},
	CreatedAt:                whereHelpertime_Time{field: "`oauth_client_redirect_urls`.`created_at`"},
	UpdatedAt:                whereHelpertime_Time{field: "`oauth_client_redirect_urls`.`updated_at`"},
}

// OauthClientRedirectURLRels is where relationship names are stored.
var OauthClientRedirectURLRels = struct {
	OauthClient string
}{
	OauthClient: "OauthClient",
}

// oauthClientRedirectURLR is where relationships are stored.
type oauthClientRedirectURLR struct {
	OauthClient *OauthClient `boil:"OauthClient" json:"OauthClient" toml:"OauthClient" yaml:"OauthClient"`
}

// NewStruct creates a new relationship struct
func (*oauthClientRedirectURLR) NewStruct() *oauthClientRedirectURLR {
	return &oauthClientRedirectURLR{}
}

// oauthClientRedirectURLL is where Load methods for each relationship are stored.
type oauthClientRedirectURLL struct{}

var (
	oauthClientRedirectURLAllColumns            = []string{"oauth_client_id", "redirect_url", "oauth_client_redirect_url_id", "created_at", "updated_at"}
	oauthClientRedirectURLColumnsWithoutDefault = []string{"oauth_client_id", "redirect_url", "oauth_client_redirect_url_id"}
	oauthClientRedirectURLColumnsWithDefault    = []string{"created_at", "updated_at"}
	oauthClientRedirectURLPrimaryKeyColumns     = []string{"oauth_client_redirect_url_id"}
	oauthClientRedirectURLGeneratedColumns      = []string{}
)

type (
	// OauthClientRedirectURLSlice is an alias for a slice of pointers to OauthClientRedirectURL.
	// This should almost always be used instead of []OauthClientRedirectURL.
	OauthClientRedirectURLSlice []*OauthClientRedirectURL
	// OauthClientRedirectURLHook is the signature for custom OauthClientRedirectURL hook methods
	OauthClientRedirectURLHook func(context.Context, boil.ContextExecutor, *OauthClientRedirectURL) error

	oauthClientRedirectURLQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	oauthClientRedirectURLType                 = reflect.TypeOf(&OauthClientRedirectURL{})
	oauthClientRedirectURLMapping              = queries.MakeStructMapping(oauthClientRedirectURLType)
	oauthClientRedirectURLPrimaryKeyMapping, _ = queries.BindMapping(oauthClientRedirectURLType, oauthClientRedirectURLMapping, oauthClientRedirectURLPrimaryKeyColumns)
	oauthClientRedirectURLInsertCacheMut       sync.RWMutex
	oauthClientRedirectURLInsertCache          = make(map[string]insertCache)
	oauthClientRedirectURLUpdateCacheMut       sync.RWMutex
	oauthClientRedirectURLUpdateCache          = make(map[string]updateCache)
	oauthClientRedirectURLUpsertCacheMut       sync.RWMutex
	oauthClientRedirectURLUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var oauthClientRedirectURLAfterSelectHooks []OauthClientRedirectURLHook

var oauthClientRedirectURLBeforeInsertHooks []OauthClientRedirectURLHook
var oauthClientRedirectURLAfterInsertHooks []OauthClientRedirectURLHook

var oauthClientRedirectURLBeforeUpdateHooks []OauthClientRedirectURLHook
var oauthClientRedirectURLAfterUpdateHooks []OauthClientRedirectURLHook

var oauthClientRedirectURLBeforeDeleteHooks []OauthClientRedirectURLHook
var oauthClientRedirectURLAfterDeleteHooks []OauthClientRedirectURLHook

var oauthClientRedirectURLBeforeUpsertHooks []OauthClientRedirectURLHook
var oauthClientRedirectURLAfterUpsertHooks []OauthClientRedirectURLHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OauthClientRedirectURL) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oauthClientRedirectURLAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OauthClientRedirectURL) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oauthClientRedirectURLBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OauthClientRedirectURL) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oauthClientRedirectURLAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OauthClientRedirectURL) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oauthClientRedirectURLBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OauthClientRedirectURL) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oauthClientRedirectURLAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OauthClientRedirectURL) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oauthClientRedirectURLBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OauthClientRedirectURL) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oauthClientRedirectURLAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OauthClientRedirectURL) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oauthClientRedirectURLBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OauthClientRedirectURL) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oauthClientRedirectURLAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOauthClientRedirectURLHook registers your hook function for all future operations.
func AddOauthClientRedirectURLHook(hookPoint boil.HookPoint, oauthClientRedirectURLHook OauthClientRedirectURLHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		oauthClientRedirectURLAfterSelectHooks = append(oauthClientRedirectURLAfterSelectHooks, oauthClientRedirectURLHook)
	case boil.BeforeInsertHook:
		oauthClientRedirectURLBeforeInsertHooks = append(oauthClientRedirectURLBeforeInsertHooks, oauthClientRedirectURLHook)
	case boil.AfterInsertHook:
		oauthClientRedirectURLAfterInsertHooks = append(oauthClientRedirectURLAfterInsertHooks, oauthClientRedirectURLHook)
	case boil.BeforeUpdateHook:
		oauthClientRedirectURLBeforeUpdateHooks = append(oauthClientRedirectURLBeforeUpdateHooks, oauthClientRedirectURLHook)
	case boil.AfterUpdateHook:
		oauthClientRedirectURLAfterUpdateHooks = append(oauthClientRedirectURLAfterUpdateHooks, oauthClientRedirectURLHook)
	case boil.BeforeDeleteHook:
		oauthClientRedirectURLBeforeDeleteHooks = append(oauthClientRedirectURLBeforeDeleteHooks, oauthClientRedirectURLHook)
	case boil.AfterDeleteHook:
		oauthClientRedirectURLAfterDeleteHooks = append(oauthClientRedirectURLAfterDeleteHooks, oauthClientRedirectURLHook)
	case boil.BeforeUpsertHook:
		oauthClientRedirectURLBeforeUpsertHooks = append(oauthClientRedirectURLBeforeUpsertHooks, oauthClientRedirectURLHook)
	case boil.AfterUpsertHook:
		oauthClientRedirectURLAfterUpsertHooks = append(oauthClientRedirectURLAfterUpsertHooks, oauthClientRedirectURLHook)
	}
}

// One returns a single oauthClientRedirectURL record from the query.
func (q oauthClientRedirectURLQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OauthClientRedirectURL, error) {
	o := &OauthClientRedirectURL{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: failed to execute a one query for oauth_client_redirect_urls")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all OauthClientRedirectURL records from the query.
func (q oauthClientRedirectURLQuery) All(ctx context.Context, exec boil.ContextExecutor) (OauthClientRedirectURLSlice, error) {
	var o []*OauthClientRedirectURL

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodels: failed to assign all query results to OauthClientRedirectURL slice")
	}

	if len(oauthClientRedirectURLAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all OauthClientRedirectURL records in the query.
func (q oauthClientRedirectURLQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to count oauth_client_redirect_urls rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q oauthClientRedirectURLQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: failed to check if oauth_client_redirect_urls exists")
	}

	return count > 0, nil
}

// OauthClient pointed to by the foreign key.
func (o *OauthClientRedirectURL) OauthClient(mods ...qm.QueryMod) oauthClientQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`oauth_client_id` = ?", o.OauthClientID),
	}

	queryMods = append(queryMods, mods...)

	query := OauthClients(queryMods...)
	queries.SetFrom(query.Query, "`oauth_clients`")

	return query
}

// LoadOauthClient allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (oauthClientRedirectURLL) LoadOauthClient(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOauthClientRedirectURL interface{}, mods queries.Applicator) error {
	var slice []*OauthClientRedirectURL
	var object *OauthClientRedirectURL

	if singular {
		object = maybeOauthClientRedirectURL.(*OauthClientRedirectURL)
	} else {
		slice = *maybeOauthClientRedirectURL.(*[]*OauthClientRedirectURL)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &oauthClientRedirectURLR{}
		}
		args = append(args, object.OauthClientID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &oauthClientRedirectURLR{}
			}

			for _, a := range args {
				if a == obj.OauthClientID {
					continue Outer
				}
			}

			args = append(args, obj.OauthClientID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`oauth_clients`),
		qm.WhereIn(`oauth_clients.oauth_client_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load OauthClient")
	}

	var resultSlice []*OauthClient
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice OauthClient")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for oauth_clients")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for oauth_clients")
	}

	if len(oauthClientRedirectURLAfterSelectHooks) != 0 {
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
		object.R.OauthClient = foreign
		if foreign.R == nil {
			foreign.R = &oauthClientR{}
		}
		foreign.R.OauthClientRedirectUrls = append(foreign.R.OauthClientRedirectUrls, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.OauthClientID == foreign.OauthClientID {
				local.R.OauthClient = foreign
				if foreign.R == nil {
					foreign.R = &oauthClientR{}
				}
				foreign.R.OauthClientRedirectUrls = append(foreign.R.OauthClientRedirectUrls, local)
				break
			}
		}
	}

	return nil
}

// SetOauthClient of the oauthClientRedirectURL to the related item.
// Sets o.R.OauthClient to related.
// Adds o to related.R.OauthClientRedirectUrls.
func (o *OauthClientRedirectURL) SetOauthClient(ctx context.Context, exec boil.ContextExecutor, insert bool, related *OauthClient) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `oauth_client_redirect_urls` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"oauth_client_id"}),
		strmangle.WhereClause("`", "`", 0, oauthClientRedirectURLPrimaryKeyColumns),
	)
	values := []interface{}{related.OauthClientID, o.OauthClientRedirectURLID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OauthClientID = related.OauthClientID
	if o.R == nil {
		o.R = &oauthClientRedirectURLR{
			OauthClient: related,
		}
	} else {
		o.R.OauthClient = related
	}

	if related.R == nil {
		related.R = &oauthClientR{
			OauthClientRedirectUrls: OauthClientRedirectURLSlice{o},
		}
	} else {
		related.R.OauthClientRedirectUrls = append(related.R.OauthClientRedirectUrls, o)
	}

	return nil
}

// OauthClientRedirectUrls retrieves all the records using an executor.
func OauthClientRedirectUrls(mods ...qm.QueryMod) oauthClientRedirectURLQuery {
	mods = append(mods, qm.From("`oauth_client_redirect_urls`"))
	return oauthClientRedirectURLQuery{NewQuery(mods...)}
}

// FindOauthClientRedirectURL retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOauthClientRedirectURL(ctx context.Context, exec boil.ContextExecutor, oauthClientRedirectURLID int, selectCols ...string) (*OauthClientRedirectURL, error) {
	oauthClientRedirectURLObj := &OauthClientRedirectURL{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `oauth_client_redirect_urls` where `oauth_client_redirect_url_id`=?", sel,
	)

	q := queries.Raw(query, oauthClientRedirectURLID)

	err := q.Bind(ctx, exec, oauthClientRedirectURLObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: unable to select from oauth_client_redirect_urls")
	}

	if err = oauthClientRedirectURLObj.doAfterSelectHooks(ctx, exec); err != nil {
		return oauthClientRedirectURLObj, err
	}

	return oauthClientRedirectURLObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OauthClientRedirectURL) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no oauth_client_redirect_urls provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(oauthClientRedirectURLColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	oauthClientRedirectURLInsertCacheMut.RLock()
	cache, cached := oauthClientRedirectURLInsertCache[key]
	oauthClientRedirectURLInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			oauthClientRedirectURLAllColumns,
			oauthClientRedirectURLColumnsWithDefault,
			oauthClientRedirectURLColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(oauthClientRedirectURLType, oauthClientRedirectURLMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(oauthClientRedirectURLType, oauthClientRedirectURLMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `oauth_client_redirect_urls` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `oauth_client_redirect_urls` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `oauth_client_redirect_urls` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, oauthClientRedirectURLPrimaryKeyColumns))
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
		return errors.Wrap(err, "dbmodels: unable to insert into oauth_client_redirect_urls")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.OauthClientRedirectURLID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to populate default values for oauth_client_redirect_urls")
	}

CacheNoHooks:
	if !cached {
		oauthClientRedirectURLInsertCacheMut.Lock()
		oauthClientRedirectURLInsertCache[key] = cache
		oauthClientRedirectURLInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the OauthClientRedirectURL.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OauthClientRedirectURL) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	oauthClientRedirectURLUpdateCacheMut.RLock()
	cache, cached := oauthClientRedirectURLUpdateCache[key]
	oauthClientRedirectURLUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			oauthClientRedirectURLAllColumns,
			oauthClientRedirectURLPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodels: unable to update oauth_client_redirect_urls, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `oauth_client_redirect_urls` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, oauthClientRedirectURLPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(oauthClientRedirectURLType, oauthClientRedirectURLMapping, append(wl, oauthClientRedirectURLPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "dbmodels: unable to update oauth_client_redirect_urls row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by update for oauth_client_redirect_urls")
	}

	if !cached {
		oauthClientRedirectURLUpdateCacheMut.Lock()
		oauthClientRedirectURLUpdateCache[key] = cache
		oauthClientRedirectURLUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q oauthClientRedirectURLQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all for oauth_client_redirect_urls")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected for oauth_client_redirect_urls")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OauthClientRedirectURLSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("dbmodels: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oauthClientRedirectURLPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `oauth_client_redirect_urls` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, oauthClientRedirectURLPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all in oauthClientRedirectURL slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected all in update all oauthClientRedirectURL")
	}
	return rowsAff, nil
}

var mySQLOauthClientRedirectURLUniqueColumns = []string{
	"oauth_client_redirect_url_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OauthClientRedirectURL) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no oauth_client_redirect_urls provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(oauthClientRedirectURLColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLOauthClientRedirectURLUniqueColumns, o)

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

	oauthClientRedirectURLUpsertCacheMut.RLock()
	cache, cached := oauthClientRedirectURLUpsertCache[key]
	oauthClientRedirectURLUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			oauthClientRedirectURLAllColumns,
			oauthClientRedirectURLColumnsWithDefault,
			oauthClientRedirectURLColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			oauthClientRedirectURLAllColumns,
			oauthClientRedirectURLPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("dbmodels: unable to upsert oauth_client_redirect_urls, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`oauth_client_redirect_urls`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `oauth_client_redirect_urls` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(oauthClientRedirectURLType, oauthClientRedirectURLMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(oauthClientRedirectURLType, oauthClientRedirectURLMapping, ret)
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
		return errors.Wrap(err, "dbmodels: unable to upsert for oauth_client_redirect_urls")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(oauthClientRedirectURLType, oauthClientRedirectURLMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to retrieve unique values for oauth_client_redirect_urls")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to populate default values for oauth_client_redirect_urls")
	}

CacheNoHooks:
	if !cached {
		oauthClientRedirectURLUpsertCacheMut.Lock()
		oauthClientRedirectURLUpsertCache[key] = cache
		oauthClientRedirectURLUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single OauthClientRedirectURL record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OauthClientRedirectURL) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodels: no OauthClientRedirectURL provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), oauthClientRedirectURLPrimaryKeyMapping)
	sql := "DELETE FROM `oauth_client_redirect_urls` WHERE `oauth_client_redirect_url_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete from oauth_client_redirect_urls")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by delete for oauth_client_redirect_urls")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q oauthClientRedirectURLQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodels: no oauthClientRedirectURLQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from oauth_client_redirect_urls")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for oauth_client_redirect_urls")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OauthClientRedirectURLSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(oauthClientRedirectURLBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oauthClientRedirectURLPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `oauth_client_redirect_urls` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, oauthClientRedirectURLPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from oauthClientRedirectURL slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for oauth_client_redirect_urls")
	}

	if len(oauthClientRedirectURLAfterDeleteHooks) != 0 {
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
func (o *OauthClientRedirectURL) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOauthClientRedirectURL(ctx, exec, o.OauthClientRedirectURLID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OauthClientRedirectURLSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OauthClientRedirectURLSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oauthClientRedirectURLPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `oauth_client_redirect_urls`.* FROM `oauth_client_redirect_urls` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, oauthClientRedirectURLPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to reload all in OauthClientRedirectURLSlice")
	}

	*o = slice

	return nil
}

// OauthClientRedirectURLExists checks if the OauthClientRedirectURL row exists.
func OauthClientRedirectURLExists(ctx context.Context, exec boil.ContextExecutor, oauthClientRedirectURLID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `oauth_client_redirect_urls` where `oauth_client_redirect_url_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, oauthClientRedirectURLID)
	}
	row := exec.QueryRowContext(ctx, sql, oauthClientRedirectURLID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: unable to check if oauth_client_redirect_urls exists")
	}

	return exists, nil
}
