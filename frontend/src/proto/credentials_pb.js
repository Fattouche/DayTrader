/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.daytrader.Response', null, global);
goog.exportSymbol('proto.daytrader.command', null, global);
goog.exportSymbol('proto.daytrader.log', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.daytrader.command = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.daytrader.command, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.daytrader.command.displayName = 'proto.daytrader.command';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.daytrader.log = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.daytrader.log, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.daytrader.log.displayName = 'proto.daytrader.log';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.daytrader.Response = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.daytrader.Response, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.daytrader.Response.displayName = 'proto.daytrader.Response';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.daytrader.command.prototype.toObject = function(opt_includeInstance) {
  return proto.daytrader.command.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.daytrader.command} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.daytrader.command.toObject = function(includeInstance, msg) {
  var obj = {
    userId: jspb.Message.getFieldWithDefault(msg, 1, ""),
    amount: +jspb.Message.getFieldWithDefault(msg, 2, 0.0),
    symbol: jspb.Message.getFieldWithDefault(msg, 3, ""),
    filename: jspb.Message.getFieldWithDefault(msg, 4, ""),
    transactionId: jspb.Message.getFieldWithDefault(msg, 5, 0),
    name: jspb.Message.getFieldWithDefault(msg, 6, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.daytrader.command}
 */
proto.daytrader.command.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.daytrader.command;
  return proto.daytrader.command.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.daytrader.command} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.daytrader.command}
 */
proto.daytrader.command.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setUserId(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setAmount(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setSymbol(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setFilename(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setTransactionId(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.daytrader.command.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.daytrader.command.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.daytrader.command} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.daytrader.command.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAmount();
  if (f !== 0.0) {
    writer.writeFloat(
      2,
      f
    );
  }
  f = message.getSymbol();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getFilename();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getTransactionId();
  if (f !== 0) {
    writer.writeInt32(
      5,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
};


/**
 * optional string user_id = 1;
 * @return {string}
 */
proto.daytrader.command.prototype.getUserId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.daytrader.command.prototype.setUserId = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional float amount = 2;
 * @return {number}
 */
proto.daytrader.command.prototype.getAmount = function() {
  return /** @type {number} */ (+jspb.Message.getFieldWithDefault(this, 2, 0.0));
};


/** @param {number} value */
proto.daytrader.command.prototype.setAmount = function(value) {
  jspb.Message.setProto3FloatField(this, 2, value);
};


/**
 * optional string symbol = 3;
 * @return {string}
 */
proto.daytrader.command.prototype.getSymbol = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.daytrader.command.prototype.setSymbol = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string filename = 4;
 * @return {string}
 */
proto.daytrader.command.prototype.getFilename = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.daytrader.command.prototype.setFilename = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional int32 transaction_id = 5;
 * @return {number}
 */
proto.daytrader.command.prototype.getTransactionId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {number} value */
proto.daytrader.command.prototype.setTransactionId = function(value) {
  jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional string name = 6;
 * @return {string}
 */
proto.daytrader.command.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.daytrader.command.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 6, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.daytrader.log.prototype.toObject = function(opt_includeInstance) {
  return proto.daytrader.log.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.daytrader.log} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.daytrader.log.toObject = function(includeInstance, msg) {
  var obj = {
    command: jspb.Message.getFieldWithDefault(msg, 1, ""),
    serverName: jspb.Message.getFieldWithDefault(msg, 2, ""),
    transactionNum: jspb.Message.getFieldWithDefault(msg, 3, 0),
    username: jspb.Message.getFieldWithDefault(msg, 4, ""),
    stockSymbol: jspb.Message.getFieldWithDefault(msg, 5, ""),
    price: +jspb.Message.getFieldWithDefault(msg, 6, 0.0),
    funds: +jspb.Message.getFieldWithDefault(msg, 7, 0.0),
    filename: jspb.Message.getFieldWithDefault(msg, 8, ""),
    cryptoKey: jspb.Message.getFieldWithDefault(msg, 9, ""),
    quoteServerTime: jspb.Message.getFieldWithDefault(msg, 10, 0),
    accountAction: jspb.Message.getFieldWithDefault(msg, 11, ""),
    errorMessage: jspb.Message.getFieldWithDefault(msg, 12, ""),
    debugMessage: jspb.Message.getFieldWithDefault(msg, 13, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.daytrader.log}
 */
proto.daytrader.log.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.daytrader.log;
  return proto.daytrader.log.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.daytrader.log} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.daytrader.log}
 */
proto.daytrader.log.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setCommand(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setServerName(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setTransactionNum(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setUsername(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setStockSymbol(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setPrice(value);
      break;
    case 7:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setFunds(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setFilename(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setCryptoKey(value);
      break;
    case 10:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setQuoteServerTime(value);
      break;
    case 11:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountAction(value);
      break;
    case 12:
      var value = /** @type {string} */ (reader.readString());
      msg.setErrorMessage(value);
      break;
    case 13:
      var value = /** @type {string} */ (reader.readString());
      msg.setDebugMessage(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.daytrader.log.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.daytrader.log.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.daytrader.log} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.daytrader.log.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCommand();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getServerName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getTransactionNum();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = message.getUsername();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getStockSymbol();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getPrice();
  if (f !== 0.0) {
    writer.writeFloat(
      6,
      f
    );
  }
  f = message.getFunds();
  if (f !== 0.0) {
    writer.writeFloat(
      7,
      f
    );
  }
  f = message.getFilename();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getCryptoKey();
  if (f.length > 0) {
    writer.writeString(
      9,
      f
    );
  }
  f = message.getQuoteServerTime();
  if (f !== 0) {
    writer.writeInt64(
      10,
      f
    );
  }
  f = message.getAccountAction();
  if (f.length > 0) {
    writer.writeString(
      11,
      f
    );
  }
  f = message.getErrorMessage();
  if (f.length > 0) {
    writer.writeString(
      12,
      f
    );
  }
  f = message.getDebugMessage();
  if (f.length > 0) {
    writer.writeString(
      13,
      f
    );
  }
};


/**
 * optional string command = 1;
 * @return {string}
 */
proto.daytrader.log.prototype.getCommand = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.daytrader.log.prototype.setCommand = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string server_name = 2;
 * @return {string}
 */
proto.daytrader.log.prototype.getServerName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.daytrader.log.prototype.setServerName = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional int32 transaction_num = 3;
 * @return {number}
 */
proto.daytrader.log.prototype.getTransactionNum = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.daytrader.log.prototype.setTransactionNum = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional string username = 4;
 * @return {string}
 */
proto.daytrader.log.prototype.getUsername = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.daytrader.log.prototype.setUsername = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string stock_symbol = 5;
 * @return {string}
 */
proto.daytrader.log.prototype.getStockSymbol = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.daytrader.log.prototype.setStockSymbol = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional float price = 6;
 * @return {number}
 */
proto.daytrader.log.prototype.getPrice = function() {
  return /** @type {number} */ (+jspb.Message.getFieldWithDefault(this, 6, 0.0));
};


/** @param {number} value */
proto.daytrader.log.prototype.setPrice = function(value) {
  jspb.Message.setProto3FloatField(this, 6, value);
};


/**
 * optional float funds = 7;
 * @return {number}
 */
proto.daytrader.log.prototype.getFunds = function() {
  return /** @type {number} */ (+jspb.Message.getFieldWithDefault(this, 7, 0.0));
};


/** @param {number} value */
proto.daytrader.log.prototype.setFunds = function(value) {
  jspb.Message.setProto3FloatField(this, 7, value);
};


/**
 * optional string filename = 8;
 * @return {string}
 */
proto.daytrader.log.prototype.getFilename = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/** @param {string} value */
proto.daytrader.log.prototype.setFilename = function(value) {
  jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional string crypto_key = 9;
 * @return {string}
 */
proto.daytrader.log.prototype.getCryptoKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/** @param {string} value */
proto.daytrader.log.prototype.setCryptoKey = function(value) {
  jspb.Message.setProto3StringField(this, 9, value);
};


/**
 * optional int64 quote_server_time = 10;
 * @return {number}
 */
proto.daytrader.log.prototype.getQuoteServerTime = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 10, 0));
};


/** @param {number} value */
proto.daytrader.log.prototype.setQuoteServerTime = function(value) {
  jspb.Message.setProto3IntField(this, 10, value);
};


/**
 * optional string account_action = 11;
 * @return {string}
 */
proto.daytrader.log.prototype.getAccountAction = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 11, ""));
};


/** @param {string} value */
proto.daytrader.log.prototype.setAccountAction = function(value) {
  jspb.Message.setProto3StringField(this, 11, value);
};


/**
 * optional string error_message = 12;
 * @return {string}
 */
proto.daytrader.log.prototype.getErrorMessage = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 12, ""));
};


/** @param {string} value */
proto.daytrader.log.prototype.setErrorMessage = function(value) {
  jspb.Message.setProto3StringField(this, 12, value);
};


/**
 * optional string debug_message = 13;
 * @return {string}
 */
proto.daytrader.log.prototype.getDebugMessage = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 13, ""));
};


/** @param {string} value */
proto.daytrader.log.prototype.setDebugMessage = function(value) {
  jspb.Message.setProto3StringField(this, 13, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.daytrader.Response.prototype.toObject = function(opt_includeInstance) {
  return proto.daytrader.Response.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.daytrader.Response} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.daytrader.Response.toObject = function(includeInstance, msg) {
  var obj = {
    message: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.daytrader.Response}
 */
proto.daytrader.Response.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.daytrader.Response;
  return proto.daytrader.Response.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.daytrader.Response} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.daytrader.Response}
 */
proto.daytrader.Response.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setMessage(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.daytrader.Response.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.daytrader.Response.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.daytrader.Response} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.daytrader.Response.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMessage();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string message = 1;
 * @return {string}
 */
proto.daytrader.Response.prototype.getMessage = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.daytrader.Response.prototype.setMessage = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


goog.object.extend(exports, proto.daytrader);