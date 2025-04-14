package db

import (
	"mail-telemetry/utils"

	"gorm.io/gorm"
)

type LoadDbInsertGormScenario struct {
	utils.Scenario
	ID        uint           `gorm:"primarykey"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type LoadDbInsertGormCredential struct {
	utils.Credential
	ID        uint           `gorm:"primarykey"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
