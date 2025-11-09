package dao

import (
	"context"
	"go-micro-todoList/app/user/repository/db/model"

	"gorm.io/gorm"
)

// define CRUD operations on user model in DB
type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) FindUserByUserName(userName string) (r *model.User, err error) {
	err = dao.Model(&model.User{}).
		Where("user_name = ?", userName).Find(&r).Error
	// record not found response
	// Find: SELECT * FROM user WHERE user_name = xxx ORDER BY id LIMIT 1
	// First/Last: SELECT * FROM user WHERE user_name = xxx
	return
}

func (dao *UserDao) CreateUser(in *model.User) error {
	return dao.Model(&model.User{}).Create(&in).Error
}
