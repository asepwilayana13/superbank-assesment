package repositories

import (
	"myapp-go/src/models"

	"gorm.io/gorm"
)

type SessionRepository struct {
	DB *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{DB: db}
}

// SAVE NEW SESSION
func (r *SessionRepository) CreateSession(session *models.Session) error {
	return r.DB.Create(session).Error
}

// UPDATE OR CREATE SESSION
func (r *SessionRepository) UpdateOrCreateSession(session *models.Session) error {
	var existingSession models.Session

	result := r.DB.Where("user_refer = ?", session.UserRefer).First(&existingSession)
	if result.RowsAffected > 0 {
		existingSession.Sessionid = session.Sessionid
		existingSession.Expires = session.Expires
		existingSession.Token = session.Token
		return r.DB.Save(&existingSession).Error
	}

	// Jika tidak ada, buat session baru
	return r.DB.Create(session).Error
}

// FIND SESSION BY ID
func (r *SessionRepository) GetSessionByID(sessionID string) (*models.Session, error) {
	var session models.Session
	result := r.DB.Where("sessionid = ?", sessionID).First(&session)
	if result.Error != nil {
		return nil, result.Error
	}
	return &session, nil
}

// DELETE SESSION (LOGOUT)
func (r *SessionRepository) DeleteSession(sessionID string) error {
	return r.DB.Where("sessionid = ?", sessionID).Delete(&models.Session{}).Error
}
