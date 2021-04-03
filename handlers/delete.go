package handlers

import (
	"net/http"
	"strconv"

	"github.com/apiorno/example/data"
	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// DeleteProduct deletes a product from the database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {

	id := getProductID(r)

	p.l.Println("[DEBUG] deleting record id", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] deleting record", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
