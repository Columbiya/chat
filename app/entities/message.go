package entities

import (
	"time"
)

type IncomingMessage struct {
	Text     string `json:"text"`
	ToUserId Guid   `json:"toUserId"`
}

type DatabaseMessage struct {
	Text       string    `json:"text"`
	ToUserId   Guid      `json:"toUserId"`
	FromUserId Guid      `json:"fromUserId"`
	DateSent   time.Time `json:"dateSent"`
}

type Message struct {
	Payload    string
	DateSent   time.Time
	FromUserId Guid
}

func NewMessage(payload *string, fromUserId string) *Message {
	return &Message{Payload: *payload, FromUserId: fromUserId, DateSent: time.Now()}
}

func NewDatabaseMessage(payload *string, fromUserId *string, toUserId *string) *DatabaseMessage {
	return &DatabaseMessage{Text: *payload, FromUserId: *fromUserId, ToUserId: *toUserId, DateSent: time.Now()}
}
