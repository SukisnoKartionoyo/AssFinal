package models

//Project Data untuk project
type Project struct {
	// gorm.Model
	ProjectID    uint32 `json:"ProjectId" gorm:"primaryKey"`
	ProjectTitle string `json:"TitleProject"`
}
