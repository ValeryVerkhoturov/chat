package main

import (
	"github.com/ValeryVerkhoturov/chat/config"
	"github.com/ValeryVerkhoturov/chat/l10n"
	"github.com/ValeryVerkhoturov/chat/utils"
	"github.com/jritsema/gotoolbox/web"
	"net/http"
)

type DataWithLocale struct {
	Data        interface{}
	PublicUrl   string
	Locale      l10n.Locale
	LocaleName  string
	TelegramUrl string
}

func index(r *http.Request) *web.Response {
	locale, localeName := utils.GetLocale(r)

	return web.HTML(http.StatusOK, html, "index.html", DataWithLocale{
		PublicUrl:   config.PublicUrl,
		Locale:      locale,
		LocaleName:  localeName,
		TelegramUrl: config.TelegramUrl,
	}, nil)
}

// /GET chat-widget
func chatWidgetGet(r *http.Request) *web.Response {
	locale, localeName := utils.GetLocale(r)

	return utils.EmbeddedJS(http.StatusOK, html, "chat-widget.html", DataWithLocale{
		PublicUrl:   config.PublicUrl,
		Locale:      locale,
		LocaleName:  localeName,
		TelegramUrl: config.TelegramUrl,
	}, nil)
}

// /GET chat-container
func chatContainerGet(r *http.Request) *web.Response {
	locale, localeName := utils.GetLocale(r)

	return web.HTML(http.StatusOK, html, "chat-container.html", DataWithLocale{
		PublicUrl:   config.PublicUrl,
		Locale:      locale,
		LocaleName:  localeName,
		TelegramUrl: config.TelegramUrl,
	}, nil)
}
