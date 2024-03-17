package raftmgr

import (
	"fmt"
	"github.com/brunowang/gokv/internal/engine"
	"github.com/hashicorp/raft"
	"io"
)

type RaftFSM struct {
}

func NewRaftFSM() *RaftFSM {
	return &RaftFSM{}
}

func (a *RaftFSM) Apply(log *raft.Log) interface{} {
	kv := engine.NewKVPair()
	if err := kv.FromBytes(log.Data); err != nil {
		fmt.Printf("fsm apply error: %v\n", err)
		return nil
	}
	engine.Set(kv.Key, kv.Value)
	return nil
}
func (a *RaftFSM) Snapshot() (raft.FSMSnapshot, error) {
	return nil, nil
}
func (a *RaftFSM) Restore(io.ReadCloser) error {
	return nil
}
