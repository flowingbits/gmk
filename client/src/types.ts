export enum GameType {
  vsFriend = "vsFriend",
  vsComputer = "vsComputer",
  debug = "debug"
}

export enum IPlayer {
  Black = 1,
  White = 2
}

export interface ICommonGameState {
  player: IPlayer;
  board: IBoard;
  blackScore: number;
  whiteScore: number;
  winner: IPlayer | 0
}

export type IBoardCell = number;

export type IBoardRow = [
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell
  ];

export type IBoard = [
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow
  ];

export interface ICoords {
  x: number;
  y: number;
}

export interface ISuggestion extends ICoords {
  state: ICommonGameState
  evaluation: number
}

export type ICellClickHandler = (coords: ICoords) => void;
