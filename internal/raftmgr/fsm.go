package raftmgr

import (
	"github.com/hashicorp/raft"
	"io"
)

type RaftFSM struct {
}

func NewRaftFSM() *RaftFSM {
	return &RaftFSM{}
}

func (a *RaftFSM) Apply(log *raft.Log) interface{} {
	return nil
}
func (a *RaftFSM) Snapshot() (raft.FSMSnapshot, error) {
	return nil, nil
}
func (a *RaftFSM) Restore(io.ReadCloser) error {
	return nil
}
