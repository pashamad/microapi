package scanner

import (
	"context"
	"github.com/micro/micro/v3/service"
	log "github.com/micro/micro/v3/service/logger"
	org "github.com/pashamad/microapi/org/proto"
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
		log.Error("Failed to resolve order meta: ", err.Error())
		return meta, err
	}

	// lookup org entity for order meta
	srv := service.New()
	client := org.NewOrgService("org", srv.Client())
	// todo: lookup by tin is provider-specific
	rsp, err := client.Lookup(context.Background(), &org.LookupRequest{
		Tin: meta.Resolved["tin"],
	})
	if err != nil {
		log.Error("error resolving org entity for receipt code: ", err)
		return meta, err
	}
	meta.Entity = *rsp.Entity

	return meta, nil
}
