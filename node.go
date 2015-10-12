package raft

type State int

const (
    Follower State = 1 << iota
    Candidate
    Leader
)

var states = [...]string {
    "follower",
    "candidate",
    "leader",
}

func (state State) String() string {
    return states[state -1]
}

type Node struct {
    id string
    state State
}

func (n *Node) State() (State) {
    return n.state
}
