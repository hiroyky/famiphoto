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

// Photo is an object representing the database table.
type Photo struct {
	PhotoID    int       `boil:"photo_id" json:"photo_id" toml:"photo_id" yaml:"photo_id"`
	Name       string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	FilePath   string    `boil:"file_path" json:"file_path" toml:"file_path" yaml:"file_path"`
	ImportedAt time.Time `boil:"imported_at" json:"imported_at" toml:"imported_at" yaml:"imported_at"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	GroupID    string    `boil:"group_id" json:"group_id" toml:"group_id" yaml:"group_id"`
	OwnerID    string    `boil:"owner_id" json:"owner_id" toml:"owner_id" yaml:"owner_id"`

	R *photoR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L photoL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PhotoColumns = struct {
	PhotoID    string
	Name       string
	FilePath   string
	ImportedAt string
	CreatedAt  string
	UpdatedAt  string
	GroupID    string
	OwnerID    string
}{
	PhotoID:    "photo_id",
	Name:       "name",
	FilePath:   "file_path",
	ImportedAt: "imported_at",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	GroupID:    "group_id",
	OwnerID:    "owner_id",
}

var PhotoTableColumns = struct {
	PhotoID    string
	Name       string
	FilePath   string
	ImportedAt string
	CreatedAt  string
	UpdatedAt  string
	GroupID    string
	OwnerID    string
}{
	PhotoID:    "photos.photo_id",
	Name:       "photos.name",
	FilePath:   "photos.file_path",
	ImportedAt: "photos.imported_at",
	CreatedAt:  "photos.created_at",
	UpdatedAt:  "photos.updated_at",
	GroupID:    "photos.group_id",
	OwnerID:    "photos.owner_id",
}

// Generated where

var PhotoWhere = struct {
	PhotoID    whereHelperint
	Name       whereHelperstring
	FilePath   whereHelperstring
	ImportedAt whereHelpertime_Time
	CreatedAt  whereHelpertime_Time
	UpdatedAt  whereHelpertime_Time
	GroupID    whereHelperstring
	OwnerID    whereHelperstring
}{
	PhotoID:    whereHelperint{field: "`photos`.`photo_id`"},
	Name:       whereHelperstring{field: "`photos`.`name`"},
	FilePath:   whereHelperstring{field: "`photos`.`file_path`"},
	ImportedAt: whereHelpertime_Time{field: "`photos`.`imported_at`"},
	CreatedAt:  whereHelpertime_Time{field: "`photos`.`created_at`"},
	UpdatedAt:  whereHelpertime_Time{field: "`photos`.`updated_at`"},
	GroupID:    whereHelperstring{field: "`photos`.`group_id`"},
	OwnerID:    whereHelperstring{field: "`photos`.`owner_id`"},
}

// PhotoRels is where relationship names are stored.
var PhotoRels = struct {
	Group     string
	Owner     string
	PhotoExif string
}{
	Group:     "Group",
	Owner:     "Owner",
	PhotoExif: "PhotoExif",
}

// photoR is where relationships are stored.
type photoR struct {
	Group     *Group     `boil:"Group" json:"Group" toml:"Group" yaml:"Group"`
	Owner     *User      `boil:"Owner" json:"Owner" toml:"Owner" yaml:"Owner"`
	PhotoExif *PhotoExif `boil:"PhotoExif" json:"PhotoExif" toml:"PhotoExif" yaml:"PhotoExif"`
}

// NewStruct creates a new relationship struct
func (*photoR) NewStruct() *photoR {
	return &photoR{}
}

// photoL is where Load methods for each relationship are stored.
type photoL struct{}

var (
	photoAllColumns            = []string{"photo_id", "name", "file_path", "imported_at", "created_at", "updated_at", "group_id", "owner_id"}
	photoColumnsWithoutDefault = []string{"name", "file_path", "imported_at", "created_at", "updated_at", "group_id", "owner_id"}
	photoColumnsWithDefault    = []string{"photo_id"}
	photoPrimaryKeyColumns     = []string{"photo_id"}
	photoGeneratedColumns      = []string{}
)

type (
	// PhotoSlice is an alias for a slice of pointers to Photo.
	// This should almost always be used instead of []Photo.
	PhotoSlice []*Photo
	// PhotoHook is the signature for custom Photo hook methods
	PhotoHook func(context.Context, boil.ContextExecutor, *Photo) error

	photoQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	photoType                 = reflect.TypeOf(&Photo{})
	photoMapping              = queries.MakeStructMapping(photoType)
	photoPrimaryKeyMapping, _ = queries.BindMapping(photoType, photoMapping, photoPrimaryKeyColumns)
	photoInsertCacheMut       sync.RWMutex
	photoInsertCache          = make(map[string]insertCache)
	photoUpdateCacheMut       sync.RWMutex
	photoUpdateCache          = make(map[string]updateCache)
	photoUpsertCacheMut       sync.RWMutex
	photoUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var photoAfterSelectHooks []PhotoHook

var photoBeforeInsertHooks []PhotoHook
var photoAfterInsertHooks []PhotoHook

var photoBeforeUpdateHooks []PhotoHook
var photoAfterUpdateHooks []PhotoHook

var photoBeforeDeleteHooks []PhotoHook
var photoAfterDeleteHooks []PhotoHook

var photoBeforeUpsertHooks []PhotoHook
var photoAfterUpsertHooks []PhotoHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Photo) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photoAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Photo) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photoBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Photo) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photoAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Photo) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photoBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Photo) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photoAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Photo) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photoBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Photo) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photoAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Photo) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photoBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Photo) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range photoAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPhotoHook registers your hook function for all future operations.
func AddPhotoHook(hookPoint boil.HookPoint, photoHook PhotoHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		photoAfterSelectHooks = append(photoAfterSelectHooks, photoHook)
	case boil.BeforeInsertHook:
		photoBeforeInsertHooks = append(photoBeforeInsertHooks, photoHook)
	case boil.AfterInsertHook:
		photoAfterInsertHooks = append(photoAfterInsertHooks, photoHook)
	case boil.BeforeUpdateHook:
		photoBeforeUpdateHooks = append(photoBeforeUpdateHooks, photoHook)
	case boil.AfterUpdateHook:
		photoAfterUpdateHooks = append(photoAfterUpdateHooks, photoHook)
	case boil.BeforeDeleteHook:
		photoBeforeDeleteHooks = append(photoBeforeDeleteHooks, photoHook)
	case boil.AfterDeleteHook:
		photoAfterDeleteHooks = append(photoAfterDeleteHooks, photoHook)
	case boil.BeforeUpsertHook:
		photoBeforeUpsertHooks = append(photoBeforeUpsertHooks, photoHook)
	case boil.AfterUpsertHook:
		photoAfterUpsertHooks = append(photoAfterUpsertHooks, photoHook)
	}
}

// One returns a single photo record from the query.
func (q photoQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Photo, error) {
	o := &Photo{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: failed to execute a one query for photos")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Photo records from the query.
func (q photoQuery) All(ctx context.Context, exec boil.ContextExecutor) (PhotoSlice, error) {
	var o []*Photo

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodels: failed to assign all query results to Photo slice")
	}

	if len(photoAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Photo records in the query.
func (q photoQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to count photos rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q photoQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: failed to check if photos exists")
	}

	return count > 0, nil
}

// Group pointed to by the foreign key.
func (o *Photo) Group(mods ...qm.QueryMod) groupQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`group_id` = ?", o.GroupID),
	}

	queryMods = append(queryMods, mods...)

	query := Groups(queryMods...)
	queries.SetFrom(query.Query, "`groups`")

	return query
}

// Owner pointed to by the foreign key.
func (o *Photo) Owner(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`user_id` = ?", o.OwnerID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "`users`")

	return query
}

// PhotoExif pointed to by the foreign key.
func (o *Photo) PhotoExif(mods ...qm.QueryMod) photoExifQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`photo_id` = ?", o.PhotoID),
	}

	queryMods = append(queryMods, mods...)

	query := PhotoExifs(queryMods...)
	queries.SetFrom(query.Query, "`photo_exif`")

	return query
}

// LoadGroup allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (photoL) LoadGroup(ctx context.Context, e boil.ContextExecutor, singular bool, maybePhoto interface{}, mods queries.Applicator) error {
	var slice []*Photo
	var object *Photo

	if singular {
		object = maybePhoto.(*Photo)
	} else {
		slice = *maybePhoto.(*[]*Photo)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &photoR{}
		}
		args = append(args, object.GroupID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &photoR{}
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
		qm.From(`groups`),
		qm.WhereIn(`groups.group_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Group")
	}

	var resultSlice []*Group
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Group")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for groups")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for groups")
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
		object.R.Group = foreign
		if foreign.R == nil {
			foreign.R = &groupR{}
		}
		foreign.R.Photos = append(foreign.R.Photos, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.GroupID == foreign.GroupID {
				local.R.Group = foreign
				if foreign.R == nil {
					foreign.R = &groupR{}
				}
				foreign.R.Photos = append(foreign.R.Photos, local)
				break
			}
		}
	}

	return nil
}

