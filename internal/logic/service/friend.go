package service

import (
	"context"
	"encoding/json"
	"kim/internal/logic/dao"
	"kim/internal/logic/errcode"
	"kim/internal/logic/global"
	"kim/internal/protocol"
	"time"
)

type FriendApplyReq struct {
	UserId   uint   `json:"user_id"`
	FriendId uint   `json:"friend_id"`
	Desc     string `json:"desc"`
}
type FriendApplyResp struct{}

type UpdateApplyReq struct {
	Id     uint `json:"id"`
	Status int8 `json:"status"`
}
type UpdateApplyResp struct {
}

type ApplicationQuestReq struct {
	Id       uint   `json:"id"`
	Question string `json:"question"`
}
type ApplicationQuestResp struct{}

type ApplicationAnsReq struct {
	Id     uint   `json:"id"`
	Answer string `json:"answer"`
}
type ApplicationAnsResp struct{}

type ApplicationInfo struct {
	UserId   uint   `json:"user_id"`
	FriendId uint   `json:"friend_id"`
	Status   int8   `json:"status"`
	Desc     string `json:"desc"`
	Question string `json:"ask"`
	Answer   string `json:"answer"`
}

type FriendApplicationService interface {
	FriendApply(req FriendApplyReq) (FriendApplyResp, error)
	UpdateFriendApplicationStatus(userId uint, req UpdateApplyReq) (UpdateApplyResp, error)
	ApplicationQuest(userId uint, req ApplicationQuestReq) (ApplicationQuestResp, error)
	ApplicationAns(userId uint, req ApplicationAnsReq) (ApplicationAnsResp, error)
	ListApplication(userId uint) ([]*ApplicationInfo, error)
}

type DeleteFriendReq struct {
	UserId   uint `json:"user_id"`
	FriendId uint `json:"friend_id"`
}
type DeleteFriendResp struct{}

type FriendService interface {
	ListFriends(userId uint) ([]*UserInfo, error)
	DeleteFriend(DeleteFriendReq) (DeleteFriendResp, error)
}

type friendApplicationService struct {
	friendApplicationDao *dao.FriendApplicationDao
	friendDao            *dao.FriendDao
}

func (f friendApplicationService) ListApplication(userId uint) ([]*ApplicationInfo, error) {
	application, err := f.friendApplicationDao.ListApplication(userId)
	if err != nil {
		return nil, err
	}
	result := make([]*ApplicationInfo, 0)
	for _, app := range application {
		result = append(result, &ApplicationInfo{
			UserId:   app.UserId,
			FriendId: app.FriendId,
			Status:   app.Status,
			Desc:     app.Desc,
			Question: app.Question,
			Answer:   app.Answer,
		})
	}
	return result, nil
}

func (f friendApplicationService) FriendApply(req FriendApplyReq) (FriendApplyResp, error) {
	isFriend, err := f.friendDao.IsFriend(req.UserId, req.FriendId)
	if err != nil {
		return FriendApplyResp{}, err
	}
	if isFriend {
		return FriendApplyResp{}, errcode.HadBeFriendErr
	}
	application, err := f.friendApplicationDao.GetApplication(req.UserId, req.FriendId)
	if err != nil {
		return FriendApplyResp{}, err
	}
	if application.ID != 0 {
		return FriendApplyResp{}, errcode.DuplicatedFriendApplyErr
	}
	app := &dao.FriendApplication{
		UserId:   req.UserId,
		FriendId: req.FriendId,
		Status:   dao.ApplyStatusWait,
		Desc:     req.Desc,
	}
	err = f.friendApplicationDao.CreateApplication(app)
	if err != nil {
		return FriendApplyResp{}, err
	}
	go func() {
		serverId, err := getUser(req.FriendId)
		if err != nil {
			return
		}
		notfi := protocol.Notification{
			ServerId:  serverId,
			UserId:    req.FriendId,
			Type:      protocol.NewFriendApplication,
			Content:   app,
			Timestamp: time.Now().UnixMilli(),
		}
		body, err := json.Marshal(notfi)
		if err != nil {
			return
		}
		err = global.Channel.PublishMsg(context.Background(), global.NotificationExchangeName, global.NotificationRoutingKey, body)
		if err != nil {
			global.Logger.Errorf("投递消息失败, error: %v", err)
		}
	}()
	return FriendApplyResp{}, nil
}

func (f friendApplicationService) UpdateFriendApplicationStatus(userId uint, req UpdateApplyReq) (UpdateApplyResp, error) {
	application, err := f.friendApplicationDao.GetApplicationById(req.Id)
	if err != nil {
		return UpdateApplyResp{}, err
	}
	if application.ID == 0 {
		return UpdateApplyResp{}, errcode.FriendApplyNotFoundErr
	}
	if userId == application.UserId {
		return UpdateApplyResp{}, errcode.CantHandleApplyErr
	}
	switch req.Status {
	case dao.ApplyStatusAgree:
		tx := global.Database.Begin()
		err = f.friendApplicationDao.UpdateApplicationTx(tx, req.Id, map[string]interface{}{"status": req.Status})
		if err != nil {
			tx.Rollback()
			return UpdateApplyResp{}, err
		}
		err = f.friendDao.CreateFriendTx(tx, &dao.Friend{
			UserId:   application.UserId,
			FriendId: application.FriendId,
		})
		if err != nil {
			tx.Rollback()
			return UpdateApplyResp{}, err
		}
		tx.Commit()
	case dao.ApplyStatusReject:
		err = f.friendApplicationDao.UpdateApplication(req.Id, map[string]interface{}{"status": req.Status})
	default:
		return UpdateApplyResp{}, errcode.ApplyStatusWrongErr
	}
	go func() {
		serverId, err := getUser(application.UserId)
		if err != nil {
			return
		}
		notfi := protocol.Notification{
			ServerId:  serverId,
			UserId:    application.UserId,
			Type:      protocol.FriendApplicationStatusChange,
			Content:   map[string]interface{}{"id": req.Id, "status": req.Status},
			Timestamp: time.Now().UnixMilli(),
		}
		body, err := json.Marshal(notfi)
		if err != nil {
			global.Logger.Errorf("好友申请状态通知-序列化失败，ID: %d, err: %v", req.Id, err)
			return
		}
		err = global.Channel.PublishMsg(context.Background(), global.NotificationExchangeName, global.NotificationRoutingKey, body)
		if err != nil {
			global.Logger.Errorf("好友申请状态通知-发送失败，ID: %d, err: %v", req.Id, err)
		}
	}()
	return UpdateApplyResp{}, nil
}

