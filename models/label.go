package models

//Label nya task
type Label struct {
	// gorm.Model
	LabelID    uint32 `json:"LabelId" gorm:"primaryKey"`
	LabelTitle string `json:"LabelTitle"`
}
