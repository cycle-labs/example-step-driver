package driver

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var ErrAccountNotFound = errors.New("account not found")

type Session interface {
	ID() uuid.UUID
	Logout(ctx context.Context) error
	GetBalance(ctx context.Context, accountID int) (string, error)
	Transfer(ctx context.Context, request Transfer) error
}

type DBSession struct {
	id   uuid.UUID
	conn *pgx.Conn
}

func NewSession(username, password, url string) (Session, error) {
	connString := fmt.Sprintf("postgresql://%s:%s@%s", username, password, url)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return &DBSession{
		id:   uuid.New(),
		conn: conn,
	}, nil
}

func (s *DBSession) ID() uuid.UUID {
	return s.id
}

func (s *DBSession) Logout(ctx context.Context) error {
	return s.conn.Close(ctx)
}

func (s *DBSession) GetBalance(ctx context.Context, accountID int) (string, error) {
	var balance sql.NullString
	error := s.conn.QueryRow(ctx, "select get_balance($1)", accountID).Scan(&balance)
	if error != nil {
		return "", fmt.Errorf("failed to get balance: %w", error)
	}
	if !balance.Valid {
		return "", ErrAccountNotFound
	}
	return balance.String, nil
}

func (s *DBSession) Transfer(ctx context.Context, request Transfer) error {
	_, err := s.conn.Exec(ctx, "call transfer($1, $2, $3::dec(15,2))",
		request.FromAccountID, request.ToAccountID, request.Amount.String())
	if err != nil {
		return fmt.Errorf("failed to transfer: %w", err)
	}
	return nil
}
