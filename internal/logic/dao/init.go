package dao

import "kim/internal/logic/global"

var tables = []interface{}{
	User{},
}

type Dao struct {
	UserDao *UserDao
}

func InitDao() *Dao {
	err := global.Database.AutoMigrate(tables...)
	if err != nil {
		panic(err)
	}
	return &Dao{
		UserDao: NewUserDao(),
	}
}
