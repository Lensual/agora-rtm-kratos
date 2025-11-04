package utils

import (
	"log"

	agrtm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/go_sdk/rtm"
)

// NewLoggerHandler 创建一个日志记录 Handler.
//
// logf 是一个日志记录函数，可自定义输出方式。
// 如果 logf 为 nil，则默认使用 [log.Printf] 函数作为日志输出。
//
// Example:
//
//	rtmConfig := agrtm.NewRtmConfig()
//	rtmConfig.EventHandler = utils.NewCompositeRtmEventHandler(
//		utils.NewLoggerHandler(nil),
//		// ...
//	)
func NewLoggerHandler(logf func(format string, a ...any)) *agrtm.RtmEventHandler {
	if logf == nil {
		logf = log.Printf
	}
	return &agrtm.RtmEventHandler{
		OnMessageEvent: func(event *agrtm.MessageEvent) {
			logf("OnMessageEvent: %+v", event)
		},
		OnPresenceEvent: func(event *agrtm.PresenceEvent) {
			logf("OnPresenceEvent: %+v", event)
		},
		OnTopicEvent: func(event *agrtm.TopicEvent) {
			logf("OnTopicEvent: %+v", event)
		},
		OnLockEvent: func(event *agrtm.LockEvent) {
			logf("OnLockEvent: %+v", event)
		},
		OnStorageEvent: func(event *agrtm.StorageEvent) {
			logf("OnStorageEvent: %+v", event)
		},
		OnJoinResult: func(requestId uint64, channelName string, userId string, errorCode int) {
			logf("OnJoinResult: requestId=%d, channel=%s, user=%s, error=%d", requestId, channelName, userId, errorCode)
		},
		OnLeaveResult: func(requestId uint64, channelName string, userId string, errorCode int) {
			logf("OnLeaveResult: requestId=%d, channel=%s, user=%s, error=%d", requestId, channelName, userId, errorCode)
		},
		OnJoinTopicResult: func(requestId uint64, channelName string, userId string, topic string, meta string, errorCode int) {
			logf("OnJoinTopicResult: requestId=%d, channel=%s, user=%s, topic=%s, meta=%s, error=%d",
				requestId, channelName, userId, topic, meta, errorCode)
		},
		OnLeaveTopicResult: func(requestId uint64, channelName string, userId string, topic string, meta string, errorCode int) {
			logf("OnLeaveTopicResult: requestId=%d, channel=%s, user=%s, topic=%s, meta=%s, error=%d",
				requestId, channelName, userId, topic, meta, errorCode)
		},
		OnSubscribeTopicResult: func(requestId uint64, channelName string, userId string, topic string, succeedUsers agrtm.UserList, failedUsers agrtm.UserList, errorCode int) {
			logf("OnSubscribeTopicResult: requestId=%d, channel=%s, user=%s, topic=%s, succeed=%+v, failed=%+v, error=%d",
				requestId, channelName, userId, topic, succeedUsers, failedUsers, errorCode)
		},
		// Deprecated: use OnLinkStateEvent instead.
		OnConnectionStateChanged: func(channelName string, state int, reason int) {
			logf("OnConnectionStateChanged: channel=%s, state=%d, reason=%d", channelName, state, reason)
		},
		OnTokenPrivilegeWillExpire: func(channelName string) {
			logf("OnTokenPrivilegeWillExpire: channel=%s", channelName)
		},
		OnSubscribeResult: func(requestId uint64, channelName string, errorCode int) {
			logf("OnSubscribeResult: requestId=%d, channel=%s, error=%d", requestId, channelName, errorCode)
		},
		OnPublishResult: func(requestId uint64, errorCode int) {
			logf("OnPublishResult: requestId=%d, error=%d", requestId, errorCode)
		},
		OnLoginResult: func(requestId uint64, errorCode int) {
			logf("OnLoginResult: requestId=%d, error=%d", requestId, errorCode)
		},
		OnSetChannelMetadataResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, errorCode int) {
			logf("OnSetChannelMetadataResult: requestId=%d, channel=%s, type=%d, error=%d", requestId, channelName, channelType, errorCode)
		},
		OnUpdateChannelMetadataResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, errorCode int) {
			logf("OnUpdateChannelMetadataResult: requestId=%d, channel=%s, type=%d, error=%d", requestId, channelName, channelType, errorCode)
		},
		OnRemoveChannelMetadataResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, errorCode int) {
			logf("OnRemoveChannelMetadataResult: requestId=%d, channel=%s, type=%d, error=%d", requestId, channelName, channelType, errorCode)
		},
		OnGetChannelMetadataResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, data *agrtm.IMetadata, errorCode int) {
			logf("OnGetChannelMetadataResult: requestId=%d, channel=%s, type=%d, data=%+v, error=%d", requestId, channelName, channelType, data, errorCode)
		},
		OnSetUserMetadataResult: func(requestId uint64, userId string, errorCode int) {
			logf("OnSetUserMetadataResult: requestId=%d, user=%s, error=%d", requestId, userId, errorCode)
		},
		OnUpdateUserMetadataResult: func(requestId uint64, userId string, errorCode int) {
			logf("OnUpdateUserMetadataResult: requestId=%d, user=%s, error=%d", requestId, userId, errorCode)
		},
		OnRemoveUserMetadataResult: func(requestId uint64, userId string, errorCode int) {
			logf("OnRemoveUserMetadataResult: requestId=%d, user=%s, error=%d", requestId, userId, errorCode)
		},
		OnGetUserMetadataResult: func(requestId uint64, userId string, data *agrtm.IMetadata, errorCode int) {
			logf("OnGetUserMetadataResult: requestId=%d, user=%s, data=%+v, error=%d", requestId, userId, data, errorCode)
		},
		OnSubscribeUserMetadataResult: func(requestId uint64, userId string, errorCode int) {
			logf("OnSubscribeUserMetadataResult: requestId=%d, user=%s, error=%d", requestId, userId, errorCode)
		},
		OnSetLockResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int) {
			logf("OnSetLockResult: requestId=%d, channel=%s, type=%d, lock=%s, error=%d", requestId, channelName, channelType, lockName, errorCode)
		},
		OnRemoveLockResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int) {
			logf("OnRemoveLockResult: requestId=%d, channel=%s, type=%d, lock=%s, error=%d", requestId, channelName, channelType, lockName, errorCode)
		},
		OnReleaseLockResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int) {
			logf("OnReleaseLockResult: requestId=%d, channel=%s, type=%d, lock=%s, error=%d", requestId, channelName, channelType, lockName, errorCode)
		},
		OnAcquireLockResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int, errorDetails string) {
			logf("OnAcquireLockResult: requestId=%d, channel=%s, type=%d, lock=%s, error=%d, details=%s", requestId, channelName, channelType, lockName, errorCode, errorDetails)
		},
		OnRevokeLockResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int) {
			logf("OnRevokeLockResult: requestId=%d, channel=%s, type=%d, lock=%s, error=%d", requestId, channelName, channelType, lockName, errorCode)
		},
		OnGetLocksResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockDetailList *agrtm.LockDetail, count uint, errorCode int) {
			logf("OnGetLocksResult: requestId=%d, channel=%s, type=%d, count=%d, locks=%+v, error=%d", requestId, channelName, channelType, count, lockDetailList, errorCode)
		},
		OnWhoNowResult: func(requestId uint64, userStateList *agrtm.UserState, count uint, nextPage string, errorCode int) {
			logf("OnWhoNowResult: requestId=%d, count=%d, next=%s, users=%+v, error=%d", requestId, count, nextPage, userStateList, errorCode)
		},
		OnGetOnlineUsersResult: func(requestId uint64, userStateList *agrtm.UserState, count uint, nextPage string, errorCode int) {
			logf("OnGetOnlineUsersResult: requestId=%d, count=%d, next=%s, users=%+v, error=%d", requestId, count, nextPage, userStateList, errorCode)
		},
		OnWhereNowResult: func(requestId uint64, channels *agrtm.ChannelInfo, count uint, errorCode int) {
			logf("OnWhereNowResult: requestId=%d, count=%d, channels=%+v, error=%d", requestId, count, channels, errorCode)
		},
		OnGetUserChannelsResult: func(requestId uint64, channels *agrtm.ChannelInfo, count uint, errorCode int) {
			logf("OnGetUserChannelsResult: requestId=%d, count=%d, channels=%+v, error=%d", requestId, count, channels, errorCode)
		},
		OnPresenceSetStateResult: func(requestId uint64, errorCode int) {
			logf("OnPresenceSetStateResult: requestId=%d, error=%d", requestId, errorCode)
		},
		OnPresenceRemoveStateResult: func(requestId uint64, errorCode int) {
			logf("OnPresenceRemoveStateResult: requestId=%d, error=%d", requestId, errorCode)
		},
		OnPresenceGetStateResult: func(requestId uint64, state *agrtm.UserState, errorCode int) {
			logf("OnPresenceGetStateResult: requestId=%d, state=%+v, error=%d", requestId, state, errorCode)
		},
		OnLinkStateEvent: func(event *agrtm.LinkStateEvent) {
			logf("OnLinkStateEvent: %+v", event)
		},
		OnLogoutResult: func(requestId uint64, errorCode int) {
			logf("OnLogoutResult: requestId=%d, error=%d", requestId, errorCode)
		},
		OnRenewTokenResult: func(requestId uint64, serverType agrtm.RtmServiceType, channelName string, errorCode int) {
			logf("OnRenewTokenResult: requestId=%d, serviceType=%d, channel=%s, error=%d", requestId, serverType, channelName, errorCode)
		},
		OnPublishTopicMessageResult: func(requestId uint64, channelName string, topic string, errorCode int) {
			logf("OnPublishTopicMessageResult: requestId=%d, channel=%s, topic=%s, error=%d", requestId, channelName, topic, errorCode)
		},
		OnUnsubscribeTopicResult: func(requestId uint64, channelName string, topic string, errorCode int) {
			logf("OnUnsubscribeTopicResult: requestId=%d, channel=%s, topic=%s, error=%d", requestId, channelName, topic, errorCode)
		},
		OnGetSubscribedUserListResult: func(requestId uint64, channelName string, topic string, user *agrtm.UserList, errorCode int) {
			logf("OnGetSubscribedUserListResult: requestId=%d, channel=%s, topic=%s, users=%+v, error=%d", requestId, channelName, topic, user, errorCode)
		},
		OnGetHistoryMessagesResult: func(requestId uint64, messageList []agrtm.HistoryMessage, newStart uint64, errorCode int) {
			logf("OnGetHistoryMessagesResult: requestId=%d, newStart=%d, messages=%+v, error=%d", requestId, newStart, messageList, errorCode)
		},
		OnUnsubscribeUserMetadataResult: func(requestId uint64, userId string, errorCode int) {
			logf("OnUnsubscribeUserMetadataResult: requestId=%d, user=%s, error=%d", requestId, userId, errorCode)
		},
	}
}
