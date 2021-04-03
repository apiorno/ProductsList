package handlers

import (
	"github.com/apiorno/example/data"
	"net/http"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 	200: productsResponse

//GetProducts handles GET requests and returns all current products
func (p *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {

	p.l.Println("[DEBUG] get all records")

	prods := data.GetProducts()

	err := data.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}

// swagger:route GET /products/{id} products product
// Returns the product matching the id
// responses:
// 		200: productsResponse
// 		404: errorResponse

//GetProductsByID handles GET requests and returns the product matching the ID
func (p *Products) GetProductsByID(rw http.ResponseWriter, h *http.Request) {

	id := getProductID(r)

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}
