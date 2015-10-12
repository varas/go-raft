package raft

import (
    "testing"
    "fmt"
)

const NbNodes = 10

var nodes []Node

func TestMain(m *testing.T){
    cluster := NewCluster(NbNodes)
    nodes = cluster.Nodes()

    fmt.Println("nodes: ", nodes)
    for idx, node := range nodes {
        fmt.Println(idx, ": ", node)
    }
}

func TestAllNodesStartAsFollowers(t *testing.T){
    for _, node := range nodes {

        state, err := node.State()
        if (err != nil) {
            t.Errorf("error recovering state of node ", node)
        }
        if (state != Follower) {
            t.Errorf("state of node is not Follower: ", state)
        }
    }
}
