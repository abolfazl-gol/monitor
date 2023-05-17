package models

import (
	"fmt"
	"time"

	"../api-monitor/pkg/errors"
)

type Monitor struct {
	ID             int64     `json:"id"`
	UserID         int64     `json:"user_id"`
	ProtocolKind   string    `json:"protocol_kind"`
	Title          string    `json:"title"`
	Method         string    `json:"method"`
	URL            string    `json:"url"`
	RequestTime    string    `json:"request_time"`
	ServerLocation string    `json:"server_location"`
	Status         string    `json:"status"`
	Enabled        bool      `json:"enabled"`
	CreatedAt      time.Time `xorm:"<-" json:"created_at"`
	UpdatedAt      time.Time `xorm:"<-" json:"updated_at"`
}

func (m *Monitor) TableName() string { return "monitors" }

func (m *Monitor) String() string {
	return fmt.Sprintf("Monitor{id:%d, user_id:%d, protocol_kind:%s, title:%s, method:%s, url:%s, request_time:%s, server_location:%s, status:%s, enabled:%v, created_at:%s, updated_at:%s,}",
		m.ID, m.UserID, m.ProtocolKind, m.Title, m.Method, m.URL, m.RequestTime,
		m.ServerLocation, m.Status, m.Enabled, m.CreatedAt, m.UpdatedAt,
	)
}

func FindMonitors() ([]*Monitor, error) {
	monitors := make([]*Monitor, 0)
	if err := engine.Find(&monitors); err != nil {
		return nil, err
	}

	return monitors, nil
}

func GetMonitorByID(id interface{}) (*Monitor, error) {
	return getMonitorBy("id", id)
}

func getMonitorBy(column string, value interface{}) (*Monitor, error) {
	monitor := new(Monitor)
	has, err := engine.Where(fmt.Sprintf("%s =?", column), value).Get(monitor)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.ErrNotFound{Resource: "monitro", Column: column, Value: value}
	}

	return monitor, nil
}

func CreateMonitor(monitor *Monitor) error {
	_, err := engine.Insert(monitor)
	return err
}

func UpdateMonitor(monitor *Monitor, cols ...string) error {
	_, err := engine.Where("id = ?", monitor.ID).Cols(cols...).Update(monitor)
	return err
}

func DeleteMonitor(id int64) error {
	_, err := engine.Delete(&Monitor{ID: id})
	return err
}
