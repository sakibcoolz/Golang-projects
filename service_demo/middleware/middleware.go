package middleware

import (
	"log"
	"service_demo/controller"
	"service_demo/database"

	"github.com/gin-gonic/gin"
)

func ContextGinToController(r *gin.Engine) *gin.Engine {
	DB, err := database.DbConnector()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(DB)

	controller := controller.GiveController(DB)

	r.GET("/", controller.IControl1.Ping)

	log.Println(controller)

	return r
}
