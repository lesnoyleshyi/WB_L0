package handlers

import (
	"L0/internal/services"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
)

type Handler struct {
	*services.Service
	NatsSub *NatsSubscription
}

const createTemplatePath = "/api/internal/handlers/templates/create.html"

const htmlTempl = `<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>L0 Product info</title>
	<style>
		.myDiv {
			border: 2px outset black;
			width: fit-content;
			padding: 10px;
			background: whitesmoke;
			text-align: left;
		}
	</style>
</head>
<body>
<h3>Input order Id</h3>
<form method="POST">
	<label>Order Id</label><br>
	<input type="text" name="id" /><br><br>
	<input type="submit" value="Get info" /><br>
</form><br>
<label>RESULT</label><br>
<div class="myDiv">
	<p>{{ .}}</p>
</div>
</body>
</html>`

func New(service *services.Service) *Handler {
	sub, err := NewNatsSubscription(service)
	if err != nil {
		log.Fatalf("Can't establish subscription. Service won't start: %v", err)
	}
	return &Handler{service, sub}
}

func (h Handler) Routes() chi.Router {
	r := chi.NewRouter()

	r.HandleFunc("/", h.htmlResp)

	return r
}

func (h Handler) htmlResp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//We will return template in any case: even if we doesn't have such order
		tmpl, err := template.New("templ").Parse(htmlTempl)
		if err != nil {
			log.Println("template goes wrong:", err)
			return
		}
		//Parse order id from form
		if err := r.ParseForm(); err != nil {
			log.Println("niceP:", err)
		}
		id := r.PostFormValue("id")
		//Get order from cache
		order, ok := h.CacheService.Get(id)
		if !ok || id == "" {
			err := tmpl.Execute(w, "No order with such id or empty id")
			if err != nil {
				log.Println("can't execute template:", err)
			}
			return
		}

		if err := tmpl.Execute(w, order); err != nil {
			log.Println("can't execute template:", err)
		}
	} else {
		http.ServeFile(w, r, createTemplatePath)
	}
}

func (h Handler) getOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	order, ok := h.CacheService.Get(id)
	if !ok {
		log.Println("sosi jopu")
	} else {
		log.Printf("%v\n", *order)
	}
}
