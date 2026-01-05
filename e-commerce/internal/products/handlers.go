package products

import (
	"log"
	"net/http"

	"github.com/agamyo168/e-commerce/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service:service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request){
	//Return JSON in HTTP Response
	products, err := h.service.ListProducts((r.Context()))
	if err != nil{
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	json.Write(w, http.StatusOK, products)
}

func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request){
	var createProductParams createProductParams; 
	json.Read(r,&createProductParams)

	products, err := h.service.CreateProduct((r.Context()),createProductParams)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return 
	} 
	json.Write(w, http.StatusOK, products)
}