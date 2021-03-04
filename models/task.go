package models

import (
	"time"
)

//Task sruck for db
type Task struct {
	// gorm.Model
	TaskID    uint      `json:"TaskID" gorm:"primary_key"`
	Title     string    `json:"Title"`
	Time      time.Time `json:"Time"`
	ProjectID uint      `json:"ProjectID" gorm:"foreignKey:ProjectID"`
	LabelID   uint      `json:"LebelID" gorm:"foreignKey:LabelID"`
	StatusID  uint      `json:"StatusID" gorm:"foreignKey:StatusID"`
}
