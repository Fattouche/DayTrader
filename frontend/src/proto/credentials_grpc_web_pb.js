/**
 * @fileoverview gRPC-Web generated client stub for daytrader
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.daytrader = require('./credentials_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.daytrader.LoggerClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.daytrader.LoggerPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.log,
 *   !proto.daytrader.Response>}
 */
const methodInfo_Logger_LogUserCommand = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.log} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.log} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.LoggerClient.prototype.logUserCommand =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.Logger/LogUserCommand',
      request,
      metadata || {},
      methodInfo_Logger_LogUserCommand,
      callback);
};


/**
 * @param {!proto.daytrader.log} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.LoggerPromiseClient.prototype.logUserCommand =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.Logger/LogUserCommand',
      request,
      metadata || {},
      methodInfo_Logger_LogUserCommand);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.log,
 *   !proto.daytrader.Response>}
 */
const methodInfo_Logger_LogQuoteServerEvent = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.log} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.log} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.LoggerClient.prototype.logQuoteServerEvent =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.Logger/LogQuoteServerEvent',
      request,
      metadata || {},
      methodInfo_Logger_LogQuoteServerEvent,
      callback);
};


/**
 * @param {!proto.daytrader.log} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.LoggerPromiseClient.prototype.logQuoteServerEvent =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.Logger/LogQuoteServerEvent',
      request,
      metadata || {},
      methodInfo_Logger_LogQuoteServerEvent);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.log,
 *   !proto.daytrader.Response>}
 */
const methodInfo_Logger_LogAccountTransaction = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.log} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.log} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.LoggerClient.prototype.logAccountTransaction =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.Logger/LogAccountTransaction',
      request,
      metadata || {},
      methodInfo_Logger_LogAccountTransaction,
      callback);
};


/**
 * @param {!proto.daytrader.log} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.LoggerPromiseClient.prototype.logAccountTransaction =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.Logger/LogAccountTransaction',
      request,
      metadata || {},
      methodInfo_Logger_LogAccountTransaction);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.log,
 *   !proto.daytrader.Response>}
 */
const methodInfo_Logger_LogSystemEvent = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.log} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.log} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.LoggerClient.prototype.logSystemEvent =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.Logger/LogSystemEvent',
      request,
      metadata || {},
      methodInfo_Logger_LogSystemEvent,
      callback);
};


/**
 * @param {!proto.daytrader.log} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.LoggerPromiseClient.prototype.logSystemEvent =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.Logger/LogSystemEvent',
      request,
      metadata || {},
      methodInfo_Logger_LogSystemEvent);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.log,
 *   !proto.daytrader.Response>}
 */
const methodInfo_Logger_LogErrorEvent = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.log} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.log} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.LoggerClient.prototype.logErrorEvent =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.Logger/LogErrorEvent',
      request,
      metadata || {},
      methodInfo_Logger_LogErrorEvent,
      callback);
};


/**
 * @param {!proto.daytrader.log} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.LoggerPromiseClient.prototype.logErrorEvent =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.Logger/LogErrorEvent',
      request,
      metadata || {},
      methodInfo_Logger_LogErrorEvent);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.log,
 *   !proto.daytrader.Response>}
 */
const methodInfo_Logger_LogDebugEvent = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.log} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.log} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.LoggerClient.prototype.logDebugEvent =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.Logger/LogDebugEvent',
      request,
      metadata || {},
      methodInfo_Logger_LogDebugEvent,
      callback);
};


/**
 * @param {!proto.daytrader.log} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.LoggerPromiseClient.prototype.logDebugEvent =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.Logger/LogDebugEvent',
      request,
      metadata || {},
      methodInfo_Logger_LogDebugEvent);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.command,
 *   !proto.daytrader.Response>}
 */
const methodInfo_Logger_DumpLogs = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.command} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.LoggerClient.prototype.dumpLogs =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.Logger/DumpLogs',
      request,
      metadata || {},
      methodInfo_Logger_DumpLogs,
      callback);
};


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.LoggerPromiseClient.prototype.dumpLogs =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.Logger/DumpLogs',
      request,
      metadata || {},
      methodInfo_Logger_DumpLogs);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.daytrader.DayTraderClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.daytrader.DayTraderPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.command,
 *   !proto.daytrader.Response>}
 */
