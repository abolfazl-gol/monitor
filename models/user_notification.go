package models

import (
	"fmt"

	"git.tulz.ir/abolfazl/fing-monitor/pkg/errors"
)

type UserNotification struct {
	ID             int64  `json:"id"`
	UserID         int64  `json:"user_id"`
	NotificationID int64  `json:"notification_id"`
	Title          string `json:"title"`
	URL            string `json:"url"`
	MethodKind     string `json:"method_kind"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Code           string `json:"code"`
	APIKey         string `json:"api_key"`
	Enabled        bool   `json:"enabled"`
}

func (ui *UserNotification) TableName() string { return "user_notifications" }

func (ui *UserNotification) String() string {
	return fmt.Sprintf("UserNotification{id:%d, user_id:%d, notification_id:%d, title:%s, url:%s, method_kind:%s, phone:%s, email:%s, code:%s, api_key:%s, enabled:%v,}",
		ui.ID, ui.UserID, ui.NotificationID, ui.Title, ui.URL, ui.MethodKind,
		ui.Phone, ui.Email, ui.Code, ui.APIKey, ui.Enabled,
	)
}

func FindUserNotifications() ([]*UserNotification, error) {
	userNotifications := make([]*UserNotification, 0)
	if err := engine.Find(&userNotifications); err != nil {
		return nil, err
	}

	return userNotifications, nil
}

func GetUserNotificationByID(id interface{}) (*UserNotification, error) {
	return getUserNotificationBy("id", id)
}

func getUserNotificationBy(column string, value interface{}) (*UserNotification, error) {
	userNotification := new(UserNotification)
	has, err := engine.Where(fmt.Sprintf("%s =?", column), value).Get(userNotification)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.ErrNotFound{Resource: "monitro", Column: column, Value: value}
	}

	return userNotification, nil
}

func CreateUserNotification(userInfo *UserNotification) error {
	_, err := engine.Insert(userInfo)
	return err
}

func UpdateUserNotification(userInfo *UserNotification, cols ...string) error {
	_, err := engine.ID(userInfo.ID).Cols(cols...).Update(userInfo)
	return err
}

func DeleteUserNotification(id int64) error {
	_, err := engine.Delete(&UserNotification{ID: id})
	return err
}
