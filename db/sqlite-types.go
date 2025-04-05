package db

import (
	"mail-telemetry/utils"

	"gorm.io/gorm"
)

// SharedStructJobs contains common job information from each job listing.
type ScenarioLoadDB struct {
	ScenarioId   string         `json:"scenarioId" binding:"required"`
	ScenarioData utils.Scenario `json:"scenarioData" binding:"required"`
}

type LoadDbInsertGorm struct {
	Scenario  ScenarioLoadDB
	ID        uint           `gorm:"primarykey"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
