package auth

import (
	"fmt"
	authapp "goCartAPI/application/auth"
	"goCartAPI/domain/shared"
	"goCartAPI/domain/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret                []byte
	accessTTL, refreshTTL time.Duration
}

type claims struct { //payload
	Role string `json:"role"`
	Kind string `json:"kind"`
	jwt.RegisteredClaims
}

func NewJWTService(secret string, accessTTL, refreshTTL time.Duration) *JWTService {
	return &JWTService{
		secret:     []byte(secret),
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
	}
}

func (s *JWTService) generate(userID shared.ID, role user.Role, kind string, ttl time.Duration) (string, error) {
	now := time.Now().UTC()
	c := claims{Role: string(role),
		Kind: kind,
		RegisteredClaims: jwt.RegisteredClaims{Subject: userID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		}}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(s.secret)
}

func (s *JWTService) GenerateAccessToken(userID shared.ID, role user.Role) (string, error) {
	return s.generate(userID, role, "access", s.accessTTL)
}
func (s *JWTService) GenerateRefreshToken(userID shared.ID, role user.Role) (string, error) {
	return s.generate(userID, role, "refresh", s.refreshTTL)
}

func (s *JWTService) ParseAccessToken(raw string) (authapp.TokenClaims, error) {
	return s.parse(raw, "access")
}

func (s *JWTService) ParseRefreshToken(raw string) (authapp.TokenClaims, error) {
	return s.parse(raw, "refresh")
}

func (s *JWTService) parse(raw, expectedKind string) (authapp.TokenClaims, error) {
	tok, err := jwt.ParseWithClaims(raw, &claims{},
		func(t *jwt.Token) (any, error) {
			if t.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("unexpected signing method")

			}
			return s.secret, nil
		},
	)
	if err != nil || tok.Valid {
		return authapp.TokenClaims{}, fmt.Errorf("invalid token")
	}

	c, ok := tok.Claims.(*claims)

	if !ok || c.Kind != expectedKind {
		return authapp.TokenClaims{}, fmt.Errorf("invalid token claims")

	}

	id, err := shared.ParseID(c.Subject)
	if err != nil {
		return authapp.TokenClaims{}, err
	}
	role := user.Role(c.Role)
	if !role.IsValid() {
		return authapp.TokenClaims{}, user.ErrInvalidRole
	}
	return authapp.TokenClaims{
		UserID: id, Role: role,
	}, nil

}
