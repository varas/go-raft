package raft

import (
    "strconv"
)

type Node struct {
    id string
}

type Cluster struct {
    nodes [10]Node
}

func NewCluster(nbNodes int) *Cluster {
    c := new(Cluster)
    for i := 0; i < nbNodes; i++ {
        c.nodes[i] = Node{"node " + strconv.Itoa(i)}
    }
    return c
}

func (c *Cluster) Nodes() [10]Node {
    return c.nodes
}
