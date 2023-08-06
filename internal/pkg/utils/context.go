package utils

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type CtxKey int

const (
	DbTrx     = "db_trx"
	UserEmail = "email"
	UserId    = "user_id"
	Uuid      = "uuid"
)

func GetTrx(ctx context.Context) (*gorm.DB, error) {
	tx, ok := ctx.Value(DbTrx).(*gorm.DB)
	if !ok {
		return nil, errors.New("can not get tx value from context")
	}
	return tx, nil
}

func GetUserId(ctx context.Context) (int, error) {
	userId, ok := ctx.Value(UserId).(int)
	if !ok {
		return 0, errors.New("can not get user id value from context")
	}

	return userId, nil
}

func GetUserIdAndTrx(ctx context.Context) (*gorm.DB, int, error) {
	tx, err := GetTrx(ctx)
	if err != nil {
		return nil, 0, err
	}

	userId, err := GetUserId(ctx)
	if err != nil {
		return nil, 0, err
	}

	return tx, userId, nil
}
