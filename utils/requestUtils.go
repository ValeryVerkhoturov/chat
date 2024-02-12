package utils

import (
	"bytes"
	"fmt"
	"github.com/ValeryVerkhoturov/chat/l10n"
	"github.com/jritsema/gotoolbox/web"
	"html/template"
	"log"
	"net/http"
)

// EmbeddedJS renders an js embedded script to a web response
func EmbeddedJS(status int, t *template.Template, template string, data interface{}, headers web.Headers) *web.Response {

	//render template to buffer
	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, template, data); err != nil {
		log.Println(err)
		return web.Empty(http.StatusInternalServerError)
	}

	var jsScript = fmt.Sprintf(`
(function() {
	var wrapper = document.createElement("div");
	wrapper.innerHTML = unescape(`+"`%s`"+`);
	document.body.appendChild(wrapper);

	// Move all scripts from wrapper to real script elements
	Array.from(wrapper.querySelectorAll('script')).forEach(function(oldScript) {
		var newScript = document.createElement('script');

		Array.from(oldScript.attributes).forEach(function(attr) {
			newScript.setAttribute(attr.name, attr.value);
		});

		if (oldScript.src) {
			newScript.src = oldScript.src;
		} else {
			newScript.textContent = oldScript.textContent;
		}
	
		oldScript.parentNode.replaceChild(newScript, oldScript);
	});
})();
        `, &buf)

	//m := minify.New()
	//m.AddFunc("application/javascript", js.Minify)
	//
	//if err := m.Minify("application/javascript", &buf, bytes.NewBufferString(jsScript)); err != nil {
	//	log.Fatalf("Failed to minify JavaScript: %v", err)
	//}
	var jsReader = bytes.NewBufferString(jsScript)

	return &web.Response{
		Status:      status,
		ContentType: "application/javascript",
		Content:     jsReader,
		Headers:     headers,
	}
}

func GetLocale(r *http.Request) (l10n.Locale, string) {
	localeName := "ru"
	queryValues := r.URL.Query()
	lang := queryValues.Get("lang")

	locale, ok := l10n.LocalesMap[lang]
	if ok {
		localeName = lang
	} else {
		locale = l10n.LocalesMap[localeName]
	}
	return locale, localeName
}
