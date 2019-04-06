package bid

import (
	"github.com/nabillarahmanizhafira/test_project/models"
)

// Controller represent the bid's controller
type Controller interface {
	GetByID(int) (models.Product, error)
	SetProduct(int, int) error
}
