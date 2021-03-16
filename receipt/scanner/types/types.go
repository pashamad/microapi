package types

import (
	"github.com/pashamad/microapi/orders/models"
	org "github.com/pashamad/microapi/org/proto"
)

type ProviderKey string

type ReceiptMeta struct {
	Provider ProviderKey
	Src      string
	Fields   map[string]string
}

// todo: bizShop, orderCart
type OrderMeta struct {
	Receipt  ReceiptMeta
	Resolved map[string]string
	Entity   org.Entity
}

type CodeParser func(string) (ReceiptMeta, error)

type MetaResolver func(meta ReceiptMeta) (OrderMeta, error)

type OrderBuilder func(meta ReceiptMeta) (models.Order, error)

type ScanContainer struct {
	Parse      CodeParser
	Resolve    MetaResolver
	BuildOrder OrderBuilder
}
