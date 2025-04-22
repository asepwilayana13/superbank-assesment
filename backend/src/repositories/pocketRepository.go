package repositories

import (
	"errors"
	"myapp-go/src/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PocketRepository struct {
	DB *gorm.DB
}

func NewPocketRepository(db *gorm.DB) *PocketRepository {
	return &PocketRepository{DB: db}
}

// CREATE POCKETE
func (r *PocketRepository) Create(pocket *models.Pocket) error {
	return r.DB.Create(pocket).Error
}

// GET ALL POCKET WITH BANK ACCOUNT
func (r *PocketRepository) GetAll(bankAccount uuid.UUID, limit, page int) ([]*models.Pocket, int64, error) {
	var pocket []*models.Pocket
	var total int64

	if err := r.DB.Model(&models.Pocket{}).
		Joins("JOIN bank_accounts ON bank_accounts.id = pockets.bank_account_id").
		Where("pockets.bank_account_id = ?", bankAccount).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Ambil data bank account dengan join ke customer
	offset := (page - 1) * limit
	if err := r.DB.Preload("BankAccount.Customer").
		Joins("JOIN bank_accounts ON bank_accounts.id = pockets.bank_account_id").
		Where("pockets.bank_account_id = ?", bankAccount).
		Limit(limit).Offset(offset).
		Find(&pocket).Error; err != nil {
		return nil, 0, err
	}

	return pocket, total, nil
}

//UPDATE POCKET
func (r *PocketRepository) Update(pocket *models.Pocket) error {
	var existingPocket models.Pocket

	// Cek apakah customer dengan ID tersebut ada
	if err := r.DB.First(&existingPocket, "id = ?", pocket.ID).Error; err != nil {
		return errors.New("customer tidak ditemukan")
	}

	// Update data customer
	return r.DB.Model(&existingPocket).Updates(pocket).Error
}

// GET POCKET BY ID
func (r *PocketRepository) GetPocketByID(id uuid.UUID) (*models.Pocket, error) {
	var pocket models.Pocket
	err := r.DB.Preload("BankAccount.Customer").
				Joins("JOIN bank_accounts ON bank_accounts.id = pockets.bank_account_id").
				Where("pockets.id = ?", id).
				First(&pocket).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // Tidak ditemukan, return nil tanpa error
	}
	if err != nil {
		return nil, err
	}
	return &pocket, nil
}

// DELETE POCKET
func (r *PocketRepository) DeletePocket(id string) error {
	return r.DB.Where("id = ?", id).Delete(&models.Pocket{}).Error
}


