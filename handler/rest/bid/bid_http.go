package bid

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nabillarahmanizhafira/test_project/common/log"
	"github.com/nabillarahmanizhafira/test_project/common/response"
	b "github.com/nabillarahmanizhafira/test_project/modules/bid"
)

// HTTPBidHandler represent the httphandler for bid
type HTTPBidHandler struct {
	BController b.Controller
}

// GetProductHandler willreturn product with specific ID
func (h *HTTPBidHandler) GetProductHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var resp response.HandlerResponse
	queryValues := r.URL.Query()
	ID := queryValues.Get("id")
	res, err := h.BController.GetByID(ID)
	if err != nil {
		log.Error(err, "error while fetching the product", res)
		resp.StatusCode = http.StatusInternalServerError
		resp.Message = "error while fetching the product"
		response.WriteAPIStandard(w, resp, err)
		return
	}
	resp.StatusCode = http.StatusOK
	resp.Message = res
	response.WriteAPIStandard(w, resp, err)
	return
}

// ProductBidHandler will act as handler for setting bid
func (h *HTTPBidHandler) ProductBidHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var resp response.HandlerResponse
	r.ParseMultipartForm(2 << 23)

	productID := r.PostForm.Get("product_id")
	if productID == "" {
		err := fmt.Errorf("product id can't be empty")
		log.Error(err)
		resp.StatusCode = http.StatusBadRequest
		resp.Message = "Bad request"
		response.WriteAPIStandard(w, resp, err)
		return
	}
	value := r.PostForm.Get("value")
	if value == "" {
		err := fmt.Errorf("value can't be empty")
		log.Error(err)
		resp.StatusCode = http.StatusBadRequest
		resp.Message = "Bad request"
		response.WriteAPIStandard(w, resp, err)
		return
	}

	err := h.BController.SetProduct(productID, value)
	if err != nil {
		log.Error(err, "Error while setting the bid")
		resp.StatusCode = http.StatusInternalServerError
		resp.Message = "There's a problem during set the value of bid"
		response.WriteAPIStandard(w, resp, err)
		return
	}

	resp.StatusCode = http.StatusOK
	resp.Message = "success"
	response.WriteAPIStandard(w, resp, err)
	return
}
