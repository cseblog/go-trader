package web

import (
	"common"
	"exchange/internal"
	"html/template"
	"net/http"
	"path/filepath"
)

type empty struct{}

var templatePath = "../src/exchange/web/html/"
var templates = []string{"welcome", "sessions", "instruments"}

var exchange = &internal.TheExchange
var t *template.Template

func StartWebServer(addr string) {
	var err error
	var paths []string

	for _, file := range templates {
		paths = append(paths, filepath.Join(templatePath, file+".html"))
	}
	t, err = template.ParseFiles(paths...)
	if err != nil {
		panic(err)
	}
	go func() {
		http.HandleFunc("/instruments", instrumentsHandler)
		http.HandleFunc("/sessions", sessionsHandler)
		http.HandleFunc("/", welcomeHandler)
		http.ListenAndServe(addr, nil)
	}()
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "welcome.html", empty{})
}

func sessionsHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["Sessions"] = exchange.ListSessions()

	t.ExecuteTemplate(w, "sessions.html", data)
}
func instrumentsHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["Symbols"] = common.IMap.AllSymbols()

	t.ExecuteTemplate(w, "instruments.html", data)
}
