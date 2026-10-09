package auth

import "context"

func (s *service) Refresh(ctx context.Context, refreshToken string) (*RefreshResult, error) {
	_ = ctx
	claims, err := s.tokens.ParseRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}
	access, err := s.tokens.GenerateAccessToken(claims.UserID, claims.Role)
	if err != nil {
		return nil, err
	}
	return &RefreshResult{AccessToken: access}, nil
}
