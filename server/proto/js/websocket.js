
// orderbooks.proto

// Code generated by protoc-gen-gotemplate. DO NOT EDIT.
// Browser mode: false


import * as orderbooks from './orderbooks_pb'
import {WebSocketClient} from '@ticketag/wsrpc-client/src/ws';




//orderbooks.orderbooks.EmptyMsg
/**
* @typedef {Object} MessageEmptyMsg
* @property {orderbooks.EmptyMsg} body this is description.
* @property {number} id ID messaggio.
* @property {string} methodName Nome metodo.
* @property {MsgFlags} flags Flag messaggio.
*/

//orderbooks.orderbooks.GetExchangesResponseMsg
/**
* @typedef {Object} MessageGetExchangesResponseMsg
* @property {orderbooks.GetExchangesResponseMsg} body this is description.
* @property {number} id ID messaggio.
* @property {string} methodName Nome metodo.
* @property {MsgFlags} flags Flag messaggio.
*/

//orderbooks.orderbooks.ExchangeMsg
/**
* @typedef {Object} MessageExchangeMsg
* @property {orderbooks.ExchangeMsg} body this is description.
* @property {number} id ID messaggio.
* @property {string} methodName Nome metodo.
* @property {MsgFlags} flags Flag messaggio.
*/

//orderbooks.orderbooks.PairMsg
/**
* @typedef {Object} MessagePairMsg
* @property {orderbooks.PairMsg} body this is description.
* @property {number} id ID messaggio.
* @property {string} methodName Nome metodo.
* @property {MsgFlags} flags Flag messaggio.
*/

//orderbooks.orderbooks.EventsQueryMsg
/**
* @typedef {Object} MessageEventsQueryMsg
* @property {orderbooks.EventsQueryMsg} body this is description.
* @property {number} id ID messaggio.
* @property {string} methodName Nome metodo.
* @property {MsgFlags} flags Flag messaggio.
*/

//orderbooks.orderbooks.SnapshotMsg
/**
* @typedef {Object} MessageSnapshotMsg
* @property {orderbooks.SnapshotMsg} body this is description.
* @property {number} id ID messaggio.
* @property {string} methodName Nome metodo.
* @property {MsgFlags} flags Flag messaggio.
*/

//orderbooks.orderbooks.Event
/**
* @typedef {Object} MessageEvent
* @property {orderbooks.Event} body this is description.
* @property {number} id ID messaggio.
* @property {string} methodName Nome metodo.
* @property {MsgFlags} flags Flag messaggio.
*/





export class OrderBooksClient {
        constructor(protocol, url, options={}, client) {
                this._client = client || new WebSocketClient(protocol, url, options)
                this._protocol = this._client.protocol;
                this._url = this._client.url;
        
                //GetExchanges (OrderBooks.GetExchanges)
                this._client.registerMessageDecoder("OrderBooks.GetExchanges",orderbooks.orderbooks.GetExchangesResponseMsg); //.orderbooks.GetExchangesResponseMsg
                this._client.registerMessageEncoder("OrderBooks.GetExchanges",orderbooks.orderbooks.EmptyMsg); //name:"EmptyMsg" 
                
                
                //QueryEvents (OrderBooks.QueryEvents)
                this._client.registerMessageDecoder("OrderBooks.QueryEvents",orderbooks.orderbooks.SnapshotMsg); //.orderbooks.SnapshotMsg
                this._client.registerMessageEncoder("OrderBooks.QueryEvents",orderbooks.orderbooks.EventsQueryMsg); //name:"EventsQueryMsg" field:<name:"exchange" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"exchange" > field:<name:"date_start" number:2 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"dateStart" > field:<name:"date_end" number:3 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"dateEnd" > field:<name:"pair" number:4 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"pair" > field:<name:"interval" number:5 label:LABEL_OPTIONAL type:TYPE_STRING oneof_index:0 json_name:"interval" > field:<name:"ticks" number:7 label:LABEL_OPTIONAL type:TYPE_INT64 oneof_index:0 json_name:"ticks" > field:<name:"snapshot_size" number:6 label:LABEL_OPTIONAL type:TYPE_INT64 json_name:"snapshotSize" > oneof_decl:<name:"snapshot_interval" > 
                
                this._client.registerServerStream("OrderBooks.QueryEvents");
                //GetLiveData (OrderBooks.GetLiveData)
                this._client.registerMessageDecoder("OrderBooks.GetLiveData",orderbooks.orderbooks.SnapshotMsg); //.orderbooks.SnapshotMsg
                this._client.registerMessageEncoder("OrderBooks.GetLiveData",orderbooks.orderbooks.EventsQueryMsg); //name:"EventsQueryMsg" field:<name:"exchange" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"exchange" > field:<name:"date_start" number:2 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"dateStart" > field:<name:"date_end" number:3 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"dateEnd" > field:<name:"pair" number:4 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"pair" > field:<name:"interval" number:5 label:LABEL_OPTIONAL type:TYPE_STRING oneof_index:0 json_name:"interval" > field:<name:"ticks" number:7 label:LABEL_OPTIONAL type:TYPE_INT64 oneof_index:0 json_name:"ticks" > field:<name:"snapshot_size" number:6 label:LABEL_OPTIONAL type:TYPE_INT64 json_name:"snapshotSize" > oneof_decl:<name:"snapshot_interval" > 
                
                this._client.registerServerStream("OrderBooks.GetLiveData");
                //GetCachedData (OrderBooks.GetCachedData)
                this._client.registerMessageDecoder("OrderBooks.GetCachedData",orderbooks.orderbooks.SnapshotMsg); //.orderbooks.SnapshotMsg
                this._client.registerMessageEncoder("OrderBooks.GetCachedData",orderbooks.orderbooks.EventsQueryMsg); //name:"EventsQueryMsg" field:<name:"exchange" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"exchange" > field:<name:"date_start" number:2 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"dateStart" > field:<name:"date_end" number:3 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"dateEnd" > field:<name:"pair" number:4 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"pair" > field:<name:"interval" number:5 label:LABEL_OPTIONAL type:TYPE_STRING oneof_index:0 json_name:"interval" > field:<name:"ticks" number:7 label:LABEL_OPTIONAL type:TYPE_INT64 oneof_index:0 json_name:"ticks" > field:<name:"snapshot_size" number:6 label:LABEL_OPTIONAL type:TYPE_INT64 json_name:"snapshotSize" > oneof_decl:<name:"snapshot_interval" > 
                
                this._client.registerServerStream("OrderBooks.GetCachedData");

        }
        
