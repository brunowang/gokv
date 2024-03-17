package main

import (
	"flag"
	"fmt"
	"github.com/brunowang/gokv/internal/raftmgr"
	"log"
	"time"
)

func main() {
	var cpath string
	flag.StringVar(&cpath, "c", "", "your config file ")
	flag.Parse()
	if cpath == "" {
		log.Fatal("config file error")
	}
	err := raftmgr.Init(cpath)
	if err != nil {
		log.Fatal(err)
	}
	for {
		leaderAddr, leaderID := raftmgr.RaftNode.LeaderWithID()
		fmt.Println("current leader:", leaderAddr, leaderID)
		time.Sleep(time.Second * 1)
	}
}
