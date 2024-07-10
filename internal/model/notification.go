package model

import (
	"awesomeProject/enums"
	"github.com/google/uuid"
	"time"
)

// trigger > sync > publishing the notification
// trigger > async >

// notification > all the changes > 500
// notification >
// 180 days old > list > 10 limit > deleted_at

type Notification struct {
	Id        uuid.UUID     `json:"id"`
	Type      enums.Type    `json:"type"`
	Coin      enums.Coin    `json:"coin"`
	Channel   enums.Channel `json:"channel"`
	Data      interface{}   `json:"data"`
	CreatedAt time.Time     `json:"created_at"`
	DeletedAt time.Time     `json:"deleted_at"`
}
