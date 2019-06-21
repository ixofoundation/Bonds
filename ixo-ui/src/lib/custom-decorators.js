export function withBaseUrl(baseUrl) {
  return target => {
    const configureBaseUrl = config => {
      return config.withBaseUrl(baseUrl);
    };

    const previousModifyConfiguation = target.prototype.modifyConfiguration || function(config) { return config; };

    target.prototype.modifyConfiguration = config => {
      return configureBaseUrl(previousModifyConfiguation.call(this, config));
    };
  };
}


export function sendJSON() {
  return target => {
    const addContentTypeJSON = config => {
      const defaults = config.defaults || {};
      return config.withDefaults( {...defaults, ...{ headers: { 'content-type': 'application/json' } } } ); //merge
    };

    const previousModifyConfiguation = target.prototype.modifyConfiguration || function(config) { return config; };

    target.prototype.modifyConfiguration = config => {
      return addContentTypeJSON(previousModifyConfiguation.call(this, config));
    };
  };
}


export function expectJSON() {
  return target => {
    const addContentTypeJSON = config => {
      const defaults = config.defaults || {};
      return config.withDefaults( {...defaults, ...{ headers: { 'accept': 'application/json' } } } ); //merge
    };

    const previousModifyConfiguation = target.prototype.modifyConfiguration || function(config) { return config; };

    target.prototype.modifyConfiguration = config => {
      return addContentTypeJSON(previousModifyConfiguation.call(this, config));
    };
  };
}


export function withBearerToken() {
  return target => {
    if (!target.prototype.acquireToken) {
      throw new Error(target.name + ' decorated with @withBearerToken() needs to implement an \'acquireToken( request )\' function.');
    }

    const addBearerTokenInterceptor = (context, config) => {
      return config.withInterceptor(bearerTokenInterceptorFactory(context));
    };

    const previousModifyConfiguation = target.prototype.modifyConfiguration || function(config) { return config; };

    target.prototype.modifyConfiguration = config => {
      return addBearerTokenInterceptor(this, previousModifyConfiguation.call(this, config));
    };
  };
}

function bearerTokenInterceptorFactory(context) {
  return {
    request: request => {
      return context.acquireToken( request )
        .then(token => {
          if (token) {
            request.headers.append('Authorization', 'Bearer ' + token);
          }
          return request;
        })
        .catch(err => {
          if (context.acquireTokenError) {
            context.acquireTokenError(err);
          }
          throw err;
        });
    }
  };
}
