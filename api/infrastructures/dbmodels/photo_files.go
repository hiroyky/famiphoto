// Code generated by SQLBoiler 4.18.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// PhotoFile is an object representing the database table.
type PhotoFile struct {
	PhotoFileID int       `boil:"photo_file_id" json:"photo_file_id" toml:"photo_file_id" yaml:"photo_file_id"`
	PhotoID     int       `boil:"photo_id" json:"photo_id" toml:"photo_id" yaml:"photo_id"`
	FileType    string    `boil:"file_type" json:"file_type" toml:"file_type" yaml:"file_type"`
	FilePath    string    `boil:"file_path" json:"file_path" toml:"file_path" yaml:"file_path"`
	ImportedAt  time.Time `boil:"imported_at" json:"imported_at" toml:"imported_at" yaml:"imported_at"`
	FileHash    string    `boil:"file_hash" json:"file_hash" toml:"file_hash" yaml:"file_hash"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *photoFileR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L photoFileL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PhotoFileColumns = struct {
	PhotoFileID string
	PhotoID     string
	FileType    string
	FilePath    string
	ImportedAt  string
	FileHash    string
	CreatedAt   string
	UpdatedAt   string
}{
	PhotoFileID: "photo_file_id",
	PhotoID:     "photo_id",
	FileType:    "file_type",
	FilePath:    "file_path",
	ImportedAt:  "imported_at",
	FileHash:    "file_hash",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

var PhotoFileTableColumns = struct {
	PhotoFileID string
	PhotoID     string
	FileType    string
	FilePath    string
	ImportedAt  string
	FileHash    string
	CreatedAt   string
	UpdatedAt   string
}{
	PhotoFileID: "photo_files.photo_file_id",
	PhotoID:     "photo_files.photo_id",
	FileType:    "photo_files.file_type",
	FilePath:    "photo_files.file_path",
	ImportedAt:  "photo_files.imported_at",
	FileHash:    "photo_files.file_hash",
	CreatedAt:   "photo_files.created_at",
	UpdatedAt:   "photo_files.updated_at",
}

// Generated where

var PhotoFileWhere = struct {
	PhotoFileID whereHelperint
	PhotoID     whereHelperint
	FileType    whereHelperstring
	FilePath    whereHelperstring
	ImportedAt  whereHelpertime_Time
	FileHash    whereHelperstring
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpertime_Time
}{
	PhotoFileID: whereHelperint{field: "`photo_files`.`photo_file_id`"},
	PhotoID:     whereHelperint{field: "`photo_files`.`photo_id`"},
	FileType:    whereHelperstring{field: "`photo_files`.`file_type`"},
	FilePath:    whereHelperstring{field: "`photo_files`.`file_path`"},
	ImportedAt:  whereHelpertime_Time{field: "`photo_files`.`imported_at`"},
	FileHash:    whereHelperstring{field: "`photo_files`.`file_hash`"},
	CreatedAt:   whereHelpertime_Time{field: "`photo_files`.`created_at`"},
	UpdatedAt:   whereHelpertime_Time{field: "`photo_files`.`updated_at`"},
}

// PhotoFileRels is where relationship names are stored.
var PhotoFileRels = struct {
	Photo string
}{
	Photo: "Photo",
}

// photoFileR is where relationships are stored.
type photoFileR struct {
	Photo *Photo `boil:"Photo" json:"Photo" toml:"Photo" yaml:"Photo"`
}

// NewStruct creates a new relationship struct
func (*photoFileR) NewStruct() *photoFileR {
	return &photoFileR{}
}

func (r *photoFileR) GetPhoto() *Photo {
	if r == nil {
		return nil
	}
	return r.Photo
}

// photoFileL is where Load methods for each relationship are stored.
type photoFileL struct{}

var (
	photoFileAllColumns            = []string{"photo_file_id", "photo_id", "file_type", "file_path", "imported_at", "file_hash", "created_at", "updated_at"}
	photoFileColumnsWithoutDefault = []string{"photo_id", "file_type", "file_path", "imported_at", "file_hash"}
	photoFileColumnsWithDefault    = []string{"photo_file_id", "created_at", "updated_at"}
	photoFilePrimaryKeyColumns     = []string{"photo_file_id"}
	photoFileGeneratedColumns      = []string{}
)

type (
	// PhotoFileSlice is an alias for a slice of pointers to PhotoFile.
	// This should almost always be used instead of []PhotoFile.
	PhotoFileSlice []*PhotoFile
	// PhotoFileHook is the signature for custom PhotoFile hook methods
	PhotoFileHook func(context.Context, boil.ContextExecutor, *PhotoFile) error

	photoFileQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	photoFileType                 = reflect.TypeOf(&PhotoFile{})
	photoFileMapping              = queries.MakeStructMapping(photoFileType)
	photoFilePrimaryKeyMapping, _ = queries.BindMapping(photoFileType, photoFileMapping, photoFilePrimaryKeyColumns)
	photoFileInsertCacheMut       sync.RWMutex
	photoFileInsertCache          = make(map[string]insertCache)
	photoFileUpdateCacheMut       sync.RWMutex
	photoFileUpdateCache          = make(map[string]updateCache)
	photoFileUpsertCacheMut       sync.RWMutex
	photoFileUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var photoFileAfterSelectMu sync.Mutex
var photoFileAfterSelectHooks []PhotoFileHook

var photoFileBeforeInsertMu sync.Mutex
var photoFileBeforeInsertHooks []PhotoFileHook
var photoFileAfterInsertMu sync.Mutex
var photoFileAfterInsertHooks []PhotoFileHook

var photoFileBeforeUpdateMu sync.Mutex
var photoFileBeforeUpdateHooks []PhotoFileHook
var photoFileAfterUpdateMu sync.Mutex
var photoFileAfterUpdateHooks []PhotoFileHook

var photoFileBeforeDeleteMu sync.Mutex
var photoFileBeforeDeleteHooks []PhotoFileHook
var photoFileAfterDeleteMu sync.Mutex
var photoFileAfterDeleteHooks []PhotoFileHook

var photoFileBeforeUpsertMu sync.Mutex
var photoFileBeforeUpsertHooks []PhotoFileHook
var photoFileAfterUpsertMu sync.Mutex
var photoFileAfterUpsertHooks []PhotoFileHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PhotoFile) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photoFileAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PhotoFile) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photoFileBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PhotoFile) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photoFileAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PhotoFile) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photoFileBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PhotoFile) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photoFileAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PhotoFile) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photoFileBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PhotoFile) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photoFileAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PhotoFile) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photoFileBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PhotoFile) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photoFileAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPhotoFileHook registers your hook function for all future operations.
func AddPhotoFileHook(hookPoint boil.HookPoint, photoFileHook PhotoFileHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		photoFileAfterSelectMu.Lock()
		photoFileAfterSelectHooks = append(photoFileAfterSelectHooks, photoFileHook)
		photoFileAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		photoFileBeforeInsertMu.Lock()
		photoFileBeforeInsertHooks = append(photoFileBeforeInsertHooks, photoFileHook)
		photoFileBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		photoFileAfterInsertMu.Lock()
		photoFileAfterInsertHooks = append(photoFileAfterInsertHooks, photoFileHook)
		photoFileAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		photoFileBeforeUpdateMu.Lock()
		photoFileBeforeUpdateHooks = append(photoFileBeforeUpdateHooks, photoFileHook)
		photoFileBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		photoFileAfterUpdateMu.Lock()
		photoFileAfterUpdateHooks = append(photoFileAfterUpdateHooks, photoFileHook)
		photoFileAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		photoFileBeforeDeleteMu.Lock()
		photoFileBeforeDeleteHooks = append(photoFileBeforeDeleteHooks, photoFileHook)
		photoFileBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		photoFileAfterDeleteMu.Lock()
		photoFileAfterDeleteHooks = append(photoFileAfterDeleteHooks, photoFileHook)
		photoFileAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		photoFileBeforeUpsertMu.Lock()
		photoFileBeforeUpsertHooks = append(photoFileBeforeUpsertHooks, photoFileHook)
		photoFileBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		photoFileAfterUpsertMu.Lock()
		photoFileAfterUpsertHooks = append(photoFileAfterUpsertHooks, photoFileHook)
		photoFileAfterUpsertMu.Unlock()
	}
}

// One returns a single photoFile record from the query.
func (q photoFileQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PhotoFile, error) {
	o := &PhotoFile{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: failed to execute a one query for photo_files")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PhotoFile records from the query.
func (q photoFileQuery) All(ctx context.Context, exec boil.ContextExecutor) (PhotoFileSlice, error) {
	var o []*PhotoFile

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodels: failed to assign all query results to PhotoFile slice")
	}

	if len(photoFileAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PhotoFile records in the query.
func (q photoFileQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to count photo_files rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q photoFileQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: failed to check if photo_files exists")
	}

	return count > 0, nil
}

// Photo pointed to by the foreign key.
func (o *PhotoFile) Photo(mods ...qm.QueryMod) photoQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`photo_id` = ?", o.PhotoID),
	}

	queryMods = append(queryMods, mods...)

	return Photos(queryMods...)
}

// LoadPhoto allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (photoFileL) LoadPhoto(ctx context.Context, e boil.ContextExecutor, singular bool, maybePhotoFile interface{}, mods queries.Applicator) error {
	var slice []*PhotoFile
	var object *PhotoFile

	if singular {
		var ok bool
		object, ok = maybePhotoFile.(*PhotoFile)
		if !ok {
			object = new(PhotoFile)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePhotoFile)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePhotoFile))
			}
		}
	} else {
		s, ok := maybePhotoFile.(*[]*PhotoFile)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePhotoFile)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePhotoFile))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &photoFileR{}
		}
		args[object.PhotoID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &photoFileR{}
			}

			args[obj.PhotoID] = struct{}{}

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
		qm.From(`photos`),
		qm.WhereIn(`photos.photo_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Photo")
	}

	var resultSlice []*Photo
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Photo")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for photos")
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

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Photo = foreign
		if foreign.R == nil {
			foreign.R = &photoR{}
		}
		foreign.R.PhotoFiles = append(foreign.R.PhotoFiles, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PhotoID == foreign.PhotoID {
				local.R.Photo = foreign
				if foreign.R == nil {
					foreign.R = &photoR{}
				}
				foreign.R.PhotoFiles = append(foreign.R.PhotoFiles, local)
				break
			}
		}
	}

	return nil
}

