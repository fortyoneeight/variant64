export type Game = {
  id: string;
  active_player: null | string;
  winning_players: Array<string>;
  losing_players: Array<string>;
  drawn_players: Array<string>;
  approved_draw_players: { [playerId: string]: boolean };
  state: string;
};
