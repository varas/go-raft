package raft

type Msg int

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
    receiverChannel chan Msg
}

func NewNode(id string, channel chan Msg) Node {
    node := new(Node)
    node.id = id
    node.state = Follower
    node.receiverChannel = channel
    return *node
}

func (n *Node) Id() (string) {
    return n.id
}

func (n *Node) State() (State) {
    return n.state
}
