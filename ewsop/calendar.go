package ewsop

import (
	"context"
	"encoding/xml"

	"github.com/Abovo-Media/go-ews"
	"github.com/Abovo-Media/go-ews/ewsxml"
)

type FindItemCalendarViewOperation struct {
	Header   ewsxml.Header
	FindItem struct {
		ewsxml.FindItem
		CalendarView    ewsxml.CalendarView
		ParentFolderIds ewsxml.ParentFolderIds
	}
}

type FindItemCalendarViewResponse struct {
	XMLName          xml.Name `xml:"FindItemResponse"`
	ResponseMessages struct {
		FindItemResponseMessage ewsxml.FindItemResponseMessage
	}
}

func (r *FindItemCalendarViewResponse) Response() *ewsxml.ResponseMessage {
	return r.ResponseMessages.FindItemResponseMessage.Response()
}

const OpGetCalendars Operation = "GetCalendars"

func GetCalendars(ctx context.Context, req ews.Requester, op *FindItemCalendarViewOperation) (*FindItemCalendarViewResponse, error) {
	ctx = setOperation(ctx, OpGetCalendars)

	if op.FindItem.Traversal == "" {
		op.FindItem.Traversal = ewsxml.Traversal_Shallow
	}
	if op.FindItem.ItemShape.BaseShape == "" {
		op.FindItem.ItemShape.BaseShape = ewsxml.BaseShape_Default
	}
	op.FindItem.ParentFolderIds.DistinguishedFolderId.Id = "calendar"

	var out FindItemCalendarViewResponse
	return &out, req.Request(ews.NewRequest(ctx, &op.Header, op.FindItem), &out)
}
