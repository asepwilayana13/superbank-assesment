package repositories

import (
	"errors"
	"myapp-go/src/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// FINDBY USERNAME
func (r *UserRepository) IsUserExists(user_id string) bool {
	var user models.User
	result := r.DB.Where("id = ?", user_id).First(&user)
	return result.RowsAffected > 0
}

// Get by ID
func (r *UserRepository) GetByUserID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.DB.Where("id = ?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FIND BY EMAIL
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Email belum terdaftar
		}
		return nil, err // Query error lain
	}

	return &user, nil // Email ditemukan
}

// CREATE USER
func (r *UserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

// Get ALL USERS WITH PAGINATION
func (r *UserRepository) GetAllUsers(limit, page int) ([]*models.User, int64, error) {
	var users []*models.User
	var total int64

	// Hitung total users
	if err := r.DB.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Ambil data users dengan limit dan offset
	offset := (page - 1) * limit
	if err := r.DB.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// Update User
func (r *UserRepository) Update(user *models.User) error {
	var existingUser models.User

	// Cek apakah customer dengan ID tersebut ada
	if err := r.DB.First(&existingUser, "id = ?", user.ID).Error; err != nil {
		return errors.New("user tidak ditemukan")
	}

	// Update data customer
	return r.DB.Model(&existingUser).Updates(user).Error
}
