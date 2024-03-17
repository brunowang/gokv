package raftmgr

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
	"net"
	"os"
	"path/filepath"
	"time"
)

var RaftNode *raft.Raft

func Init(path string) error {
	sysConfig, err := LoadConfig(path)
	if err != nil {
		return err
	}
	fmt.Println(sysConfig)
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(sysConfig.ServerID)
	config.Logger = hclog.New(&hclog.LoggerOptions{
		Name:   sysConfig.ServerName,
		Level:  hclog.LevelFromString("DEBUG"),
		Output: os.Stderr,
	})

	//logStore保存配置
	os.MkdirAll(filepath.Dir(sysConfig.LogStore.Absolute().String()), 0755)
	logStore, err := raftboltdb.NewBoltStore(sysConfig.LogStore.Absolute().String())
	if err != nil {
		return err
	}

	//保存节点信息
	os.MkdirAll(filepath.Dir(sysConfig.StableStore.Absolute().String()), 0755)
	stableStore, err := raftboltdb.NewBoltStore(sysConfig.StableStore.Absolute().String())
	if err != nil {
		return err
	}
	//不存储快照
	snapshotStore := raft.NewDiscardSnapshotStore()

	// 节点之间的通信
	addr, err := net.ResolveTCPAddr("tcp", sysConfig.Transport)
	transport, err := raft.NewTCPTransport(addr.String(), addr, 5, time.Second*10, os.Stdout)
	if err != nil {
		panic(err)
	}
	fsm := NewRaftFSM()

	RaftNode, err = raft.NewRaft(config, fsm, logStore, stableStore, snapshotStore, transport)
	if err != nil {
		return err
	}
	configuration := raft.Configuration{
		Servers: sysConfig.Servers,
	}

	RaftNode.BootstrapCluster(configuration)
	return nil
}
