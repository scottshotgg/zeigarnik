// Code generated by sqlc. DO NOT EDIT.

package sql

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createReminderStmt, err = db.PrepareContext(ctx, createReminder); err != nil {
		return nil, fmt.Errorf("error preparing query CreateReminder: %w", err)
	}
	if q.deleteReminderStmt, err = db.PrepareContext(ctx, deleteReminder); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteReminder: %w", err)
	}
	if q.getReminderByIDStmt, err = db.PrepareContext(ctx, getReminderByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetReminderByID: %w", err)
	}
	if q.updateReminderStmt, err = db.PrepareContext(ctx, updateReminder); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateReminder: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createReminderStmt != nil {
		if cerr := q.createReminderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createReminderStmt: %w", cerr)
		}
	}
	if q.deleteReminderStmt != nil {
		if cerr := q.deleteReminderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteReminderStmt: %w", cerr)
		}
	}
	if q.getReminderByIDStmt != nil {
		if cerr := q.getReminderByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getReminderByIDStmt: %w", cerr)
		}
	}
	if q.updateReminderStmt != nil {
		if cerr := q.updateReminderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateReminderStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                  DBTX
	tx                  *sql.Tx
	createReminderStmt  *sql.Stmt
	deleteReminderStmt  *sql.Stmt
	getReminderByIDStmt *sql.Stmt
	updateReminderStmt  *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                  tx,
		tx:                  tx,
		createReminderStmt:  q.createReminderStmt,
		deleteReminderStmt:  q.deleteReminderStmt,
		getReminderByIDStmt: q.getReminderByIDStmt,
		updateReminderStmt:  q.updateReminderStmt,
	}
}