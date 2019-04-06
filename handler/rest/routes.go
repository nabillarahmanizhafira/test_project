package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nabillarahmanizhafira/test_project/handler/rest/bid"
	bcon "github.com/nabillarahmanizhafira/test_project/modules/bid/controller"
	brepo "github.com/nabillarahmanizhafira/test_project/modules/bid/repository"
)

//InitRoutes set all routes here
func InitRoutes(router *httprouter.Router) {
	// standard
	router.GET("/ping", ping)
	router.MethodNotAllowed = http.HandlerFunc(notfound)

	// bid
	handlerBid := &bid.HTTPBidHandler{
		BController: bcon.NewBidController(brepo.NewBidRedis()),
	}
	router.GET("/get-product", handlerBid.GetProductHandler)
	router.POST("/bid-product", handlerBid.ProductBidHandler)
}

func ping(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("You're connected with me :))"))
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("404 Not Found"))
}
