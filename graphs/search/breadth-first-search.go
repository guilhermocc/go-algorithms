package search

import "go-algorithms/graphs"

// breadthSearch returns the size of the *minimum* path until the targetNodeID, or -1 if there is no path to target.
// time complexity: O(V+E), where V is the number of nodes and E is the number of edges.
func breadthSearch(initialNode *graphs.Node, targetNodeID int) int {
	if initialNode.Id == targetNodeID {
		return 0
	}
	initialNode.Visited = true
	initialNode.Distance = 0
	nodesQueue := []*graphs.Node{initialNode}
	for len(nodesQueue) != 0 {
		var currentNode *graphs.Node
		currentNode, nodesQueue = nodesQueue[0], nodesQueue[1:]
		for _, neighbour := range currentNode.Neighbours {
			if !neighbour.Visited {
				if neighbour.Id == targetNodeID {
					return currentNode.Distance + 1
				}
				neighbour.Visited = true
				neighbour.Distance = currentNode.Distance + 1
				nodesQueue = append(nodesQueue, neighbour)
			}
		}
	}
	return -1
}
