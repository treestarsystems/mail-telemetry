package utils

import "gorm.io/gorm"

type Credential struct {
	Name         string `json:"name" binding:"required"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type Scenario struct {
	Name                    string `json:"name" binding:"required"`
	Type                    string `json:"type" binding:"required"`
	EnableTestVirtruEncrypt string `json:"enableTestVirtruEncrypt"`
	EnableTestDLP           string `json:"enableTestDLP"`
	FromEmail               string `json:"fromEmail" binding:"required"`
	ToEmail                 string `json:"toEmail" binding:"required"`
	Description             string `json:"description"`
	AttachmentFilePath      string `json:"attachmentFilePath"`
	Hosts                   string `json:"hosts"`
	Ports                   string `json:"ports"`
	Endpoints               string `json:"endpoints"`
	ClientId                string `json:"clientId"`
	ClientSecret            string `json:"clientSecret"`
	TenantId                string `json:"tenantId"`
	SmtpUsername            string `json:"smtpUsername"`
	SmtpPassword            string `json:"smtpPassword"`
	FileLastModified        string `json:"fileLastModified" binding:"required"`
}

type LoadDbInsertGormScenario struct {
	Scenario
	ID        uint           `gorm:"primarykey"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type LoadDbInsertGormCredential struct {
	Credential
	ID        uint           `gorm:"primarykey"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// ScenarioDetail Types: O365
type ScenarioDetailsO365 struct {
	Scenario Scenario         `json:"scenario" binding:"required"`
	Auth     ScenarioAuthO365 `json:"scenarioAuth" binding:"required"`
	Host     ScenarioHostO365 `json:"scenarioHost" binding:"required"`
}

type ScenarioAuthO365 struct {
	ClientId     string `json:"clientId" binding:"required"`
	ClientSecret string `json:"clientSecret" binding:"required"`
	TenantId     string `json:"tenantId" binding:"required"`
}

type ScenarioHostO365 struct {
	Address  string `json:"address" binding:"required"`
	Port     uint   `json:"port" binding:"required"`
	Endpoint string `json:"endpoint" binding:"required"`
}

// ScenarioDetail Types: SMTP
type ScenarioDetailsSMTP struct {
	Scenario Scenario         `json:"scenario" binding:"required"`
	Auth     ScenarioAuthSMTP `json:"scenarioAuth" binding:"required"`
	Host     ScenarioHostSMTP `json:"scenarioHost" binding:"required"`
}

type ScenarioAuthSMTP struct {
	Username string `json:"Username" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

type ScenarioHostSMTP struct {
	Address string `json:"address" binding:"required"`
	Port    uint   `json:"port" binding:"required"`
}
