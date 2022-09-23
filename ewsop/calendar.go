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
		CalendarView ewsxml.CalendarView

		// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/parentfolderids
		ParentFolderIds struct {
			DistinguishedFolderId ewsxml.DistinguishedFolderId
		} `xml:"m:ParentFolderIds"`
	}
}

type FindItemCalendarViewResponse struct {
	XMLName          xml.Name `xml:"FindItemResponse"`
	ResponseMessages struct {
		FindItemResponseMessage ewsxml.FindItemResponseMessage
	}
}

func GetCalendars(ctx context.Context, req ews.Requester, op *FindItemCalendarViewOperation) (*FindItemCalendarViewResponse, error) {
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