func (f friendApplicationService) ApplicationQuest(userId uint, req ApplicationQuestReq) (ApplicationQuestResp, error) {
	application, err := f.friendApplicationDao.GetApplicationById(req.Id)
	if err != nil {
		return ApplicationQuestResp{}, err
	}
	if application.ID == 0 {
		return ApplicationQuestResp{}, errcode.FriendApplyNotFoundErr
	}
	if userId == application.UserId {
		return ApplicationQuestResp{}, errcode.CantHandleApplyErr
	}
	err = f.friendApplicationDao.UpdateApplication(req.Id, map[string]interface{}{"question": req.Question})
	if err != nil {
		return ApplicationQuestResp{}, err
	}
	go func() {
		serverId, err := getUser(application.UserId)
		if err != nil {
			return
		}
		notfi := protocol.Notification{
			ServerId:  serverId,
			UserId:    application.UserId,
			Type:      protocol.FriendApplicationStatusChange,
			Content:   map[string]interface{}{"id": req.Id, "question": req.Question},
			Timestamp: time.Now().UnixMilli(),
		}
		body, err := json.Marshal(notfi)
		if err != nil {
			global.Logger.Errorf("好友申请提问通知-序列化失败，ID: %d, err: %v", req.Id, err)
			return
		}
		err = global.Channel.PublishMsg(context.Background(), global.NotificationExchangeName, global.NotificationRoutingKey, body)
		if err != nil {
			global.Logger.Errorf("好友申请提问通知-发送失败，ID: %d, err: %v", req.Id, err)
		}
	}()
	return ApplicationQuestResp{}, nil
}

func (f friendApplicationService) ApplicationAns(userId uint, req ApplicationAnsReq) (ApplicationAnsResp, error) {
	application, err := f.friendApplicationDao.GetApplicationById(req.Id)
	if err != nil {
		return ApplicationAnsResp{}, err
	}
	if application.ID == 0 {
		return ApplicationAnsResp{}, errcode.FriendApplyNotFoundErr
	}
	if userId == application.FriendId {
		return ApplicationAnsResp{}, errcode.CantHandleApplyErr
	}
	err = f.friendApplicationDao.UpdateApplication(req.Id, map[string]interface{}{"answer": req.Answer})
	if err != nil {
		return ApplicationAnsResp{}, err
	}
	go func() {
		serverId, err := getUser(application.FriendId)
		if err != nil {
			return
		}
		notfi := protocol.Notification{
			ServerId:  serverId,
			UserId:    application.FriendId,
			Type:      protocol.FriendApplicationStatusChange,
			Content:   map[string]interface{}{"id": req.Id, "answer": req.Answer},
			Timestamp: time.Now().UnixMilli(),
		}
		body, err := json.Marshal(notfi)
		if err != nil {
			global.Logger.Errorf("好友申请回答通知-序列化失败，ID: %d, err: %v", req.Id, err)
			return
		}
		err = global.Channel.PublishMsg(context.Background(), global.NotificationExchangeName, global.NotificationRoutingKey, body)
		if err != nil {
			global.Logger.Errorf("好友申请回答通知-发送失败，ID: %d, err: %v", req.Id, err)
		}
	}()
	return ApplicationAnsResp{}, nil
}

type friendService struct {
	friendDao *dao.FriendDao
	userDao   *dao.UserDao
}

func (f friendService) ListFriends(userId uint) ([]*UserInfo, error) {
	friends, err := f.friendDao.ListFriends(userId)
	if err != nil {
		return nil, err
	}
	userIds := make([]uint, 0, len(friends))
	for _, friend := range friends {
		userIds = append(userIds, friend.FriendId)
	}
	userInfo := make([]*UserInfo, 0, len(friends))
	users, err := f.userDao.GetUserBatch(userIds...)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		userInfo = append(userInfo, &UserInfo{
			UserId:   user.ID,
			UserName: user.UserName,
		})
	}
	return userInfo, nil
}

func (f friendService) DeleteFriend(req DeleteFriendReq) (DeleteFriendResp, error) {
	return DeleteFriendResp{}, f.friendDao.DeleteFriend(req.UserId, req.FriendId)
}

func NewFriendApplicationService(dao *dao.FriendApplicationDao, friendDao *dao.FriendDao) FriendApplicationService {
	return &friendApplicationService{
		friendDao:            friendDao,
		friendApplicationDao: dao,
	}
}

func NewFriendService(friendDao *dao.FriendDao, userDao *dao.UserDao) FriendService {
	return &friendService{
		friendDao: friendDao,
		userDao:   userDao,
	}
}
