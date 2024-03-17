package raftmgr

import (
	"fmt"
	"github.com/hashicorp/raft"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type OSPath string

func (p OSPath) Absolute() OSPath {
	path := string(p)
	if strings.HasPrefix(path, "/") {
		return p
	}
	dir, _ := os.Getwd()
	return OSPath(fmt.Sprintf("%s/%s", dir, path))
}

func (p OSPath) String() string {
	return string(p)
}

type Config struct {
	ServerName  string        `mapstructure:"server_name"`
	ServerID    string        `mapstructure:"server_id"`
	LogStore    OSPath        `mapstructure:"log_store"`
	StableStore OSPath        `mapstructure:"stable_store"`
	Transport   string        `mapstructure:"transport"`
	Servers     []raft.Server `mapstructure:"servers"`
}

func NewConfig() *Config {
	return &Config{}
}

func LoadConfig(path string) (*Config, error) {
	vip := viper.New()
	vip.SetConfigFile(path)
	if err := vip.ReadInConfig(); err != nil {
		fmt.Printf("failed to load config file: %s, error: %v\n", path, err)
		return nil, err
	}

	config := NewConfig()
	if err := vip.Unmarshal(config); err != nil {
		fmt.Printf("failed to parse config file: %s, error: %v\n", path, err)
		return nil, err
	}
	return config, nil
}
