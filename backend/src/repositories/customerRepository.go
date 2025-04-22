package repositories

import (
	"errors"
	"myapp-go/src/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{DB: db}
}

// CREATE CUSTOMER
func (r *CustomerRepository) Create(customer *models.Customer) error {
	return r.DB.Create(customer).Error
}

// Update CUSTOMER
func (r *CustomerRepository) Update(customer *models.Customer) error {
	var existingCustomer models.Customer

	if err := r.DB.First(&existingCustomer, "id = ?", customer.ID).Error; err != nil {
		return errors.New("customer tidak ditemukan")
	}

	// Update data customer
	return r.DB.Model(&existingCustomer).Updates(customer).Error
}

// GET CUSTOMER BY ID
func (r *CustomerRepository) GetByUserID(userID uuid.UUID) (*models.Customer, error) {
	var customer models.Customer
	err := r.DB.Preload("BankAccount").
				Where("user_id = ?", userID).First(&customer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

// GET CUSTOMER BY EMAIL
func (r *CustomerRepository) GetByEmail(email string) (*models.Customer, error) {
	var customer models.Customer
	err := r.DB.Where("email = ?", email).First(&customer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

// GET ALL CUSTOMER WITH PAGINATION
func (r *CustomerRepository) GetAllCustomer(limit, page int, search string) ([]*models.Customer, int64, error) {
	var customers []*models.Customer
	var total int64

	query := r.DB.Model(&models.Customer{}).Where("full_name ILIKE ? OR address ILIKE ?", "%"+search+"%", "%"+search+"%")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Preload("BankAccount").
		Limit(limit).
		Offset(offset).
		Find(&customers).Error; err != nil {
		return nil, 0, err
	}
	return customers, total, nil
}



