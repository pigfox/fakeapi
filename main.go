package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	Domain   string
	Protocol string
	Port     string
	Slug     string
}

type Payload struct {
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	CCNumber     string  `json:"cc_number"`
	CCExpiration string  `json:"cc_expiration"`
	CCCVV        string  `json:"cc_cvv"`
	Amount       float32 `json:"amount"`
}

type Response struct {
	Result string `json:"result"`
}

func main() {
	config := Config{Domain: "fakeapi.com", Protocol: "http://", Port: ":6666", Slug: "/"}
	router := mux.NewRouter()
	router.HandleFunc(config.Slug, func(w http.ResponseWriter, r *http.Request) {
		d := json.NewDecoder(r.Body)
		p := &Payload{}
		err := d.Decode(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			r := Response{Result: "success"}
			res, err := json.Marshal(r)
			if err != nil {
				fmt.Println(err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write(res)
		}
		fmt.Println(p)
	}).Methods(http.MethodPost)

	fmt.Println("Server running on " + config.Protocol + config.Domain + config.Port)

	err := http.ListenAndServe(config.Port, router)
	if err != nil {
		log.Fatal(err)
	}
}
