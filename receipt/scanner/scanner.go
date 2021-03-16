package scanner

import (
	"github.com/google/uuid"
	log "github.com/micro/micro/v3/service/logger"
	"github.com/pashamad/microapi/orders/models"
	"strconv"
	"time"
)

func GetOrder(code string) (*models.Order, error) {
	// parse code
	receipt, err := Parse(code)
	if err != nil {
		return nil, err
	}

	// resolve order receipt
	meta, err := ResolveMeta(receipt)
	if err != nil {
		return nil, err
	}
	log.Infof("Resolved order receipt: %s", meta.Resolved)

	// todo: build real order
	amount, _ := strconv.ParseFloat(meta.Resolved["amount"], 32)
	order := &models.Order{
		ID:          uuid.UUID{},
		Amount:      float32(amount),
		Description: "",
		CreatedAt:   time.Time{},
	}
	return order, nil
}