// SetPhoto of the photoFile to the related item.
// Sets o.R.Photo to related.
// Adds o to related.R.PhotoFiles.
func (o *PhotoFile) SetPhoto(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Photo) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `photo_files` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"photo_id"}),
		strmangle.WhereClause("`", "`", 0, photoFilePrimaryKeyColumns),
	)
	values := []interface{}{related.PhotoID, o.PhotoFileID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PhotoID = related.PhotoID
	if o.R == nil {
		o.R = &photoFileR{
			Photo: related,
		}
	} else {
		o.R.Photo = related
	}

	if related.R == nil {
		related.R = &photoR{
			PhotoFiles: PhotoFileSlice{o},
		}
	} else {
		related.R.PhotoFiles = append(related.R.PhotoFiles, o)
	}

	return nil
}

// PhotoFiles retrieves all the records using an executor.
func PhotoFiles(mods ...qm.QueryMod) photoFileQuery {
	mods = append(mods, qm.From("`photo_files`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`photo_files`.*"})
	}

	return photoFileQuery{q}
}

// FindPhotoFile retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPhotoFile(ctx context.Context, exec boil.ContextExecutor, photoFileID int, selectCols ...string) (*PhotoFile, error) {
	photoFileObj := &PhotoFile{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `photo_files` where `photo_file_id`=?", sel,
	)

	q := queries.Raw(query, photoFileID)

	err := q.Bind(ctx, exec, photoFileObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: unable to select from photo_files")
	}

	if err = photoFileObj.doAfterSelectHooks(ctx, exec); err != nil {
		return photoFileObj, err
	}

	return photoFileObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PhotoFile) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no photo_files provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(photoFileColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	photoFileInsertCacheMut.RLock()
	cache, cached := photoFileInsertCache[key]
	photoFileInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			photoFileAllColumns,
			photoFileColumnsWithDefault,
			photoFileColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(photoFileType, photoFileMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(photoFileType, photoFileMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `photo_files` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `photo_files` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `photo_files` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, photoFilePrimaryKeyColumns))
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
		return errors.Wrap(err, "dbmodels: unable to insert into photo_files")
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

	o.PhotoFileID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == photoFileMapping["photo_file_id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.PhotoFileID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to populate default values for photo_files")
	}

