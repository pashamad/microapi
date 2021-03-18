package scanner

import (
	"context"
	"github.com/google/uuid"
	log "github.com/micro/micro/v3/service/logger"
	"github.com/pashamad/microapi/auth/user"
	"github.com/pashamad/microapi/orders/models"
	"github.com/pashamad/microapi/receipt/scanner/types"
	"strconv"
	"time"
)

// @todo do we really need to pass a context here
func GetOrder(ctx context.Context, code string) (*models.Order, error) {
	// get authenticated user
	userId, err := user.GetAppUser(ctx)
	//userId, err := user.GetAppUserTest()
	if err != nil {
		return nil, types.UserNotAuthenticated
	}
	log.Infof("User authenticated with id \"%s\"", userId)

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

	entity := meta.Entity
	log.Debugf("Org entity for receipt: %+v", entity)

	// todo: build order
	amount, _ := strconv.ParseFloat(meta.Resolved["amount"], 32)
	order := &models.Order{
		ID:          uuid.UUID{},
		Amount:      float32(amount),
		Description: "",
		CreatedAt:   time.Time{},
	}
	return order, nil
}
