package dao

import "kim/internal/logic/global"

var tables = []interface{}{
	User{}, FriendApplication{}, Friend{},
}

type Dao struct {
	UserDao              *UserDao
	FriendApplicationDao *FriendApplicationDao
	FriendDao            *FriendDao
}

func InitDao() *Dao {
	err := global.Database.AutoMigrate(tables...)
	if err != nil {
		panic(err)
	}
	return &Dao{
		UserDao:              NewUserDao(),
		FriendApplicationDao: NewFriendApplicationDao(),
		FriendDao:            NewFriendDao(),
	}
}
