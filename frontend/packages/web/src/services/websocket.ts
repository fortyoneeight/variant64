import { HttpServiceParams, RoutesConfig } from '../types';

export class WebSocketService {
  url: string;
  routesConfig: RoutesConfig;

  constructor(params: HttpServiceParams) {
    this.url = params.url;
    this.routesConfig = params.routesConfig;

    this.initializeConnection();
  }

  private initializeConnection() {
    let socket = new WebSocket(this.url);

    socket.onopen = function (e) {
      console.log('[open] Connection established');
    };

    socket.onmessage = function (event) {
      console.log(`[message] Data received from server: ${event.data}`);
    };

    socket.onclose = function (event) {
      if (event.wasClean) {
        console.log(`[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`);
      } else {
        // e.g. server process killed or network down
        // event.code is usually 1006 in this case
        console.log('[close] Connection died');
      }
    };

    socket.onerror = function (error) {
      console.log(`[error]`);
    };
  }
}
