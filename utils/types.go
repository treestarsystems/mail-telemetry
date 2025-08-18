package utils

import "gorm.io/gorm"

type Scenario struct {
	Name                    string `json:"name" binding:"required"`
	Type                    string `json:"type" binding:"required"`
	EnableTestVirtruEncrypt string `json:"enableTestVirtruEncrypt"`
	EnableTestDLP           string `json:"enableTestDLP"`
	FromEmails              string `json:"fromEmails" binding:"required"`
	ToEmails                string `json:"toEmails" binding:"required"`
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

type ScenarioTable struct {
	Name                    string `json:"name" binding:"required"`
	Type                    string `json:"type" binding:"required"`
	EnableTestVirtruEncrypt string `json:"enableTestVirtruEncrypt"`
	EnableTestDLP           string `json:"enableTestDLP"`
	FromEmails              string `json:"fromEmails" binding:"required"`
	ToEmails                string `json:"toEmails" binding:"required"`
	Description             string `json:"description"`
	AttachmentFilePath      string `json:"attachmentFilePath"`
	Hosts                   string `json:"hosts"`
	Ports                   string `json:"ports"`
	Endpoints               string `json:"endpoints"`
	FileLastModified        string `json:"fileLastModified" binding:"required"`
}

type LoadDbInsertGormScenario struct {
	ScenarioTable
	ID        uint           `gorm:"primarykey"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type LoadDbInsertGormCredential struct {
	ScenarioAuth
	ID        uint           `gorm:"primarykey"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ScenarioHost struct {
	InstanceURI          string `json:"instanceUri" binding:"required"`
	OriginHostName       string `json:"originHostName" binding:"required"`
	OriginLocalIpAddress string `json:"originLocalIpAddress" binding:"required"`
}

// ScenarioDetail Types: O365
type ScenarioDetails struct {
	Details ScenarioTable `json:"details" binding:"required"`
	// Auth     ScenarioAuth    `json:"scenarioAuth" binding:"required"`
	Host    ScenarioHost    `json:"scenarioHost" binding:"required"`
	Message ScenarioMessage `json:"scenarioMessage" binding:"required"`
	Errors  []string        `json:"errors" binding:"required"`
}

// Scenario Detail Message
type ScenarioMessage struct {
	ID            string   `json:"id" binding:"required"`
	FromEmails    []string `json:"fromEmails" binding:"required"`
	ToEmails      []string `json:"toEmails" binding:"required"`
	Subject       string   `json:"subject" binding:"required"`
	BodyPlainText string   `json:"bodyPlainText" binding:"required"`
	BodyHTML      string   `json:"bodyHtml" binding:"required"`
}

type ScenarioAuth struct {
	ClientId                            string `json:"clientId" binding:"required"`
	ClientSecret                        string `json:"clientSecret" binding:"required"`
	TenantId                            string `json:"tenantId" binding:"required"`
	CredentialName                      string `json:"credentialName" binding:"required"`
	GraphApiToken                       string `json:"graphApiToken" binding:"required"`
	TokenExpireAtTimeStampMilliseconds  int64  `json:"TokenExpireAtTimeStampMilliseconds" binding:"required"`
	TokenUpdatedAtTimeStampMilliseconds int64  `json:"TokenUpdatedAtTimeStampMilliseconds" binding:"required"`
	SmtpUsername                        string `json:"smtpUsername" binding:"required"`
	SmtpPassword                        string `json:"smtpPassword" binding:"required"`
}
