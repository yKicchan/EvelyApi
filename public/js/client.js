// This module exports functions that give access to the EvelyApi API hosted at localhost:8888.
// It uses the axios javascript library for making the actual HTTP requests.
define(['axios'] , function (axios) {
  function merge(obj1, obj2) {
    var obj3 = {};
    for (var attrname in obj1) { obj3[attrname] = obj1[attrname]; }
    for (var attrname in obj2) { obj3[attrname] = obj2[attrname]; }
    return obj3;
  }

  return function (scheme, host, timeout) {
    scheme = scheme || 'http';
    host = host || 'localhost:8888';
    timeout = timeout || 20000;

    // Client is the object returned by this module.
    var client = axios;

    // URL prefix for all API requests.
    var urlPrefix = scheme + '://' + host;

  // イベント作成
  // path is the request path, the format is "/api/develop/v2/events"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.createEvents = function (path, data, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // レビュー投稿
  // path is the request path, the format is "/api/develop/v2/reviews/:event_id"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.createReviews = function (path, data, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // イベント削除
  // path is the request path, the format is "/api/develop/v2/events/:event_id"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.deleteEvents = function (path, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'delete',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // イベント複数取得
  // path is the request path, the format is "/api/develop/v2/events"
  // category, keyword, limit, offset are used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.listEvents = function (path, category, keyword, limit, offset, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      params: {
        category: category,
        keyword: keyword,
        limit: limit,
        offset: offset
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // レビューの一覧取得
  // path is the request path, the format is "/api/develop/v2/reviews/:event_id"
  // limit, offset are used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.listReviews = function (path, limit, offset, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      params: {
        limit: limit,
        offset: offset
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // イベント編集
  // path is the request path, the format is "/api/develop/v2/events/:event_id"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.modifyEvents = function (path, data, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'put',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 自分のイベント一覧を取得する
  // path is the request path, the format is "/api/develop/v2/events/my_list"
  // limit, offset are used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.my_listEvents = function (path, limit, offset, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      params: {
        limit: limit,
        offset: offset
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 近くのイベントを検索する
  // path is the request path, the format is "/api/develop/v2/events/nearby"
  // category, lat, limit, lng, offset, range are used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.nearbyEvents = function (path, category, lat, limit, lng, offset, range, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      params: {
        category: category,
        lat: lat,
        limit: limit,
        lng: lng,
        offset: offset,
        range: range
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 近くにイベントがあれば通知する
  // path is the request path, the format is "/api/develop/v2/events/notify"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.notifyEvents = function (path, data, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // ピンを外す
  // path is the request path, the format is "/api/develop/v2/pins/off"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.offPins = function (path, data, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'put',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // ピンする
  // path is the request path, the format is "/api/develop/v2/pins/on"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.onPins = function (path, data, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'put',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // ユーザーのピンしたイベント一覧を取得する
  // path is the request path, the format is "/api/develop/v2/events/pin/:user_id"
  // limit, offset are used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.pinEvents = function (path, limit, offset, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      params: {
        limit: limit,
        offset: offset
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 新規登録用のメール送信
  // path is the request path, the format is "/api/develop/v2/auth/signup/send_mail"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.send_mailAuth = function (path, data, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // イベント情報取得
  // path is the request path, the format is "/api/develop/v2/events/detail"
  // ids is used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.showEvents = function (path, ids, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      params: {
        ids: ids
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // アカウント情報取得
  // path is the request path, the format is "/api/develop/v2/users/:user_id"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.showUsers = function (path, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // ログイン
  // path is the request path, the format is "/api/develop/v2/auth/signin"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.signinAuth = function (path, data, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 新規登録
  // path is the request path, the format is "/api/develop/v2/auth/signup"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.signupAuth = function (path, data, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // イベントの開催フラグを更新する
  // path is the request path, the format is "/api/develop/v2/events/update"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.updateEvents = function (path, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // インスタンスIDの登録・更新
認証ありで登録ユーザーを、認証なしでゲストユーザーを登録・更新する
  // path is the request path, the format is "/api/develop/v2/users/update/token"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.updateUsers = function (path, data, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // ファイルアップロード
  // path is the request path, the format is "/api/develop/v2/files/upload"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.uploadFiles = function (path, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 新規登録時のトークンのチェック
  // path is the request path, the format is "/api/develop/v2/auth/signup/verify_token"
  // token is used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.verify_tokenAuth = function (path, token, config) {
    var cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      params: {
        token: token
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }
  return client;
  };
});
