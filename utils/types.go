package utils

type Scenario struct {
	Name               string `json:"name" binding:"required"`
	Type               string `json:"type" binding:"required"`
	CredentialLocation string `json:"credentialLocation" binding:"required"`
	FromEmail          string `json:"fromEmail" binding:"required"`
	ToEmail            string `json:"toEmail" binding:"required"`
	Description        string `json:"description"`
}
