package services

import (
	"awesomeProject/enums"
	"awesomeProject/internal/logger"
	"awesomeProject/internal/model"
	"awesomeProject/repositories"
	"github.com/google/uuid"
	"time"
)

type NotificationRequest struct {
	Type     enums.Type      `json:"type" binding:"required"`
	Coin     enums.Coin      `json:"coin" binding:"required"`
	Channels []enums.Channel `json:"channels" binding:"required"`
	Data     interface{}     `json:"data" binding:"required"`
}

type NotificationServiceI interface {
	Create(notificationReq *NotificationRequest) error
	List(opts repositories.GetNotificationsOpts) ([]model.Notification, error)
	Delete(id uuid.UUID) error
}

type NotificationService struct {
	repo repositories.NotificationRepoI
}

func NewNotificationService() NotificationServiceI {
	return &NotificationService{
		repo: repositories.NewNotificationRepo(),
	}
}

func (s *NotificationService) Create(notificationReq *NotificationRequest) error {
	for _, channel := range notificationReq.Channels {
		notification := model.Notification{
			Id:        uuid.New(),
			Type:      notificationReq.Type,
			Coin:      notificationReq.Coin,
			Channel:   channel,
			Data:      notificationReq.Data,
			CreatedAt: time.Now(),
		}

		err := s.repo.SaveNotification(notification)
		if err != nil {
			logger.W(nil, "error while saving notification", logger.Field("channel", channel))
		}
	}

	return nil
}

func (s *NotificationService) List(opts repositories.GetNotificationsOpts) ([]model.Notification, error) {
	notifications, err := s.repo.GetNotifications(opts)

	if err != nil {
		logger.E(nil, err, "Error while fetching notifications", logger.Field("opts", opts))
		return nil, err
	}

	return notifications, nil
}

func (s *NotificationService) Delete(id uuid.UUID) error {
	notification, err := s.repo.GetNotificationById(id)
	if err != nil {
		logger.W(nil, "Error while fetching notification to update",
			logger.Field("err", err),
			logger.Field("id", id))
		return err
	}

	notification.DeletedAt = time.Now()
	err = s.repo.UpdateNotification(notification)

	if err != nil {
		logger.W(nil, "Error while updating card details",
			logger.Field("err", err),
			logger.Field("id", id))
		return err
	}

	return nil
}
