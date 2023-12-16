package service

import "kim/internal/logic/dao"

type Service struct {
	UserService UserService
}

func InitService(dao *dao.Dao) *Service {
	return &Service{
		UserService: NewUserService(dao.UserDao),
	}
}
