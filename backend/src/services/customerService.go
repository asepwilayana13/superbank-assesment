package services

import (
	"errors"
	"fmt"
	"math"

	// "net/http"
	"time"

	"github.com/google/uuid"

	"myapp-go/src/dto"
	"myapp-go/src/models"
	"myapp-go/src/repositories"
	"myapp-go/src/utils"
	"myapp-go/src/utils/httpError"
)

type CustomerService struct {
	CustomerRepo    *repositories.CustomerRepository
	UserRepo    *repositories.UserRepository
	BankAccountRepo   *repositories.BankAccountRepository
}

func NewCustomerService(customerRepo *repositories.CustomerRepository, userRepo *repositories.UserRepository, bankAccountRepo *repositories.BankAccountRepository) *CustomerService {
	return &CustomerService{
		CustomerRepo:   customerRepo,
		UserRepo:   userRepo,
		BankAccountRepo: bankAccountRepo,
	}
}

// CREATE CUSTOMER
func (s *CustomerService) Create(req dto.CreateCustomerRequest) (*models.Customer, *models.User, error) {

	// Find By Email
	user, err := s.UserRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, nil, err
	}

	if user != nil {
		return nil, nil, fmt.Errorf("email %s sudah terdaftar", req.Email)
	}


	// change format date string to date
	dateOfBirth,  err  := utils.FormatDate(req.DateOfBirth)
	if err != nil {
		fmt.Println("Format tanggal salah:", err)
	} else {
		fmt.Println("Tanggal lahir:", dateOfBirth)
	}

	password := utils.GenerateRandomPassword(8)
	fmt.Println(password)

	// Hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, nil, httpError.HashPasswordError
	}

	newUser := &models.User{
		ID:       	 	uuid.New(),
		Username:    	req.FullName,
		Email: 				req.Email,
		Role: 				"user",
		Password:			hashedPassword,
		CreatedAt:   	time.Now(),
		UpdatedAt:   	time.Now(),
	}

	if err := s.UserRepo.CreateUser(newUser); err != nil {
		return nil, nil, err
	}
	
	newCustomer := &models.Customer{
		ID:       	 uuid.New(),
		UserID:      newUser.ID,
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
		DateOfBirth: dateOfBirth,
		Address:     req.Address,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.CustomerRepo.Create(newCustomer); err != nil {
		return nil, nil, err
	}

	newBankAccount := &models.BankAccount{
		ID:       	 		uuid.New(),
		CustomerID:  		newCustomer.ID,
		AccountNumber: 	utils.GenerateAccountNumber(),
		AccountType: 		req.AccountType,
		Balance: 				float64(req.Balance),
		Currency: 			req.Currency,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.BankAccountRepo.Create(newBankAccount); err != nil {
		return nil, nil, err
	}

	return newCustomer, newUser, nil
}

// UPDATE CUSTOMER
func (s *CustomerService) Update(req dto.UpdateCustomerRequest, userId string) (*models.Customer, error) {

	userUUID, err := uuid.Parse(userId)
	if err != nil {
    return nil, errors.New("invalid bank account format")
	}

	user, err := s.UserRepo.GetByUserID(userUUID) 
	if err != nil {
		return nil, err
	}

	customer, err := s.CustomerRepo.GetByUserID(userUUID) 
	if err != nil {
		return nil, err
	}

	// change format date string to date
	dateOfBirth,  err  := utils.FormatDate(req.DateOfBirth)
	if err != nil {
		fmt.Println("Format tanggal salah:", err)
	} else {
		fmt.Println("Tanggal lahir:", dateOfBirth)
	}

	updateUser := &models.User{
		ID:       	 	user.ID,
		Username:    	req.FullName,
		Email: 				req.Email,
		Role: 				"user",
		Password:			user.Password,
		UpdatedAt:   	time.Now(),
	}

	if err := s.UserRepo.Update(updateUser); err != nil {
		return nil, err
	}

	updateCustomer := &models.Customer{
		ID:       	 customer.ID,
		UserID:      customer.UserID,
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
		DateOfBirth: dateOfBirth,
		Address:     req.Address,
		UpdatedAt:   time.Now(),
	}

	if err := s.CustomerRepo.Update(updateCustomer); err != nil {
		return nil, err
	}

	newBankAccount := &models.BankAccount{
		ID:       	 		customer.BankAccount.ID,
		CustomerID:  		customer.ID,
		AccountNumber: 	customer.BankAccount.AccountNumber,
		AccountType: 		req.AccountType,
		Currency: 			req.Currency,
		Balance: 				customer.BankAccount.Balance,
		UpdatedAt:   		time.Now(),
	}

	if err := s.BankAccountRepo.Update(newBankAccount); err != nil {
		return nil, err
	}

	return updateCustomer, nil
}

// GET CUSTOMER BY USER ID
func (s *CustomerService) GetByUserID(user_id string) (*models.Customer, error) {
	userUUID, err := uuid.Parse(user_id)
	if err != nil {
    return nil, errors.New("invalid user ID format")
	}
	
	customer, err := s.CustomerRepo.GetByUserID(userUUID)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (s *CustomerService) GetAllCustomer(search string, limit, page int) (*models.CustomersList, error) {
	customer, total, err := s.CustomerRepo.GetAllCustomer(limit, page, search)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	hasMore := page < totalPages

	return &models.CustomersList {
		TotalCount: int(total),
		TotalPages: totalPages,
		Page:       page,
		Size:       limit,
		HasMore:    hasMore,
		Customer:  customer,
	}, nil
}
