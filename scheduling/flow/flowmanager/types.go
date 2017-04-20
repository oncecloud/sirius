package flowmanager

import "github.com/oncecloud/sirius/scheduling/flow/flowgraph"

// TaskMapping is a 1:1 mapping from Task Node to Resource Node
type TaskMapping map[flowgraph.NodeID]flowgraph.NodeID
