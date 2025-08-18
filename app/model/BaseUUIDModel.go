package model

import "github.com/google/uuid"

type BaseUUID struct {
	ID uuid.UUID `json:"id"`
}

func (b *BaseUUID) GenerateUUID() {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
}
