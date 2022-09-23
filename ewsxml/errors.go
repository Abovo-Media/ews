package ewsxml

// The ResponseCode element provides status information about the request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/responsecode
type ResponseCode string

func (r ResponseCode) String() string { return string(r) }

//goland:noinspection GoUnusedConst
const (
	// NoError indicates no error occurred for the request.
	NoError ResponseCode = "NoError"

	// ErrorAccessDenied occurs when the calling account does not have the
	// rights to perform the requested action.
	ErrorAccessDenied ResponseCode = "ErrorAccessDenied"

	// ErrorAccessModeSpecified is for internal use only. This error is not
	// returned.
	ErrorAccessModeSpecified ResponseCode = "ErrorAccessModeSpecified"

	// ErrorAccountDisabled occurs when the account in question has been
	// disabled.
	ErrorAccountDisabled ResponseCode = "ErrorAccountDisabled"

	// ErrorAddDelegatesFailed occurs when a list with added delegates cannot
	// be saved.
	ErrorAddDelegatesFailed ResponseCode = "ErrorAddDelegatesFailed"

	// ErrorAddressSpaceNotFound occurs when the address space record, or
	// Domain Name System (DNS) domain name, for cross-forest availability
	// could not be found in the Active Directory database.
	ErrorAddressSpaceNotFound ResponseCode = "ErrorAddressSpaceNotFound"

	// ErrorADOperation occurs when the operation failed because of
	// communication problems with Active Directory Domain Services (AD DS).
	ErrorADOperation ResponseCode = "ErrorADOperation"

	// ErrorADSessionFilter is returned when a ResolveNames operation request
	// specifies a name that is not valid.
	ErrorADSessionFilter ResponseCode = "ErrorADSessionFilter"

	// ErrorADUnavailable occurs when AD DS is unavailable. Try your request
	// again later.
	ErrorADUnavailable ResponseCode = "ErrorADUnavailable"

	// ErrorAffectedTaskOccurrencesRequired indicates that the
	// AffectedTaskOccurrences attribute was not specified. When the
	// DeleteItem element is used to delete at least one item that is a task,
	// and regardless of whether that task is recurring or not, the
	// AffectedTaskOccurrences attribute has to be specified so that
	// DeleteItem can determine whether to delete the current occurrence
	// or the entire series.
	ErrorAffectedTaskOccurrencesRequired ResponseCode = "ErrorAffectedTaskOccurrencesRequired"

	// ErrorArchiveFolderPathCreation indicates an error in archive folder path
	// creation.
	ErrorArchiveFolderPathCreation ResponseCode = "ErrorArchiveFolderPathCreation"

	// ErrorArchiveMailboxNotEnabled indicates that the archive mailbox was not
	// enabled.
	ErrorArchiveMailboxNotEnabled ResponseCode = "ErrorArchiveMailboxNotEnabled"

	// ErrorArchiveMailboxServiceDiscoveryFailed indicates that archive mailbox
	// service discovery failed.
	// This error was introduced in Exchange 2013.
	ErrorArchiveMailboxServiceDiscoveryFailed ResponseCode = "ErrorArchiveMailboxServiceDiscoveryFailed"

	// ErrorAttachmentNestLevelLimitExceeded specifies that an attempt was made
	// to create an item with more than 10 nested attachments. This value was
	// introduced in Exchange LogClientStart 2010 Service Pack 2 (SP2).
	ErrorAttachmentNestLevelLimitExceeded ResponseCode = "ErrorAttachmentNestLevelLimitExceeded"

	// ErrorAttachmentSizeLimitExceeded is returned by the CreateAttachment
	// element if an attempt is made to create an attachment with size
	// exceeding TryInt32.MaxValue, in bytes. The GetAttachment element returns
	// this error if an attempt to retrieve an existing attachment with size
	// exceeding TryInt32.MaxValue, in bytes.
	ErrorAttachmentSizeLimitExceeded ResponseCode = "ErrorAttachmentSizeLimitExceeded"

	// ErrorAutoDiscoverFailed indicates that Exchange Web Services tried to
	// determine the location of a cross-forest computer that is running
	// Exchange 2010 that has the Client Access server role installed by using
	// the Autodiscover service, but the call to the Autodiscover service
	// failed.
	ErrorAutoDiscoverFailed ResponseCode = "ErrorAutoDiscoverFailed"

	// ErrorAvailabilityConfigNotFound indicates that the availability
	// configuration information for the local forest is missing from AD DS.
	ErrorAvailabilityConfigNotFound ResponseCode = "ErrorAvailabilityConfigNotFound"

	// ErrorBatchProcessingStopped indicates that an exception occurred while
	// processing an item and that exception is likely to occur for the items
	// that follow. Requests may include multiple items; for example, a GetItem
	// operation request might include multiple identifiers. In general, items
	// are processed one at a time. If an exception occurs while processing an
	// item and that exception is likely to occur for the items that follow,
	// items that follow will not be processed. The following are examples of
	// errors that will stop processing for items that follow:
	//  - ErrorAccessDenied
	//  - ErrorAccountDisabled
	//  - ErrorADUnavailable
	//  - ErrorADOperation
	//  - ErrorConnectionFailed
	//  - ErrorMailboxStoreUnavailable
	//  - ErrorMailboxMoveInProgress
	//  - ErrorPasswordChangeRequired
	//  - ErrorPasswordExpired
	//  - ErrorQuotaExceeded
	//  - ErrorInsufficientResources
	ErrorBatchProcessingStopped ResponseCode = "ErrorBatchProcessingStopped"

	// ErrorCalendarCannotMoveOrCopyOccurrence occurs when an attempt is made
	// to move or copy an occurrence of a recurring calendar item.
	ErrorCalendarCannotMoveOrCopyOccurrence ResponseCode = "ErrorCalendarCannotMoveOrCopyOccurrence"

	// ErrorCalendarCannotUpdateDeletedItem occurs when an attempt is made to
	// update a calendar item that is located in the Deleted Items folder and
	// when meeting updates or cancellations are to be sent according to the
	// value of the SendMeetingInvitationsOrCancellations attribute.
	// The following are the possible values for this attribute:
	//  - SendToAllAndSaveCopy (SendMeetingInvitationsOrCancellations_SendToAllAndSaveCopy)
	//  - SendToChangedAndSaveCopy (SendMeetingInvitationsOrCancellations_SendToChangedAndSaveCopy)
	//  - SendOnlyToAll (SendMeetingInvitationsOrCancellations_SendOnlyToAll)
	//  - SendOnlyToChanged (SendMeetingInvitationsOrCancellations_SendOnlyToChanged)
	// However, such an update is allowed only when the value of this attribute
	// is set to SendToNone (SendMeetingInvitationsOrCancellations_SendToNone).
	ErrorCalendarCannotUpdateDeletedItem ResponseCode = "ErrorCalendarCannotUpdateDeletedItem"

	// ErrorCalendarCannotUseIdForOccurrenceId occurs when the UpdateItem,
	// GetItem, DeleteItem, MoveItem, CopyItem, or SendItem operation is called
	// and the ID that was specified is not an occurrence ID of any recurring
	// calendar item.
	ErrorCalendarCannotUseIdForOccurrenceId ResponseCode = "ErrorCalendarCannotUseIdForOccurrenceId"

	// ErrorCalendarCannotUseIdForRecurringMasterId occurs when the UpdateItem,
	// GetItem, DeleteItem, MoveItem, CopyItem, or SendItem operation is called
	// and the ID that was specified is not an ID of any recurring master item.
	ErrorCalendarCannotUseIdForRecurringMasterId ResponseCode = "ErrorCalendarCannotUseIdForRecurringMasterId"

	// ErrorCalendarDurationIsTooLong occurs during a CreateItem or UpdateItem
	// operation when a calendar item duration is longer than the maximum
	// allowed, which is currently 5 years.
	ErrorCalendarDurationIsTooLong ResponseCode = "ErrorCalendarDurationIsTooLong"

	// ErrorCalendarEndDateIsEarlierThanStartDate occurs when a calendar End
	// time is set to the same time or after the Start time.
	ErrorCalendarEndDateIsEarlierThanStartDate ResponseCode = "ErrorCalendarEndDateIsEarlierThanStartDate"

	// ErrorCalendarFolderIsInvalidForCalendarView occurs when the specified
	// folder for a FindItem operation with a CalendarView element is not of
	// calendar folder type.
	ErrorCalendarFolderIsInvalidForCalendarView ResponseCode = "ErrorCalendarFolderIsInvalidForCalendarView"

	// This response code is not used.
	// ErrorCalendarInvalidAttributeValue ResponseCode = "ErrorCalendarInvalidAttributeValue"

	// ErrorCalendarInvalidDayForTimeChangePattern occurs during a CreateItem
	// or UpdateItem operation when invalid values of Day, WeekendDay, and
	// Weekday are used to define the time change pattern.
	ErrorCalendarInvalidDayForTimeChangePattern ResponseCode = "ErrorCalendarInvalidDayForTimeChangePattern"

	// ErrorCalendarInvalidDayForWeeklyRecurrence occurs during a CreateItem or
	// UpdateItem operation when invalid values of Day, WeekDay, and WeekendDay
	// are used to specify the weekly recurrence.
	ErrorCalendarInvalidDayForWeeklyRecurrence ResponseCode = "ErrorCalendarInvalidDayForWeeklyRecurrence"

	// ErrorCalendarInvalidPropertyState occurs when the state of a calendar
	// item recurrence binary large object (BLOB) in the Exchange store is
	// invalid.
	ErrorCalendarInvalidPropertyState ResponseCode = "ErrorCalendarInvalidPropertyState"

	// This response code is not used.
	// ErrorCalendarInvalidPropertyValue ResponseCode = "ErrorCalendarInvalidPropertyValue"

	// ErrorCalendarInvalidRecurrence occurs when the specified recurrence
	// cannot be created.
	ErrorCalendarInvalidRecurrence ResponseCode = "ErrorCalendarInvalidRecurrence"

	// ErrorCalendarInvalidTimeZone occurs when an invalid time zone is
	// encountered.
	ErrorCalendarInvalidTimeZone ResponseCode = "ErrorCalendarInvalidTimeZone"

	// ErrorCalendarIsCancelledForAccept indicates that a calendar item has
	// been canceled.
	ErrorCalendarIsCancelledForAccept ResponseCode = "ErrorCalendarIsCancelledForAccept"

	// ErrorCalendarIsCancelledForDecline indicates that a calendar item has
	// been canceled.
	ErrorCalendarIsCancelledForDecline ResponseCode = "ErrorCalendarIsCancelledForDecline"

	// ErrorCalendarIsCancelledForRemove indicates that a calendar item has
	// been canceled.
	ErrorCalendarIsCancelledForRemove ResponseCode = "ErrorCalendarIsCancelledForRemove"

	// ErrorCalendarIsCancelledForTentative indicates that a calendar item has
	// been canceled.
	ErrorCalendarIsCancelledForTentative ResponseCode = "ErrorCalendarIsCancelledForTentative"

	// ErrorCalendarIsDelegatedForAccept indicates that the AcceptItem element
	// is invalid for a calendar item or meeting request in a delegated
	// scenario.
	ErrorCalendarIsDelegatedForAccept ResponseCode = "ErrorCalendarIsDelegatedForAccept"

	// ErrorCalendarIsDelegatedForDecline indicates that the DeclineItem
	// element is invalid for a calendar item or meeting request in a delegated
	// scenario.
	ErrorCalendarIsDelegatedForDecline ResponseCode = "ErrorCalendarIsDelegatedForDecline"

	// ErrorCalendarIsDelegatedForRemove indicates that the RemoveItem element
	// is invalid for a meeting cancellation in a delegated scenario.
	ErrorCalendarIsDelegatedForRemove ResponseCode = "ErrorCalendarIsDelegatedForRemove"

	// ErrorCalendarIsDelegatedForTentative indicates that the
	// TentativelyAcceptItem element is invalid for a calendar item or meeting
	// request in a delegated scenario.
	ErrorCalendarIsDelegatedForTentative ResponseCode = "ErrorCalendarIsDelegatedForTentative"

	// ErrorCalendarIsNotOrganizer indicates that the operation (currently
	// CancelItem) on the calendar item is not valid for an attendee. Only the
	// meeting organizer can cancel the meeting.
	ErrorCalendarIsNotOrganizer ResponseCode = "ErrorCalendarIsNotOrganizer"

	// ErrorCalendarIsOrganizerForAccept indicates that AcceptItem is invalid
	// for the organizer's calendar item.
	ErrorCalendarIsOrganizerForAccept ResponseCode = "ErrorCalendarIsOrganizerForAccept"

	// ErrorCalendarIsOrganizerForDecline indicates that DeclineItem is invalid
	// for the organizer's calendar item.
	ErrorCalendarIsOrganizerForDecline ResponseCode = "ErrorCalendarIsOrganizerForDecline"

	// ErrorCalendarIsOrganizerForRemove indicates that RemoveItem is invalid
	// for the organizer's calendar item. To remove a meeting from the calendar,
	// the organizer must use CancelCalendarItem.
	ErrorCalendarIsOrganizerForRemove ResponseCode = "ErrorCalendarIsOrganizerForRemove"

	// ErrorCalendarIsOrganizerForTentative indicates that
	// TentativelyAcceptItem is invalid for the organizer's calendar item.
	ErrorCalendarIsOrganizerForTentative ResponseCode = "ErrorCalendarIsOrganizerForTentative"

	// ErrorCalendarMeetingRequestIsOutOfDate indicates that a meeting request
	// is out-of-date and cannot be updated.
	ErrorCalendarMeetingRequestIsOutOfDate ResponseCode = "ErrorCalendarMeetingRequestIsOutOfDate"

	// ErrorCalendarOccurrenceIndexIsOutOfRecurrenceRange indicates that the
	// occurrence index does not point to an occurrence within the current
	// recurrence. For example, if your recurrence pattern defines a set of
	// three meeting occurrences and you try to access the fifth occurrence,
	// this response code will result.
	ErrorCalendarOccurrenceIndexIsOutOfRecurrenceRange ResponseCode = "ErrorCalendarOccurrenceIndexIsOutOfRecurrenceRange"

	// ErrorCalendarOccurrenceIsDeletedFromRecurrence indicates that any
	// operation on a deleted occurrence (addressed via recurring master ID and
	// occurrence index) is invalid.
	ErrorCalendarOccurrenceIsDeletedFromRecurrence ResponseCode = "ErrorCalendarOccurrenceIsDeletedFromRecurrence"

	// ErrorCalendarOutOfRange is reported on CreateItem and UpdateItem
	// operations for calendar items or task recurrence properties when the
	// property value is out of range. For example, specifying the fifteenth
	// week of the month will result in this response code.
	ErrorCalendarOutOfRange ResponseCode = "ErrorCalendarOutOfRange"

	// ErrorCalendarViewRangeTooBig occurs when Start to End range for the
	// CalendarView element is more than the maximum allowed, currently 2 years.
	ErrorCalendarViewRangeTooBig ResponseCode = "ErrorCalendarViewRangeTooBig"

	// ErrorCallerIsInvalidADAccount indicates that the requesting account is
	// not a valid account in the directory database.
	ErrorCallerIsInvalidADAccount ResponseCode = "ErrorCallerIsInvalidADAccount"

	// ErrorCannotArchiveCalendarContactTaskFolderException indicates that an
	// attempt was made to archive a calendar contact task folder.
	ErrorCannotArchiveCalendarContactTaskFolderException ResponseCode = "ErrorCannotArchiveCalendarContactTaskFolderException"

	// ErrorCannotArchiveItemsInPublicFolders indicates that an attempt was
	// made to archive items in public folders.
	ErrorCannotArchiveItemsInPublicFolders ResponseCode = "ErrorCannotArchiveItemsInPublicFolders"

	// ErrorCannotArchiveItemsInArchiveMailbox indicates that attempt was made
	// to archive items in the archive mailbox.
	ErrorCannotArchiveItemsInArchiveMailbox ResponseCode = "ErrorCannotArchiveItemsInArchiveMailbox"

	// ErrorCannotCreateCalendarItemInNonCalendarFolder occurs when a calendar
	// item is being created and the SavedItemFolderId attribute refers to a
	// non-calendar folder.
	ErrorCannotCreateCalendarItemInNonCalendarFolder ResponseCode = "ErrorCannotCreateCalendarItemInNonCalendarFolder"

	// ErrorCannotCreateContactInNonContactFolder occurs when a contact is
	// being created and the SavedItemFolderId attribute refers to a
	// non-contact folder.
	ErrorCannotCreateContactInNonContactFolder ResponseCode = "ErrorCannotCreateContactInNonContactFolder"

	// ErrorCannotCreatePostItemInNonMailFolder indicates that a post item
	// cannot be created in a folder other than a mail folder, such as Calendar,
	// Contact, Tasks, Notes, and so on.
	ErrorCannotCreatePostItemInNonMailFolder ResponseCode = "ErrorCannotCreatePostItemInNonMailFolder"

	// ErrorCannotCreateTaskInNonTaskFolder occurs when a task is being created
	// and the SavedItemFolderId attribute refers to a non-task folder.
	ErrorCannotCreateTaskInNonTaskFolder ResponseCode = "ErrorCannotCreateTaskInNonTaskFolder"

	// ErrorCannotDeleteObject occurs when the item or folder to delete cannot
	// be deleted.
	ErrorCannotDeleteObject ResponseCode = "ErrorCannotDeleteObject"

	// ErrorCannotDeleteTaskOccurrence is returned by the DeleteItem operation
	// when it fails to delete the current occurrence of a recurring task. This
	// can only happen if the AffectedTaskOccurrences attribute has been set to
	// SpecifiedOccurrenceOnly.
	ErrorCannotDeleteTaskOccurrence ResponseCode = "ErrorCannotDeleteTaskOccurrence"

	// ErrorCannotDisableMandatoryExtension indicates that an attempt was made
	// to disable a mandatory extension.
	ErrorCannotDisableMandatoryExtension ResponseCode = "ErrorCannotDisableMandatoryExtension"

	// ErrorCannotEmptyFolder must be returned when the server cannot empty a
	// folder.
	ErrorCannotEmptyFolder ResponseCode = "ErrorCannotEmptyFolder"

	// ErrorCannotGetSourceFolderPath indicates that the source folder path
	// could not be retrieved.
	ErrorCannotGetSourceFolderPath ResponseCode = "ErrorCannotGetSourceFolderPath"

	// ErrorCannotGetExternalEcpUrl specifies that the server could not
	// retrieve the external URL for Outlook Web App Options.
	ErrorCannotGetExternalEcpUrl ResponseCode = "ErrorCannotGetExternalEcpUrl"

	// ErrorCannotOpenFileAttachment is returned by the GetAttachment operation
	// if it cannot retrieve the body of a file attachment.
	ErrorCannotOpenFileAttachment ResponseCode = "ErrorCannotOpenFileAttachment"

	// ErrorCannotSetCalendarPermissionOnNonCalendarFolder indicates that the
	// caller tried to set calendar permissions on a non-calendar folder.
	ErrorCannotSetCalendarPermissionOnNonCalendarFolder ResponseCode = "ErrorCannotSetCalendarPermissionOnNonCalendarFolder"

	// ErrorCannotSetNonCalendarPermissionOnCalendarFolder indicates that the
	// caller tried to set non-calendar permissions on a calendar folder.
	ErrorCannotSetNonCalendarPermissionOnCalendarFolder ResponseCode = "ErrorCannotSetNonCalendarPermissionOnCalendarFolder"

	// ErrorCannotSetPermissionUnknownEntries indicates that you cannot set
	// unknown permissions in a permissions set.
	ErrorCannotSetPermissionUnknownEntries ResponseCode = "ErrorCannotSetPermissionUnknownEntries"

	// ErrorCannotSpecifySearchFolderAsSourceFolder indicates that an attempt
	// was made to specify the search folder as the source folder.
	ErrorCannotSpecifySearchFolderAsSourceFolder ResponseCode = "ErrorCannotSpecifySearchFolderAsSourceFolder"

	// ErrorCannotUseFolderIdForItemId occurs when a request that requires an
	// item identifier is given a folder identifier.
	ErrorCannotUseFolderIdForItemId ResponseCode = "ErrorCannotUseFolderIdForItemId"

	// ErrorCannotUseItemIdForFolderId occurs when a request that requires a
	// folder identifier is given an item identifier.
	ErrorCannotUseItemIdForFolderId ResponseCode = "ErrorCannotUseItemIdForFolderId"

	// ErrorChangeKeyRequired code has been replaced by
	// ErrorChangeKeyRequiredForWriteOperations
	ErrorChangeKeyRequired ResponseCode = "ErrorChangeKeyRequired"

	// ErrorChangeKeyRequiredForWriteOperations is returned when the change key
	// for an item is missing or stale. For SendItem, UpdateItem, and
	// UpdateFolder operations, the caller must pass in a correct and current
	// change key for the item. Note that this is the case with UpdateItem even
	// when conflict resolution is set to always overwrite.
	ErrorChangeKeyRequiredForWriteOperations ResponseCode = "ErrorChangeKeyRequiredForWriteOperations"

	// ErrorClientDisconnected specifies that the client was disconnected.
	ErrorClientDisconnected ResponseCode = "ErrorClientDisconnected"

	// ErrorClientIntentInvalidStateDefinition is intended for internal use only.
	ErrorClientIntentInvalidStateDefinition ResponseCode = "ErrorClientIntentInvalidStateDefinition"

	// ErrorClientIntentNotFound is intended for internal use only.
	ErrorClientIntentNotFound ResponseCode = "ErrorClientIntentNotFound"

	// ErrorConnectionFailed occurs when Exchange Web Services cannot connect
	// to the mailbox.
	ErrorConnectionFailed ResponseCode = "ErrorConnectionFailed"

	// ErrorContainsFilterWrongType indicates that the property that was
	// inspected for a Contains filter is not a string type.
	ErrorContainsFilterWrongType ResponseCode = "ErrorContainsFilterWrongType"

	// ErrorContentConversionFailed is returned by the GetItem operation when
	// Exchange Web Services is unable to retrieve the MIME content for the
	// item requested. The CreateItem operation returns this error when
	// Exchange Web Services is unable to create the item from the supplied
	// MIME content. Usually this is an indication that the item property is
	// corrupted or truncated.
	ErrorContentConversionFailed ResponseCode = "ErrorContentConversionFailed"

	// ErrorContentIndexingNotEnabled occurs when a search request is made
	// using the QueryString option and content indexing is not enabled for the
	// target mailbox.
	ErrorContentIndexingNotEnabled ResponseCode = "ErrorContentIndexingNotEnabled"

	// ErrorCorruptData occurs when the data is corrupted and cannot be
	// processed.
	ErrorCorruptData ResponseCode = "ErrorCorruptData"

	// ErrorCreateItemAccessDenied occurs when the caller does not have
	// permission to create the item.
	ErrorCreateItemAccessDenied ResponseCode = "ErrorCreateItemAccessDenied"

	// ErrorCreateManagedFolderPartialCompletion occurs when one or more of the
	// managed folders that were specified in the CreateManagedFolder operation
	// request failed to be created. Search for each folder to determine which
	// folders were created and which folders do not exist.
	ErrorCreateManagedFolderPartialCompletion ResponseCode = "ErrorCreateManagedFolderPartialCompletion"

	// ErrorCreateSubfolderAccessDenied occurs when the calling account does
	// not have the permissions required to create the subfolder.
	ErrorCreateSubfolderAccessDenied ResponseCode = "ErrorCreateSubfolderAccessDenied"

	// ErrorCrossMailboxMoveCopy occurs when an attempt is made to move an item
	// or folder from one mailbox to another. If the source mailbox and
	// destination mailbox are different, you will get this error.
	ErrorCrossMailboxMoveCopy ResponseCode = "ErrorCrossMailboxMoveCopy"

	// ErrorCrossSiteRequest indicates that the request is not allowed because
	// the Client Access server that should service the request is in a
	// different site.
	ErrorCrossSiteRequest ResponseCode = "ErrorCrossSiteRequest"

	// ErrorDataSizeLimitExceeded can occur in the following scenarios:
	//  - An attempt is made to access or write a property on an item and the
	//    property value is too large.
	//  - The base64 encoded MIME content length within the request XML exceeds
	//    the limit.
	//  - The size of the body of an existing item body exceeds the limit.
	//  - The consumer tries to set an HTML or text body whose length (or
	//    combined length in the case of append) exceeds the limit.
	ErrorDataSizeLimitExceeded ResponseCode = "ErrorDataSizeLimitExceeded"

	// ErrorDataSourceOperation occurs when the underlying data provider fails
	// to complete the operation.
	ErrorDataSourceOperation ResponseCode = "ErrorDataSourceOperation"

	// ErrorDelegateAlreadyExists occurs in an AddDelegate operation when the
	// specified user already exists in the list of delegates.
	ErrorDelegateAlreadyExists ResponseCode = "ErrorDelegateAlreadyExists"

	// ErrorDelegateCannotAddOwner occurs in an AddDelegate operation when the
	// specified user to be added is the owner of the mailbox.
	ErrorDelegateCannotAddOwner ResponseCode = "ErrorDelegateCannotAddOwner"

	// ErrorDelegateMissingConfiguration occurs in a GetDelegate operation when
	// either there is no delegate information on the local FreeBusy message or
	// no Active Directory public delegate (no "public delegate" or no "Send On
	// Behalf" entry in AD DS).
	ErrorDelegateMissingConfiguration ResponseCode = "ErrorDelegateMissingConfiguration"

	// ErrorDelegateNoUser occurs when a specified user cannot be mapped to a
	// user in AD DS.
	ErrorDelegateNoUser ResponseCode = "ErrorDelegateNoUser"

	// ErrorDelegateValidationFailed occurs in the AddDelegate operation when
	// an added delegate user is not valid.
	ErrorDelegateValidationFailed ResponseCode = "ErrorDelegateValidationFailed"

	// ErrorDeleteDistinguishedFolder occurs when an attempt is made to delete
	// a distinguished folder.
	ErrorDeleteDistinguishedFolder ResponseCode = "ErrorDeleteDistinguishedFolder"

	// ErrorDeleteItemsFailed is not used.
	// ErrorDeleteItemsFailed ResponseCode = "ErrorDeleteItemsFailed"

	// ErrorDeleteUnifiedMessagingPromptFailed is intended for internal use
	// only.
	ErrorDeleteUnifiedMessagingPromptFailed ResponseCode = "ErrorDeleteUnifiedMessagingPromptFailed"

	// ErrorDistinguishedUserNotSupported indicates that a distinguished user
	// ID is not valid for the operation. DistinguishedUserType should not be
	// present in the request.
	ErrorDistinguishedUserNotSupported ResponseCode = "ErrorDistinguishedUserNotSupported"

	// ErrorDistributionListMemberNotExist indicates that a request
	// distribution list member does not exist in the distribution list.
	ErrorDistributionListMemberNotExist ResponseCode = "ErrorDistributionListMemberNotExist"

	// ErrorDuplicateInputFolderNames occurs when duplicate folder names are
	// specified within the FolderNames element of the CreateManagedFolder
	// operation request.
	ErrorDuplicateInputFolderNames ResponseCode = "ErrorDuplicateInputFolderNames"

	// ErrorDuplicateSOAPHeader indicates that there are duplicate SOAP headers.
	ErrorDuplicateSOAPHeader ResponseCode = "ErrorDuplicateSOAPHeader"

	// ErrorDuplicateUserIdsSpecified indicates that a duplicate user ID has
	// been found in a permission set, either Default or Anonymous are set more
	// than once, or there are duplicate SIDs or recipients.
	ErrorDuplicateUserIdsSpecified ResponseCode = "ErrorDuplicateUserIdsSpecified"

	// ErrorEmailAddressMismatch occurs when a request attempts to
	// create/update the search parameters of a search folder. For example,
	// this can occur when a search folder is created in the mailbox but the
	// search folder is directed to look in another mailbox.
	ErrorEmailAddressMismatch ResponseCode = "ErrorEmailAddressMismatch"

	// ErrorEventNotFound occurs when the event that is associated with a
	// watermark is deleted before the event is returned. When this error is
	// returned, the subscription is also deleted.
	ErrorEventNotFound ResponseCode = "ErrorEventNotFound"

	// ErrorExceededConnectionCount -iIndicates that there are more concurrent
	// requests against the server than are allowed by a user's policy.
	ErrorExceededConnectionCount ResponseCode = "ErrorExceededConnectionCount"

	// ErrorExceededSubscriptionCount indicates that a user's throttling policy
	// maximum subscription count has been exceeded.
	ErrorExceededSubscriptionCount ResponseCode = "ErrorExceededSubscriptionCount"

	// ErrorExceededFindCountLimit indicates that a search operation call has
	// exceeded the total number of items that can be returned.
	ErrorExceededFindCountLimit ResponseCode = "ErrorExceededFindCountLimit"

	// ErrorExpiredSubscription occurs if the GetEvents operation is called as
	// a subscription is being deleted because it has expired.
	ErrorExpiredSubscription ResponseCode = "ErrorExpiredSubscription"

	// ErrorExtensionNotFound indicates that the extension was not found.
	ErrorExtensionNotFound ResponseCode = "ErrorExtensionNotFound"

	// ErrorFolderCorrupt occurs when the folder is corrupted and cannot be
	// saved.
	ErrorFolderCorrupt ResponseCode = "ErrorFolderCorrupt"

	// ErrorFolderExists occurs when an attempt is made to create a folder that
	// has the same name as another folder in the same parent. Duplicate folder
	// names are not allowed.
	ErrorFolderExists ResponseCode = "ErrorFolderExists"

	// ErrorFolderNotFound indicates that the folder ID that was specified does
	// not correspond to a valid folder, or that the delegate does not have
	// permission to access the folder.
	ErrorFolderNotFound ResponseCode = "ErrorFolderNotFound"

	// ErrorFolderPropertRequestFailed indicates that the requested property
	// could not be retrieved. This does not indicate that the property does
	// not exist, but that the property was corrupted in some way so that the
	// retrieval failed.
	ErrorFolderPropertRequestFailed ResponseCode = "ErrorFolderPropertRequestFailed"

	// ErrorFolderSave indicates that the folder could not be created or
	// updated because of an invalid state.
	ErrorFolderSave ResponseCode = "ErrorFolderSave"

	// ErrorFolderSaveFailed indicates that the folder could not be created or
	// updated because of an invalid state.
	ErrorFolderSaveFailed ResponseCode = "ErrorFolderSaveFailed"

	// ErrorFolderSavePropertyError indicates that the folder could not be
	// created or updated because of invalid property values. The response code
	// lists which properties caused the problem.
	ErrorFolderSavePropertyError ResponseCode = "ErrorFolderSavePropertyError"

	// ErrorFreeBusyDLLimitReached indicates that the maximum group member
	// count has been reached for obtaining free/busy information for a
	// distribution list.
	ErrorFreeBusyDLLimitReached ResponseCode = "ErrorFreeBusyDLLimitReached"

	// ErrorFreeBusyGenerationFailed is returned when free/busy information
	// cannot be retrieved because of an intervening failure.
	ErrorFreeBusyGenerationFailed ResponseCode = "ErrorFreeBusyGenerationFailed"

	// ErrorGetServerSecurityDescriptorFailed is not used.
	// ErrorGetServerSecurityDescriptorFailed ResponseCode = "ErrorGetServerSecurityDescriptorFailed"

	// ErrorImContactLimitReached is returned when new instant messaging (IM)
	// contacts cannot be added because the maximum number of contacts has been
	// reached. This error was introduced in Exchange LogClientStart 2013.
	ErrorImContactLimitReached ResponseCode = "ErrorImContactLimitReached"

	// ErrorImGroupDisplayNameAlreadyExists is returned when an attempt is made
	// to add a group display name when an existing group already has the same
	// display name. This error was introduced in Exchange 2013.
	ErrorImGroupDisplayNameAlreadyExists ResponseCode = "ErrorImGroupDisplayNameAlreadyExists"

	// ErrorImGroupLimitReached is returned when new IM groups cannot be added
	// because the maximum number of groups has been reached.
	// This error was introduced in Exchange 2013.
	ErrorImGroupLimitReached ResponseCode = "ErrorImGroupLimitReached"

	// ErrorImpersonateUserDenied is returned in the server-to-server
	// authorization case for Exchange Impersonation when the caller does not
	// have the proper rights to impersonate the specific user in question.
	// This error maps to the ms-Exch-EPI-May-Impersonate extended Active
	// Directory right.
	ErrorImpersonateUserDenied ResponseCode = "ErrorImpersonateUserDenied"

	// ErrorImpersonationDenied is returned in the server-to-server
	// authorization for Exchange Impersonation when the caller does not have
	// the proper rights to impersonate through the Client Access server that
	// they are making the request against. This maps to the
	// ms-Exch-EPI-Impersonation extended Active Directory right.
	ErrorImpersonationDenied ResponseCode = "ErrorImpersonationDenied"

	// ErrorImpersonationFailed indicates that there was an unexpected error
	// when an attempt was made to perform server-to-server authentication.
	// This response code typically indicates either that the service account
	// that is running the Exchange Web Services application pool is configured
	// incorrectly, that Exchange Web Services cannot talk to the directory, or
	// that a trust between forests is not correctly configured.
	ErrorImpersonationFailed ResponseCode = "ErrorImpersonationFailed"

	// ErrorIncorrectSchemaVersion indicates that the request was valid for the
	// current Exchange LogClientStart version but was invalid for the request server
	// version that was specified.
	ErrorIncorrectSchemaVersion ResponseCode = "ErrorIncorrectSchemaVersion"

	// ErrorIncorrectUpdatePropertyCount indicates that each change description
	// in the UpdateItem or UpdateFolder elements must list only one property
	// to update.
	ErrorIncorrectUpdatePropertyCount ResponseCode = "ErrorIncorrectUpdatePropertyCount"

	// ErrorIndividualMailboxLimitReached occurs when the request contains too
	// many attendees to resolve. By default, the maximum number of attendees
	// to resolve is 100.
	ErrorIndividualMailboxLimitReached ResponseCode = "ErrorIndividualMailboxLimitReached"

	// ErrorInsufficientResources occurs when the mailbox server is overloaded.
	// Try your request again later.
	ErrorInsufficientResources ResponseCode = "ErrorInsufficientResources"

	// ErrorInternalServerError indicates that Exchange Web Services
	// encountered an error that it could not recover from, and a more specific
	// response code that is associated with the error that occurred does not
	// exist.
	ErrorInternalServerError ResponseCode = "ErrorInternalServerError"

	// ErrorInternalServerTransientError indicates that an internal server
	// error occurred and that you should try your request again later.
	ErrorInternalServerTransientError ResponseCode = "ErrorInternalServerTransientError"

	// ErrorInvalidAccessLevel indicates that the level of access that the
	// caller has on the free/busy data is invalid.
	ErrorInvalidAccessLevel ResponseCode = "ErrorInvalidAccessLevel"

	// ErrorInvalidArgument indicates an error caused by all invalid arguments
	// passed to the GetMessageTrackingReport operation. This error is returned
	// in the following scenarios:
	//  - The user specified in the sending-as parameter does not exist in the
	//    directory.
	//  - The user specified in the sending-as parameter is not unique in the
	//    directory.
	//  - The sending-as address is empty.
	//  - The sending-as address is not a valid e-mail address.
	ErrorInvalidArgument ResponseCode = "ErrorInvalidArgument"

	// ErrorInvalidAttachmentId is returned by the GetAttachment operation or
	// the DeleteAttachment operation when an attachment that corresponds to
	// the specified ID is not found.
	ErrorInvalidAttachmentId ResponseCode = "ErrorInvalidAttachmentId"

	// ErrorInvalidAttachmentSubfilter occurs when you try to bind to an
	// existing search folder by using a complex attachment table restriction.
	// Exchange Web Services only supports simple contains filters against the
	// attachment table. If you try to bind to an existing search folder that
	// has a more complex attachment table restriction (a subfilter), Exchange
	// Web Services cannot render the XML for that filter and returns this
	// response code.
	// Note that you can still call the GetFolder operation on the folder, but
	// do not request the SearchParameters element.
	ErrorInvalidAttachmentSubfilter ResponseCode = "ErrorInvalidAttachmentSubfilter"

	// ErrorInvalidAttachmentSubfilterTextFilter occurs when you try to bind to
	// an existing search folder by using a complex attachment table
	// restriction. Exchange Web Services only supports simple contains filters
	// against the attachment table.
	// If you try to bind to an existing search folder that has a more complex
	// attachment table restriction, Exchange Web Services cannot render the
	// XML for that filter. In this case, the attachment subfilter contains a
	// text filter, but it is not looking at the attachment display name.
	// Note that you can still call the GetFolder operation on the folder, but
	// do not request the SearchParameters element.
	ErrorInvalidAttachmentSubfilterTextFilter ResponseCode = "ErrorInvalidAttachmentSubfilterTextFilter"

	// ErrorInvalidAuthorizationContext indicates that the authorization
	// procedure for the requestor failed.
	ErrorInvalidAuthorizationContext ResponseCode = "ErrorInvalidAuthorizationContext"

	// ErrorInvalidChangeKey occurs when a consumer passes in a folder or item
	// identifier with a change key that cannot be parsed. For example, this
	// could be invalid base64 content or an empty string.
	ErrorInvalidChangeKey ResponseCode = "ErrorInvalidChangeKey"

	// ErrorInvalidClientSecurityContext indicates that there was an internal
	// error when an attempt was made to resolve the identity of the caller.
	ErrorInvalidClientSecurityContext ResponseCode = "ErrorInvalidClientSecurityContext"

	// ErrorInvalidCompleteDate is returned when an attempt is made to set the
	// CompleteDate element value of a task to a time in the future. When it is
	// converted to the local time of the Client Access server, the
	// CompleteDate of a task cannot be set to a value that is later than the
	// local time on the Client Access server.
	ErrorInvalidCompleteDate ResponseCode = "ErrorInvalidCompleteDate"

	// ErrorInvalidContactEmailAddress indicates that an invalid e-mail address
	// was provided for a contact.
	ErrorInvalidContactEmailAddress ResponseCode = "ErrorInvalidContactEmailAddress"

	// ErrorInvalidContactEmailIndex indicates that an invalid e-mail index
	// value was provided for an e-mail entry.
	ErrorInvalidContactEmailIndex ResponseCode = "ErrorInvalidContactEmailIndex"

	// ErrorInvalidCrossForestCredentials occurs when the credentials that are
	// used to proxy a request to a different directory service forest fail
	// authentication.
	ErrorInvalidCrossForestCredentials ResponseCode = "ErrorInvalidCrossForestCredentials"

	// ErrorInvalidDelegatePermission indicates that the specified folder
	// permissions are invalid.
	ErrorInvalidDelegatePermission ResponseCode = "ErrorInvalidDelegatePermission"

	// ErrorInvalidDelegateUserId indicates that the specified delegate user ID
	// is invalid.
	ErrorInvalidDelegateUserId ResponseCode = "ErrorInvalidDelegateUserId"

	// ErrorInvalidExchangeImpersonationHeaderData occurs during Exchange
	// Impersonation when a caller does not specify a UPN, an e-mail address,
	// or a user SID. This will also occur if the caller specifies one or more
	// of those and the values are empty.
	ErrorInvalidExchangeImpersonationHeaderData ResponseCode = "ErrorInvalidExchangeImpersonationHeaderData"

	// ErrorInvalidExcludesRestriction occurs when the bitmask that was passed
	// into an Excludes element restriction is unable to be parsed.
	ErrorInvalidExcludesRestriction ResponseCode = "ErrorInvalidExcludesRestriction"

	// ErrorInvalidExpressionTypeForSubFilter is not used.
	// ErrorInvalidExpressionTypeForSubFilter ResponseCode = "ErrorInvalidExpressionTypeForSubFilter"

	// ErrorInvalidExtendedProperty occurs when the following events take place:
	//  - The caller tries to use an extended property that is not supported by
	//    Exchange Web Services.
	//  - The caller passes in an invalid combination of attribute values for
	//    an extended property. This also occurs if no attributes are passed.
	//    Only certain combinations are allowed.
	ErrorInvalidExtendedProperty ResponseCode = "ErrorInvalidExtendedProperty"

	// ErrorInvalidExtendedPropertyValue occurs when the value section of an
	// extended property does not match the type of the property.
	// For example, setting an extended property that has PropertyType="String"
	// to an array of integers will result in this error. Any string
	// representation that is not coercible into the type that is specified for
	// the extended property will give this error.
	ErrorInvalidExtendedPropertyValue ResponseCode = "ErrorInvalidExtendedPropertyValue"

	// ErrorInvalidExternalSharingSubscriber indicates that the sharing
	// invitation sender did not create the sharing invitation metadata.
	ErrorInvalidExternalSharingSubscriber ResponseCode = "ErrorInvalidExternalSharingSubscriber"

	// ErrorInvalidExternalSharingInitiator indicates that a sharing message is
	// not intended for the caller.
	// Specifies that the sharing invitation sender did not create the sharing
	// invitation metadata. This error code MUST be returned if the sharing
	// invitation sender did not create the sharing invitation metadata.
	ErrorInvalidExternalSharingInitiator ResponseCode = "ErrorInvalidExternalSharingInitiator"

	// ErrorInvalidFederatedOrganizationId indicates that the requestor's
	// organization federation objects are not correctly configured.
	ErrorInvalidFederatedOrganizationId ResponseCode = "ErrorInvalidFederatedOrganizationId"

	// ErrorInvalidFolderId occurs when the folder ID is corrupt.
	ErrorInvalidFolderId ResponseCode = "ErrorInvalidFolderId"

	// ErrorInvalidFolderTypeForOperation indicates that the specified folder
	// type is invalid for the current operation. For example, you cannot
	// create a Search folder in a public folder.
	ErrorInvalidFolderTypeForOperation ResponseCode = "ErrorInvalidFolderTypeForOperation"

	// ErrorInvalidFractionalPagingParameters occurs in fractional paging when
	// the user has specified one of the following:
	//  - A numerator that is greater than the denominator
	//  - A numerator that is less than zero
	//  - A denominator that is less than or equal to zero.
	ErrorInvalidFractionalPagingParameters ResponseCode = "ErrorInvalidFractionalPagingParameters"

	// ErrorInvalidGetSharingFolderRequest indicates that the DataType and
	// ShareFolderId elements are both present in a request.
	ErrorInvalidGetSharingFolderRequest ResponseCode = "ErrorInvalidGetSharingFolderRequest"

	// ErrorInvalidFreeBusyViewType occurs when the GetUserAvailability
	// operation is called with a FreeBusyViewType of None.
	ErrorInvalidFreeBusyViewType ResponseCode = "ErrorInvalidFreeBusyViewType"

	// ErrorInvalidId indicates that the ID and/or change key is malformed.
	ErrorInvalidId ResponseCode = "ErrorInvalidId"

	// ErrorInvalidIdEmpty occurs when the caller specifies an Id attribute
	// that is empty.
	ErrorInvalidIdEmpty ResponseCode = "ErrorInvalidIdEmpty"

	// ErrorInvalidLikeRequest occurs when the item can't be liked. Versions of
	// Exchange starting with build number 15.00.0913.09 include this value.
	ErrorInvalidLikeRequest ResponseCode = "ErrorInvalidLikeRequest"

	// ErrorInvalidIdMalformed occurs when the caller specifies an Id attribute
	// that is malformed.
	ErrorInvalidIdMalformed ResponseCode = "ErrorInvalidIdMalformed"

	// ErrorInvalidIdMalformedEwsLegacyIdFormat indicates that a folder or item
	// ID that is using the Exchange 2007 format was specified for a request
	// with a version of Exchange 2007 SP1 or later. You cannot use Exchange
	// 2007 IDs in Exchange 2007 SP1 or later requests. You have to use the
	// ConvertId operation to convert them first.
	ErrorInvalidIdMalformedEwsLegacyIdFormat ResponseCode = "ErrorInvalidIdMalformedEwsLegacyIdFormat"

	// ErrorInvalidIdMonikerTooLong occurs when the caller specifies an Id
	// attribute that is too long.
	ErrorInvalidIdMonikerTooLong ResponseCode = "ErrorInvalidIdMonikerTooLong"

	// ErrorInvalidIdNotAnItemAttachmentId is returned whenever an ID that is
	// not an item attachment ID is passed to a Web service method that expects
	// an attachment ID.
	ErrorInvalidIdNotAnItemAttachmentId ResponseCode = "ErrorInvalidIdNotAnItemAttachmentId"

	// ErrorInvalidIdReturnedByResolveNames occurs when a contact in your
	// mailbox is corrupt.
	ErrorInvalidIdReturnedByResolveNames ResponseCode = "ErrorInvalidIdReturnedByResolveNames"

	// ErrorInvalidIdStoreObjectIdTooLong occurs when the caller specifies an
	// Id attribute that is too long.
	ErrorInvalidIdStoreObjectIdTooLong ResponseCode = "ErrorInvalidIdStoreObjectIdTooLong"

	// ErrorInvalidIdTooManyAttachmentLevels is returned when the attachment
	// hierarchy on an item exceeds the maximum of 255 levels deep.
	ErrorInvalidIdTooManyAttachmentLevels ResponseCode = "ErrorInvalidIdTooManyAttachmentLevels"

	// ErrorInvalidIdXml is not used.
	// ErrorInvalidIdXml ResponseCode = "ErrorInvalidIdXml"

	// ErrorInvalidImContactId is returned when the specified IM contact
	// identifier does not represent a valid identifier.
	// This error was introduced in Exchange 2013.
	ErrorInvalidImContactId ResponseCode = "ErrorInvalidImContactId"

	// ErrorInvalidImDistributionGroupSmtpAddress is returned when the
	// specified IM distribution group SMTP address identifier does not
	// represent a valid identifier.
	// This error was introduced in Exchange 2013.
	ErrorInvalidImDistributionGroupSmtpAddress ResponseCode = "ErrorInvalidImDistributionGroupSmtpAddress"

	// ErrorInvalidImGroupId is returned when the specified IM group identifier
	// does not represent a valid identifier.
	// This error was introduced in Exchange 2013.
	ErrorInvalidImGroupId ResponseCode = "ErrorInvalidImGroupId"

	// ErrorInvalidIndexedPagingParameters occurs if the offset for indexed
	// paging is negative.
	ErrorInvalidIndexedPagingParameters ResponseCode = "ErrorInvalidIndexedPagingParameters"

	// ErrorInvalidInternetHeaderChildNodes is not used.
	// ErrorInvalidInternetHeaderChildNodes ResponseCode = "ErrorInvalidInternetHeaderChildNodes"

	// ErrorInvalidItemForOperationArchiveItem indicates that the item was
	// invalid for an ArchiveItem operation.
	ErrorInvalidItemForOperationArchiveItem ResponseCode = "ErrorInvalidItemForOperationArchiveItem"

	// ErrorInvalidItemForOperationAcceptItem occurs when an attempt is made to
	// use an AcceptItem response object for an item type other than a meeting
	// request or a calendar item, or when an attempt is made to accept a
	// calendar item occurrence that is in the Deleted Items folder.
	ErrorInvalidItemForOperationAcceptItem ResponseCode = "ErrorInvalidItemForOperationAcceptItem"

	// ErrorInvalidItemForOperationCancelItem occurs when an attempt is made to
	// use a CancelItem response object on an item type other than a calendar
	// item.
	ErrorInvalidItemForOperationCancelItem ResponseCode = "ErrorInvalidItemForOperationCancelItem"

	// ErrorInvalidItemForOperationCreateItemAttachment is returned when an
	// attempt is made to create an item attachment of an unsupported type.
	// Supported item types for item attachments include the following objects:
	//  - Item
	//  - Message
	//  - CalendarItem
	//  - Task
	//  - Contact
	// For example, if you try to create a MeetingMessage attachment, you'll
	// encounter this response code.
	ErrorInvalidItemForOperationCreateItemAttachment ResponseCode = "ErrorInvalidItemForOperationCreateItemAttachment"

	// ErrorInvalidItemForOperationCreateItem is returned from a CreateItem
	// operation if the request contains an unsupported item type.
	// Supported items include the following objects:
	//  - Item
	//  - Message
	//  - CalendarItem
	//  - Task
	//  - Contact
	// Certain types are created as a side effect of doing something else.
	// Meeting messages, for example, are created when you send a calendar
	// item to attendees; they are not explicitly created.
	ErrorInvalidItemForOperationCreateItem ResponseCode = "ErrorInvalidItemForOperationCreateItem"

	// ErrorInvalidItemForOperationDeclineItem occurs when an attempt is made
	// to use a DeclineItem response object for an item type other than a
	// meeting request or a calendar item, or when an attempt is made to
	// decline a calendar item occurrence that is in the Deleted Items folder.
	ErrorInvalidItemForOperationDeclineItem ResponseCode = "ErrorInvalidItemForOperationDeclineItem"

	// ErrorInvalidItemForOperationExpandDL indicates that the ExpandDL
	// operation is valid only for private distribution lists.
	ErrorInvalidItemForOperationExpandDL ResponseCode = "ErrorInvalidItemForOperationExpandDL"

	// ErrorInvalidItemForOperationRemoveItem is returned from a RemoveItem
	// response object if the request specifies an item that is not a meeting
	// cancellation item.
	ErrorInvalidItemForOperationRemoveItem ResponseCode = "ErrorInvalidItemForOperationRemoveItem"

	// ErrorInvalidItemForOperationSendItem is returned from a SendItem
	// operation if the request specifies an item that is not a message item.
	ErrorInvalidItemForOperationSendItem ResponseCode = "ErrorInvalidItemForOperationSendItem"

	// ErrorInvalidItemForOperationTentative occurs when an attempt is made to
	// use TentativelyAcceptItem for an item type other than a meeting request
	// or a calendar item, or when an attempt is made to tentatively accept a
	// calendar item occurrence that is in the Deleted Items folder.
	ErrorInvalidItemForOperationTentative ResponseCode = "ErrorInvalidItemForOperationTentative"

	// ErrorInvalidLogonType is for internal use only. This error is not returned.
	// ErrorInvalidLogonType ResponseCode = "ErrorInvalidLogonType"

	// ErrorInvalidMailbox indicates that the CreateItem operation or the
	// UpdateItem operation failed while creating or updating a personal
	// distribution list.
	ErrorInvalidMailbox ResponseCode = "ErrorInvalidMailbox"

	// ErrorInvalidManagedFolderProperty occurs when the structure of the
	// managed folder is corrupted and cannot be rendered.
	ErrorInvalidManagedFolderProperty ResponseCode = "ErrorInvalidManagedFolderProperty"

	// ErrorInvalidManagedFolderQuota occurs when the quota that is set on the
	// managed folder is less than zero, which indicates a corrupted managed
	// folder.
	ErrorInvalidManagedFolderQuota ResponseCode = "ErrorInvalidManagedFolderQuota"

	// ErrorInvalidManagedFolderSize occurs when the size that is set on the
	// managed folder is less than zero, which indicates a corrupted managed
	// folder.
	ErrorInvalidManagedFolderSize ResponseCode = "ErrorInvalidManagedFolderSize"

	// ErrorInvalidMergedFreeBusyInterval occurs when the supplied merged
	// free/busy internal value is invalid. The default minimum value is
	// 5 minutes. The default maximum value is 1440 minutes.
	ErrorInvalidMergedFreeBusyInterval ResponseCode = "ErrorInvalidMergedFreeBusyInterval"

	// ErrorInvalidNameForNameResolution occurs when the name is invalid for
	// the ResolveNames operation. For example, a zero length string, a single
	// space, a comma, and a dash are all invalid names.
	ErrorInvalidNameForNameResolution ResponseCode = "ErrorInvalidNameForNameResolution"

	// ErrorInvalidNetworkServiceContext indicates an error in the Network
	// Service account on the Client Access server.
	ErrorInvalidNetworkServiceContext ResponseCode = "ErrorInvalidNetworkServiceContext"

	// ErrorInvalidOofParameter is not used.
	ErrorInvalidOofParameter ResponseCode = "ErrorInvalidOofParameter"

	// ErrorInvalidOperation is a general error that is used when the requested
	// operation is invalid. For example, you cannot do the following:
	//  - Perform a "Deep" traversal by using the FindFolder operation on a
	//    public folder.
	//  - Move or copy the public folder root.
	//  - Delete an associated item by using any mode except "Hard" delete.
	//  - Move or copy an associated item.
	ErrorInvalidOperation ResponseCode = "ErrorInvalidOperation"

	// ErrorInvalidOrganizationRelationshipForFreeBusy indicates that a caller
	// requested free/busy information for a user in another organization but
	// the organizational relationship does not have free/busy enabled.
	ErrorInvalidOrganizationRelationshipForFreeBusy ResponseCode = "ErrorInvalidOrganizationRelationshipForFreeBusy"

	// ErrorInvalidPagingMaxRows occurs when a consumer passes in a zero or a
	// negative value for the maximum rows to be returned.
	ErrorInvalidPagingMaxRows ResponseCode = "ErrorInvalidPagingMaxRows"

	// ErrorInvalidParentFolder occurs when a consumer passes in an invalid
	// parent folder for an operation. For example, this error is returned if
	// you try to create a folder within a search folder.
	ErrorInvalidParentFolder ResponseCode = "ErrorInvalidParentFolder"

	// ErrorInvalidPercentCompleteValue is returned when an attempt is made to
	// set a task completion percentage to an invalid value. The value must be
	// between 0 and 100 (inclusive).
	ErrorInvalidPercentCompleteValue ResponseCode = "ErrorInvalidPercentCompleteValue"

	// ErrorInvalidPermissionSettings indicates that the permission level is
	// inconsistent with the permission settings.
	ErrorInvalidPermissionSettings ResponseCode = "ErrorInvalidPermissionSettings"

	// ErrorInvalidPhoneCallId indicates that the caller identifier is not
	// valid.
	ErrorInvalidPhoneCallId ResponseCode = "ErrorInvalidPhoneCallId"

	// ErrorInvalidPhoneNumber indicates that the telephone number is not
	// correct or does not fit the dial plan rules.
	ErrorInvalidPhoneNumber ResponseCode = "ErrorInvalidPhoneNumber"

	// ErrorInvalidPropertyAppend occurs when the property that you are trying
	// to append to does not support appending. The following are the only
	// properties that support appending:
	//  - Recipient collections (ToRecipients, CcRecipients, BccRecipients)
	//  - Attendee collections (RequiredAttendees, OptionalAttendees, Resources)
	//  - Body
	//  - ReplyTo
	// In addition, this error occurs when a message body is appended if the
	// format specified in the request does not match the format of the item in
	// the store.
	ErrorInvalidPropertyAppend ResponseCode = "ErrorInvalidPropertyAppend"

	// ErrorInvalidPropertyDelete occurs if the delete operation is specified
	// in an UpdateItem operation or UpdateFolder operation call for a property
	// that does not support the delete operation. For example, you cannot
	// delete the ItemId element of the Item object.
	ErrorInvalidPropertyDelete ResponseCode = "ErrorInvalidPropertyDelete"

	// ErrorInvalidPropertyForExists occurs if the consumer passes in one of
	// the flag properties in an Exists filter. For example, this error occurs
	// if the IsRead or IsFromMe flags are specified in the Exists element. The
	// request should use the IsEqualTo element instead for these as they are
	// flags and therefore part of a single property.
	ErrorInvalidPropertyForExists ResponseCode = "ErrorInvalidPropertyForExists"

	// ErrorInvalidPropertyForOperation occurs when the property that you are
	// trying to manipulate does not support the operation that is being
	// performed on it.
	ErrorInvalidPropertyForOperation ResponseCode = "ErrorInvalidPropertyForOperation"

	// ErrorInvalidPropertyRequest occurs if a property that is specified in
	// the request is not available for the item type. For example, this error
	// is returned if a property that is only available on calendar items is
	// requested in a GetItem operation call for a message or is updated in an
	// UpdateItem operation call for a message. This occurs in the following
	// operations:
	//  - GetItem operation
	//  - GetFolder operation
	//  - UpdateItem operation
	//  - UpdateFolder operation
	ErrorInvalidPropertyRequest ResponseCode = "ErrorInvalidPropertyRequest"

	// ErrorInvalidPropertySet indicates that the property that you are trying
	// to manipulate does not support the operation that is being performed on
	// it. For example, this error is returned if the property that you are
	// trying to set is read-only.
	ErrorInvalidPropertySet ResponseCode = "ErrorInvalidPropertySet"

	// ErrorInvalidPropertyUpdateSentMessage occurs during an UpdateItem
	// operation when a client tries to update certain properties of a message
	// that has already been sent.
	// For example, the following properties cannot be updated on a sent
	// message:
	//  - IsReadReceiptRequested
	//  - IsDeliveryReceiptRequested
	ErrorInvalidPropertyUpdateSentMessage ResponseCode = "ErrorInvalidPropertyUpdateSentMessage"

	// ErrorInvalidProxySecurityContext is not used.
	// ErrorInvalidProxySecurityContext ResponseCode = "ErrorInvalidProxySecurityContext"

	// ErrorInvalidPullSubscriptionId occurs if you call the GetEvents
	// operation or the Unsubscribe operation by using a push subscription ID.
	// To unsubscribe from a push subscription, you must respond to a push
	// request with an unsubscribe response, or disconnect your Web service and
	// wait for the push notifications to time out.
	ErrorInvalidPullSubscriptionId ResponseCode = "ErrorInvalidPullSubscriptionId"

	// ErrorInvalidPushSubscriptionUrl is returned by the Subscribe operation
	// when it creates a "push" subscription and indicates that the URL that is
	// included in the request is invalid. The following conditions must be met
	// for Exchange Web Services to accept the URL:
	//  - String length > 0 and < 2083.
	//  - Protocol is http or https.
	//  - The URL can be parsed by the URI Microsoft .NET Framework class.
	ErrorInvalidPushSubscriptionUrl ResponseCode = "ErrorInvalidPushSubscriptionUrl"

	// ErrorInvalidRecipients indicates that the recipient collection on your
	// message or the attendee collection on your calendar item is invalid.
	// For example, this error will be returned when an attempt is made to send
	// an item that has no recipients.
	ErrorInvalidRecipients ResponseCode = "ErrorInvalidRecipients"

	// ErrorInvalidRecipientSubfilter indicates that the search folder has a
	// recipient table filter that Exchange Web Services cannot represent.
	// To get around this error, retrieve the folder without requesting the
	// search parameters.
	ErrorInvalidRecipientSubfilter ResponseCode = "ErrorInvalidRecipientSubfilter"

	// ErrorInvalidRecipientSubfilterComparison indicates that the search
	// folder has a recipient table filter that Exchange Web Services cannot
	// represent. To get around this error, retrieve the folder without
	// requesting the search parameters.
	ErrorInvalidRecipientSubfilterComparison ResponseCode = "ErrorInvalidRecipientSubfilterComparison"

	// ErrorInvalidRecipientSubfilterOrder indicates that the search folder has
	// a recipient table filter that Exchange Web Services cannot represent.
	// To get around this error, retrieve the folder without requesting the
	// search parameters.
	ErrorInvalidRecipientSubfilterOrder ResponseCode = "ErrorInvalidRecipientSubfilterOrder"

	// ErrorInvalidRecipientSubfilterTextFilter indicates that the search
	// folder has a recipient table filter that Exchange Web Services cannot
	// represent. To get around this error, retrieve the folder without
	// requesting the search parameters.
	ErrorInvalidRecipientSubfilterTextFilter ResponseCode = "ErrorInvalidRecipientSubfilterTextFilter"

	// ErrorInvalidReferenceItem is returned from the CreateItem operation for
	// Forward and Reply response objects in the following scenarios:
	//  - The referenced item identifier is not a Message, a CalendarItem, or
	//    a descendant of a Message or CalendarItem.
	//  - The reference item identifier is for a CalendarItem and the organizer
	//    is trying to Reply or ReplyAll to himself.
	//  - The message is a draft and Reply or ReplyAll is selected.
	//  - The reference item, for SuppressReadReceipt, is not a Message or a
	//    descendant of a Message.
	ErrorInvalidReferenceItem ResponseCode = "ErrorInvalidReferenceItem"

	// ErrorInvalidRequest occurs when the SOAP request has a SOAP action
	// header, but nothing in the SOAP body. Note that the SOAP Action header
	// is not required as Exchange Web Services can determine the method to
	// call from the local name of the root element in the SOAP body.
	ErrorInvalidRequest ResponseCode = "ErrorInvalidRequest"

	// ErrorInvalidRestriction is not used.
	// ErrorInvalidRestriction ResponseCode = "ErrorInvalidRestriction"

	// ErrorInvalidRetentionTagTypeMismatch is returned when the specified
	// retention tag has an incorrect action associated with it.
	// This error was introduced in Exchange 2013.
	ErrorInvalidRetentionTagTypeMismatch ResponseCode = "ErrorInvalidRetentionTagTypeMismatch"

	// ErrorInvalidRetentionTagInvisible is returned when an attempt is made to
	// set a nonexistent or invisible tag on a PolicyTag property.
	// This error was introduced in Exchange 2013.
	ErrorInvalidRetentionTagInvisible ResponseCode = "ErrorInvalidRetentionTagInvisible"

	// ErrorInvalidRetentionTagInheritance is returned when an attempt is made
	// to set an implicit tag on the PolicyTag property.
	// This error was introduced in Exchange 2013.
	ErrorInvalidRetentionTagInheritance ResponseCode = "ErrorInvalidRetentionTagInheritance"

	// ErrorInvalidRetentionTagIdGuid indicates that the retention tag GUID was
	// invalid.
	ErrorInvalidRetentionTagIdGuid ResponseCode = "ErrorInvalidRetentionTagIdGuid"

	// ErrorInvalidRoutingType occurs if the routing type that is passed for a
	// RoutingType (EmailAddressType) element is invalid. Typically, the
	// routing type is set to Simple Mail Transfer Protocol (SMTP).
	ErrorInvalidRoutingType ResponseCode = "ErrorInvalidRoutingType"

	// ErrorInvalidScheduledOofDuration occurs if the specified duration end
	// time is not greater than the start time, or if the end time does not
	// occur in the future.
	ErrorInvalidScheduledOofDuration ResponseCode = "ErrorInvalidScheduledOofDuration"

	// ErrorInvalidSchemaVersionForMailboxVersion indicates that a proxy
	// request that was sent to another server is not able to service the
	// request due to a versioning mismatch.
	ErrorInvalidSchemaVersionForMailboxVersion ResponseCode = "ErrorInvalidSchemaVersionForMailboxVersion"

	// ErrorInvalidSecurityDescriptor indicates that the Exchange security
	// descriptor on the Calendar folder in the store is corrupted.
	ErrorInvalidSecurityDescriptor ResponseCode = "ErrorInvalidSecurityDescriptor"

	// ErrorInvalidSendItemSaveSettings occurs during an attempt to send an
	// item where the SavedItemFolderId is specified in the request but the
	// SaveItemToFolder property is set to false.
	ErrorInvalidSendItemSaveSettings ResponseCode = "ErrorInvalidSendItemSaveSettings"

	// ErrorInvalidSerializedAccessToken indicates that the token that was
	// passed in the header is malformed, does not refer to a valid account in
	// the directory, or is missing the primary group ConnectingSID.
	ErrorInvalidSerializedAccessToken ResponseCode = "ErrorInvalidSerializedAccessToken"

	// ErrorInvalidSharingData indicates that the sharing metadata is not
	// valid. This can be caused by invalid XML.
	ErrorInvalidSharingData ResponseCode = "ErrorInvalidSharingData"

	// ErrorInvalidSharingMessage indicates that the sharing message is not
	// valid. This can be caused by a missing property.
	ErrorInvalidSharingMessage ResponseCode = "ErrorInvalidSharingMessage"

	// ErrorInvalidSid occurs when an invalid SID is passed in a request.
	ErrorInvalidSid ResponseCode = "ErrorInvalidSid"

	// ErrorInvalidSIPUri indicates that the SIP name, dial plan, or the phone
	// number are invalid SIP URIs.
	ErrorInvalidSIPUri ResponseCode = "ErrorInvalidSIPUri"

	// ErrorInvalidServerVersion indicates that an invalid request server
	// version was specified in the request.
	ErrorInvalidServerVersion ResponseCode = "ErrorInvalidServerVersion"

	// ErrorInvalidSmtpAddress occurs when the SMTP address cannot be parsed.
	ErrorInvalidSmtpAddress ResponseCode = "ErrorInvalidSmtpAddress"

	// ErrorInvalidSubfilterType is not used.
	// ErrorInvalidSubfilterType ResponseCode = "ErrorInvalidSubfilterType"

	// ErrorInvalidSubfilterTypeNotAttendeeType is not used.
	// ErrorInvalidSubfilterTypeNotAttendeeType ResponseCode = "ErrorInvalidSubfilterTypeNotAttendeeType"

	// ErrorInvalidSubfilterTypeNotRecipientType is not used.
	// ErrorInvalidSubfilterTypeNotRecipientType ResponseCode = "ErrorInvalidSubfilterTypeNotRecipientType"

	// ErrorInvalidSubscription indicates that the subscription is no longer
	// valid. This could be because the Client Access server is restarting or
	// because the subscription is expired.
	ErrorInvalidSubscription ResponseCode = "ErrorInvalidSubscription"

	// ErrorInvalidSubscriptionRequest indicates that the subscribe request
	// included multiple public folder IDs. A subscription can include multiple
	// folders from the same mailbox or one public folder ID.
	ErrorInvalidSubscriptionRequest ResponseCode = "ErrorInvalidSubscriptionRequest"

	// ErrorInvalidSyncStateData is returned by SyncFolderItems or
	// SyncFolderHierarchy if the SyncState data that is passed is invalid. To
	// fix this error, you must resynchronize without the sync state. Make sure
	// that if you are persisting sync state BLOBs, you are not accidentally
	// truncating the BLOB.
	ErrorInvalidSyncStateData ResponseCode = "ErrorInvalidSyncStateData"

	// ErrorInvalidTimeInterval indicates that the specified time interval is
	// invalid. The start time must be greater than or equal to the end time.
	ErrorInvalidTimeInterval ResponseCode = "ErrorInvalidTimeInterval"

	// ErrorInvalidUserInfo indicates that an internally inconsistent UserId
	// was specified for a permissions operation. For example, if a
	// distinguished user ID is specified (Default or Anonymous), this error is
	// returned if you also try to specify a SID, or primary SMTP address or
	// display name for this user.
	ErrorInvalidUserInfo ResponseCode = "ErrorInvalidUserInfo"

	// ErrorInvalidUserOofSettings indicates that the user Out of Office (OOF)
	// settings are invalid because of a missing internal or external reply.
	ErrorInvalidUserOofSettings ResponseCode = "ErrorInvalidUserOofSettings"

	// ErrorInvalidUserPrincipalName occurs during Exchange Impersonation. The
	// caller passed in an invalid UPN in the SOAP header that was not
	// accessible in the directory.
	ErrorInvalidUserPrincipalName ResponseCode = "ErrorInvalidUserPrincipalName"

	// ErrorInvalidUserSid occurs when an invalid SID is passed in a request.
	ErrorInvalidUserSid ResponseCode = "ErrorInvalidUserSid"

	// ErrorInvalidUserSidMissingUPN is not used.
	ErrorInvalidUserSidMissingUPN ResponseCode = "ErrorInvalidUserSidMissingUPN"

	// ErrorInvalidValueForProperty indicates that the comparison value in the
	// restriction is invalid for the property you are comparing against. For
	// example, the comparison value of DateTimeCreated > true would return
	// this response code. This response code is also returned if you specify
	// an enumeration property in the comparison, but the value that you are
	// comparing against is not a valid value for that enumeration.
	ErrorInvalidValueForProperty ResponseCode = "ErrorInvalidValueForProperty"

	// ErrorInvalidWatermark is caused by an invalid watermark.
	ErrorInvalidWatermark ResponseCode = "ErrorInvalidWatermark"

	// ErrorIPGatewayNotFound indicates that a valid VoIP gateway is not
	// available.
	ErrorIPGatewayNotFound ResponseCode = "ErrorIPGatewayNotFound"

	// ErrorIrresolvableConflict indicates that conflict resolution was unable
	// to resolve changes for the properties. The items in the store may have
	// been changed and have to be updated. Retrieve the updated change key and
	// try again.
	ErrorIrresolvableConflict ResponseCode = "ErrorIrresolvableConflict"

	// ErrorItemCorrupt indicates that the state of the object is corrupted and
	// cannot be retrieved. When you are retrieving an item, only certain
	// elements will be in this state, such as Body and MimeContent. Omit these
	// elements and retry the operation.
	ErrorItemCorrupt ResponseCode = "ErrorItemCorrupt"

	// ErrorItemNotFound occurs when the item was not found or you do not have
	// permission to access the item.
	ErrorItemNotFound ResponseCode = "ErrorItemNotFound"

	// ErrorItemPropertyRequestFailed occurs if a property request on an item
	// fails. The property may exist, but it could not be retrieved.
	ErrorItemPropertyRequestFailed ResponseCode = "ErrorItemPropertyRequestFailed"

	// ErrorItemSave occurs when attempts to save the item or folder fail.
	ErrorItemSave ResponseCode = "ErrorItemSave"

	// ErrorItemSavePropertyError occurs when attempts to save the item or
	// folder fail because of invalid property values. The response code
	// includes the path of the invalid properties.
	ErrorItemSavePropertyError ResponseCode = "ErrorItemSavePropertyError"

	// ErrorLegacyMailboxFreeBusyViewTypeNotMerged is not used.
	// ErrorLegacyMailboxFreeBusyViewTypeNotMerged ResponseCode = "ErrorLegacyMailboxFreeBusyViewTypeNotMerged"

	// ErrorLocalServerObjectNotFound is not used.
	// ErrorLocalServerObjectNotFound ResponseCode = "ErrorLocalServerObjectNotFound"

	// ErrorLogonAsNetworkServiceFailed indicates that the Availability service
	// was unable to log on as the network service to proxy requests to the
	// appropriate sites or forests. This response typically indicates a
	// configuration error.
	ErrorLogonAsNetworkServiceFailed ResponseCode = "ErrorLogonAsNetworkServiceFailed"

	// ErrorMailboxConfiguration indicates that the mailbox information in
	// AD DS is configured incorrectly.
	ErrorMailboxConfiguration ResponseCode = "ErrorMailboxConfiguration"

	// ErrorMailboxDataArrayEmpty indicates that the MailboxDataArray element
	// in the request is empty. You must supply at least one mailbox identifier.
	ErrorMailboxDataArrayEmpty ResponseCode = "ErrorMailboxDataArrayEmpty"

	// ErrorMailboxDataArrayTooBig occurs when more than 100 entries are
	// supplied in a MailboxDataArray element..
	ErrorMailboxDataArrayTooBig ResponseCode = "ErrorMailboxDataArrayTooBig"

	// ErrorMailboxFailover indicates that an attempt to access a mailbox
	// failed because the mailbox is in a failover process.
	ErrorMailboxFailover ResponseCode = "ErrorMailboxFailover"

	// ErrorMailboxHoldNotFound indicates that the mailbox hold was not found.
	ErrorMailboxHoldNotFound ResponseCode = "ErrorMailboxHoldNotFound"

	// ErrorMailboxLogonFailed  occurs when the connection to the mailbox to
	// get the calendar view information failed.
	ErrorMailboxLogonFailed ResponseCode = "ErrorMailboxLogonFailed"

	// ErrorMailboxMoveInProgress  indicates that the mailbox is being moved to
	// a different mailbox store or server. This error can also indicate that
	// the mailbox is on another server or mailbox database.
	ErrorMailboxMoveInProgress ResponseCode = "ErrorMailboxMoveInProgress"

	// ErrorMailboxStoreUnavailable indicates that one of the following error
	// conditions occurred:
	//  - The mailbox store is corrupt.
	//  - The mailbox store is being stopped.
	//  - The mailbox store is offline.
	//  - A network error occurred when an attempt was made to access the
	//    mailbox store.
	//  - The mailbox store is overloaded and cannot accept any more
	//    connections.
	//  - The mailbox store has been paused.
	ErrorMailboxStoreUnavailable ResponseCode = "ErrorMailboxStoreUnavailable"

	// ErrorMailRecipientNotFound occurs if the MailboxData element information
	// cannot be mapped to a valid mailbox account.
	ErrorMailRecipientNotFound ResponseCode = "ErrorMailRecipientNotFound"

	// ErrorMailTipsDisabled indicates that mail tips are disabled.
	ErrorMailTipsDisabled ResponseCode = "ErrorMailTipsDisabled"

	// ErrorManagedFolderAlreadyExists occurs if the managed folder that you
	// are trying to create already exists in a mailbox.
	ErrorManagedFolderAlreadyExists ResponseCode = "ErrorManagedFolderAlreadyExists"

	// ErrorManagedFolderNotFound occurs when the folder name that was
	// specified in the request does not map to a managed folder definition in
	// AD DS. You can only create instances of managed folders for folders that
	// are defined in AD DS. Check the name and try again.
	ErrorManagedFolderNotFound ResponseCode = "ErrorManagedFolderNotFound"

	// ErrorManagedFoldersRootFailure indicates that the managed folders root
	// was deleted from the mailbox or that a folder exists in the same parent
	// folder that has the name of the managed folder root. This will also
	// occur if the attempt to create the root managed folder fails.
	ErrorManagedFoldersRootFailure ResponseCode = "ErrorManagedFoldersRootFailure"

	// ErrorMeetingSuggestionGenerationFailed indicates that the suggestions
	// engine encountered a problem when it was trying to generate the
	// suggestions.
	ErrorMeetingSuggestionGenerationFailed ResponseCode = "ErrorMeetingSuggestionGenerationFailed"

	// ErrorMessageDispositionRequired occurs if the MessageDisposition
	// attribute is not set. This attribute is required for the following:
	//  - The CreateItem operation and the UpdateItem operation when the item
	//    being created or updated is a Message.
	//  - CancelCalendarItem, AcceptItem, DeclineItem, or TentativelyAcceptItem
	//    response objects.
	ErrorMessageDispositionRequired ResponseCode = "ErrorMessageDispositionRequired"

	// ErrorMessageSizeExceeded indicates that the message that you are trying
	// to send exceeds the allowed limits.
	ErrorMessageSizeExceeded ResponseCode = "ErrorMessageSizeExceeded"

	// ErrorMessageTrackingNoSuchDomain indicates that the given domain cannot
	// be found.
	ErrorMessageTrackingNoSuchDomain ResponseCode = "ErrorMessageTrackingNoSuchDomain"

	// ErrorMessageTrackingPermanentError indicates that the message tracking
	// service cannot track the message.
	ErrorMessageTrackingPermanentError ResponseCode = "ErrorMessageTrackingPermanentError"

	// ErrorMessageTrackingTransientError indicates that the message tracking
	// service is either down or busy. This error code indicates a transient
	// error. Clients can retry to connect to the server when this error is
	// received.
	ErrorMessageTrackingTransientError ResponseCode = "ErrorMessageTrackingTransientError"

	// ErrorMimeContentConversionFailed occurs when the MIME content is not a
	// valid iCal for a CreateItem operation. For a GetItem operation, this
	// response indicates that the MIME content could not be generated.
	ErrorMimeContentConversionFailed ResponseCode = "ErrorMimeContentConversionFailed"

	// ErrorMimeContentInvalid occurs when the MIME content is invalid.
	ErrorMimeContentInvalid ResponseCode = "ErrorMimeContentInvalid"

	// ErrorMimeContentInvalidBase64String occurs when the MIME content in the
	// request is not a valid base 64 string.
	ErrorMimeContentInvalidBase64String ResponseCode = "ErrorMimeContentInvalidBase64String"

	// ErrorMissingArgument indicates that a required argument was missing from
	// the request. The response message text indicates which argument to check.
	ErrorMissingArgument ResponseCode = "ErrorMissingArgument"

	// ErrorMissingEmailAddress indicates that you specified a distinguished
	// folder ID in the request, but the account that made the request does not
	// have a mailbox on the system. In that case, you must supply a Mailbox
	// sub-element under DistinguishedFolderId.
	ErrorMissingEmailAddress ResponseCode = "ErrorMissingEmailAddress"

	// ErrorMissingEmailAddressForManagedFolder indicates that you specified a
	// distinguished folder ID in the request, but the account that made the
	// request does not have a mailbox on the system. In that case, you must
	// supply a Mailbox sub-element under DistinguishedFolderId. This response
	// is returned from the CreateManagedFolder operation.
	ErrorMissingEmailAddressForManagedFolder ResponseCode = "ErrorMissingEmailAddressForManagedFolder"

	// ErrorMissingInformationEmailAddress occurs if the EmailAddress
	// (NonEmptyStringType) element is missing.
	ErrorMissingInformationEmailAddress ResponseCode = "ErrorMissingInformationEmailAddress"

	// ErrorMissingInformationReferenceItemId occurs if the ReferenceItemId is
	// missing.
	ErrorMissingInformationReferenceItemId ResponseCode = "ErrorMissingInformationReferenceItemId"

	// ErrorMissingInformationSharingFolderId code is never returned.
	// ErrorMissingInformationSharingFolderId ResponseCode = "ErrorMissingInformationSharingFolderId"

	// ErrorMissingItemForCreateItemAttachment is returned when an attempt is
	// made to not include the item element in the ItemAttachment element of a
	// CreateAttachment operation request.
	ErrorMissingItemForCreateItemAttachment ResponseCode = "ErrorMissingItemForCreateItemAttachment"

	// ErrorMissingManagedFolderId occurs when the policy IDs property,
	// property tag 0x6732, for the folder is missing. You should consider this
	// a corrupted folder.
	ErrorMissingManagedFolderId ResponseCode = "ErrorMissingManagedFolderId"

	// ErrorMissingRecipients indicates that you tried to send an item without
	// including recipients. Note that if you call the CreateItem operation
	// with a message disposition that causes the message to be sent, you will
	// get the following response code: ErrorInvalidRecipients.
	ErrorMissingRecipients ResponseCode = "ErrorMissingRecipients"

	// ErrorMissingUserIdInformation indicates that a UserId has not been fully
	// specified in a permissions set.
	ErrorMissingUserIdInformation ResponseCode = "ErrorMissingUserIdInformation"

	// ErrorMoreThanOneAccessModeSpecified indicates that you have specified
	// more than one ExchangeImpersonation element value within a request.
	ErrorMoreThanOneAccessModeSpecified ResponseCode = "ErrorMoreThanOneAccessModeSpecified"

	// ErrorMoveCopyFailed indicates that the move or copy operation failed.
	// Moving occurs in the CreateItem operation when you accept a meeting
	// request that is in the Deleted Items folder. In addition, if you decline
	// a meeting request, cancel a calendar item, or remove a meeting from your
	// calendar, it is moved to the Deleted Items folder.
	ErrorMoveCopyFailed ResponseCode = "ErrorMoveCopyFailed"

	// ErrorMoveDistinguishedFolder occurs if you try to move a distinguished
	// folder.
	ErrorMoveDistinguishedFolder ResponseCode = "ErrorMoveDistinguishedFolder"

	// ErrorMultiLegacyMailboxAccess occurs when a request attempts to access
	// multiple mailbox servers.
	// This error was introduced in Exchange 2013.
	ErrorMultiLegacyMailboxAccess ResponseCode = "ErrorMultiLegacyMailboxAccess"

	// ErrorNameResolutionMultipleResults occurs if the ResolveNames operation
	// returns more than one result or the ambiguous name that you specified
	// matches more than one object in the directory. The response code
	// includes the matched names in the response data.
	ErrorNameResolutionMultipleResults ResponseCode = "ErrorNameResolutionMultipleResults"

	// ErrorNameResolutionNoMailbox indicates that the caller does not have a
	// mailbox on the system. The ResolveNames operation or ExpandDL operation
	// is invalid for connecting a user without a mailbox.
	ErrorNameResolutionNoMailbox ResponseCode = "ErrorNameResolutionNoMailbox"

	// ErrorNameResolutionNoResults indicates that the ResolveNames operation
	// returns no results.
	ErrorNameResolutionNoResults ResponseCode = "ErrorNameResolutionNoResults"

	// ErrorNoApplicableProxyCASServersAvailable code MUST be returned when the
	// Web service cannot find a server to handle the request.
	ErrorNoApplicableProxyCASServersAvailable ResponseCode = "ErrorNoApplicableProxyCASServersAvailable"

	// ErrorNoCalendar occurs if there is no Calendar folder for the mailbox.
	ErrorNoCalendar ResponseCode = "ErrorNoCalendar"

	// ErrorNoDestinationCASDueToKerberosRequirements indicates that the
	// request referred to a mailbox in another Active Directory site, but no
	// Client Access servers in the destination site were configured for
	// Windows Authentication, and therefore the request could not be proxied.
	ErrorNoDestinationCASDueToKerberosRequirements ResponseCode = "ErrorNoDestinationCASDueToKerberosRequirements"

	// ErrorNoDestinationCASDueToSSLRequirements indicates that the request
	// referred to a mailbox in another Active Directory site, but no Client
	// Access servers in the destination site were configured for SSL
	// connections, and therefore the request could not be proxied.
	ErrorNoDestinationCASDueToSSLRequirements ResponseCode = "ErrorNoDestinationCASDueToSSLRequirements"

	// ErrorNoDestinationCASDueToVersionMismatch indicates that the request
	// referred to a mailbox in another Active Directory site, but no Client
	// Access servers in the destination site were of an acceptable product
	// version to receive the request, and therefore the request could not be
	// proxied.
	ErrorNoDestinationCASDueToVersionMismatch ResponseCode = "ErrorNoDestinationCASDueToVersionMismatch"

	// ErrorNoFolderClassOverride occurs if you set the FolderClass element
	// when you are creating an item other than a generic folder. For typed
	// folders such as CalendarFolder and TasksFolder, the folder class is
	// implied. Setting the folder class to a different folder type by using
	// the UpdateFolder operation results in the ErrorObjectTypeChanged
	// response. Instead, use a generic folder type but set the folder class
	// to the value that you require. Exchange Web Services will create the
	// correct strongly typed folder.
	ErrorNoFolderClassOverride ResponseCode = "ErrorNoFolderClassOverride"

	// ErrorNoFreeBusyAccess indicates that the caller does not have free/busy
	// viewing rights on the Calendar folder in question.
	ErrorNoFreeBusyAccess ResponseCode = "ErrorNoFreeBusyAccess"

	// ErrorNonExistentMailbox occurs in the following scenarios:
	//  - The e-mail address is empty in CreateManagedFolder.
	//  - The e-mail address does not refer to a valid account in a request
	//  that takes an e-mail address in the body or in the SOAP header, such as
	//  in an Exchange Impersonation call.
	ErrorNonExistentMailbox ResponseCode = "ErrorNonExistentMailbox"

	// ErrorNonPrimarySmtpAddress occurs when a caller passes in a non-primary
	// SMTP address. The response includes the correct SMTP address to use.
	ErrorNonPrimarySmtpAddress ResponseCode = "ErrorNonPrimarySmtpAddress"

	// ErrorNoPropertyTagForCustomProperties indicates that MAPI properties in
	// the custom range, 0x8000 and greater, cannot be referenced by property
	// tags. You must use the EWS Managed API PropertySetIdproperty or the EWS
	// ExtendedFieldURI element with the PropertySetId attribute.
	ErrorNoPropertyTagForCustomProperties ResponseCode = "ErrorNoPropertyTagForCustomProperties"

	// ErrorNoPublicFolderReplicaAvailable is not used.
	// ErrorNoPublicFolderReplicaAvailable ResponseCode = "ErrorNoPublicFolderReplicaAvailable"

	// ErrorNoPublicFolderServerAvailable code MUST be returned if no public
	// folder server is available or if the caller does not have a home public
	// server.
	ErrorNoPublicFolderServerAvailable ResponseCode = "ErrorNoPublicFolderServerAvailable"

	// ErrorNoRespondingCASInDestinationSite indicates that the request
	// referred to a mailbox in another Active Directory site, but none of the
	// Client Access servers in that site responded, and therefore the request
	// could not be proxied.
	ErrorNoRespondingCASInDestinationSite ResponseCode = "ErrorNoRespondingCASInDestinationSite"

	// ErrorNotAllowedExternalSharingByPolicy indicates that the caller tried
	// to grant permissions in its calendar or contacts folder to a user in
	// another organization, and the attempt failed.
	ErrorNotAllowedExternalSharingByPolicy ResponseCode = "ErrorNotAllowedExternalSharingByPolicy"

	// ErrorNotDelegate indicates that the user is not a delegate for the
	// mailbox. It is returned by the GetDelegate operation, the RemoveDelegate
	// operation, and the UpdateDelegate operation when the specified delegate
	// user is not found in the list of delegates.
	ErrorNotDelegate ResponseCode = "ErrorNotDelegate"

	// ErrorNotEnoughMemory indicates that the operation could not be completed
	// because of insufficient memory.
	ErrorNotEnoughMemory ResponseCode = "ErrorNotEnoughMemory"

	// ErrorNotSupportedSharingMessage indicates that the sharing message is
	// not supported.
	ErrorNotSupportedSharingMessage ResponseCode = "ErrorNotSupportedSharingMessage"

	// ErrorObjectTypeChanged occurs if the object type changed.
	ErrorObjectTypeChanged ResponseCode = "ErrorObjectTypeChanged"

	// ErrorOccurrenceCrossingBoundary occurs when the Start or End time of an
	// occurrence is updated so that the occurrence is scheduled to happen
	// earlier or later than the corresponding previous or next occurrence.
	ErrorOccurrenceCrossingBoundary ResponseCode = "ErrorOccurrenceCrossingBoundary"

	// ErrorOccurrenceTimeSpanTooBig indicates that the time allotment for a
	// given occurrence overlaps with another occurrence of the same recurring
	// item. This response also occurs when the length in minutes of a given
	// occurrence is larger than TryInt32.MaxValue.
	ErrorOccurrenceTimeSpanTooBig ResponseCode = "ErrorOccurrenceTimeSpanTooBig"

	// ErrorOperationNotAllowedWithPublicFolderRoot indicates that the current
	// operation is not valid for the public folder root.
	ErrorOperationNotAllowedWithPublicFolderRoot ResponseCode = "ErrorOperationNotAllowedWithPublicFolderRoot"

	// ErrorOrganizationNotFederated indicates that the requester's
	// organization is not federated so the requester cannot create sharing
	// messages to send to an external user or cannot accept sharing messages
	// received from an external user.
	ErrorOrganizationNotFederated ResponseCode = "ErrorOrganizationNotFederated"

	// ErrorParentFolderIdRequired is not used.
	// ErrorParentFolderIdRequired ResponseCode = "ErrorParentFolderIdRequired"

	// ErrorParentFolderNotFound occurs in the CreateFolder operation when the
	// parent folder is not found.
	ErrorParentFolderNotFound ResponseCode = "ErrorParentFolderNotFound"

	// ErrorPasswordChangeRequired indicates that you must change your password
	// before you can access this mailbox. This occurs when a new account has
	// been created and the administrator indicated that the user must change
	// the password at first logon. You cannot update the password by using
	// Exchange Web Services. You must use a tool such as Microsoft Office
	// Outlook Web App to change your password.
	ErrorPasswordChangeRequired ResponseCode = "ErrorPasswordChangeRequired"

	// ErrorPasswordExpired indicates that the password has expired. You cannot
	// change the password by using Exchange Web Services. You must use a tool
	// such as Outlook Web App to change your password.
	ErrorPasswordExpired ResponseCode = "ErrorPasswordExpired"

	// ErrorPermissionNotAllowedByPolicy indicates that the requester tried to
	// grant permissions in its calendar or contacts folder to an external user
	// but the sharing policy assigned to the requester indicates that the
	// requested permission level is higher than what the sharing policy allows.
	ErrorPermissionNotAllowedByPolicy ResponseCode = "ErrorPermissionNotAllowedByPolicy"

	// ErrorPhoneNumberNotDialable indicates that the telephone number was not
	// in the correct form.
	ErrorPhoneNumberNotDialable ResponseCode = "ErrorPhoneNumberNotDialable"

	// ErrorPropertyUpdate indicates that the update failed because of invalid
	// property values. The response message includes the invalid property
	// paths.
	ErrorPropertyUpdate ResponseCode = "ErrorPropertyUpdate"

	// ErrorPromptPublishingOperationFailed is intended for internal use only.
	// This error was introduced in Exchange 2013.
	ErrorPromptPublishingOperationFailed ResponseCode = "ErrorPromptPublishingOperationFailed"

	// ErrorPropertyValidationFailure is not used.
	// ErrorPropertyValidationFailure ResponseCode = "ErrorPropertyValidationFailure"

	// ErrorProxiedSubscriptionCallFailure indicates that the request referred
	// to a subscription that exists on another Client Access server, but an
	// attempt to proxy the request to that Client Access server failed.
	ErrorProxiedSubscriptionCallFailure ResponseCode = "ErrorProxiedSubscriptionCallFailure"

	// ErrorProxyCallFailed is not used.
	// ErrorProxyCallFailed ResponseCode = "ErrorProxyCallFailed"

	// ErrorProxyGroupSidLimitExceeded indicates that the request referred to a
	// mailbox in another Active Directory site, and the original caller is a
	// member of more than 3,000 groups.
	ErrorProxyGroupSidLimitExceeded ResponseCode = "ErrorProxyGroupSidLimitExceeded"

	// ErrorProxyRequestNotAllowed indicates that the request that Exchange Web
	// Services sent to another Client Access server when trying to fulfill a
	// GetUserAvailabilityRequest request was invalid. This response code
	// typically indicates that a configuration or rights error has occurred,
	// or that someone tried unsuccessfully to mimic an availability proxy
	// request.
	ErrorProxyRequestNotAllowed ResponseCode = "ErrorProxyRequestNotAllowed"

	// ErrorProxyRequestProcessingFailed indicates that Exchange Web Services
	// tried to proxy an availability request to another Client Access server
	// for fulfillment, but the request failed. This response can be caused by
	// network connectivity issues or request timeout issues.
	ErrorProxyRequestProcessingFailed ResponseCode = "ErrorProxyRequestProcessingFailed"

	// ErrorProxyServiceDiscoveryFailed code must be returned if the Web
	// service cannot determine whether the request is to run on the target
	// server or will be proxied to another server.
	ErrorProxyServiceDiscoveryFailed ResponseCode = "ErrorProxyServiceDiscoveryFailed"

	// ErrorProxyTokenExpired is not used.
	// ErrorProxyTokenExpired ResponseCode = "ErrorProxyTokenExpired"

	// ErrorPublicFolderMailboxDiscoveryFailed occurs when the public folder
	// mailbox URL cannot be found. This error is intended for internal use
	// only. This error was introduced in Exchange 2013.
	ErrorPublicFolderMailboxDiscoveryFailed ResponseCode = "ErrorPublicFolderMailboxDiscoveryFailed"

	// ErrorPublicFolderOperationFailed occurs when an attempt is made to
	// access a public folder and the attempt is unsuccessful.
	// This error was introduced in Exchange 2013.
	ErrorPublicFolderOperationFailed ResponseCode = "ErrorPublicFolderOperationFailed"

	// ErrorPublicFolderRequestProcessingFailed occurs when the recipient that
	// was passed to the GetUserAvailability operation is located on a computer
	// that is running a version of Exchange LogClientStart that is earlier than
	// Exchange 2007, and the request to retrieve free/busy information for the
	// recipient from the public folder server failed.
	ErrorPublicFolderRequestProcessingFailed ResponseCode = "ErrorPublicFolderRequestProcessingFailed"

	// ErrorPublicFolderServerNotFound occurs when the recipient that was
	// passed to the GetUserAvailability operation is located on a computer
	// that is running a version of Exchange LogClientStart that is earlier than
	// Exchange 2007, and the request to retrieve free/busy information for
	// the recipient from the public folder server failed because the
	// organizational unit did not have an associated public folder server.
	ErrorPublicFolderServerNotFound ResponseCode = "ErrorPublicFolderServerNotFound"

	// ErrorPublicFolderSyncException occurs when a synchronization operation
	// succeeds against the primary public folder mailbox but does not succeed
	// against the secondary public folder mailbox.
	// This error was introduced in Exchange 2013.
	ErrorPublicFolderSyncException ResponseCode = "ErrorPublicFolderSyncException"

	// ErrorQueryFilterTooLong indicates that the search folder restriction may
	// be valid, but it is not supported by EWS. Exchange Web Services limits
	// restrictions to contain a maximum of 255 filter expressions. If you try
	// to bind to an existing search folder that exceeds 255, this response
	// code is returned.
	ErrorQueryFilterTooLong ResponseCode = "ErrorQueryFilterTooLong"

	// ErrorQuotaExceeded occurs when the mailbox quota is exceeded.
	ErrorQuotaExceeded ResponseCode = "ErrorQuotaExceeded"

	// ErrorReadEventsFailed is returned by the GetEvents operation or push
	// notifications when a failure occurs while retrieving event information.
	// When this error is returned, the subscription is deleted. Re-create the
	// event synchronization based on a last known watermark.
	ErrorReadEventsFailed ResponseCode = "ErrorReadEventsFailed"

	// ErrorReadReceiptNotPending is returned by the CreateItem operation if an
	// attempt was made to suppress a read receipt when the message sender did
	// not request a read receipt on the message or if the message is in the
	// Junk E-mail folder.
	ErrorReadReceiptNotPending ResponseCode = "ErrorReadReceiptNotPending"

	// ErrorRecurrenceEndDateTooBig occurs when the end date for the recurrence
	// is after 9/1/4500.
	ErrorRecurrenceEndDateTooBig ResponseCode = "ErrorRecurrenceEndDateTooBig"

	// ErrorRecurrenceHasNoOccurrence occurs when the specified recurrence does
	// not have any occurrence instances in the specified range.
	ErrorRecurrenceHasNoOccurrence ResponseCode = "ErrorRecurrenceHasNoOccurrence"

	// ErrorRemoveDelegatesFailed indicates that the delegate list failed to be
	// saved after delegates were removed.
	ErrorRemoveDelegatesFailed ResponseCode = "ErrorRemoveDelegatesFailed"

	// ErrorRequestAborted is not used.
	// ErrorRequestAborted ResponseCode = "ErrorRequestAborted"

	// ErrorRequestStreamTooBig occurs when the request stream is larger than
	// 400 KB.
	ErrorRequestStreamTooBig ResponseCode = "ErrorRequestStreamTooBig"

	// ErrorRequiredPropertyMissing is returned when a required property is
	// missing in a CreateAttachment operation request. The missing property
	// URI is included in the response.
	ErrorRequiredPropertyMissing ResponseCode = "ErrorRequiredPropertyMissing"

	// ErrorResolveNamesInvalidFolderType indicates that the caller has
	// specified a folder that is not a contacts folder to the ResolveNames
	// operation.
	ErrorResolveNamesInvalidFolderType ResponseCode = "ErrorResolveNamesInvalidFolderType"

	// ErrorResolveNamesOnlyOneContactsFolderAllowed indicates that the caller
	// has specified more than one contacts folder to the ResolveNames
	// operation.
	ErrorResolveNamesOnlyOneContactsFolderAllowed ResponseCode = "ErrorResolveNamesOnlyOneContactsFolderAllowed"

	// ErrorResponseSchemaValidation is not used.
	ErrorResponseSchemaValidation ResponseCode = "ErrorResponseSchemaValidation"

	// ErrorRestrictionTooLong occurs if the restriction contains more than
	// 255 nodes.
	ErrorRestrictionTooLong ResponseCode = "ErrorRestrictionTooLong"

	// ErrorRestrictionTooComplex occurs when the restriction cannot be
	// evaluated by Exchange Web Services.
	ErrorRestrictionTooComplex ResponseCode = "ErrorRestrictionTooComplex"

	// ErrorResultSetTooBig indicates that the number of calendar entries for
	// a given recipient exceeds the allowed limit of 1000. Reduce the window
	// and try again.
	ErrorResultSetTooBig ResponseCode = "ErrorResultSetTooBig"

	// ErrorSavedItemFolderNotFound occurs when the SavedItemFolderId is not
	// found.
	ErrorSavedItemFolderNotFound ResponseCode = "ErrorSavedItemFolderNotFound"

	// ErrorSchemaValidation occurs when the request cannot be validated
	// against the schema.
	ErrorSchemaValidation ResponseCode = "ErrorSchemaValidation"

	// ErrorSearchFolderNotInitialized indicates that the search folder was
	// created, but the search criteria were never set on the folder. This only
	// occurs when you access corrupted search folders that were created by
	// using another API or client. To fix this error, use the UpdateFolder
	// operation to set the SearchParameters element to include the restriction
	// that should be on the folder.
	ErrorSearchFolderNotInitialized ResponseCode = "ErrorSearchFolderNotInitialized"

	// ErrorSendAsDenied occurs when both of the following conditions occur:
	//  - A user has been granted CanActAsOwner permissions but is not granted
	//    delegate rights on the principal's mailbox.
	//  - The same user tries to create and send an e-mail message in the
	//    principal's mailbox by using the SendAndSaveCopy option.
	// The result is an ErrorSendAsDenied error and the creation of the e-mail
	// message in the principal's Drafts folder.
	ErrorSendAsDenied ResponseCode = "ErrorSendAsDenied"

	// ErrorSendMeetingCancellationsRequired is returned by the DeleteItem
	// operation if the SendMeetingCancellations attribute is missing from the
	// request and the item to delete is a calendar item.
	ErrorSendMeetingCancellationsRequired ResponseCode = "ErrorSendMeetingCancellationsRequired"

	// ErrorSendMeetingInvitationsOrCancellationsRequired is returned by the
	// UpdateItem operation if the SendMeetingInvitationsOrCancellations
	// attribute is missing from the request and the item to update is a
	// calendar item.
	ErrorSendMeetingInvitationsOrCancellationsRequired ResponseCode = "ErrorSendMeetingInvitationsOrCancellationsRequired"

	// ErrorSendMeetingInvitationsRequired is returned by the CreateItem
	// operation if the SendMeetingInvitations attribute is missing from the
	// request and the item to create is a calendar item.
	ErrorSendMeetingInvitationsRequired ResponseCode = "ErrorSendMeetingInvitationsRequired"

	// ErrorSentMeetingRequestUpdate indicates that after the organizer sends
	// a meeting request, the request cannot be updated. To modify the meeting,
	// modify the calendar item, not the meeting request.
	ErrorSentMeetingRequestUpdate ResponseCode = "ErrorSentMeetingRequestUpdate"

	// ErrorSentTaskRequestUpdate indicates that after the task initiator sends
	// a task request, that request cannot be updated.
	ErrorSentTaskRequestUpdate ResponseCode = "ErrorSentTaskRequestUpdate"

	// ErrorServerBusy occurs when the server is busy.
	ErrorServerBusy ResponseCode = "ErrorServerBusy"

	// ErrorServiceDiscoveryFailed indicates that Exchange Web Services tried
	// to proxy a user availability request to the appropriate forest for the
	// recipient, but it could not determine where to send the request because
	// of a service discovery failure.
	ErrorServiceDiscoveryFailed ResponseCode = "ErrorServiceDiscoveryFailed"

	// ErrorSharingNoExternalEwsAvailable indicates that the external URL
	// property has not been set in the Active Directory database.
	// It specifies that the external URL property has not been set in the
	// Active Directory database. This error code MUST be returned if the
	// external URL property has not been set in the Active Directory database.
	ErrorSharingNoExternalEwsAvailable ResponseCode = "ErrorSharingNoExternalEwsAvailable"

	// ErrorSharingSynchronizationFailed indicates that an attempt at
	// synchronizing a sharing folder failed. This error code is returned when
	// the following occurs:
	//  - The subscription for a sharing folder is not found.
	//  - The sharing folder is not found.
	//  - The corresponding directory user is not found.
	//  - The user no longer exists.
	//  - The appointment is invalid.
	//  - The contact item is invalid.
	//  - There is a communication failure with the remote server.
	ErrorSharingSynchronizationFailed ResponseCode = "ErrorSharingSynchronizationFailed"

	// ErrorStaleObject occurs in an UpdateItem operation or a SendItem
	// operation when the change key is not up-to-date or was not supplied.
	// Call the GetItem operation to retrieve an updated change key and then
	// try the operation again.
	ErrorStaleObject ResponseCode = "ErrorStaleObject"

	// ErrorSubmissionQuotaExceeded Indicates that a user cannot immediately
	// send more requests because the submission quota has been reached.
	ErrorSubmissionQuotaExceeded ResponseCode = "ErrorSubmissionQuotaExceeded"

	// ErrorSubscriptionAccessDenied occurs when you try to access a
	// subscription by using an account that did not create that subscription.
	// Each subscription can only be accessed by the creator of the
	// subscription.
	ErrorSubscriptionAccessDenied ResponseCode = "ErrorSubscriptionAccessDenied"

	// ErrorSubscriptionDelegateAccessNotSupported indicates that you cannot
	// create a subscription if you are not the owner or do not have owner
	// access to the mailbox.
	ErrorSubscriptionDelegateAccessNotSupported ResponseCode = "ErrorSubscriptionDelegateAccessNotSupported"

	// ErrorSubscriptionNotFound occurs if the subscription that corresponds to
	// the specified SubscriptionId (GetEvents) is not found. The subscription
	// may have expired, the Exchange Web Services process may have been
	// restarted, or an invalid subscription was passed in. If the subscription
	// was valid, re-create the subscription with the latest watermark. This is
	// returned by the Unsubscribe operation or the GetEvents operation
	// responses.
	ErrorSubscriptionNotFound ResponseCode = "ErrorSubscriptionNotFound"

	// ErrorSubscriptionUnsubsribed code must be returned when a request is
	// made for a subscription that has been unsubscribed.
	ErrorSubscriptionUnsubsribed ResponseCode = "ErrorSubscriptionUnsubsribed"

	// ErrorSyncFolderNotFound is returned by the SyncFolderItems operation if
	// the parent folder that is specified cannot be found.
	ErrorSyncFolderNotFound ResponseCode = "ErrorSyncFolderNotFound"

	// ErrorTeamMailboxNotFound indicates that a team mailbox was not found.
	// This error was introduced in Exchange 2013.
	ErrorTeamMailboxNotFound ResponseCode = "ErrorTeamMailboxNotFound"

	// ErrorTeamMailboxNotLinkedToSharePoint indicates that a team mailbox was
	// found but that it is not linked to a SharePoint LogClientStart.
	// This error was introduced in Exchange 2013.
	ErrorTeamMailboxNotLinkedToSharePoint ResponseCode = "ErrorTeamMailboxNotLinkedToSharePoint"

	// ErrorTeamMailboxUrlValidationFailed indicates that a team mailbox was
	// found but that the link to the SharePoint LogClientStart is not valid.
	// This error was introduced in Exchange 2013.
	ErrorTeamMailboxUrlValidationFailed ResponseCode = "ErrorTeamMailboxUrlValidationFailed"

	// ErrorTeamMailboxNotAuthorizedOwner code is not used.
	// This error was introduced in Exchange 2013.
	ErrorTeamMailboxNotAuthorizedOwner ResponseCode = "ErrorTeamMailboxNotAuthorizedOwner"

	// ErrorTeamMailboxActiveToPendingDelete code is not used.
	// This error was introduced in Exchange 2013.
	ErrorTeamMailboxActiveToPendingDelete ResponseCode = "ErrorTeamMailboxActiveToPendingDelete"

	// ErrorTeamMailboxFailedSendingNotifications indicates that an attempt to
	// send a notification to the team mailbox owners was unsuccessful.
	// This error was introduced in Exchange 2013.
	ErrorTeamMailboxFailedSendingNotifications ResponseCode = "ErrorTeamMailboxFailedSendingNotifications"

	// ErrorTeamMailboxErrorUnknown indicates a general error that can occur
	// when trying to access a team mailbox. Try submitting the request at a
	// later time. This error was introduced in Exchange 2013.
	ErrorTeamMailboxErrorUnknown ResponseCode = "ErrorTeamMailboxErrorUnknown"

	// ErrorTimeIntervalTooBig indicates that the time window that was
	// specified is larger than the allowed limit. By default, the allowed
	// limit is 42.
	ErrorTimeIntervalTooBig ResponseCode = "ErrorTimeIntervalTooBig"

	// ErrorTimeoutExpired occurs when there is not enough time to complete the
	// processing of the request.
	ErrorTimeoutExpired ResponseCode = "ErrorTimeoutExpired"

	// ErrorTimeZone indicates that there is a time zone error.
	ErrorTimeZone ResponseCode = "ErrorTimeZone"

	// ErrorToFolderNotFound indicates that the destination folder does not
	// exist.
	ErrorToFolderNotFound ResponseCode = "ErrorToFolderNotFound"

	// ErrorTokenSerializationDenied occurs if the caller tries to do a Token
	// serialization request but does not have the
	// ms-Exch-EPI-TokenSerialization right on the Client Access server.
	ErrorTokenSerializationDenied ResponseCode = "ErrorTokenSerializationDenied"

	// ErrorTooManyObjectsOpened occurs when the internal limit on open objects
	// has been exceeded.
	ErrorTooManyObjectsOpened ResponseCode = "ErrorTooManyObjectsOpened"

	// ErrorUnifiedMessagingDialPlanNotFound indicates that a user's dial plan
	// is not available.
	ErrorUnifiedMessagingDialPlanNotFound ResponseCode = "ErrorUnifiedMessagingDialPlanNotFound"

	// ErrorUnifiedMessagingReportDataNotFound is intended for internal use
	// only. This error was introduced in Exchange 2013.
	ErrorUnifiedMessagingReportDataNotFound ResponseCode = "ErrorUnifiedMessagingReportDataNotFound"

	// ErrorUnifiedMessagingPromptNotFound is intended for internal use only.
	// This error was introduced in Exchange 2013.
	ErrorUnifiedMessagingPromptNotFound ResponseCode = "ErrorUnifiedMessagingPromptNotFound"

	// ErrorUnifiedMessagingRequestFailed indicates that the user could not be
	// found.
	ErrorUnifiedMessagingRequestFailed ResponseCode = "ErrorUnifiedMessagingRequestFailed"

	// ErrorUnifiedMessagingServerNotFound indicates that a valid server for
	// the dial plan can be found to handle the request.
	ErrorUnifiedMessagingServerNotFound ResponseCode = "ErrorUnifiedMessagingServerNotFound"

	// ErrorUnableToGetUserOofSettings is not used.
	// ErrorUnableToGetUserOofSettings ResponseCode = "ErrorUnableToGetUserOofSettings"

	// ErrorUnableToRemoveImContactFromGroup occurs when an unsuccessful
	// attempt is made to remove an IM contact from a group.
	// This error was introduced in Exchange 2013.
	ErrorUnableToRemoveImContactFromGroup ResponseCode = "ErrorUnableToRemoveImContactFromGroup"

	// ErrorUnsupportedCulture occurs when you try to set the Culture property
	// to a value that is not parsable by the System.Globalization.CultureInfo
	// class.
	ErrorUnsupportedCulture ResponseCode = "ErrorUnsupportedCulture"

	// ErrorUnsupportedMapiPropertyType occurs when a caller tries to use
	// extended properties of types object, object array, error, or null.
	ErrorUnsupportedMapiPropertyType ResponseCode = "ErrorUnsupportedMapiPropertyType"

	// ErrorUnsupportedMimeConversion occurs when you are trying to retrieve or
	// set MIME content for an item other than a PostItem, Message, or
	// CalendarItem object.
	ErrorUnsupportedMimeConversion ResponseCode = "ErrorUnsupportedMimeConversion"

	// ErrorUnsupportedPathForQuery occurs when the caller passes a property
	// that is invalid for a query. This can occur when calculated properties
	// are used.
	ErrorUnsupportedPathForQuery ResponseCode = "ErrorUnsupportedPathForQuery"

	// ErrorUnsupportedPathForSortGroup occurs when the caller passes a
	// property that is invalid for a sort or group by property. This can occur
	// when calculated properties are used.
	ErrorUnsupportedPathForSortGroup ResponseCode = "ErrorUnsupportedPathForSortGroup"

	// ErrorUnsupportedPropertyDefinition is not used.
	ErrorUnsupportedPropertyDefinition ResponseCode = "ErrorUnsupportedPropertyDefinition"

	// ErrorUnsupportedQueryFilter indicates that the search folder restriction
	// may be valid, but it is not supported by EWS.
	ErrorUnsupportedQueryFilter ResponseCode = "ErrorUnsupportedQueryFilter"

	// ErrorUnsupportedRecurrence indicates that the specified recurrence is
	// not supported for tasks.
	ErrorUnsupportedRecurrence ResponseCode = "ErrorUnsupportedRecurrence"

	// ErrorUnsupportedSubFilter is not used.
	// ErrorUnsupportedSubFilter ResponseCode = "ErrorUnsupportedSubFilter"

	// ErrorUnsupportedTypeForConversion indicates that Exchange Web Services
	// found a property type in the store but it cannot generate XML for the
	// property type.
	ErrorUnsupportedTypeForConversion ResponseCode = "ErrorUnsupportedTypeForConversion"

	// ErrorUpdateDelegatesFailed indicates that the delegate list failed to be
	// saved after delegates were updated.
	ErrorUpdateDelegatesFailed ResponseCode = "ErrorUpdateDelegatesFailed"

	// ErrorUpdatePropertyMismatch occurs when the single property path that is
	// listed in a change description does not match the single property that
	// is being set within the actual Item or Folder object.
	ErrorUpdatePropertyMismatch ResponseCode = "ErrorUpdatePropertyMismatch"

	// ErrorUserNotUnifiedMessagingEnabled indicates that the requester is not
	// enabled.
	ErrorUserNotUnifiedMessagingEnabled ResponseCode = "ErrorUserNotUnifiedMessagingEnabled"

	// ErrorUserNotAllowedByPolicy indicates that the requester tried to grant
	// permissions in its calendar or contacts folder to an external user but
	// the sharing policy assigned to the requester indicates that the domain
	// of the external user is not listed in the policy.
	ErrorUserNotAllowedByPolicy ResponseCode = "ErrorUserNotAllowedByPolicy"

	// ErrorUserWithoutFederatedProxyAddress indicates that the requester's
	// organization has a set of federated domains but the requester's
	// organization does not have any SMTP proxy addresses with one of the
	// federated domains.
	ErrorUserWithoutFederatedProxyAddress ResponseCode = "ErrorUserWithoutFederatedProxyAddress"

	// ErrorValueOutOfRange indicates that a calendar view start date or end
	// date was set to 1/1/0001 12:00:00 AM or 12/31/9999 11:59:59 PM.
	ErrorValueOutOfRange ResponseCode = "ErrorValueOutOfRange"

	// ErrorVirusDetected indicates that the Exchange store detected a virus in
	// the message.
	ErrorVirusDetected ResponseCode = "ErrorVirusDetected"

	// ErrorVirusMessageDeleted indicates that the Exchange store detected a
	// virus in the message and deleted it.
	ErrorVirusMessageDeleted ResponseCode = "ErrorVirusMessageDeleted"

	// ErrorVoiceMailNotImplemented is not used.
	// ErrorVoiceMailNotImplemented ResponseCode = "ErrorVoiceMailNotImplemented"

	// ErrorWebRequestInInvalidState is not used.
	// ErrorWebRequestInInvalidState ResponseCode = "ErrorWebRequestInInvalidState"

	// ErrorWin32InteropError indicates that there was an internal failure
	// during communication with unmanaged code.
	ErrorWin32InteropError ResponseCode = "ErrorWin32InteropError"

	// ErrorWorkingHoursSaveFailed is not used.
	// ErrorWorkingHoursSaveFailed ResponseCode = "ErrorWorkingHoursSaveFailed"

	// ErrorWorkingHoursXmlMalformed is not used.
	// ErrorWorkingHoursXmlMalformed ResponseCode = "ErrorWorkingHoursXmlMalformed"

	// ErrorWrongServerVersion indicates that a request can only be made to a
	// server that is the same version as the mailbox server.
	ErrorWrongServerVersion ResponseCode = "ErrorWrongServerVersion"

	// ErrorWrongServerVersionDelegate indicates that a request was made by a
	// delegate that has a different server version than the principal's
	// mailbox server.
	ErrorWrongServerVersionDelegate ResponseCode = "ErrorWrongServerVersionDelegate"

	// ErrorMissingInformationSharingFolderId is never returned.
	// ErrorMissingInformationSharingFolderId ResponseCode = "ErrorMissingInformationSharingFolderId"

	// ErrorDuplicateSOAPHeader specifies that there are duplicate SOAP headers. (duplicate)
	// ErrorDuplicateSOAPHeader ResponseCode = "ErrorDuplicateSOAPHeader"

	// ErrorSharingSynchronizationFailed specifies that an attempt at
	// synchronizing a sharing folder failed. This error code MUST be returned
	// when:
	//  - The subscription for a sharing folder is not found.
	//  - The sharing folder was not found.
	//  - The corresponding directory user was not found.
	//  - The user no longer exists.
	//  - The appointment is invalid.
	//  - The contact item is invalid.
	//  - There was a communication failure with the remote server.
	// (duplicate)
	// ErrorSharingSynchronizationFailed ResponseCode = "ErrorSharingSynchronizationFailed"

	// ErrorSharingNoExternalEwsAvailable specifies that the external URL
	// property has not been set in the Active Directory database. This error
	// code MUST be returned if the external URL property has not been set in
	// the Active Directory database. (duplicate)
	// ErrorSharingNoExternalEwsAvailable ResponseCode = "ErrorSharingNoExternalEwsAvailable"

	// ErrorFreeBusyDLLimitReached specifies that the maximum group member
	// count has been reached for obtaining free/busy information for a
	// distribution list. This error MUST be returned when the maximum group
	// member count has been reached for obtaining free/busy information for a
	// distribution list. (duplicate)
	// ErrorFreeBusyDLLimitReached ResponseCode = "ErrorFreeBusyDLLimitReached"

	// ErrorInvalidGetSharingFolderRequest specifies that the DataType and
	// ShareFolderId element are both present in a request. This error code
	// MUST be returned if the DataType and ShareFolderId element are both
	// present in a request. (duplicate)
	// ErrorInvalidGetSharingFolderRequest ResponseCode = "ErrorInvalidGetSharingFolderRequest"

	// ErrorNotAllowedExternalSharingByPolicy specifies that the caller
	// attempted to grant permissions in its calendar or contacts folder to a
	// user in another organization and the attempt failed. This error code
	// MUST be returned when the sharing policy is disabled for the caller or
	// when the sharing policy assigned to the caller disallows sharing for
	// the requested level or the requested folder type. (duplicate)
	// ErrorNotAllowedExternalSharingByPolicy ResponseCode = "ErrorNotAllowedExternalSharingByPolicy"

	// ErrorUserNotAllowedByPolicy specifies that the requestor attempted to
	// grant permissions in its calendar or contacts folder to an external user
	// but the sharing policy assigned to the requestor specifies that the
	// domain of the external user is not listed in the policy. (duplicate)
	// ErrorUserNotAllowedByPolicy ResponseCode = "ErrorUserNotAllowedByPolicy"

	// ErrorPermissionNotAllowedByPolicy specifies that the requestor attempted
	// to grant permissions in its calendar or contacts folder to an external
	// user but the sharing policy assigned to the requestor specifies that the
	// requested permission level is higher is than what the sharing policy
	// allows. (duplicate)
	// ErrorPermissionNotAllowedByPolicy ResponseCode = "ErrorPermissionNotAllowedByPolicy"

	// ErrorOrganizationNotFederated specifies that the requestor's
	// organization is not federated so the requestor cannot create sharing
	// messages to send to an external user or cannot accept sharing messages
	// received from an external user. This error code MUST be returned if the
	// requestor's organization is not federated. (duplicate)
	// ErrorOrganizationNotFederated ResponseCode = "ErrorOrganizationNotFederated"

	// ErrorMailboxFailover specifies that an attempt to access a mailbox
	// failed because the mailbox is in a failover process. (duplicate)
	// ErrorMailboxFailover ResponseCode = "ErrorMailboxFailover"

	// ErrorMessageTrackingPermanentError specifies that the message tracking
	// service cannot track the message. (duplicate)
	// ErrorMessageTrackingPermanentError ResponseCode = "ErrorMessageTrackingPermanentError"

	// ErrorMessageTrackingTransientError specifies that either the message
	// tracking service is down or busy. This error code specifies a transient
	// error. Clients can retry to connect to the server when this error is
	// received. (duplicate)
	// ErrorMessageTrackingTransientError ResponseCode = "ErrorMessageTrackingTransientError"

	// ErrorMessageTrackingNoSuchDomain specifies that the given domain cannot
	// be found. (duplicate)
	// ErrorMessageTrackingNoSuchDomain ResponseCode = "ErrorMessageTrackingNoSuchDomain"

	// ErrorUserWithoutFederatedProxyAddress specifies that the requestor's
	// organization has a set of federated domains but the requestor's
	// organization does not have any SMTP proxy addresses with one of the
	// federated domains. (duplicate)
	// ErrorUserWithoutFederatedProxyAddress ResponseCode = "ErrorUserWithoutFederatedProxyAddress"

	// ErrorInvalidOrganizationRelationshipForFreeBusy specifies that a caller
	// requested free/busy information for a user in another organization but
	// the organizational relationship does not have free/busy enabled.
	// (duplicate)
	// ErrorInvalidOrganizationRelationshipForFreeBusy ResponseCode = "ErrorInvalidOrganizationRelationshipForFreeBusy"

	// ErrorInvalidFederatedOrganizationId specifies that the requestor's
	// organization federation objects are not properly configured. (duplicate)
	// ErrorInvalidFederatedOrganizationId ResponseCode = "ErrorInvalidFederatedOrganizationId"

	// ErrorInvalidExternalSharingSubscriber that a sharing message is not
	// intended for the caller. (duplicate)
	// ErrorInvalidExternalSharingSubscriber ResponseCode = "ErrorInvalidExternalSharingSubscriber"

	// ErrorInvalidSharingData that the sharing metadata is not valid. This can
	// be caused by invalid XML. (duplicate)
	// ErrorInvalidSharingData ResponseCode = "ErrorInvalidSharingData"

	// ErrorInvalidSharingMessage that the sharing message is not valid. This
	// can be caused by a missing property. (duplicate)
	// ErrorInvalidSharingMessage ResponseCode = "ErrorInvalidSharingMessage"

	// Specifies that the sharing message is not supported. (duplicate)
	// ErrorNotSupportedSharingMessage ResponseCode = "ErrorNotSupportedSharingMessage"

	// ErrorApplyConversationActionFailed MUST be returned if an action cannot
	// be applied to one or more items in the conversation.
	ErrorApplyConversationActionFailed ResponseCode = "ErrorApplyConversationActionFailed"

	// ErrorInboxRulesValidationError MUST be returned if any rule does not
	// validate.
	ErrorInboxRulesValidationError ResponseCode = "ErrorInboxRulesValidationError"

	// ErrorOutlookRuleBlobExists MUST be returned when an attempt to manage
	// Inbox rules occurs after another client has accessed the Inbox rules.
	ErrorOutlookRuleBlobExists ResponseCode = "ErrorOutlookRuleBlobExists"

	// ErrorRulesOverQuota MUST be returned when a user's rule quota has been
	// exceeded.
	ErrorRulesOverQuota ResponseCode = "ErrorRulesOverQuota"

	// ErrorNewEventStreamConnectionOpened MUST be returned to the first
	// subscription connection if a second subscription connection is opened.
	ErrorNewEventStreamConnectionOpened ResponseCode = "ErrorNewEventStreamConnectionOpened"

	// ErrorMissedNotificationEvents MUST be returned when event notifications
	// are missed.
	ErrorMissedNotificationEvents ResponseCode = "ErrorMissedNotificationEvents"

	// ErrorDuplicateLegacyDistinguishedName is returned when there are
	// duplicate legacy distinguished names in Active Directory Domain
	// Services (AD DS). This error was introduced in Exchange 2013.
	ErrorDuplicateLegacyDistinguishedName ResponseCode = "ErrorDuplicateLegacyDistinguishedName"

	// ErrorInvalidClientAccessTokenRequest indicates that a request to get a
	// client access token was not valid.
	// This error was introduced in Exchange 2013.
	ErrorInvalidClientAccessTokenRequest ResponseCode = "ErrorInvalidClientAccessTokenRequest"

	// ErrorNoSpeechDetected is intended for internal use only.
	// This error was introduced in Exchange 2013.
	ErrorNoSpeechDetected ResponseCode = "ErrorNoSpeechDetected"

	// ErrorUMServerUnavailable is intended for internal use only.
	// This error was introduced in Exchange 2013.
	ErrorUMServerUnavailable ResponseCode = "ErrorUMServerUnavailable"

	// ErrorRecipientNotFound is intended for internal use only.
	// This error was introduced in Exchange 2013.
	ErrorRecipientNotFound ResponseCode = "ErrorRecipientNotFound"

	// ErrorRecognizerNotInstalled is intended for internal use only.
	// This error was introduced in Exchange 2013.
	ErrorRecognizerNotInstalled ResponseCode = "ErrorRecognizerNotInstalled"

	// ErrorSpeechGrammarError is intended for internal use only.
	// This error was introduced in Exchange 2013.
	ErrorSpeechGrammarError ResponseCode = "ErrorSpeechGrammarError"

	// ErrorInvalidManagementRoleHeader is returned if the ManagementRole
	// header in the SOAP header is incorrect.
	// This error was introduced in Exchange 2013.
	ErrorInvalidManagementRoleHeader ResponseCode = "ErrorInvalidManagementRoleHeader"

	// ErrorLocationServicesDisabled is intended for internal use only.
	// This error was introduced in Exchange 2013.
	ErrorLocationServicesDisabled ResponseCode = "ErrorLocationServicesDisabled"

	// ErrorLocationServicesRequestTimedOut is intended for internal use only.
	// This error was introduced in Exchange 2013.
	ErrorLocationServicesRequestTimedOut ResponseCode = "ErrorLocationServicesRequestTimedOut"

	// ErrorLocationServicesRequestFailed is intended for internal use only.
	// This error was introduced in Exchange 2013.
	ErrorLocationServicesRequestFailed ResponseCode = "ErrorLocationServicesRequestFailed"

	// ErrorLocationServicesInvalidRequest is intended for internal use only.
	// This error was introduced in Exchange 2013.
	ErrorLocationServicesInvalidRequest ResponseCode = "ErrorLocationServicesInvalidRequest"

	// ErrorWeatherServiceDisabled is intended for internal use only.
	ErrorWeatherServiceDisabled ResponseCode = "ErrorWeatherServiceDisabled"

	// ErrorMailboxScopeNotAllowedWithoutQueryString is returned when a scoped
	// search attempt is performed without using a QueryString (String) element
	// for a content indexing search. This is applicable to the SearchMailboxes
	// and FindConversation operations.
	// This error was introduced in Exchange 2013.
	ErrorMailboxScopeNotAllowedWithoutQueryString ResponseCode = "ErrorMailboxScopeNotAllowedWithoutQueryString"

	// ErrorArchiveMailboxSearchFailed is returned when an archive mailbox
	// search is unsuccessful.
	// This error was introduced in Exchange 2013.
	ErrorArchiveMailboxSearchFailed ResponseCode = "ErrorArchiveMailboxSearchFailed"

	// ErrorArchiveMailboxServiceDiscoveryFailed is returned when the URL of an
	// archive mailbox is not discoverable.
	// This error was introduced in Exchange 2013. (duplicate)
	// ErrorArchiveMailboxServiceDiscoveryFailed ResponseCode = "ErrorArchiveMailboxServiceDiscoveryFailed"

	// ErrorGetRemoteArchiveFolderFailed occurs when the operation to get the
	// remote archive mailbox folder failed.
	ErrorGetRemoteArchiveFolderFailed ResponseCode = "ErrorGetRemoteArchiveFolderFailed"

	// ErrorFindRemoteArchiveFolderFailed occurs when the operation to find the
	// remote archive mailbox folder failed.
	ErrorFindRemoteArchiveFolderFailed ResponseCode = "ErrorFindRemoteArchiveFolderFailed"

	// ErrorGetRemoteArchiveItemFailed occurs when the operation to get the
	// remote archive mailbox item failed.
	ErrorGetRemoteArchiveItemFailed ResponseCode = "ErrorGetRemoteArchiveItemFailed"

	// ErrorExportRemoteArchiveItemsFailed occurs when the operation to export
	// remote archive mailbox items failed.
	ErrorExportRemoteArchiveItemsFailed ResponseCode = "ErrorExportRemoteArchiveItemsFailed"

	// ErrorInvalidPhotoSize is returned if an invalid photo size is requested
	// from the server. This error was introduced in Exchange 2013.
	ErrorInvalidPhotoSize ResponseCode = "ErrorInvalidPhotoSize"

	// ErrorSearchQueryHasTooManyKeywords is returned when an unexpected photo
	// size is requested in a GetUserPhoto operation request.
	// This error was introduced in Exchange 2013.
	ErrorSearchQueryHasTooManyKeywords ResponseCode = "ErrorSearchQueryHasTooManyKeywords"

	// ErrorSearchTooManyMailboxes is returned when a SearchMailboxes operation
	// request contains too many mailboxes to search.
	// This error was introduced in Exchange 2013.
	ErrorSearchTooManyMailboxes ResponseCode = "ErrorSearchTooManyMailboxes"

	// ErrorInvalidRetentionTagNone indicates that no retention tags were found
	// for this user. This error was introduced in Exchange 2013.
	ErrorInvalidRetentionTagNone ResponseCode = "ErrorInvalidRetentionTagNone"

	// ErrorDiscoverySearchesDisabled is returned when discovery searches are
	// disabled on a tenant or server.
	// This error was introduced in Exchange 2013.
	ErrorDiscoverySearchesDisabled ResponseCode = "ErrorDiscoverySearchesDisabled"

	// ErrorCalendarSeekToConditionNotSupported occurs when attempting to
	// invoke the FindItem operation with a SeekToConditionPageItemView for
	// fetching calendar items, which is not supported.
	ErrorCalendarSeekToConditionNotSupported ResponseCode = "ErrorCalendarSeekToConditionNotSupported"

	// ErrorCalendarIsGroupMailboxForAccept is intended for internal use only.
	// ErrorCalendarIsGroupMailboxForAccept ResponseCode = "ErrorCalendarIsGroupMailboxForAccept"

	// ErrorCalendarIsGroupMailboxForDecline is intended for internal use only.
	// ErrorCalendarIsGroupMailboxForDecline ResponseCode = "ErrorCalendarIsGroupMailboxForDecline"

	// ErrorCalendarIsGroupMailboxForTentative is intended for internal use
	// only.
	// ErrorCalendarIsGroupMailboxForTentative ResponseCode = "ErrorCalendarIsGroupMailboxForTentative"

	// ErrorCalendarIsGroupMailboxForSuppressReadReceipt is intended for
	// internal use only.
	// ErrorCalendarIsGroupMailboxForSuppressReadReceipt ResponseCode = "ErrorCalendarIsGroupMailboxForSuppressReadReceipt"

	// ErrorOrganizationAccessBlocked is marked for removal.
	// ErrorOrganizationAccessBlocked ResponseCode = "ErrorOrganizationAccessBlocked"

	// ErrorInvalidLicense indicates the user doesn't have a valid license.
	ErrorInvalidLicense ResponseCode = "ErrorInvalidLicense"

	// ErrorMessagePerFolderCountReceiveQuotaExceeded indicates the message per
	// folder receive quota has been exceeded.
	ErrorMessagePerFolderCountReceiveQuotaExceeded ResponseCode = "ErrorMessagePerFolderCountReceiveQuotaExceeded"
)
