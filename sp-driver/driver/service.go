package driver

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrSessionNotFound = errors.New("session not found")

type DriverService struct {
	sessions map[uuid.UUID]Session
}

func NewDriverService() *DriverService {
	return &DriverService{
		sessions: make(map[uuid.UUID]Session),
	}
}

func (s *DriverService) CreateSession(ctx context.Context, username, password, url string) (uuid.UUID, error) {
	session, err := NewSession(username, password, url)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create session: %w", err)
	}
	s.sessions[session.ID()] = session
	return session.ID(), nil
}

func (s *DriverService) CloseSession(ctx context.Context, id uuid.UUID) error {
	if session, ok := s.sessions[id]; ok {
		session.Logout(ctx)
		delete(s.sessions, id)
		return nil
	}
	return ErrSessionNotFound
}

func (s *DriverService) GetBalance(ctx context.Context, sessionID uuid.UUID, accountID int) (string, error) {
	if session, ok := s.sessions[sessionID]; ok {
		return session.GetBalance(ctx, accountID)
	}
	return "", ErrSessionNotFound
}

func (s *DriverService) Transfer(ctx context.Context, sessionID uuid.UUID, request Transfer) error {
	if session, ok := s.sessions[sessionID]; ok {
		return session.Transfer(ctx, request)
	}
	return ErrSessionNotFound
}
