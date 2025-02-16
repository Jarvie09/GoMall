package service

import (
	"GoMall/app/user/biz/dal/mysql"
	"GoMall/app/user/model"
	"fmt"

	user "GoMall/rpc_gen/kitex_gen/user"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	fmt.Println("到达User注册方法")
	fmt.Println(req)
	// Finish your business logic.
	if req.Email == "" || req.Password == "" || req.PasswordConform == "" {
		return nil, errors.New("Email or Password is empty")
	}
	fmt.Println("0")
	if req.Password != req.PasswordConform {
		fmt.Println("6666")
		return nil, errors.New("password not match")
	}
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	fmt.Println("1")
	if err != nil {
		return nil, err
	}
	fmt.Println("2")
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(passwordHashed),
	}

	err = model.Creat(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}
	fmt.Println("3")
	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
