package scanner

import (
	"github.com/pashamad/microapi/receipt/scanner/impl"
	"github.com/pashamad/microapi/receipt/scanner/types"
)

var (
	parsers = map[types.ProviderKey]types.CodeParser{
		types.FnsRu: impl.FnsRuCodeParser,
	}
)

var Parse types.CodeParser = func(code string) (types.ReceiptMeta, error) {
	var meta types.ReceiptMeta
	// try each scanner
	for prov, parser := range parsers {
		meta, err := parser(code)
		if err == types.UnfitParser {
			continue
		} else if err != nil {
			return meta, err
		} else if &meta != nil {
			meta.Provider = prov
			meta.Src = code
			return meta, nil
		}
	}
	// failed to parse
	return meta, types.UnfitCode
}
