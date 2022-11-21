import { Board, BoardClocks, Player } from './board';

export interface BoardState {
  board: Board;
  active_player: Player;
  clocks: BoardClocks;
}
