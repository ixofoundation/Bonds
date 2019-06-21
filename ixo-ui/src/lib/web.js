import { inject } from 'aurelia-framework';
import { HttpClient } from 'aurelia-fetch-client';
import env from '../environment';
import { withBaseUrl, sendJSON, expectJSON, withBearerToken } from 'lib/custom-decorators';

@withBaseUrl(env.chainServiceUrl)
@sendJSON()
@expectJSON()
@withBearerToken()
export class ApiClient extends HttpClient {
  acquireToken( request ) {
    return '';
  }
}

@inject(ApiClient)
export class Web {
  constructor( apiClient ) {
    apiClient.configure( config => apiClient.modifyConfiguration( config ) );

    apiClient.configure(config => {
      config.withInterceptor( this.interceptor );
    });
    this.apiClient = apiClient;
  }

  get interceptor() {
    let me = this;
    return {
      async request(request) { return me.interceptRequest(request); },
      async response(response) { return me.interceptResponse(response); }
    };
  }


  interceptRequest(request) {
    return request;
  }

  interceptResponse(response) {
    return response;
  }

  async fetch(config) {
    let realConfig;
    if (typeof config === 'object') {
      realConfig = config;
    } else if (typeof config === 'function') {
      realConfig = new WebLibConfig();
      realConfig.url = '';
      realConfig.method = 'GET';
      realConfig.requestParams = [];
      realConfig.requestBody = undefined;
      let cnf = config(realConfig);
      if (WebLibConfig.prototype.isPrototypeOf(cnf)) {
        realConfig = cnf;
      }
    } else {
      throw new Error('invalid config');
    }

    let { url, method, requestParams, requestBody } = realConfig;
    url = this.addRequestParams(url, requestParams);

    const response = await this.apiClient.fetch(url, { method: method, body: JSON.stringify(requestBody) });

    if (response.status >= 200 && response.status <= 300) {
      let contentType = response.headers.get('content-type');
      if (contentType && contentType.includes('application/json')) {
        try {
          return await response.json();
        } catch (e) {
          throw response.body;
        }
      } else {
        console.log(`No JSON - content type is: ${contentType}`);
        let result = await response.text();
        return result;
      }
    } else {
      return Promise.reject({ status: response.status });
    }
  }


  addRequestParams(url, requestParams) {
    if (requestParams && requestParams.length > 0) {
      let urlWithParams = url + '?';
      for (let i = 0; i < requestParams.length; i++) {
        urlWithParams = urlWithParams + requestParams[i].key;
        urlWithParams = urlWithParams + '=';
        urlWithParams = urlWithParams + requestParams[i].value;
        if (i < requestParams.length - 1) {
          urlWithParams = urlWithParams + '&';
        }
      }
      return urlWithParams;
    }
    return url;
  }
}


class WebLibConfig {
  constructor() {
    this.url = '';
    this.method = 'GET';
    this.requestParams = [];
    this.requestBody = undefined;
  }
  withUrl(url) {
    this.url = url;
    return this;
  }
  withMethod(method) {
    this.method = method;
    return this;
  }
  withParams(requestParams) {
    this.requestParams = requestParams;
    return this;
  }
  withBody(requestBody) {
    this.requestBody = requestBody;
    return this;
  }
}
