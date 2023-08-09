package storage

import (
	"context"
	"pro-link-api/internal/model"
)

type (
	IAccountStorage interface {
		IStorage[*model.Account]
		FindByEmailOrName(ctx context.Context, email string, username string) (*model.Account, error)
		FindByVerificationCode(ctx context.Context, code string) (*model.Account, error)
	}

	AccountStorage struct {
		AbstractStorage[*model.Account]
	}
)

func NewAccountStorage(s *Storage) *AccountStorage {
	return &AccountStorage{
		AbstractStorage[*model.Account]{
			db:        s.db,
			tableName: model.AccountTableName,
		},
	}
}

func (s *AccountStorage) FindByEmailOrName(ctx context.Context, email string, username string) (result *model.Account, err error) {
	db := s.db.WithContext(ctx)

	if email != "" {
		db.Where("acc_email =?", email)
	}

	if username != "" {
		db.Where("acc_username =?", username)
	}

	err = db.Find(&result).Error

	return result, err
}

func (s *AccountStorage) FindByVerificationCode(ctx context.Context, code string) (result *model.Account, err error) {
	db := s.db.WithContext(ctx)

	err = db.Where("acc_verification_code = ?", code).Find(&result).Error

	return result, err
}
