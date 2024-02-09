package service_test

import (
	"testTaskMedods/config"
	"testTaskMedods/internal/domain"
	"testTaskMedods/internal/service"
	"testing"

	"github.com/google/uuid"
)

type mockRepository struct{}

func (m *mockRepository) GetAllSessions() ([]domain.Session, error) {
	return []domain.Session{}, nil
}

func (m *mockRepository) Create(session domain.Session) error {
	return nil
}

func (m *mockRepository) GetSessionById(sessionId uuid.UUID) (domain.Session, error) {
	return domain.Session{}, nil
}

func (m *mockRepository) Update(session domain.Session) error {
	return nil
}

func (m *mockRepository) DeleteSessionById(sessionId uuid.UUID) error {
	return nil
}

func TestCreateSession(t *testing.T) {
	mockRepo := &mockRepository{}
	conf := config.Config{} // Вам нужно определить конфигурацию
	service := service.NewService(mockRepo, conf)

	guid := uuid.New()
	tokenPair, err := service.Create(guid)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if tokenPair.AccessToken == "" {
		t.Error("Expected non-empty access token")
	}

	if len(tokenPair.RefreshToken) == 0 {
		t.Error("Expected non-empty refresh token")
	}
}

func TestUpdateSession(t *testing.T) {
	mockRepo := &mockRepository{}
	conf := config.Config{} // Вам нужно определить конфигурацию
	service := service.NewService(mockRepo, conf)

	sessionId := uuid.New()
	tokenPair, err := service.Update(sessionId)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if tokenPair.AccessToken == "" {
		t.Error("Expected non-empty access token")
	}

	if len(tokenPair.RefreshToken) == 0 {
		t.Error("Expected non-empty refresh token")
	}
}

func TestGetAllSessions(t *testing.T) {
	mockRepo := &mockRepository{}
	conf := config.Config{} // Вам нужно определить конфигурацию
	service := service.NewService(mockRepo, conf)

	sessions, err := service.GetAllSessions()

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(sessions) != 0 {
		t.Errorf("Expected empty slice, but got: %v", sessions)
	}
}
