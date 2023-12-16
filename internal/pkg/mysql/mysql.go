package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Config struct {
	Host            string `json:"host" yaml:"host"`
	Port            string `json:"port" yaml:"port"`
	User            string `json:"user" yaml:"user"`
	Password        string `json:"password" yaml:"password"`
	Database        string `json:"database" yaml:"database"`
	MaxIdleConns    int    `json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns    int    `json:"max_open_conns" yaml:"max_open_con"`
	ConnMaxIdleTime int    `json:"conn_max_idle_time" yaml:"conn_max_idle_time"`
}

type Mysql struct {
	*gorm.DB
}

func InitMysql(cfg Config) *Mysql {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDb.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDb.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime) * time.Second)

	return &Mysql{db}
}
