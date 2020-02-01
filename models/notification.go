package models

import "fmt"

type Information struct {
	ID              int64  `json:""`
	InformationName string `json:"information_name"`
	Enabled         bool   `jsno:"enabled"`
}

func (i *Information) TableName() string { return "informations" }

func (i *Information) String() string {
	return fmt.Sprintf("Information{id:%d, information_name:%s enabled:%v}",
		i.ID, i.InformationName, i.Enabled,
	)
}
