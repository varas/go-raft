package raft

import (
    "testing"
    "fmt"
)

func TestAllNodesStartAsFollowers(t *testing.T){
    cluster := NewCluster(10)
    nodes := cluster.Nodes()

    fmt.Println("nodes: ", nodes)
    for k, v := range nodes {
        fmt.Println(k, ": ", v)
    }
}
