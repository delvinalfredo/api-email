package Controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"mail.blast/Config"
	"mail.blast/Models"

	"mail.blast/dto"
)

func PostSendEmail(c *gin.Context) {
	var mail dto.ReqPublisher
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()

	err := c.ShouldBindJSON(&mail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	var publisher Models.Publisher
	Models.GetPublisherByID(&publisher, db, mail.ID)

	var accountConfig Models.EmailAccount
	Models.GetEmailByID(&accountConfig, publisher.EmailAccountID, db)

	var mailTemplate Models.EmailTemplate
	Models.GetTemplateByID(&mailTemplate, publisher.EmailTemplateID, db)

	var participants []Models.PublisherParticipant
	Models.GetParticipantByID(&participants, publisher.ID, db)

	for _, participant := range participants {

		AUTH_EMAIL := accountConfig.Email
		AUTH_PASSWORD := accountConfig.Password
		SMTP_HOST := os.Getenv("SMTP_HOST")
		SMTP_PORT, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
		m := gomail.NewMessage()
		m.SetHeader("From", accountConfig.Email)
		m.SetAddressHeader("To", participant.Email, participant.FirstName)
		m.SetHeader("Subject", mailTemplate.Subject)
		m.SetBody("text/html", mailTemplate.HTMLTemplate)

		emailLog := Models.PublisherParticipant{}
		parsedID := fmt.Sprintf("%v", publisher.ID)
		emailLog.Status = "SUCCESS"
		emailLog.PublisherID = parsedID
		
		d := gomail.NewDialer(SMTP_HOST, SMTP_PORT, AUTH_EMAIL, AUTH_PASSWORD)
		if err := d.DialAndSend(m); err != nil {
			log.Println(err)
			emailLog.Status = "FAILED"
			db.Model(&participant).Update("status", emailLog.Status)
		}
		db.Model(&participant).Update("status", emailLog.Status)
	}
	c.JSON(http.StatusOK, dto.ResCommon{
		Message: "Email sent!",
		Status:  true,
	})
}
