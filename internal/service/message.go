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

func (s *sMessage) buildPBMsg(action pb.ACTION, code pb.CODE) *pb.Msg {
	return &pb.Msg{
		Sequence: 0,
		Action:   action,
		Code:     code,
	}
}
