package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Heyy World!!")

		a, error := ioutil.ReadAll(r.Body)

		if error != nil {
			http.Error(rw, "sorry", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Data %s", a)
	})

	http.HandleFunc("/seeyou", func(http.ResponseWriter, *http.Request) {
		log.Println("see you soon")
	})

	http.ListenAndServe(":8080", nil)
}