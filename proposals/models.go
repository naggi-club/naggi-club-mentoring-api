package proposals

import (
	"github.com/jinzhu/gorm"
	"github.com/naggi-club/naggi-club-mentoring-api/common"
)

type (
	Proposal struct {
		gorm.Model
		Name                string `form:"name" json:"name" binding:"required" gorm:"not null"`
		Email               string `form:"email" json:"email" binding:"required" gorm:"unique;not null"`
		PrograminExperience bool   `form:"programming_experience" json:"programming_experience" binding:"required" gorm:"not null"`
		RequestMotivation   int    `form:"request_motivation" json:"request_motivation" binding:"required" gorm:"not null"`
		Content             string `form:"content" json:"content" binding:"required" gorm:"not null"`
	}
)

func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&Proposal{})
}
