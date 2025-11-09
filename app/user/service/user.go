package service

import (
	"context"
	"errors"
	"go-micro-todoList/app/user/repository/db/dao"
	"go-micro-todoList/app/user/repository/db/model"
	"go-micro-todoList/idl/pb"
	"go-micro-todoList/pkg/e"
	"sync"
)

type UserServ struct {
}

var UserSevIns *UserServ
var UserSrvOnce sync.Once

// GetUserServ Lazy Singleton Pattern. lazy-loading --> Eager Singleton Pattern
func GetUserServ() *UserServ {
	UserSrvOnce.Do(func() {
		UserSevIns = &UserServ{}
	})
	return UserSevIns
}

/*
// GetUserServEager Eager Singleton Pattern: cause problems of concurrency
func GetUserServEager() *UserServ {
	if UserSevIns == nil {
		return new(UserServ)
	}
	return UserSevIns
}

*/

func (u *UserServ) UserLogin(ctx context.Context, req *pb.UserRequest, resp *pb.UserResponse) (err error) {
	resp.Code = e.Success
	user, err := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err != nil {
		return
	}
	if user.ID == 0 {
		err = errors.New("User Not Found.")
		resp.Code = e.Error
		return
	}
	if !user.CheckPassword(req.Password) {
		err = errors.New("User Password Error.")
		resp.Code = e.Error
		return
	}

	resp.UserDetail = BuildUser(user)
	return

}

func (u *UserServ) UserRegister(ctx context.Context, req *pb.UserRequest, resp *pb.UserResponse) (err error) {
	resp.Code = e.Success
	if req.Password != req.PasswordConfirm {
		err = errors.New("User Password Error.")
		resp.Code = e.Error
		return
	}

	user, err := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err != nil {
		return
	}
	if user.ID > 0 {
		err = errors.New("User Name Already Exists.")
		resp.Code = e.Error
		return
	}

	user = &model.User{
		UserName: req.UserName,
	}
	// hashing
	if err = user.SetPassword(req.Password); err != nil {
		resp.Code = e.Error
		return
	}

	if err = dao.NewUserDao(ctx).CreateUser(user); err != nil {
		resp.Code = e.Error
		return
	}
	return
}

func BuildUser(item *model.User) *pb.UserModel {
	return &pb.UserModel{
		Id:        uint32(item.ID),
		UserName:  item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
		DeletedAt: 0,
	}
}
