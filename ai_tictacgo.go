package main

import "math"

const level_Max int = 9

type Ai struct {
	FinishGame bool
	ToeBoard   [][]string
}

type Node struct {
	node      Ai
	heuristic int
}

func (e Node) Node(nodeConstructor Ai, value int) {
	e.heuristic = value
	e.node = nodeConstructor
}

func (e Ai) Ai(finish bool, board [][]string) {
	e.FinishGame = finish
	e.ToeBoard = board
}

func heuristica(e Ai, level int) int {
	return 0
}

func finalNode(nodeToAvaluate Node) bool {
	return false
}

func nextNode(nodeToAvaluate Node) Node {
	return nodeToAvaluate
}

func (e Ai) MinMax(alpha int, beta int, node Node, level int) Node {
	var result Node
	var value int
	var next Node
	var aux Node

	if finalNode(node) {
		if (level % 2) == 0 { //machine
			result.Node(e, math.MinInt64)
			return result
		} else { //player
			result.Node(e, math.MaxUint64)
			return result
		}
	} else if level == level_Max {
		value = heuristica(e, level)
		result.Node(e, value)
		return result
	} else {
		for alpha < beta {
			next = nextNode(node)
			aux = e.MinMax(alpha, beta, next, level+1)
			if (level % 2) == 0 {
				if result.heuristic > alpha {
					alpha = result.heuristic
					result.node = aux.node
				}
			} else {
				if result.heuristic < beta {
					beta = result.heuristic
					result.node = aux.node
				}
			}
		}

		return result
	}
}
