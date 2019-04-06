package bid

import (
	"net/http"
	"strconv"

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
	res, err := h.BController.GetByID(1)
	if err != nil {
		log.Error(err, "error while fetching the product", res)
		resp.StatusCode = http.StatusInternalServerError
		resp.Message = "error while fetching the product"
		response.WriteAPIStandard(w, resp, err)
		return
	}
	resp.StatusCode = http.StatusOK
	resp.Message = res
	return
}

// ProductBidHandler will act as handler for setting bid
func (h *HTTPBidHandler) ProductBidHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var resp response.HandlerResponse
	r.ParseMultipartForm(2 << 23)

	productID, err := strconv.Atoi(r.PostForm.Get("product_id"))
	if err != nil {
		log.Error(err, "error while parse the productID", productID)
		resp.StatusCode = http.StatusBadRequest
		resp.Message = "Bad request"
		response.WriteAPIStandard(w, resp, err)
		return
	}
	value, err := strconv.Atoi(r.PostForm.Get("value"))
	if err != nil {
		log.Error(err, "error while parse the product", value)
		resp.StatusCode = http.StatusBadRequest
		resp.Message = "Bad request"
		response.WriteAPIStandard(w, resp, err)
		return
	}

	err = h.BController.SetProduct(productID, value)
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

// func bidProductHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
// 	r.ParseMultipartForm(2 << 23)

// 	//TODO: Convert ke int
// 	productID := r.PostForm.Get("product_id")
// 	price := r.PostForm.Get("price")

// 	err := bid.BidProduct(productID, price)
// 	if err != nil {
// 		w.WriteHeader(500)
// 		w.Write([]byte("Internal Server Error"))
// 	}

// 	w.Write([]byte("OK"))

// 	return nil
// }
