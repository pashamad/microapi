package scanner

import (
	"github.com/micro/micro/v3/service/logger"
	"github.com/pashamad/microapi/receipt/scanner/impl"
	"github.com/pashamad/microapi/receipt/scanner/types"
)

var (
	resolvers = map[types.ProviderKey]types.MetaResolver{
		types.FnsRu: impl.FnsRruMetaResolver,
	}
)

var ResolveMeta types.MetaResolver = func(receipt types.ReceiptMeta) (meta types.OrderMeta, err error) {
	// resolve order meta
	resolver := resolvers[receipt.Provider]
	meta, err = resolver(receipt)
	if err != nil {
		logger.Error("Failed to resolve order meta: ", err.Error())
		return meta, err
	}
	return meta, nil
}
