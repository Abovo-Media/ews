package ewsop

import (
	"context"

	"github.com/Abovo-Media/go-ews"
	"github.com/Abovo-Media/go-ews/ewsxml"
)

type GetServerTimeZonesOperation struct {
	Header             ewsxml.Header
	GetServerTimeZones ewsxml.GetServerTimeZones
}

type GetServerTimeZonesResponse struct {
	ewsxml.Response
	RoomLists struct {
		Address []ewsxml.Mailbox `xml:"Address"`
	} `xml:"RoomLists"`
}

func GetServerTimeZones(ctx context.Context, req ews.Requester, op *GetServerTimeZonesOperation) (*GetServerTimeZonesResponse, error) {
	var out GetServerTimeZonesResponse
	return &out, req.Request(ews.NewRequest(ctx, &op.Header, &op.GetServerTimeZones), &out)
}
