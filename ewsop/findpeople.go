package ewsop

import (
	"context"

	"github.com/Abovo-Media/go-ews"
	"github.com/Abovo-Media/go-ews/ewsxml"
)

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/findpeople-operation
type FindPeopleOperation struct {
	Header     ewsxml.Header
	FindPeople ewsxml.FindPeople
}

type FindPeopleResponse struct {
	ewsxml.ResponseMessage
	People []ewsxml.Persona `xml:"People>Persona"`
}

const OpFindPeople Operation = "FindPeople"

func FindPeople(ctx context.Context, req ews.Requester, op *FindPeopleOperation) (*FindPeopleResponse, error) {
	ctx = setOperation(ctx, OpFindPeople)

	if op.FindPeople.PersonaShape != nil && op.FindPeople.PersonaShape.BaseShape == "" {
		op.FindPeople.PersonaShape.BaseShape = ewsxml.BaseShape_Default
	}
	if op.FindPeople.IndexedPageItemView.BasePoint == "" {
		op.FindPeople.IndexedPageItemView.BasePoint = ewsxml.BasePoint_Beginning
	}

	var out FindPeopleResponse
	return &out, req.Request(ews.NewRequest(ctx, &op.Header, op.FindPeople), &out)
}
