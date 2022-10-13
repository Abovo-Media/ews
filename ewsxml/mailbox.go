package ewsxml

// The RoutingType element represents the routing protocol for the recipient.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/routingtype-emailaddress
type RoutingType string

func (s RoutingType) String() string { return string(s) }

// The MailboxType element represents the type of mailbox that is represented by
// the e-mail address.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxtype
type MailboxType string

func (s MailboxType) String() string { return string(s) }

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	RoutingType_Smtp RoutingType = "SMTP"
	RoutingType_EX   RoutingType = "EX"

	// MailboxType_Mailbox represents a mail-enabled Active Directory object.
	MailboxType_Mailbox MailboxType = "Mailbox"
	// MailboxType_PublicDL represents a public distribution list.
	MailboxType_PublicDL MailboxType = "PublicDL"
	// MailboxType_PrivateDL represents a private distribution list in a user's
	// mailbox.
	MailboxType_PrivateDL MailboxType = "PrivateDL"
	// MailboxType_Contact represents a contact in a user's mailbox.
	MailboxType_Contact MailboxType = "Contact"
	// MailboxType_PublicFolder represents a public folder.
	MailboxType_PublicFolder MailboxType = "PublicFolder"
	// MailboxType_Unknown represents an unknown type of mailbox.
	MailboxType_Unknown MailboxType = "Unknown"
	// MailboxType_OneOff represents a one-off member of a personal distribution
	// list.
	MailboxType_OneOff MailboxType = "OneOff"
	// MailboxType_GroupMailbox represents a group mailbox.
	MailboxType_GroupMailbox MailboxType = "GroupMailbox"
)

// The Mailbox element identifies a mail-enabled Active Directory object.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailbox
type Mailbox struct {
	// XMLName      xml.Name `xml:"http://schemas.microsoft.com/exchange/services/2006/types Mailbox"`
	Name         string `xml:",omitempty"`
	EmailAddress string
	RoutingType  RoutingType `xml:",omitempty"`
	MailboxType  MailboxType `xml:",omitempty"`
	ItemId       *ItemId     `xml:",omitempty"`
}

func EmailMailbox(email string) *Mailbox {
	return &Mailbox{EmailAddress: email}
}

// OneMailbox is a wrapper with only a single Mailbox element inside.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sender
// type OneMailbox struct {
// 	Mailbox Mailbox `xml:"Mailbox"`
// }
//
// func One(m Mailbox) *OneMailbox { return &OneMailbox{Mailbox: m} }

type Attendee struct {
	Mailbox Mailbox
	// <ResponseType/>
	// <LastResponseTime/>
}

type Attendees struct {
	Attendee []Attendee
}

func NewAttendees(a ...Attendee) *Attendees {
	return &Attendees{Attendee: a}
}

func (a *Attendees) AddEmailAddress(email ...string) *Attendees {
	for _, e := range email {
		a.Attendee = append(a.Attendee, Attendee{Mailbox: *EmailMailbox(e)})
	}
	return a
}
