package parser

import (
	goerr "errors"
	"github.com/micro/micro/v3/service/errors"
)

var (
	ErrUnfitCode = errors.BadRequest("UNFIT_CODE_FORMAT", "Unsupported code string passed")
)

var (
	parsers = map[string]Parser{
		"fnsRu": FnsRu,
	}
)

var UnfitParser = goerr.New("UNFIT_PARSER")

type ReceiptMeta struct {
	Type   string
	Src    string
	Fields map[string]string
}

type Parser func(string) (ReceiptMeta, error)

func Parse(data string) (ReceiptMeta, error) {
	var meta ReceiptMeta
	// try each parser
	for key, parser := range parsers {
		meta, err := parser(data)
		if err == UnfitParser {
			continue
		} else if err != nil {
			return meta, err
		} else if &meta != nil {
			meta.Type = key
			meta.Src = data
			return meta, nil
		}
	}
	// failed to parse
	return meta, ErrUnfitCode
}
