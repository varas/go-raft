package raft

import (
    "testing"
)

func TestAllNodesStartAsFollowers(t *testing.T){
    cluster := NewCluster(10)
    nodes := cluster.Nodes
    println(nodes)
}
