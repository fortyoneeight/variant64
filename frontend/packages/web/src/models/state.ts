import { Board, BoardClocks } from './board';
import { Player } from './player';

export interface BoardState {
  board: Board;
  active_player: Player;
  clocks: BoardClocks;
}
