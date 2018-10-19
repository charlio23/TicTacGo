package main

type Ai struct {
	FinishGame bool
	ToeBoard   [][]string
}

func (e Ai) Ai(finish bool, board [][]string) {
	e.FinishGame = finish
	e.ToeBoard = board
}

func (e Ai) MixMax() {

}

func heuristica(ToeBoard [][]string) int {
	return 0
}
