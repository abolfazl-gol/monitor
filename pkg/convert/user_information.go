package convert

import (
	"git.tulz.ir/abolfazl/fing-monitor/models"
	monitoring "git.tulz.ir/abolfazl/fing-monitor/monitor/v1"
)

func UserNotificationToProto(userInfo *models.UserNotification) *monitoring.UserNotification {
	return &monitoring.UserNotification{
		Id:             userInfo.ID,
		UserId:         userInfo.UserID,
		NotificationId: userInfo.NotificationID,
		Title:          userInfo.Title,
		Url:            userInfo.URL,
		MethodKind:     userInfo.MethodKind,
		Phone:          userInfo.Phone,
		Email:          userInfo.Email,
		Code:           userInfo.Code,
		ApiKey:         userInfo.APIKey,
		Enabled:        userInfo.Enabled,
	}
}
