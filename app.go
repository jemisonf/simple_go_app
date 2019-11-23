package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

var random_generator = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	http.HandleFunc("/healthcheck", healthCheck)
	http.HandleFunc("/message", message)

	log.Print("starting server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func message(w http.ResponseWriter, r *http.Request) {
	log.Print("request at /message")
	messages := []string{
		// "Hello world",
		// "It's a lovely day",
		// "It's not a lovely day",
		"What are you still doing here?",
	}

	message := messages[random_generator.Int()%len(messages)]

	_, err := w.Write([]byte(message))
	if err != nil {
		log.Fatal(err.Error())
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	log.Print("request at /healthcheck")
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Fatal(err.Error())
	}
}
