package ewsxml

import (
	"encoding/xml"
)

// DistinguishedFolderId_Id identifies a default folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/distinguishedfolderid
type DistinguishedFolderId_Id string

func (d DistinguishedFolderId_Id) String() string { return string(d) }

const (
	// DistinguishedFolderId_Calendar represents the Calendar folder.
	DistinguishedFolderId_Calendar DistinguishedFolderId_Id = "calendar"
	// DistinguishedFolderId_Contacts represents the Contacts folder.
	DistinguishedFolderId_Contacts DistinguishedFolderId_Id = "contacts"
	// DistinguishedFolderId_DeletedItems represents the Deleted Items folder.
	DistinguishedFolderId_DeletedItems DistinguishedFolderId_Id = "deleteditems"
	// DistinguishedFolderId_Drafts represents the Drafts folder.
	DistinguishedFolderId_Drafts DistinguishedFolderId_Id = "drafts"
	// DistinguishedFolderId_Inbox represents the Inbox folder.
	DistinguishedFolderId_Inbox DistinguishedFolderId_Id = "inbox"
	// DistinguishedFolderId_Journal represents the Journal folder.
	DistinguishedFolderId_Journal DistinguishedFolderId_Id = "journal"
	// DistinguishedFolderId_Notes represents the Notes folder.
	DistinguishedFolderId_Notes DistinguishedFolderId_Id = "notes"
	// DistinguishedFolderId_Outbox represents the Outbox folder.
	DistinguishedFolderId_Outbox DistinguishedFolderId_Id = "outbox"
	// DistinguishedFolderId_SentItems represents the Sent Items folder.
	DistinguishedFolderId_SentItems DistinguishedFolderId_Id = "sentitems"
	// DistinguishedFolderId_Tasks represents the Tasks folder.
	DistinguishedFolderId_Tasks DistinguishedFolderId_Id = "tasks"
	// DistinguishedFolderId_MsgFolderRoot represents the message folder root.
	DistinguishedFolderId_MsgFolderRoot DistinguishedFolderId_Id = "msgfolderroot"
	// DistinguishedFolderId_Root represents the root of the mailbox.
	DistinguishedFolderId_Root DistinguishedFolderId_Id = "root"
	// DistinguishedFolderId_JunkEmail represents the Junk Email folder.
	DistinguishedFolderId_JunkEmail DistinguishedFolderId_Id = "junkemail"
	// DistinguishedFolderId_SearchFolders represents the Search Folders folder.
	DistinguishedFolderId_SearchFolders DistinguishedFolderId_Id = "searchfolders"
	// DistinguishedFolderId_Voicemail represents the Voice Mail folder.
	DistinguishedFolderId_Voicemail DistinguishedFolderId_Id = "voicemail"
	// DistinguishedFolderId_RecoverableItemsRoot represents the dumpster root folder.
	DistinguishedFolderId_RecoverableItemsRoot DistinguishedFolderId_Id = "recoverableitemsroot"
	// DistinguishedFolderId_RecoverableItemsDeletions represents the dumpster deletions folder.
	DistinguishedFolderId_RecoverableItemsDeletions DistinguishedFolderId_Id = "recoverableitemsdeletions"
	// DistinguishedFolderId_RecoverableItemsVersions represents the dumpster versions folder.
	DistinguishedFolderId_RecoverableItemsVersions DistinguishedFolderId_Id = "recoverableitemsversions"
	// DistinguishedFolderId_RecoverableItemsPurges represents the dumpster purges folder.
	DistinguishedFolderId_RecoverableItemsPurges DistinguishedFolderId_Id = "recoverableitemspurges"
	// DistinguishedFolderId_ArchiveRoot represents the root archive folder.
	DistinguishedFolderId_ArchiveRoot DistinguishedFolderId_Id = "archiveroot"
	// DistinguishedFolderId_ArchiveMsgFolderRoot represents the root archive message folder.
	DistinguishedFolderId_ArchiveMsgFolderRoot DistinguishedFolderId_Id = "archivemsgfolderroot"
	// DistinguishedFolderId_ArchiveDeletedItems represents the archive deleted items folder.
	DistinguishedFolderId_ArchiveDeletedItems DistinguishedFolderId_Id = "archivedeleteditems"
	// DistinguishedFolderId_ArchiveInbox represents the archive Inbox folder. Versions of Exchange starting with build number 15.00.0913.09 include this value.
	DistinguishedFolderId_ArchiveInbox DistinguishedFolderId_Id = "archiveinbox"
	// DistinguishedFolderId_ArchiveRecoverableItemsRoot represents the archive recoverable items root folder.
	DistinguishedFolderId_ArchiveRecoverableItemsRoot DistinguishedFolderId_Id = "archiverecoverableitemsroot"
	// DistinguishedFolderId_ArchiveRecoverableItemsDeletions represents the archive recoverable items deletions folder.
	DistinguishedFolderId_ArchiveRecoverableItemsDeletions DistinguishedFolderId_Id = "archiverecoverableitemsdeletions"
	// DistinguishedFolderId_ArchiveRecoverableItemsVersions represents the archive recoverable items versions folder.
	DistinguishedFolderId_ArchiveRecoverableItemsVersions DistinguishedFolderId_Id = "archiverecoverableitemsversions"
	// DistinguishedFolderId_ArchiveRecoverableItemsPurges represents the archive recoverable items purges folder.
	DistinguishedFolderId_ArchiveRecoverableItemsPurges DistinguishedFolderId_Id = "archiverecoverableitemspurges"
	// DistinguishedFolderId_SyncIssues represents the sync issues folder.
	DistinguishedFolderId_SyncIssues DistinguishedFolderId_Id = "syncissues"
	// DistinguishedFolderId_Conflicts represents the conflicts folder.
	DistinguishedFolderId_Conflicts DistinguishedFolderId_Id = "conflicts"
	// DistinguishedFolderId_LocalFailures represents the local failures folder.
	DistinguishedFolderId_LocalFailures DistinguishedFolderId_Id = "localfailures"
	// DistinguishedFolderId_ServerFailures represents the server failures folder.
	DistinguishedFolderId_ServerFailures DistinguishedFolderId_Id = "serverfailures"
	// DistinguishedFolderId_RecipientCache represents the recipient cache folder.
	DistinguishedFolderId_RecipientCache DistinguishedFolderId_Id = "recipientcache"
	// DistinguishedFolderId_QuickContacts represents the quick contacts folder.
	DistinguishedFolderId_QuickContacts DistinguishedFolderId_Id = "quickcontacts"
	// DistinguishedFolderId_ConversationHistory represents the conversation history folder.
	DistinguishedFolderId_ConversationHistory DistinguishedFolderId_Id = "conversationhistory"
	// DistinguishedFolderId_AdminAuditLogs represents the admin audit logs folder.
	DistinguishedFolderId_AdminAuditLogs DistinguishedFolderId_Id = "adminauditlogs"
	// DistinguishedFolderId_TodoSearch represents the todo search folder.
	DistinguishedFolderId_TodoSearch DistinguishedFolderId_Id = "todosearch"
	// DistinguishedFolderId_MyContacts represents the My Contacts folder.
	DistinguishedFolderId_MyContacts DistinguishedFolderId_Id = "mycontacts"
	// DistinguishedFolderId_Directory represents the directory folder.
	DistinguishedFolderId_Directory DistinguishedFolderId_Id = "directory"
	// DistinguishedFolderId_IMContactList represents the IM contact list folder.
	DistinguishedFolderId_IMContactList DistinguishedFolderId_Id = "imcontactlist"
	// DistinguishedFolderId_PeopleConnect represents the people connect folder.
	DistinguishedFolderId_PeopleConnect DistinguishedFolderId_Id = "peopleconnect"
	// DistinguishedFolderId_Favorites represents the Favorites folder.
	DistinguishedFolderId_Favorites DistinguishedFolderId_Id = "favorites"
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
// by name.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/distinguishedfolderid
type DistinguishedFolderId struct {
	Id        DistinguishedFolderId_Id `xml:",attr"`
	ChangeKey string                   `xml:",attr,omitempty"`
	Mailbox   *Mailbox                 `xml:",omitempty"`
}

func (d *DistinguishedFolderId) WithId(id DistinguishedFolderId_Id) *DistinguishedFolderId {
	d.Id = id
	return d
}

func (d *DistinguishedFolderId) WithMailbox(mb *Mailbox) *DistinguishedFolderId {
	d.Mailbox = mb
	return d
}

// The ParentFolderId element represents the identifier of the parent folder
// that contains the item or folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/parentfolderid
type ParentFolderId struct {
	XMLName   xml.Name `xml:"m:ParentFolderId"`
	Id        string   `xml:",attr"`
	ChangeKey string   `xml:",attr,omitempty"`
}
