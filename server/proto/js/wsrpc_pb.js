/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
import * as $protobuf from "protobufjs/minimal";

// Common aliases
const $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
const $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

export const orderbooks = $root.orderbooks = (() => {

    /**
     * Namespace orderbooks.
     * @exports orderbooks
     * @namespace
     */
    const orderbooks = {};

    orderbooks.OrderBooks = (function() {

        /**
         * Constructs a new OrderBooks service.
         * @memberof orderbooks
         * @classdesc Represents an OrderBooks
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function OrderBooks(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (OrderBooks.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = OrderBooks;

        /**
         * Creates new OrderBooks service using the specified rpc implementation.
         * @function create
         * @memberof orderbooks.OrderBooks
         * @static
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         * @returns {OrderBooks} RPC service. Useful where requests and/or responses are streamed.
         */
        OrderBooks.create = function create(rpcImpl, requestDelimited, responseDelimited) {
            return new this(rpcImpl, requestDelimited, responseDelimited);
        };

        /**
         * Callback as used by {@link orderbooks.OrderBooks#getExchanges}.
         * @memberof orderbooks.OrderBooks
         * @typedef GetExchangesCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {orderbooks.GetExchangesResponseMsg} [response] GetExchangesResponseMsg
         */

        /**
         * Calls GetExchanges.
         * @function getExchanges
         * @memberof orderbooks.OrderBooks
         * @instance
         * @param {orderbooks.IEmptyMsg} request EmptyMsg message or plain object
         * @param {orderbooks.OrderBooks.GetExchangesCallback} callback Node-style callback called with the error, if any, and GetExchangesResponseMsg
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(OrderBooks.prototype.getExchanges = function getExchanges(request, callback) {
            return this.rpcCall(getExchanges, $root.orderbooks.EmptyMsg, $root.orderbooks.GetExchangesResponseMsg, request, callback);
        }, "name", { value: "GetExchanges" });

        /**
         * Calls GetExchanges.
         * @function getExchanges
         * @memberof orderbooks.OrderBooks
         * @instance
         * @param {orderbooks.IEmptyMsg} request EmptyMsg message or plain object
         * @returns {Promise<orderbooks.GetExchangesResponseMsg>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link orderbooks.OrderBooks#queryEvents}.
         * @memberof orderbooks.OrderBooks
         * @typedef QueryEventsCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {orderbooks.SnapshotMsg} [response] SnapshotMsg
         */

        /**
         * Calls QueryEvents.
         * @function queryEvents
         * @memberof orderbooks.OrderBooks
         * @instance
         * @param {orderbooks.IEventsQueryMsg} request EventsQueryMsg message or plain object
         * @param {orderbooks.OrderBooks.QueryEventsCallback} callback Node-style callback called with the error, if any, and SnapshotMsg
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(OrderBooks.prototype.queryEvents = function queryEvents(request, callback) {
            return this.rpcCall(queryEvents, $root.orderbooks.EventsQueryMsg, $root.orderbooks.SnapshotMsg, request, callback);
        }, "name", { value: "QueryEvents" });

        /**
         * Calls QueryEvents.
         * @function queryEvents
         * @memberof orderbooks.OrderBooks
         * @instance
         * @param {orderbooks.IEventsQueryMsg} request EventsQueryMsg message or plain object
         * @returns {Promise<orderbooks.SnapshotMsg>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link orderbooks.OrderBooks#getLiveData}.
         * @memberof orderbooks.OrderBooks
         * @typedef GetLiveDataCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {orderbooks.SnapshotMsg} [response] SnapshotMsg
         */

        /**
         * Calls GetLiveData.
         * @function getLiveData
         * @memberof orderbooks.OrderBooks
         * @instance
         * @param {orderbooks.IEventsQueryMsg} request EventsQueryMsg message or plain object
         * @param {orderbooks.OrderBooks.GetLiveDataCallback} callback Node-style callback called with the error, if any, and SnapshotMsg
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(OrderBooks.prototype.getLiveData = function getLiveData(request, callback) {
            return this.rpcCall(getLiveData, $root.orderbooks.EventsQueryMsg, $root.orderbooks.SnapshotMsg, request, callback);
        }, "name", { value: "GetLiveData" });

        /**
         * Calls GetLiveData.
         * @function getLiveData
         * @memberof orderbooks.OrderBooks
         * @instance
         * @param {orderbooks.IEventsQueryMsg} request EventsQueryMsg message or plain object
         * @returns {Promise<orderbooks.SnapshotMsg>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link orderbooks.OrderBooks#getCachedData}.
         * @memberof orderbooks.OrderBooks
         * @typedef GetCachedDataCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {orderbooks.SnapshotMsg} [response] SnapshotMsg
         */

        /**
         * Calls GetCachedData.
         * @function getCachedData
         * @memberof orderbooks.OrderBooks
         * @instance
         * @param {orderbooks.IEventsQueryMsg} request EventsQueryMsg message or plain object
         * @param {orderbooks.OrderBooks.GetCachedDataCallback} callback Node-style callback called with the error, if any, and SnapshotMsg
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(OrderBooks.prototype.getCachedData = function getCachedData(request, callback) {
            return this.rpcCall(getCachedData, $root.orderbooks.EventsQueryMsg, $root.orderbooks.SnapshotMsg, request, callback);
        }, "name", { value: "GetCachedData" });

        /**
         * Calls GetCachedData.
         * @function getCachedData
         * @memberof orderbooks.OrderBooks
         * @instance
         * @param {orderbooks.IEventsQueryMsg} request EventsQueryMsg message or plain object
         * @returns {Promise<orderbooks.SnapshotMsg>} Promise
         * @variation 2
         */

        return OrderBooks;
    })();

    orderbooks.EmptyMsg = (function() {

        /**
         * Properties of an EmptyMsg.
         * @memberof orderbooks
         * @interface IEmptyMsg
         */

        /**
         * Constructs a new EmptyMsg.
         * @memberof orderbooks
         * @classdesc Represents an EmptyMsg.
         * @implements IEmptyMsg
         * @constructor
         * @param {orderbooks.IEmptyMsg=} [properties] Properties to set
         */
        function EmptyMsg(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new EmptyMsg instance using the specified properties.
         * @function create
         * @memberof orderbooks.EmptyMsg
         * @static
         * @param {orderbooks.IEmptyMsg=} [properties] Properties to set
         * @returns {orderbooks.EmptyMsg} EmptyMsg instance
         */
        EmptyMsg.create = function create(properties) {
            return new EmptyMsg(properties);
        };

        /**
         * Encodes the specified EmptyMsg message. Does not implicitly {@link orderbooks.EmptyMsg.verify|verify} messages.
         * @function encode
         * @memberof orderbooks.EmptyMsg
         * @static
         * @param {orderbooks.IEmptyMsg} message EmptyMsg message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        EmptyMsg.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            return writer;
        };

        /**
         * Encodes the specified EmptyMsg message, length delimited. Does not implicitly {@link orderbooks.EmptyMsg.verify|verify} messages.
         * @function encodeDelimited
         * @memberof orderbooks.EmptyMsg
         * @static
         * @param {orderbooks.IEmptyMsg} message EmptyMsg message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        EmptyMsg.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an EmptyMsg message from the specified reader or buffer.
         * @function decode
         * @memberof orderbooks.EmptyMsg
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {orderbooks.EmptyMsg} EmptyMsg
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        EmptyMsg.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.orderbooks.EmptyMsg();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an EmptyMsg message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof orderbooks.EmptyMsg
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {orderbooks.EmptyMsg} EmptyMsg
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        EmptyMsg.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an EmptyMsg message.
         * @function verify
         * @memberof orderbooks.EmptyMsg
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        EmptyMsg.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            return null;
        };

        /**
         * Creates an EmptyMsg message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof orderbooks.EmptyMsg
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {orderbooks.EmptyMsg} EmptyMsg
         */
        EmptyMsg.fromObject = function fromObject(object) {
            if (object instanceof $root.orderbooks.EmptyMsg)
                return object;
            return new $root.orderbooks.EmptyMsg();
        };

        /**
         * Creates a plain object from an EmptyMsg message. Also converts values to other types if specified.
         * @function toObject
         * @memberof orderbooks.EmptyMsg
         * @static
         * @param {orderbooks.EmptyMsg} message EmptyMsg
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        EmptyMsg.toObject = function toObject() {
            return {};
        };

        /**
         * Converts this EmptyMsg to JSON.
         * @function toJSON
         * @memberof orderbooks.EmptyMsg
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        EmptyMsg.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return EmptyMsg;
    })();

    orderbooks.GetExchangesResponseMsg = (function() {

        /**
         * Properties of a GetExchangesResponseMsg.
         * @memberof orderbooks
         * @interface IGetExchangesResponseMsg
         * @property {number|null} [count] GetExchangesResponseMsg count
         * @property {Array.<orderbooks.IExchangeMsg>|null} [exchanges] GetExchangesResponseMsg exchanges
         */

        /**
         * Constructs a new GetExchangesResponseMsg.
         * @memberof orderbooks
         * @classdesc Represents a GetExchangesResponseMsg.
         * @implements IGetExchangesResponseMsg
         * @constructor
         * @param {orderbooks.IGetExchangesResponseMsg=} [properties] Properties to set
         */
        function GetExchangesResponseMsg(properties) {
            this.exchanges = [];
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * GetExchangesResponseMsg count.
         * @member {number} count
         * @memberof orderbooks.GetExchangesResponseMsg
         * @instance
         */
        GetExchangesResponseMsg.prototype.count = 0;

        /**
         * GetExchangesResponseMsg exchanges.
         * @member {Array.<orderbooks.IExchangeMsg>} exchanges
         * @memberof orderbooks.GetExchangesResponseMsg
         * @instance
         */
        GetExchangesResponseMsg.prototype.exchanges = $util.emptyArray;

        /**
         * Creates a new GetExchangesResponseMsg instance using the specified properties.
         * @function create
         * @memberof orderbooks.GetExchangesResponseMsg
         * @static
         * @param {orderbooks.IGetExchangesResponseMsg=} [properties] Properties to set
         * @returns {orderbooks.GetExchangesResponseMsg} GetExchangesResponseMsg instance
         */
        GetExchangesResponseMsg.create = function create(properties) {
            return new GetExchangesResponseMsg(properties);
        };

        /**
         * Encodes the specified GetExchangesResponseMsg message. Does not implicitly {@link orderbooks.GetExchangesResponseMsg.verify|verify} messages.
         * @function encode
         * @memberof orderbooks.GetExchangesResponseMsg
         * @static
         * @param {orderbooks.IGetExchangesResponseMsg} message GetExchangesResponseMsg message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        GetExchangesResponseMsg.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.count != null && Object.hasOwnProperty.call(message, "count"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.count);
            if (message.exchanges != null && message.exchanges.length)
                for (let i = 0; i < message.exchanges.length; ++i)
                    $root.orderbooks.ExchangeMsg.encode(message.exchanges[i], writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified GetExchangesResponseMsg message, length delimited. Does not implicitly {@link orderbooks.GetExchangesResponseMsg.verify|verify} messages.
         * @function encodeDelimited
         * @memberof orderbooks.GetExchangesResponseMsg
         * @static
         * @param {orderbooks.IGetExchangesResponseMsg} message GetExchangesResponseMsg message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        GetExchangesResponseMsg.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a GetExchangesResponseMsg message from the specified reader or buffer.
         * @function decode
         * @memberof orderbooks.GetExchangesResponseMsg
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {orderbooks.GetExchangesResponseMsg} GetExchangesResponseMsg
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        GetExchangesResponseMsg.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.orderbooks.GetExchangesResponseMsg();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.count = reader.uint32();
                    break;
                case 2:
                    if (!(message.exchanges && message.exchanges.length))
                        message.exchanges = [];
                    message.exchanges.push($root.orderbooks.ExchangeMsg.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a GetExchangesResponseMsg message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof orderbooks.GetExchangesResponseMsg
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {orderbooks.GetExchangesResponseMsg} GetExchangesResponseMsg
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        GetExchangesResponseMsg.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a GetExchangesResponseMsg message.
         * @function verify
         * @memberof orderbooks.GetExchangesResponseMsg
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        GetExchangesResponseMsg.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.count != null && message.hasOwnProperty("count"))
                if (!$util.isInteger(message.count))
                    return "count: integer expected";
            if (message.exchanges != null && message.hasOwnProperty("exchanges")) {
                if (!Array.isArray(message.exchanges))
                    return "exchanges: array expected";
                for (let i = 0; i < message.exchanges.length; ++i) {
                    let error = $root.orderbooks.ExchangeMsg.verify(message.exchanges[i]);
                    if (error)
                        return "exchanges." + error;
                }
            }
            return null;
        };

        /**
         * Creates a GetExchangesResponseMsg message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof orderbooks.GetExchangesResponseMsg
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {orderbooks.GetExchangesResponseMsg} GetExchangesResponseMsg
         */
        GetExchangesResponseMsg.fromObject = function fromObject(object) {
            if (object instanceof $root.orderbooks.GetExchangesResponseMsg)
                return object;
            let message = new $root.orderbooks.GetExchangesResponseMsg();
            if (object.count != null)
                message.count = object.count >>> 0;
            if (object.exchanges) {
                if (!Array.isArray(object.exchanges))
                    throw TypeError(".orderbooks.GetExchangesResponseMsg.exchanges: array expected");
                message.exchanges = [];
                for (let i = 0; i < object.exchanges.length; ++i) {
                    if (typeof object.exchanges[i] !== "object")
                        throw TypeError(".orderbooks.GetExchangesResponseMsg.exchanges: object expected");
                    message.exchanges[i] = $root.orderbooks.ExchangeMsg.fromObject(object.exchanges[i]);
                }
            }
            return message;
        };

        /**
         * Creates a plain object from a GetExchangesResponseMsg message. Also converts values to other types if specified.
         * @function toObject
         * @memberof orderbooks.GetExchangesResponseMsg
         * @static
         * @param {orderbooks.GetExchangesResponseMsg} message GetExchangesResponseMsg
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        GetExchangesResponseMsg.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.arrays || options.defaults)
                object.exchanges = [];
            if (options.defaults)
                object.count = 0;
            if (message.count != null && message.hasOwnProperty("count"))
                object.count = message.count;
            if (message.exchanges && message.exchanges.length) {
                object.exchanges = [];
                for (let j = 0; j < message.exchanges.length; ++j)
                    object.exchanges[j] = $root.orderbooks.ExchangeMsg.toObject(message.exchanges[j], options);
            }
            return object;
        };

        /**
         * Converts this GetExchangesResponseMsg to JSON.
         * @function toJSON
         * @memberof orderbooks.GetExchangesResponseMsg
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        GetExchangesResponseMsg.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return GetExchangesResponseMsg;
    })();

    orderbooks.ExchangeMsg = (function() {

        /**
         * Properties of an ExchangeMsg.
         * @memberof orderbooks
         * @interface IExchangeMsg
         * @property {string|null} [id] ExchangeMsg id
         * @property {string|null} [name] ExchangeMsg name
         * @property {google.protobuf.ITimestamp|null} [dateStart] ExchangeMsg dateStart
         * @property {google.protobuf.ITimestamp|null} [dateEnd] ExchangeMsg dateEnd
         * @property {Array.<orderbooks.IPairMsg>|null} [pairs] ExchangeMsg pairs
         */

        /**
         * Constructs a new ExchangeMsg.
         * @memberof orderbooks
         * @classdesc Represents an ExchangeMsg.
         * @implements IExchangeMsg
         * @constructor
         * @param {orderbooks.IExchangeMsg=} [properties] Properties to set
         */
        function ExchangeMsg(properties) {
            this.pairs = [];
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * ExchangeMsg id.
         * @member {string} id
         * @memberof orderbooks.ExchangeMsg
         * @instance
         */
        ExchangeMsg.prototype.id = "";

        /**
         * ExchangeMsg name.
         * @member {string} name
         * @memberof orderbooks.ExchangeMsg
         * @instance
         */
        ExchangeMsg.prototype.name = "";

        /**
         * ExchangeMsg dateStart.
         * @member {google.protobuf.ITimestamp|null|undefined} dateStart
         * @memberof orderbooks.ExchangeMsg
         * @instance
         */
        ExchangeMsg.prototype.dateStart = null;

        /**
         * ExchangeMsg dateEnd.
         * @member {google.protobuf.ITimestamp|null|undefined} dateEnd
         * @memberof orderbooks.ExchangeMsg
         * @instance
         */
        ExchangeMsg.prototype.dateEnd = null;

        /**
         * ExchangeMsg pairs.
         * @member {Array.<orderbooks.IPairMsg>} pairs
         * @memberof orderbooks.ExchangeMsg
         * @instance
         */
        ExchangeMsg.prototype.pairs = $util.emptyArray;

        /**
         * Creates a new ExchangeMsg instance using the specified properties.
         * @function create
         * @memberof orderbooks.ExchangeMsg
         * @static
         * @param {orderbooks.IExchangeMsg=} [properties] Properties to set
         * @returns {orderbooks.ExchangeMsg} ExchangeMsg instance
         */
        ExchangeMsg.create = function create(properties) {
            return new ExchangeMsg(properties);
        };

        /**
         * Encodes the specified ExchangeMsg message. Does not implicitly {@link orderbooks.ExchangeMsg.verify|verify} messages.
         * @function encode
         * @memberof orderbooks.ExchangeMsg
         * @static
         * @param {orderbooks.IExchangeMsg} message ExchangeMsg message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ExchangeMsg.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.id != null && Object.hasOwnProperty.call(message, "id"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.id);
            if (message.name != null && Object.hasOwnProperty.call(message, "name"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.name);
            if (message.dateStart != null && Object.hasOwnProperty.call(message, "dateStart"))
                $root.google.protobuf.Timestamp.encode(message.dateStart, writer.uint32(/* id 3, wireType 2 =*/26).fork()).ldelim();
            if (message.dateEnd != null && Object.hasOwnProperty.call(message, "dateEnd"))
                $root.google.protobuf.Timestamp.encode(message.dateEnd, writer.uint32(/* id 4, wireType 2 =*/34).fork()).ldelim();
            if (message.pairs != null && message.pairs.length)
                for (let i = 0; i < message.pairs.length; ++i)
                    $root.orderbooks.PairMsg.encode(message.pairs[i], writer.uint32(/* id 5, wireType 2 =*/42).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified ExchangeMsg message, length delimited. Does not implicitly {@link orderbooks.ExchangeMsg.verify|verify} messages.
         * @function encodeDelimited
         * @memberof orderbooks.ExchangeMsg
         * @static
         * @param {orderbooks.IExchangeMsg} message ExchangeMsg message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ExchangeMsg.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an ExchangeMsg message from the specified reader or buffer.
         * @function decode
         * @memberof orderbooks.ExchangeMsg
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {orderbooks.ExchangeMsg} ExchangeMsg
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        ExchangeMsg.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.orderbooks.ExchangeMsg();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                case 2:
                    message.name = reader.string();
                    break;
                case 3:
                    message.dateStart = $root.google.protobuf.Timestamp.decode(reader, reader.uint32());
                    break;
                case 4:
                    message.dateEnd = $root.google.protobuf.Timestamp.decode(reader, reader.uint32());
                    break;
                case 5:
                    if (!(message.pairs && message.pairs.length))
                        message.pairs = [];
                    message.pairs.push($root.orderbooks.PairMsg.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an ExchangeMsg message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof orderbooks.ExchangeMsg
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {orderbooks.ExchangeMsg} ExchangeMsg
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        ExchangeMsg.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an ExchangeMsg message.
         * @function verify
         * @memberof orderbooks.ExchangeMsg
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        ExchangeMsg.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.id != null && message.hasOwnProperty("id"))
                if (!$util.isString(message.id))
                    return "id: string expected";
            if (message.name != null && message.hasOwnProperty("name"))
                if (!$util.isString(message.name))
                    return "name: string expected";
            if (message.dateStart != null && message.hasOwnProperty("dateStart")) {
                let error = $root.google.protobuf.Timestamp.verify(message.dateStart);
                if (error)
                    return "dateStart." + error;
            }
            if (message.dateEnd != null && message.hasOwnProperty("dateEnd")) {
                let error = $root.google.protobuf.Timestamp.verify(message.dateEnd);
                if (error)
                    return "dateEnd." + error;
            }
            if (message.pairs != null && message.hasOwnProperty("pairs")) {
                if (!Array.isArray(message.pairs))
                    return "pairs: array expected";
                for (let i = 0; i < message.pairs.length; ++i) {
                    let error = $root.orderbooks.PairMsg.verify(message.pairs[i]);
                    if (error)
                        return "pairs." + error;
                }
            }
            return null;
        };

        /**
         * Creates an ExchangeMsg message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof orderbooks.ExchangeMsg
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {orderbooks.ExchangeMsg} ExchangeMsg
         */
        ExchangeMsg.fromObject = function fromObject(object) {
            if (object instanceof $root.orderbooks.ExchangeMsg)
                return object;
            let message = new $root.orderbooks.ExchangeMsg();
            if (object.id != null)
                message.id = String(object.id);
            if (object.name != null)
                message.name = String(object.name);
            if (object.dateStart != null) {
                if (typeof object.dateStart !== "object")
                    throw TypeError(".orderbooks.ExchangeMsg.dateStart: object expected");
                message.dateStart = $root.google.protobuf.Timestamp.fromObject(object.dateStart);
            }
            if (object.dateEnd != null) {
                if (typeof object.dateEnd !== "object")
                    throw TypeError(".orderbooks.ExchangeMsg.dateEnd: object expected");
                message.dateEnd = $root.google.protobuf.Timestamp.fromObject(object.dateEnd);
            }
            if (object.pairs) {
                if (!Array.isArray(object.pairs))
                    throw TypeError(".orderbooks.ExchangeMsg.pairs: array expected");
                message.pairs = [];
                for (let i = 0; i < object.pairs.length; ++i) {
                    if (typeof object.pairs[i] !== "object")
                        throw TypeError(".orderbooks.ExchangeMsg.pairs: object expected");
                    message.pairs[i] = $root.orderbooks.PairMsg.fromObject(object.pairs[i]);
                }
            }
            return message;
        };

        /**
         * Creates a plain object from an ExchangeMsg message. Also converts values to other types if specified.
         * @function toObject
         * @memberof orderbooks.ExchangeMsg
         * @static
         * @param {orderbooks.ExchangeMsg} message ExchangeMsg
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        ExchangeMsg.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.arrays || options.defaults)
                object.pairs = [];
            if (options.defaults) {
                object.id = "";
                object.name = "";
                object.dateStart = null;
                object.dateEnd = null;
            }
            if (message.id != null && message.hasOwnProperty("id"))
                object.id = message.id;
            if (message.name != null && message.hasOwnProperty("name"))
                object.name = message.name;
            if (message.dateStart != null && message.hasOwnProperty("dateStart"))
                object.dateStart = $root.google.protobuf.Timestamp.toObject(message.dateStart, options);
            if (message.dateEnd != null && message.hasOwnProperty("dateEnd"))
                object.dateEnd = $root.google.protobuf.Timestamp.toObject(message.dateEnd, options);
            if (message.pairs && message.pairs.length) {
                object.pairs = [];
                for (let j = 0; j < message.pairs.length; ++j)
                    object.pairs[j] = $root.orderbooks.PairMsg.toObject(message.pairs[j], options);
            }
            return object;
        };

        /**
         * Converts this ExchangeMsg to JSON.
         * @function toJSON
         * @memberof orderbooks.ExchangeMsg
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        ExchangeMsg.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return ExchangeMsg;
    })();

    orderbooks.PairMsg = (function() {

        /**
         * Properties of a PairMsg.
         * @memberof orderbooks
         * @interface IPairMsg
         * @property {string|null} [base] PairMsg base
         * @property {string|null} [quote] PairMsg quote
         * @property {google.protobuf.ITimestamp|null} [dateStart] PairMsg dateStart
         * @property {google.protobuf.ITimestamp|null} [dateEnd] PairMsg dateEnd
         */

        /**
         * Constructs a new PairMsg.
         * @memberof orderbooks
         * @classdesc Represents a PairMsg.
         * @implements IPairMsg
         * @constructor
         * @param {orderbooks.IPairMsg=} [properties] Properties to set
         */
        function PairMsg(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * PairMsg base.
         * @member {string} base
         * @memberof orderbooks.PairMsg
         * @instance
         */
        PairMsg.prototype.base = "";

        /**
         * PairMsg quote.
         * @member {string} quote
         * @memberof orderbooks.PairMsg
         * @instance
         */
        PairMsg.prototype.quote = "";

        /**
         * PairMsg dateStart.
         * @member {google.protobuf.ITimestamp|null|undefined} dateStart
         * @memberof orderbooks.PairMsg
         * @instance
         */
        PairMsg.prototype.dateStart = null;

        /**
         * PairMsg dateEnd.
         * @member {google.protobuf.ITimestamp|null|undefined} dateEnd
         * @memberof orderbooks.PairMsg
         * @instance
         */
        PairMsg.prototype.dateEnd = null;

        /**
         * Creates a new PairMsg instance using the specified properties.
         * @function create
         * @memberof orderbooks.PairMsg
         * @static
         * @param {orderbooks.IPairMsg=} [properties] Properties to set
         * @returns {orderbooks.PairMsg} PairMsg instance
         */
        PairMsg.create = function create(properties) {
            return new PairMsg(properties);
        };

        /**
         * Encodes the specified PairMsg message. Does not implicitly {@link orderbooks.PairMsg.verify|verify} messages.
         * @function encode
         * @memberof orderbooks.PairMsg
         * @static
         * @param {orderbooks.IPairMsg} message PairMsg message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        PairMsg.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.base != null && Object.hasOwnProperty.call(message, "base"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.base);
            if (message.quote != null && Object.hasOwnProperty.call(message, "quote"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.quote);
            if (message.dateStart != null && Object.hasOwnProperty.call(message, "dateStart"))
                $root.google.protobuf.Timestamp.encode(message.dateStart, writer.uint32(/* id 3, wireType 2 =*/26).fork()).ldelim();
            if (message.dateEnd != null && Object.hasOwnProperty.call(message, "dateEnd"))
                $root.google.protobuf.Timestamp.encode(message.dateEnd, writer.uint32(/* id 4, wireType 2 =*/34).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified PairMsg message, length delimited. Does not implicitly {@link orderbooks.PairMsg.verify|verify} messages.
         * @function encodeDelimited
         * @memberof orderbooks.PairMsg
         * @static
         * @param {orderbooks.IPairMsg} message PairMsg message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        PairMsg.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a PairMsg message from the specified reader or buffer.
         * @function decode
         * @memberof orderbooks.PairMsg
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {orderbooks.PairMsg} PairMsg
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        PairMsg.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.orderbooks.PairMsg();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.base = reader.string();
                    break;
                case 2:
                    message.quote = reader.string();
                    break;
                case 3:
                    message.dateStart = $root.google.protobuf.Timestamp.decode(reader, reader.uint32());
                    break;
                case 4:
                    message.dateEnd = $root.google.protobuf.Timestamp.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a PairMsg message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof orderbooks.PairMsg
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {orderbooks.PairMsg} PairMsg
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        PairMsg.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a PairMsg message.
         * @function verify
         * @memberof orderbooks.PairMsg
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        PairMsg.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.base != null && message.hasOwnProperty("base"))
                if (!$util.isString(message.base))
                    return "base: string expected";
            if (message.quote != null && message.hasOwnProperty("quote"))
                if (!$util.isString(message.quote))
                    return "quote: string expected";
            if (message.dateStart != null && message.hasOwnProperty("dateStart")) {
                let error = $root.google.protobuf.Timestamp.verify(message.dateStart);
                if (error)
                    return "dateStart." + error;
            }
            if (message.dateEnd != null && message.hasOwnProperty("dateEnd")) {
                let error = $root.google.protobuf.Timestamp.verify(message.dateEnd);
                if (error)
                    return "dateEnd." + error;
            }
            return null;
        };

        /**
         * Creates a PairMsg message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof orderbooks.PairMsg
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {orderbooks.PairMsg} PairMsg
         */
        PairMsg.fromObject = function fromObject(object) {
            if (object instanceof $root.orderbooks.PairMsg)
                return object;
            let message = new $root.orderbooks.PairMsg();
            if (object.base != null)
                message.base = String(object.base);
            if (object.quote != null)
                message.quote = String(object.quote);
            if (object.dateStart != null) {
                if (typeof object.dateStart !== "object")
                    throw TypeError(".orderbooks.PairMsg.dateStart: object expected");
                message.dateStart = $root.google.protobuf.Timestamp.fromObject(object.dateStart);
            }
            if (object.dateEnd != null) {
                if (typeof object.dateEnd !== "object")
                    throw TypeError(".orderbooks.PairMsg.dateEnd: object expected");
                message.dateEnd = $root.google.protobuf.Timestamp.fromObject(object.dateEnd);
            }
            return message;
        };

        /**
         * Creates a plain object from a PairMsg message. Also converts values to other types if specified.
         * @function toObject
         * @memberof orderbooks.PairMsg
         * @static
         * @param {orderbooks.PairMsg} message PairMsg
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        PairMsg.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.base = "";
                object.quote = "";
                object.dateStart = null;
                object.dateEnd = null;
            }
            if (message.base != null && message.hasOwnProperty("base"))
                object.base = message.base;
            if (message.quote != null && message.hasOwnProperty("quote"))
                object.quote = message.quote;
            if (message.dateStart != null && message.hasOwnProperty("dateStart"))
                object.dateStart = $root.google.protobuf.Timestamp.toObject(message.dateStart, options);
            if (message.dateEnd != null && message.hasOwnProperty("dateEnd"))
                object.dateEnd = $root.google.protobuf.Timestamp.toObject(message.dateEnd, options);
            return object;
        };

        /**
         * Converts this PairMsg to JSON.
         * @function toJSON
         * @memberof orderbooks.PairMsg
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        PairMsg.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return PairMsg;
    })();

    orderbooks.EventsQueryMsg = (function() {

        /**
         * Properties of an EventsQueryMsg.
         * @memberof orderbooks
         * @interface IEventsQueryMsg
         * @property {string|null} [exchange] EventsQueryMsg exchange
         * @property {google.protobuf.ITimestamp|null} [dateStart] EventsQueryMsg dateStart
         * @property {google.protobuf.ITimestamp|null} [dateEnd] EventsQueryMsg dateEnd
         * @property {string|null} [pair] EventsQueryMsg pair
         * @property {string|null} [interval] EventsQueryMsg interval
         * @property {number|Long|null} [ticks] EventsQueryMsg ticks
         * @property {number|Long|null} [snapshotSize] EventsQueryMsg snapshotSize
         */

        /**
         * Constructs a new EventsQueryMsg.
         * @memberof orderbooks
         * @classdesc Represents an EventsQueryMsg.
         * @implements IEventsQueryMsg
         * @constructor
         * @param {orderbooks.IEventsQueryMsg=} [properties] Properties to set
         */
        function EventsQueryMsg(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * EventsQueryMsg exchange.
         * @member {string} exchange
         * @memberof orderbooks.EventsQueryMsg
         * @instance
         */
        EventsQueryMsg.prototype.exchange = "";

        /**
         * EventsQueryMsg dateStart.
         * @member {google.protobuf.ITimestamp|null|undefined} dateStart
         * @memberof orderbooks.EventsQueryMsg
         * @instance
         */
        EventsQueryMsg.prototype.dateStart = null;

        /**
         * EventsQueryMsg dateEnd.
         * @member {google.protobuf.ITimestamp|null|undefined} dateEnd
         * @memberof orderbooks.EventsQueryMsg
         * @instance
         */
        EventsQueryMsg.prototype.dateEnd = null;

        /**
         * EventsQueryMsg pair.
         * @member {string} pair
         * @memberof orderbooks.EventsQueryMsg
         * @instance
         */
        EventsQueryMsg.prototype.pair = "";

        /**
         * EventsQueryMsg interval.
         * @member {string} interval
         * @memberof orderbooks.EventsQueryMsg
         * @instance
         */
        EventsQueryMsg.prototype.interval = "";

        /**
         * EventsQueryMsg ticks.
         * @member {number|Long} ticks
         * @memberof orderbooks.EventsQueryMsg
         * @instance
         */
        EventsQueryMsg.prototype.ticks = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * EventsQueryMsg snapshotSize.
         * @member {number|Long} snapshotSize
         * @memberof orderbooks.EventsQueryMsg
         * @instance
         */
        EventsQueryMsg.prototype.snapshotSize = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        // OneOf field names bound to virtual getters and setters
        let $oneOfFields;

        /**
         * EventsQueryMsg snapshotInterval.
         * @member {"interval"|"ticks"|undefined} snapshotInterval
         * @memberof orderbooks.EventsQueryMsg
         * @instance
         */
        Object.defineProperty(EventsQueryMsg.prototype, "snapshotInterval", {
            get: $util.oneOfGetter($oneOfFields = ["interval", "ticks"]),
            set: $util.oneOfSetter($oneOfFields)
        });

        /**
         * Creates a new EventsQueryMsg instance using the specified properties.
         * @function create
         * @memberof orderbooks.EventsQueryMsg
         * @static
         * @param {orderbooks.IEventsQueryMsg=} [properties] Properties to set
         * @returns {orderbooks.EventsQueryMsg} EventsQueryMsg instance
         */
        EventsQueryMsg.create = function create(properties) {
            return new EventsQueryMsg(properties);
        };

        /**
         * Encodes the specified EventsQueryMsg message. Does not implicitly {@link orderbooks.EventsQueryMsg.verify|verify} messages.
         * @function encode
         * @memberof orderbooks.EventsQueryMsg
         * @static
         * @param {orderbooks.IEventsQueryMsg} message EventsQueryMsg message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        EventsQueryMsg.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.exchange != null && Object.hasOwnProperty.call(message, "exchange"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.exchange);
            if (message.dateStart != null && Object.hasOwnProperty.call(message, "dateStart"))
                $root.google.protobuf.Timestamp.encode(message.dateStart, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            if (message.dateEnd != null && Object.hasOwnProperty.call(message, "dateEnd"))
                $root.google.protobuf.Timestamp.encode(message.dateEnd, writer.uint32(/* id 3, wireType 2 =*/26).fork()).ldelim();
            if (message.pair != null && Object.hasOwnProperty.call(message, "pair"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.pair);
            if (message.interval != null && Object.hasOwnProperty.call(message, "interval"))
                writer.uint32(/* id 5, wireType 2 =*/42).string(message.interval);
            if (message.snapshotSize != null && Object.hasOwnProperty.call(message, "snapshotSize"))
                writer.uint32(/* id 6, wireType 0 =*/48).int64(message.snapshotSize);
            if (message.ticks != null && Object.hasOwnProperty.call(message, "ticks"))
                writer.uint32(/* id 7, wireType 0 =*/56).int64(message.ticks);
            return writer;
        };

        /**
         * Encodes the specified EventsQueryMsg message, length delimited. Does not implicitly {@link orderbooks.EventsQueryMsg.verify|verify} messages.
         * @function encodeDelimited
         * @memberof orderbooks.EventsQueryMsg
         * @static
         * @param {orderbooks.IEventsQueryMsg} message EventsQueryMsg message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        EventsQueryMsg.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an EventsQueryMsg message from the specified reader or buffer.
         * @function decode
         * @memberof orderbooks.EventsQueryMsg
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {orderbooks.EventsQueryMsg} EventsQueryMsg
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        EventsQueryMsg.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.orderbooks.EventsQueryMsg();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.exchange = reader.string();
                    break;
                case 2:
                    message.dateStart = $root.google.protobuf.Timestamp.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.dateEnd = $root.google.protobuf.Timestamp.decode(reader, reader.uint32());
                    break;
                case 4:
                    message.pair = reader.string();
                    break;
                case 5:
                    message.interval = reader.string();
                    break;
                case 7:
                    message.ticks = reader.int64();
                    break;
                case 6:
                    message.snapshotSize = reader.int64();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an EventsQueryMsg message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof orderbooks.EventsQueryMsg
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {orderbooks.EventsQueryMsg} EventsQueryMsg
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        EventsQueryMsg.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an EventsQueryMsg message.
         * @function verify
         * @memberof orderbooks.EventsQueryMsg
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        EventsQueryMsg.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            let properties = {};
            if (message.exchange != null && message.hasOwnProperty("exchange"))
                if (!$util.isString(message.exchange))
                    return "exchange: string expected";
            if (message.dateStart != null && message.hasOwnProperty("dateStart")) {
                let error = $root.google.protobuf.Timestamp.verify(message.dateStart);
                if (error)
                    return "dateStart." + error;
            }
            if (message.dateEnd != null && message.hasOwnProperty("dateEnd")) {
                let error = $root.google.protobuf.Timestamp.verify(message.dateEnd);
                if (error)
                    return "dateEnd." + error;
            }
            if (message.pair != null && message.hasOwnProperty("pair"))
                if (!$util.isString(message.pair))
                    return "pair: string expected";
            if (message.interval != null && message.hasOwnProperty("interval")) {
                properties.snapshotInterval = 1;
                if (!$util.isString(message.interval))
                    return "interval: string expected";
            }
            if (message.ticks != null && message.hasOwnProperty("ticks")) {
                if (properties.snapshotInterval === 1)
                    return "snapshotInterval: multiple values";
                properties.snapshotInterval = 1;
                if (!$util.isInteger(message.ticks) && !(message.ticks && $util.isInteger(message.ticks.low) && $util.isInteger(message.ticks.high)))
                    return "ticks: integer|Long expected";
            }
            if (message.snapshotSize != null && message.hasOwnProperty("snapshotSize"))
                if (!$util.isInteger(message.snapshotSize) && !(message.snapshotSize && $util.isInteger(message.snapshotSize.low) && $util.isInteger(message.snapshotSize.high)))
                    return "snapshotSize: integer|Long expected";
            return null;
        };

        /**
         * Creates an EventsQueryMsg message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof orderbooks.EventsQueryMsg
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {orderbooks.EventsQueryMsg} EventsQueryMsg
         */
        EventsQueryMsg.fromObject = function fromObject(object) {
            if (object instanceof $root.orderbooks.EventsQueryMsg)
                return object;
            let message = new $root.orderbooks.EventsQueryMsg();
            if (object.exchange != null)
                message.exchange = String(object.exchange);
            if (object.dateStart != null) {
                if (typeof object.dateStart !== "object")
                    throw TypeError(".orderbooks.EventsQueryMsg.dateStart: object expected");
                message.dateStart = $root.google.protobuf.Timestamp.fromObject(object.dateStart);
            }
            if (object.dateEnd != null) {
                if (typeof object.dateEnd !== "object")
                    throw TypeError(".orderbooks.EventsQueryMsg.dateEnd: object expected");
                message.dateEnd = $root.google.protobuf.Timestamp.fromObject(object.dateEnd);
            }
            if (object.pair != null)
                message.pair = String(object.pair);
            if (object.interval != null)
                message.interval = String(object.interval);
            if (object.ticks != null)
                if ($util.Long)
                    (message.ticks = $util.Long.fromValue(object.ticks)).unsigned = false;
                else if (typeof object.ticks === "string")
                    message.ticks = parseInt(object.ticks, 10);
                else if (typeof object.ticks === "number")
                    message.ticks = object.ticks;
                else if (typeof object.ticks === "object")
                    message.ticks = new $util.LongBits(object.ticks.low >>> 0, object.ticks.high >>> 0).toNumber();
            if (object.snapshotSize != null)
                if ($util.Long)
                    (message.snapshotSize = $util.Long.fromValue(object.snapshotSize)).unsigned = false;
                else if (typeof object.snapshotSize === "string")
                    message.snapshotSize = parseInt(object.snapshotSize, 10);
                else if (typeof object.snapshotSize === "number")
                    message.snapshotSize = object.snapshotSize;
                else if (typeof object.snapshotSize === "object")
                    message.snapshotSize = new $util.LongBits(object.snapshotSize.low >>> 0, object.snapshotSize.high >>> 0).toNumber();
            return message;
        };

        /**
         * Creates a plain object from an EventsQueryMsg message. Also converts values to other types if specified.
         * @function toObject
         * @memberof orderbooks.EventsQueryMsg
         * @static
         * @param {orderbooks.EventsQueryMsg} message EventsQueryMsg
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        EventsQueryMsg.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.exchange = "";
                object.dateStart = null;
                object.dateEnd = null;
                object.pair = "";
                if ($util.Long) {
                    let long = new $util.Long(0, 0, false);
                    object.snapshotSize = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.snapshotSize = options.longs === String ? "0" : 0;
            }
            if (message.exchange != null && message.hasOwnProperty("exchange"))
                object.exchange = message.exchange;
            if (message.dateStart != null && message.hasOwnProperty("dateStart"))
                object.dateStart = $root.google.protobuf.Timestamp.toObject(message.dateStart, options);
            if (message.dateEnd != null && message.hasOwnProperty("dateEnd"))
                object.dateEnd = $root.google.protobuf.Timestamp.toObject(message.dateEnd, options);
            if (message.pair != null && message.hasOwnProperty("pair"))
                object.pair = message.pair;
            if (message.interval != null && message.hasOwnProperty("interval")) {
                object.interval = message.interval;
                if (options.oneofs)
                    object.snapshotInterval = "interval";
            }
            if (message.snapshotSize != null && message.hasOwnProperty("snapshotSize"))
                if (typeof message.snapshotSize === "number")
                    object.snapshotSize = options.longs === String ? String(message.snapshotSize) : message.snapshotSize;
                else
                    object.snapshotSize = options.longs === String ? $util.Long.prototype.toString.call(message.snapshotSize) : options.longs === Number ? new $util.LongBits(message.snapshotSize.low >>> 0, message.snapshotSize.high >>> 0).toNumber() : message.snapshotSize;
            if (message.ticks != null && message.hasOwnProperty("ticks")) {
                if (typeof message.ticks === "number")
                    object.ticks = options.longs === String ? String(message.ticks) : message.ticks;
                else
                    object.ticks = options.longs === String ? $util.Long.prototype.toString.call(message.ticks) : options.longs === Number ? new $util.LongBits(message.ticks.low >>> 0, message.ticks.high >>> 0).toNumber() : message.ticks;
                if (options.oneofs)
                    object.snapshotInterval = "ticks";
            }
            return object;
        };

        /**
         * Converts this EventsQueryMsg to JSON.
         * @function toJSON
         * @memberof orderbooks.EventsQueryMsg
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        EventsQueryMsg.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return EventsQueryMsg;
    })();

    /**
     * SnapshotMode enum.
     * @name orderbooks.SnapshotMode
     * @enum {number}
     * @property {number} TIME=0 TIME value
     * @property {number} TICKS=1 TICKS value
     */
    orderbooks.SnapshotMode = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "TIME"] = 0;
        values[valuesById[1] = "TICKS"] = 1;
        return values;
    })();

    /**
     * OrderSide enum.
     * @name orderbooks.OrderSide
     * @enum {number}
     * @property {number} BID=0 BID value
     * @property {number} ASK=1 ASK value
     */
    orderbooks.OrderSide = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "BID"] = 0;
        values[valuesById[1] = "ASK"] = 1;
        return values;
    })();

    /**
     * EventType enum.
     * @name orderbooks.EventType
     * @enum {number}
     * @property {number} INIT=0 INIT value
     * @property {number} ADD=1 ADD value
     * @property {number} CHANGE=2 CHANGE value
     * @property {number} REMOVE=3 REMOVE value
     */
    orderbooks.EventType = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "INIT"] = 0;
        values[valuesById[1] = "ADD"] = 1;
        values[valuesById[2] = "CHANGE"] = 2;
        values[valuesById[3] = "REMOVE"] = 3;
        return values;
    })();

    orderbooks.SnapshotMsg = (function() {

        /**
         * Properties of a SnapshotMsg.
         * @memberof orderbooks
         * @interface ISnapshotMsg
         * @property {google.protobuf.ITimestamp|null} [timestamp] SnapshotMsg timestamp
         * @property {string|null} [exchange] SnapshotMsg exchange
         * @property {string|null} [pair] SnapshotMsg pair
         * @property {Array.<orderbooks.IEvent>|null} [orders] SnapshotMsg orders
         * @property {Array.<orderbooks.IEvent>|null} [events] SnapshotMsg events
         * @property {number|Long|null} [sessionId] SnapshotMsg sessionId
         * @property {number|Long|null} [counter] SnapshotMsg counter
         */

        /**
         * Constructs a new SnapshotMsg.
         * @memberof orderbooks
         * @classdesc Represents a SnapshotMsg.
         * @implements ISnapshotMsg
         * @constructor
         * @param {orderbooks.ISnapshotMsg=} [properties] Properties to set
         */
        function SnapshotMsg(properties) {
            this.orders = [];
            this.events = [];
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * SnapshotMsg timestamp.
         * @member {google.protobuf.ITimestamp|null|undefined} timestamp
         * @memberof orderbooks.SnapshotMsg
         * @instance
         */
        SnapshotMsg.prototype.timestamp = null;

        /**
         * SnapshotMsg exchange.
         * @member {string} exchange
         * @memberof orderbooks.SnapshotMsg
         * @instance
         */
        SnapshotMsg.prototype.exchange = "";

        /**
         * SnapshotMsg pair.
         * @member {string} pair
         * @memberof orderbooks.SnapshotMsg
         * @instance
         */
        SnapshotMsg.prototype.pair = "";

        /**
         * SnapshotMsg orders.
         * @member {Array.<orderbooks.IEvent>} orders
         * @memberof orderbooks.SnapshotMsg
         * @instance
         */
        SnapshotMsg.prototype.orders = $util.emptyArray;

        /**
         * SnapshotMsg events.
         * @member {Array.<orderbooks.IEvent>} events
         * @memberof orderbooks.SnapshotMsg
         * @instance
         */
        SnapshotMsg.prototype.events = $util.emptyArray;

        /**
         * SnapshotMsg sessionId.
         * @member {number|Long} sessionId
         * @memberof orderbooks.SnapshotMsg
         * @instance
         */
        SnapshotMsg.prototype.sessionId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * SnapshotMsg counter.
         * @member {number|Long} counter
         * @memberof orderbooks.SnapshotMsg
         * @instance
         */
        SnapshotMsg.prototype.counter = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Creates a new SnapshotMsg instance using the specified properties.
         * @function create
         * @memberof orderbooks.SnapshotMsg
         * @static
         * @param {orderbooks.ISnapshotMsg=} [properties] Properties to set
         * @returns {orderbooks.SnapshotMsg} SnapshotMsg instance
         */
        SnapshotMsg.create = function create(properties) {
            return new SnapshotMsg(properties);
        };

        /**
         * Encodes the specified SnapshotMsg message. Does not implicitly {@link orderbooks.SnapshotMsg.verify|verify} messages.
         * @function encode
         * @memberof orderbooks.SnapshotMsg
         * @static
         * @param {orderbooks.ISnapshotMsg} message SnapshotMsg message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        SnapshotMsg.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.timestamp != null && Object.hasOwnProperty.call(message, "timestamp"))
                $root.google.protobuf.Timestamp.encode(message.timestamp, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            if (message.exchange != null && Object.hasOwnProperty.call(message, "exchange"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.exchange);
            if (message.pair != null && Object.hasOwnProperty.call(message, "pair"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.pair);
            if (message.orders != null && message.orders.length)
                for (let i = 0; i < message.orders.length; ++i)
                    $root.orderbooks.Event.encode(message.orders[i], writer.uint32(/* id 4, wireType 2 =*/34).fork()).ldelim();
            if (message.events != null && message.events.length)
                for (let i = 0; i < message.events.length; ++i)
                    $root.orderbooks.Event.encode(message.events[i], writer.uint32(/* id 5, wireType 2 =*/42).fork()).ldelim();
            if (message.sessionId != null && Object.hasOwnProperty.call(message, "sessionId"))
                writer.uint32(/* id 6, wireType 0 =*/48).int64(message.sessionId);
            if (message.counter != null && Object.hasOwnProperty.call(message, "counter"))
                writer.uint32(/* id 7, wireType 0 =*/56).int64(message.counter);
            return writer;
        };

        /**
         * Encodes the specified SnapshotMsg message, length delimited. Does not implicitly {@link orderbooks.SnapshotMsg.verify|verify} messages.
         * @function encodeDelimited
         * @memberof orderbooks.SnapshotMsg
         * @static
         * @param {orderbooks.ISnapshotMsg} message SnapshotMsg message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        SnapshotMsg.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a SnapshotMsg message from the specified reader or buffer.
         * @function decode
         * @memberof orderbooks.SnapshotMsg
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {orderbooks.SnapshotMsg} SnapshotMsg
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        SnapshotMsg.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.orderbooks.SnapshotMsg();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.timestamp = $root.google.protobuf.Timestamp.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.exchange = reader.string();
                    break;
                case 3:
                    message.pair = reader.string();
                    break;
                case 4:
                    if (!(message.orders && message.orders.length))
                        message.orders = [];
                    message.orders.push($root.orderbooks.Event.decode(reader, reader.uint32()));
                    break;
                case 5:
                    if (!(message.events && message.events.length))
                        message.events = [];
                    message.events.push($root.orderbooks.Event.decode(reader, reader.uint32()));
                    break;
                case 6:
                    message.sessionId = reader.int64();
                    break;
                case 7:
                    message.counter = reader.int64();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a SnapshotMsg message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof orderbooks.SnapshotMsg
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {orderbooks.SnapshotMsg} SnapshotMsg
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        SnapshotMsg.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a SnapshotMsg message.
         * @function verify
         * @memberof orderbooks.SnapshotMsg
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        SnapshotMsg.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.timestamp != null && message.hasOwnProperty("timestamp")) {
                let error = $root.google.protobuf.Timestamp.verify(message.timestamp);
                if (error)
                    return "timestamp." + error;
            }
            if (message.exchange != null && message.hasOwnProperty("exchange"))
                if (!$util.isString(message.exchange))
                    return "exchange: string expected";
            if (message.pair != null && message.hasOwnProperty("pair"))
                if (!$util.isString(message.pair))
                    return "pair: string expected";
            if (message.orders != null && message.hasOwnProperty("orders")) {
                if (!Array.isArray(message.orders))
                    return "orders: array expected";
                for (let i = 0; i < message.orders.length; ++i) {
                    let error = $root.orderbooks.Event.verify(message.orders[i]);
                    if (error)
                        return "orders." + error;
                }
            }
            if (message.events != null && message.hasOwnProperty("events")) {
                if (!Array.isArray(message.events))
                    return "events: array expected";
                for (let i = 0; i < message.events.length; ++i) {
                    let error = $root.orderbooks.Event.verify(message.events[i]);
                    if (error)
                        return "events." + error;
                }
            }
            if (message.sessionId != null && message.hasOwnProperty("sessionId"))
                if (!$util.isInteger(message.sessionId) && !(message.sessionId && $util.isInteger(message.sessionId.low) && $util.isInteger(message.sessionId.high)))
                    return "sessionId: integer|Long expected";
            if (message.counter != null && message.hasOwnProperty("counter"))
                if (!$util.isInteger(message.counter) && !(message.counter && $util.isInteger(message.counter.low) && $util.isInteger(message.counter.high)))
                    return "counter: integer|Long expected";
            return null;
        };

        /**
         * Creates a SnapshotMsg message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof orderbooks.SnapshotMsg
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {orderbooks.SnapshotMsg} SnapshotMsg
         */
        SnapshotMsg.fromObject = function fromObject(object) {
            if (object instanceof $root.orderbooks.SnapshotMsg)
                return object;
            let message = new $root.orderbooks.SnapshotMsg();
            if (object.timestamp != null) {
                if (typeof object.timestamp !== "object")
                    throw TypeError(".orderbooks.SnapshotMsg.timestamp: object expected");
                message.timestamp = $root.google.protobuf.Timestamp.fromObject(object.timestamp);
            }
            if (object.exchange != null)
                message.exchange = String(object.exchange);
            if (object.pair != null)
                message.pair = String(object.pair);
            if (object.orders) {
                if (!Array.isArray(object.orders))
                    throw TypeError(".orderbooks.SnapshotMsg.orders: array expected");
                message.orders = [];
                for (let i = 0; i < object.orders.length; ++i) {
                    if (typeof object.orders[i] !== "object")
                        throw TypeError(".orderbooks.SnapshotMsg.orders: object expected");
                    message.orders[i] = $root.orderbooks.Event.fromObject(object.orders[i]);
                }
            }
            if (object.events) {
                if (!Array.isArray(object.events))
                    throw TypeError(".orderbooks.SnapshotMsg.events: array expected");
                message.events = [];
                for (let i = 0; i < object.events.length; ++i) {
                    if (typeof object.events[i] !== "object")
                        throw TypeError(".orderbooks.SnapshotMsg.events: object expected");
                    message.events[i] = $root.orderbooks.Event.fromObject(object.events[i]);
                }
            }
            if (object.sessionId != null)
                if ($util.Long)
                    (message.sessionId = $util.Long.fromValue(object.sessionId)).unsigned = false;
                else if (typeof object.sessionId === "string")
                    message.sessionId = parseInt(object.sessionId, 10);
                else if (typeof object.sessionId === "number")
                    message.sessionId = object.sessionId;
                else if (typeof object.sessionId === "object")
                    message.sessionId = new $util.LongBits(object.sessionId.low >>> 0, object.sessionId.high >>> 0).toNumber();
            if (object.counter != null)
                if ($util.Long)
                    (message.counter = $util.Long.fromValue(object.counter)).unsigned = false;
                else if (typeof object.counter === "string")
                    message.counter = parseInt(object.counter, 10);
                else if (typeof object.counter === "number")
                    message.counter = object.counter;
                else if (typeof object.counter === "object")
                    message.counter = new $util.LongBits(object.counter.low >>> 0, object.counter.high >>> 0).toNumber();
            return message;
        };

        /**
         * Creates a plain object from a SnapshotMsg message. Also converts values to other types if specified.
         * @function toObject
         * @memberof orderbooks.SnapshotMsg
         * @static
         * @param {orderbooks.SnapshotMsg} message SnapshotMsg
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        SnapshotMsg.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.arrays || options.defaults) {
                object.orders = [];
                object.events = [];
            }
            if (options.defaults) {
                object.timestamp = null;
                object.exchange = "";
                object.pair = "";
                if ($util.Long) {
                    let long = new $util.Long(0, 0, false);
                    object.sessionId = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.sessionId = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    let long = new $util.Long(0, 0, false);
                    object.counter = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.counter = options.longs === String ? "0" : 0;
            }
            if (message.timestamp != null && message.hasOwnProperty("timestamp"))
                object.timestamp = $root.google.protobuf.Timestamp.toObject(message.timestamp, options);
            if (message.exchange != null && message.hasOwnProperty("exchange"))
                object.exchange = message.exchange;
            if (message.pair != null && message.hasOwnProperty("pair"))
                object.pair = message.pair;
            if (message.orders && message.orders.length) {
                object.orders = [];
                for (let j = 0; j < message.orders.length; ++j)
                    object.orders[j] = $root.orderbooks.Event.toObject(message.orders[j], options);
            }
            if (message.events && message.events.length) {
                object.events = [];
                for (let j = 0; j < message.events.length; ++j)
                    object.events[j] = $root.orderbooks.Event.toObject(message.events[j], options);
            }
            if (message.sessionId != null && message.hasOwnProperty("sessionId"))
                if (typeof message.sessionId === "number")
                    object.sessionId = options.longs === String ? String(message.sessionId) : message.sessionId;
                else
                    object.sessionId = options.longs === String ? $util.Long.prototype.toString.call(message.sessionId) : options.longs === Number ? new $util.LongBits(message.sessionId.low >>> 0, message.sessionId.high >>> 0).toNumber() : message.sessionId;
            if (message.counter != null && message.hasOwnProperty("counter"))
                if (typeof message.counter === "number")
                    object.counter = options.longs === String ? String(message.counter) : message.counter;
                else
                    object.counter = options.longs === String ? $util.Long.prototype.toString.call(message.counter) : options.longs === Number ? new $util.LongBits(message.counter.low >>> 0, message.counter.high >>> 0).toNumber() : message.counter;
            return object;
        };

        /**
         * Converts this SnapshotMsg to JSON.
         * @function toJSON
         * @memberof orderbooks.SnapshotMsg
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        SnapshotMsg.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return SnapshotMsg;
    })();

    orderbooks.Event = (function() {

        /**
         * Properties of an Event.
         * @memberof orderbooks
         * @interface IEvent
         * @property {number|Long|null} [timestamp] Event timestamp
         * @property {orderbooks.OrderSide|null} [orderSide] Event orderSide
         * @property {number|null} [amount] Event amount
         * @property {number|null} [price] Event price
         * @property {orderbooks.EventType|null} [eventType] Event eventType
         */

        /**
         * Constructs a new Event.
         * @memberof orderbooks
         * @classdesc Represents an Event.
         * @implements IEvent
         * @constructor
         * @param {orderbooks.IEvent=} [properties] Properties to set
         */
        function Event(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Event timestamp.
         * @member {number|Long} timestamp
         * @memberof orderbooks.Event
         * @instance
         */
        Event.prototype.timestamp = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Event orderSide.
         * @member {orderbooks.OrderSide} orderSide
         * @memberof orderbooks.Event
         * @instance
         */
        Event.prototype.orderSide = 0;

        /**
         * Event amount.
         * @member {number} amount
         * @memberof orderbooks.Event
         * @instance
         */
        Event.prototype.amount = 0;

        /**
         * Event price.
         * @member {number} price
         * @memberof orderbooks.Event
         * @instance
         */
        Event.prototype.price = 0;

        /**
         * Event eventType.
         * @member {orderbooks.EventType} eventType
         * @memberof orderbooks.Event
         * @instance
         */
        Event.prototype.eventType = 0;

        /**
         * Creates a new Event instance using the specified properties.
         * @function create
         * @memberof orderbooks.Event
         * @static
         * @param {orderbooks.IEvent=} [properties] Properties to set
         * @returns {orderbooks.Event} Event instance
         */
        Event.create = function create(properties) {
            return new Event(properties);
        };

        /**
         * Encodes the specified Event message. Does not implicitly {@link orderbooks.Event.verify|verify} messages.
         * @function encode
         * @memberof orderbooks.Event
         * @static
         * @param {orderbooks.IEvent} message Event message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Event.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.timestamp != null && Object.hasOwnProperty.call(message, "timestamp"))
                writer.uint32(/* id 1, wireType 0 =*/8).int64(message.timestamp);
            if (message.orderSide != null && Object.hasOwnProperty.call(message, "orderSide"))
                writer.uint32(/* id 2, wireType 0 =*/16).int32(message.orderSide);
            if (message.amount != null && Object.hasOwnProperty.call(message, "amount"))
                writer.uint32(/* id 3, wireType 5 =*/29).float(message.amount);
            if (message.price != null && Object.hasOwnProperty.call(message, "price"))
                writer.uint32(/* id 4, wireType 5 =*/37).float(message.price);
            if (message.eventType != null && Object.hasOwnProperty.call(message, "eventType"))
                writer.uint32(/* id 5, wireType 0 =*/40).int32(message.eventType);
            return writer;
        };

        /**
         * Encodes the specified Event message, length delimited. Does not implicitly {@link orderbooks.Event.verify|verify} messages.
         * @function encodeDelimited
         * @memberof orderbooks.Event
         * @static
         * @param {orderbooks.IEvent} message Event message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Event.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an Event message from the specified reader or buffer.
         * @function decode
         * @memberof orderbooks.Event
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {orderbooks.Event} Event
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Event.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.orderbooks.Event();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.timestamp = reader.int64();
                    break;
                case 2:
                    message.orderSide = reader.int32();
                    break;
                case 3:
                    message.amount = reader.float();
                    break;
                case 4:
                    message.price = reader.float();
                    break;
                case 5:
                    message.eventType = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an Event message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof orderbooks.Event
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {orderbooks.Event} Event
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Event.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an Event message.
         * @function verify
         * @memberof orderbooks.Event
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Event.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.timestamp != null && message.hasOwnProperty("timestamp"))
                if (!$util.isInteger(message.timestamp) && !(message.timestamp && $util.isInteger(message.timestamp.low) && $util.isInteger(message.timestamp.high)))
                    return "timestamp: integer|Long expected";
            if (message.orderSide != null && message.hasOwnProperty("orderSide"))
                switch (message.orderSide) {
                default:
                    return "orderSide: enum value expected";
                case 0:
                case 1:
                    break;
                }
            if (message.amount != null && message.hasOwnProperty("amount"))
                if (typeof message.amount !== "number")
                    return "amount: number expected";
            if (message.price != null && message.hasOwnProperty("price"))
                if (typeof message.price !== "number")
                    return "price: number expected";
            if (message.eventType != null && message.hasOwnProperty("eventType"))
                switch (message.eventType) {
                default:
                    return "eventType: enum value expected";
                case 0:
                case 1:
                case 2:
                case 3:
                    break;
                }
            return null;
        };

        /**
         * Creates an Event message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof orderbooks.Event
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {orderbooks.Event} Event
         */
        Event.fromObject = function fromObject(object) {
            if (object instanceof $root.orderbooks.Event)
                return object;
            let message = new $root.orderbooks.Event();
            if (object.timestamp != null)
                if ($util.Long)
                    (message.timestamp = $util.Long.fromValue(object.timestamp)).unsigned = false;
                else if (typeof object.timestamp === "string")
                    message.timestamp = parseInt(object.timestamp, 10);
                else if (typeof object.timestamp === "number")
                    message.timestamp = object.timestamp;
                else if (typeof object.timestamp === "object")
                    message.timestamp = new $util.LongBits(object.timestamp.low >>> 0, object.timestamp.high >>> 0).toNumber();
            switch (object.orderSide) {
            case "BID":
            case 0:
                message.orderSide = 0;
                break;
            case "ASK":
            case 1:
                message.orderSide = 1;
                break;
            }
            if (object.amount != null)
                message.amount = Number(object.amount);
            if (object.price != null)
                message.price = Number(object.price);
            switch (object.eventType) {
            case "INIT":
            case 0:
                message.eventType = 0;
                break;
            case "ADD":
            case 1:
                message.eventType = 1;
                break;
            case "CHANGE":
            case 2:
                message.eventType = 2;
                break;
            case "REMOVE":
            case 3:
                message.eventType = 3;
                break;
            }
            return message;
        };

        /**
         * Creates a plain object from an Event message. Also converts values to other types if specified.
         * @function toObject
         * @memberof orderbooks.Event
         * @static
         * @param {orderbooks.Event} message Event
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Event.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                if ($util.Long) {
                    let long = new $util.Long(0, 0, false);
                    object.timestamp = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.timestamp = options.longs === String ? "0" : 0;
                object.orderSide = options.enums === String ? "BID" : 0;
                object.amount = 0;
                object.price = 0;
                object.eventType = options.enums === String ? "INIT" : 0;
            }
            if (message.timestamp != null && message.hasOwnProperty("timestamp"))
                if (typeof message.timestamp === "number")
                    object.timestamp = options.longs === String ? String(message.timestamp) : message.timestamp;
                else
                    object.timestamp = options.longs === String ? $util.Long.prototype.toString.call(message.timestamp) : options.longs === Number ? new $util.LongBits(message.timestamp.low >>> 0, message.timestamp.high >>> 0).toNumber() : message.timestamp;
            if (message.orderSide != null && message.hasOwnProperty("orderSide"))
                object.orderSide = options.enums === String ? $root.orderbooks.OrderSide[message.orderSide] : message.orderSide;
            if (message.amount != null && message.hasOwnProperty("amount"))
                object.amount = options.json && !isFinite(message.amount) ? String(message.amount) : message.amount;
            if (message.price != null && message.hasOwnProperty("price"))
                object.price = options.json && !isFinite(message.price) ? String(message.price) : message.price;
            if (message.eventType != null && message.hasOwnProperty("eventType"))
                object.eventType = options.enums === String ? $root.orderbooks.EventType[message.eventType] : message.eventType;
            return object;
        };

        /**
         * Converts this Event to JSON.
         * @function toJSON
         * @memberof orderbooks.Event
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Event.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Event;
    })();

    return orderbooks;
})();

export const google = $root.google = (() => {

    /**
     * Namespace google.
     * @exports google
     * @namespace
     */
    const google = {};

    google.protobuf = (function() {

        /**
         * Namespace protobuf.
         * @memberof google
         * @namespace
         */
        const protobuf = {};

        protobuf.Timestamp = (function() {

            /**
             * Properties of a Timestamp.
             * @memberof google.protobuf
             * @interface ITimestamp
             * @property {number|Long|null} [seconds] Timestamp seconds
             * @property {number|null} [nanos] Timestamp nanos
             */

            /**
             * Constructs a new Timestamp.
             * @memberof google.protobuf
             * @classdesc Represents a Timestamp.
             * @implements ITimestamp
             * @constructor
             * @param {google.protobuf.ITimestamp=} [properties] Properties to set
             */
            function Timestamp(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * Timestamp seconds.
             * @member {number|Long} seconds
             * @memberof google.protobuf.Timestamp
             * @instance
             */
            Timestamp.prototype.seconds = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

            /**
             * Timestamp nanos.
             * @member {number} nanos
             * @memberof google.protobuf.Timestamp
             * @instance
             */
            Timestamp.prototype.nanos = 0;

            /**
             * Creates a new Timestamp instance using the specified properties.
             * @function create
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {google.protobuf.ITimestamp=} [properties] Properties to set
             * @returns {google.protobuf.Timestamp} Timestamp instance
             */
            Timestamp.create = function create(properties) {
                return new Timestamp(properties);
            };

            /**
             * Encodes the specified Timestamp message. Does not implicitly {@link google.protobuf.Timestamp.verify|verify} messages.
             * @function encode
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {google.protobuf.ITimestamp} message Timestamp message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Timestamp.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.seconds != null && Object.hasOwnProperty.call(message, "seconds"))
                    writer.uint32(/* id 1, wireType 0 =*/8).int64(message.seconds);
                if (message.nanos != null && Object.hasOwnProperty.call(message, "nanos"))
                    writer.uint32(/* id 2, wireType 0 =*/16).int32(message.nanos);
                return writer;
            };

            /**
             * Encodes the specified Timestamp message, length delimited. Does not implicitly {@link google.protobuf.Timestamp.verify|verify} messages.
             * @function encodeDelimited
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {google.protobuf.ITimestamp} message Timestamp message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Timestamp.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a Timestamp message from the specified reader or buffer.
             * @function decode
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {google.protobuf.Timestamp} Timestamp
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Timestamp.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.google.protobuf.Timestamp();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.seconds = reader.int64();
                        break;
                    case 2:
                        message.nanos = reader.int32();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a Timestamp message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {google.protobuf.Timestamp} Timestamp
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Timestamp.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a Timestamp message.
             * @function verify
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            Timestamp.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.seconds != null && message.hasOwnProperty("seconds"))
                    if (!$util.isInteger(message.seconds) && !(message.seconds && $util.isInteger(message.seconds.low) && $util.isInteger(message.seconds.high)))
                        return "seconds: integer|Long expected";
                if (message.nanos != null && message.hasOwnProperty("nanos"))
                    if (!$util.isInteger(message.nanos))
                        return "nanos: integer expected";
                return null;
            };

            /**
             * Creates a Timestamp message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {google.protobuf.Timestamp} Timestamp
             */
            Timestamp.fromObject = function fromObject(object) {
                if (object instanceof $root.google.protobuf.Timestamp)
                    return object;
                let message = new $root.google.protobuf.Timestamp();
                if (object.seconds != null)
                    if ($util.Long)
                        (message.seconds = $util.Long.fromValue(object.seconds)).unsigned = false;
                    else if (typeof object.seconds === "string")
                        message.seconds = parseInt(object.seconds, 10);
                    else if (typeof object.seconds === "number")
                        message.seconds = object.seconds;
                    else if (typeof object.seconds === "object")
                        message.seconds = new $util.LongBits(object.seconds.low >>> 0, object.seconds.high >>> 0).toNumber();
                if (object.nanos != null)
                    message.nanos = object.nanos | 0;
                return message;
            };

            /**
             * Creates a plain object from a Timestamp message. Also converts values to other types if specified.
             * @function toObject
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {google.protobuf.Timestamp} message Timestamp
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            Timestamp.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults) {
                    if ($util.Long) {
                        let long = new $util.Long(0, 0, false);
                        object.seconds = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                    } else
                        object.seconds = options.longs === String ? "0" : 0;
                    object.nanos = 0;
                }
                if (message.seconds != null && message.hasOwnProperty("seconds"))
                    if (typeof message.seconds === "number")
                        object.seconds = options.longs === String ? String(message.seconds) : message.seconds;
                    else
                        object.seconds = options.longs === String ? $util.Long.prototype.toString.call(message.seconds) : options.longs === Number ? new $util.LongBits(message.seconds.low >>> 0, message.seconds.high >>> 0).toNumber() : message.seconds;
                if (message.nanos != null && message.hasOwnProperty("nanos"))
                    object.nanos = message.nanos;
                return object;
            };

            /**
             * Converts this Timestamp to JSON.
             * @function toJSON
             * @memberof google.protobuf.Timestamp
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            Timestamp.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return Timestamp;
        })();

        return protobuf;
    })();

    return google;
})();

export { $root as default };
