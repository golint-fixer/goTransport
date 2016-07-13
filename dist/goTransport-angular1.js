var goTransport;
(function (goTransport) {
    "use strict";
    angular
        .module("goTransport", ['bd.sockjs']);
})(goTransport || (goTransport = {}));
var goTransport;
(function (goTransport) {
    class Callback {
    }
    goTransport.Callback = Callback;
})(goTransport || (goTransport = {}));
var goTransport;
(function (goTransport) {
    "use strict";
    class GoTransport {
        constructor($q, $timeout) {
            this.$q = $q;
            this.$timeout = $timeout;
            this.connectedPromise = $q.defer();
            this.callback = new goTransport.Callback();
        }
        static GetInstance($q, $timeout) {
            if (!GoTransport.instance)
                GoTransport.instance = new GoTransport($q, $timeout);
            return GoTransport.instance;
        }
        connect(url) {
            if (this.socket == null) {
                this.socket = Socket.Adapter.getSocket("SockJSClient", url, this);
            }
            return this.connectedPromise.promise;
        }
        connected() {
            this.connectedPromise.resolve();
        }
        message(data) {
            var message = goTransport.Message.fromJSON(data);
            console.log('receiving', message);
        }
        disconnected(code, reason, wasClean) {
            this.connectedPromise.reject(reason);
        }
        method(methodName, parameters) {
            var message = new goTransport.Message(new goTransport.MethodHandler(methodName, parameters));
            return message.send();
        }
        onConnect() {
            return this.connectedPromise.promise;
        }
    }
    goTransport.GoTransport = GoTransport;
    angular
        .module("goTransport")
        .factory("goTransport", ["$q", "$timeout", GoTransport.GetInstance]);
})(goTransport || (goTransport = {}));
var goTransport;
(function (goTransport) {
    (function (MessageType) {
        MessageType[MessageType["MessageTypeMethod"] = 0] = "MessageTypeMethod";
        MessageType[MessageType["MessageTypeMethodResult"] = 1] = "MessageTypeMethodResult";
        MessageType[MessageType["MessageTypePub"] = 2] = "MessageTypePub";
    })(goTransport.MessageType || (goTransport.MessageType = {}));
    var MessageType = goTransport.MessageType;
    class Message {
        constructor(messageHandler) {
            this.type = messageHandler.getMessageType();
            this.id = Message.current_id++;
            this.data = messageHandler.toJSON();
            if (!Message.promises) {
                Message.promises = [];
            }
        }
        send(timeout = 3000) {
            Message.promises[this.id] = goTransport.GoTransport.instance.$q.defer();
            goTransport.GoTransport.instance.connectedPromise.promise.then(function () {
                goTransport.GoTransport.instance.socket.send(JSON.stringify(this.toJSON()));
                goTransport.GoTransport.instance.$timeout(function () {
                    if (Message.promises[this.id].promise.$$state.status == 0) {
                        console.log("Timed out");
                        Message.promises[this.id].reject("Timed out");
                    }
                }.bind(this), timeout);
            }.bind(this));
            return Message.promises[this.id].promise;
        }
        toJSON() {
            return Object.assign({}, this, {});
        }
        static fromJSON(json) {
            if (typeof json === 'string') {
                return JSON.parse(json, Message.reviver);
            }
            else {
                let message = Object.create(Message.prototype);
                return Object.assign(message, json, {});
            }
        }
        static reviver(key, value) {
            return key === "" ? Message.fromJSON(value) : value;
        }
    }
    Message.current_id = 0;
    goTransport.Message = Message;
})(goTransport || (goTransport = {}));
var goTransport;
(function (goTransport) {
    class MessageHandler {
        constructor(messageType) {
            this.messageType = messageType;
        }
        getMessageType() {
            return this.messageType;
        }
    }
    goTransport.MessageHandler = MessageHandler;
})(goTransport || (goTransport = {}));
var goTransport;
(function (goTransport) {
    class MethodHandler extends goTransport.MessageHandler {
        constructor(name, parameters) {
            super(goTransport.MessageType.MessageTypeMethod);
            this.name = name;
            this.parameters = parameters;
        }
        validate() {
            return null;
        }
        run() {
            return null;
        }
        toJSON() {
            return Object.assign({}, this, {});
        }
    }
    goTransport.MethodHandler = MethodHandler;
})(goTransport || (goTransport = {}));
var Socket;
(function (Socket) {
    class SockJSClient {
        constructor(url, delegate) {
            this.delegate = delegate;
            this.connection = new SockJS(url);
            this.connection.onopen = function (e) { this.open(e); }.bind(this);
            this.connection.onclose = function (e) { this.disconnect(e); }.bind(this);
            this.connection.onmessage = function (e) { this.message(e); }.bind(this);
        }
        static getInstance(url, delegate) {
            if (!SockJSClient.instance) {
                SockJSClient.instance = new SockJSClient(url, delegate);
            }
            return SockJSClient.instance;
        }
        open(e) {
            this.delegate.connected();
        }
        disconnect(e) {
            this.delegate.disconnected(e.code, e.reason, e.wasClean);
        }
        message(e) {
            this.delegate.message(e.data);
        }
        send(data) {
            this.connection.send(data);
        }
        close() {
            this.connection.close();
        }
    }
    Socket.SockJSClient = SockJSClient;
})(Socket || (Socket = {}));
var Socket;
(function (Socket) {
    class Adapter {
        static getSocket(type, url, delegate) {
            switch (type) {
                case "SockJSClient":
                    return Socket.SockJSClient.getInstance(url, delegate);
                default:
                    throw ("Invalid socket type:" + type);
            }
        }
    }
    Socket.Adapter = Adapter;
})(Socket || (Socket = {}));
//# sourceMappingURL=goTransport-angular1.js.map