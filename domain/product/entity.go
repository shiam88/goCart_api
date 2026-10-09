package product

import (
	"strings"
	"time"
	"unicode"

	"goCartAPI/domain/shared"
)

type Product struct {
	id          shared.ID
	categoryID  shared.ID
	name        string
	slug        string
	sku         string //Stock Keeping Unit
	description string
	price       shared.Money
	imageURL    string
	active      bool
	createdAt   time.Time
	updatedAt   time.Time
}

func NewProduct(categoryID shared.ID, name, sku, description string, price shared.Money, imageURL string) (*Product, error) {
	if shared.IsZeroID(categoryID) {
		return nil, ErrInvalidCategoryID
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, ErrInvalidName
	}
	sku = strings.ToUpper(strings.TrimSpace(sku))
	if sku == "" {
		return nil, ErrInvalidSKU
	}
	if price.Amount() <= 0 {
		return nil, ErrInvalidPrice
	}
	now := time.Now().UTC()
	return &Product{id: shared.NewID(), categoryID: categoryID, name: name, slug: slugify(name), sku: sku, description: strings.TrimSpace(description), price: price, imageURL: strings.TrimSpace(imageURL), active: true, createdAt: now, updatedAt: now}, nil
}

func Restore(id, categoryID shared.ID, name, slug, sku, description string, price shared.Money, imageURL string, active bool, createdAt, updatedAt time.Time) (*Product, error) {
	if shared.IsZeroID(id) {
		return nil, ErrInvalidProductID
	}
	p, err := NewProduct(categoryID, name, sku, description, price, imageURL)
	if err != nil {
		return nil, err
	}
	p.id = id
	p.slug = strings.TrimSpace(slug)
	if p.slug == "" {
		p.slug = slugify(name)
	}
	p.active = active
	p.createdAt = createdAt.UTC()
	p.updatedAt = updatedAt.UTC()
	return p, nil
}

func (p *Product) ID() shared.ID         { return p.id }
func (p *Product) CategoryID() shared.ID { return p.categoryID }
func (p *Product) Name() string          { return p.name }
func (p *Product) Slug() string          { return p.slug }
func (p *Product) SKU() string           { return p.sku }
func (p *Product) Description() string   { return p.description }
func (p *Product) Price() shared.Money   { return p.price }
func (p *Product) ImageURL() string      { return p.imageURL }

func (p *Product) IsActive() bool {
	return p.active
}
func (p *Product) CreatedAt() time.Time { return p.createdAt }
func (p *Product) UpdatedAt() time.Time { return p.updatedAt }

func (p *Product) Rename(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return ErrInvalidName
	}
	p.name = name
	p.slug = slugify(name)
	p.touch()
	return nil
}
func (p *Product) ChangePrice(price shared.Money) error {
	if price.Amount() <= 0 {
		return ErrInvalidPrice
	}
	p.price = price
	p.touch()
	return nil
}
func (p *Product) ChangeCategory(categoryID shared.ID) error {
	if shared.IsZeroID(categoryID) {
		return ErrInvalidCategoryID
	}
	p.categoryID = categoryID
	p.touch()
	return nil
}
func (p *Product) UpdateDetails(description, imageURL string) {
	p.description = strings.TrimSpace(description)
	p.imageURL = strings.TrimSpace(imageURL)
	p.touch()
}
func (p *Product) Activate()   { p.active = true; p.touch() }
func (p *Product) Deactivate() { p.active = false; p.touch() }
func (p *Product) touch()      { p.updatedAt = time.Now().UTC() }

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
