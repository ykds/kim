package dao

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"kim/internal/logic/global"
	"kim/internal/pkg/errors"
)

const (
	ApplyStatusWait   = 1
	ApplyStatusAgree  = 1 << 1
	ApplyStatusReject = 1 << 2
)

type FriendApplication struct {
	gorm.Model
	UserId   uint   `json:"user_id" gorm:"user_id"`
	FriendId uint   `json:"friend_id" gorm:"friend_id"`
	Status   int8   `json:"status" gorm:"status"`
	Desc     string `json:"desc" gorm:"desc"`
	Question string `json:"ask" gorm:"question"`
	Answer   string `json:"answer" gorm:"answer"`
}

func (fa *FriendApplication) TableName() string {
	return "tbl_friendapplication"
}

type FriendApplicationDao struct{}

func NewFriendApplicationDao() *FriendApplicationDao {
	return &FriendApplicationDao{}
}

func (fad *FriendApplicationDao) CreateApplication(app *FriendApplication) error {
	err := global.Database.Create(app).Error
	if err != nil {
		err = errors.Wrap(err, "创建好友申请失败")
	}
	return err
}

func (fad *FriendApplicationDao) GetApplication(userId, friendId uint) (*FriendApplication, error) {
	app := &FriendApplication{}
	err := global.Database.First(app, "user_id=? AND friend_id=?", userId, friendId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return app, nil
		}
		err = errors.Wrap(err, "获取好友申请失败")
	}
	return app, err
}

func (fad *FriendApplicationDao) GetApplicationById(id uint) (*FriendApplication, error) {
	app := &FriendApplication{}
	err := global.Database.First(app, "id=?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return app, nil
		}
		err = errors.Wrap(err, "获取好友申请失败")
	}
	return app, err
}

func (fad *FriendApplicationDao) ListApplication(userId uint) ([]*FriendApplication, error) {
	apps := make([]*FriendApplication, 0)
	err := global.Database.Find(apps, "user_id=?", userId).Error
	if err != nil {
		err = errors.Wrap(err, "获取好友申请列表失败")
	}
	return apps, err
}

func (fad *FriendApplicationDao) UpdateApplication(id uint, update map[string]interface{}) error {
	return fad.UpdateApplicationTx(global.Database.DB, id, update)
}

func (fad *FriendApplicationDao) UpdateApplicationTx(tx *gorm.DB, id uint, update map[string]interface{}) error {
	err := tx.Where("id=?", id).Updates(update).Error
	if err != nil {
		err = errors.Wrap(err, "更新好友申请失败")
	}
	return err
}

type Friend struct {
	gorm.Model
	UserId   uint `json:"user_id" gorm:"user_id;uniqueIndex:friendship1,priority:1;uniqueIndex:friendship2,priority:2"`
	FriendId uint `json:"friend_id" gorm:"friend_id;uniqueIndex:friendship1,priority:2;uniqueIndex:friendship2,priority:1"`
}

func (f *Friend) TableName() string {
	return "tbl_friend"
}

type FriendDao struct{}

func NewFriendDao() *FriendDao {
	return &FriendDao{}
}

func (fd *FriendDao) CreateFriendTx(tx *gorm.DB, friend *Friend) error {
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "friend_id"}},
		DoNothing: true,
	}).Create(&[]*Friend{{UserId: friend.FriendId, FriendId: friend.UserId}, friend}).Error
	if err != nil {
		err = errors.Wrap(err, "添加好友关系失败")
	}
	return err
}

func (fd *FriendDao) DeleteFriend(userId uint, friendId uint) error {
	err := global.Database.Delete(&Friend{}, "(user_id=? AND friend_id=?) OR (friend_id=? AND user_id=?)", userId, friendId, userId, friendId).Error
	if err != nil {
		err = errors.Wrap(err, "删除好友失败")
	}
	return err
}

func (fd *FriendDao) ListFriends(userId uint) ([]*Friend, error) {
	friends := make([]*Friend, 0)
	err := global.Database.Find(friends, "user_id=?", userId).Error
	if err != nil {
		err = errors.Wrap(err, "")
	}
	return friends, err
}
