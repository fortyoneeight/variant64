import { Board, BoardClocks, BoardPlayer } from './board';

export interface BoardState {
  board: Board;
  active_player: BoardPlayer;
  clocks: BoardClocks;
}
