package Controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"mail.blast/Config"
	"mail.blast/Models"
	"mail.blast/dto"
)

func Login(c *gin.Context) {
	var user dto.ReqLogin
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	var currentuser Models.User
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	db.Where("username=?", user.Username).Find(&currentuser, db)
	if currentuser.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "account not found",
		})
		c.Abort()
		return
	}
	if currentuser.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "wrong username or password",
		})
		c.Abort()
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(currentuser.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte("secret"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func GetUser(c *gin.Context) {
	var user []Models.User
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	err := Models.GetAllUser(&user, db)
	response := []dto.ResUser{}
	for _, user := range user {
		tmp := dto.ResUser{
			ID:       user.ID,
			Username: user.Username,
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

func Register(c *gin.Context) {
	var user Models.User
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	c.ShouldBindJSON(&user)
	err := Models.Register(&user, db)
	// res := db.Model(&Models.EmailAccount{}).Create(&email)
	if err != nil {
		c.ShouldBindJSON(&user)
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
			Data:    user,
		})
	}
}

func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	err := Models.GetUserByID(&user, id, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		c.Abort()
		return
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {
	var user Models.User
	db := Config.Connect()
	con, _ := db.DB()
	defer con.Close()
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&user, id, db)
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
