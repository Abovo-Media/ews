package ewsxml

import (
	"encoding/xml"
)

type FieldUri string

const (
	// FieldUri_Folder_FolderId identifies the FolderId property.
	FieldUri_Folder_FolderId FieldUri = "folder:FolderId"
	// FieldUri_Folder_ParentFolderId identifies the ParentFolderId property.
	FieldUri_Folder_ParentFolderId FieldUri = "folder:ParentFolderId"
	// FieldUri_Folder_DisplayName identifies the DisplayName property.
	FieldUri_Folder_DisplayName FieldUri = "folder:DisplayName"
	// FieldUri_Folder_UnreadCount identifies the UnreadCount property.
	FieldUri_Folder_UnreadCount FieldUri = "folder:UnreadCount"
	// FieldUri_Folder_TotalCount identifies the TotalCount property.
	FieldUri_Folder_TotalCount FieldUri = "folder:TotalCount"
	// FieldUri_Folder_ChildFolderCount identifies the ChildFolderCount property.
	FieldUri_Folder_ChildFolderCount FieldUri = "folder:ChildFolderCount"
	// FieldUri_Folder_FolderClass identifies the FolderClass property.
	FieldUri_Folder_FolderClass FieldUri = "folder:FolderClass"
	// FieldUri_Folder_SearchParameters identifies the SearchParameters property.
	FieldUri_Folder_SearchParameters FieldUri = "folder:SearchParameters"
	// FieldUri_Folder_ManagedFolderInformation identifies the ManagedFolderInformation property.
	FieldUri_Folder_ManagedFolderInformation FieldUri = "folder:ManagedFolderInformation"
	// FieldUri_Folder_PermissionSet identifies the PermissionSet property.
	FieldUri_Folder_PermissionSet FieldUri = "folder:PermissionSet"
	// FieldUri_Folder_EffectiveRights identifies the EffectiveRights property.
	FieldUri_Folder_EffectiveRights FieldUri = "folder:EffectiveRights"
	// FieldUri_Folder_SharingEffectiveRights identifies the SharingEffectiveRights property.
	FieldUri_Folder_SharingEffectiveRights FieldUri = "folder:SharingEffectiveRights"
	// FieldUri_Item_ItemId identifies the ItemId property.
	FieldUri_Item_ItemId FieldUri = "item:ItemId"
	// FieldUri_Item_ParentFolderId identifies the ParentFolderId property.
	FieldUri_Item_ParentFolderId FieldUri = "item:ParentFolderId"
	// FieldUri_Item_ItemClass identifies the ItemClass property.
	FieldUri_Item_ItemClass FieldUri = "item:ItemClass"
	// FieldUri_Item_MimeContent identifies the MimeContent property.
	FieldUri_Item_MimeContent FieldUri = "item:MimeContent"
	// FieldUri_Item_Attachments identifies the Attachments property.
	FieldUri_Item_Attachments FieldUri = "item:Attachments"
	// FieldUri_Item_Subject identifies the Subject property.
	FieldUri_Item_Subject FieldUri = "item:Subject"
	// FieldUri_Item_DateTimeReceived identifies the DateTimeReceived property.
	FieldUri_Item_DateTimeReceived FieldUri = "item:DateTimeReceived"
	// FieldUri_Item_Size identifies the Size property.
	FieldUri_Item_Size FieldUri = "item:Size"
	// FieldUri_Item_Categories identifies the Categories property.
	FieldUri_Item_Categories FieldUri = "item:Categories"
	// FieldUri_Item_HasAttachments identifies the HasAttachments property.
	FieldUri_Item_HasAttachments FieldUri = "item:HasAttachments"
	// FieldUri_Item_Importance identifies the Importance property.
	FieldUri_Item_Importance FieldUri = "item:Importance"
	// FieldUri_Item_InReplyTo identifies the InReplyTo property.
	FieldUri_Item_InReplyTo FieldUri = "item:InReplyTo"
	// FieldUri_Item_InternetMessageHeaders identifies the InternetMessageHeaders property.
	FieldUri_Item_InternetMessageHeaders FieldUri = "item:InternetMessageHeaders"
	// FieldUri_Item_IsAssociated identifies the IsAssociated property.
	FieldUri_Item_IsAssociated FieldUri = "item:IsAssociated"
	// FieldUri_Item_IsDraft identifies the IsDraft property.
	FieldUri_Item_IsDraft FieldUri = "item:IsDraft"
	// FieldUri_Item_IsFromMe identifies the IsFromMe property.
	FieldUri_Item_IsFromMe FieldUri = "item:IsFromMe"
	// FieldUri_Item_IsResend identifies the IsResend property.
	FieldUri_Item_IsResend FieldUri = "item:IsResend"
	// FieldUri_Item_IsSubmitted identifies the IsSubmitted property.
	FieldUri_Item_IsSubmitted FieldUri = "item:IsSubmitted"
	// FieldUri_Item_IsUnmodified identifies the IsUnmodified property.
	FieldUri_Item_IsUnmodified FieldUri = "item:IsUnmodified"
	// FieldUri_Item_DateTimeSent identifies the DateTimeSent property.
	FieldUri_Item_DateTimeSent FieldUri = "item:DateTimeSent"
	// FieldUri_Item_DateTimeCreated identifies the DateTimeCreated property.
	FieldUri_Item_DateTimeCreated FieldUri = "item:DateTimeCreated"
	// FieldUri_Item_Body identifies the Body property.
	FieldUri_Item_Body FieldUri = "item:Body"
	// FieldUri_Item_ResponseObjects identifies the ResponseObjects property.
	FieldUri_Item_ResponseObjects FieldUri = "item:ResponseObjects"
	// FieldUri_Item_Sensitivity identifies the Sensitivity property.
	FieldUri_Item_Sensitivity FieldUri = "item:Sensitivity"
	// FieldUri_Item_ReminderDueBy identifies the ReminderDueBy property.
	FieldUri_Item_ReminderDueBy FieldUri = "item:ReminderDueBy"
	// FieldUri_Item_ReminderIsSet identifies the ReminderIsSet property.
	FieldUri_Item_ReminderIsSet FieldUri = "item:ReminderIsSet"
	// FieldUri_Item_ReminderNextTime identifies the ReminderNextTime property.
	FieldUri_Item_ReminderNextTime FieldUri = "item:ReminderNextTime"
	// FieldUri_Item_ReminderMinutesBeforeStart identifies the ReminderMinutesBeforeStart property.
	FieldUri_Item_ReminderMinutesBeforeStart FieldUri = "item:ReminderMinutesBeforeStart"
	// FieldUri_Item_DisplayTo identifies the DisplayTo property.
	FieldUri_Item_DisplayTo FieldUri = "item:DisplayTo"
	// FieldUri_Item_DisplayCc identifies the DisplayCc property.
	FieldUri_Item_DisplayCc FieldUri = "item:DisplayCc"
	// FieldUri_Item_Culture identifies the Culture property.
	FieldUri_Item_Culture FieldUri = "item:Culture"
	// FieldUri_Item_EffectiveRights identifies the EffectiveRights property.
	FieldUri_Item_EffectiveRights FieldUri = "item:EffectiveRights"
	// FieldUri_Item_LastModifiedName identifies the LastModifiedName property.
	FieldUri_Item_LastModifiedName FieldUri = "item:LastModifiedName"
	// FieldUri_Item_LastModifiedTime identifies the LastModifiedTime property.
	FieldUri_Item_LastModifiedTime FieldUri = "item:LastModifiedTime"
	// FieldUri_Item_ConversationId identifies the ConversationId property.
	FieldUri_Item_ConversationId FieldUri = "item:ConversationId"
	// FieldUri_Item_UniqueBody identifies the UniqueBody property.
	FieldUri_Item_UniqueBody FieldUri = "item:UniqueBody"
	// FieldUri_Item_Flag identifies the Flag property.
	FieldUri_Item_Flag FieldUri = "item:Flag"
	// FieldUri_Item_StoreEntryId identifies the StoreEntryId property.
	FieldUri_Item_StoreEntryId FieldUri = "item:StoreEntryId"
	// FieldUri_Item_InstanceKey identifies the InstanceKey property.
	FieldUri_Item_InstanceKey FieldUri = "item:InstanceKey"
	// FieldUri_Item_NormalizedBody identifies the NormalizedBody property.
	FieldUri_Item_NormalizedBody FieldUri = "item:NormalizedBody"
	// FieldUri_Item_EntityExtractionResult identifies the EntityExtractionResult property.
	FieldUri_Item_EntityExtractionResult FieldUri = "item:EntityExtractionResult"
	// FieldUri_itemPolicyTag identifies the PolicyTag property.
	FieldUri_itemPolicyTag FieldUri = "itemPolicyTag"
	// FieldUri_Item_ArchiveTag identifies the ArchiveTag property.
	FieldUri_Item_ArchiveTag FieldUri = "item:ArchiveTag"
	// FieldUri_Item_RetentionDate identifies the RetentionDate property.
	FieldUri_Item_RetentionDate FieldUri = "item:RetentionDate"
	// FieldUri_Item_Preview identifies the Preview property.
	FieldUri_Item_Preview FieldUri = "item:Preview"
	// FieldUri_Item_NextPredictedAction identifies the NextPredictedAction property.
	FieldUri_Item_NextPredictedAction FieldUri = "item:NextPredictedAction"
	// FieldUri_Item_GroupingAction identifies the GroupingAction property.
	FieldUri_Item_GroupingAction FieldUri = "item:GroupingAction"
	// FieldUri_Item_PredictedActionReasons identifies the PredictedActionReasons property
	FieldUri_Item_PredictedActionReasons FieldUri = "item:PredictedActionReasons"
	// FieldUri_Item_RightsManagementLicenseData identifies the RightsManagementLicenseData property.
	FieldUri_Item_RightsManagementLicenseData FieldUri = "item:RightsManagementLicenseData"
	// FieldUri_Item_BlockStatus identifies the BlockStatus property.
	FieldUri_Item_BlockStatus FieldUri = "item:BlockStatus"
	// FieldUri_Item_HasBlockedImages identifies the HasBlockedImages property.
	FieldUri_Item_HasBlockedImages FieldUri = "item:HasBlockedImages"
	// FieldUri_Item_WebClientReadFormQueryString identifies the WebClientReadFormQueryString property.
	FieldUri_Item_WebClientReadFormQueryString FieldUri = "item:WebClientReadFormQueryString"
	// FieldUri_Item_WebClientEditFormQueryString identifies the WebClientEditFormQueryString property.
	FieldUri_Item_WebClientEditFormQueryString FieldUri = "item:WebClientEditFormQueryString"
	// FieldUri_Item_TextBody identifies the TextBody property.
	FieldUri_Item_TextBody FieldUri = "item:TextBody"
	// FieldUri_Item_IconIndex identifies the IconIndex property.
	FieldUri_Item_IconIndex FieldUri = "item:IconIndex"
	// FieldUri_Item_MimeContentUTF8 identifies the MimeContentUTF8 property.
	FieldUri_Item_MimeContentUTF8 FieldUri = "item:MimeContentUTF8"
	// FieldUri_Message_ConversationIndex identifies the ConversationIndex property.
	FieldUri_Message_ConversationIndex FieldUri = "message:ConversationIndex"
	// FieldUri_Message_ConversationTopic identifies the ConversationTopic property.
	FieldUri_Message_ConversationTopic FieldUri = "message:ConversationTopic"
	// FieldUri_Message_InternetMessageId identifies the InternetMessageId property.
	FieldUri_Message_InternetMessageId FieldUri = "message:InternetMessageId"
	// FieldUri_Message_IsRead identifies the IsRead property.
	FieldUri_Message_IsRead FieldUri = "message:IsRead"
	// FieldUri_Message_IsResponseRequested identifies the IsResponseRequested property.
	FieldUri_Message_IsResponseRequested FieldUri = "message:IsResponseRequested"
	// FieldUri_Message_IsReadReceiptRequested identifies the IsReadReceiptRequested property.
	FieldUri_Message_IsReadReceiptRequested FieldUri = "message:IsReadReceiptRequested"
	// FieldUri_Message_IsDeliveryReceiptRequested identifies the IsDeliveryReceiptRequested property.
	FieldUri_Message_IsDeliveryReceiptRequested FieldUri = "message:IsDeliveryReceiptRequested"
	// FieldUri_Message_ReceivedBy identifies the ReceivedBy property.
	FieldUri_Message_ReceivedBy FieldUri = "message:ReceivedBy"
	// FieldUri_Message_ReceivedRepresenting identifies the ReceivedRepresenting property.
	FieldUri_Message_ReceivedRepresenting FieldUri = "message:ReceivedRepresenting"
	// FieldUri_Message_References identifies the References property.
	FieldUri_Message_References FieldUri = "message:References"
	// FieldUri_Message_ReplyTo identifies the ReplyTo property.
	FieldUri_Message_ReplyTo FieldUri = "message:ReplyTo"
	// FieldUri_Message_From identifies the From property.
	FieldUri_Message_From FieldUri = "message:From"
	// FieldUri_Message_Sender identifies the Sender property.
	FieldUri_Message_Sender FieldUri = "message:Sender"
	// FieldUri_Message_ToRecipients identifies the ToRecipients property.
	FieldUri_Message_ToRecipients FieldUri = "message:ToRecipients"
	// FieldUri_Message_CcRecipients identifies the CcRecipients property.
	FieldUri_Message_CcRecipients FieldUri = "message:CcRecipients"
	// FieldUri_Message_BccRecipients identifies the BccRecipients property.
	FieldUri_Message_BccRecipients FieldUri = "message:BccRecipients"
	// FieldUri_Message_ApprovalRequestData identifies the ApprovalRequestData property.
	FieldUri_Message_ApprovalRequestData FieldUri = "message:ApprovalRequestData"
	// FieldUri_Message_VotingInformation identifies the VotingInformation property.
	FieldUri_Message_VotingInformation FieldUri = "message:VotingInformation"
	// FieldUri_Message_ReminderMessageData identifies the ReminderMessageData property.
	FieldUri_Message_ReminderMessageData FieldUri = "message:ReminderMessageData"
	// FieldUri_Meeting_AssociatedCalendarItemId identifies the AssociatedCalendarItemId property.
	FieldUri_Meeting_AssociatedCalendarItemId FieldUri = "meeting:AssociatedCalendarItemId"
	// FieldUri_Meeting_IsDelegated identifies the IsDelegated property.
	FieldUri_Meeting_IsDelegated FieldUri = "meeting:IsDelegated"
	// FieldUri_Meeting_IsOutOfDate identifies the IsOutOfDate property.
	FieldUri_Meeting_IsOutOfDate FieldUri = "meeting:IsOutOfDate"
	// FieldUri_Meeting_HasBeenProcessed identifies the HasBeenProcessed property.
	FieldUri_Meeting_HasBeenProcessed FieldUri = "meeting:HasBeenProcessed"
	// FieldUri_Meeting_ResponseType identifies the ResponseType property.
	FieldUri_Meeting_ResponseType FieldUri = "meeting:ResponseType"
	// FieldUri_Meeting_ProposedStart identifies the ProposedStart property.
	FieldUri_Meeting_ProposedStart FieldUri = "meeting:ProposedStart"
	// FieldUri_Meeting_ProposedEnd identifies the ProposedEnd property.
	FieldUri_Meeting_ProposedEnd FieldUri = "meeting:ProposedEnd"
	// FieldUri_MeetingRequest_MeetingRequestType identifies the MeetingRequestType property.
	FieldUri_MeetingRequest_MeetingRequestType FieldUri = "meetingRequest:MeetingRequestType"
	// FieldUri_MeetingRequest_IntendedFreeBusyStatus identifies the IntendedFreeBusyStatus property.
	FieldUri_MeetingRequest_IntendedFreeBusyStatus FieldUri = "meetingRequest:IntendedFreeBusyStatus"
	// FieldUri_MeetingRequest_ChangeHighlights identifies the ChangeHighlights property.
	FieldUri_MeetingRequest_ChangeHighlights FieldUri = "meetingRequest:ChangeHighlights"
	// FieldUri_Calendar_Start identifies the Start property.
	FieldUri_Calendar_Start FieldUri = "calendar:Start"
	// FieldUri_Calendar_End identifies the End property.
	FieldUri_Calendar_End FieldUri = "calendar:End"
	// FieldUri_Calendar_OriginalStart identifies the OriginalStart property.
	FieldUri_Calendar_OriginalStart FieldUri = "calendar:OriginalStart"
	// FieldUri_Calendar_StartWallClock identifies the StartWallClock property.
	FieldUri_Calendar_StartWallClock FieldUri = "calendar:StartWallClock"
	// FieldUri_Calendar_EndWallClock identifies the EndWallClock property.
	FieldUri_Calendar_EndWallClock FieldUri = "calendar:EndWallClock"
	// FieldUri_Calendar_StartTimeZoneId identifies the StartTimeZoneId property.
	FieldUri_Calendar_StartTimeZoneId FieldUri = "calendar:StartTimeZoneId"
	// FieldUri_Calendar_EndTimeZoneId identifies the EndTimeZoneId property.
	FieldUri_Calendar_EndTimeZoneId FieldUri = "calendar:EndTimeZoneId"
	// FieldUri_Calendar_IsAllDayEvent identifies the IsAllDayEvent property.
	FieldUri_Calendar_IsAllDayEvent FieldUri = "calendar:IsAllDayEvent"
	// FieldUri_Calendar_LegacyFreeBusyStatus identifies the LegacyFreeBusyStatus property.
	FieldUri_Calendar_LegacyFreeBusyStatus FieldUri = "calendar:LegacyFreeBusyStatus"
	// FieldUri_Calendar_Location identifies the Location property.
	FieldUri_Calendar_Location FieldUri = "calendar:Location"
	// FieldUri_Calendar_When identifies the When property.
	FieldUri_Calendar_When FieldUri = "calendar:When"
	// FieldUri_Calendar_IsMeeting identifies the IsMeeting property.
	FieldUri_Calendar_IsMeeting FieldUri = "calendar:IsMeeting"
	// FieldUri_Calendar_IsCancelled identifies the IsCancelled property.
	FieldUri_Calendar_IsCancelled FieldUri = "calendar:IsCancelled"
	// FieldUri_Calendar_IsRecurring identifies the IsRecurring property.
	FieldUri_Calendar_IsRecurring FieldUri = "calendar:IsRecurring"
	// FieldUri_Calendar_MeetingRequestWasSent identifies the MeetingRequestWasSent property.
	FieldUri_Calendar_MeetingRequestWasSent FieldUri = "calendar:MeetingRequestWasSent"
	// FieldUri_Calendar_IsResponseRequested identifies the IsResponseRequested property.
	FieldUri_Calendar_IsResponseRequested FieldUri = "calendar:IsResponseRequested"
	// FieldUri_Calendar_CalendarItemType identifies the CalendarItemType property.
	FieldUri_Calendar_CalendarItemType FieldUri = "calendar:CalendarItemType"
	// FieldUri_Calendar_MyResponseType identifies the MyResponseType property.
	FieldUri_Calendar_MyResponseType FieldUri = "calendar:MyResponseType"
	// FieldUri_Calendar_Organizer identifies the Organizer property.
	FieldUri_Calendar_Organizer FieldUri = "calendar:Organizer"
	// FieldUri_Calendar_RequiredAttendees identifies the RequiredAttendees property.
	FieldUri_Calendar_RequiredAttendees FieldUri = "calendar:RequiredAttendees"
	// FieldUri_Calendar_OptionalAttendees identifies the OptionalAttendees property.
	FieldUri_Calendar_OptionalAttendees FieldUri = "calendar:OptionalAttendees"
	// FieldUri_Calendar_Resources identifies the Resources property.
	FieldUri_Calendar_Resources FieldUri = "calendar:Resources"
	// FieldUri_Calendar_ConflictingMeetingCount identifies the ConflictingMeetingCount property.
	FieldUri_Calendar_ConflictingMeetingCount FieldUri = "calendar:ConflictingMeetingCount"
	// FieldUri_Calendar_AdjacentMeetingCount identifies the AdjacentMeetingCount property.
	FieldUri_Calendar_AdjacentMeetingCount FieldUri = "calendar:AdjacentMeetingCount"
	// FieldUri_Calendar_ConflictingMeetings identifies the ConflictingMeetings property.
	FieldUri_Calendar_ConflictingMeetings FieldUri = "calendar:ConflictingMeetings"
	// FieldUri_Calendar_AdjacentMeetings identifies the AdjacentMeetings property.
	FieldUri_Calendar_AdjacentMeetings FieldUri = "calendar:AdjacentMeetings"
	// FieldUri_Calendar_Duration identifies the Duration property.
	FieldUri_Calendar_Duration FieldUri = "calendar:Duration"
	// FieldUri_Calendar_TimeZone identifies the TimeZone property.
	FieldUri_Calendar_TimeZone FieldUri = "calendar:TimeZone"
	// FieldUri_Calendar_AppointmentReplyTime identifies the AppointmentReplyTime property.
	FieldUri_Calendar_AppointmentReplyTime FieldUri = "calendar:AppointmentReplyTime"
	// FieldUri_Calendar_AppointmentSequenceNumber identifies the AppointmentSequenceNumber property.
	FieldUri_Calendar_AppointmentSequenceNumber FieldUri = "calendar:AppointmentSequenceNumber"
	// FieldUri_Calendar_AppointmentState identifies the AppointmentState property.
	FieldUri_Calendar_AppointmentState FieldUri = "calendar:AppointmentState"
	// FieldUri_Calendar_Recurrence identifies the Recurrence property.
	FieldUri_Calendar_Recurrence FieldUri = "calendar:Recurrence"
	// FieldUri_Calendar_FirstOccurrence identifies the FirstOccurrence property.
	FieldUri_Calendar_FirstOccurrence FieldUri = "calendar:FirstOccurrence"
	// FieldUri_Calendar_LastOccurrence identifies the LastOccurrence property.
	FieldUri_Calendar_LastOccurrence FieldUri = "calendar:LastOccurrence"
	// FieldUri_Calendar_ModifiedOccurrences identifies the ModifiedOccurrences property.
	FieldUri_Calendar_ModifiedOccurrences FieldUri = "calendar:ModifiedOccurrences"
	// FieldUri_Calendar_DeletedOccurrences identifies the DeletedOccurrences property.
	FieldUri_Calendar_DeletedOccurrences FieldUri = "calendar:DeletedOccurrences"
	// FieldUri_Calendar_MeetingTimeZone identifies the MeetingTimeZone property.
	FieldUri_Calendar_MeetingTimeZone FieldUri = "calendar:MeetingTimeZone"
	// FieldUri_Calendar_ConferenceType identifies the ConferenceType property.
	FieldUri_Calendar_ConferenceType FieldUri = "calendar:ConferenceType"
	// FieldUri_Calendar_AllowNewTimeProposal identifies the AllowNewTimeProposal property.
	FieldUri_Calendar_AllowNewTimeProposal FieldUri = "calendar:AllowNewTimeProposal"
	// FieldUri_Calendar_IsOnlineMeeting identifies the IsOnlineMeeting property.
	FieldUri_Calendar_IsOnlineMeeting FieldUri = "calendar:IsOnlineMeeting"
	// FieldUri_Calendar_MeetingWorkspaceUrl identifies the MeetingWorkspaceUrl property.
	FieldUri_Calendar_MeetingWorkspaceUrl FieldUri = "calendar:MeetingWorkspaceUrl"
	// FieldUri_Calendar_NetShowUrl identifies the NetShowUrl property.
	FieldUri_Calendar_NetShowUrl FieldUri = "calendar:NetShowUrl"
	// FieldUri_Calendar_UID identifies the UID property.
	FieldUri_Calendar_UID FieldUri = "calendar:UID"
	// FieldUri_Calendar_RecurrenceId identifies the RecurrenceId property.
	FieldUri_Calendar_RecurrenceId FieldUri = "calendar:RecurrenceId"
	// FieldUri_Calendar_DateTimeStamp identifies the DateTimeStamp property.
	FieldUri_Calendar_DateTimeStamp FieldUri = "calendar:DateTimeStamp"
	// FieldUri_Calendar_StartTimeZone identifies the StartTimeZone property.
	FieldUri_Calendar_StartTimeZone FieldUri = "calendar:StartTimeZone"
	// FieldUri_Calendar_EndTimeZone identifies the EndTimeZone property.
	FieldUri_Calendar_EndTimeZone FieldUri = "calendar:EndTimeZone"
	// FieldUri_Calendar_JoinOnlineMeetingUrl identifies the JoinOnlineMeetingUrl property.
	FieldUri_Calendar_JoinOnlineMeetingUrl FieldUri = "calendar:JoinOnlineMeetingUrl"
	// FieldUri_Calendar_OnlineMeetingSettings identifies the OnlineMeetingSettings property.
	FieldUri_Calendar_OnlineMeetingSettings FieldUri = "calendar:OnlineMeetingSettings"
	// FieldUri_Calendar_IsOrganizer identifies the IsOrganizer property.
	FieldUri_Calendar_IsOrganizer FieldUri = "calendar:IsOrganizer"
	// FieldUri_Task_ActualWork identifies the ActualWork property.
	FieldUri_Task_ActualWork FieldUri = "task:ActualWork"
	// FieldUri_Task_AssignedTime identifies the AssignedTime property.
	FieldUri_Task_AssignedTime FieldUri = "task:AssignedTime"
	// FieldUri_Task_BillingInformation identifies the BillingInformation property.
	FieldUri_Task_BillingInformation FieldUri = "task:BillingInformation"
	// FieldUri_Task_ChangeCount identifies the ChangeCount property.
	FieldUri_Task_ChangeCount FieldUri = "task:ChangeCount"
	// FieldUri_Task_Companies identifies the Companies property.
	FieldUri_Task_Companies FieldUri = "task:Companies"
	// FieldUri_Task_CompleteDate identifies the CompleteDate property.
	FieldUri_Task_CompleteDate FieldUri = "task:CompleteDate"
	// FieldUri_Task_Contacts identifies the Contacts property.
	FieldUri_Task_Contacts FieldUri = "task:Contacts"
	// FieldUri_Task_DelegationState identifies the DelegationState property.
	FieldUri_Task_DelegationState FieldUri = "task:DelegationState"
	// FieldUri_Task_Delegator identifies the Delegator property.
	FieldUri_Task_Delegator FieldUri = "task:Delegator"
	// FieldUri_Task_DueDate identifies the DueDate property.
	FieldUri_Task_DueDate FieldUri = "task:DueDate"
	// FieldUri_Task_IsAssignmentEditable identifies the IsAssignmentEditable property.
	FieldUri_Task_IsAssignmentEditable FieldUri = "task:IsAssignmentEditable"
	// FieldUri_Task_IsComplete identifies the IsComplete property.
	FieldUri_Task_IsComplete FieldUri = "task:IsComplete"
	// FieldUri_Task_IsRecurring identifies the IsRecurring property.
	FieldUri_Task_IsRecurring FieldUri = "task:IsRecurring"
	// FieldUri_Task_IsTeamTask identifies the IsTeamTask property.
	FieldUri_Task_IsTeamTask FieldUri = "task:IsTeamTask"
	// FieldUri_Task_Mileage identifies the Mileage property.
	FieldUri_Task_Mileage FieldUri = "task:Mileage"
	// FieldUri_Task_Owner identifies the Owner property.
	FieldUri_Task_Owner FieldUri = "task:Owner"
	// FieldUri_Task_PercentComplete identifies the PercentComplete property.
	FieldUri_Task_PercentComplete FieldUri = "task:PercentComplete"
	// FieldUri_Task_Recurrence identifies the Recurrence property.
	FieldUri_Task_Recurrence FieldUri = "task:Recurrence"
	// FieldUri_Task_StartDate identifies the StartDate property.
	FieldUri_Task_StartDate FieldUri = "task:StartDate"
	// FieldUri_Task_Status identifies the Status property.
	FieldUri_Task_Status FieldUri = "task:Status"
	// FieldUri_Task_StatusDescription identifies the StatusDescription property.
	FieldUri_Task_StatusDescription FieldUri = "task:StatusDescription"
	// FieldUri_Task_TotalWork identifies the TotalWork property.
	FieldUri_Task_TotalWork FieldUri = "task:TotalWork"
	// FieldUri_Contacts_Alias identifies the Alias property. This property was introduced in Exchange Server 2010 Service Pack 2 (SP2).
	FieldUri_Contacts_Alias FieldUri = "contacts:Alias"
	// FieldUri_Contacts_AssistantName identifies the AssistantName property.
	FieldUri_Contacts_AssistantName FieldUri = "contacts:AssistantName"
	// FieldUri_Contacts_Birthday identifies the Birthday property.
	FieldUri_Contacts_Birthday FieldUri = "contacts:Birthday"
	// FieldUri_Contacts_BusinessHomePage identifies the BusinessHomePage property.
	FieldUri_Contacts_BusinessHomePage FieldUri = "contacts:BusinessHomePage"
	// FieldUri_Contacts_Children identifies the Children property.
	FieldUri_Contacts_Children FieldUri = "contacts:Children"
	// FieldUri_Contacts_Companies identifies the Companies property.
	FieldUri_Contacts_Companies FieldUri = "contacts:Companies"
	// FieldUri_Contacts_CompanyName identifies the CompanyName property.
	FieldUri_Contacts_CompanyName FieldUri = "contacts:CompanyName"
	// FieldUri_Contacts_CompleteName identifies the CompleteName property.
	FieldUri_Contacts_CompleteName FieldUri = "contacts:CompleteName"
	// FieldUri_Contacts_ContactSource identifies the ContactSource property.
	FieldUri_Contacts_ContactSource FieldUri = "contacts:ContactSource"
	// FieldUri_Contacts_Culture identifies the Culture property.
	FieldUri_Contacts_Culture FieldUri = "contacts:Culture"
	// FieldUri_Contacts_Department identifies the Department property.
	FieldUri_Contacts_Department FieldUri = "contacts:Department"
	// FieldUri_Contacts_DisplayName identifies the DisplayName property.
	FieldUri_Contacts_DisplayName FieldUri = "contacts:DisplayName"
	// FieldUri_Contacts_DirectoryId identifies the DirectoryId property. This property was introduced in Exchange 2010 SP2.
	FieldUri_Contacts_DirectoryId FieldUri = "contacts:DirectoryId"
	// FieldUri_Contacts_DirectReports identifies the DirectReports property. This property was introduced in Exchange 2010 SP2.
	FieldUri_Contacts_DirectReports FieldUri = "contacts:DirectReports"
	// FieldUri_Contacts_EmailAddresses identifies the EmailAddresses property.
	FieldUri_Contacts_EmailAddresses FieldUri = "contacts:EmailAddresses"
	// FieldUri_Contacts_FileAs identifies the FileAs property.
	FieldUri_Contacts_FileAs FieldUri = "contacts:FileAs"
	// FieldUri_Contacts_FileAsMapping identifies the FileAsMapping property.
	FieldUri_Contacts_FileAsMapping FieldUri = "contacts:FileAsMapping"
	// FieldUri_Contacts_Generation identifies the Generation property.
	FieldUri_Contacts_Generation FieldUri = "contacts:Generation"
	// FieldUri_Contacts_GivenName identifies the GivenName property.
	FieldUri_Contacts_GivenName FieldUri = "contacts:GivenName"
	// FieldUri_Contacts_ImAddresses identifies the ImAddresses property.
	FieldUri_Contacts_ImAddresses FieldUri = "contacts:ImAddresses"
	// FieldUri_Contacts_Initials identifies the Initials property.
	FieldUri_Contacts_Initials FieldUri = "contacts:Initials"
	// FieldUri_Contacts_JobTitle identifies the JobTitle property.
	FieldUri_Contacts_JobTitle FieldUri = "contacts:JobTitle"
	// FieldUri_Contacts_Manager identifies the Manager property.
	FieldUri_Contacts_Manager FieldUri = "contacts:Manager"
	// FieldUri_Contacts_ManagerMailbox identifies the ManagerMailbox property. This property was introduced in Exchange 2010 SP2.
	FieldUri_Contacts_ManagerMailbox FieldUri = "contacts:ManagerMailbox"
	// FieldUri_Contacts_MiddleName identifies the MiddleName property.
	FieldUri_Contacts_MiddleName FieldUri = "contacts:MiddleName"
	// FieldUri_Contacts_Mileage identifies the Mileage property.
	FieldUri_Contacts_Mileage FieldUri = "contacts:Mileage"
	// FieldUri_Contacts_MSExchangeCertificate identifies the MSExchangeCertificate property. This property was introduced in Exchange 2010 SP2.
	FieldUri_Contacts_MSExchangeCertificate FieldUri = "contacts:MSExchangeCertificate"
	// FieldUri_Contacts_Nickname identifies the Nickname property.
	FieldUri_Contacts_Nickname FieldUri = "contacts:Nickname"
	// FieldUri_Contacts_Notes identifies the Notes property. This property was introduced with in Exchange 2010 SP2.
	FieldUri_Contacts_Notes FieldUri = "contacts:Notes"
	// FieldUri_Contacts_OfficeLocation identifies the OfficeLocation property.
	FieldUri_Contacts_OfficeLocation FieldUri = "contacts:OfficeLocation"
	// FieldUri_Contacts_PhoneNumbers identifies the PhoneNumbers property.
	FieldUri_Contacts_PhoneNumbers FieldUri = "contacts:PhoneNumbers"
	// FieldUri_Contacts_PhoneticFullName identifies the PhoneticFullName property. This property was introduced in Exchange 2010 SP2.
	FieldUri_Contacts_PhoneticFullName FieldUri = "contacts:PhoneticFullName"
	// FieldUri_Contacts_PhoneticFirstName identifies the PhoneticFirstName property. This property was introduced in Exchange 2010 SP2.
	FieldUri_Contacts_PhoneticFirstName FieldUri = "contacts:PhoneticFirstName"
	// FieldUri_Contacts_PhoneticLastName identifies the PhoneticLastName property. This property was introduced in Exchange 2010 SP2.
	FieldUri_Contacts_PhoneticLastName FieldUri = "contacts:PhoneticLastName"
	// FieldUri_Contacts_Photo identifies the Photo property. This property was introduced in Exchange 2010 SP2.
	FieldUri_Contacts_Photo FieldUri = "contacts:Photo"
	// FieldUri_Contacts_PhysicalAddresses identifies the PhysicalAddresses property.
	FieldUri_Contacts_PhysicalAddresses FieldUri = "contacts:PhysicalAddresses"
	// FieldUri_Contacts_PostalAddressIndex identifies the PostalAddressIndex property.
	FieldUri_Contacts_PostalAddressIndex FieldUri = "contacts:PostalAddressIndex"
	// FieldUri_Contacts_Profession identifies the Profession property.
	FieldUri_Contacts_Profession FieldUri = "contacts:Profession"
	// FieldUri_Contacts_SpouseName identifies the SpouseName property.
	FieldUri_Contacts_SpouseName FieldUri = "contacts:SpouseName"
	// FieldUri_Contacts_Surname identifies the Surname property.
	FieldUri_Contacts_Surname FieldUri = "contacts:Surname"
	// FieldUri_Contacts_WeddingAnniversary identifies the WeddingAnniversary property.
	FieldUri_Contacts_WeddingAnniversary FieldUri = "contacts:WeddingAnniversary"
	// FieldUri_Contacts_UserSMIMECertificate identifies the UserSMIMECertificate property. This property was introduced in Exchange 2010 SP2.
	FieldUri_Contacts_UserSMIMECertificate FieldUri = "contacts:UserSMIMECertificate"
	// FieldUri_Contacts_HasPicture identifies the HasPicture property.
	FieldUri_Contacts_HasPicture FieldUri = "contacts:HasPicture"
	// FieldUri_DistributionList_Members identifies the Members property.
	FieldUri_DistributionList_Members FieldUri = "distributionlist:Members"
	// FieldUri_PostItem_PostedTime identifies the PostedTime property.
	FieldUri_PostItem_PostedTime FieldUri = "postitem:PostedTime"
	// FieldUri_Conversation_ConversationId identifies the ConversationId property.
	FieldUri_Conversation_ConversationId FieldUri = "conversation:ConversationId"
	// FieldUri_Conversation_ConversationTopic identifies the ConversationTopic property.
	FieldUri_Conversation_ConversationTopic FieldUri = "conversation:ConversationTopic"
	// FieldUri_Conversation_UniqueRecipients identifies the UniqueRecipients property.
	FieldUri_Conversation_UniqueRecipients FieldUri = "conversation:UniqueRecipients"
	// FieldUri_Conversation_GlobalUniqueRecipients identifies the GlobalUniqueRecipients property.
	FieldUri_Conversation_GlobalUniqueRecipients FieldUri = "conversation:GlobalUniqueRecipients"
	// FieldUri_Conversation_UniqueUnreadSenders identifies the UniqueUnreadSenders property.
	FieldUri_Conversation_UniqueUnreadSenders FieldUri = "conversation:UniqueUnreadSenders"
	// FieldUri_Conversation_GlobalUniqueUnreadSenders identifies the GlobalUniqueUnreadSenders property.
	FieldUri_Conversation_GlobalUniqueUnreadSenders FieldUri = "conversation:GlobalUniqueUnreadSenders"
	// FieldUri_Conversation_UniqueSenders identifies the UniqueSenders property.
	FieldUri_Conversation_UniqueSenders FieldUri = "conversation:UniqueSenders"
	// FieldUri_Conversation_GlobalUniqueSenders identifies the GlobalUniqueSenders property.
	FieldUri_Conversation_GlobalUniqueSenders FieldUri = "conversation:GlobalUniqueSenders"
	// FieldUri_Conversation_LastDeliveryTime identifies the LastDeliveryTime property.
	FieldUri_Conversation_LastDeliveryTime FieldUri = "conversation:LastDeliveryTime"
	// FieldUri_Conversation_GlobalLastDeliveryTime identifies the GlobalLastDeliveryTime property.
	FieldUri_Conversation_GlobalLastDeliveryTime FieldUri = "conversation:GlobalLastDeliveryTime"
	// FieldUri_Conversation_Categories identifies the Categories property.
	FieldUri_Conversation_Categories FieldUri = "conversation:Categories"
	// FieldUri_Conversation_GlobalCategories identifies the GlobalCategories property.
	FieldUri_Conversation_GlobalCategories FieldUri = "conversation:GlobalCategories"
	// FieldUri_Conversation_FlagStatus identifies the FlagStatus property.
	FieldUri_Conversation_FlagStatus FieldUri = "conversation:FlagStatus"
	// FieldUri_Conversation_GlobalFlagStatus identifies the GlobalFlagStatus property.
	FieldUri_Conversation_GlobalFlagStatus FieldUri = "conversation:GlobalFlagStatus"
	// FieldUri_Conversation_HasAttachments identifies the HasAttachments property.
	FieldUri_Conversation_HasAttachments FieldUri = "conversation:HasAttachments"
	// FieldUri_Conversation_GlobalHasAttachments identifies the GlobalHasAttachments property.
	FieldUri_Conversation_GlobalHasAttachments FieldUri = "conversation:GlobalHasAttachments"
	// FieldUri_Conversation_HasIrm identifies the HasIrm property.
	FieldUri_Conversation_HasIrm FieldUri = "conversation:HasIrm"
	// FieldUri_Conversation_GlobalHasIrm identifies the GlobalHasIrm property.
	FieldUri_Conversation_GlobalHasIrm FieldUri = "conversation:GlobalHasIrm"
	// FieldUri_Conversation_MessageCount identifies the MessageCount property.
	FieldUri_Conversation_MessageCount FieldUri = "conversation:MessageCount"
	// FieldUri_Conversation_GlobalMessageCount identifies the GlobalMessageCount property.
	FieldUri_Conversation_GlobalMessageCount FieldUri = "conversation:GlobalMessageCount"
	// FieldUri_Conversation_UnreadCount identifies the UnreadCount property.
	FieldUri_Conversation_UnreadCount FieldUri = "conversation:UnreadCount"
	// FieldUri_Conversation_GlobalUnreadCount identifies the GlobalUnreadCount property.
	FieldUri_Conversation_GlobalUnreadCount FieldUri = "conversation:GlobalUnreadCount"
	// FieldUri_Conversation_Size identifies the Size property.
	FieldUri_Conversation_Size FieldUri = "conversation:Size"
	// FieldUri_Conversation_GlobalSize identifies the GlobalSize property.
	FieldUri_Conversation_GlobalSize FieldUri = "conversation:GlobalSize"
	// FieldUri_Conversation_ItemClasses identifies the ItemClasses property.
	FieldUri_Conversation_ItemClasses FieldUri = "conversation:ItemClasses"
	// FieldUri_Conversation_GlobalItemClasses identifies the GlobalItemClasses property.
	FieldUri_Conversation_GlobalItemClasses FieldUri = "conversation:GlobalItemClasses"
	// FieldUri_Conversation_Importance identifies the Importance property.
	FieldUri_Conversation_Importance FieldUri = "conversation:Importance"
	// FieldUri_Conversation_GlobalImportance identifies the GlobalImportance property.
	FieldUri_Conversation_GlobalImportance FieldUri = "conversation:GlobalImportance"
	// FieldUri_Conversation_ItemIds identifies the ItemIds property.
	FieldUri_Conversation_ItemIds FieldUri = "conversation:ItemIds"
	// FieldUri_Conversation_GlobalItemIds identifies the GlobalItemIds property.
	FieldUri_Conversation_GlobalItemIds FieldUri = "conversation:GlobalItemIds"
	// FieldUri_Conversation_LastModifiedTime identifies the LastModifiedTime property.
	FieldUri_Conversation_LastModifiedTime FieldUri = "conversation:LastModifiedTime"
	// FieldUri_Conversation_InstanceKey identifies the InstanceKey property.
	FieldUri_Conversation_InstanceKey FieldUri = "conversation:InstanceKey"
	// FieldUri_Conversation_Preview identifies the Preview property.
	FieldUri_Conversation_Preview FieldUri = "conversation:Preview"
	// FieldUri_Conversation_GlobalParentFolderId identifies the GlobalParentFolderId property.
	FieldUri_Conversation_GlobalParentFolderId FieldUri = "conversation:GlobalParentFolderId"
	// FieldUri_Conversation_NextPredictedAction identifies the NextPredictedAction property.
	FieldUri_Conversation_NextPredictedAction FieldUri = "conversation:NextPredictedAction"
	// FieldUri_Conversation_GroupingAction identifies the GroupingAction property.
	FieldUri_Conversation_GroupingAction FieldUri = "conversation:GroupingAction"
	// FieldUri_Conversation_IconIndex identifies the IconIndex property.
	FieldUri_Conversation_IconIndex FieldUri = "conversation:IconIndex"
	// FieldUri_Conversation_GlobalIconIndex identifies the GlobalIconIndex property.
	FieldUri_Conversation_GlobalIconIndex FieldUri = "conversation:GlobalIconIndex"
	// FieldUri_Conversation_DraftItemIds identifies the DraftItemIds property.
	FieldUri_Conversation_DraftItemIds FieldUri = "conversation:DraftItemIds"
	// FieldUri_Persona_PersonaId identifies the PersonaId property.
	FieldUri_Persona_PersonaId FieldUri = "persona:PersonaId"
	// FieldUri_Persona_PersonaType identifies the PersonaType property.
	FieldUri_Persona_PersonaType FieldUri = "persona:PersonaType"
	// FieldUri_Persona_GivenName identifies the GivenName property.
	FieldUri_Persona_GivenName FieldUri = "persona:GivenName"
	// FieldUri_Persona_CompanyName identifies the CompanyName property.
	FieldUri_Persona_CompanyName FieldUri = "persona:CompanyName"
	// FieldUri_Persona_Surname identifies the Surname property.
	FieldUri_Persona_Surname FieldUri = "persona:Surname"
	// FieldUri_Persona_DisplayName identifies the DisplayName property.
	FieldUri_Persona_DisplayName FieldUri = "persona:DisplayName"
	// FieldUri_Persona_EmailAddress identifies the EmailAddress property.
	FieldUri_Persona_EmailAddress FieldUri = "persona:EmailAddress"
	// FieldUri_Persona_FileAs identifies the FileAs property.
	FieldUri_Persona_FileAs FieldUri = "persona:FileAs"
	// FieldUri_Persona_HomeCity identifies the HomeCity property.
	FieldUri_Persona_HomeCity FieldUri = "persona:HomeCity"
	// FieldUri_Persona_CreationTime identifies the CreationTime property.
	FieldUri_Persona_CreationTime FieldUri = "persona:CreationTime"
	// FieldUri_Persona_RelevanceScore identifies the RelevanceScore property.
	FieldUri_Persona_RelevanceScore FieldUri = "persona:RelevanceScore"
	// FieldUri_Persona_WorkCity identifies the WorkCity property.
	FieldUri_Persona_WorkCity FieldUri = "persona:WorkCity"
	// FieldUri_Persona_PersonaObjectStatus identifies the PersonaObjectStatus property.
	FieldUri_Persona_PersonaObjectStatus FieldUri = "persona:PersonaObjectStatus"
	// FieldUri_Persona_FileAsId identifies the FileAsId property.
	FieldUri_Persona_FileAsId FieldUri = "persona:FileAsId"
	// FieldUri_Persona_DisplayNamePrefix identifies the DisplayNamePrefix property.
	FieldUri_Persona_DisplayNamePrefix FieldUri = "persona:DisplayNamePrefix"
	// FieldUri_Persona_YomiCompanyName identifies the YomiCompanyName property.
	FieldUri_Persona_YomiCompanyName FieldUri = "persona:YomiCompanyName"
	// FieldUri_Persona_YomiFirstName identifies the YomiFirstName property.
	FieldUri_Persona_YomiFirstName FieldUri = "persona:YomiFirstName"
	// FieldUri_Persona_YomiLastName identifies the YomiLastName property.
	FieldUri_Persona_YomiLastName FieldUri = "persona:YomiLastName"
	// FieldUri_Persona_Title identifies the Title property.
	FieldUri_Persona_Title FieldUri = "persona:Title"
	// FieldUri_Persona_EmailAddresses identifies the EmailAddresses property.
	FieldUri_Persona_EmailAddresses FieldUri = "persona:EmailAddresses"
	// FieldUri_Persona_PhoneNumber identifies the PhoneNumber property.
	FieldUri_Persona_PhoneNumber FieldUri = "persona:PhoneNumber"
	// FieldUri_Persona_ImAddress identifies the ImAddress property.
	FieldUri_Persona_ImAddress FieldUri = "persona:ImAddress"
	// FieldUri_Persona_ImAddresses identifies the ImAddresses property.
	FieldUri_Persona_ImAddresses FieldUri = "persona:ImAddresses"
	// FieldUri_Persona_ImAddresses2 identifies the ImAddresses2 property.
	FieldUri_Persona_ImAddresses2 FieldUri = "persona:ImAddresses2"
	// FieldUri_Persona_ImAddresses3 identifies the ImAddresses3 property.
	FieldUri_Persona_ImAddresses3 FieldUri = "persona:ImAddresses3"
	// FieldUri_Persona_FolderIds identifies the FolderIds property.
	FieldUri_Persona_FolderIds FieldUri = "persona:FolderIds"
	// FieldUri_Persona_Attributions identifies the Attributions property.
	FieldUri_Persona_Attributions FieldUri = "persona:Attributions"
	// FieldUri_Persona_DisplayNames identifies the DisplayNames property.
	FieldUri_Persona_DisplayNames FieldUri = "persona:DisplayNames"
	// FieldUri_Persona_Initials identifies the Initials property.
	FieldUri_Persona_Initials FieldUri = "persona:Initials"
	// FieldUri_Persona_FileAses identifies the FileAses property.
	FieldUri_Persona_FileAses FieldUri = "persona:FileAses"
	// FieldUri_Persona_FileAsIds identifies the FileAsIds property.
	FieldUri_Persona_FileAsIds FieldUri = "persona:FileAsIds"
	// FieldUri_Persona_DisplayNamePrefixes identifies the DisplayNamePrefixes property.
	FieldUri_Persona_DisplayNamePrefixes FieldUri = "persona:DisplayNamePrefixes"
	// FieldUri_Persona_GivenNames identifies the GivenNames property.
	FieldUri_Persona_GivenNames FieldUri = "persona:GivenNames"
	// FieldUri_Persona_MiddleNames identifies the MiddleNames property.
	FieldUri_Persona_MiddleNames FieldUri = "persona:MiddleNames"
	// FieldUri_Persona_Surnames identifies the Surnames property.
	FieldUri_Persona_Surnames FieldUri = "persona:Surnames"
	// FieldUri_Persona_Generations identifies the Generations property.
	FieldUri_Persona_Generations FieldUri = "persona:Generations"
	// FieldUri_Persona_Nicknames identifies the Nicknames property.
	FieldUri_Persona_Nicknames FieldUri = "persona:Nicknames"
	// FieldUri_Persona_YomiCompanyNames identifies the YomiCompanyNames property.
	FieldUri_Persona_YomiCompanyNames FieldUri = "persona:YomiCompanyNames"
	// FieldUri_Persona_YomiFirstNames identifies the YomiFirstNames property.
	FieldUri_Persona_YomiFirstNames FieldUri = "persona:YomiFirstNames"
	// FieldUri_Persona_YomiLastNames identifies the YomiLastNames property.
	FieldUri_Persona_YomiLastNames FieldUri = "persona:YomiLastNames"
	// FieldUri_Persona_BusinessPhoneNumbers identifies the BusinessPhoneNumbers property.
	FieldUri_Persona_BusinessPhoneNumbers FieldUri = "persona:BusinessPhoneNumbers"
	// FieldUri_Persona_BusinessPhoneNumbers2 identifies the BusinessPhoneNumbers2 property.
	FieldUri_Persona_BusinessPhoneNumbers2 FieldUri = "persona:BusinessPhoneNumbers2"
	// FieldUri_Persona_HomePhones identifies the HomePhones property.
	FieldUri_Persona_HomePhones FieldUri = "persona:HomePhones"
	// FieldUri_Persona_HomePhones2 identifies the HomePhones2 property.
	FieldUri_Persona_HomePhones2 FieldUri = "persona:HomePhones2"
	// FieldUri_Persona_MobilePhones identifies the MobilePhones property.
	FieldUri_Persona_MobilePhones FieldUri = "persona:MobilePhones"
	// FieldUri_Persona_MobilePhones2 identifies the MobilePhones2 property.
	FieldUri_Persona_MobilePhones2 FieldUri = "persona:MobilePhones2"
	// FieldUri_Persona_AssistantPhoneNumbers identifies the AssistantPhoneNumbers property.
	FieldUri_Persona_AssistantPhoneNumbers FieldUri = "persona:AssistantPhoneNumbers"
	// FieldUri_Persona_CallbackPhones identifies the CallbackPhones property.
	FieldUri_Persona_CallbackPhones FieldUri = "persona:CallbackPhones"
	// FieldUri_Persona_CarPhones identifies the CarPhones property.
	FieldUri_Persona_CarPhones FieldUri = "persona:CarPhones"
	// FieldUri_Persona_HomeFaxes identifies the HomeFaxes property.
	FieldUri_Persona_HomeFaxes FieldUri = "persona:HomeFaxes"
	// FieldUri_Persona_OrganizationMainPhones identifies the OrganizationMainPhones property.
	FieldUri_Persona_OrganizationMainPhones FieldUri = "persona:OrganizationMainPhones"
	// FieldUri_Persona_OtherFaxes identifies the OtherFaxes property.
	FieldUri_Persona_OtherFaxes FieldUri = "persona:OtherFaxes"
	// FieldUri_Persona_OtherTelephones identifies the OtherTelephones property.
	FieldUri_Persona_OtherTelephones FieldUri = "persona:OtherTelephones"
	// FieldUri_Persona_OtherPhones2 identifies the OtherPhones2 property.
	FieldUri_Persona_OtherPhones2 FieldUri = "persona:OtherPhones2"
	// FieldUri_Persona_Pagers identifies the Pagers property.
	FieldUri_Persona_Pagers FieldUri = "persona:Pagers"
	// FieldUri_Persona_RadioPhones identifies the RadioPhones property.
	FieldUri_Persona_RadioPhones FieldUri = "persona:RadioPhones"
	// FieldUri_Persona_TelexNumbers identifies the TelexNumbers property.
	FieldUri_Persona_TelexNumbers FieldUri = "persona:TelexNumbers"
	// FieldUri_Persona_WorkFaxes identifies the WorkFaxes property.
	FieldUri_Persona_WorkFaxes FieldUri = "persona:WorkFaxes"
	// FieldUri_Persona_Emails1 identifies the Emails1 property.
	FieldUri_Persona_Emails1 FieldUri = "persona:Emails1"
	// FieldUri_Persona_Emails2 identifies the Emails2 property.
	FieldUri_Persona_Emails2 FieldUri = "persona:Emails2"
	// FieldUri_Persona_Emails3 identifies the Emails3 property.
	FieldUri_Persona_Emails3 FieldUri = "persona:Emails3"
	// FieldUri_Persona_BusinessHomePages identifies the BusinessHomePages property.
	FieldUri_Persona_BusinessHomePages FieldUri = "persona:BusinessHomePages"
	// FieldUri_Persona_School identifies the School property.
	FieldUri_Persona_School FieldUri = "persona:School"
	// FieldUri_Persona_PersonalHomePages identifies the PersonalHomePages property.
	FieldUri_Persona_PersonalHomePages FieldUri = "persona:PersonalHomePages"
	// FieldUri_Persona_OfficeLocations identifies the OfficeLocations property.
	FieldUri_Persona_OfficeLocations FieldUri = "persona:OfficeLocations"
	// FieldUri_Persona_BusinessAddresses identifies the BusinessAddresses property.
	FieldUri_Persona_BusinessAddresses FieldUri = "persona:BusinessAddresses"
	// FieldUri_Persona_HomeAddresses identifies the HomeAddresses property.
	FieldUri_Persona_HomeAddresses FieldUri = "persona:HomeAddresses"
	// FieldUri_Persona_OtherAddresses identifies the OtherAddresses property.
	FieldUri_Persona_OtherAddresses FieldUri = "persona:OtherAddresses"
	// FieldUri_Persona_Titles identifies the Titles property.
	FieldUri_Persona_Titles FieldUri = "persona:Titles"
	// FieldUri_Persona_Departments identifies the Departments property.
	FieldUri_Persona_Departments FieldUri = "persona:Departments"
	// FieldUri_Persona_CompanyNames identifies the CompanyNames property.
	FieldUri_Persona_CompanyNames FieldUri = "persona:CompanyNames"
	// FieldUri_Persona_Managers identifies the Managers property.
	FieldUri_Persona_Managers FieldUri = "persona:Managers"
	// FieldUri_Persona_AssistantNames identifies the AssistantNames property.
	FieldUri_Persona_AssistantNames FieldUri = "persona:AssistantNames"
	// FieldUri_Persona_Professions identifies the Professions property.
	FieldUri_Persona_Professions FieldUri = "persona:Professions"
	// FieldUri_Persona_SpouseNames identifies the SpouseNames property.
	FieldUri_Persona_SpouseNames FieldUri = "persona:SpouseNames"
	// FieldUri_Persona_Hobbies identifies the Hobbies property.
	FieldUri_Persona_Hobbies FieldUri = "persona:Hobbies"
	// FieldUri_Persona_WeddingAnniversaries identifies the WeddingAnniversaries property.
	FieldUri_Persona_WeddingAnniversaries FieldUri = "persona:WeddingAnniversaries"
	// FieldUri_Persona_Birthdays identifies the Birthdays property.
	FieldUri_Persona_Birthdays FieldUri = "persona:Birthdays"
	// FieldUri_Persona_Children identifies the Children property.
	FieldUri_Persona_Children FieldUri = "persona:Children"
	// FieldUri_Persona_Locations identifies the Locations property.
	FieldUri_Persona_Locations FieldUri = "persona:Locations"
	// FieldUri_Persona_ExtendedProperties identifies the ExtendedProperties property.
	FieldUri_Persona_ExtendedProperties FieldUri = "persona:ExtendedProperties"
	// FieldUri_Persona_PostalAddress identifies the PostalAddress property.
	FieldUri_Persona_PostalAddress FieldUri = "persona:PostalAddress"
	// FieldUri_Persona_Bodies identifies the Bodies property.
	FieldUri_Persona_Bodies FieldUri = "persona:Bodies"
)

type FieldURI struct {
	XMLName  xml.Name `xml:"FieldURI"`
	FieldURI FieldUri `xml:",attr"`
}
