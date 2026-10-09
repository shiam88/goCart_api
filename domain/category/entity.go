package category

import (
	"goCartAPI/domain/shared"
	"strings"
	"time"
	"unicode"
)

type Category struct {
	id        shared.ID
	name      string
	slug      string
	active    bool
	createdAt time.Time
	updatedAt time.Time
}

func NewCategory(name string) (*Category, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, ErrInvalidName
	}
	now := time.Now().UTC()
	return &Category{id: shared.NewID(), name: name, slug: slugify(name), active: true, createdAt: now, updatedAt: now}, nil
}

func Restore(id shared.ID, name, slug string, active bool, createdAt, updatedAt time.Time) (*Category, error) {
	if shared.IsZeroID(id) {
		return nil, ErrInvalidCategoryID
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, ErrInvalidName
	}
	if strings.TrimSpace(slug) == "" {
		slug = slugify(name)
	}
	return &Category{id: id, name: name, slug: slug, active: active, createdAt: createdAt.UTC(), updatedAt: updatedAt.UTC()}, nil
}

func (c *Category) ID() shared.ID {
	return c.id
}
func (c *Category) Name() string {
	return c.name
}

func (c *Category) Slug() string {
	return c.slug
}

func (c *Category) IsActive() bool {
	return c.active
}

func (c *Category) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Category) UpdatedAt() time.Time {
	return c.updatedAt
}

func (c *Category) Rename(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return ErrInvalidName
	}
	c.name = name
	c.slug = slugify(name)
	c.updatedAt = time.Now().UTC()
	return nil
}
func (c *Category) Activate() {
	c.active = true
	c.updatedAt = time.Now().UTC()
}

func (c *Category) Deactivate() {
	c.active = false
	c.updatedAt = time.Now().UTC()
}

func slugify(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	var b strings.Builder
	lastDash := false
	for _, r := range value {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			b.WriteRune(r)
			lastDash = false
		case unicode.IsSpace(r) || r == '-' || r == '_':
			if b.Len() > 0 && !lastDash {
				b.WriteByte('-')
				lastDash = true
			}
		}
	}
	return strings.Trim(b.String(), "-")
}
