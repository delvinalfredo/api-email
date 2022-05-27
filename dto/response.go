package dto

import "time"

type ResCommon struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type ResEmailAccount struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Email       string `json:"email"`
}

type ResEmailTemplate struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

type ResPublisher struct {
	ID               uint      `json:"id"`
	EmailAccount     string    `json:"emailAccount"`
	EmailTemplate    string    `json:"emailTemplate"`
	SendingStartDate time.Time `json:"sendingStartDate"`
	SendingEndDate   time.Time `json:"sendingEndDate"`
	SendingStartTime time.Time `json:"sendingStartTime"`
	SendingEndTime   time.Time `json:"sendingEndTime"`
	UpdatedAt        time.Time `json:"update"`
}

type ResPublisherParticipant struct {
	PublisherID string `json:"publisher"`
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Status      string `json:"status"`
}

type ResUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type ResToken struct {
	Token string `json:"token"`
	//RefreshToken string `json:"refreshToken"`
	ExpiredAt uint `json:"expiredAt"`
}
