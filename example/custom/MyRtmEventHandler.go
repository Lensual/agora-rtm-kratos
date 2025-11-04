package main

import (
	agrtm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/go_sdk/rtm"
	"github.com/go-kratos/kratos/v2/log"
)

func NewMyRtmEventHandler() *agrtm.RtmEventHandler {
	return &agrtm.RtmEventHandler{
		OnMessageEvent: func(event *agrtm.MessageEvent) {
			log.Infow(
				"callback", "OnMessageEvent",
				"event", event,
			)
		},
		OnPresenceEvent: func(event *agrtm.PresenceEvent) {
			log.Infow(
				"callback", "OnPresenceEvent",
				"event", event,
			)
		},
		OnTopicEvent: func(event *agrtm.TopicEvent) {
			log.Infow(
				"callback", "OnTopicEvent",
				"event", event,
			)
		},
		OnLockEvent: func(event *agrtm.LockEvent) {
			log.Infow(
				"callback", "OnLockEvent",
				"event", event,
			)
		},
		OnStorageEvent: func(event *agrtm.StorageEvent) {
			log.Infow(
				"callback", "OnStorageEvent",
				"event", event,
			)
		},
		OnJoinResult: func(requestId uint64, channelName string, userId string, errorCode int) {
			log.Infow(
				"callback", "OnJoinResult",
				"requestId", requestId,
				"channelName", channelName,
				"userId", userId,
				"errorCode", errorCode,
			)
		},
		OnLeaveResult: func(requestId uint64, channelName string, userId string, errorCode int) {
			log.Infow(
				"callback", "OnLeaveResult",
				"requestId", requestId,
				"channelName", channelName,
				"userId", userId,
				"errorCode", errorCode,
			)
		},
		OnJoinTopicResult: func(requestId uint64, channelName string, userId string, topic string, meta string, errorCode int) {
			log.Infow(
				"callback", "OnJoinTopicResult",
				"requestId", requestId,
				"channelName", channelName,
				"userId", userId,
				"topic", topic,
				"meta", meta,
				"errorCode", errorCode,
			)
		},
		OnLeaveTopicResult: func(requestId uint64, channelName string, userId string, topic string, meta string, errorCode int) {
			log.Infow(
				"callback", "OnLeaveTopicResult",
				"requestId", requestId,
				"channelName", channelName,
				"userId", userId,
				"topic", topic,
				"meta", meta,
				"errorCode", errorCode,
			)
		},
		OnSubscribeTopicResult: func(requestId uint64, channelName string, userId string, topic string, succeedUsers agrtm.UserList, failedUsers agrtm.UserList, errorCode int) {
			log.Infow(
				"callback", "OnSubscribeTopicResult",
				"requestId", requestId,
				"channelName", channelName,
				"userId", userId,
				"topic", topic,
				"succeedUsers", succeedUsers,
				"failedUsers", failedUsers,
				"errorCode", errorCode,
			)
		},
		OnConnectionStateChanged: func(channelName string, state int, reason int) {
			log.Infow(
				"callback", "OnConnectionStateChanged",
				"channelName", channelName,
				"state", state,
				"reason", reason,
			)
		},
		OnTokenPrivilegeWillExpire: func(channelName string) {
			log.Infow(
				"callback", "OnTokenPrivilegeWillExpire",
				"channelName", channelName,
			)
		},
		OnSubscribeResult: func(requestId uint64, channelName string, errorCode int) {
			log.Infow(
				"callback", "OnSubscribeResult",
				"requestId", requestId,
				"channelName", channelName,
				"errorCode", errorCode,
			)
		},
		OnPublishResult: func(requestId uint64, errorCode int) {
			log.Infow(
				"callback", "OnPublishResult",
				"requestId", requestId,
				"errorCode", errorCode,
			)
		},
		OnLoginResult: func(requestId uint64, errorCode int) {
			log.Infow(
				"callback", "OnLoginResult",
				"requestId", requestId,
				"errorCode", errorCode,
			)
		},
		OnSetChannelMetadataResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, errorCode int) {
			log.Infow(
				"callback", "OnSetChannelMetadataResult",
				"requestId", requestId,
				"channelName", channelName,
				"channelType", channelType,
				"errorCode", errorCode,
			)
		},
		OnUpdateChannelMetadataResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, errorCode int) {
			log.Infow(
				"callback", "OnUpdateChannelMetadataResult",
				"requestId", requestId,
				"channelName", channelName,
				"channelType", channelType,
				"errorCode", errorCode,
			)
		},
		OnRemoveChannelMetadataResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, errorCode int) {
			log.Infow(
				"callback", "OnRemoveChannelMetadataResult",
				"requestId", requestId,
				"channelName", channelName,
				"channelType", channelType,
				"errorCode", errorCode,
			)
		},
		OnGetChannelMetadataResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, data *agrtm.IMetadata, errorCode int) {
			log.Infow(
				"callback", "OnGetChannelMetadataResult",
				"requestId", requestId,
				"channelName", channelName,
				"channelType", channelType,
				"data", data,
				"errorCode", errorCode,
			)
		},
		OnSetUserMetadataResult: func(requestId uint64, userId string, errorCode int) {
			log.Infow(
				"callback", "OnSetUserMetadataResult",
				"requestId", requestId,
				"userId", userId,
				"errorCode", errorCode,
			)
		},
		OnUpdateUserMetadataResult: func(requestId uint64, userId string, errorCode int) {
			log.Infow(
				"callback", "OnUpdateUserMetadataResult",
				"requestId", requestId,
				"userId", userId,
				"errorCode", errorCode,
			)
		},
		OnRemoveUserMetadataResult: func(requestId uint64, userId string, errorCode int) {
			log.Infow(
				"callback", "OnRemoveUserMetadataResult",
				"requestId", requestId,
				"userId", userId,
				"errorCode", errorCode,
			)
		},
		OnGetUserMetadataResult: func(requestId uint64, userId string, data *agrtm.IMetadata, errorCode int) {
			log.Infow(
				"callback", "OnGetUserMetadataResult",
				"requestId", requestId,
				"userId", userId,
				"data", data,
				"errorCode", errorCode,
			)
		},
		OnSubscribeUserMetadataResult: func(requestId uint64, userId string, errorCode int) {
			log.Infow(
				"callback", "OnSubscribeUserMetadataResult",
				"requestId", requestId,
				"userId", userId,
				"errorCode", errorCode,
			)
		},
		OnSetLockResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int) {
			log.Infow(
				"callback", "OnSetLockResult",
				"requestId", requestId,
				"channelName", channelName,
				"channelType", channelType,
				"lockName", lockName,
				"errorCode", errorCode,
			)
		},
		OnRemoveLockResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int) {
			log.Infow(
				"callback", "OnRemoveLockResult",
				"requestId", requestId,
				"channelName", channelName,
				"channelType", channelType,
				"lockName", lockName,
				"errorCode", errorCode,
			)
		},
		OnReleaseLockResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int) {
			log.Infow(
				"callback", "OnReleaseLockResult",
				"requestId", requestId,
				"channelName", channelName,
				"channelType", channelType,
				"lockName", lockName,
				"errorCode", errorCode,
			)
		},
		OnAcquireLockResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int, errorDetails string) {
			log.Infow(
				"callback", "OnAcquireLockResult",
				"requestId", requestId,
				"channelName", channelName,
				"channelType", channelType,
				"lockName", lockName,
				"errorCode", errorCode,
				"errorDetails", errorDetails,
			)
		},
		OnRevokeLockResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int) {
			log.Infow(
				"callback", "OnRevokeLockResult",
				"requestId", requestId,
				"channelName", channelName,
				"channelType", channelType,
				"lockName", lockName,
				"errorCode", errorCode,
			)
		},
		OnGetLocksResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockDetailList *agrtm.LockDetail, count uint, errorCode int) {
			log.Infow(
				"callback", "OnGetLocksResult",
				"requestId", requestId,
				"channelName", channelName,
				"channelType", channelType,
				"lockDetailList", lockDetailList,
				"count", count,
				"errorCode", errorCode,
			)
		},
		OnWhoNowResult: func(requestId uint64, userStateList *agrtm.UserState, count uint, nextPage string, errorCode int) {
			log.Infow(
				"callback", "OnWhoNowResult",
				"requestId", requestId,
				"userStateList", userStateList,
				"count", count,
				"nextPage", nextPage,
				"errorCode", errorCode,
			)
		},
		OnGetOnlineUsersResult: func(requestId uint64, userStateList *agrtm.UserState, count uint, nextPage string, errorCode int) {
			log.Infow(
				"callback", "OnGetOnlineUsersResult",
				"requestId", requestId,
				"userStateList", userStateList,
				"count", count,
				"nextPage", nextPage,
				"errorCode", errorCode,
			)
		},
		OnWhereNowResult: func(requestId uint64, channels *agrtm.ChannelInfo, count uint, errorCode int) {
			log.Infow(
				"callback", "OnWhereNowResult",
				"requestId", requestId,
				"channels", channels,
				"count", count,
				"errorCode", errorCode,
			)
		},
		OnGetUserChannelsResult: func(requestId uint64, channels *agrtm.ChannelInfo, count uint, errorCode int) {
			log.Infow(
				"callback", "OnGetUserChannelsResult",
				"requestId", requestId,
				"channels", channels,
				"count", count,
				"errorCode", errorCode,
			)
		},
		OnPresenceSetStateResult: func(requestId uint64, errorCode int) {
			log.Infow(
				"callback", "OnPresenceSetStateResult",
				"requestId", requestId,
				"errorCode", errorCode,
			)
		},
		OnPresenceRemoveStateResult: func(requestId uint64, errorCode int) {
			log.Infow(
				"callback", "OnPresenceRemoveStateResult",
				"requestId", requestId,
				"errorCode", errorCode,
			)
		},
		OnPresenceGetStateResult: func(requestId uint64, state *agrtm.UserState, errorCode int) {
			log.Infow(
				"callback", "OnPresenceGetStateResult",
				"requestId", requestId,
				"state", state,
				"errorCode", errorCode,
			)
		},
		OnLinkStateEvent: func(event *agrtm.LinkStateEvent) {
			log.Infow(
				"callback", "OnLinkStateEvent",
				"event", event,
			)
		},
		OnLogoutResult: func(requestId uint64, errorCode int) {
			log.Infow(
				"callback", "OnLogoutResult",
				"requestId", requestId,
				"errorCode", errorCode,
			)
		},
		OnRenewTokenResult: func(requestId uint64, serverType agrtm.RtmServiceType, channelName string, errorCode int) {
			log.Infow(
				"callback", "OnRenewTokenResult",
				"requestId", requestId,
				"serverType", serverType,
				"channelName", channelName,
				"errorCode", errorCode,
			)
		},
		OnPublishTopicMessageResult: func(requestId uint64, channelName string, topic string, errorCode int) {
			log.Infow(
				"callback", "OnPublishTopicMessageResult",
				"requestId", requestId,
				"channelName", channelName,
				"topic", topic,
				"errorCode", errorCode,
			)
		},
		OnUnsubscribeTopicResult: func(requestId uint64, channelName string, topic string, errorCode int) {
			log.Infow(
				"callback", "OnUnsubscribeTopicResult",
				"requestId", requestId,
				"channelName", channelName,
				"topic", topic,
				"errorCode", errorCode,
			)
		},
		OnGetSubscribedUserListResult: func(requestId uint64, channelName string, topic string, user *agrtm.UserList, errorCode int) {
			log.Infow(
				"callback", "OnGetSubscribedUserListResult",
				"requestId", requestId,
				"channelName", channelName,
				"topic", topic,
				"user", user,
				"errorCode", errorCode,
			)
		},
		OnGetHistoryMessagesResult: func(requestId uint64, messageList []agrtm.HistoryMessage, newStart uint64, errorCode int) {
			log.Infow(
				"callback", "OnGetHistoryMessagesResult",
				"requestId", requestId,
				"messageList", messageList,
				"newStart", newStart,
				"errorCode", errorCode,
			)
		},
		OnUnsubscribeUserMetadataResult: func(requestId uint64, userId string, errorCode int) {
			log.Infow(
				"callback", "OnUnsubscribeUserMetadataResult",
				"requestId", requestId,
				"userId", userId,
				"errorCode", errorCode,
			)
		},
	}
}
