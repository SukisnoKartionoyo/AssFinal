package models

//Status grom model
type Status struct {
	// gorm.Model
	StatusID   uint   `json:"StatusID" gorm:"primaryKey"`
	StatusDesc string `json:"StatusDesc"`
}
