// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type ExampleAuthor struct {
	ID   int64
	Name string
	Bio  pgtype.Text
}
