package main

import (
	"embed"
	"fmt"
	"github.com/ValeryVerkhoturov/chat/requestUtils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"html/template"
)

var (
	//go:embed all:templates/*
	templateFS embed.FS

	//parsed templates
	html *template.Template
)

func main() {

	var err error
	html, err = requestUtils.TemplateParseFSRecursive(templateFS, ".html", true, nil)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Use(cors.Default()) // AllowAllOrigins true

	router.Static("/images/", "./images")
	router.StaticFile("/css/output-css", "./css/output.css")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", index)
	router.GET("/index.html", index)
	router.GET("/chat-widget", chatWidget)
	router.GET("/chat-container", chatContainer)

	fmt.Println("Listening on http://localhost:8080")
	err = router.Run()
	if err != nil {
		fmt.Println(err)
	}
}
