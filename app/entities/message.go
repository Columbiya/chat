package entities

import (
	"time"
)

type Message struct {
	Text     string
	DateSent time.Time
	room *Room
}