package shared

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return uuid.New()

}

func ParseID(value string) (ID, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil, ErrInvalidID
	}

	return id, nil
}

func IsZeroID(id ID) bool {
	return id == uuid.Nil
}
