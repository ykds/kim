package service

import "kim/internal/logic/dao"

type Service struct {
	UserService              UserService
	FriendService            FriendService
	FriendApplicationService FriendApplicationService
}

func InitService(dao *dao.Dao) *Service {
	return &Service{
		UserService:              NewUserService(dao.UserDao),
		FriendService:            NewFriendService(dao.FriendDao, dao.UserDao),
		FriendApplicationService: NewFriendApplicationService(dao.FriendApplicationDao, dao.FriendDao),
	}
}
