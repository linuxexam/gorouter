package router

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	appsMu sync.RWMutex
	apps   = make(map[string]App)
)

type App struct {
	UrlPattern  string
	HttpHandler http.Handler
}

func RegisterApp(name string, handler http.Handler) {
	appsMu.Lock()
	defer appsMu.Unlock()
	if handler == nil {
		panic("web: nil is not valid handler.")
	}
	if _, dup := apps[name]; dup {
		panic("web: Register called twice for handler " + name)
	}
	app := App{"/" + name + "/", handler}
	apps[name] = app
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	for name, app := range apps {
		fmt.Fprintln(w, name+" is serving "+app.UrlPattern)
	}
}

func Run(addr string) error {
	http.HandleFunc("/", rootHandler)
	for _, app := range apps {
		http.Handle(app.UrlPattern, app.HttpHandler)
	}
	return http.ListenAndServe(addr, nil)
}
