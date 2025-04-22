package services_interface

import (
	"myapp-go/src/dto"
	"myapp-go/src/models"
)

type ICustomerService interface {
	Create(req dto.CreateCustomerRequest) (*models.Customer, *models.User, error)
	Update(req dto.UpdateCustomerRequest, id string) (*models.Customer, error)
	GetByUserID(userID string) (*models.Customer, error)
	GetAllCustomer(search string, limit, page int) (*models.CustomersList, error)
}