package services

import (
	"math"
	"myapp-go/src/models"
	"myapp-go/src/repositories"
)

type UserService struct {
	UserRepo    *repositories.UserRepository
	SessionRepo *repositories.SessionRepository
}

func NewUserService(userRepo *repositories.UserRepository, sessionRepo *repositories.SessionRepository) *UserService {
	return &UserService{
		UserRepo:    userRepo,
		SessionRepo: sessionRepo,
	}
}

// gET ALL USER WITH PAGINATION
func (s *UserService) GetAllUsers(limit, page int) (*models.UsersList, error) {
	users, total, err := s.UserRepo.GetAllUsers(limit, page)
	if err != nil {
		return nil, err
	}

	// Hitung total halaman
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	// Tentukan apakah masih ada halaman berikutnya
	hasMore := page < totalPages

	return &models.UsersList {
		TotalCount: int(total),
		TotalPages: totalPages,
		Page:       page,
		Size:       limit,
		HasMore:    hasMore,
		Users:      users,
	}, nil
}
