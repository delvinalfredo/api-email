package Controllers

import (
	"net/http"
	"strconv"

	"mail.blast/Config"
	"mail.blast/Models"
	"mail.blast/dto"

	"github.com/gin-gonic/gin"
)

func GetTemplate(c *gin.Context) {
	var template []Models.EmailTemplate
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	err := Models.GetAllTemplate(&template, db)
	response := []dto.ResEmailTemplate{}
	for _, template := range template {
		tmp := dto.ResEmailTemplate{
			ID:          template.ID,
			Name:        template.Name,
			Description: template.Description,
			Subject:     template.Subject,
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

func CreateTemplate(c *gin.Context) {
	var template Models.EmailTemplate
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	c.ShouldBindJSON(&template)
	err := Models.CreateTemplate(&template, db)
	// res := db.Model(&Models.E	mailAccount{}).Create(&email)
	if err != nil {
		c.ShouldBindJSON(&template)
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
			Data:    template,
		})
	}
}

// func GetTemplateByID(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var template Models.EmailTemplate
// 	db := Config.Connect()
// 	con, _ := db.DB()
// 	defer con.Close()
// 	err := Models.GetTemplateByID(&template, id, db)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  http.StatusBadRequest,
// 			"message": err.Error(),
// 		})
// 		c.Abort()
// 		return
// 	} else {
// 		c.JSON(http.StatusOK, template)
// 	}
// }

func GetTemplateByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var template Models.EmailTemplate
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	parsedId, _ := strconv.Atoi(id)
	err := Models.GetTemplateByID(&template, uint(parsedId), db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		c.Abort()
		return
	} else {
		tmp := dto.ResEmailTemplate{
			ID:          template.ID,
			Name:        template.Name,
			Description: template.Description,
			Subject:     template.Subject,
		}
		c.JSON(http.StatusOK, tmp)
	}

}

func UpdateTemplate(c *gin.Context) {
	var template Models.EmailTemplate
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	id := c.Params.ByName("id")
	parsedId, _ := strconv.Atoi(id)
	err := Models.GetTemplateByID(&template, uint(parsedId), db)
	if err != nil {
		c.JSON(http.StatusNotFound, template)
	}
	c.ShouldBindJSON(&template)
	err = Models.UpdateTemplate(&template, id, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		c.Abort()
		return
	} else {
		c.JSON(http.StatusOK, template)
	}
}

func DeleteTemplate(c *gin.Context) {
	var template Models.EmailTemplate
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	id := c.Params.ByName("id")
	err := Models.DeleteTemplate(&template, id, db)
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
