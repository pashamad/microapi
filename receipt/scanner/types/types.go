package types

import "github.com/pashamad/microapi/orders/models"

type ProviderKey string

type ReceiptMeta struct {
	Provider ProviderKey
	Src      string
	Fields   map[string]string
}

// todo: biz, cart, amount
type OrderMeta struct {
	Receipt  ReceiptMeta
	Resolved map[string]string
}

type CodeParser func(string) (ReceiptMeta, error)

type MetaResolver func(meta ReceiptMeta) (OrderMeta, error)

type OrderBuilder func(meta ReceiptMeta) (models.Order, error)

type ScanContainer struct {
	Parse      CodeParser
	Resolve    MetaResolver
	BuildOrder OrderBuilder
}
