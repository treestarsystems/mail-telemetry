package utils

type Scenario struct {
	Name        string `bson:"name" json:"name" binding:"required"`
	Type        string `bson:"type" json:"type" binding:"required"`
	FromEmail   string `bson:"fromEmail" json:"fromEmail" binding:"required"`
	ToEmail     string `bson:"toEmail" json:"toEmail" binding:"required"`
	Description string `bson:"description" json:"description"`
}
