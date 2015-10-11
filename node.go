package raft

const StateFollower = "follower"

type Node struct {
    id string
    state string
}

func (n *Node) State() (string, error) {
    return n.state, nil
}
