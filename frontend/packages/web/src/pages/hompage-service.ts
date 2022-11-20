import { HttpService } from '../services';
import {
  AppActions,
  CreatePlayerResponse,
  CreateRoomResponse,
  GetPlayerResponse,
  GetRoomResponse,
  GetRoomsResponse,
  JoinRoomResponse,
  LeaveRoomResponse,
  RoutesParams,
} from '../types';

export class HomepageService {
  private httpservice: HttpService;
  constructor(httpservice: HttpService) {
    this.httpservice = httpservice;
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
}
