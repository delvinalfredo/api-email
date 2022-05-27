package Controllers

import (
	"net/http"
	"strconv"

	"mail.blast/Config"
	"mail.blast/Models"
	"mail.blast/dto"
	"mail.blast/query_model"

	"github.com/gin-gonic/gin"
)

func GetPublisher(c *gin.Context) {
	var publisher []query_model.PublisherQuery
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	err := Models.GetAllPublisher(&publisher, db)
	response := []dto.ResPublisher{}
	for _, publisher := range publisher {
		tmp := dto.ResPublisher{
			ID:               publisher.ID,
			// Status:           publisher.Status,
			EmailAccount:     publisher.EmailAccount,
			EmailTemplate:    publisher.EmailTemplate,
			SendingStartDate: publisher.SendingStartDate,
			SendingEndDate:   publisher.SendingEndDate,
			SendingStartTime: publisher.SendingStartTime,
			SendingEndTime:   publisher.SendingEndTime,
			UpdatedAt:        publisher.UpdatedAt,
		}
		response = append(response, tmp)
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		c.Abort()
		return
	} else {
		c.JSON(http.StatusOK, response)
	}
}

func CreatePublisher(c *gin.Context) {
	var publisher dto.ReqPublisher
	var publisherModel Models.Publisher
	// id := c.Params.ByName("id")
	c.ShouldBindJSON(&publisher)
	publisherModel.ID = publisher.ID
	publisherModel.EmailAccountID = publisher.EmailAccountID
	publisherModel.EmailTemplateID = publisher.EmailTemplateID
	// publisherModel.Status = publisher.Status
	publisherModel.SendingStartDate = publisher.SendingStartDate
	publisherModel.SendingEndDate = publisher.SendingEndDate
	publisherModel.SendingStartTime = publisher.SendingStartTime
	publisherModel.SendingEndTime = publisher.SendingEndTime
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	Models.CreatePublisher(&publisherModel, db)
	
	dataResponse := dto.ResPublisher{
		ID: publisherModel.ID,
		// Status: publisherModel.Status,
		// EmailAccount: publisherModel.EmailAccount.Email,
		// EmailTemplate: publisherModel.EmailTemplate.Name,
		SendingStartDate: publisherModel.SendingStartDate,
		SendingEndDate: publisherModel.SendingEndDate,
		SendingStartTime: publisherModel.SendingStartTime,
		SendingEndTime: publisherModel.SendingEndTime,
		
	}
	c.JSON(http.StatusCreated, dto.ResCommon{
		Message: "data sudah dibuat",
		Status:  true,
		Data:    dataResponse,
	})
}

func GetPublisherByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var publisher Models.Publisher
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	parsedId, _ := strconv.Atoi(id)
	err := Models.GetPublisherByID(&publisher, db, uint(parsedId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		c.Abort()
		return
	} else {
		parsedId, _ := strconv.Atoi(id)
		err := Models.GetEmailByID(&publisher.EmailAccount, uint(parsedId), db)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		parsedId, _ = strconv.Atoi(id)
		err = Models.GetTemplateByID(&publisher.EmailTemplate, uint(parsedId), db)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		tmp := dto.ResPublisher{
			ID:               publisher.ID,
			// Status:           publisher.Status,
			EmailAccount:     publisher.EmailAccount.Email,
			EmailTemplate:    publisher.EmailTemplate.Name,
			SendingStartDate: publisher.SendingStartDate,
			SendingEndDate:   publisher.SendingEndDate,
			SendingStartTime: publisher.SendingStartTime,
			SendingEndTime:   publisher.SendingEndTime,
			UpdatedAt:        publisher.UpdatedAt,
		}
		c.JSON(http.StatusOK, tmp)
	}
}

func UpdatePublisher(c *gin.Context) {
	var publisher Models.Publisher
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	id := c.Params.ByName("id")
	parsedId, _ := strconv.Atoi(id)
	err := Models.GetPublisherByID(&publisher, db, uint(parsedId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	c.ShouldBindJSON(&publisher)
	err = Models.UpdatePublisher(&publisher, id, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		c.Abort()
		return
	} else {
		c.JSON(http.StatusOK, publisher)
	}
}

func DeletePublisher(c *gin.Context) {
	var publisher Models.Publisher
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	id := c.Params.ByName("id")
	err := Models.DeletePublisher(&publisher, id, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		c.Abort()
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
