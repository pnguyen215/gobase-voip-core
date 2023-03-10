package config

const (
	AmiListenerEventDeviceStateChange          = "DeviceStateChange" // Raised when a device state changes. This differs from the ExtensionStatus event because this event is raised for all device state changes, not only for changes that affect dialplan hints.
	AmiListenerEventNewChannel                 = "Newchannel"        // Raised when a new channel is created.
	AmiListenerEventNewState                   = "Newstate"          // Raised when a channel's state changes.
	AmiListenerEventSuccessfulAuth             = "SuccessfulAuth"    // Raised when a request successfully authenticates with a service
	AmiListenerEventNewExtension               = "NewExten"          // Raised when a channel enters a new context, extension, priority.
	AmiListenerEventNewCallerId                = "NewCallerid"       // Raised when a channel receives new Caller ID information.
	AmiListenerEventNewConnectedLine           = "NewConnectedLine"  // Raised when a channel's connected line information is changed.
	AmiListenerEventDialBegin                  = "DialBegin"         // Raised when a dial action has started
	AmiListenerEventUserEvent                  = "UserEvent"         // A user defined event raised from the dialplan.
	AmiListenerEventBridgeCreate               = "BridgeCreate"      // Raised when a bridge is created
	AmiListenerEventBridgeEnter                = "BridgeEnter"       // Raised when a channel enters a bridge.
	AmiListenerEventHangupRequest              = "HangupRequest"     // Raised when a hangup is requested
	AmiListenerEventBridgeLeave                = "BridgeLeave"       // Raised when a channel leaves a bridge.
	AmiListenerEventBridgeDestroy              = "BridgeDestroy"     // Raised when a bridge is destroyed.
	AmiListenerEventHangup                     = "Hangup"            // Raised when a channel is hung up
	AmiListenerEventSoftHangupRequest          = "SoftHangupRequest" // Raised when a soft hangup is requested with a specific cause code.
	AmiListenerEventQueueParams                = "QueueParams"
	AmiListenerEventQueueMember                = "QueueMember"
	AmiListenerEventQueueStatusComplete        = "QueueStatusComplete"
	AmiListenerEventQueueMemberPause           = "QueueMemberPause"  // Raised when a member is paused/unpaused in the queue.
	AmiListenerEventLocalBridge                = "LocalBridge"       // Raised when two halves of a Local Channel form a bridge.
	AmiListenerEventDialEnd                    = "DialEnd"           // Raised when a dial action has completed.
	AmiListenerEventConfBridgeJoin             = "ConfbridgeJoin"    // Raised when a channel joins a Confbridge conference.
	AmiListenerEventConfBridgeTalking          = "ConfbridgeTalking" // Raised when a confbridge participant begins or ends talking.
	AmiListenerEventConfBridgeKick             = "ConfbridgeKick"
	AmiListenerEventConfBridgeLeave            = "ConfbridgeLeave" // Raised when a channel leaves a Confbridge conference.
	AmiListenerEventMessageWaiting             = "MessageWaiting"  // Raised when the state of messages in a voicemail mailbox has changed or when a channel has finished interacting with a mailbox
	AmiListenerEventCdr                        = "Cdr"             // Raised when a CDR is generated. The Cdr event is only raised when the cdr_manager backend is loaded and registered with the CDR engine.
	AmiListenerEventExtensionStatus            = "ExtensionStatus" // Raised when a hint changes due to a device state change.
	AmiListenerEventConnect                    = "Connect"
	AmiListenerEventDisconnect                 = "Disconnect"
	AmiListenerEventDial                       = "Dial"
	AmiListenerEventBridge                     = "Bridge"
	AmiListenerEventRename                     = "Rename"            // Raised when the name of a channel is changed.
	AmiListenerEventVarSet                     = "VarSet"            // Raised when a variable local to the gosub stack frame is set due to a subroutine call.
	AmiListenerEventParkedCall                 = "ParkedCall"        // Raised when a channel is parked.
	AmiListenerEventParkedCallGiveUp           = "ParkedCallGiveUp"  // Raised when a channel leaves a parking lot because it hung up without being answered
	AmiListenerEventParkedCallTimeOut          = "ParkedCallTimeOut" // Raised when a channel leaves a parking lot due to reaching the time limit of being parked.
	AmiListenerEventUnParkedCall               = "UnparkedCall"      // Raised when a channel leaves a parking lot because it was retrieved from the parking lot and reconnected.
	AmiListenerEventJoin                       = "Join"
	AmiListenerEventLeave                      = "Leave"
	AmiListenerEventQueueMemberStatus          = "QueueMemberStatus"  // Raised when a Queue member's status has changed.
	AmiListenerEventQueueMemberPenalty         = "QueueMemberPenalty" // Raised when a member's penalty is changed.
	AmiListenerEventQueueMemberAdded           = "QueueMemberAdded"   // Raised when a member is added to the queue.
	AmiListenerEventQueueMemberRemoved         = "QueueMemberRemoved" // Raised when a member is removed from the queue.
	AmiListenerEventAbstractMeetMe             = "AbstractMeetMe"
	AmiListenerEventOriginateResponse          = "OriginateResponse" // Raised in response to an Originate command
	AmiListenerEventOriginate                  = "Originate"
	AmiListenerEventAgents                     = "Agents"        // Response event in a series to the Agents AMI action containing information about a defined agent.
	AmiListenerEventAgentCalled                = "AgentCalled"   // Raised when an queue member is notified of a caller in the queue
	AmiListenerEventAgentConnect               = "AgentConnect"  // Raised when a queue member answers and is bridged to a caller in the queue.
	AmiListenerEventAgentComplete              = "AgentComplete" // Raised when a queue member has finished servicing a caller in the queue.
	AmiListenerEventAgentCallbackLogin         = "AgentCallbackLogin"
	AmiListenerEventAgentCallbackLogoff        = "AgentCallbackLogoff"
	AmiListenerEventAgentLogin                 = "AgentLogin"  // Raised when an Agent has logged in.
	AmiListenerEventAgentLogoff                = "AgentLogoff" // Raised when an Agent has logged off.
	AmiListenerEventHoldedCall                 = "HoldedCall"
	AmiListenerEventPeerStatus                 = "PeerStatus" // Raised when the state of a peer changes.
	AmiListenerEventPeerlistComplete           = "PeerlistComplete"
	AmiListenerEventPeerEntry                  = "PeerEntry"
	AmiListenerEventStatus                     = "Status"            // Raised in response to a Status command.
	AmiListenerEventStatusComplete             = "StatusComplete"    // Raised in response to a Status command.
	AmiListenerEventAgentRingNoAnswer          = "AgentRingNoAnswer" // Raised when a queue member is notified of a caller in the queue and fails to answer.
	AmiListenerEventHold                       = "Hold"              // Raised when a channel goes on hold.
	AmiListenerEventChannelUpdate              = "ChannelUpdate"
	AmiListenerEventDialState                  = "DialState"       // Raised when dial status has changed.
	AmiListenerEventInvalidPassword            = "InvalidPassword" // Raised when a request provides an invalid password during an authentication attempt
	AmiListenerEventMusicOnHold                = "MusicOnHold"
	AmiListenerEventPickup                     = "Pickup" // Raised when a call pickup occurs
	AmiListenerEventPriEvent                   = "PriEvent"
	AmiListenerEventQueue                      = "Queue"
	AmiListenerEventAgentsComplete             = "AgentsComplete" // Final response event in a series of events to the Agents AMI action.
	AmiListenerEventUnHold                     = "Unhold"         // Raised when a channel goes off hold.
	AmiListenerEventDbGetResponse              = "DbGetResponse"
	AmiListenerEventCommon                     = "Common"
	AmiListenerEventHangupHandlerPush          = "HangupHandlerPush"          // Raised when a hangup handler is added to the handler stack by the CHANNEL() function.
	AmiListenerEventHangupHandlerRun           = "HangupHandlerRun"           // Raised when a hangup handler is about to be called.
	AmiListenerEventAgentDump                  = "AgentDump"                  // Raised when a queue member hangs up on a caller in the queue
	AmiListenerEventAGIExecEnd                 = "AGIExecEnd"                 // Raised when a received AGI command completes processing.
	AmiListenerEventAGIExecStart               = "AGIExecStart"               // Raised when a received AGI command starts processing.
	AmiListenerEventAlarm                      = "Alarm"                      // Raised when an alarm is set on a DAHDI channel.
	AmiListenerEventAlarmClear                 = "AlarmClear"                 // Raised when an alarm is cleared on a DAHDI channel.
	AmiListenerEventAOCD                       = "AOC-D"                      // Raised when an Advice of Charge message is sent during a call.
	AmiListenerEventAOCE                       = "AOC-E"                      // Raised when an Advice of Charge message is sent at the end of a call.
	AmiListenerEventAOCS                       = "AOC-S"                      // Raised when an Advice of Charge message is sent at the beginning of a call.
	AmiListenerEventAorDetail                  = "AorDetail"                  // Provide details about an Address of Record (AoR) section.
	AmiListenerEventAorList                    = "AorList"                    // Provide details about an Address of Record (AoR) section.
	AmiListenerEventAorListComplete            = "AorListComplete"            // Provide final information about an aor list
	AmiListenerEventAsyncAGIEnd                = "AsyncAGIEnd"                // Raised when a channel stops AsyncAGI command processing.
	AmiListenerEventAsyncAGIExec               = "AsyncAGIExec"               // Raised when AsyncAGI completes an AGI command.
	AmiListenerEventAsyncAGIStart              = "AsyncAGIStart"              // Raised when a channel starts AsyncAGI command processing.
	AmiListenerEventAttendedTransfer           = "AttendedTransfer"           // Raised when an attended transfer is complete
	AmiListenerEventAuthDetail                 = "AuthDetail"                 // Provide details about an authentication section.
	AmiListenerEventAuthList                   = "AuthList"                   // Provide details about an Address of Record (Auth) section.
	AmiListenerEventAuthListComplete           = "AuthListComplete"           // Provide final information about an auth list.
	AmiListenerEventAuthMethodNotAllowed       = "AuthMethodNotAllowed"       // Raised when a request used an authentication method not allowed by the service
	AmiListenerEventBlindTransfer              = "BlindTransfer"              // Raised when a blind transfer is complete.
	AmiListenerEventBridgeInfoChannel          = "BridgeInfoChannel"          // Information about a channel in a bridge.
	AmiListenerEventBridgeInfoComplete         = "BridgeInfoComplete"         // Information about a bridge.
	AmiListenerEventBridgeMerge                = "BridgeMerge"                // Raised when two bridges are merged
	AmiListenerEventBridgeVideoSourceUpdate    = "BridgeVideoSourceUpdate"    // Raised when the channel that is the source of video in a bridge changes
	AmiListenerEventCel                        = "CEL"                        // Raised when a Channel Event Log is generated for a channel.
	AmiListenerEventChallengeResponseFailed    = "ChallengeResponseFailed"    // Raised when a request's attempt to authenticate has been challenged, and the request failed the authentication challenge
	AmiListenerEventChallengeSent              = "ChallengeSent"              // Raised when an Asterisk service sends an authentication challenge to a request.
	AmiListenerEventChannelTalkingStart        = "ChannelTalkingStart"        // Raised when talking is detected on a channel.
	AmiListenerEventChannelTalkingStop         = "ChannelTalkingStop"         // Raised when talking is no longer detected on a channel.
	AmiListenerEventChanSpyStart               = "ChanSpyStart"               // Raised when one channel begins spying on another channel.
	AmiListenerEventChanSpyStop                = "ChanSpyStop"                // Raised when a channel has stopped spying.
	AmiListenerEventConfbridgeEnd              = "ConfbridgeEnd"              // Raised when a conference ends.
	AmiListenerEventConfbridgeJoin             = "ConfbridgeJoin"             // Raised when a channel joins a Confbridge conference.
	AmiListenerEventConfbridgeLeave            = "ConfbridgeLeave"            // Raised when a channel leaves a Confbridge conference.
	AmiListenerEventConfbridgeList             = "ConfbridgeList"             // Raised as part of the ConfbridgeList action response list.
	AmiListenerEventConfbridgeMute             = "ConfbridgeMute"             // Raised when a Confbridge participant mutes.
	AmiListenerEventConfbridgeRecord           = "ConfbridgeRecord"           // Raised when a conference starts recording.
	AmiListenerEventConfbridgeStart            = "ConfbridgeStart"            // Raised when a conference starts.
	AmiListenerEventConfbridgeStopRecord       = "ConfbridgeStopRecord"       // Raised when a conference that was recording stops recording.
	AmiListenerEventConfbridgeTalking          = "ConfbridgeTalking"          // Raised when a confbridge participant begins or ends talking.
	AmiListenerEventConfbridgeUnMute           = "ConfbridgeUnmute"           // Raised when a confbridge participant unmutes.
	AmiListenerEventContactList                = "ContactList"                // Provide details about a contact section.
	AmiListenerEventContactListComplete        = "ContactListComplete"        // Provide final information about a contact list.
	AmiListenerEventContactStatus              = "ContactStatus"              // Raised when the state of a contact changes
	AmiListenerEventContactStatusDetail        = "ContactStatusDetail"        // Provide details about a contact's status.
	AmiListenerEventCoreShowChannel            = "CoreShowChannel"            // Raised in response to a CoreShowChannels command.
	AmiListenerEventCoreShowChannelsComplete   = "CoreShowChannelsComplete"   // Raised at the end of the CoreShowChannel list produced by the CoreShowChannels command.
	AmiListenerEventDAHDIChannel               = "DAHDIChannel"               // Raised when a DAHDI channel is created or an underlying technology is associated with a DAHDI channel
	AmiListenerEventDeviceStateListComplete    = "DeviceStateListComplete"    // Indicates the end of the list the current known extension states.
	AmiListenerEventDNDState                   = "DNDState"                   // page 317, Raised when the Do Not Disturb state is changed on a DAHDI channel.
	AmiListenerEventDTMFBegin                  = "DTMFBegin"                  // Raised when a DTMF digit has started on a channel
	AmiListenerEventDTMFEnd                    = "DTMFEnd"                    // Raised when a DTMF digit has ended on a channel.
	AmiListenerEventEndpointDetail             = "EndpointDetail"             // Provide details about an endpoint section.
	AmiListenerEventEndpointDetailComplete     = "EndpointDetailComplete"     // Provide final information about endpoint details
	AmiListenerEventEndpointList               = "EndpointList"               // Provide details about a contact's status.
	AmiListenerEventEndpointListComplete       = "EndpointListComplete"       // Provide final information about an endpoint list
	AmiListenerEventExtensionStateListComplete = "ExtensionStateListComplete" // Indicates the end of the list the current known extension states.
	AmiListenerEventFailedACL                  = "FailedACL"                  // Raised when a request violates an ACL check.
	AmiListenerEventFAXSession                 = "FAXSession"                 // Raised in response to FAXSession manager command
	AmiListenerEventFAXSessionsComplete        = "FAXSessionsComplete"        // Raised when all FAXSession events are completed for a FAXSessions command
	AmiListenerEventFAXSessionsEntry           = "FAXSessionsEntry"           // A single list item for the FAXSessions AMI command
	AmiListenerEventFAXStats                   = "FAXStats"                   // Raised in response to FAXStats manager command
	AmiListenerEventFAXStatus                  = "FAXStatus"                  // Raised periodically during a fax transmission.
	AmiListenerEventFullyBooted                = "FullyBooted"                // Raised when all Asterisk initialization procedures have finished
	AmiListenerEventHangupHandlerPop           = "HangupHandlerPop"           // Raised when a hangup handler is removed from the handler stack by the CHANNEL() function.
	AmiListenerEventIdentifyDetail             = "IdentifyDetail"             // Provide details about an identify section.
	AmiListenerEventInvalidAccountID           = "InvalidAccountID"           // Raised when a request fails an authentication check due to an invalid account ID.
	AmiListenerEventInvalidTransport           = "InvalidTransport"           // Raised when a request attempts to use a transport not allowed by the Asterisk service.
	AmiListenerEventLoad                       = "Load"                       // Raised when a module has been loaded in Asterisk.
	AmiListenerEventLoadAverageLimit           = "LoadAverageLimit"           // Raised when a request fails because a configured load average limit has been reached.
	AmiListenerEventLocalOptimizationBegin     = "LocalOptimizationBegin"     // Raised when two halves of a Local Channel begin to optimize themselves out of the media path
	AmiListenerEventLocalOptimizationEnd       = "LocalOptimizationEnd"       // Raised when two halves of a Local Channel have finished optimizing themselves out of the media path.
	AmiListenerEventLogChannel                 = "LogChannel"                 // Raised when a logging channel is re-enabled after a reload operation.
	AmiListenerEventMCID                       = "MCID"                       // Published when a malicious call ID request arrives
	AmiListenerEventMeetMeEnd                  = "MeetmeEnd"                  // Raised when a MeetMe conference ends
	AmiListenerEventMeetMeJoin                 = "MeetmeJoin"                 // Raised when a user joins a MeetMe conference.
	AmiListenerEventMeetMeLeave                = "MeetmeLeave"                // Raised when a user leaves a MeetMe conference.
	AmiListenerEventMeetMeMute                 = "MeetmeMute"                 // Raised when a MeetMe user is muted or unmuted.
	AmiListenerEventMeetMeTalking              = "MeetmeTalking"              // Raised when a MeetMe user begins or ends talking.
	AmiListenerEventMeetMeTalkRequest          = "MeetmeTalkRequest"          // Raised when a MeetMe user has started talking.
	AmiListenerEventMemoryLimit                = "MemoryLimit"                // Raised when a request fails due to an internal memory allocation failure.
	AmiListenerEventMiniVoiceMail              = "MiniVoiceMail"              // Raised when a notification is sent out by a MiniVoiceMail application
	AmiListenerEventMonitorStart               = "MonitorStart"               // Raised when monitoring has started on a channel.
	AmiListenerEventMonitorStop                = "MonitorStop"                // Raised when monitoring has stopped on a channel
	AmiListenerEventMusicOnHoldStart           = "MusicOnHoldStart"           // Raised when music on hold has started on a channel.
	AmiListenerEventMusicOnHoldStop            = "MusicOnHoldStop"            // Raised when music on hold has stopped on a channel.
	AmiListenerEventMWIGet                     = "MWIGet"                     // Raised in response to a MWIGet command.
	AmiListenerEventMWIGetComplete             = "MWIGetComplete"             // Raised in response to a MWIGet command.
	AmiListenerEventNewAccountCode             = "NewAccountCode"             // Raised when a Channel's AccountCode is changed.
	AmiListenerEventParkedCallSwap             = "ParkedCallSwap"             // Raised when a channel takes the place of a previously parked channel
	AmiListenerEventPresenceStateChange        = "PresenceStateChange"        // Raised when a presence state changes
	AmiListenerEventPresenceStateListComplete  = "PresenceStateListComplete"  // Indicates the end of the list the current known extension states.
	AmiListenerEventPresenceStatus             = "PresenceStatus"             // Raised when a hint changes due to a presence state change.
	AmiListenerEventQueueCallerAbandon         = "QueueCallerAbandon"         // Raised when a caller abandons the queue.
	AmiListenerEventQueueCallerJoin            = "QueueCallerJoin"            // Raised when a caller joins a Queue.
	AmiListenerEventQueueCallerLeave           = "QueueCallerLeave"           // Raised when a caller leaves a Queue.
	AmiListenerEventQueueMemberRinginuse       = "QueueMemberRinginuse"       // Raised when a member's ringinuse setting is changed.
	AmiListenerEventReceiveFAX                 = "ReceiveFAX"                 // Raised when a receive fax operation has completed.
	AmiListenerEventRegistry                   = "Registry"                   // Raised when an outbound registration completes.
	AmiListenerEventReload                     = "Reload"                     // Raised when a module has been reloaded in Asterisk.
	AmiListenerEventRequestBadFormat           = "RequestBadFormat"           // Raised when a request is received with bad formatting.
	AmiListenerEventRequestNotAllowed          = "RequestNotAllowed"          // Raised when a request is not allowed by the service.
	AmiListenerEventRequestNotSupported        = "RequestNotSupported"        // Raised when a request fails due to some aspect of the requested item not being supported by the service.
	AmiListenerEventRTCPReceived               = "RTCPReceived"               // Raised when an RTCP packet is received.
	AmiListenerEventRTCPSent                   = "RTCPSent"                   // Raised when an RTCP packet is sent.
	AmiListenerEventSendFAX                    = "SendFAX"                    // Raised when a send fax operation has completed.
	AmiListenerEventSessionLimit               = "SessionLimit"               // Raised when a request fails due to exceeding the number of allowed concurrent sessions for that service.
	AmiListenerEventSessionTimeout             = "SessionTimeout"             // Raised when a SIP session times out.
	AmiListenerEventShutdown                   = "Shutdown"                   // Raised when Asterisk is shutdown or restarted.
	AmiListenerEventSIPQualifyPeerDone         = "SIPQualifyPeerDone"         // Raised when SIPQualifyPeer has finished qualifying the specified peer
	AmiListenerEventSpanAlarm                  = "SpanAlarm"                  // Raised when an alarm is set on a DAHDI span.
	AmiListenerEventSpanAlarmClear             = "SpanAlarmClear"             // Raised when an alarm is cleared on a DAHDI span.
	AmiListenerEventTransportDetail            = "TransportDetail"            // Provide details about an authentication section.
	AmiListenerEventUnexpectedAddress          = "UnexpectedAddress"          // Raised when a request has a different source address then what is expected for a session already in progress with a service
	AmiListenerEventUnload                     = "Unload"                     // Raised when a module has been unloaded in Asterisk.
)
