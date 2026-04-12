package repository

import (
	"context"
	"user-service/src/common/model"
	"gorm.io/gorm"
)

type User interface{
	Find(ctx context.Context, whereClause string, args [] any) (*model.User, error)
}

type userImpl	struct{
	db *gorm.DB
}

func NewUser(db *gorm.DB) User {
	return &userImpl{db: db}
}

func (r *userImpl) Find(ctx context.Context, whereClause string, args [] any) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where(whereClause, args...).First(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if user.ID == "" {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}