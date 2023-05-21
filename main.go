package main

import (
	"embed"
	_ "embed"
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed i18n/active.*.toml
var LocaleFS embed.FS

func main() {
	bundle := i18n.NewBundle(language.Japanese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	_, err := bundle.LoadMessageFileFS(LocaleFS, "i18n/active.ja.toml")
	if err != nil {
		log.Fatal(err)
	}
	bundle.LoadMessageFileFS(LocaleFS, "i18n/active.en.toml")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		langFromValue := r.FormValue("lang")
		langQuery := r.URL.Query().Get("lang")
		accept := r.Header.Get("Accept-Language")
		fmt.Printf("`%s` `%s` `%s`\n", langFromValue, langQuery, accept)
		localizer := i18n.NewLocalizer(bundle, langFromValue, langQuery, accept)
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
