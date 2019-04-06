package bid

import (
	"github.com/nabillarahmanizhafira/test_project/models"
)

// Repository represent the bid's repository contract
type Repository interface {
	GetByID(int) (models.Product, error)
	SetProduct(int, int) error
}
