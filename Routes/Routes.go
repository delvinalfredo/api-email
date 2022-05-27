package Routes

import (
	"mail.blast/Controllers"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000", "http://localhost"}
	// config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	// config.ExposeHeaders = []string{"Content-Length", "Content-Type"}

	// router.Use(cors.New(config))

	// v1 := router.Group("mail/v1")
	// {
	// 	v1.GET("/mail", Controllers.GetAccount)
	// 	v1.POST("/mail", Controllers.CreateAccount)
	// }
	template := router.Group("api/talent/template")
	{
		template.GET("/", Controllers.GetTemplate)
		template.POST("/", Controllers.CreateTemplate)
		template.GET("/:id", Controllers.GetTemplateByID)
		template.PUT("/:id", Controllers.UpdateTemplate)
		template.DELETE("/:id", Controllers.DeleteTemplate)
	}
	email := router.Group("api/talent/email")
	{
		email.GET("/", Controllers.GetAccount)
		email.POST("/", Controllers.CreateAccount)
		email.GET("/:id", Controllers.GetAccountByID)
		email.PUT("/:id", Controllers.UpdateAccount)
		email.DELETE("/:id", Controllers.DeleteAccount)
	}
	publisher := router.Group("api/talent/publisher")
	{
		publisher.GET("/", Controllers.GetPublisher)
		publisher.POST("/", Controllers.CreatePublisher)
		publisher.GET("/:id", Controllers.GetPublisherByID)
		publisher.PUT("/:id", Controllers.UpdatePublisher)
		publisher.DELETE("/:id", Controllers.DeletePublisher)
	}

	participant := router.Group("api/talent/participant")
	{
		participant.GET("/", Controllers.GetParticipant)
		participant.POST("/add", Controllers.CreateParticipant)
		participant.POST("/", Controllers.CreateParticipantExcel)
		participant.GET("/:id", Controllers.GetParticipantByID)
		participant.PUT("/:id", Controllers.UpdateParticipant)
		participant.DELETE("/:id", Controllers.DeleteParticipant)
	}

	login := router.Group("api/talent")
	{
		login.POST("/login", Controllers.Login)
		login.POST("/register", Controllers.Register)
		login.GET("/login", Controllers.GetUser)
		login.DELETE("/login/:id", Controllers.DeleteUser)
	}

	router.POST("api/talent/send-email/", Controllers.PostSendEmail)
	// router.POST("/upload", Controllers.Upload)
	return router
}
