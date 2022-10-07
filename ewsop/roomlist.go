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
	ewsxml.ResponseMessage
	RoomLists struct {
		Address []ewsxml.Mailbox `xml:"Address"`
	} `xml:"RoomLists"`
}

const OpGetRoomLists Operation = "GetRoomLists"

func GetRoomLists(ctx context.Context, req ews.Requester, op *GetRoomListsOperation) (*GetRoomListsResponse, error) {
	var out GetRoomListsResponse
	return &out, req.Request(
		ews.NewRequest(setOperation(ctx, OpGetRoomLists), &op.Header, op.body),
		&out,
	)
}
