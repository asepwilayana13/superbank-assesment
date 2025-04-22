package repositories

import (
	//"errors"
	"errors"
	"myapp-go/src/models"

	// "github.com/google/uuid"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BankAccountRepository struct {
	DB *gorm.DB
}

func NewBankAccountRepository(db *gorm.DB) *BankAccountRepository {
	return &BankAccountRepository{DB: db}
}

// CREATE CUSTOMER
func (r *BankAccountRepository) Create(bankaccount *models.BankAccount) error {
	return r.DB.Create(bankaccount).Error
}

func (r *BankAccountRepository) GetAll(user_id uuid.UUID, limit, page int) ([]*models.BankAccount, int64, error) {
	var bankAccounts []*models.BankAccount
	var total int64

	if err := r.DB.Model(&models.BankAccount{}).
		Joins("JOIN customers ON customers.id = bank_accounts.customer_id").
		Where("customers.user_id = ?", user_id).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.DB.Preload("Customer").
		Joins("JOIN customers ON customers.id = bank_accounts.customer_id").
		Where("customers.user_id = ?", user_id).
		Limit(limit).Offset(offset).
		Find(&bankAccounts).Error; err != nil {
		return nil, 0, err
	}

	return bankAccounts, total, nil
}

//Update CUSTOMER
func (r *BankAccountRepository) Update(customer *models.BankAccount) error {
	var existingBankAccount models.BankAccount

	if err := r.DB.First(&existingBankAccount, "id = ?", customer.ID).Error; err != nil {
		return errors.New("customer tidak ditemukan")
	}

	// Update data customer
	return r.DB.Model(&existingBankAccount).Updates(customer).Error
}

// GET BANK ACCOUNT BY ID
func (r *BankAccountRepository) GetBankAccountByID(id uuid.UUID) (*models.BankAccount, error) {
	var bankAccount models.BankAccount
	err := r.DB.Preload("Customer").Preload("Pockets").Preload("TermDeposits").
				Where("bank_accounts.id = ?", id).
				First(&bankAccount).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// Menghitung total balance dari pockets
	totalPocketBalance := 0
	for _, pocket := range bankAccount.Pockets {
		totalPocketBalance += int(pocket.Balance)
	}
	bankAccount.TotalBalance.TotalPocketBalance = float64(totalPocketBalance)

	// Menghitung total deposite dari termdeposite
	totalDepositeBalance := 0
	for _, deposite := range bankAccount.TermDeposits {
		totalDepositeBalance += int(deposite.DepositAmount)
	}
	bankAccount.TotalBalance.TotalDepositeBalance = float64(totalDepositeBalance)

	bankAccount.TotalBalance.TotalBalance = bankAccount.Balance + float64(totalDepositeBalance) + float64(totalPocketBalance)

	return &bankAccount, nil
}

// DELETE BANK ACCOUNT
func (r *BankAccountRepository) DeleteBankAccount(id string) error {
	return r.DB.Where("id = ?", id).Delete(&models.BankAccount{}).Error
}


