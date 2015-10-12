package raft

type State int

const (
    Follower State = 1 << iota
    Candidate
    Leader
)
func (state State) String() string {
    switch state {
    case Follower:
        return "follower"
    case Candidate:
        return "candidate"
    case Leader:
        return "leader"
    }
    return "unknown-state"
}

type Node struct {
    id string
    state State
}

func (n *Node) State() (State, error) {
    return n.state, nil
}
