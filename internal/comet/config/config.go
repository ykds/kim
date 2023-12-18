package config

import (
	"kim/internal/pkg/etcd"
	"kim/internal/pkg/log"
)

type ServerConfig struct {
	ID                int32  `json:"id" yaml:"id"`
	Host              string `json:"host" yaml:"host"`
	Port              string `json:"port" yaml:"port"`
	LogicHost         string `json:"logic_host" yaml:"logic_host"`
	LogicPort         string `json:"logic_port" yaml:"logic_port"`
	BucketNum         int    `json:"bucket_num" yaml:"bucket_num"`
	HeartBeatInterval int    `json:"heart_beat_interval" yaml:"heart_beat_interval"`
}

type GrpcConfig struct {
	Host string `json:"host" yaml:"host"`
	Port string `json:"port" yaml:"port"`
}

type Config struct {
	Server     ServerConfig `json:"server" yaml:"server"`
	Log        log.Config   `json:"log" yaml:"log"`
	GrpcServer GrpcConfig   `json:"grpc_server" yaml:"grpc_server"`
	Etcd       etcd.Config  `json:"etcd" yaml:"etcd"`
}
