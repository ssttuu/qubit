/**
 * @fileoverview
 * @enhanceable
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_api_annotations_pb = require('../google/api/annotations_pb.js');
var geometry_geometry_pb = require('../geometry/geometry_pb.js');
goog.exportSymbol('proto.render_operators.RenderOperator', null, global);
goog.exportSymbol('proto.render_operators.RenderOperatorRequest', null, global);

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
proto.render_operators.RenderOperatorRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.render_operators.RenderOperatorRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.render_operators.RenderOperatorRequest.displayName = 'proto.render_operators.RenderOperatorRequest';
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
proto.render_operators.RenderOperatorRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.render_operators.RenderOperatorRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.render_operators.RenderOperatorRequest} msg The msg instance to transform.
 * @return {!Object}
 */
proto.render_operators.RenderOperatorRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    renderKey: jspb.Message.getFieldWithDefault(msg, 1, "")
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
 * @return {!proto.render_operators.RenderOperatorRequest}
 */
proto.render_operators.RenderOperatorRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.render_operators.RenderOperatorRequest;
  return proto.render_operators.RenderOperatorRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.render_operators.RenderOperatorRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.render_operators.RenderOperatorRequest}
 */
proto.render_operators.RenderOperatorRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setRenderKey(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Class method variant: serializes the given message to binary data
 * (in protobuf wire format), writing to the given BinaryWriter.
 * @param {!proto.render_operators.RenderOperatorRequest} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.render_operators.RenderOperatorRequest.serializeBinaryToWriter = function(message, writer) {
  message.serializeBinaryToWriter(writer);
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.render_operators.RenderOperatorRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  this.serializeBinaryToWriter(writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the message to binary data (in protobuf wire format),
 * writing to the given BinaryWriter.
 * @param {!jspb.BinaryWriter} writer
 */
proto.render_operators.RenderOperatorRequest.prototype.serializeBinaryToWriter = function (writer) {
  var f = undefined;
  f = this.getRenderKey();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string render_key = 1;
 * @return {string}
 */
proto.render_operators.RenderOperatorRequest.prototype.getRenderKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.render_operators.RenderOperatorRequest.prototype.setRenderKey = function(value) {
  jspb.Message.setField(this, 1, value);
};



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
proto.render_operators.RenderOperator = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.render_operators.RenderOperator, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.render_operators.RenderOperator.displayName = 'proto.render_operators.RenderOperator';
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
proto.render_operators.RenderOperator.prototype.toObject = function(opt_includeInstance) {
  return proto.render_operators.RenderOperator.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.render_operators.RenderOperator} msg The msg instance to transform.
 * @return {!Object}
 */
proto.render_operators.RenderOperator.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    sceneId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    sceneVersion: jspb.Message.getFieldWithDefault(msg, 3, 0),
    operatorId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    time: jspb.Message.getFieldWithDefault(msg, 5, ""),
    boundingBox: (f = msg.getBoundingBox()) && geometry_geometry_pb.BoundingBox2D.toObject(includeInstance, f),
    inputIds: msg.getInputIds_asB64(),
    parameters: msg.getParameters_asB64()
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
 * @return {!proto.render_operators.RenderOperator}
 */
proto.render_operators.RenderOperator.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.render_operators.RenderOperator;
  return proto.render_operators.RenderOperator.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.render_operators.RenderOperator} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.render_operators.RenderOperator}
 */
proto.render_operators.RenderOperator.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setSceneId(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setSceneVersion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setOperatorId(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setTime(value);
      break;
    case 6:
      var value = new geometry_geometry_pb.BoundingBox2D;
      reader.readMessage(value,geometry_geometry_pb.BoundingBox2D.deserializeBinaryFromReader);
      msg.setBoundingBox(value);
      break;
    case 7:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setInputIds(value);
      break;
    case 8:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setParameters(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Class method variant: serializes the given message to binary data
 * (in protobuf wire format), writing to the given BinaryWriter.
 * @param {!proto.render_operators.RenderOperator} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.render_operators.RenderOperator.serializeBinaryToWriter = function(message, writer) {
  message.serializeBinaryToWriter(writer);
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.render_operators.RenderOperator.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  this.serializeBinaryToWriter(writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the message to binary data (in protobuf wire format),
 * writing to the given BinaryWriter.
 * @param {!jspb.BinaryWriter} writer
 */
proto.render_operators.RenderOperator.prototype.serializeBinaryToWriter = function (writer) {
  var f = undefined;
  f = this.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = this.getSceneId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = this.getSceneVersion();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = this.getOperatorId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = this.getTime();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = this.getBoundingBox();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      geometry_geometry_pb.BoundingBox2D.serializeBinaryToWriter
    );
  }
  f = this.getInputIds_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      7,
      f
    );
  }
  f = this.getParameters_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      8,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.render_operators.RenderOperator.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.render_operators.RenderOperator.prototype.setId = function(value) {
  jspb.Message.setField(this, 1, value);
};


/**
 * optional string scene_id = 2;
 * @return {string}
 */
proto.render_operators.RenderOperator.prototype.getSceneId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.render_operators.RenderOperator.prototype.setSceneId = function(value) {
  jspb.Message.setField(this, 2, value);
};


/**
 * optional int32 scene_version = 3;
 * @return {number}
 */
proto.render_operators.RenderOperator.prototype.getSceneVersion = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.render_operators.RenderOperator.prototype.setSceneVersion = function(value) {
  jspb.Message.setField(this, 3, value);
};


/**
 * optional string operator_id = 4;
 * @return {string}
 */
proto.render_operators.RenderOperator.prototype.getOperatorId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.render_operators.RenderOperator.prototype.setOperatorId = function(value) {
  jspb.Message.setField(this, 4, value);
};


/**
 * optional string time = 5;
 * @return {string}
 */
proto.render_operators.RenderOperator.prototype.getTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.render_operators.RenderOperator.prototype.setTime = function(value) {
  jspb.Message.setField(this, 5, value);
};


/**
 * optional geometry.BoundingBox2D bounding_box = 6;
 * @return {?proto.geometry.BoundingBox2D}
 */
proto.render_operators.RenderOperator.prototype.getBoundingBox = function() {
  return /** @type{?proto.geometry.BoundingBox2D} */ (
    jspb.Message.getWrapperField(this, geometry_geometry_pb.BoundingBox2D, 6));
};


/** @param {?proto.geometry.BoundingBox2D|undefined} value */
proto.render_operators.RenderOperator.prototype.setBoundingBox = function(value) {
  jspb.Message.setWrapperField(this, 6, value);
};


proto.render_operators.RenderOperator.prototype.clearBoundingBox = function() {
  this.setBoundingBox(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.render_operators.RenderOperator.prototype.hasBoundingBox = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional bytes input_ids = 7;
 * @return {!(string|Uint8Array)}
 */
proto.render_operators.RenderOperator.prototype.getInputIds = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * optional bytes input_ids = 7;
 * This is a type-conversion wrapper around `getInputIds()`
 * @return {string}
 */
proto.render_operators.RenderOperator.prototype.getInputIds_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getInputIds()));
};


/**
 * optional bytes input_ids = 7;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getInputIds()`
 * @return {!Uint8Array}
 */
proto.render_operators.RenderOperator.prototype.getInputIds_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getInputIds()));
};


/** @param {!(string|Uint8Array)} value */
proto.render_operators.RenderOperator.prototype.setInputIds = function(value) {
  jspb.Message.setField(this, 7, value);
};


/**
 * optional bytes parameters = 8;
 * @return {!(string|Uint8Array)}
 */
proto.render_operators.RenderOperator.prototype.getParameters = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * optional bytes parameters = 8;
 * This is a type-conversion wrapper around `getParameters()`
 * @return {string}
 */
proto.render_operators.RenderOperator.prototype.getParameters_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getParameters()));
};


/**
 * optional bytes parameters = 8;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getParameters()`
 * @return {!Uint8Array}
 */
proto.render_operators.RenderOperator.prototype.getParameters_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getParameters()));
};


/** @param {!(string|Uint8Array)} value */
proto.render_operators.RenderOperator.prototype.setParameters = function(value) {
  jspb.Message.setField(this, 8, value);
};


goog.object.extend(exports, proto.render_operators);
