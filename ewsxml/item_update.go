package ewsxml

import (
	"encoding/xml"
)

type ConflictResolution string

func (s ConflictResolution) String() string { return string(s) }

type SendMeetingInvitationsOrCancellations string

func (s SendMeetingInvitationsOrCancellations) String() string { return string(s) }

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	// ConflictResolution_NeverOverwrite indicates if there is a conflict, the
	// update operation fails and an error is returned.
	ConflictResolution_NeverOverwrite ConflictResolution = "NeverOverwrite"
	// ConflictResolution_AutoResolve indicates the update operation
	// automatically resolves any conflict.
	ConflictResolution_AutoResolve ConflictResolution = "AutoResolve"
	// ConflictResolution_AlwaysOverwrite indicates if there is a conflict, the
	// update operation will overwrite information.
	ConflictResolution_AlwaysOverwrite ConflictResolution = "AlwaysOverwrite"

	// SendMeetingInvitationsOrCancellations_SendToNone indicates the calendar
	// item is updated but updates are not sent to attendees.
	SendMeetingInvitationsOrCancellations_SendToNone SendMeetingInvitationsOrCancellations = "SendToNone"
	// SendMeetingInvitationsOrCancellations_SendOnlyToAll indicates the
	// calendar item is updated and the meeting update is sent to all attendees
	// but is not saved in the Sent Items folder.
	SendMeetingInvitationsOrCancellations_SendOnlyToAll SendMeetingInvitationsOrCancellations = "SendOnlyToAll"
	// SendMeetingInvitationsOrCancellations_SendOnlyToChanged indicates the
	// calendar item is updated and the meeting update is sent only to
	// attendees that are affected by the change in the meeting.
	SendMeetingInvitationsOrCancellations_SendOnlyToChanged SendMeetingInvitationsOrCancellations = "SendOnlyToChanged"
	// SendMeetingInvitationsOrCancellations_SendToAllAndSaveCopy indicates the
	// calendar item is updated, the meeting update is sent to all attendees,
	// and a copy is saved in the Sent Items folder.
	SendMeetingInvitationsOrCancellations_SendToAllAndSaveCopy SendMeetingInvitationsOrCancellations = "SendToAllAndSaveCopy"
	// SendMeetingInvitationsOrCancellations_SendToChangedAndSaveCopy indicates
	// the calendar item is updated, the meeting update is sent to all
	// attendees that are affected by the change in the meeting, and a copy is
	// saved in the Sent Items folder.
	SendMeetingInvitationsOrCancellations_SendToChangedAndSaveCopy SendMeetingInvitationsOrCancellations = "SendToChangedAndSaveCopy"
)

// The UpdateItem element defines a request to update an item in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updateitem
type UpdateItem struct {
	XMLName                               xml.Name                              `xml:"m:UpdateItem"`
	ConflictResolution                    ConflictResolution                    `xml:",attr"`
	MessageDisposition                    MessageDisposition                    `xml:",attr"`
	SendMeetingInvitationsOrCancellations SendMeetingInvitationsOrCancellations `xml:",attr"`
	SuppressReadReceipts                  bool                                  `xml:",attr"`

	SavedItemFolderId SavedItemFolderId
	ItemChanges       []ItemChange `xml:"m:ItemChanges"`
}

// The ItemChange element contains an item identifier and the updates to apply
// to the item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemchange
type ItemChange struct {
	XMLName               xml.Name `xml:"m:ItemChange"`
	ItemId                ItemId
	OccurrenceItemId      OccurrenceItemId
	RecurringMasterItemId RecurringMasterItemId
	Updates               Updates
}

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updates-item
type Updates struct {
}
