package minimax

import (
	"gomoku/minimax/heuristic"
	"container/heap"
	"gomoku/game"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Minimax(state game.State, width, depth int) Moves {
	if depth == 0 {
		return []Move{
			{
				State:      state,
				Evaluation: heuristic.Evaluation(state.Board, state.BlackScore, state.WhiteScore),
			},
		}
	}
	// get all possible moves (cells adjacent to occupied cells
	moveCoords := PossibleMoves(state.Board, 1)
	moves := make(Moves, len(moveCoords))
	// traverse all moves and order them by priority
	for i, coords := range moveCoords {
		stateAfterMove := state.MakeMoveImmut(coords)
		moves[i] = Move{
			Coords: coords,
			State:  stateAfterMove,
			Evaluation: heuristic.Evaluation(stateAfterMove.Board,
				stateAfterMove.BlackScore, stateAfterMove.WhiteScore),
		}
	}
	heap.Init(&moves)
	bestMovesLen := Min(width, moves.Len())
	bestMoves := make(Moves, bestMovesLen)
	// take best moves and proceed with them, discard the rest
	for i := 0; i < bestMovesLen; i++ {
		move := heap.Pop(&moves).(Move)
		move.Evaluation = Minimax(move.State, width, depth-1)[0].Evaluation
		bestMoves[i] = move
	}
	heap.Init(&bestMoves)
	return bestMoves.Slice(bestMoves.Len())
}