const methodInfo_DayTrader_CreateUser = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.command} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.DayTraderClient.prototype.createUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.DayTrader/CreateUser',
      request,
      metadata || {},
      methodInfo_DayTrader_CreateUser,
      callback);
};


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.DayTraderPromiseClient.prototype.createUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.DayTrader/CreateUser',
      request,
      metadata || {},
      methodInfo_DayTrader_CreateUser);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.command,
 *   !proto.daytrader.Response>}
 */
const methodInfo_DayTrader_Add = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.command} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.DayTraderClient.prototype.add =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.DayTrader/Add',
      request,
      metadata || {},
      methodInfo_DayTrader_Add,
      callback);
};


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.DayTraderPromiseClient.prototype.add =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.DayTrader/Add',
      request,
      metadata || {},
      methodInfo_DayTrader_Add);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.command,
 *   !proto.daytrader.Response>}
 */
const methodInfo_DayTrader_Quote = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.command} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.DayTraderClient.prototype.quote =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.DayTrader/Quote',
      request,
      metadata || {},
      methodInfo_DayTrader_Quote,
      callback);
};


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.DayTraderPromiseClient.prototype.quote =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.DayTrader/Quote',
      request,
      metadata || {},
      methodInfo_DayTrader_Quote);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.command,
 *   !proto.daytrader.Response>}
 */
const methodInfo_DayTrader_Buy = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.command} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.DayTraderClient.prototype.buy =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.DayTrader/Buy',
      request,
      metadata || {},
      methodInfo_DayTrader_Buy,
      callback);
};


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.DayTraderPromiseClient.prototype.buy =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.DayTrader/Buy',
      request,
      metadata || {},
      methodInfo_DayTrader_Buy);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.command,
 *   !proto.daytrader.Response>}
 */
const methodInfo_DayTrader_Sell = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.command} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.DayTraderClient.prototype.sell =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.DayTrader/Sell',
      request,
      metadata || {},
      methodInfo_DayTrader_Sell,
      callback);
};


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.DayTraderPromiseClient.prototype.sell =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.DayTrader/Sell',
      request,
      metadata || {},
      methodInfo_DayTrader_Sell);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.command,
 *   !proto.daytrader.Response>}
 */
const methodInfo_DayTrader_CommitBuy = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.command} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.DayTraderClient.prototype.commitBuy =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.DayTrader/CommitBuy',
      request,
      metadata || {},
      methodInfo_DayTrader_CommitBuy,
      callback);
};


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.DayTraderPromiseClient.prototype.commitBuy =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.DayTrader/CommitBuy',
      request,
      metadata || {},
      methodInfo_DayTrader_CommitBuy);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.command,
 *   !proto.daytrader.Response>}
 */
const methodInfo_DayTrader_CommitSell = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.command} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.DayTraderClient.prototype.commitSell =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.DayTrader/CommitSell',
      request,
      metadata || {},
      methodInfo_DayTrader_CommitSell,
      callback);
};


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.DayTraderPromiseClient.prototype.commitSell =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.DayTrader/CommitSell',
      request,
      metadata || {},
      methodInfo_DayTrader_CommitSell);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.command,
 *   !proto.daytrader.Response>}
 */
const methodInfo_DayTrader_CancelBuy = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.command} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.DayTraderClient.prototype.cancelBuy =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.DayTrader/CancelBuy',
      request,
      metadata || {},
      methodInfo_DayTrader_CancelBuy,
      callback);
};


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.DayTraderPromiseClient.prototype.cancelBuy =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.DayTrader/CancelBuy',
      request,
      metadata || {},
      methodInfo_DayTrader_CancelBuy);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.command,
 *   !proto.daytrader.Response>}
 */
const methodInfo_DayTrader_CancelSell = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.command} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.DayTraderClient.prototype.cancelSell =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.DayTrader/CancelSell',
      request,
      metadata || {},
      methodInfo_DayTrader_CancelSell,
      callback);
};


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.DayTraderPromiseClient.prototype.cancelSell =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.DayTrader/CancelSell',
      request,
      metadata || {},
      methodInfo_DayTrader_CancelSell);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.command,
 *   !proto.daytrader.Response>}
 */
