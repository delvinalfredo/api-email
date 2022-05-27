package Controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"mail.blast/Config"
	"mail.blast/Models"
	"mail.blast/dto"

	"github.com/gin-gonic/gin"
)

func GetParticipant(c *gin.Context) {
	var participant []Models.PublisherParticipant
	// var publisher Models.Publisher
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	err := Models.GetAllParticipant(&participant, db)
	response := []dto.ResPublisherParticipant{}
	for _, participant := range participant {
		tmp := dto.ResPublisherParticipant{
			PublisherID: participant.PublisherID,
			ID:          participant.ID,
			Title:       participant.Title,
			FirstName:   participant.FirstName,
			LastName:    participant.LastName,
			Email:       participant.Email,
			Status:      participant.Status,
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

func CreateParticipantExcel(c *gin.Context) {
	// id := c.Params.ByName("id")
	var reqParticipant dto.ReqPublisherParticipant
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	c.ShouldBindJSON(&reqParticipant)

	parsedID := fmt.Sprintf("%v", reqParticipant.PublisherID)
	// participant.ID = reqParticipant.ParticipantData.ID
	for i := 0; i < len(reqParticipant.ParticipantData); i++ {
		var participant Models.PublisherParticipant
		participant.PublisherID = parsedID
		participant.Title = reqParticipant.ParticipantData[i].Title
		participant.FirstName = reqParticipant.ParticipantData[i].FirstName
		participant.LastName = reqParticipant.ParticipantData[i].LastName
		participant.Email = reqParticipant.ParticipantData[i].Email
		// parsedId, _ := strconv.Atoi(id)
		err := Models.CreateParticipant(&participant, db)
		if err != nil {
			c.ShouldBindJSON(&participant)
			c.JSON(http.StatusBadRequest, dto.ResCommon{
				Status:  false,
				Message: err.Error(),
			})
			c.Abort()
			return
		}
	}

	c.JSON(http.StatusCreated, dto.ResCommon{
		Message: "data sudah dibuat",
		Status:  true,
	})

}

func CreateParticipant(c *gin.Context) {
	var participant Models.PublisherParticipant
	var reqParticipant dto.ReqParticipantAdd
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	c.ShouldBindJSON(&reqParticipant)

	parsedID := fmt.Sprintf("%v", reqParticipant.PublisherID)
	participant.PublisherID = parsedID
	participant.ID = reqParticipant.ID
	participant.Title = reqParticipant.Title
	participant.FirstName = reqParticipant.FirstName
	participant.LastName = reqParticipant.LastName
	participant.Email = reqParticipant.Email

	err := Models.CreateParticipant(&participant, db)
	if err != nil {
		c.ShouldBindJSON(&participant)
		c.JSON(http.StatusBadRequest, dto.ResCommon{
			Status:  false,
			Message: err.Error(),
		})
		c.Abort()
		return
	} else {
		res := dto.ResPublisherParticipant{
			PublisherID: participant.PublisherID,
			ID:          participant.ID,
			Title:       participant.Title,
			FirstName:   participant.FirstName,
			LastName:    participant.LastName,
			Email:       participant.Email,
			Status:      participant.Status,
		}
		c.JSON(http.StatusCreated, dto.ResCommon{
			Message: "data sudah dibuat",
			Status:  true,
			Data:    res,
		})
	}
}

func GetParticipantByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var participant []Models.PublisherParticipant
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	parsedId, _ := strconv.Atoi(id)
	err := Models.GetParticipantIDList(&participant, uint(parsedId), db)
	response := []dto.ResPublisherParticipant{}
	for _, participant := range participant {
		tmp := dto.ResPublisherParticipant{
			PublisherID: participant.PublisherID,
			ID:          participant.ID,
			Title:       participant.Title,
			FirstName:   participant.FirstName,
			LastName:    participant.LastName,
			Email:       participant.Email,
			Status:      participant.Status,
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

func UpdateParticipant(c *gin.Context) {
	var participant Models.PublisherParticipant
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	id := c.Params.ByName("id")
	parsedId, _ := strconv.Atoi(id)
	err := Models.GetParticipantID(&participant, uint(parsedId), db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	c.ShouldBindJSON(&participant)
	err = Models.UpdateParticipant(&participant, id, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		c.Abort()
		return
	} else {
		c.JSON(http.StatusOK, participant)
	}
}

func DeleteParticipant(c *gin.Context) {
	var template Models.PublisherParticipant
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	id := c.Params.ByName("id")
	err := Models.DeleteParticipant(&template, id, db)
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
