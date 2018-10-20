package main

import (
	"strconv"
)

// human
var huPlayer string = "O"

// ai
var aiPlayer string = "X"

func win(board [9]string, player string) bool {

	if (board[0] == player && board[1] == player && board[2] == player) ||
		(board[3] == player && board[4] == player && board[5] == player) ||
		(board[6] == player && board[7] == player && board[8] == player) ||
		(board[0] == player && board[3] == player && board[6] == player) ||
		(board[1] == player && board[4] == player && board[7] == player) ||
		(board[2] == player && board[5] == player && board[8] == player) ||
		(board[0] == player && board[4] == player && board[8] == player) ||
		(board[2] == player && board[4] == player && board[6] == player) {
		return true
	} else {
		return false
	}
}

func filter(board [9]string, aux []string, value int) []string {
	if value < 0 {
		return aux
	} else {
		if board[value] == "X" || board[value] == "O" {
			aux = filter(board, aux, value-1)
		} else {
			i := append(aux, board[value])
			aux = filter(board, i, value-1)
		}
	}
	return aux
}

func emptyIndexes(board [9]string) []string {
	aux := make([]string, 0)
	return filter(board, aux, len(board)-1)
}

type Move struct {
	score int
	index string
}

func MinMax(board [9]string, player string) Move {
	var move Move
	var availSpots []string = emptyIndexes(board)

	if win(board, huPlayer) {
		move.score = -10
		return move
	} else if win(board, aiPlayer) {
		move.score = 10
		return move
	} else if len(availSpots) == 0 {
		move.score = 0
		return move
	}
	var moves []Move

	for i := 0; i < len(availSpots); i++ {
		var aux Move
		j, _ := strconv.Atoi(availSpots[i])
		aux.index = board[j]
		board[j] = player
		if player == aiPlayer {
			result := MinMax(board, huPlayer)
			aux.score = result.score
		} else {
			result := MinMax(board, aiPlayer)
			aux.score = result.score
		}

		board[j] = aux.index
		moves = append(moves, aux)
	}

	var bestMove int

	if player == aiPlayer {
		bestScore := -10000
		for i := 0; i < len(moves); i++ {
			if moves[i].score > bestScore {
				bestScore = moves[i].score
				bestMove = i
			}
		}
	} else {
		bestScore := 10000
		for i := 0; i < len(moves); i++ {
			if moves[i].score < bestScore {
				bestScore = moves[i].score
				bestMove = i
			}
		}
	}

	return moves[bestMove]
}

func callAi(board [9]string) string {

	return MinMax(board, aiPlayer).index
}
