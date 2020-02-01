package convert

import (
	"git.tulz.ir/abolfazl/fing-monitor/models"
	monitoring "git.tulz.ir/abolfazl/fing-monitor/monitor/v1"
)

func MonitorToProto(monitor *models.Monitor) *monitoring.Monitor {
	return &monitoring.Monitor{
		Id:             monitor.ID,
		UserId:         monitor.UserID,
		ProtocolKind:   monitor.ProtocolKind,
		Title:          monitor.Title,
		Method:         monitor.Method,
		Url:            monitor.URL,
		RequestTime:    monitor.RequestTime,
		ServerLocation: monitor.ServerLocation,
		Enabled:        monitor.Enabled,
	}
}
