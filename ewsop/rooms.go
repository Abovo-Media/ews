package ewsop

import (
	"context"

	"github.com/Abovo-Media/go-ews"
	"github.com/Abovo-Media/go-ews/ewsxml"
)

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getrooms-operation
type GetRoomsOperation struct {
	Header   ewsxml.Header
	GetRooms ewsxml.GetRooms
}

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getroomsresponse
type GetRoomsResponse struct {
	ewsxml.Response
	Rooms []struct {
		Name         string
		EmailAddress string
		RoutingType  ewsxml.RoutingType
		MailboxType  ewsxml.MailboxType
	} `xml:"Rooms>Room>Id"`
}

func GetRooms(ctx context.Context, req ews.Requester, op *GetRoomsOperation) (*GetRoomsResponse, error) {
	var out GetRoomsResponse
	return &out, req.Request(ews.NewRequest(ctx, &op.Header, op.GetRooms), &out)
}
