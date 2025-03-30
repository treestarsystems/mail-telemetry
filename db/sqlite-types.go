package db

import (
	"gorm.io/gorm"
)

// SharedStructJobs contains common job information from each job listing.
type Scenario struct {
	ScenarioId string `bson:"job_id" json:"jobId" binding:"required"`
}

type LoadDbInsertGorm struct {
	Scenario
	ID        uint           `gorm:"primarykey"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
