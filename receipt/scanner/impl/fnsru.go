package impl

import (
	"encoding/json"
	"fmt"
	"github.com/micro/micro/v3/service/errors"
	"github.com/micro/micro/v3/service/logger"
	"github.com/pashamad/microapi/receipt/scanner/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var FnsRuCodeParser = func(code string) (m types.ReceiptMeta, e error) {
	logger.Info("Parsing code string with fnsRu parser: ", code)

	u, e := url.ParseQuery(code)
	if e != nil {
		logger.Error("failed to parse query: ", e.Error())
		return m, e
	}

	fn, fp, i, s, t := u.Get("fn"), u.Get("fp"), u.Get("i"), u.Get("s"), u.Get("i")
	if len(fn) > 0 && len(fp) > 0 && len(i) > 0 && len(s) > 0 || len(t) > 0 {
		m.Fields = make(map[string]string, 5)
		m.Fields["fn"] = fn[:0]
		m.Fields["fp"] = fp[:0]
		m.Fields["i"] = i[:0]
		m.Fields["s"] = s[:0]
		m.Fields["t"] = t[:0]
		return m, nil
	}

	return m, types.UnfitParser
}

var FnsRruMetaResolver types.MetaResolver = func(receipt types.ReceiptMeta) (meta types.OrderMeta, err error) {

	// prepare request // todo: modular back-ends
	u := "https://proverkacheka.com/check/get"
	v := url.Values{}
	v.Add("fn", receipt.Fields["fn"])
	v.Add("fp", receipt.Fields["fp"])
	v.Add("fd", receipt.Fields["fd"])

	// send post with urlencoded data // todo: build from values map
	req := strings.NewReader(receipt.Src)
	res, err := http.Post(u, "application/x-www-form-urlencoded", req)
	if err != nil {
		return meta, err
	}
	//noinspection GoUnhandledErrorResult
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	// check response and handle error
	var summary summary
	err = json.Unmarshal(body, &summary)
	if err != nil {
		return meta, err
	}
	if summary.Code == 5 {
		return meta, types.CodeNotResolved
	} else if summary.Code == 4 {
		return meta, types.CodeResolveTimeout
	} else if summary.Code != 1 {
		return meta, errors.InternalServerError("RESOLVE_ERROR_UNKNOWN",
			"error occurred while resolving receipt code; error code %d; full response: %s",
			summary.Code, string(body))
	}

	// unmarshal successful response
	var doc doc
	err = json.Unmarshal(body, &doc)
	if err != nil {
		return meta, err
	}

	// copy resolved data to result meta
	meta.Resolved = make(map[string]string, 3)
	b := doc.Data.Body
	meta.Resolved["user"] = strings.TrimSpace(b.User)
	meta.Resolved["inn"] = strings.TrimSpace(b.INN)
	meta.Resolved["amount"] = fmt.Sprintf("%.2f", float64(b.Amount)/100)

	return meta, nil
}

// todo: create separate routine
type summary struct {
	Code int `json:"code"`
}
type doc struct {
	Code int  `json:"code"`
	Data data `json:"data"`
}
type data struct {
	Body body `json:"json"`
}
type body struct {
	Code   int    `json:"code"`
	User   string `json:"user"`
	INN    string `json:"userInn"`
	Amount int    `json:"totalSum"`
}
