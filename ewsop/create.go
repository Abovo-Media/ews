package ewsop

import (
	"context"

	"github.com/Abovo-Media/go-ews"
	"github.com/Abovo-Media/go-ews/ewsxml"
)

type CreateItemOperation struct {
	Header     ewsxml.Header
	CreateItem ewsxml.CreateItem
}

type CreateItemResponse struct {
	ResponseMessages []struct {
		ewsxml.Response
		CalendarItems []struct {
			ItemId ewsxml.ItemId
		} `xml:"Items>CalendarItem"`
	} `xml:"ResponseMessages>CreateItemResponseMessage"`
}

func CreateCalendarItem(ctx context.Context, req ews.Requester, op *CreateItemOperation, ci ...ewsxml.CalendarItem) (*CreateItemResponse, error) {
	if op == nil {
		op = new(CreateItemOperation)
	}
	if op.CreateItem.SendMeetingInvitations == "" {
		op.CreateItem.SendMeetingInvitations = ewsxml.SendMeetingInvitations_SendToAllAndSaveCopy
	}

	if op.CreateItem.SavedItemFolderId.FolderId == nil {
		if op.CreateItem.SavedItemFolderId.DistinguishedFolderId == nil {
			op.CreateItem.SavedItemFolderId.DistinguishedFolderId = new(ewsxml.DistinguishedFolderId)
		}
		op.CreateItem.SavedItemFolderId.DistinguishedFolderId.Id = "calendar"
	}

	op.CreateItem.Items.CalendarItem = append(op.CreateItem.Items.CalendarItem, ci...)

	var out CreateItemResponse
	return &out, req.Request(ews.NewRequest(ctx, &op.Header, op.CreateItem), &out)
}
