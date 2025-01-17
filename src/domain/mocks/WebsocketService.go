// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "Refractor/domain"
	broadcast "Refractor/pkg/broadcast"

	mock "github.com/stretchr/testify/mock"

	net "net"
)

// WebsocketService is an autogenerated mock type for the WebsocketService type
type WebsocketService struct {
	mock.Mock
}

// Broadcast provides a mock function with given fields: message
func (_m *WebsocketService) Broadcast(message *domain.WebsocketMessage) {
	_m.Called(message)
}

// BroadcastServerMessage provides a mock function with given fields: message, serverID, authChecker
func (_m *WebsocketService) BroadcastServerMessage(message *domain.WebsocketMessage, serverID int64, authChecker domain.AuthChecker) error {
	ret := _m.Called(message, serverID, authChecker)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.WebsocketMessage, int64, domain.AuthChecker) error); ok {
		r0 = rf(message, serverID, authChecker)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateClient provides a mock function with given fields: userID, conn
func (_m *WebsocketService) CreateClient(userID string, conn net.Conn) {
	_m.Called(userID, conn)
}

// HandlePlayerJoin provides a mock function with given fields: fields, serverID, game
func (_m *WebsocketService) HandlePlayerJoin(fields broadcast.Fields, serverID int64, game domain.Game) {
	_m.Called(fields, serverID, game)
}

// HandlePlayerListUpdate provides a mock function with given fields: serverID, players, game
func (_m *WebsocketService) HandlePlayerListUpdate(serverID int64, players []*domain.OnlinePlayer, game domain.Game) {
	_m.Called(serverID, players, game)
}

// HandlePlayerQuit provides a mock function with given fields: fields, serverID, game
func (_m *WebsocketService) HandlePlayerQuit(fields broadcast.Fields, serverID int64, game domain.Game) {
	_m.Called(fields, serverID, game)
}

// HandleServerStatusChange provides a mock function with given fields: serverID, status
func (_m *WebsocketService) HandleServerStatusChange(serverID int64, status string) {
	_m.Called(serverID, status)
}

// SendDirectMessage provides a mock function with given fields: message, userID
func (_m *WebsocketService) SendDirectMessage(message *domain.WebsocketMessage, userID string) {
	_m.Called(message, userID)
}

// StartPool provides a mock function with given fields:
func (_m *WebsocketService) StartPool() {
	_m.Called()
}

// SubscribeChatSend provides a mock function with given fields: sub
func (_m *WebsocketService) SubscribeChatSend(sub domain.ChatSendSubscriber) {
	_m.Called(sub)
}
