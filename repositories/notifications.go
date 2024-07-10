package repositories

import (
	"awesomeProject/internal/db"
	"awesomeProject/internal/logger"
	"awesomeProject/internal/model"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GetNotificationsOpts struct {
	Limit  int
	Offset int
}

type NotificationRepoI interface {
	SaveNotification(notification model.Notification) error
	GetNotifications(opts GetNotificationsOpts) ([]model.Notification, error)
	UpdateNotification(notification *model.Notification) error
	GetNotificationById(id uuid.UUID) (*model.Notification, error)
}

type NotificationRepo struct {
	DB *gorm.DB
}

func NewNotificationRepo() NotificationRepoI {
	return &NotificationRepo{DB: db.Get()}
}

func (r *NotificationRepo) SaveNotification(notification model.Notification) error {
	fmt.Println(notification)
	err := r.DB.Create(notification).Error
	if err != nil {
		logger.E(nil, err, "Error while saving notification")
		return err
	}

	return nil
}

func (r *NotificationRepo) GetNotifications(opts GetNotificationsOpts) ([]model.Notification, error) {
	var notifications []model.Notification
	query := r.DB.Order("created_at desc")

	if opts.Limit != 0 {
		query = query.Limit(opts.Limit).Offset(opts.Offset)
	}

	query = query.Where("deleted_at is null")

	err := query.Find(&notifications).Error
	if err != nil {
		logger.E(nil, err, "Error while fetching notifications")
		return nil, err
	}

	return notifications, nil
}

func (r *NotificationRepo) GetNotificationById(id uuid.UUID) (*model.Notification, error) {
	var notification model.Notification

	queryErr := r.DB.Where(model.Notification{Id: id}).First(&notification).Error
	if queryErr != nil {
		logger.E(nil, queryErr, "Error while fetching notification details", logger.Field("id", id))
		return nil, queryErr
	}

	return &notification, nil
}

func (r *NotificationRepo) UpdateNotification(notification *model.Notification) error {
	err := r.DB.UpdateColumns(notification).Error
	if err != nil {
		logger.E(nil, err, "Error while updating card")
		return err
	}

	return nil
}
