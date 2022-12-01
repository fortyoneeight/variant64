import {
  HttpService,
  WebSocketService,
  GetRoomsResponse,
  CreateRoomResponse,
  GetRoomResponse,
  JoinRoomResponse,
  LeaveRoomResponse,
  CreatePlayerResponse,
  GetPlayerResponse,
  SubscribeGameUpdatesCommand,
  AppActions,
  StartGameResponse,
  ConcedeResponse,
  RoutesParams,
} from '../services';

export class HomepageService {
  private httpservice: HttpService;
  private websocketservice: WebSocketService;
  constructor(httpservice: HttpService, websocketservice: WebSocketService) {
    this.httpservice = httpservice;
    this.websocketservice = websocketservice;
  }

  getRooms() {
    return this.httpservice.request<GetRoomsResponse>({
      action: AppActions.GET_ROOMS,
    });
  }

  createRoom(roomName: string) {
    return this.httpservice.request<CreateRoomResponse>({
      action: AppActions.CREATE_ROOM,
      body: {
        [RoutesParams.ROOM_NAME]: roomName,
      },
    });
  }

  getRoom(roomID: string) {
    return this.httpservice.request<GetRoomResponse>({
      action: AppActions.GET_ROOM,
      params: { [RoutesParams.ROOM_ID]: roomID },
    });
  }

  joinRoom(roomID: string, playerID: string) {
    return this.httpservice.request<JoinRoomResponse>({
      action: AppActions.JOIN_ROOM,
      params: {
        [RoutesParams.ROOM_ID]: roomID,
      },
      body: {
        [RoutesParams.PLAYER_ID]: playerID,
      },
    });
  }

  leaveRoom(roomID: string, playerID: string) {
    return this.httpservice.request<LeaveRoomResponse>({
      action: AppActions.LEAVE_ROOM,
      params: {
        [RoutesParams.ROOM_ID]: roomID,
      },
      body: {
        [RoutesParams.PLAYER_ID]: playerID,
      },
    });
  }

  startRoom(roomID: string, clockMillis: number) {
    return this.httpservice.request<StartGameResponse>({
      action: AppActions.START_ROOM,
      body: {
        [RoutesParams.ROOM_ID]: roomID,
        [RoutesParams.PLAYER_TIME_MILLIS]: clockMillis,
      },
    });
  }

  createPlayer(displayName: string) {
    return this.httpservice.request<CreatePlayerResponse>({
      action: AppActions.CREATE_PLAYER,
      body: {
        [RoutesParams.PLAYER_DISPLAY_NAME]: displayName,
      },
    });
  }

  getPlayer(playerID: string) {
    return this.httpservice.request<GetPlayerResponse>({
      action: AppActions.GET_PLAYER,
      params: {
        [RoutesParams.PLAYER_ID]: playerID,
      },
    });
  }

  subscribeToGameUpdates(gameID: string) {
    this.websocketservice.send<SubscribeGameUpdatesCommand>({
      action: AppActions.SUBSCRIBE_GAME_UPDATES,
      body: {
        [RoutesParams.GAME_ID]: gameID,
      },
    });
  }

  concedeGame(gameID: string, playerID: string) {
    return this.httpservice.request<ConcedeResponse>({
      action: AppActions.CONCEDE_GAME,
      params: {
        [RoutesParams.GAME_ID]: gameID,
      },
      body: {
        [RoutesParams.PLAYER_ID]: playerID,
      },
    });
  }
}
