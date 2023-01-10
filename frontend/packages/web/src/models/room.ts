export type Room = {
  id: string;
  name: string;
  players: { [playerID: string]: string };
  game_id: string;
};
