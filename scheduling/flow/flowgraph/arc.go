

package flowgraph

//Enum for flow arc type
type ArcType int

const (
	ArcTypeOther ArcType = iota
	ArcTypeRunning
)

// Represents an arc in the scheduling flow graph.
type Arc struct {
	Src     NodeID
	Dst     NodeID
	SrcNode *Node
	DstNode *Node

	CapLowerBound uint64
	CapUpperBound uint64
	Cost          int64
	Type          ArcType
}

// Constructor equivalent in go
func NewArc(srcNode, dstNode *Node) *Arc {
	a := &Arc{
		Src:     srcNode.ID,
		Dst:     dstNode.ID,
		SrcNode: srcNode,
		DstNode: dstNode,
	}
	return a
}