const methodInfo_DayTrader_SetBuyAmount = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.command} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.DayTraderClient.prototype.setBuyAmount =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.DayTrader/SetBuyAmount',
      request,
      metadata || {},
      methodInfo_DayTrader_SetBuyAmount,
      callback);
};


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.DayTraderPromiseClient.prototype.setBuyAmount =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.DayTrader/SetBuyAmount',
      request,
      metadata || {},
      methodInfo_DayTrader_SetBuyAmount);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.command,
 *   !proto.daytrader.Response>}
 */
const methodInfo_DayTrader_SetSellAmount = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.command} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.DayTraderClient.prototype.setSellAmount =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.DayTrader/SetSellAmount',
      request,
      metadata || {},
      methodInfo_DayTrader_SetSellAmount,
      callback);
};


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.DayTraderPromiseClient.prototype.setSellAmount =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.DayTrader/SetSellAmount',
      request,
      metadata || {},
      methodInfo_DayTrader_SetSellAmount);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.command,
 *   !proto.daytrader.Response>}
 */
const methodInfo_DayTrader_SetBuyTrigger = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.command} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.DayTraderClient.prototype.setBuyTrigger =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.DayTrader/SetBuyTrigger',
      request,
      metadata || {},
      methodInfo_DayTrader_SetBuyTrigger,
      callback);
};


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.DayTraderPromiseClient.prototype.setBuyTrigger =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.DayTrader/SetBuyTrigger',
      request,
      metadata || {},
      methodInfo_DayTrader_SetBuyTrigger);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.command,
 *   !proto.daytrader.Response>}
 */
const methodInfo_DayTrader_SetSellTrigger = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.command} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.DayTraderClient.prototype.setSellTrigger =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.DayTrader/SetSellTrigger',
      request,
      metadata || {},
      methodInfo_DayTrader_SetSellTrigger,
      callback);
};


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.DayTraderPromiseClient.prototype.setSellTrigger =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.DayTrader/SetSellTrigger',
      request,
      metadata || {},
      methodInfo_DayTrader_SetSellTrigger);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.command,
 *   !proto.daytrader.Response>}
 */
const methodInfo_DayTrader_CancelSetSell = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.command} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.DayTraderClient.prototype.cancelSetSell =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.DayTrader/CancelSetSell',
      request,
      metadata || {},
      methodInfo_DayTrader_CancelSetSell,
      callback);
};


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.DayTraderPromiseClient.prototype.cancelSetSell =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.DayTrader/CancelSetSell',
      request,
      metadata || {},
      methodInfo_DayTrader_CancelSetSell);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.command,
 *   !proto.daytrader.Response>}
 */
const methodInfo_DayTrader_CancelSetBuy = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.command} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.DayTraderClient.prototype.cancelSetBuy =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.DayTrader/CancelSetBuy',
      request,
      metadata || {},
      methodInfo_DayTrader_CancelSetBuy,
      callback);
};


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.DayTraderPromiseClient.prototype.cancelSetBuy =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.DayTrader/CancelSetBuy',
      request,
      metadata || {},
      methodInfo_DayTrader_CancelSetBuy);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.command,
 *   !proto.daytrader.Response>}
 */
const methodInfo_DayTrader_DumpLog = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.command} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.DayTraderClient.prototype.dumpLog =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.DayTrader/DumpLog',
      request,
      metadata || {},
      methodInfo_DayTrader_DumpLog,
      callback);
};


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.DayTraderPromiseClient.prototype.dumpLog =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.DayTrader/DumpLog',
      request,
      metadata || {},
      methodInfo_DayTrader_DumpLog);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.daytrader.command,
 *   !proto.daytrader.Response>}
 */
const methodInfo_DayTrader_DisplaySummary = new grpc.web.AbstractClientBase.MethodInfo(
  proto.daytrader.Response,
  /** @param {!proto.daytrader.command} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.daytrader.Response.deserializeBinary
);


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.daytrader.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.daytrader.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.daytrader.DayTraderClient.prototype.displaySummary =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/daytrader.DayTrader/DisplaySummary',
      request,
      metadata || {},
      methodInfo_DayTrader_DisplaySummary,
      callback);
};


/**
 * @param {!proto.daytrader.command} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.daytrader.Response>}
 *     A native promise that resolves to the response
 */
proto.daytrader.DayTraderPromiseClient.prototype.displaySummary =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/daytrader.DayTrader/DisplaySummary',
      request,
      metadata || {},
      methodInfo_DayTrader_DisplaySummary);
};


module.exports = proto.daytrader;

