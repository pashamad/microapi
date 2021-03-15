package models

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID          uuid.UUID
	Amount      float32
	Description string
	CreatedAt   time.Time
}
