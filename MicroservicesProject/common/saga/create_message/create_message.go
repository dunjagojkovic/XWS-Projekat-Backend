package create_message

import (
	"time"
)

type MessageDetails struct {
	Id       string
	Text     string
	Sender   string
	Receiver string
	Time     time.Time
	Status   string
	ChatId   string
}

type CreateMessageCommandType int8

const (
	CheckBlocking CreateMessageCommandType = iota
	CancelMessage
	ApproveMessage
	UnknownCommand
)

type CreateMessageCommand struct {
	Message MessageDetails
	Type    CreateMessageCommandType
}

type CreateMessageReplyType int8

const (
	UserBlocked CreateMessageReplyType = iota
	UserNotBlocked
	MessageCancelled
	MessageApproved
	UnknownReply
)

type CreateMessageReply struct {
	Message MessageDetails
	Type    CreateMessageReplyType
}
