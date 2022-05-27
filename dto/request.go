package dto

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type ReqPublisher struct {
	ID               uint      `json:"id"`
	Status           string    `json:"status"`
	EmailAccountID   uint      `json:"emailAccountId"`
	EmailTemplateID  uint      `json:"emailTemplateId"`
	SendingStartDate time.Time `json:"sendingStartDate"`
	SendingEndDate   time.Time `json:"sendingEndDate"`
	SendingStartTime time.Time `json:"sendingStartTime"`
	SendingEndTime   time.Time `json:"sendingEndTime"`
}

type ReqEmailAccount struct {
	ID       uint   `json:"emailAccountId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ReqEmailTemplate struct {
	ID           uint   `json:"emailTemplateId"`
	Name         string `json:"name"`
	Subject      string `json:"subject"`
	HTMLTemplate string `json:"htmlTemplate"`
}

type ReqPublisherParticipant struct {
	PublisherID uint `json:"publisherId"`
	ParticipantData []ReqParticipantData `json:"participantData"`
}
type ReqParticipantData struct {
	Title     string `json:"title"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type ReqParticipantAdd struct {
	PublisherID uint   `json:"publisherId"`
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
}

type ReqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtClaims struct {
	Id string `json:"userId"`
	jwt.StandardClaims
}

type ReqEmail struct {
	From         string   `json:"from" binding:"required"`
	To           []string `json:"to" binding:"required"`
	Subject      string   `json:"subject" binding:"required"`
	Content      string   `json:"content" binding:"required"`
	AttachmentID []string `json:"attachments"`
}
