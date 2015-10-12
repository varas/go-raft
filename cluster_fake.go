package raft

import (
    "strconv"
)

type cluster struct {
    nodes []Node
}

func NewCluster(nbNodes int) *cluster {
    c := new(cluster)
    c.nodes = make([]Node, nbNodes)
    for i := 0; i < nbNodes; i++ {
        c.nodes[i] = Node{"node " + strconv.Itoa(i), Follower}
    }
    return c
}

func (c *cluster) Nodes() []Node {
    return c.nodes
}
