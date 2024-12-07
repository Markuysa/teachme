package client

import (
	"context"
	"fmt"

	"gitlab.com/coinhubs/balance/internal/domain"
)

func (s *service) SignUpRequest(
	ctx context.Context,
	request v1.,
) (domain.User, error) {
	user, err := s.repos.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to create user: %w", err)
	}
}
