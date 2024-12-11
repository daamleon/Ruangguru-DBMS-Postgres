package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)
}

type sessionsRepoImpl struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (s *sessionsRepoImpl) AddSessions(session model.Session) error {
	err := s.db.Create(&session).Error
	return err
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	err := s.db.Where("token = ?", token).Delete(&model.Session{}).Error
	return err
}

func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	err := s.db.Model(&model.Session{}).Where("username = ?", session.Username).Updates(session).Error
	return err
}

func (s *sessionsRepoImpl) SessionAvailName(Username string) error {
	var session model.Session
	err := s.db.Where("username = ?", Username).First(&session).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *sessionsRepoImpl) SessionAvailToken(Token string) (model.Session, error) {
	var session model.Session
	err := s.db.Where("token = ?", Token).First(&session).Error
	if err != nil {
		return model.Session{}, err
	}
	return session, nil
}

