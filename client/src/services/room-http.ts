export interface HttpServiceParams {
  url: string;
}

export class RoomHttpService {
  url: string;

  constructor(params: HttpServiceParams) {
    this.url = params.url;
  }

  createRoom() {
    console.log('created room');
  }

  getRoom() {
    console.log('get room');
  }

  joinRoom() {
    console.log('join room');
  }

  startRoom() {
    console.log('start room');
  }
}
