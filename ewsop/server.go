package ewsop

import (
	"context"

	ews "github.com/Abovo-Media/go-ews"
	"github.com/Abovo-Media/go-ews/ewsxml"
)

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getservertimezones-operation
type GetServerTimeZonesOperation struct {
	Header             ewsxml.Header
	GetServerTimeZones ewsxml.GetServerTimeZones
}

type GetServerTimeZonesResponse struct {
	ResponseMessages struct {
		GetServerTimeZonesResponseMessage struct {
			ewsxml.ResponseMessage
			TimeZoneDefinitions []ewsxml.TimeZoneDefinition
		}
	}
}

func (r *GetServerTimeZonesResponse) Response() *ewsxml.ResponseMessage {
	return r.ResponseMessages.GetServerTimeZonesResponseMessage.Response()
}

const OpGetServerTimeZones Operation = "GetServerTimeZones"

func GetServerTimeZones(ctx context.Context, req ews.Requester, op *GetServerTimeZonesOperation) (*GetServerTimeZonesResponse, error) {
	var out GetServerTimeZonesResponse
	return &out, req.Request(
		ews.NewRequest(setOperation(ctx, OpGetServerTimeZones), &op.Header, &op.GetServerTimeZones),
		&out,
	)
}
