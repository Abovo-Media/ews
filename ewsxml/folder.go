package ewsxml

import (
	"encoding/xml"
)

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/parentfolderids
type ParentFolderIds struct {
	XMLName               xml.Name `xml:"m:ParentFolderIds"`
	DistinguishedFolderId DistinguishedFolderId
}

// The FolderId element contains the identifier and change key of a folder
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/folderid
type FolderId struct {
	Id        string `xml:",attr"`
	ChangeKey string `xml:",attr,omitempty"`
}

// The DistinguishedFolderId element identifies folders that can be referenced
// by name. If you do not use this element, you must use the FolderId element to
// identify a folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/distinguishedfolderid
type DistinguishedFolderId struct {
	Id        string   `xml:",attr"`
	ChangeKey string   `xml:",attr,omitempty"`
	Mailbox   *Mailbox `xml:",omitempty"`
}

func (d *DistinguishedFolderId) WithId(s string) {
	if d == nil {
		x := DistinguishedFolderId{}
		*d = x
	}
	d.Id = s
}

// func (d *DistinguishedFolderId) WithMailbox(mb *Mailbox) {
// 	if d == nil {
// 		*d = new(DistinguishedFolderId)
// 	}
// }

func NewDistinguishedFolderIdWithMailbox(mb *Mailbox) *DistinguishedFolderId {
	return &DistinguishedFolderId{Mailbox: mb}
}

// The ParentFolderId element represents the identifier of the parent folder
// that contains the item or folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/parentfolderid
type ParentFolderId struct {
	XMLName   xml.Name `xml:"m:ParentFolderId"`
	Id        string   `xml:",attr"`
	ChangeKey string   `xml:",attr,omitempty"`
}
