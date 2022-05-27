package query_model

import (
	"time"

	"gorm.io/gorm"
)

type PublisherQuery struct {
	gorm.Model
	SendingStartDate time.Time `gorm:"column:sending_start_date;not null" json:"sendingStartDate"`
	SendingEndDate   time.Time `gorm:"column:sending_end_date;not null" json:"sendingEndDate"`
	SendingStartTime time.Time `gorm:"column:sending_start_time;not null" json:"sendingStartTime"`
	SendingEndTime   time.Time `gorm:"column:sending_end_time;not null" json:"sendingEndTime"`
	OnlyWeekdays     bool      `gorm:"column:only_weekdays;type:bool;default:false;not null" json:"onlyWeekdays"`
	Status           string    `gorm:"column:status;type:varchar(30);not null" json:"status"`
	ErrorStatus      string    `gorm:"column:error_status;type:varchar(30)" json:"errorStatus"`
	EmailTemplateID  uint      `json:"emailTemplateId" gorm:"column:email_template_id"`
	EmailTemplate    string    `gorm:"column:name"`
	EmailAccountID   uint      `json:"emailAccountId" gorm:"column:email_account_id"`
	EmailAccount     string    `gorm:"column:email"`
}

