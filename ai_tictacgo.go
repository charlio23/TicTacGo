package main

import "strconv"

var origBoard = {0,1,2,3,4,5,6,7,8}

// human
var huPlayer = "O";

// ai
var aiPlayer = "X";

func win(board []string, player string) bool {

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

func filter(board []string, aux [] string, value int) [] string{
	if board==nil{
		return aux
	}else{
		i := append(aux,board[value])
		filter(board,i,value-1)
	}
	return nil
}

func choose(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func emptyIndexies(board []string) []string{
	var aux []string
	return filter(board, aux, len(board))
}

func heuristica(board []string, availSpots []string)int{
	if (win(board, huPlayer)){
		return -10
	} else if (win(board, aiPlayer)){
		return 10
	} else if (len(availSpots) == 0){
		return 0
	}
}

type Move struct{
	score int
	index string
}

func movement (availSpots []string, player string, newBoard []string){
	var moves = {}
	var i int
	i=0
	// loop through available spots
	for (i < len(availSpots)){
		//create an object for each and store the index of that spot
		var move Move
		var j int = i

		move.index = newBoard[strconv.Atoi(availSpots.[i])]

		// set the empty spot to the current player
		newBoard[availSpots[i]] = player;

		/*collect the score resulted from calling minimax
		  on the opponent of the current player*/
		if (player == aiPlayer){
			var result = MinMax(newBoard, huPlayer);
			move.score = result.score;
		}
		else{
			var result = MinMax(newBoard, aiPlayer);
			move.score = result.score;
		}

		// reset the spot to empty
		newBoard[availSpots[i]] = move.index;

		// push the object to the array
		append(move,moves);
	i++;
	}
}

func MinMax(board []string, player string) Move{
	var availSpots = emptyIndexies(board)

	var heur = heuristica(board,availSpots)

	movement(availSpots,player,board)
	/*var result Node
	var value int
	var next []string
	var aux Node

	if finalNode(node) {
		if (level % 2) == 0 { //machine
			result.Node(e, math.MinInt64)
			return result
		} else { //player
			result.Node(e, math.MaxInt64)
			return result
		}
	} else if level == level_Max {
		value = heuristica(e, level)
		result.Node(e, value)
		return result
	} else {
		i:= len(e.ToeBoard)
		for alpha < beta && i>0 {
			next = emptyIndexies(e)
			aux = next.node.MinMax(alpha, beta, next, level+1)
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

		return result*/


		return 0
		}

