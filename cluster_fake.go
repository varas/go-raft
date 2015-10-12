package raft

import (
    "strconv"
)

type cluster struct {
    nodes []Node
    channel chan Msg
}

func NewCluster(nbNodes int) *cluster {
    c := new(cluster)
    c.channel = make(chan Msg)
    c.nodes = make([]Node, nbNodes)
    for i := 0; i < nbNodes; i++ {
        c.nodes[i] = NewNode("node " + strconv.Itoa(i), c.channel)
    }
    return c
}

func (c *cluster) Nodes() []Node {
    return c.nodes
}
