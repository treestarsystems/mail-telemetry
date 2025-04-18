package utils

type Credential struct {
	Name         string `json:"name" binding:"required"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}
type Scenario struct {
	Name               string `json:"name" binding:"required"`
	Type               string `json:"type" binding:"required"`
	CredentialLocation string `json:"credentialLocation" binding:"required"`
	FromEmail          string `json:"fromEmail" binding:"required"`
	ToEmail            string `json:"toEmail" binding:"required"`
	Description        string `json:"description"`
	AttachmentFilePath string `json:"attachmentFilePath"`
	FileLastModified   string `json:"fileLastModified" binding:"required"`
}
