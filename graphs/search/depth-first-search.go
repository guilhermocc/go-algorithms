package search

import "go-algorithms/graphs"

// depthSearchRec returns the size of a (not necessarily the minimum) path until the targetNodeID, or -1 if there is no path to target
// time complexity: O(V+E), where V is the number of nodes and E is the number of edges.
func depthSearchRec(currentNode *graphs.Node, targetNodeID int) int {
	// mark as visited to avoid cycles
	currentNode.Visited = true
	if currentNode.Id == targetNodeID {
		return 0
	}
	for _, neighbour := range currentNode.Neighbours {
		if !neighbour.Visited {
			pathFromNeighbour := 1 + depthSearchRec(neighbour, targetNodeID)
			if pathFromNeighbour > 0 {
				return pathFromNeighbour
			}
		}
	}
	return -1
}
