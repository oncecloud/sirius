// edge.go
package graph

type Edge struct {
	src           uint64
	dst           uint64
	cost          int64
	capUpperBound uint64
	capLowBound   uint64
}
