package repositories

import (
	"context"
	"github.com/hiroyky/famiphoto/drivers/mysql"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type GroupUserRepository interface {
	CreateGroupUsers(ctx context.Context, groupUsers dbmodels.GroupUserSlice) error
	DeleteGroupUsers(ctx context.Context, groupUsers dbmodels.GroupUserSlice) error
}

func NewGroupUserRepository(db mysql.SQLExecutor) GroupUserRepository {
	return &groupUserRepository{
		db: db,
	}
}

type groupUserRepository struct {
	db mysql.SQLExecutor
}

func (r *groupUserRepository) CreateGroupUsers(ctx context.Context, groupUsers dbmodels.GroupUserSlice) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.New(errors.TxnBeginFatal, err)
	}

	if err := r.createGroupUsersTxn(ctx, tx, groupUsers); err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.New(errors.TxnRollbackFatal, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.New(errors.TxnRollbackFatal, err)
		}
		return err
	}
	return nil
}

func (r *groupUserRepository) createGroupUsersTxn(ctx context.Context, tx boil.ContextExecutor, groupUsers dbmodels.GroupUserSlice) error {
	for _, gu := range groupUsers {
		if exit, err := dbmodels.GroupUserExists(ctx, tx, gu.GroupID, gu.UserID); err != nil {
			return err
		} else if exit {
			continue
		}
		if err := gu.Insert(ctx, tx, boil.Infer()); err != nil {
			return err
		}
	}
	return nil
}

func (r *groupUserRepository) DeleteGroupUsers(ctx context.Context, groupUsers dbmodels.GroupUserSlice) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.New(errors.TxnBeginFatal, err)
	}

	if err := r.deleteGroupUsersTxn(ctx, tx, groupUsers); err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.New(errors.TxnRollbackFatal, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.New(errors.TxnRollbackFatal, err)
		}
		return err
	}
	return nil
}

func (r *groupUserRepository) deleteGroupUsersTxn(ctx context.Context, tx boil.ContextExecutor, groupUsers dbmodels.GroupUserSlice) error {
	for _, gu := range groupUsers {
		if exit, err := dbmodels.GroupUserExists(ctx, tx, gu.GroupID, gu.UserID); err != nil {
			return err
		} else if !exit {
			continue
		}

		if _, err := gu.Delete(ctx, tx); err != nil {
			return err
		}
	}
	return nil
}
