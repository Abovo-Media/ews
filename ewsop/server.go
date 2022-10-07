package ewsop

import (
	"context"

	ews "github.com/Abovo-Media/go-ews"
	"github.com/Abovo-Media/go-ews/ewsxml"
)

type GetServerTimeZonesOperation struct {
	Header             ewsxml.Header
	GetServerTimeZones ewsxml.GetServerTimeZones
}

// todo: finish response
type GetServerTimeZonesResponse struct {
	ewsxml.ResponseMessage
}

const OpGetServerTimeZones Operation = "GetServerTimeZones"

func GetServerTimeZones(ctx context.Context, req ews.Requester, op *GetServerTimeZonesOperation) (*GetServerTimeZonesResponse, error) {
	var out GetServerTimeZonesResponse
	return &out, req.Request(
		ews.NewRequest(setOperation(ctx, OpGetServerTimeZones), &op.Header, &op.GetServerTimeZones),
		&out,
	)
}
