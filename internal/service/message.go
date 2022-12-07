package service

import "gof/internal/pb"

type (
	sMessage struct{}
)

var (
	insMessage = sMessage{}
)

func Message() *sMessage {
	return &insMessage
}

func (s *sMessage) BuildMsg(sessionId, action string) {
	msg := pb.Msg{}
	msg.Action = action
	msg.SessionId = sessionId
	msg.Code = 0
	msg.Sequence = 0
}
