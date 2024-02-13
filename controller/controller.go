package controller

import (
	"bytes"
	"embed"
	"github.com/ValeryVerkhoturov/chat/config"
	"github.com/ValeryVerkhoturov/chat/i18n"
	"github.com/ValeryVerkhoturov/chat/requestUtils"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)

var (
	//go:embed all:templates/*
	templateFS embed.FS

	//parsed templates
	html *template.Template
)

func init() {
	var err error
	html, err = requestUtils.TemplateParseFSRecursive(templateFS, ".html", true, nil)
	if err != nil {
		panic(err)
	}
}

type TemplateData struct {
	APIVersion  int
	Data        interface{}
	PublicUrl   string
	Locale      i18n.Locale
	LocaleName  string
	TelegramUrl string
}

func Index(c *gin.Context) {
	locale, localeName := requestUtils.GetLocale(c)

	c.HTML(http.StatusOK, "index.html", TemplateData{
		APIVersion:  1,
		PublicUrl:   config.PublicUrl,
		Locale:      locale,
		LocaleName:  localeName,
		TelegramUrl: config.TelegramUrl,
	})
}

func ChatWidgetV1(c *gin.Context) {
	locale, localeName := requestUtils.GetLocale(c)

	c.Header("Content-Type", "application/javascript")

	var buf bytes.Buffer
	err := html.ExecuteTemplate(&buf, "chat-widget.html", TemplateData{
		APIVersion:  1,
		PublicUrl:   config.PublicUrl,
		Locale:      locale,
		LocaleName:  localeName,
		TelegramUrl: config.TelegramUrl,
	})
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	var jsContent = requestUtils.WrapHTMLWithEmbeddingJS(buf)

	c.String(http.StatusOK, jsContent)
}

func ChatContainerV1(c *gin.Context) {
	locale, localeName := requestUtils.GetLocale(c)

	c.HTML(http.StatusOK, "chat-container.html", TemplateData{
		APIVersion:  1,
		PublicUrl:   config.PublicUrl,
		Locale:      locale,
		LocaleName:  localeName,
		TelegramUrl: config.TelegramUrl,
	})
}
