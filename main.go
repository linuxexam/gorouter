package main

import (
	_ "github.com/linuxexam/gorouter/app1"
	_ "github.com/linuxexam/gorouter/app2"
	"github.com/linuxexam/gorouter/router"
	"log"
)

func main() {
	log.Fatal(router.Run(":8080"))
}
