package services

import (
	"myapp-go/src/models"
	"myapp-go/src/repositories"
	"myapp-go/src/utils"
	"myapp-go/src/utils/httpError"

	"time"

	"github.com/badoux/checkmail"
	"github.com/google/uuid"
)

type AuthService struct {
	UserRepo    *repositories.UserRepository
	SessionRepo *repositories.SessionRepository
}

func NewAuthService(userRepo *repositories.UserRepository, sessionRepo *repositories.SessionRepository) *AuthService {
	return &AuthService{
		UserRepo:    userRepo,
		SessionRepo: sessionRepo,
	}
}

// Buat user baru
func (s *AuthService) RegisterUser(username, email, password string) (*models.User, *models.Session, error) {
	// Validasi email
	if err := checkmail.ValidateFormat(email); err != nil {
		return nil, nil, httpError.InvalidEmail
	}

	// Cek jika user sudah ada
	if s.UserRepo.IsUserExists(username) {
		return nil, nil, httpError.UserAlreadyExists
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, nil, httpError.HashPasswordError
	}

	// Buat user baru
	user := &models.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
		ID:       uuid.New(),
	}

	// create database
	err = s.UserRepo.CreateUser(user)
	if err != nil {
		return nil, nil, httpError.CreateUserError
	}

	// create session
	session := &models.Session{
		UserRefer:    user.ID,
		Sessionid: uuid.New(),
		Expires: time.Now().Add(5 * 24 * time.Hour),
	}

	// save session
	err = s.SessionRepo.CreateSession(session)
	if err != nil {
		return nil, nil, httpError.CreateSessionError
	}

	return user, session, nil
}

// Buat user baru
func (s *AuthService) LoginUser(email, password string) (*models.User, *models.Session, error) {

	// Validasi email
	if err := checkmail.ValidateFormat(email); err != nil {
		return nil, nil, httpError.InvalidEmail
	}

	/// Find user by username
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return nil, nil, httpError.NotFound
	}

	// Check password
	if !utils.ComparePasswords(user.Password, []byte(password)) {
		return nil, nil, httpError.InvalidPassword
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID.String())
	if err != nil {
		return nil, nil, httpError.InvalidGenerateToken
	}

	// Create session
	session := &models.Session{
		UserRefer: user.ID,
		Expires:   utils.SessionExpires(),
		Sessionid: uuid.New(),
		Token: token,
	}

	// Save session
	if err := s.SessionRepo.UpdateOrCreateSession(session); err != nil {
		return nil, nil, err
	}
		
	return user, session, nil
}

// Logout user dengan menghapus session
func (s *AuthService) Logout(sessionID string) error {
	// Cek apakah session ada
	session, err := s.SessionRepo.GetSessionByID(sessionID)
	if err != nil {
		return httpError.NotFoundSession
	}

	// Hapus session
	return s.SessionRepo.DeleteSession(session.Sessionid.String())
}