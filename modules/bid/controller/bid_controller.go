package controller

import (
	"github.com/nabillarahmanizhafira/test_project/models"
	"github.com/nabillarahmanizhafira/test_project/modules/bid"
)

type bidController struct {
	bidRepository bid.Repository
}

// NewBidController return an instance of bidController
func NewBidController(bRepo bid.Repository) bid.Controller {
	return &bidController{
		bidRepository: bRepo,
	}
}

func (bc *bidController) GetByID(ID int) (res models.Product, err error) {
	return
}

func (bc *bidController) SetProduct(ID, value int) (err error) {
	return
}
