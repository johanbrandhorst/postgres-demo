// Code generated by sqlc. DO NOT EDIT.

package users

import (
	"context"

	"github.com/jackc/pgtype"
)

type Querier interface {
	AddUser(ctx context.Context, role Role) (User, error)
	DeleteUser(ctx context.Context, id pgtype.UUID) (User, error)
	GetUser(ctx context.Context, id pgtype.UUID) (User, error)
}

var _ Querier = (*Queries)(nil)
