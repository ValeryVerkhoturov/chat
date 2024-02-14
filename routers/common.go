package routers

import (
	"embed"
	"github.com/ValeryVerkhoturov/chat/config"
	"github.com/ValeryVerkhoturov/chat/utils/requestUtils"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

var (
	//go:embed all:templates/*
	TemplateFS embed.FS

	// HTML parsed templates
	HTML *template.Template
)

func init() {
	var err error
	HTML, err = requestUtils.TemplateParseFSRecursive(TemplateFS, ".html", true, nil)
	if err != nil {
		panic(err)
	}
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
