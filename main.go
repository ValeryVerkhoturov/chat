package main

import (
	"github.com/ValeryVerkhoturov/chat/config"
	"github.com/ValeryVerkhoturov/chat/routers"
	v1Routers "github.com/ValeryVerkhoturov/chat/routers/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
)

func port() string {
	port := config.Port
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	router := gin.Default()
	router.Use(cors.Default()) // AllowAllOrigins true
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.Static("/images/", "./public/images")
	router.StaticFile("/css/output-css", "./public/css/output.css")
	router.LoadHTMLGlob("routers/templates/*")

	router.GET("/", routers.Index)
	router.GET("/index.html", routers.Index)
	router.GET("/v1/chat-widget", v1Routers.ChatWidget)
	router.GET("/v1/chat-container", v1Routers.ChatContainer)

	log.Info("Starting Server on http://localhost" + port())

	err := router.Run(port())
	if err != nil {
		log.Error(err)
	}
}
