package auth

import "context"

// v1 logout is stateless: the client discards its tokens.
func (s *service) Logout(ctx context.Context) error {
	_ = ctx
	return nil
}
