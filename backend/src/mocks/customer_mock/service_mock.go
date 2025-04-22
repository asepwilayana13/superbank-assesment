package customer_mock

import (
	"myapp-go/src/dto"
	"myapp-go/src/models"
	"github.com/stretchr/testify/mock"
)

type CustomerServiceMock struct {
	mock.Mock
}

func (m *CustomerServiceMock) Create(req dto.CreateCustomerRequest) (*models.Customer, *models.User, error) {
	args := m.Called(req)
	return args.Get(0).(*models.Customer), args.Get(1).(*models.User), args.Error(2)
}

func (m *CustomerServiceMock) Update(req dto.UpdateCustomerRequest, userId string) (*models.Customer, error) {
	args := m.Called(req, userId)
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *CustomerServiceMock) GetByUserID(userId string) (*models.Customer, error) {
	args := m.Called(userId)
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *CustomerServiceMock) GetAllCustomer(search string, limit, page int) (*models.CustomersList, error) {
	args := m.Called(search, limit, page)
	return args.Get(0).(*models.CustomersList), args.Error(1)
}



