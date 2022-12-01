import { WebSocketRequestEvent, WebSocketServiceParams } from './types';

export class WebSocketService {
  url: string;
  private socket: WebSocket;

  constructor(params: WebSocketServiceParams) {
    this.url = params.url;

    this.socket = this.initializeConnection();
  }

  private initializeConnection(): WebSocket {
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

    return socket;
  }

  send<T>(event: WebSocketRequestEvent) {
    const { action, body } = event;

    const command = {
      action,
      ...(body as T),
    };

    const serialized = JSON.stringify(command);

    console.log(`[WEB_SOCKET_COMMAND] ${action}`, body);
    return this.socket.send(serialized);
  }
}
