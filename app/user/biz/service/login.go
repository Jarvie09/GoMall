package service

import (
	"GoMall/app/user/biz/dal/mysql"
	"GoMall/app/user/model"
	user "GoMall/rpc_gen/kitex_gen/user"
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	fmt.Println("到达User登录方法")
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("Email or Password is empty")
	}
	row, err := model.GetByEmail(mysql.DB, req.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	resp = &user.LoginResp{
		UserId: int32(row.ID),
	}
	return resp, nil
}
