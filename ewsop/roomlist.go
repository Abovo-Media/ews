package ewsop

import (
	"context"
	"encoding/xml"

	"github.com/Abovo-Media/go-ews"
	"github.com/Abovo-Media/go-ews/ewsxml"
)

type GetRoomListsOperation struct {
	Header ewsxml.Header
	body   struct {
		XMLName xml.Name `xml:"m:GetRoomLists"`
	}
}

type GetRoomListsResponse struct {
	ewsxml.Response
	RoomLists struct {
		Address []ewsxml.Mailbox `xml:"Address"`
	} `xml:"RoomLists"`
}

func GetRoomLists(ctx context.Context, req ews.Requester, op *GetRoomListsOperation) (*GetRoomListsResponse, error) {
	var out GetRoomListsResponse
	return &out, req.Request(ews.NewRequest(ctx, &op.Header, op.body), &out)
}
