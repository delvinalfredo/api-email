package Models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type EmailAccount struct {
	gorm.Model
	Name          string `gorm:"column:name;type:varchar(30);not null" json:"name"`
	Description   string `gorm:"column:description;type:varchar(255);not null" json:"description"`
	Email         string `gorm:"column:email;type:varchar(255);not null" json:"email"`
	Username      string `gorm:"column:username;type:varchar(30);not null" json:"username"`
	Password      string `gorm:"column:password;type:varchar(30);not null" json:"password"`
	IsTestAccount bool   `gorm:"column:is_test_account;type:bool;default:false" json:"isTestAccount"`
}

type EmailTemplate struct {
	gorm.Model
	Name         string `gorm:"column:name;type:varchar(30);not null" json:"name"`
	Description  string `gorm:"column:description;type:varchar(255);not null" json:"description"`
	Subject      string `gorm:"column:subject;type:varchar(255);not null" json:"subject"`
	HTMLTemplate string `gorm:"column:html_template;type:longtext;not null" json:"htmlTemplate"`
}

type Publisher struct {
	gorm.Model
	SendingStartDate time.Time     `gorm:"column:sending_start_date;not null" json:"sendingStartDate"`
	SendingEndDate   time.Time     `gorm:"column:sending_end_date;not null" json:"sendingEndDate"`
	SendingStartTime time.Time     `gorm:"column:sending_start_time;not null" json:"sendingStartTime"`
	SendingEndTime   time.Time     `gorm:"column:sending_end_time;not null" json:"sendingEndTime"`
	OnlyWeekdays     bool          `gorm:"column:only_weekdays;type:bool;default:false;not null" json:"onlyWeekdays"`
	Status           string        `gorm:"column:status;type:varchar(30);not null" json:"status"`
	ErrorStatus      string        `gorm:"column:error_status;type:varchar(30)" json:"errorStatus"`
	EmailTemplateID  uint          `json:"emailTemplateId" gorm:"column:email_template_id"`
	EmailTemplate    EmailTemplate `gorm:"foreignKey:email_template_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	EmailAccountID   uint          `json:"emailAccountId" gorm:"column:email_account_id"`
	EmailAccount     EmailAccount  `gorm:"foreignKey:email_account_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type PublisherParticipant struct {
	gorm.Model
	PublisherID string    `gorm:"column:publisher_id" json:"publisherId"`
	Publisher   Publisher `gorm:"foreignKey:publisher_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Title       string    `gorm:"column:title;type:varchar(255)" json:"title"`
	FirstName   string    `gorm:"column:first_name;type:varchar(30)" json:"firstName"`
	LastName    string    `gorm:"column:last_name;type:varchar(30)" json:"lastName"`
	Email       string    `gorm:"column:email;type:varchar(255)" json:"email"`
	Status      string    `gorm:"column:status;type:varchar(30);not null" json:"status"`
}

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

// type PublisherData struct {
// 	gorm.Model
// 	EmailAccountID       int
// 	EmailTemplateID      int
// 	PublisherID          int
// 	PublisherRecipientID int
// 	EmailAccount         EmailAccount
// 	EmailTemplate        EmailTemplate
// 	Publisher            Publisher
// 	PublisherRecipient   PublisherRecipient
// }
