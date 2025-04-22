package services

import (
	"errors"
	"math"
	"myapp-go/src/dto"
	"myapp-go/src/models"
	"myapp-go/src/repositories"
	"time"

	"github.com/google/uuid"
)

type BankAccountService struct {
	BankAccountRepo    *repositories.BankAccountRepository
	CustomerRepo    *repositories.CustomerRepository
}

func NewBankAccountService(bankAccountRepo *repositories.BankAccountRepository, customerRepo *repositories.CustomerRepository) *BankAccountService {
	return &BankAccountService{
		BankAccountRepo:   bankAccountRepo,
		CustomerRepo:   customerRepo,
	}
}

// Create customer
func (s *BankAccountService) Create(req dto.CreateBARequest, user_id string) (*models.BankAccount, error) {

	// Konversi userID dari string ke UUID
	userUUID, err := uuid.Parse(user_id)
	if err != nil {
    return nil, errors.New("invalid user ID format")
	}

	customer, err := s.CustomerRepo.GetByUserID(userUUID) 
	if err != nil {
		return nil, err
	}

	newBankAccount := &models.BankAccount{
		ID:       	 uuid.New(),
		CustomerID:  customer.ID,
		AccountNumber: req.AccountNumber,
		AccountType: req.AccountType,
		Balance: req.Balance,
		Currency: req.Currency,
		Customer:  customer,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.BankAccountRepo.Create(newBankAccount); err != nil {
		return nil, err
	}

	return newBankAccount, nil
}

func (s *BankAccountService) GetAllBankAccount(user_id string, limit, page int) (*models.BankAccountList, error) {

	userUUID, err := uuid.Parse(user_id)
	if err != nil {
    return nil, errors.New("invalid user ID format")
	}


	bankAccount, total, err := s.BankAccountRepo.GetAll(userUUID, limit, page)
	if err != nil {
		return nil, err
	}

	// Hitung total halaman
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	// Tentukan apakah masih ada halaman berikutnya
	hasMore := page < totalPages

	return &models.BankAccountList {
		TotalCount: int(total),
		TotalPages: totalPages,
		Page:       page,
		Size:       limit,
		HasMore:    hasMore,
		BankAccount:    bankAccount,
	}, nil
}

// Update Bank Account
func (s *BankAccountService) Update(req dto.UpdateBARequest, user_id string, bankAccountId string) (*models.BankAccount, error) {

	// Konversi userID dari string ke UUID
	userUUID, err := uuid.Parse(user_id)
	if err != nil {
    return nil, errors.New("invalid user ID format")
	}

	customer, err := s.CustomerRepo.GetByUserID(userUUID) 
	if err != nil {
		return nil, err
	}

	// Konversi bankAccountId dari string ke UUID
	bankAccountUUID, err := uuid.Parse(bankAccountId)
	if err != nil {
    return nil, errors.New("invalid bank account format")
	}

	bankAccount, err := s.BankAccountRepo.GetBankAccountByID(bankAccountUUID) 
	if err != nil {
		return nil, err
	}

	newBankAccount := &models.BankAccount{
		ID:       	 bankAccount.ID,
		CustomerID:  bankAccount.CustomerID,
		AccountNumber: req.AccountNumber,
		AccountType: req.AccountType,
		Balance: req.Balance,
		Currency: req.Currency,
		Customer: customer,
		UpdatedAt:   time.Now(),
	}

	if err := s.BankAccountRepo.Update(newBankAccount); err != nil {
		return nil, err
	}

	return newBankAccount, nil
}

func (s *BankAccountService) GetOne(bankAccountId string) (*models.BankAccount, error) {

	bankAccountUUID, err := uuid.Parse(bankAccountId)
	if err != nil {
    return nil, errors.New("invalid bank account format")
	}

	bankAccount, err := s.BankAccountRepo.GetBankAccountByID(bankAccountUUID) 
	if err != nil {
		return nil, err
	}

	return bankAccount, nil
}

func (s *BankAccountService) Delete(bankAccountId string) (error) {
	
	bankAccountUUID, err := uuid.Parse(bankAccountId)
	if err != nil {
    return errors.New("invalid bank account format")
	}

	bankAccount, err := s.BankAccountRepo.GetBankAccountByID(bankAccountUUID) 
	if err != nil {
		return err
	}

	return s.BankAccountRepo.DeleteBankAccount(bankAccount.ID.String())
}
