package service

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"kim/internal/logic/dao"
	"kim/internal/logic/errcode"
	"kim/internal/logic/global"
	"kim/internal/pkg/errors"
	"kim/internal/pkg/jwt"
)

var _ UserService = userService{}

type RegisterReq struct {
	MobilePhone     string `json:"mobile_phone"`
	UserName        string `json:"user_name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	VerifyCode      string `json:"verify_code"`
}
type RegisterResp struct{}

type LoginReq struct {
	MobilePhone string `json:"mobile_phone"`
	Password    string `json:"password"`
}
type LoginResp struct {
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
}

type LogoutReq struct {
	UserId uint
}
type LogoutResp struct{}

type UserService interface {
	Register(req RegisterReq) (RegisterResp, error)
	Login(req LoginReq) (LoginResp, error)
	Logout(req LogoutReq) (LogoutResp, error)
}

type userService struct {
	userDao *dao.UserDao
}

func NewUserService(userDao *dao.UserDao) UserService {
	return &userService{userDao: userDao}
}

func (u userService) Register(req RegisterReq) (RegisterResp, error) {
	if req.MobilePhone == "" {
		return RegisterResp{}, errcode.MobileEmptyErr
	}
	user, err := u.userDao.GetUserByMobile(req.MobilePhone)
	if err != nil {
		return RegisterResp{}, err
	}
	if user.ID != 0 {
		return RegisterResp{}, errcode.RegisteredErr
	}
	if req.Password != req.ConfirmPassword {
		return RegisterResp{}, errcode.PasswordNotSameErr
	}
	salt := randStr()
	_, err = u.userDao.CreateUser(dao.User{
		UserName:    req.UserName,
		MobilePhone: req.MobilePhone,
		Avatar:      "",
		Password:    hashPassword(req.Password, salt),
		Salt:        salt,
	})
	return RegisterResp{}, err
}

func randStr() string {
	randBytes := make([]byte, 8)
	_, err := rand.Read(randBytes)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", randBytes)
}

func hashPassword(password string, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	return string(hash.Sum(nil))
}

func comparePassword(pwd1 string, pwd2 string, salt string) bool {
	return hashPassword(pwd1, salt) == pwd2
}

func (u userService) Login(req LoginReq) (LoginResp, error) {
	if req.MobilePhone == "" {
		return LoginResp{}, errcode.MobileEmptyErr
	}
	user, err := u.userDao.GetUserByMobile(req.MobilePhone)
	if err != nil {
		return LoginResp{}, err
	}
	if user.ID == 0 {
		return LoginResp{}, errcode.NotRegisterErr
	}
	if !comparePassword(req.Password, user.Password, user.Salt) {
		return LoginResp{}, errcode.PasswordNotMatchErr
	}

	token, err := jwt.NewToken(user.ID)
	if err != nil {
		err = errors.Wrap(err, "登录失败")
		return LoginResp{}, err
	}
	if addOnlineUser(user.ID) != nil {
		err = errors.Wrap(err, "登录失败")
		return LoginResp{}, err
	}
	return LoginResp{
		UserId: user.ID,
		Token:  token,
	}, nil
}

func (u userService) Logout(req LogoutReq) (LogoutResp, error) {
	if err := rmOnlineUser(req.UserId); err != nil {
		err = errors.Wrap(err, "登录失败")
		return LogoutResp{}, err
	}
	return LogoutResp{}, nil
}

func addOnlineUser(userId uint) error {
	return global.Redis.SetBit(context.Background(), "online_users", int64(userId), 1).Err()
}

func rmOnlineUser(userId uint) error {
	return global.Redis.SetBit(context.Background(), "online_users", int64(userId), 0).Err()
}