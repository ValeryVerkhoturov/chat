package main

import (
	"bytes"
	"github.com/ValeryVerkhoturov/chat/config"
	"github.com/ValeryVerkhoturov/chat/l10n"
	"github.com/ValeryVerkhoturov/chat/requestUtils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type DataWithLocale struct {
	Data        interface{}
	PublicUrl   string
	Locale      l10n.Locale
	LocaleName  string
	TelegramUrl string
}

func index(c *gin.Context) {
	locale, localeName := requestUtils.GetLocale(c)

	c.HTML(http.StatusOK, "index.html", DataWithLocale{
		PublicUrl:   config.PublicUrl,
		Locale:      locale,
		LocaleName:  localeName,
		TelegramUrl: config.TelegramUrl,
	})
}

func chatWidget(c *gin.Context) {
	locale, localeName := requestUtils.GetLocale(c)

	c.Header("Content-Type", "application/javascript")

	var buf bytes.Buffer
	err := html.ExecuteTemplate(&buf, "chat-widget.html", DataWithLocale{
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

func chatContainer(c *gin.Context) {
	locale, localeName := requestUtils.GetLocale(c)

	c.HTML(http.StatusOK, "chat-container.html", DataWithLocale{
		PublicUrl:   config.PublicUrl,
		Locale:      locale,
		LocaleName:  localeName,
		TelegramUrl: config.TelegramUrl,
	})
}
