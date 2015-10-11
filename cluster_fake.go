package raft

import (
    "strconv"
)

type Cluster struct {
    nodes [10]Node
}

func NewCluster(nbNodes int) *Cluster {
    c := new(Cluster)
    for i := 0; i < nbNodes; i++ {
        c.nodes[i] = Node{"node " + strconv.Itoa(i), StateFollower}
    }
    return c
}

func (c *Cluster) Nodes() [10]Node {
    return c.nodes
}
