package core

import (
	"context"
	"errors"
	"user/model"
	"user/services"

	"github.com/jinzhu/gorm"
)

func BuildUser(item model.User) *services.UserModel {
	userModel := services.UserModel{
		ID:        uint32(item.ID),
		UserName:  item.Username,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &userModel
}

func (*UserService) UserLogin(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	var user model.User
	resp.Code = 200
	if err := model.DB.Where("user_name=?", req.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			resp.Code = 400
			return nil
		}
		resp.Code = 500
		return nil
	}
	if user.CheckPassword(req.Password) == false {
		resp.Code = 400
		return nil
	}
	resp.UserDetail = BuildUser(user)
	return nil
}

func (*UserService) UserRegister(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	if req.Password != req.PasswordConfirm {
		err := errors.New("the password is incorrect")
		return err
	}
	count := 0
	if err := model.DB.Model(&model.User{}).Where("user_name=?", req.UserName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		err := errors.New("the account is already exist")
		return err
	}
	user := model.User{
		Username: req.UserName,
	}
	if err := user.SetPassword(req.Password); err != nil {
		return err
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return err
	}
	resp.UserDetail = BuildUser(user)
	return nil

}
