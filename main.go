package main

import (
	"fmt"
	"github.com/ValeryVerkhoturov/chat/config"
	"github.com/ValeryVerkhoturov/chat/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(cors.Default()) // AllowAllOrigins true
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.Static("/images/", "./images")
	router.StaticFile("/css/output-css", "./css/output.css")
	router.LoadHTMLGlob("controller/templates/*")

	router.GET("/", controller.Index)
	router.GET("/index.html", controller.Index)
	router.GET("/v1/chat-widget", controller.ChatWidgetV1)
	router.GET("/v1/chat-container", controller.ChatContainerV1)

	var port = config.Port
	if config.Port == "" {
		port = "8080"
	}
	fmt.Println("Listening on http://" + config.Host + ":" + port)
	err := router.Run(":" + port)
	if err != nil {
		fmt.Println(err)
	}
}
