package ewsxml

import (
	"encoding/xml"
)

type SearchScope string

func (s SearchScope) String() string { return string(s) }

type ContactDataShape string

func (s ContactDataShape) String() string { return string(s) }

const (
	// SearchScope_ActiveDirectory indicates only the Active Directory
	// directory service is searched.
	SearchScope_ActiveDirectory SearchScope = "ActiveDirectory"
	// SearchScope_ActiveDirectoryContacts indicates Active Directory is
	// searched first, and then the contact folders that are specified in the
	// ParentFolderIds property are searched.
	SearchScope_ActiveDirectoryContacts SearchScope = "ActiveDirectoryContacts"
	// SearchScope_Contacts indicates only the contact folders that are
	// identified by the ParentFolderIds property are searched.
	SearchScope_Contacts SearchScope = "Contacts"
	// SearchScope_ContactsActiveDirectory indicates Contact folders that are
	// identified by the ParentFolderIds property are searched first and then
	// Active Directory is searched.
	SearchScope_ContactsActiveDirectory SearchScope = "ContactsActiveDirectory"

	// ContactDataShape_IdOnly indicates the contact item identifier property
	// is returned.
	ContactDataShape_IdOnly
	// ContactDataShape_Default indicates the Default set of contact item
	// properties is returned. For more information, see Response shapes in EWS.
	ContactDataShape_Default
	// ContactDataShape_AllProperties indicates the AllProperties set of
	// contact item properties are returned.For more information, see Response
	// shapes in EWS.
	ContactDataShape_AllProperties
)

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/resolvenames
type ResolveNames struct {
	XMLName               xml.Name         `xml:"ResolveNames"`
	ReturnFullContactData bool             `xml:",attr"`
	SearchScope           SearchScope      `xml:",attr"`
	ContactDataShape      ContactDataShape `xml:",attr"`
	ParentFolderIds       []string         `xml:",omitempty"`
	UnresolvedEntry       string           `xml:",omitempty"`
}
