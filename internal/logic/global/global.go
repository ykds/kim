package global

import (
	"kim/internal/logic/config"
	"kim/internal/pkg/log"
	"kim/internal/pkg/mysql"
	"kim/internal/pkg/redis"
)

const (
	Token = "token"
)

var (
	Conf *config.Config

	Database *mysql.Mysql
	Redis    *redis.Redis
	Logger   *log.Logger
)
