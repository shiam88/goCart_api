package user

import (
	"goCartAPI/domain/shared"
	"net/mail"
	"strings"
	"time"
)

type User struct {
	id           shared.ID
	name         string
	email        string
	passwordHash string
	role         Role
	createdAt    time.Time
	updatedAt    time.Time
}

func NewUser(name, email, passwordHash string) (*User, error) {
	now := time.Now().UTC()
	return restoreUser(shared.NewID(), name, email, passwordHash, RoleCustomer, now, now)
}

func Restore(id shared.ID, name, email, passwordHash string, role Role, createdAt, updatedAt time.Time) (*User, error) {
	return restoreUser(id, name, email, passwordHash, role, createdAt, updatedAt)
}

func restoreUser(id shared.ID, name, email, passwordHash string, role Role, createdAt, updatedAt time.Time) (*User, error) {
	if shared.IsZeroID(id) {
		return nil, ErrInvalidUserID
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, ErrInvalidName
	}
	email = strings.ToLower(strings.TrimSpace(email))
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, ErrInvalidEmail
	}
	if strings.TrimSpace(passwordHash) == "" {
		return nil, ErrMissingPasswordHash
	}
	if !role.IsValid() {
		return nil, ErrInvalidRole
	}
	return &User{id: id, name: name, email: email, passwordHash: passwordHash, role: role, createdAt: createdAt.UTC(), updatedAt: updatedAt.UTC()}, nil
}

func (u *User) ID() shared.ID {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

func (u *User) PasswordHash() string {
	return u.passwordHash
}
func (u *User) Role() Role {
	return u.role
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}
func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) Rename(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return ErrInvalidName
	}
	u.name = name
	u.touch()
	return nil
}

func (u *User) ChangePasswordHash(hash string) error {
	if strings.TrimSpace(hash) == "" {
		return ErrMissingPasswordHash
	}
	u.passwordHash = hash
	u.touch()
	return nil
}

func (u *User) ChangeRole(role Role) error {
	if !role.IsValid() {
		return ErrInvalidRole
	}
	u.role = role
	u.touch()
	return nil
}

func (u *User) touch() {
	u.updatedAt = time.Now().UTC()
}
