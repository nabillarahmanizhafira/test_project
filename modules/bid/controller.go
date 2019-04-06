package bid

import (
	"github.com/nabillarahmanizhafira/test_project/models"
)

// Controller represent the bid's controller
type Controller interface {
	GetByID(string) (models.Product, error)
	SetProduct(string, string) error
}
