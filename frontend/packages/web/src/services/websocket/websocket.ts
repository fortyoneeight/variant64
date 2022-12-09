import { json } from 'stream/consumers';
import { WebSocketRequestEvent, WebSocketServiceParams } from './types';

export class WebSocketService {
  url: string;
  private socket: WebSocket;

  private subscribers: Record<string, (event: any) => void>;
  constructor(params: WebSocketServiceParams) {
    this.url = params.url;
    this.socket = this.initializeConnection();
    this.subscribers = {};
  }

  private initializeConnection(): WebSocket {
    let socket = new WebSocket(this.url);

    socket.onopen = (e) => {
      console.log('[open] Connection established');
    };

    socket.onmessage = (event) => {
      console.info(`[message] Data received from server: ${event.data}`);
      Object.values(this.subscribers).forEach((cb) => cb(JSON.parse(event.data)));
    };

    socket.onclose = (event) => {
      if (event.wasClean) {
        console.log(`[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`);
      } else {
        // e.g. server process killed or network down
        // event.code is usually 1006 in this case
        console.log('[close] Connection died');
      }
    };

    socket.onerror = (error) => {
      console.log(`[error]`, error);
    };

    return socket;
  }

  send<T>(event: WebSocketRequestEvent) {
    const { action, channel, body } = event;

    const command = {
      command: action,
      channel,
      body: JSON.stringify(body),
    };

    const serialized = JSON.stringify(command);

    console.log(`[WEB_SOCKET_COMMAND] ${action}`, body);
    return this.socket.send(serialized);
  }

  subscribe(name: string, cb: (event: any) => void) {
    this.subscribers[name] = cb;
  }
}
