   
package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"data"
)

// items are a http.Handler
type Items struct {
	l *log.Logger
}


func NewItems(l *log.Logger) *Items {
	return &Items{l}
}

//INTERFACE

func (i *Items) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// request handling
	if r.Method == http.MethodGet {
		i.getItems(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		i.addItem(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		i.l.Println("Push", r.URL.Path)
		
		register := registerexp.MustCompile(`/([0-9]+)`)
		reg := register.FindAllStringSubmatch(r.URL.Path, -1)

		if len(reg) != 1 {
			i.l.Println("Invalid , more than one id")
			http.Error(rw, "Invalid", http.StatusBadRequest)
			return
		}

		if len(reg[0]) != 2 {
			i.l.Println("Invalid , more than one capture group")
			http.Error(rw, "Invalid", http.StatusBadRequest)
			return
		}

		id_String := reg[0][1]
		id, error := strconv.Atoi(id_String)
		if error != nil {
			i.l.Println("Invalid, unable to convert into number", id_String)
			http.Error(rw, "Invalid", http.StatusBadRequest)
			return
		}

		i.updateItems(id, rw, r)
		return
	}

	
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

//get and return

func (i *Items) getItems(rw http.ResponseWriter, r *http.Request) {
	i.l.Println("GET items")

	
	li := data.GetItems()

	
	error := li.ToJSON(rw)
	if error != nil {
		http.Error(rw, "Unable", http.StatusInternalServerError)
	}
}

func (i *Items) addItem(rw http.ResponseWriter, r *http.Request) {
	i.l.Println("Handle POST Product")

	it := &data.Item{}

	error := it.FromJSON(r.Body)
	if error != nil {
		http.Error(rw, "Unable", http.StatusBadRequest)
	}

	data.Additem(it)
}

func (i Items) updateItems(id int, rw http.ResponseWriter, r *http.Request) {
	i.l.Println("Push item")

	it := &data.item{}

	error := it.FromJSON(r.Body)
	if error != nil {
		http.Error(rw, "Unable", http.StatusBadRequest)
	}

	error = data.Updateitem(id, it)
	if error == data.ErrorItemNotFound {
		http.Error(rw, "Item not found", http.StatusNotFound)
		return
	}

	if error != nil {
		http.Error(rw, "Item not found", http.StatusInternalServerError)
		return
	}
}
