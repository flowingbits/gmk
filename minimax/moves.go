package minimax

import (
	"gomoku/board"
	"fmt"
	"container/heap"
	"gomoku/game"
)

type Move struct {
	board.Coords
	State      game.State `json:"state"`
	Evaluation int64      `json:"evaluation"`
}

type Moves []Move

func (pq Moves) Len() int {
	return len(pq)
}

func (pq Moves) Less(i, j int) bool {
	if pq.Len() == 0 {
		return true
	}
	if pq[0].State.Player == MIN_PLAYER {
		return pq[i].Evaluation < pq[j].Evaluation
	}
	if pq[0].State.Player == MAX_PLAYER {
		return pq[i].Evaluation > pq[j].Evaluation
	}
	panic(fmt.Errorf("pq.Player must be one of %d, %d, got %d", MIN_PLAYER,
		MAX_PLAYER, pq[0].State.Player))
}

func (pq Moves) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *Moves) Push(x interface{}) {
	item := x.(Move)
	*pq = append(*pq, item)
}

func (pq *Moves) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
func (pq *Moves) Slice(n int) []Move {
	movesLen := Min(n, pq.Len())
	moves := make([]Move, movesLen)
	for i := range moves {
		moves[i] = heap.Pop(pq).(Move)
	}
	return moves
}
