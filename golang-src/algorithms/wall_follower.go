package algorithms

import "server/maze"

// WallFollowerAlgorithm performs the wall follower algorithm on the grid
func WallFollowerAlgorithm(grid [][]maze.Node, startNode, endNode *maze.Node) []maze.Node {
	visitedNodesInOrder := []maze.Node{}
	startNode.Distance = 0
	currentNode := startNode
	var previousNode *maze.Node

	for currentNode != endNode {
		currentNode.IsVisited = true
		currentNode.NoOfVisits++
		visitedNodesInOrder = append(visitedNodesInOrder, *currentNode)

		neighbors := getUnvisitedNeighbors(currentNode, grid)
		var nextNode *maze.Node

		for _, neighbor := range neighbors {
			if neighbor != previousNode && !neighbor.IsWall {
				nextNode = neighbor
				break
			}
		}

		if nextNode != nil {
			nextNode.Distance = currentNode.Distance + 1
			nextNode.PreviousNode = currentNode

			nextNode.PreviousNodeIndex = currentNode.Index // Set PreviousNodeIndex

			previousNode = currentNode
			currentNode = nextNode
		} else {
			if previousNode != nil {
				currentNode = previousNode
				previousNode = currentNode.PreviousNode

				// Update PreviousNodeIndex during backtracking
				if previousNode != nil {
					currentNode.PreviousNodeIndex = previousNode.Index
				}
			} else {
				break
			}
		}
	}

	return visitedNodesInOrder
}
