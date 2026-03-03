package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/your-moon/grape-shared/entities"
	"gorm.io/gorm"
)

type refreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *refreshTokenRepository {
	return &refreshTokenRepository{db: db}
}

func (r *refreshTokenRepository) Create(token *entities.RefreshTokenUUID) error {
	return r.db.Create(token).Error
}

func (r *refreshTokenRepository) GetByToken(token string) (*entities.RefreshTokenUUID, error) {
	var refreshToken entities.RefreshTokenUUID
	err := r.db.Where("token = ? AND revoked = ? AND expires_at > ?", token, false, time.Now()).First(&refreshToken).Error
	if err != nil {
		return nil, err
	}
	return &refreshToken, nil
}

func (r *refreshTokenRepository) Revoke(token string) error {
	return r.db.Model(&entities.RefreshTokenUUID{}).Where("token = ?", token).Update("revoked", true).Error
}

func (r *refreshTokenRepository) RevokeAllForUser(userID uuid.UUID) error {
	return r.db.Model(&entities.RefreshTokenUUID{}).Where("user_id = ? AND revoked = ?", userID, false).Update("revoked", true).Error
}
