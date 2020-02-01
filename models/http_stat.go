package models

import (
	"fmt"
	"time"
)

type HTTPStat struct {
	ID         int64     `json:"id"`
	TriggerID  int64     `json:"trigger_id"`
	DNS        int       `json:"dns"`
	TCP        int       `json:"tcp"`
	TLS        int       `json:"tls"`
	Processing int       `json:"processing"`
	Transfer   int       `json:"transfer"`
	StatusCode int       `json:"status_code"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `xorm:"<-" json:"created_at"`
}

func (h *HTTPStat) TableName() string { return "http_stats" }

func (h *HTTPStat) String() string {
	return fmt.Sprintf("HTTPStat{id:%d, trigger_id:%d, dns:%d, tcp:%d, tls:%d, processing:%d, transfer:%d, status_code:%d, status:%s, created_at:%s,}",
		h.ID, h.TriggerID, h.DNS, h.TCP, h.TLS, h.Processing, h.Transfer,
		h.StatusCode, h.Status, h.CreatedAt,
	)
}

func CreateHttpStat(httpStat *HTTPStat) error {
	_, err := engine.Insert(httpStat)
	return err
}
