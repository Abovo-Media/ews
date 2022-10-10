package ewsxml

import (
	"time"
)

// The Sensitivity element indicates the sensitivity level of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sensitivity
type Sensitivity string

func (s Sensitivity) String() string { return string(s) }

// The LegacyFreeBusyStatus element represents the free/busy status of the
// calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/legacyfreebusystatus
type LegacyFreeBusyStatus string

func (s LegacyFreeBusyStatus) String() string { return string(s) }

// The CalendarItemType element represents the type of a calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendaritemtype
type CalendarItemType string

func (s CalendarItemType) String() string { return string(s) }

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	Sensitivity_Normal       Sensitivity = "Normal"
	Sensitivity_Personal     Sensitivity = "Personal"
	Sensitivity_Private      Sensitivity = "Private"
	Sensitivity_Confidential Sensitivity = "Confidential"

	LegacyFreeBusyStatus_Free             LegacyFreeBusyStatus = "Free"
	LegacyFreeBusyStatus_Tentative        LegacyFreeBusyStatus = "Tentative"
	LegacyFreeBusyStatus_Busy             LegacyFreeBusyStatus = "Busy"
	LegacyFreeBusyStatus_OOF              LegacyFreeBusyStatus = "OOF"
	LegacyFreeBusyStatus_WorkingElsewhere LegacyFreeBusyStatus = "WorkingElsewhere"
	LegacyFreeBusyStatus_NoData           LegacyFreeBusyStatus = "NoData"

	// CalendarItemType_Single indicates the item is not associated with a
	// recurring calendar item.
	CalendarItemType_Single CalendarItemType = "Single"
	// CalendarItemType_Occurrence indicates the item is an occurrence of a
	// recurring calendar item.
	CalendarItemType_Occurrence CalendarItemType = "Occurrence"
	// CalendarItemType_Exception indicates the item is an exception to a
	// recurring calendar item.
	CalendarItemType_Exception CalendarItemType = "Exception"
	// CalendarItemType_RecurringMaster indicates the item is master for a set
	// of recurring calendar items.
	CalendarItemType_RecurringMaster CalendarItemType = "RecurringMaster"
)

// The CalendarItem element represents an Exchange calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendaritem
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createitem-operation-calendar-item
type CalendarItem struct {
	// MimeContent                  string
	ItemId         *ItemId `xml:",omitempty"`
	ParentFolderId *ItemId `xml:",omitempty"`
	// ItemClass                    string
	Subject string
	// Sensitivity *Sensitivity
	Body *Body `xml:",omitempty"`
	// Attachments                  string
	// DateTimeReceived             string
	// Size                         string
	// Categories                   string
	// InReplyTo                    string
	// IsSubmitted                  string
	// IsDraft                      string
	// IsFromMe                     string
	// IsResend                     string
	// IsUnmodified                 string
	// InternetMessageHeaders       string
	// DateTimeSent                 string
	// DateTimeCreated              string
	// ResponseObjects              string
	// ReminderDueBy                string
	ReminderIsSet              bool               `xml:",omitempty"`
	ReminderMinutesBeforeStart Minutes            `xml:",omitempty"`
	DisplayCc                  ConcatenatedString `xml:",omitempty"`
	DisplayTo                  ConcatenatedString `xml:",omitempty"`
	// HasAttachments             bool
	// ExtendedProperty             string
	// Culture                      string
	Start time.Time
	End   time.Time
	// OriginalStart                string
	IsAllDayEvent        bool
	LegacyFreeBusyStatus *LegacyFreeBusyStatus `xml:",omitempty"`
	Location             *string               `xml:",omitempty"`
	// When                         string
	// IsMeeting                    string
	// IsCancelled                  string
	// IsRecurring                  string
	// MeetingRequestWasSent        string
	// IsResponseRequested          string
	CalendarItemType CalendarItemType `xml:",omitempty"`
	// MyResponseType               string
	Organizer         *Mailbox   `xml:"Organizer>Mailbox,omitempty"`
	RequiredAttendees *Attendees `xml:",omitempty"`
	OptionalAttendees *Attendees `xml:",omitempty"`
	// Resources         []Attendee // []Attendees
	// ConflictingMeetingCount      string
	// AdjacentMeetingCount         string
	// ConflictingMeetings          string
	// AdjacentMeetings             string
	// Duration                     string
	// TimeZone                     string
	// AppointmentReplyTime         string
	// AppointmentSequenceNumber    string
	// AppointmentState             string
	// Recurrence                   string
	// FirstOccurrence              string
	// LastOccurrence               string
	// ModifiedOccurrences          string
	// DeletedOccurrences           string
	// MeetingTimeZone              string
	// StartTimeZone                string
	// EndTimeZone                  string
	// ConferenceType               string
	// AllowNewTimeProposal         string
	// IsOnlineMeeting              string
	// MeetingWorkspaceUrl          string
	// NetShowUrl                   string
	// EffectiveRights              string
	// LastModifiedName             string
	// LastModifiedTime             string
	// IsAssociated bool
	// WebClientReadFormQueryString string
	// WebClientEditFormQueryString string
	// ConversationId               string
	// UniqueBody                   string
}

func (ci *CalendarItem) SetReminder(d time.Duration) {
	ci.ReminderIsSet = true
	ci.ReminderMinutesBeforeStart = Minutes(d)
}
