package models

import (
	"fmt"
	"time"

	"git.tulz.ir/abolfazl/fing-monitor/pkg/errors"
)

type Trigger struct {
	ID               int64     `json:"id"`
	UserID           int64     `json:"user_id"`
	MonitorID        int64     `json:"monitor_id"`
	TriggerKind      string    `json:"trigger_kind"`
	NotificationKind string    `json:"notification_kind"`
	ErrTolerance     string    `json:"err_tolerance"`
	Warning          string    `json:"warning"`
	Status           string    `json:"status"`
	Enabled          bool      `json:"enabled"`
	CreatedAt        time.Time `xorm:"<-" json:"created_at"`
	UpdateAt         time.Time `xorm:"<-" json:"updated_at"`
}

func (t *Trigger) TableName() string { return "triggers" }

func (t *Trigger) String() string {
	return fmt.Sprintf("Trigger{id:%d, user_id:%d, monitor_id:%d, trigger_kind:%s, notification_kind:%s, err_tolerance:%s, warning:%s, status:%s, enabled:%v, created_at:%s, updated_at:%s,}",
		t.ID, t.UserID, t.MonitorID, t.TriggerKind, t.NotificationKind,
		t.ErrTolerance, t.Warning, t.Status, t.Enabled, t.CreatedAt, t.UpdateAt,
	)
}

func FindTriggers() ([]*Trigger, error) {
	triggers := make([]*Trigger, 0)
	if err := engine.Find(&triggers); err != nil {
		return nil, err
	}

	return triggers, nil
}

func GetTriggerByID(id interface{}) (*Trigger, error) {
	return getTriggerBy("id", id)
}

func getTriggerBy(column string, value interface{}) (*Trigger, error) {
	trigger := new(Trigger)
	has, err := engine.Where(fmt.Sprintf("%s =?", column), value).Get(trigger)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.ErrNotFound{Resource: "monitro", Column: column, Value: value}
	}

	return trigger, nil
}

func CreateTrigger(trigger *Trigger) error {
	_, err := engine.Insert(trigger)
	return err
}

func UpdateTrigger(trigger *Trigger, cols ...string) error {
	_, err := engine.Where("id = ?", trigger.ID).Cols(cols...).Update(trigger)
	return err
}

func DeleteTrigger(id int64) error {
	_, err := engine.Delete(&Trigger{ID: id})
	return err
}