CacheNoHooks:
	if !cached {
		photoFileInsertCacheMut.Lock()
		photoFileInsertCache[key] = cache
		photoFileInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PhotoFile.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PhotoFile) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	photoFileUpdateCacheMut.RLock()
	cache, cached := photoFileUpdateCache[key]
	photoFileUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			photoFileAllColumns,
			photoFilePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodels: unable to update photo_files, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `photo_files` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, photoFilePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(photoFileType, photoFileMapping, append(wl, photoFilePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "dbmodels: unable to update photo_files row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by update for photo_files")
	}

	if !cached {
		photoFileUpdateCacheMut.Lock()
		photoFileUpdateCache[key] = cache
		photoFileUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q photoFileQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all for photo_files")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected for photo_files")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PhotoFileSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), photoFilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `photo_files` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, photoFilePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all in photoFile slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected all in update all photoFile")
	}
	return rowsAff, nil
}

var mySQLPhotoFileUniqueColumns = []string{
	"photo_file_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PhotoFile) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no photo_files provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(photoFileColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLPhotoFileUniqueColumns, o)

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

	photoFileUpsertCacheMut.RLock()
	cache, cached := photoFileUpsertCache[key]
	photoFileUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			photoFileAllColumns,
			photoFileColumnsWithDefault,
			photoFileColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			photoFileAllColumns,
			photoFilePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("dbmodels: unable to upsert photo_files, could not build update column list")
		}

		ret := strmangle.SetComplement(photoFileAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`photo_files`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `photo_files` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(photoFileType, photoFileMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(photoFileType, photoFileMapping, ret)
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
		return errors.Wrap(err, "dbmodels: unable to upsert for photo_files")
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

	o.PhotoFileID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == photoFileMapping["photo_file_id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(photoFileType, photoFileMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to retrieve unique values for photo_files")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to populate default values for photo_files")
	}

CacheNoHooks:
	if !cached {
		photoFileUpsertCacheMut.Lock()
		photoFileUpsertCache[key] = cache
		photoFileUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PhotoFile record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PhotoFile) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodels: no PhotoFile provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), photoFilePrimaryKeyMapping)
	sql := "DELETE FROM `photo_files` WHERE `photo_file_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete from photo_files")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by delete for photo_files")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q photoFileQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodels: no photoFileQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from photo_files")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for photo_files")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PhotoFileSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(photoFileBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), photoFilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `photo_files` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, photoFilePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from photoFile slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for photo_files")
	}

	if len(photoFileAfterDeleteHooks) != 0 {
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
func (o *PhotoFile) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPhotoFile(ctx, exec, o.PhotoFileID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PhotoFileSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PhotoFileSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), photoFilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `photo_files`.* FROM `photo_files` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, photoFilePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to reload all in PhotoFileSlice")
	}

	*o = slice

	return nil
}

// PhotoFileExists checks if the PhotoFile row exists.
func PhotoFileExists(ctx context.Context, exec boil.ContextExecutor, photoFileID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `photo_files` where `photo_file_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, photoFileID)
	}
	row := exec.QueryRowContext(ctx, sql, photoFileID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: unable to check if photo_files exists")
	}

	return exists, nil
}

// Exists checks if the PhotoFile row exists.
func (o *PhotoFile) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PhotoFileExists(ctx, exec, o.PhotoFileID)
}
