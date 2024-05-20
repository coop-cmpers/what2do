package helpers

import "github.com/gofrs/uuid"

// Iterate until we generate a valid UUID v4
func GenerateUUID() uuid.UUID {
	id, err := uuid.NewV4()

	for err != nil {
		id, err = uuid.NewV4()
	}

	return id
}
