package services

import (
	"awesomeProject/internal/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification interface {
	Create(notification Notification) error
	List(page int, per_page int) []Notification
	Delete(id uuid.UUID) error
}

type NotificationService struct {
	db *gorm.DB
}

func NewNotificationService() Notification {
	return &NotificationService{
		db: db.Get(),
	}
}

func (NotificationService) Create(notification Notification) error {
	//TODO implement me
	panic("implement me")
}

func (NotificationService) List(page int, per_page int) []Notification {
	//TODO implement me
	panic("implement me")
}

func (NotificationService) Delete(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
