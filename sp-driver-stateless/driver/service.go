package driver

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/cockroachdb/apd/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrAccountNotFound = errors.New("account not found")

type DriverService struct {
	pool *pgxpool.Pool
}

func NewDriverService(connPool *pgxpool.Pool) *DriverService {
	return &DriverService{
		pool: connPool,
	}
}

func (s *DriverService) GetBalance(ctx context.Context, accountID int) (string, error) {
	var balance sql.NullString
	error := s.pool.QueryRow(ctx, "select get_balance($1)", accountID).Scan(&balance)
	if error != nil {
		return "", fmt.Errorf("failed to get balance: %w", error)
	}
	if !balance.Valid {
		return "", ErrAccountNotFound
	}
	return balance.String, nil
}

func (s *DriverService) Transfer(ctx context.Context, fromAccountID int, toAccountID int, amount apd.Decimal) error {
	_, err := s.pool.Exec(ctx, "call transfer($1, $2, $3::dec(15,2))", fromAccountID, toAccountID, amount.String())
	if err != nil {
		return fmt.Errorf("failed to transfer: %w", err)
	}
	return nil
}
