package errcode

import (
	"kim/internal/pkg/errors"
)

var (
	MobileEmpty      = 100000
	NotRegister      = 100001
	PasswordNotMatch = 100002
	InvalidToken     = 100003
	Registered       = 100004
	PasswordNotSame  = 100005
)

var (
	MobileEmptyErr      = errors.NewError(MobileEmpty, "手机号不能为空")
	NotRegisterErr      = errors.NewError(NotRegister, "用户未注册")
	PasswordNotMatchErr = errors.NewError(PasswordNotMatch, "密码错误")
	InvalidTokenErr     = errors.NewError(InvalidToken, "无效的token")
	RegisteredErr       = errors.NewError(Registered, "手机号已注册")
	PasswordNotSameErr  = errors.NewError(PasswordNotSame, "两次密码不一致")
)
