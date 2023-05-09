package main

import (
	"log"
	"net/http"

	"github.com/ducduong-c0x12c/cproxy"
)

func main() {
	handler := cproxy.New()
	log.Println("Listening on:", "*:8080")
	_ = http.ListenAndServe(":8080", handler)
}
