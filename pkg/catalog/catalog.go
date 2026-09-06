package catalog

import "strings"

// PackageMeta defines the metadata and known declarations of a wa-core proto package.
type PackageMeta struct {
	Dir       string
	File      string
	Package   string
	GoPackage string
	Messages  []string
	Enums     []string
	Imports   []string
}

// Packages contains the complete catalog of all standard packages in wa-core/proto.
var Packages = map[string]PackageMeta{
	"instamadilloAddMessage": {
		Dir: "instamadilloAddMessage",
		File: "InstamadilloAddMessage.proto",
		Package: "InstamadilloAddMessage",
		GoPackage: "go.mau.fi/whatsmeow/proto/instamadilloAddMessage",
		Messages: []string{
			"AddMessageContent",
			"AddMessageMetadata",
			"AddMessagePayload",
			"EphemeralityParams",
			"ForwardingParams",
			"Like",
			"OpenMessageMicroSecondTimestamp",
			"Placeholder",
			"PrivateReplyInfo",
			"ReceiverFetchXma",
			"RepliedToMessage",
		},
		Enums: []string{
		},
		Imports: []string{
			"instamadilloCoreTypeActionLog/InstamadilloCoreTypeActionLog.proto",
			"instamadilloCoreTypeAdminMessage/InstamadilloCoreTypeAdminMessage.proto",
			"instamadilloCoreTypeCollection/InstamadilloCoreTypeCollection.proto",
			"instamadilloCoreTypeLink/InstamadilloCoreTypeLink.proto",
			"instamadilloCoreTypeMedia/InstamadilloCoreTypeMedia.proto",
			"instamadilloCoreTypeText/InstamadilloCoreTypeText.proto",
			"instamadilloXmaContentRef/InstamadilloXmaContentRef.proto",
		},
	},
	"instamadilloCoreTypeActionLog": {
		Dir: "instamadilloCoreTypeActionLog",
		File: "InstamadilloCoreTypeActionLog.proto",
		Package: "InstamadilloCoreTypeActionLog",
		GoPackage: "go.mau.fi/whatsmeow/proto/instamadilloCoreTypeActionLog",
		Messages: []string{
			"ActionLog",
			"ActionLogReaction",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"instamadilloCoreTypeAdminMessage": {
		Dir: "instamadilloCoreTypeAdminMessage",
		File: "InstamadilloCoreTypeAdminMessage.proto",
		Package: "InstamadilloCoreTypeAdminMessage",
		GoPackage: "go.mau.fi/whatsmeow/proto/instamadilloCoreTypeAdminMessage",
		Messages: []string{
			"AdminMessage",
			"DeviceAdminMessage",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"instamadilloCoreTypeCollection": {
		Dir: "instamadilloCoreTypeCollection",
		File: "InstamadilloCoreTypeCollection.proto",
		Package: "InstamadilloCoreTypeCollection",
		GoPackage: "go.mau.fi/whatsmeow/proto/instamadilloCoreTypeCollection",
		Messages: []string{
			"Collection",
		},
		Enums: []string{
		},
		Imports: []string{
			"instamadilloCoreTypeMedia/InstamadilloCoreTypeMedia.proto",
		},
	},
	"instamadilloCoreTypeLink": {
		Dir: "instamadilloCoreTypeLink",
		File: "InstamadilloCoreTypeLink.proto",
		Package: "InstamadilloCoreTypeLink",
		GoPackage: "go.mau.fi/whatsmeow/proto/instamadilloCoreTypeLink",
		Messages: []string{
			"ImageUrl",
			"Link",
			"LinkContext",
		},
		Enums: []string{
		},
		Imports: []string{
			"instamadilloCoreTypeMedia/InstamadilloCoreTypeMedia.proto",
		},
	},
	"instamadilloCoreTypeMedia": {
		Dir: "instamadilloCoreTypeMedia",
		File: "InstamadilloCoreTypeMedia.proto",
		Package: "InstamadilloCoreTypeMedia",
		GoPackage: "go.mau.fi/whatsmeow/proto/instamadilloCoreTypeMedia",
		Messages: []string{
			"AvatarSticker",
			"CommonMediaTransport",
			"Gif",
			"Media",
			"Raven",
			"RavenContent",
			"StaticPhoto",
			"Thumbnail",
			"Video",
			"VideoExtraMetadata",
			"Voice",
		},
		Enums: []string{
			"PjpegScanConfiguration",
		},
		Imports: []string{
		},
	},
	"instamadilloCoreTypeText": {
		Dir: "instamadilloCoreTypeText",
		File: "InstamadilloCoreTypeText.proto",
		Package: "InstamadilloCoreTypeText",
		GoPackage: "go.mau.fi/whatsmeow/proto/instamadilloCoreTypeText",
		Messages: []string{
			"AnimatedEmojiCharacterRange",
			"CommandRangeData",
			"FormattedText",
			"PowerUpsData",
			"Text",
		},
		Enums: []string{
		},
		Imports: []string{
			"instamadilloCoreTypeMedia/InstamadilloCoreTypeMedia.proto",
		},
	},
	"instamadilloDeleteMessage": {
		Dir: "instamadilloDeleteMessage",
		File: "InstamadilloDeleteMessage.proto",
		Package: "InstamadilloDeleteMessage",
		GoPackage: "go.mau.fi/whatsmeow/proto/instamadilloDeleteMessage",
		Messages: []string{
			"DeleteMessagePayload",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"instamadilloSupplementMessage": {
		Dir: "instamadilloSupplementMessage",
		File: "InstamadilloSupplementMessage.proto",
		Package: "InstamadilloSupplementMessage",
		GoPackage: "go.mau.fi/whatsmeow/proto/instamadilloSupplementMessage",
		Messages: []string{
			"ContentView",
			"EditText",
			"MediaInterventions",
			"MediaReaction",
			"OriginalTransportPayload",
			"Reaction",
			"SupplementMessageContent",
			"SupplementMessagePayload",
		},
		Enums: []string{
		},
		Imports: []string{
			"instamadilloCoreTypeMedia/InstamadilloCoreTypeMedia.proto",
		},
	},
	"instamadilloTransportPayload": {
		Dir: "instamadilloTransportPayload",
		File: "InstamadilloTransportPayload.proto",
		Package: "InstamadilloTransportPayload",
		GoPackage: "go.mau.fi/whatsmeow/proto/instamadilloTransportPayload",
		Messages: []string{
			"Franking",
			"TransportPayload",
		},
		Enums: []string{
			"PayloadCreator",
		},
		Imports: []string{
			"instamadilloAddMessage/InstamadilloAddMessage.proto",
			"instamadilloDeleteMessage/InstamadilloDeleteMessage.proto",
			"instamadilloSupplementMessage/InstamadilloSupplementMessage.proto",
		},
	},
	"instamadilloXmaContentRef": {
		Dir: "instamadilloXmaContentRef",
		File: "InstamadilloXmaContentRef.proto",
		Package: "InstamadilloXmaContentRef",
		GoPackage: "go.mau.fi/whatsmeow/proto/instamadilloXmaContentRef",
		Messages: []string{
			"ReceiverFetchXmaClipFetchParams",
			"ReceiverFetchXmaCommentFetchParams",
			"ReceiverFetchXmaFeedFetchParams",
			"ReceiverFetchXmaFetchParams",
			"ReceiverFetchXmaLiveFetchParams",
			"ReceiverFetchXmaLocationShareFetchParams",
			"ReceiverFetchXmaMediaNoteFetchParams",
			"ReceiverFetchXmaNoteFetchParams",
			"ReceiverFetchXmaProfileFetchParams",
			"ReceiverFetchXmaReelsAudioFetchParams",
			"ReceiverFetchXmaSocialContextFetchParams",
			"ReceiverFetchXmaStoryFetchParams",
			"XmaContentRef",
		},
		Enums: []string{
			"MediaNoteFetchParamsMessageType",
			"ReceiverFetchContentType",
			"XmaActionType",
		},
		Imports: []string{
		},
	},
	"waAICommon": {
		Dir: "waAICommon",
		File: "WAWebProtobufsAICommon.proto",
		Package: "WAWebProtobufsAICommon",
		GoPackage: "go.mau.fi/whatsmeow/proto/waAICommon",
		Messages: []string{
			"AIHomeState",
			"AIMediaCollectionMessage",
			"AIMediaCollectionMetadata",
			"AIMetadataOperation",
			"AIProvenance",
			"AIRegenerateMetadata",
			"AIRichResponseUnifiedResponse",
			"AISubscriptionUpsellMetadata",
			"AIThreadInfo",
			"BotAgeCollectionMetadata",
			"BotAgentDeepLinkMetadata",
			"BotAgentMetadata",
			"BotCapabilityMetadata",
			"BotCommandMetadata",
			"BotDocumentMessageMetadata",
			"BotFeedbackMessage",
			"BotGroupMetadata",
			"BotGroupParticipantMetadata",
			"BotHistoryShareMetadata",
			"BotImagineMetadata",
			"BotInfrastructureDiagnostics",
			"BotLinkedAccount",
			"BotLinkedAccountsMetadata",
			"BotMediaMetadata",
			"BotMemoryFact",
			"BotMemoryMetadata",
			"BotMemuMetadata",
			"BotMessageOrigin",
			"BotMessageOriginMetadata",
			"BotMessageSharingInfo",
			"BotMetadata",
			"BotMetricsMetadata",
			"BotModeSelectionMetadata",
			"BotModelMetadata",
			"BotPluginMetadata",
			"BotProgressIndicatorMetadata",
			"BotPromotionMessageMetadata",
			"BotPromptSuggestion",
			"BotPromptSuggestions",
			"BotPttPromptMetadata",
			"BotQuotaMetadata",
			"BotReminderMetadata",
			"BotRenderingConfigMetadata",
			"BotRenderingMetadata",
			"BotResolvedToolCallMetadata",
			"BotSessionMetadata",
			"BotSignatureVerificationMetadata",
			"BotSignatureVerificationUseCaseProof",
			"BotSourcesMetadata",
			"BotSuggestedPromptMetadata",
			"BotUnifiedResponseMutation",
			"ForwardedAIBotMessageInfo",
			"HatchMetadataSync",
			"InThreadSurveyMetadata",
			"SessionTransparencyMetadata",
		},
		Enums: []string{
			"AISubscriptionRequestType",
			"BotMetricsEntryPoint",
			"BotMetricsThreadEntryPoint",
			"BotSessionSource",
			"SessionTransparencyType",
		},
		Imports: []string{
			"waCommon/WACommon.proto",
		},
	},
	"waAICommonDeprecated": {
		Dir: "waAICommonDeprecated",
		File: "WAAICommonDeprecated.proto",
		Package: "WAAICommonDeprecated",
		GoPackage: "go.mau.fi/whatsmeow/proto/waAICommonDeprecated",
		Messages: []string{
			"AIRichResponseCodeMetadata",
			"AIRichResponseContentItemsMetadata",
			"AIRichResponseDynamicMetadata",
			"AIRichResponseGridImageMetadata",
			"AIRichResponseImageURL",
			"AIRichResponseInlineImageMetadata",
			"AIRichResponseLatexMetadata",
			"AIRichResponseMapMetadata",
			"AIRichResponseSubMessage",
			"AIRichResponseTableMetadata",
		},
		Enums: []string{
			"AIRichResponseMessageType",
			"AIRichResponseSubMessageType",
		},
		Imports: []string{
		},
	},
	"waAdv": {
		Dir: "waAdv",
		File: "WAAdv.proto",
		Package: "WAAdv",
		GoPackage: "go.mau.fi/whatsmeow/proto/waAdv",
		Messages: []string{
			"ADVDeviceIdentity",
			"ADVKeyIndexList",
			"ADVSignedDeviceIdentity",
			"ADVSignedDeviceIdentityHMAC",
			"ADVSignedKeyIndexList",
		},
		Enums: []string{
			"ADVEncryptionType",
		},
		Imports: []string{
		},
	},
	"waAea": {
		Dir: "waAea",
		File: "WAWebProtobufsAea.proto",
		Package: "WAWebProtobufsAea",
		GoPackage: "go.mau.fi/whatsmeow/proto/waAea",
		Messages: []string{
			"NonE2EEAttestation",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waArmadilloApplication": {
		Dir: "waArmadilloApplication",
		File: "WAArmadilloApplication.proto",
		Package: "WAArmadilloApplication",
		GoPackage: "go.mau.fi/whatsmeow/proto/waArmadilloApplication",
		Messages: []string{
			"Armadillo",
		},
		Enums: []string{
		},
		Imports: []string{
			"waArmadilloXMA/WAArmadilloXMA.proto",
			"waCommon/WACommon.proto",
		},
	},
	"waArmadilloBackupCommon": {
		Dir: "waArmadilloBackupCommon",
		File: "WAArmadilloBackupCommon.proto",
		Package: "WAArmadilloBackupCommon",
		GoPackage: "go.mau.fi/whatsmeow/proto/waArmadilloBackupCommon",
		Messages: []string{
			"FrankingMetadata",
			"Metadata",
			"Subprotocol",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waArmadilloBackupMessage": {
		Dir: "waArmadilloBackupMessage",
		File: "WAArmadilloBackupMessage.proto",
		Package: "WAArmadilloBackupMessage",
		GoPackage: "go.mau.fi/whatsmeow/proto/waArmadilloBackupMessage",
		Messages: []string{
			"BackupMessage",
		},
		Enums: []string{
		},
		Imports: []string{
			"waArmadilloBackupCommon/WAArmadilloBackupCommon.proto",
		},
	},
	"waArmadilloICDC": {
		Dir: "waArmadilloICDC",
		File: "WAArmadilloICDC.proto",
		Package: "WAArmadilloICDC",
		GoPackage: "go.mau.fi/whatsmeow/proto/waArmadilloICDC",
		Messages: []string{
			"ICDCIdentityList",
			"SignedICDCIdentityList",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waArmadilloMiTransportAdminMessage": {
		Dir: "waArmadilloMiTransportAdminMessage",
		File: "WAArmadilloMiTransportAdminMessage.proto",
		Package: "WAArmadilloMiTransportAdminMessage",
		GoPackage: "go.mau.fi/whatsmeow/proto/waArmadilloMiTransportAdminMessage",
		Messages: []string{
			"MiTransportAdminMessage",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waArmadilloTransportEvent": {
		Dir: "waArmadilloTransportEvent",
		File: "WAArmadilloTransportEvent.proto",
		Package: "WAArmadilloTransportEvent",
		GoPackage: "go.mau.fi/whatsmeow/proto/waArmadilloTransportEvent",
		Messages: []string{
			"TransportEvent",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waArmadilloXMA": {
		Dir: "waArmadilloXMA",
		File: "WAArmadilloXMA.proto",
		Package: "WAArmadilloXMA",
		GoPackage: "go.mau.fi/whatsmeow/proto/waArmadilloXMA",
		Messages: []string{
			"ExtendedContentMessage",
		},
		Enums: []string{
		},
		Imports: []string{
			"waCommon/WACommon.proto",
		},
	},
	"waBotMetadata": {
		Dir: "waBotMetadata",
		File: "WABotMetadata.proto",
		Package: "WABotMetadata",
		GoPackage: "go.mau.fi/whatsmeow/proto/waBotMetadata",
		Messages: []string{
			"AIThreadInfo",
			"BotAgeCollectionMetadata",
			"BotAvatarMetadata",
			"BotCapabilityMetadata",
			"BotImagineMetadata",
			"BotLinkedAccount",
			"BotLinkedAccountsMetadata",
			"BotMediaMetadata",
			"BotMemoryFact",
			"BotMemoryMetadata",
			"BotMemuMetadata",
			"BotMessageOrigin",
			"BotMessageOriginMetadata",
			"BotMetadata",
			"BotMetricsMetadata",
			"BotModeSelectionMetadata",
			"BotModelMetadata",
			"BotPluginMetadata",
			"BotProgressIndicatorMetadata",
			"BotPromotionMessageMetadata",
			"BotPromptSuggestion",
			"BotPromptSuggestions",
			"BotQuotaMetadata",
			"BotReminderMetadata",
			"BotRenderingMetadata",
			"BotSessionMetadata",
			"BotSignatureVerificationMetadata",
			"BotSignatureVerificationUseCaseProof",
			"BotSourcesMetadata",
			"BotSuggestedPromptMetadata",
			"BotUnifiedResponseMutation",
			"InThreadSurveyMetadata",
		},
		Enums: []string{
			"BotMetricsEntryPoint",
			"BotMetricsThreadEntryPoint",
			"BotSessionSource",
		},
		Imports: []string{
			"waCommon/WACommon.proto",
		},
	},
	"waCert": {
		Dir: "waCert",
		File: "WACert.proto",
		Package: "WACert",
		GoPackage: "go.mau.fi/whatsmeow/proto/waCert",
		Messages: []string{
			"CertChain",
			"NoiseCertificate",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waChatLockSettings": {
		Dir: "waChatLockSettings",
		File: "WAWebProtobufsChatLockSettings.proto",
		Package: "WAWebProtobufsChatLockSettings",
		GoPackage: "go.mau.fi/whatsmeow/proto/waChatLockSettings",
		Messages: []string{
			"ChatLockSettings",
		},
		Enums: []string{
		},
		Imports: []string{
			"waUserPassword/WAWebProtobufsUserPassword.proto",
		},
	},
	"waCommon": {
		Dir: "waCommon",
		File: "WACommon.proto",
		Package: "WACommon",
		GoPackage: "go.mau.fi/whatsmeow/proto/waCommon",
		Messages: []string{
			"Command",
			"LimitSharing",
			"Mention",
			"MessageKey",
			"MessageText",
			"SubProtocol",
		},
		Enums: []string{
			"FutureProofBehavior",
		},
		Imports: []string{
		},
	},
	"waCommonParameterised": {
		Dir: "waCommonParameterised",
		File: "WACommonParameterised.proto",
		Package: "WACommonParameterised",
		GoPackage: "go.mau.fi/whatsmeow/proto/waCommonParameterised",
		Messages: []string{
			"Command",
			"Mention",
			"MessageKey",
			"MessageText",
			"SubProtocol",
		},
		Enums: []string{
			"FutureProofBehavior",
		},
		Imports: []string{
		},
	},
	"waCompanionReg": {
		Dir: "waCompanionReg",
		File: "WACompanionReg.proto",
		Package: "WACompanionReg",
		GoPackage: "go.mau.fi/whatsmeow/proto/waCompanionReg",
		Messages: []string{
			"ClientPairingProps",
			"CompanionCommitment",
			"CompanionEphemeralIdentity",
			"DeviceProps",
			"EncryptedPairingRequest",
			"PairingRequest",
			"PrimaryEphemeralIdentity",
			"ProloguePayload",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waConsumerApplication": {
		Dir: "waConsumerApplication",
		File: "WAConsumerApplication.proto",
		Package: "WAConsumerApplication",
		GoPackage: "go.mau.fi/whatsmeow/proto/waConsumerApplication",
		Messages: []string{
			"ConsumerApplication",
		},
		Enums: []string{
		},
		Imports: []string{
			"waCommon/WACommon.proto",
		},
	},
	"waConsumerApplicationParameterised": {
		Dir: "waConsumerApplicationParameterised",
		File: "WAConsumerApplicationParameterised.proto",
		Package: "WAConsumerApplicationParameterised",
		GoPackage: "go.mau.fi/whatsmeow/proto/waConsumerApplicationParameterised",
		Messages: []string{
			"ConsumerApplication",
		},
		Enums: []string{
		},
		Imports: []string{
			"waCommonParameterised/WACommonParameterised.proto",
		},
	},
	"waDeviceCapabilities": {
		Dir: "waDeviceCapabilities",
		File: "WAWebProtobufsDeviceCapabilities.proto",
		Package: "WAWebProtobufsDeviceCapabilities",
		GoPackage: "go.mau.fi/whatsmeow/proto/waDeviceCapabilities",
		Messages: []string{
			"DeviceCapabilities",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waE2E": {
		Dir: "waE2E",
		File: "WAWebProtobufsE2E.proto",
		Package: "WAWebProtobufsE2E",
		GoPackage: "go.mau.fi/whatsmeow/proto/waE2E",
		Messages: []string{
			"AIQueryFanout",
			"AIRichResponseMessage",
			"ActionLink",
			"AlbumMessage",
			"AppStateFatalExceptionNotification",
			"AppStateSyncKey",
			"AppStateSyncKeyData",
			"AppStateSyncKeyFingerprint",
			"AppStateSyncKeyId",
			"AppStateSyncKeyRequest",
			"AppStateSyncKeyShare",
			"AudioMessage",
			"BCallMessage",
			"BotHistoryShareSyncMetadata",
			"ButtonsMessage",
			"ButtonsResponseMessage",
			"Call",
			"CallLogMessage",
			"CancelPaymentRequestMessage",
			"Chat",
			"ChatCustomImageWallpaper",
			"ChatDefaultWallpaper",
			"ChatSolidColorWallpaper",
			"ChatStockImageWallpaper",
			"ChatThemeSetting",
			"CloudAPIThreadControlNotification",
			"CommentMessage",
			"ConditionalRevealMessage",
			"ContactMessage",
			"ContactsArrayMessage",
			"ContextInfo",
			"DeclinePaymentRequestMessage",
			"DeviceListMetadata",
			"DeviceSentMessage",
			"DisappearingMode",
			"DocumentMessage",
			"EmbeddedContent",
			"EmbeddedMessage",
			"EmbeddedMusic",
			"EncCommentMessage",
			"EncEventResponseMessage",
			"EncReactionMessage",
			"EventInviteMessage",
			"EventMessage",
			"EventResponseMessage",
			"ExtendedTextMessage",
			"FullHistorySyncOnDemandConfig",
			"FullHistorySyncOnDemandRequestMetadata",
			"FutureProofMessage",
			"GroupInviteMessage",
			"GroupMention",
			"GroupRootKeyShare",
			"GroupRootKeyShareEntry",
			"HighlyStructuredMessage",
			"HistoryShareMessageEntry",
			"HistorySyncMessageAccessStatus",
			"HistorySyncNotification",
			"HydratedTemplateButton",
			"ImageMessage",
			"InitialSecurityNotificationSettingSync",
			"InteractiveAnnotation",
			"InteractiveMessage",
			"InteractiveResponseMessage",
			"InvoiceMessage",
			"KeepInChatMessage",
			"LIDMigrationMappingSyncMessage",
			"LinkPreviewMetadata",
			"ListMessage",
			"ListResponseMessage",
			"LiveLocationMessage",
			"Location",
			"LocationMessage",
			"MMSThumbnailMetadata",
			"MarkAsVerifiedAction",
			"MediaDomainInfo",
			"MediaNotifyMessage",
			"MemberLabel",
			"Message",
			"MessageAssociation",
			"MessageContextInfo",
			"MessageHistoryBundle",
			"MessageHistoryMetadata",
			"MessageHistoryNotice",
			"MessageSecretMessage",
			"Money",
			"MusicMessage",
			"NewsletterAdminInviteMessage",
			"NewsletterFollowerInviteMessage",
			"OrderMessage",
			"PaymentBackground",
			"PaymentExtendedMetadata",
			"PaymentInviteMessage",
			"PaymentLinkMetadata",
			"PaymentReminderMessage",
			"PeerDataOperationRequestMessage",
			"PeerDataOperationRequestResponseMessage",
			"PinInChatMessage",
			"PlaceholderMessage",
			"Point",
			"PollAddOptionMessage",
			"PollCreationMessage",
			"PollEncValue",
			"PollResultSnapshotMessage",
			"PollUpdateMessage",
			"PollUpdateMessageMetadata",
			"PollVoteMessage",
			"ProcessedVideo",
			"ProductMessage",
			"ProtocolMessage",
			"QuestionResponseMessage",
			"ReactionMessage",
			"RequestPaymentMessage",
			"RequestPhoneNumberMessage",
			"RequestWelcomeMessageMetadata",
			"RootSecretDistributeMessage",
			"ScheduledCallCreationMessage",
			"ScheduledCallEditMessage",
			"SecretEncryptedMessage",
			"SendPaymentMessage",
			"SenderKeyDistributionMessage",
			"SplitPaymentMessage",
			"SplitPaymentParticipant",
			"SplitPaymentUpdateMessage",
			"StatusLinkPreviewMetadata",
			"StatusNotificationMessage",
			"StatusQuestionAnswerMessage",
			"StatusQuotedMessage",
			"StatusStickerInteractionMessage",
			"StickerMessage",
			"StickerPackMessage",
			"StickerSyncRMRMessage",
			"TapLinkAction",
			"TemplateButton",
			"TemplateButtonReplyMessage",
			"TemplateMessage",
			"ThreadID",
			"URLMetadata",
			"UrlTrackingMap",
			"VideoEndCard",
			"VideoMessage",
			"AccountLinkingOpaqueData",
			"ChatRowOpaqueData",
			"DeviceConsistencyCodeMessage",
			"IdentityKeyPairStructure",
			"KeyExchangeMessage",
			"MsgOpaqueData",
			"MsgRowOpaqueData",
			"PreKeyRecordStructure",
			"PreKeySignalMessage",
			"RecordStructure",
			"SenderKeyMessage",
			"SenderKeyRecordStructure",
			"SenderKeyStateStructure",
			"SessionStructure",
			"SignalMessage",
			"SignedPreKeyRecordStructure",
		},
		Enums: []string{
			"HistorySyncType",
			"InsightDeliveryState",
			"KeepType",
			"MediaKeyDomain",
			"PeerDataOperationRequestType",
			"PollContentType",
			"PollType",
			"WebLinkRenderConfig",
		},
		Imports: []string{
			"waAICommon/WAWebProtobufsAICommon.proto",
			"waAICommonDeprecated/WAAICommonDeprecated.proto",
			"waAdv/WAAdv.proto",
			"waAea/WAWebProtobufsAea.proto",
			"waCommon/WACommon.proto",
			"waCompanionReg/WACompanionReg.proto",
			"waMmsRetry/WAMmsRetry.proto",
			"waServerSync/WAWebProtobufsServerSync.proto",
			"waStatusAttributions/WAStatusAttributions.proto",
		},
	},
	"waE2EGuest": {
		Dir: "waE2EGuest",
		File: "WAWebProtobufsE2EGuest.proto",
		Package: "WAWebProtobufsE2EGuest",
		GoPackage: "go.mau.fi/whatsmeow/proto/waE2EGuest",
		Messages: []string{
			"Message",
			"MessageContextInfo",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waEphemeral": {
		Dir: "waEphemeral",
		File: "WAWebProtobufsEphemeral.proto",
		Package: "WAWebProtobufsEphemeral",
		GoPackage: "go.mau.fi/whatsmeow/proto/waEphemeral",
		Messages: []string{
			"EphemeralSetting",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waFingerprint": {
		Dir: "waFingerprint",
		File: "WAFingerprint.proto",
		Package: "WAFingerprint",
		GoPackage: "go.mau.fi/whatsmeow/proto/waFingerprint",
		Messages: []string{
			"CombinedFingerprint",
			"FingerprintData",
		},
		Enums: []string{
			"HostedState",
		},
		Imports: []string{
		},
	},
	"waGroupHistory": {
		Dir: "waGroupHistory",
		File: "WAWebProtobufsGroupHistory.proto",
		Package: "WAWebProtobufsGroupHistory",
		GoPackage: "go.mau.fi/whatsmeow/proto/waGroupHistory",
		Messages: []string{
			"GroupHistory",
			"GroupHistoryWithMessageBytes",
			"UnCountedAssociatedMessageList",
			"UnCountedAssociatedMessageListWithMessageBytes",
			"WebMessageInfoWithMessageBytes",
		},
		Enums: []string{
		},
		Imports: []string{
			"waCommon/WACommon.proto",
			"waE2E/WAWebProtobufsE2E.proto",
			"waWeb/WAWebProtobufsWeb.proto",
		},
	},
	"waHistorySync": {
		Dir: "waHistorySync",
		File: "WAWebProtobufsHistorySync.proto",
		Package: "WAWebProtobufsHistorySync",
		GoPackage: "go.mau.fi/whatsmeow/proto/waHistorySync",
		Messages: []string{
			"Account",
			"AutoDownloadSettings",
			"AvatarUserSettings",
			"Conversation",
			"GlobalSettings",
			"GroupParticipant",
			"HistorySync",
			"HistorySyncMsg",
			"IdentityVerificationState",
			"InlineContact",
			"NotificationSettings",
			"PastParticipant",
			"PastParticipants",
			"PhoneNumberToLIDMapping",
			"Pushname",
			"StickerMetadata",
			"WallpaperSettings",
		},
		Enums: []string{
			"MediaVisibility",
			"PrivacySystemMessage",
		},
		Imports: []string{
			"waChatLockSettings/WAWebProtobufsChatLockSettings.proto",
			"waCommon/WACommon.proto",
			"waE2E/WAWebProtobufsE2E.proto",
			"waSyncAction/WAWebProtobufSyncAction.proto",
			"waWeb/WAWebProtobufsWeb.proto",
		},
	},
	"waLidMigrationSyncPayload": {
		Dir: "waLidMigrationSyncPayload",
		File: "WAWebProtobufLidMigrationSyncPayload.proto",
		Package: "WAWebProtobufLidMigrationSyncPayload",
		GoPackage: "go.mau.fi/whatsmeow/proto/waLidMigrationSyncPayload",
		Messages: []string{
			"LIDMigrationMapping",
			"LIDMigrationMappingSyncPayload",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waMediaEntryData": {
		Dir: "waMediaEntryData",
		File: "WAMediaEntryData.proto",
		Package: "WAMediaEntryData",
		GoPackage: "go.mau.fi/whatsmeow/proto/waMediaEntryData",
		Messages: []string{
			"MediaEntry",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waMediaTransport": {
		Dir: "waMediaTransport",
		File: "WAMediaTransport.proto",
		Package: "WAMediaTransport",
		GoPackage: "go.mau.fi/whatsmeow/proto/waMediaTransport",
		Messages: []string{
			"AudioTransport",
			"ContactTransport",
			"DocumentTransport",
			"ImageTransport",
			"StickerTransport",
			"VideoTransport",
			"WAMediaTransport",
		},
		Enums: []string{
		},
		Imports: []string{
			"waCommon/WACommon.proto",
		},
	},
	"waMmsRetry": {
		Dir: "waMmsRetry",
		File: "WAMmsRetry.proto",
		Package: "WAMmsRetry",
		GoPackage: "go.mau.fi/whatsmeow/proto/waMmsRetry",
		Messages: []string{
			"MediaRetryNotification",
			"ServerErrorReceipt",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waMsgApplication": {
		Dir: "waMsgApplication",
		File: "WAMsgApplication.proto",
		Package: "WAMsgApplication",
		GoPackage: "go.mau.fi/whatsmeow/proto/waMsgApplication",
		Messages: []string{
			"MessageApplication",
		},
		Enums: []string{
		},
		Imports: []string{
			"waCommon/WACommon.proto",
		},
	},
	"waMsgTransport": {
		Dir: "waMsgTransport",
		File: "WAMsgTransport.proto",
		Package: "WAMsgTransport",
		GoPackage: "go.mau.fi/whatsmeow/proto/waMsgTransport",
		Messages: []string{
			"DeviceListMetadata",
			"MessageTransport",
		},
		Enums: []string{
		},
		Imports: []string{
			"waCommon/WACommon.proto",
		},
	},
	"waMultiDevice": {
		Dir: "waMultiDevice",
		File: "WAMultiDevice.proto",
		Package: "WAMultiDevice",
		GoPackage: "go.mau.fi/whatsmeow/proto/waMultiDevice",
		Messages: []string{
			"MultiDevice",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waQuickPromotionSurfaces": {
		Dir: "waQuickPromotionSurfaces",
		File: "WAWebProtobufsQuickPromotionSurfaces.proto",
		Package: "WAWebProtobufsQuickPromotionSurfaces",
		GoPackage: "go.mau.fi/whatsmeow/proto/waQuickPromotionSurfaces",
		Messages: []string{
			"QP",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waReporting": {
		Dir: "waReporting",
		File: "WAWebProtobufsReporting.proto",
		Package: "WAWebProtobufsReporting",
		GoPackage: "go.mau.fi/whatsmeow/proto/waReporting",
		Messages: []string{
			"Config",
			"Field",
			"Reportable",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waRoutingInfo": {
		Dir: "waRoutingInfo",
		File: "WAWebProtobufsRoutingInfo.proto",
		Package: "WAWebProtobufsRoutingInfo",
		GoPackage: "go.mau.fi/whatsmeow/proto/waRoutingInfo",
		Messages: []string{
			"RoutingInfo",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waServerSync": {
		Dir: "waServerSync",
		File: "WAWebProtobufsServerSync.proto",
		Package: "WAWebProtobufsServerSync",
		GoPackage: "go.mau.fi/whatsmeow/proto/waServerSync",
		Messages: []string{
			"CoexStateSync",
			"ExitCode",
			"ExternalBlobReference",
			"KeyId",
			"SyncdIndex",
			"SyncdMutation",
			"SyncdMutations",
			"SyncdPatch",
			"SyncdRecord",
			"SyncdSnapshot",
			"SyncdValue",
			"SyncdVersion",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waStatusAttributions": {
		Dir: "waStatusAttributions",
		File: "WAStatusAttributions.proto",
		Package: "WAStatusAttributions",
		GoPackage: "go.mau.fi/whatsmeow/proto/waStatusAttributions",
		Messages: []string{
			"StatusAttribution",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waSyncAction": {
		Dir: "waSyncAction",
		File: "WAWebProtobufSyncAction.proto",
		Package: "WAWebProtobufSyncAction",
		GoPackage: "go.mau.fi/whatsmeow/proto/waSyncAction",
		Messages: []string{
			"AgentAction",
			"AiThreadRenameAction",
			"AndroidUnsupportedActions",
			"ArchiveChatAction",
			"AutoOrganizeBusinessChatSetting",
			"AvatarUpdatedAction",
			"BizAISettingsNudgeAction",
			"BotWelcomeRequestAction",
			"BroadcastListParticipant",
			"BubbleLockMessageAction",
			"BusinessBroadcastAssociationAction",
			"BusinessBroadcastCampaignAction",
			"BusinessBroadcastInsightsAction",
			"BusinessBroadcastListAction",
			"CallLogAction",
			"CallLogRecord",
			"ChatAssignmentAction",
			"ChatAssignmentOpenedStatusAction",
			"ClearChatAction",
			"CoexV2VersionAction",
			"ContactAction",
			"CtwaPerCustomerDataSharingAction",
			"CustomPaymentMethod",
			"CustomPaymentMethodMetadata",
			"CustomPaymentMethodsAction",
			"CustomerDataAction",
			"DeleteChatAction",
			"DeleteIndividualCallLogAction",
			"DeleteMessageForMeAction",
			"DetectedOutcomesStatusAction",
			"ExternalWebBetaAction",
			"FavoritesAction",
			"InteractiveMessageAction",
			"KeyExpiration",
			"LabelAssociationAction",
			"LabelEditAction",
			"LabelReorderingAction",
			"LabelSublistAction",
			"LidContactAction",
			"LocaleSetting",
			"LockChatAction",
			"MaibaAIFeaturesControlAction",
			"MarkChatAsReadAction",
			"MarketingMessageAction",
			"MarketingMessageBroadcastAction",
			"MerchantPaymentPartnerAction",
			"MusicUserIdAction",
			"MuteAction",
			"NctSaltSyncAction",
			"NewsletterSavedInterestsAction",
			"NoteEditAction",
			"NotificationActivitySettingAction",
			"NuxAction",
			"OutContactAction",
			"PatchDebugData",
			"PaymentInfoAction",
			"PaymentTosAction",
			"PinAction",
			"PnForLidChatAction",
			"PrimaryFeature",
			"PrimaryVersionAction",
			"PrivacySettingChannelsPersonalisedRecommendationAction",
			"PrivacySettingDisableLinkPreviewsAction",
			"PrivacySettingRelayAllCalls",
			"PrivateProcessingSettingAction",
			"PushNameSetting",
			"QuickReplyAction",
			"RecentEmojiWeight",
			"RecentEmojiWeightsAction",
			"RemoveRecentStickerAction",
			"SettingsSyncAction",
			"StarAction",
			"StatusPostOptInNotificationPreferencesAction",
			"StatusPrivacyAction",
			"StickerAction",
			"SubscriptionAction",
			"SubscriptionsSyncV2Action",
			"SyncActionData",
			"SyncActionMessage",
			"SyncActionMessageRange",
			"SyncActionValue",
			"ThreadPinAction",
			"TimeFormatAction",
			"UGCBot",
			"UnarchiveChatsSetting",
			"UserStatusMuteAction",
			"UsernameChatStartModeAction",
			"WASARootSecretAction",
			"WaffleAccountLinkStateAction",
			"WamoUserIdentifierAction",
		},
		Enums: []string{
			"BusinessBroadcastCampaignStatus",
			"CollectionName",
			"CollectionName",
			"CollectionNameStr",
			"MutationName",
			"MutationProps",
		},
		Imports: []string{
			"waChatLockSettings/WAWebProtobufsChatLockSettings.proto",
			"waCommon/WACommon.proto",
			"waDeviceCapabilities/WAWebProtobufsDeviceCapabilities.proto",
		},
	},
	"waSyncdSnapshotRecovery": {
		Dir: "waSyncdSnapshotRecovery",
		File: "WAWebProtobufsSyncdSnapshotRecovery.proto",
		Package: "WAWebProtobufsSyncdSnapshotRecovery",
		GoPackage: "go.mau.fi/whatsmeow/proto/waSyncdSnapshotRecovery",
		Messages: []string{
			"SyncdPlainTextRecord",
			"SyncdSnapshotRecovery",
			"SyncdVersion",
		},
		Enums: []string{
		},
		Imports: []string{
			"waSyncAction/WAWebProtobufSyncAction.proto",
		},
	},
	"waUserPassword": {
		Dir: "waUserPassword",
		File: "WAWebProtobufsUserPassword.proto",
		Package: "WAWebProtobufsUserPassword",
		GoPackage: "go.mau.fi/whatsmeow/proto/waUserPassword",
		Messages: []string{
			"UserPassword",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waVnameCert": {
		Dir: "waVnameCert",
		File: "WAWebProtobufsVnameCert.proto",
		Package: "WAWebProtobufsVnameCert",
		GoPackage: "go.mau.fi/whatsmeow/proto/waVnameCert",
		Messages: []string{
			"BizAccountLinkInfo",
			"BizAccountPayload",
			"BizIdentityInfo",
			"LocalizedName",
			"VerifiedNameCertificate",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waWa6": {
		Dir: "waWa6",
		File: "WAWebProtobufsWa6.proto",
		Package: "WAWebProtobufsWa6",
		GoPackage: "go.mau.fi/whatsmeow/proto/waWa6",
		Messages: []string{
			"ClientPayload",
			"HandshakeMessage",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waWeb": {
		Dir: "waWeb",
		File: "WAWebProtobufsWeb.proto",
		Package: "WAWebProtobufsWeb",
		GoPackage: "go.mau.fi/whatsmeow/proto/waWeb",
		Messages: []string{
			"Citation",
			"CommentMetadata",
			"EventAdditionalMetadata",
			"EventResponse",
			"GroupHistoryBundleInfo",
			"GroupHistoryIndividualMessageInfo",
			"InteractiveMessageAdditionalMetadata",
			"KeepInChat",
			"LegacyMessage",
			"MediaData",
			"MessageAddOn",
			"MessageAddOnContextInfo",
			"NotificationMessageInfo",
			"PaymentInfo",
			"PhotoChange",
			"PinInChat",
			"PollAdditionalMetadata",
			"PollUpdate",
			"PremiumMessageInfo",
			"QuarantinedMessage",
			"Reaction",
			"ReportingTokenInfo",
			"ScheduledMessageMetadata",
			"StatusMentionMessage",
			"StatusPSA",
			"UserReceipt",
			"WebFeatures",
			"WebMessageInfo",
			"WebNotificationsInfo",
		},
		Enums: []string{
		},
		Imports: []string{
			"waCommon/WACommon.proto",
			"waE2E/WAWebProtobufsE2E.proto",
		},
	},
	"waWebLabyrinthWaWasm": {
		Dir: "waWebLabyrinthWaWasm",
		File: "WAWebLabyrinthWaWasm.proto",
		Package: "WAWebLabyrinthWaWasm",
		GoPackage: "go.mau.fi/whatsmeow/proto/waWebLabyrinthWaWasm",
		Messages: []string{
			"CreateBackupInput",
			"CreateBackupOutput",
			"DecryptMessageInput",
			"DecryptMessageOutput",
			"DeriveMessageKeyInput",
			"DeriveMessageKeyOutput",
			"DeviceOutput",
			"EncryptMessageInput",
			"EncryptMessageOutput",
			"EncryptedSecretValuesOutput",
			"Epoch0Output",
			"LabyrinthWaCommand",
			"OrfThreadIdInput",
			"OrfThreadIdOutput",
			"RotateEpochInput",
			"RotateEpochMemberEdge",
			"RotateEpochMemberInput",
			"RotateEpochOutput",
			"VirtualDeviceOutput",
		},
		Enums: []string{
		},
		Imports: []string{
		},
	},
	"waWinUIApi": {
		Dir: "waWinUIApi",
		File: "WAWinUIApi.proto",
		Package: "WAWinUIApi",
		GoPackage: "go.mau.fi/whatsmeow/proto/waWinUIApi",
		Messages: []string{
			"PositronChat",
			"PositronContact",
			"PositronData",
			"PositronGroupMetadata",
			"PositronGroupParticipants",
			"PositronMessage",
			"PositronReaction",
		},
		Enums: []string{
			"PositronDataSource",
		},
		Imports: []string{
		},
	},
}

var TypeToPackage = make(map[string]string)

func init() {
	// 1. Initial pass in deterministic order
	for dir, pkg := range Packages {
		for _, m := range pkg.Messages {
			TypeToPackage[m] = dir
		}
		for _, e := range pkg.Enums {
			TypeToPackage[e] = dir
		}
	}

	// 2. High priority core packages override parameterised variants/stubs
	primaryOverrides := []string{
		"waCommon",
		"waConsumerApplication",
		"waServerSync",
		"waAICommon",
		"waE2E",
	}
	for _, dir := range primaryOverrides {
		if pkg, ok := Packages[dir]; ok {
			for _, m := range pkg.Messages {
				TypeToPackage[m] = dir
			}
			for _, e := range pkg.Enums {
				TypeToPackage[e] = dir
			}
		}
	}
}

// IsScalarType returns true if the type is a primitive protobuf scalar.
func IsScalarType(t string) bool {
	switch strings.ToLower(t) {
	case "double", "float", "int32", "int64", "uint32", "uint64",
		"sint32", "sint64", "fixed32", "fixed64", "sfixed32", "sfixed64",
		"bool", "string", "bytes":
		return true
	default:
		return false
	}
}

// LookupPackageForType finds the package directory for a given type name.
func LookupPackageForType(typeName string) string {
	clean := strings.TrimPrefix(typeName, ".")
	if IsScalarType(clean) {
		return ""
	}
	if clean == "BackwardEdge" {
		return "waWebLabyrinthWaWasm"
	}
	if pkg, ok := TypeToPackage[clean]; ok {
		return pkg
	}
	if strings.HasPrefix(typeName, "Instamadillo") {
		return "instamadilloAddMessage"
	}
	if strings.HasPrefix(typeName, "ADV") {
		return "waAdv"
	}
	if strings.HasPrefix(typeName, "Bot") || strings.HasPrefix(typeName, "AI") || strings.HasPrefix(typeName, "BizAI") {
		return "waAICommon"
	}
	if strings.HasPrefix(typeName, "HistorySync") {
		return "waHistorySync"
	}
	if strings.HasPrefix(typeName, "SyncAction") {
		return "waSyncAction"
	}
	return "waE2E"
}

