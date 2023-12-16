package dao

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"kim/internal/logic/global"
)

type User struct {
	gorm.Model
	UserName    string
	MobilePhone string
	Avatar      string
	Password    string `json:"password"`
	Salt        string `json:"salt"`
}

func (u *User) TableName() string {
	return "tbl_users"
}

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (u *UserDao) GetUser(id uint) (*User, error) {
	user := &User{}
	err := global.Database.First(user, "id=?", id).Error
	if err != nil {
		err = errors.Wrap(err, "获取用户失败")
	}
	return user, err
}

func (u *UserDao) GetUserByMobile(mobile string) (*User, error) {
	user := &User{}
	err := global.Database.First(user, "mobile_phone=?", mobile).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, nil
		}
		err = errors.Wrap(err, "获取用户失败")
	}
	return user, err
}

func (u *UserDao) CreateUser(user User) (uint, error) {
	err := global.Database.Create(&user).Error
	if err != nil {
		err = errors.Wrap(err, "创建用户失败")
	}
	return user.ID, err
}
