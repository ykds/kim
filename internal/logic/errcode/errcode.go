package errcode

import (
	"kim/internal/pkg/errors"
)

var (
	MobileEmpty           = 100000
	NotRegister           = 100001
	PasswordNotMatch      = 100002
	InvalidToken          = 100003
	Registered            = 100004
	PasswordNotSame       = 100005
	DuplicatedFriendApply = 100006
	FriendApplyNotFound   = 100007
	ApplyStatusWrong      = 100008
	CantHandleApply       = 100009
	HadBeFriend           = 100010
)

var (
	MobileEmptyErr           = errors.NewError(MobileEmpty, "手机号不能为空")
	NotRegisterErr           = errors.NewError(NotRegister, "用户未注册")
	PasswordNotMatchErr      = errors.NewError(PasswordNotMatch, "密码错误")
	InvalidTokenErr          = errors.NewError(InvalidToken, "无效的token")
	RegisteredErr            = errors.NewError(Registered, "手机号已注册")
	PasswordNotSameErr       = errors.NewError(PasswordNotSame, "两次密码不一致")
	DuplicatedFriendApplyErr = errors.NewError(DuplicatedFriendApply, "重复好友申请")
	FriendApplyNotFoundErr   = errors.NewError(FriendApplyNotFound, "好友申请记录不存在")
	ApplyStatusWrongErr      = errors.NewError(ApplyStatusWrong, "好友申请状态错误")
	CantHandleApplyErr       = errors.NewError(CantHandleApply, "你不能执行该动作")
	HadBeFriendErr           = errors.NewError(HadBeFriend, "已经添加好友")
)
