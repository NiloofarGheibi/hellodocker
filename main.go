package main

import (
	"github.com/gotschmarcel/goserv"
	"log"
	"net/http"
)

type model struct {
	Message string `"json:message"`
}
func main() {

	server := goserv.NewServer()

	// Handle Get Request
	server.Get("/message", func(w http.ResponseWriter, r *http.Request) {
		data := model{Message:"Hello Docker"}
		if err := goserv.WriteJSON(w, data); err != nil {
			goserv.Context(r).Error(err, http.StatusInternalServerError)
			return
		}
	})

	log.Fatalln(server.Listen(":60234"))
}