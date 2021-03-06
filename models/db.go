// Code generated by sqlc. DO NOT EDIT.

package models

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
	if q.createCategoryStmt, err = db.PrepareContext(ctx, createCategory); err != nil {
		return nil, fmt.Errorf("error preparing query CreateCategory: %w", err)
	}
	if q.createFilmStmt, err = db.PrepareContext(ctx, createFilm); err != nil {
		return nil, fmt.Errorf("error preparing query CreateFilm: %w", err)
	}
	if q.deleteFilmStmt, err = db.PrepareContext(ctx, deleteFilm); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteFilm: %w", err)
	}
	if q.getAllCategoryStmt, err = db.PrepareContext(ctx, getAllCategory); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllCategory: %w", err)
	}
	if q.getAllFilmStmt, err = db.PrepareContext(ctx, getAllFilm); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllFilm: %w", err)
	}
	if q.getFilmStmt, err = db.PrepareContext(ctx, getFilm); err != nil {
		return nil, fmt.Errorf("error preparing query GetFilm: %w", err)
	}
	if q.updateFilmStmt, err = db.PrepareContext(ctx, updateFilm); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateFilm: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createCategoryStmt != nil {
		if cerr := q.createCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createCategoryStmt: %w", cerr)
		}
	}
	if q.createFilmStmt != nil {
		if cerr := q.createFilmStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createFilmStmt: %w", cerr)
		}
	}
	if q.deleteFilmStmt != nil {
		if cerr := q.deleteFilmStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteFilmStmt: %w", cerr)
		}
	}
	if q.getAllCategoryStmt != nil {
		if cerr := q.getAllCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllCategoryStmt: %w", cerr)
		}
	}
	if q.getAllFilmStmt != nil {
		if cerr := q.getAllFilmStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllFilmStmt: %w", cerr)
		}
	}
	if q.getFilmStmt != nil {
		if cerr := q.getFilmStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getFilmStmt: %w", cerr)
		}
	}
	if q.updateFilmStmt != nil {
		if cerr := q.updateFilmStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateFilmStmt: %w", cerr)
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
	db                 DBTX
	tx                 *sql.Tx
	createCategoryStmt *sql.Stmt
	createFilmStmt     *sql.Stmt
	deleteFilmStmt     *sql.Stmt
	getAllCategoryStmt *sql.Stmt
	getAllFilmStmt     *sql.Stmt
	getFilmStmt        *sql.Stmt
	updateFilmStmt     *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                 tx,
		tx:                 tx,
		createCategoryStmt: q.createCategoryStmt,
		createFilmStmt:     q.createFilmStmt,
		deleteFilmStmt:     q.deleteFilmStmt,
		getAllCategoryStmt: q.getAllCategoryStmt,
		getAllFilmStmt:     q.getAllFilmStmt,
		getFilmStmt:        q.getFilmStmt,
		updateFilmStmt:     q.updateFilmStmt,
	}
}