// LoadOwner allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (photoL) LoadOwner(ctx context.Context, e boil.ContextExecutor, singular bool, maybePhoto interface{}, mods queries.Applicator) error {
	var slice []*Photo
	var object *Photo

	if singular {
		object = maybePhoto.(*Photo)
	} else {
		slice = *maybePhoto.(*[]*Photo)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &photoR{}
		}
		args = append(args, object.OwnerID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &photoR{}
			}

			for _, a := range args {
				if a == obj.OwnerID {
					continue Outer
				}
			}

			args = append(args, obj.OwnerID)

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
		object.R.Owner = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.OwnerPhotos = append(foreign.R.OwnerPhotos, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.OwnerID == foreign.UserID {
				local.R.Owner = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.OwnerPhotos = append(foreign.R.OwnerPhotos, local)
				break
			}
		}
	}

	return nil
}

// LoadPhotoExif allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-1 relationship.
func (photoL) LoadPhotoExif(ctx context.Context, e boil.ContextExecutor, singular bool, maybePhoto interface{}, mods queries.Applicator) error {
	var slice []*Photo
	var object *Photo

	if singular {
		object = maybePhoto.(*Photo)
	} else {
		slice = *maybePhoto.(*[]*Photo)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &photoR{}
		}
		args = append(args, object.PhotoID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &photoR{}
			}

			for _, a := range args {
				if a == obj.PhotoID {
					continue Outer
				}
			}

			args = append(args, obj.PhotoID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`photo_exif`),
		qm.WhereIn(`photo_exif.photo_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load PhotoExif")
	}

	var resultSlice []*PhotoExif
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice PhotoExif")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for photo_exif")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for photo_exif")
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
		object.R.PhotoExif = foreign
		if foreign.R == nil {
			foreign.R = &photoExifR{}
		}
		foreign.R.Photo = object
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PhotoID == foreign.PhotoID {
				local.R.PhotoExif = foreign
				if foreign.R == nil {
					foreign.R = &photoExifR{}
				}
				foreign.R.Photo = local
				break
			}
		}
	}

	return nil
}

// SetGroup of the photo to the related item.
// Sets o.R.Group to related.
// Adds o to related.R.Photos.
func (o *Photo) SetGroup(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Group) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `photos` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"group_id"}),
		strmangle.WhereClause("`", "`", 0, photoPrimaryKeyColumns),
	)
	values := []interface{}{related.GroupID, o.PhotoID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.GroupID = related.GroupID
	if o.R == nil {
		o.R = &photoR{
			Group: related,
		}
	} else {
		o.R.Group = related
	}

	if related.R == nil {
		related.R = &groupR{
			Photos: PhotoSlice{o},
		}
	} else {
		related.R.Photos = append(related.R.Photos, o)
	}

	return nil
}

// SetOwner of the photo to the related item.
// Sets o.R.Owner to related.
// Adds o to related.R.OwnerPhotos.
func (o *Photo) SetOwner(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `photos` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"owner_id"}),
		strmangle.WhereClause("`", "`", 0, photoPrimaryKeyColumns),
	)
	values := []interface{}{related.UserID, o.PhotoID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OwnerID = related.UserID
	if o.R == nil {
		o.R = &photoR{
			Owner: related,
		}
	} else {
		o.R.Owner = related
	}

	if related.R == nil {
		related.R = &userR{
			OwnerPhotos: PhotoSlice{o},
		}
	} else {
		related.R.OwnerPhotos = append(related.R.OwnerPhotos, o)
	}

	return nil
}

// SetPhotoExif of the photo to the related item.
// Sets o.R.PhotoExif to related.
// Adds o to related.R.Photo.
func (o *Photo) SetPhotoExif(ctx context.Context, exec boil.ContextExecutor, insert bool, related *PhotoExif) error {
	var err error

	if insert {
		related.PhotoID = o.PhotoID

		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE `photo_exif` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, []string{"photo_id"}),
			strmangle.WhereClause("`", "`", 0, photoExifPrimaryKeyColumns),
		)
		values := []interface{}{o.PhotoID, related.PhotoID}

		if boil.IsDebug(ctx) {
			writer := boil.DebugWriterFrom(ctx)
			fmt.Fprintln(writer, updateQuery)
			fmt.Fprintln(writer, values)
		}
		if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PhotoID = o.PhotoID

	}

	if o.R == nil {
		o.R = &photoR{
			PhotoExif: related,
		}
	} else {
		o.R.PhotoExif = related
	}

	if related.R == nil {
		related.R = &photoExifR{
			Photo: o,
		}
	} else {
		related.R.Photo = o
	}
	return nil
}

// Photos retrieves all the records using an executor.
func Photos(mods ...qm.QueryMod) photoQuery {
	mods = append(mods, qm.From("`photos`"))
	return photoQuery{NewQuery(mods...)}
}

// FindPhoto retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPhoto(ctx context.Context, exec boil.ContextExecutor, photoID int, selectCols ...string) (*Photo, error) {
	photoObj := &Photo{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `photos` where `photo_id`=?", sel,
	)

	q := queries.Raw(query, photoID)

	err := q.Bind(ctx, exec, photoObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: unable to select from photos")
	}

	if err = photoObj.doAfterSelectHooks(ctx, exec); err != nil {
		return photoObj, err
	}

	return photoObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Photo) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no photos provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(photoColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	photoInsertCacheMut.RLock()
	cache, cached := photoInsertCache[key]
	photoInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			photoAllColumns,
			photoColumnsWithDefault,
			photoColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(photoType, photoMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(photoType, photoMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `photos` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `photos` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `photos` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, photoPrimaryKeyColumns))
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
		return errors.Wrap(err, "dbmodels: unable to insert into photos")
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

	o.PhotoID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == photoMapping["photo_id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.PhotoID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to populate default values for photos")
	}

CacheNoHooks:
	if !cached {
		photoInsertCacheMut.Lock()
		photoInsertCache[key] = cache
		photoInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Photo.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Photo) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	photoUpdateCacheMut.RLock()
	cache, cached := photoUpdateCache[key]
	photoUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			photoAllColumns,
			photoPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodels: unable to update photos, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `photos` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, photoPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(photoType, photoMapping, append(wl, photoPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "dbmodels: unable to update photos row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by update for photos")
	}

	if !cached {
		photoUpdateCacheMut.Lock()
		photoUpdateCache[key] = cache
		photoUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q photoQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all for photos")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected for photos")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PhotoSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), photoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `photos` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, photoPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all in photo slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected all in update all photo")
	}
	return rowsAff, nil
}

var mySQLPhotoUniqueColumns = []string{
	"photo_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Photo) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no photos provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(photoColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLPhotoUniqueColumns, o)

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

	photoUpsertCacheMut.RLock()
	cache, cached := photoUpsertCache[key]
	photoUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			photoAllColumns,
			photoColumnsWithDefault,
			photoColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			photoAllColumns,
			photoPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("dbmodels: unable to upsert photos, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`photos`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `photos` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(photoType, photoMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(photoType, photoMapping, ret)
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
		return errors.Wrap(err, "dbmodels: unable to upsert for photos")
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

	o.PhotoID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == photoMapping["photo_id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(photoType, photoMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to retrieve unique values for photos")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to populate default values for photos")
	}

CacheNoHooks:
	if !cached {
		photoUpsertCacheMut.Lock()
		photoUpsertCache[key] = cache
		photoUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Photo record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Photo) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodels: no Photo provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), photoPrimaryKeyMapping)
	sql := "DELETE FROM `photos` WHERE `photo_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete from photos")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by delete for photos")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q photoQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodels: no photoQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from photos")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for photos")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PhotoSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(photoBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), photoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `photos` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, photoPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from photo slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for photos")
	}

	if len(photoAfterDeleteHooks) != 0 {
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
func (o *Photo) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPhoto(ctx, exec, o.PhotoID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PhotoSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PhotoSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), photoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `photos`.* FROM `photos` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, photoPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to reload all in PhotoSlice")
	}

	*o = slice

	return nil
}

// PhotoExists checks if the Photo row exists.
func PhotoExists(ctx context.Context, exec boil.ContextExecutor, photoID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `photos` where `photo_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, photoID)
	}
	row := exec.QueryRowContext(ctx, sql, photoID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: unable to check if photos exists")
	}

	return exists, nil
}