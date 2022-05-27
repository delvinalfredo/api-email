package Controllers

import (
	"net/http"
	"strconv"

	"mail.blast/Config"
	"mail.blast/Models"
	"mail.blast/dto"

	"github.com/gin-gonic/gin"
)

func GetAccount(c *gin.Context) {
	var account []Models.EmailAccount
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	err := Models.GetAllEmail(&account, db)
	response := []dto.ResEmailAccount{}
	for _, account := range account {
		tmp := dto.ResEmailAccount{
			ID:          account.ID,
			Name:        account.Name,
			Description: account.Description,
			Email:       account.Email,
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

func CreateAccount(c *gin.Context) {
	var account Models.EmailAccount
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	c.ShouldBindJSON(&account)
	err := Models.CreateEmail(&account, db)
	if err != nil {
		c.ShouldBindJSON(&account)
		c.JSON(http.StatusBadRequest, dto.ResCommon{
			Status:  false,
			Message: err.Error(),
		})
		c.Abort()
		return
	} else {
		c.JSON(http.StatusCreated, dto.ResCommon{
			Message: "data sudah dibuat",
			Status:  true,
			Data:    account,
		})
	}
}

func GetAccountByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var account Models.EmailAccount
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	parsedId, _ := strconv.Atoi(id)
	err := Models.GetEmailByID(&account, uint(parsedId), db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		c.Abort()
		return
	} else {
		tmp := dto.ResEmailAccount{
			ID:          account.ID,
			Name:        account.Name,
			Description: account.Description,
			Email:       account.Email,
		}
		c.JSON(http.StatusOK, tmp)
	}
}

func UpdateAccount(c *gin.Context) {
	var account Models.EmailAccount
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	id := c.Params.ByName("id")
	parsedId, _ := strconv.Atoi(id)
	err := Models.GetEmailByID(&account, uint(parsedId), db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	c.ShouldBindJSON(&account)
	err = Models.UpdateEmail(&account, id, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		c.Abort()
		return
	} else {
		c.JSON(http.StatusOK, account)
	}
}

func DeleteAccount(c *gin.Context) {
	var account Models.EmailAccount
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	id := c.Params.ByName("id")
	err := Models.DeleteEmail(&account, id, db)
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
