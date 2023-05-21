package main

import (
	"embed"
	_ "embed"
	"encoding/json"
	"log"
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed i18n/active.*.json
var LocaleFS embed.FS

func main() {
	bundle := i18n.NewBundle(language.Japanese)
	bundle.RegisterUnmarshalFunc("toml", json.Unmarshal)
	must := func(_ *i18n.MessageFile, err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	must(bundle.LoadMessageFileFS(LocaleFS, "i18n/active.ja.json"))
	must(bundle.LoadMessageFileFS(LocaleFS, "i18n/active.en.json"))

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		langQuery := r.URL.Query().Get("lang")
		accept := r.Header.Get("Accept-Language")

		localizer := i18n.NewLocalizer(bundle, langQuery, accept)
		l, err := localizer.Localize(&i18n.LocalizeConfig{
			MessageID:    "Greet",
			TemplateData: map[string]string{"Name": "John"},
		})
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(l))
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
