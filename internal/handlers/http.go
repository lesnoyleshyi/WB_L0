package handlers

import (
	myTemplates "L0/internal/handlers/templates"
	"html/template"
	"log"
	"net/http"
)

func (h Handler) htmlResp(w http.ResponseWriter, r *http.Request) {
	//We will return template in any case
	tmpl, err := template.New("heh").Parse(myTemplates.GetOrderById)
	if err != nil {
		log.Println("template goes wrong:", err)
		return
	}
	if r.Method == "POST" {
		//Parse order id from form
		if err := r.ParseForm(); err != nil {
			log.Println("can't parse form:", err)
		}
		id := r.PostFormValue("id")
		//Get order from cache
		order, ok := h.CacheService.Get(id)
		if !ok || id == "" {
			err := tmpl.Execute(w, "No order with such id or empty id")
			if err != nil {
				log.Println("htmlResp can't execute template for POST:", err)
			}
			return
		}
		if err := tmpl.Execute(w, order); err != nil {
			log.Println("htmlResp can't execute template for POST:", err)
		}
	} else {
		if err := tmpl.Execute(w, "order data will be here"); err != nil {
			log.Println("htmlResp can't execute template for GET:", err)
		}
	}
}
