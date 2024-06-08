package handlers

import (
	"log"
	"microservices/micro-service/coffee/data"
	"microservices/micro-service/coffee/utils"
	"net/http"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) Products {
	return Products{l}
}

func (p Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		 p.GetProducts(rw, r)
	case http.MethodPost:
		 p.AddProducts(rw, r)
	case http.MethodPut:
		 p.UpdateProduct(rw, r)
	default: 
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
	// method is Get
	// if r.Method == http.MethodGet {
	// 	p.GetProducts(rw, r)
	// 	return
	// }
	// if r.Method == http.MethodPost {
	// 	p.AddProducts(rw, r)
	// 	return
	// }
	// if r.Method == http.MethodPut {
	// 	p.UpdateProduct(rw, r)
	// 	return
	// }
	// if method is not handled
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	pl := data.GetProducts()
	pl.HandleToJson(rw)
	// err := pl.ToJson(rw)
	// if err != nil {
	// 	p.l.Println(err)
	// 	http.Error(rw, utils.InternalServerError, http.StatusInternalServerError)
	// 	return
	// }
	return
}

func (p *Products) GetProduct(rw http.ResponseWriter, r *http.Request) {
	strId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(strId)
	pl, _ := data.GetProduct(id)
	pl.ToJson(rw)
	return
}
func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	prod.HandleFromJson(rw, r.Body)
	// err := prod.FromJson(r.Body)
	// if err != nil {
	// 	p.l.Println("Error-here ", err)
	// 	http.Error(rw, utils.InternalServerError, http.StatusInternalServerError)
	// 	return
	// }
	data.AddProduct(prod)
	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(utils.ProductAdded))
	return
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	prod.HandleFromJson(rw, r.Body)
	// err := prod.FromJson(r.Body)
	// if err != nil {
	// 	log.Println("Error-here ", err)
	// 	http.Error(rw, utils.InternalServerError, http.StatusInternalServerError)
	// 	return
	// }

	data.UpdateProduct(prod)
	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(utils.ProductUpdated))
	return
}
