package ewsxml

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/message-ex15websvcsotherref
type Message struct {
	// MimeContent string
	// ItemId         ItemId
	// ParentFolderId ParentFolderId
	ItemClass   string
	Subject     string
	Sensitivity Sensitivity
	Body        Body
	// Attachments                  string
	// DateTimeReceived             string
	// Size                         string
	// Categories                   string
	// Importance                   string
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
	// ReminderIsSet                string
	// ReminderMinutesBeforeStart   string
	// DisplayCc                    string
	// DisplayTo                    string
	// HasAttachments               string
	// ExtendedProperty             string
	// Culture                      string
	Sender       Mailbox `xml:"Sender>Mailbox"`
	ToRecipients []Mailbox
	// CcRecipients                 string
	// BccRecipients                string
	// IsReadReceiptRequested       string
	// IsDeliveryReceiptRequested   string
	// ConversationIndex            string
	// ConversationTopic            string
	// From                         string
	// InternetMessageId            string
	// IsRead                       string
	// IsResponseRequested          string
	// References                   string
	// ReplyTo                      string
	// EffectiveRights              string
	// ReceivedBy                   string
	// ReceivedRepresenting         string
	// LastModifiedName             string
	// LastModifiedTime             string
	// IsAssociated                 string
	// WebClientReadFormQueryString string
	// WebClientEditFormQueryString string
	// ConversationId               string
	// UniqueBody                   string
	// ReminderMessageData          string
}

// The BodyType element identifies how the body text is formatted in the
// response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/bodytype
type BodyType string

func (s BodyType) String() string { return string(s) }

// The Body element specifies the body of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/body
type Body struct {
	BodyType    BodyType `xml:",attr"`
	IsTruncated bool     `xml:",attr"`
	Contents    []byte   `xml:",chardata"`
}

func (b *Body) Html(contents []byte) {
	b.BodyType = BodyType_HTML
	b.Contents = contents
}

func (b *Body) Text(contents []byte) {
	b.BodyType = BodyType_Text
	b.Contents = contents
}
