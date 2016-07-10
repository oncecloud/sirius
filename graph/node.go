package graph

const (
	Root = 1 + iota
	Sink
	Machine
)

type Node struct {
	id    uint64
	jobId int64
}
