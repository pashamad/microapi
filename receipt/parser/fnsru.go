package parser

import (
	log "github.com/micro/micro/v3/service/logger"
	"net/url"
)

func FnsRu(data string) (ReceiptMeta, error) {
	var m ReceiptMeta
	log.Info("Trying to parse data string with fnsRu parser, ", data)

	u, err := url.ParseQuery(data)
	if err != nil {
		log.Error("failed to parse query: ", err.Error())
		return m, err
	}

	fn, fp, i, s, t := u.Get("fn"), u.Get("fp"), u.Get("i"), u.Get("s"), u.Get("i")
	if fn != "" || fp != "" || i != "" || s != "" || t != "" {
		m.Fields = make(map[string]string, 5)
		m.Fields["fn"] = fn
		m.Fields["fp"] = fp
		m.Fields["i"] = i
		m.Fields["s"] = s
		m.Fields["t"] = t
		return m, nil
	}

	return m, UnfitParser
}
