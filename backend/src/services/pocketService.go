package services

import (
	"errors"
	//"fmt"
	"math"
	"myapp-go/src/dto"
	"myapp-go/src/models"
	"myapp-go/src/repositories"
	"time"

	"github.com/google/uuid"
)

type PocketService struct {
	PocketRepo    *repositories.PocketRepository
	BankAccountRepo    *repositories.BankAccountRepository
}

func NewPocketService(pocketRepo *repositories.PocketRepository, bankAccountRepo *repositories.BankAccountRepository) *PocketService {
	return &PocketService{
		PocketRepo:   pocketRepo,
		BankAccountRepo:   bankAccountRepo,
	}
}

// CREATE POCKET
func (s *PocketService) Create(req dto.CreatePocketRequest, user_id string) (*models.Pocket, error) {

	bankAccountUUID, err := uuid.Parse(req.BankAccountID)
	if err != nil {
    return nil, errors.New("invalid Bank Account Id format")
	}

	bankAccount, err := s.BankAccountRepo.GetBankAccountByID(bankAccountUUID) 
	if err != nil {
		return nil, err
	}

	newPocket := &models.Pocket{
		ID:       	 uuid.New(),
		BankAccountID: bankAccount.ID,
		PocketName: req.PocketName,
		Balance: req.Balance,
		GoalAmount: req.GoalAmount,
		BankAccount: *bankAccount,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.PocketRepo.Create(newPocket); err != nil {
		return nil, err
	}

	return newPocket, nil
}

// GET ALL POCKET
func (s *PocketService) GetAll(bank_account_id string, limit, page int) (*models.PocketList, error) {
	bankAccountUUID, err := uuid.Parse(bank_account_id)
	if err != nil {
    return nil, errors.New("invalid user ID format")
	}

	pocket, total, err := s.PocketRepo.GetAll(bankAccountUUID, limit, page)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	hasMore := page < totalPages

	return &models.PocketList {
		TotalCount: int(total),
		TotalPages: totalPages,
		Page:       page,
		Size:       limit,
		HasMore:    hasMore,
		Pockets:    pocket,
	}, nil
}

// Update Pocket
func (s *PocketService) Update(req dto.UpdatePocketRequest, pocketId string) (*models.Pocket, error) {

	pocketUUID, err := uuid.Parse(pocketId)
	if err != nil {
    return nil, errors.New("invalid user ID format")
	}

	pocket, err := s.PocketRepo.GetPocketByID(pocketUUID) 
	if err != nil {
		return nil, err
	}


	bankAccount, err := s.BankAccountRepo.GetBankAccountByID(pocket.BankAccountID) 
	if err != nil {
		return nil, err
	}

	newPocket := &models.Pocket{
		ID:       	 pocket.ID,
		BankAccountID: bankAccount.ID,
		PocketName: req.PocketName,
		Balance: req.Balance,
		GoalAmount: req.GoalAmount,
		BankAccount: *bankAccount,
		UpdatedAt:   time.Now(),
	}

	if err := s.PocketRepo.Update(newPocket); err != nil {
		return nil, err
	}

	return newPocket, nil
}

// GET ONE POCKET BY ID
func (s *PocketService) GetOneById(pocketId string) (*models.Pocket, error) {
	// Konversi pocketId dari string ke UUID
	pocketUUID, err := uuid.Parse(pocketId)
	if err != nil {
    return nil, errors.New("invalid bank account format")
	}

	pocket, err := s.PocketRepo.GetPocketByID(pocketUUID) 
	if err != nil {
		return nil, err
	}

	return pocket, nil
}

// DELETE POCKET
func (s *PocketService) Delete(pocketId string) (error) {
	// Konversi pocketId dari string ke UUID
	pocketUUID, err := uuid.Parse(pocketId)
	if err != nil {
    return errors.New("invalid bank account format")
	}

	bankAccount, err := s.PocketRepo.GetPocketByID(pocketUUID) 
	if err != nil {
		return err
	}

	return s.PocketRepo.DeletePocket(bankAccount.ID.String())
}
