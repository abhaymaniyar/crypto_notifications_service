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
	Id        uuid.UUID     `gorm:"type:uuid;primary_key" json:"id"`
	Type      enums.Type    `gorm:"type:varchar(10);not null" json:"type"`
	Coin      enums.Coin    `gorm:"coin:varchar(10);not null" json:"coin"`
	Channel   enums.Channel `gorm:"channel:varchar(10);not null" json:"channel"`
	Data      interface{}   `gorm:"type:text" json:"description"`
	CreatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}