        /**
        * Constructs a new EmptyMsg.
        * @memberof OrderBooksClient
        * @classdesc Represents a EmptyMsg.
        * @implements orderbooks.IEmptyMsg
        * @constructor
        * @param {orderbooks.IEmptyMsg=} [properties] Properties to set
        */
        NewEmptyMsg(properties){
                return orderbooks.orderbooks.EmptyMsg(properties)
        }
        
        /**
        * Constructs a new GetExchangesResponseMsg.
        * @memberof OrderBooksClient
        * @classdesc Represents a GetExchangesResponseMsg.
        * @implements orderbooks.IGetExchangesResponseMsg
        * @constructor
        * @param {orderbooks.IGetExchangesResponseMsg=} [properties] Properties to set
        */
        NewGetExchangesResponseMsg(properties){
                return orderbooks.orderbooks.GetExchangesResponseMsg(properties)
        }
        
        /**
        * Constructs a new ExchangeMsg.
        * @memberof OrderBooksClient
        * @classdesc Represents a ExchangeMsg.
        * @implements orderbooks.IExchangeMsg
        * @constructor
        * @param {orderbooks.IExchangeMsg=} [properties] Properties to set
        */
        NewExchangeMsg(properties){
                return orderbooks.orderbooks.ExchangeMsg(properties)
        }
        
        /**
        * Constructs a new PairMsg.
        * @memberof OrderBooksClient
        * @classdesc Represents a PairMsg.
        * @implements orderbooks.IPairMsg
        * @constructor
        * @param {orderbooks.IPairMsg=} [properties] Properties to set
        */
        NewPairMsg(properties){
                return orderbooks.orderbooks.PairMsg(properties)
        }
        
        /**
        * Constructs a new EventsQueryMsg.
        * @memberof OrderBooksClient
        * @classdesc Represents a EventsQueryMsg.
        * @implements orderbooks.IEventsQueryMsg
        * @constructor
        * @param {orderbooks.IEventsQueryMsg=} [properties] Properties to set
        */
        NewEventsQueryMsg(properties){
                return orderbooks.orderbooks.EventsQueryMsg(properties)
        }
        
        /**
        * Constructs a new SnapshotMsg.
        * @memberof OrderBooksClient
        * @classdesc Represents a SnapshotMsg.
        * @implements orderbooks.ISnapshotMsg
        * @constructor
        * @param {orderbooks.ISnapshotMsg=} [properties] Properties to set
        */
        NewSnapshotMsg(properties){
                return orderbooks.orderbooks.SnapshotMsg(properties)
        }
        
        /**
        * Constructs a new Event.
        * @memberof OrderBooksClient
        * @classdesc Represents a Event.
        * @implements orderbooks.IEvent
        * @constructor
        * @param {orderbooks.IEvent=} [properties] Properties to set
        */
        NewEvent(properties){
                return orderbooks.orderbooks.Event(properties)
        }
        

        // ----------------------------------------------

        /**
        * Send a message.
        * @memberof OrderBooksClient
        * @param {orderbooks.IEmptyMsg=} [inputVal] Properties to set
        * @return {Promise<MessageGetExchangesResponseMsg>}
        */
        SendGetExchanges(inputVal){
                return this._client.send("OrderBooks.GetExchanges",inputVal, null).promise;
        }
        
        OnGetExchanges(func){
                this._client.addMessageHandler("OrderBooks.GetExchanges",func);
        }

        /**
        * Send a message.
        * @memberof OrderBooksClient
        * @param {orderbooks.IEventsQueryMsg=} [inputVal] Properties to set
        * @return {Promise<MessageSnapshotMsg>}
        */
        SendQueryEvents(inputVal){
                return this._client.send("OrderBooks.QueryEvents",inputVal, null).promise;
        }
        
        OnQueryEvents(func){
                this._client.addMessageHandler("OrderBooks.QueryEvents",func);
        }

        /**
        * Send a message.
        * @memberof OrderBooksClient
        * @param {orderbooks.IEventsQueryMsg=} [inputVal] Properties to set
        * @return {Promise<MessageSnapshotMsg>}
        */
        SendGetLiveData(inputVal){
                return this._client.send("OrderBooks.GetLiveData",inputVal, null).promise;
        }
        
        OnGetLiveData(func){
                this._client.addMessageHandler("OrderBooks.GetLiveData",func);
        }

        /**
        * Send a message.
        * @memberof OrderBooksClient
        * @param {orderbooks.IEventsQueryMsg=} [inputVal] Properties to set
        * @return {Promise<MessageSnapshotMsg>}
        */
        SendGetCachedData(inputVal){
                return this._client.send("OrderBooks.GetCachedData",inputVal, null).promise;
        }
        
        OnGetCachedData(func){
                this._client.addMessageHandler("OrderBooks.GetCachedData",func);
        }

        
}



export default {OrderBooksClient, }