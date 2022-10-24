import axios from 'axios';
import {
  CreateRoomRequest,
  CreateRoomResponse,
  GetRoomsResponse,
  JoinRoomRequest,
  JoinRoomResponse,
  StartRoomRequest,
  StartRoomResponse,
} from '../types';
export interface HttpServiceParams {
  url: string;
}

export class RoomHttpService {
  url: string;

  constructor(params: HttpServiceParams) {
    this.url = params.url + '/api';
  }

  createRoom(body: CreateRoomRequest): PromiseLike<CreateRoomResponse> {
    return axios.post(this.url + '/room', body);
  }

  getRooms(): PromiseLike<GetRoomsResponse> {
    return axios.get(this.url + '/rooms');
  }

  joinRoom(roomName: string, body: JoinRoomRequest): PromiseLike<JoinRoomResponse> {
    return axios.post(this.url + `/room/${roomName}/join`, body);
  }

  startRoom(roomName: string, body: StartRoomRequest): StartRoomResponse {
    return axios.post(this.url + `/room/${roomName}/start`, body);
  }
}
