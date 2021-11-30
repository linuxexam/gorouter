package app2

import (
	"github.com/linuxexam/gorouter/router"
	"fmt"
	"net/http"
)

func init() {
	router.RegisterApp("app2", http.HandlerFunc(handleFunc))
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, from app2")
}
