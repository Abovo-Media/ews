package ews

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"io"
	"strconv"
	"sync/atomic"

	"github.com/Abovo-Media/go-ews/ewsxml"
	"github.com/go-pogo/errors"
	"github.com/go-pogo/writing"
)

const panicNilBody = "ews.NewRequest: body must be a non-nil value"

type Request struct {
	ctx  context.Context
	head *ewsxml.Header
	body interface{}
}

func NewRequest(ctx context.Context, head *ewsxml.Header, body interface{}) *Request {
	if ctx == nil {
		ctx = context.Background()
	}
	if head == nil {
		head = new(ewsxml.Header)
	}
	if body == nil {
		panic(panicNilBody)
	}

	return &Request{
		ctx:  context.WithValue(ctx, requestIdKey{}, genRequestId()),
		head: head,
		body: body,
	}
}

func (r *Request) Header() *ewsxml.Header { return r.head }
func (r *Request) Body() interface{}      { return r.body }

//goland:noinspection HttpUrlsUsage
var (
	soapStart = []byte(xml.Header + `<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
		xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"
		xmlns:m="http://schemas.microsoft.com/exchange/services/2006/messages"
		xmlns="http://schemas.microsoft.com/exchange/services/2006/types">`)

	soapBodyStart = []byte(`<soap:Body>`)
	soapEnd       = []byte(`</soap:Body></soap:Envelope>`)
)

func (r *Request) WriteTo(w io.Writer) (int64, error) {
	x, err := xml.Marshal(r.head)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	cw := writing.ToCountingStringWriter(w)
	_, _ = cw.Write(soapStart)
	_, _ = cw.Write(x)
	_, _ = cw.Write(soapBodyStart)

	switch b := r.body.(type) {
	case []byte:
		_, _ = cw.Write(b)

	case string:
		_, _ = cw.WriteString(b)

	case io.WriterTo:
		_, _ = b.WriteTo(cw)

	default:
		if x, err = xml.Marshal(b); err != nil {
			return 0, errors.WithStack(err)
		} else {
			_, _ = cw.Write(x)
		}
	}

	_, _ = cw.Write(soapEnd)
	return int64(cw.Count()), errors.Combine(cw.Errors()...)
}

type (
	requestIdKey struct{}
	attemptKey   struct{}
)

func RequestId(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(requestIdKey{}).(string)
	return v, ok
}

func RequestAttempt(ctx context.Context) (uint8, bool) {
	v, ok := ctx.Value(attemptKey{}).(uint8)
	return v, ok
}

var atomicRequestId atomic.Uint64

func genRequestId() string {
	uid := strconv.FormatUint(atomicRequestId.Add(1), 10)
	h := md5.Sum([]byte(uid))
	return hex.EncodeToString(h[:])
}
