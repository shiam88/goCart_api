package cart

import (
	"goCartAPI/domain/shared"
	"time"
)

type Cart struct {
	id        shared.ID
	userID    shared.ID
	currency  string
	items     map[shared.ID]*Item
	createdAt time.Time
	updatedAt time.Time
}

func NewCart(userID shared.ID, currency string) (*Cart, error) {
	if shared.IsZeroID(userID) {
		return nil, ErrInvalidUserID
	}
	if _, err := shared.ZeroMoney(currency); err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	return &Cart{id: shared.NewID(), userID: userID, currency: currency, items: make(map[shared.ID]*Item), createdAt: now, updatedAt: now}, nil
}

func Restore(id, userID shared.ID, currency string, items []*Item, createdAt, updatedAt time.Time) (*Cart, error) {
	if shared.IsZeroID(id) {
		return nil, ErrInvalidCartID
	}
	if shared.IsZeroID(userID) {
		return nil, ErrInvalidUserID
	}
	if _, err := shared.ZeroMoney(currency); err != nil {
		return nil, err
	}
	itemMap := make(map[shared.ID]*Item, len(items))
	for _, item := range items {
		if item == nil {
			continue
		}
		if item.UnitPrice().Currency() != currency {
			return nil, shared.ErrCurrencyMismatch
		}
		copyItem := *item
		itemMap[item.ProductID()] = &copyItem
	}
	return &Cart{id: id, userID: userID, currency: currency, items: itemMap, createdAt: createdAt.UTC(), updatedAt: updatedAt.UTC()}, nil
}

func (c *Cart) ID() shared.ID        { return c.id }
func (c *Cart) UserID() shared.ID    { return c.userID }
func (c *Cart) Currency() string     { return c.currency }
func (c *Cart) CreatedAt() time.Time { return c.createdAt }
func (c *Cart) UpdatedAt() time.Time { return c.updatedAt }
func (c *Cart) IsEmpty() bool        { return len(c.items) == 0 }
func (c *Cart) ItemCount() int       { return len(c.items) }

func (c *Cart) Items() []Item {
	items := make([]Item, 0, len(c.items))
	for _, item := range c.items {
		if item != nil {
			items = append(items, *item)
		}
	}
	return items
}

func (c *Cart) TotalQuantity() int {
	total := 0
	for _, item := range c.items {
		if item != nil {
			total += item.Quantity()
		}
	}
	return total
}

func (c *Cart) AddItem(productID shared.ID, productName string, unitPrice shared.Money, quantity int) error {
	if unitPrice.Currency() != c.currency {
		return shared.ErrCurrencyMismatch
	}
	if quantity <= 0 {
		return ErrInvalidQuantity
	}
	if existing, ok := c.items[productID]; ok {
		if err := existing.UpdateQuantity(existing.Quantity() + quantity); err != nil {
			return err
		}
		c.touch()
		return nil
	}
	item, err := NewItem(productID, productName, unitPrice, quantity)
	if err != nil {
		return err
	}
	c.items[productID] = item
	c.touch()
	return nil
}

func (c *Cart) UpdateQuantity(productID shared.ID, quantity int) error {
	item, ok := c.items[productID]
	if !ok {
		return ErrItemNotFound
	}
	if err := item.UpdateQuantity(quantity); err != nil {
		return err
	}
	c.touch()
	return nil
}

func (c *Cart) RemoveItem(productID shared.ID) error {
	if _, ok := c.items[productID]; !ok {
		return ErrItemNotFound
	}
	delete(c.items, productID)
	c.touch()
	return nil
}

func (c *Cart) Clear() {
	c.items = make(map[shared.ID]*Item)
	c.touch()
}

func (c *Cart) Total() (shared.Money, error) {
	total, err := shared.ZeroMoney(c.currency)
	if err != nil {
		return shared.Money{}, err
	}
	for _, item := range c.items {
		if item == nil {
			continue
		}
		subtotal, err := item.Subtotal()
		if err != nil {
			return shared.Money{}, err
		}
		total, err = total.Add(subtotal)
		if err != nil {
			return shared.Money{}, err
		}
	}
	return total, nil
}

func (c *Cart) touch() { c.updatedAt = time.Now().UTC() }
