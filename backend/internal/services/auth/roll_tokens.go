package auth

import (
	"context"
	"fmt"

	"github.com/ksusonic/niti/backend/internal/models"
)

func (s *Service) RollTokens(ctx context.Context, refresh *models.RefreshToken) (*models.JWTokens, error) {
	token, err := s.GenerateTokens(ctx, refresh.UserID)
	if err != nil {
		return nil, fmt.Errorf("generate tokens: %w", err)
	}

	err = s.refreshTokenRepo.Revoke(ctx, refresh.JTI)
	if err != nil {
		return nil, fmt.Errorf("revoke old token: %w", err)
	}

	return &token, nil
}
