package raft

import (
    "strconv"
)

type cluster struct {
    nodes [10]Node
}

func NewCluster(nbNodes int) *cluster {
    c := new(cluster)
    for i := 0; i < nbNodes; i++ {
        c.nodes[i] = Node{"node " + strconv.Itoa(i), StateFollower}
    }
    return c
}

func (c *cluster) Nodes() [10]Node {
    return c.nodes
}
