package data

import (
	"encoding/json"
	"io"
	"log"
	"microservices/micro-service/coffee/utils"
	"net/http"
	"time"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products is a collection of Product
type Products []*Product

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

func GetNextId() int {
	if len(productList) == 0 {
		return 1
	}
	pl := productList[len(productList)-1]
	id := pl.ID + 1
	return id
}
func AddProduct(p *Product) {
	p.ID = GetNextId()
	p.CreatedOn = time.Now().UTC().String()
	p.UpdatedOn = time.Now().UTC().String()
	p.DeletedOn = ""
	productList = append(productList, p)
}
func UpdateProduct(p *Product) {
	for i := 0; i < len(productList); i++ {
		if productList[i].ID == p.ID {
			productList[i].Name = p.Name
			productList[i].Description = p.Description
			productList[i].Price = p.Price
			productList[i].SKU = p.SKU
			productList[i].UpdatedOn = time.Now().UTC().String()
		}
	}
	// for _, pr := range productList {
	// 	log.Println("ID: ", pr.ID)
	// 	log.Println("Name: ", pr.Name)
	// 	log.Println("Description: ", pr.Description)
	// 	log.Println("Price: ", pr.Price)
	// 	log.Println("SKU: ", pr.SKU)
	// 	log.Println("CreatedOn: ", pr.CreatedOn)
	// 	log.Println("UpdatedOn: ", pr.UpdatedOn)
	// 	log.Println("DeletedOn: ", pr.DeletedOn)
	// 	log.Println("-------------------------------")
	// }
	return
}

// FromJson De-serializes the contents of the collection from JSON
func (p *Product) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (prod *Product) HandleFromJson(rw http.ResponseWriter, data io.Reader) {
	err := prod.FromJson(data)
	if err != nil {
		log.Println("Error-here ", err)
		http.Error(rw, utils.InternalServerError, http.StatusInternalServerError)
		return
	}
}

func (prod *Products) HandleToJson(rw http.ResponseWriter) {
	err := prod.ToJson(rw)
	if err != nil {
		log.Println("Error-here ", err)
		http.Error(rw, utils.InternalServerError, http.StatusInternalServerError)
		return
	}
}

// productList is a hard coded list of products for this
// example data source
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
