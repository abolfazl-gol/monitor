package convert

import (
	"../api-monitor/models"
	monitoring "../api-monitor/monitor/v1"
)

func TriggerToProto(trigger *models.Trigger) *monitoring.Trigger {
	return &monitoring.Trigger{
		Id:               trigger.ID,
		UserId:           trigger.UserID,
		MonitorId:        trigger.MonitorID,
		TriggerKind:      trigger.TriggerKind,
		NotificationKind: trigger.NotificationKind,
		ErrTolerance:     trigger.ErrTolerance,
		Warning:          trigger.Warning,
		Status:           trigger.Status,
		Enabled:          trigger.Enabled,
	}
}
