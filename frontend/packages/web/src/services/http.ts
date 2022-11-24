import axios from 'axios';
import { AppActions, HttpServiceParams, RoutesConfig } from '../types';

interface requestEvent {
  action: AppActions;
  params?: any;
  body?: any;
}
export class HttpService {
  url: string;
  routesConfig: RoutesConfig;

  constructor(params: HttpServiceParams) {
    this.url = params.url;
    this.routesConfig = params.routesConfig;
  }

  fetchRequest<T>(eventName: string, method: string, path: string, body?: any): Promise<T> {
    return axios
      .request({
        method: method,
        url: this.url + path,
        data: body || {},
      })
      .then((res) => {
        console.log(`[${this.routesConfig.name}_${eventName}] response: `, res);
        return res.data;
      })
      .catch((err) => {
        err['isError'] = true;
        console.error(`[${this.routesConfig.name}_${eventName}] error: `, err);
        throw err;
      });
  }

  request<T>(event: requestEvent): Promise<T> {
    const { action, params } = event;
    const { path, method } = this.routesConfig?.routes[action];

    return this.fetchRequest<T>(action, method, path(params), event.body);
  }
}
