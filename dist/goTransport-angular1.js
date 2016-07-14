var goTransport;
(function (goTransport) {
    class Client {
        constructor($q, $timeout) {
            this.$q = $q;
            this.$timeout = $timeout;
            this.connected = $q.defer();
            this.messageManager = new goTransport.MessageManager(this);
        }
        connect(url) {
            return this.messageManager.connect(url);
        }
        method(methodName, parameters) {
            return null;
        }
        onConnect() {
            return this.connected.promise;
        }
    }
    goTransport.Client = Client;
})(goTransport || (goTransport = {}));
var goTransport;
(function (goTransport) {
    class Message {
        constructor(type) {
            this.type = type;
            this.id = Message.current_id++;
        }
    }
    Message.current_id = 0;
    goTransport.Message = Message;
})(goTransport || (goTransport = {}));
var goTransport;
(function (goTransport) {
    class MessageBuilder {
        constructor(testType) {
            this.testType = testType;
        }
        build() {
            return new this.testType();
        }
    }
    goTransport.MessageBuilder = MessageBuilder;
})(goTransport || (goTransport = {}));
var goTransport;
(function (goTransport) {
    (function (MessageType) {
        MessageType[MessageType["MessageTypeMethod"] = 0] = "MessageTypeMethod";
        MessageType[MessageType["MessageTypeMethodResult"] = 1] = "MessageTypeMethodResult";
        MessageType[MessageType["MessageTypePub"] = 2] = "MessageTypePub";
    })(goTransport.MessageType || (goTransport.MessageType = {}));
    var MessageType = goTransport.MessageType;
    class MessageDefinitions {
        static set(type, definition) {
            if (!definition || !definition.prototype) {
                console.warn("Invalid message definition set for type", type);
                return;
            }
            MessageDefinitions.definitions[type] = definition;
        }
        static get(type, data) {
            var definition = MessageDefinitions.definitions[type];
            if (definition === undefined) {
                console.warn("Invalid messageType requested", type);
                return null;
            }
            let messageBuilder = new goTransport.MessageBuilder(definition);
            var message = messageBuilder.build();
            Object.assign(message, JSON.parse(data), {});
            return message;
        }
    }
    MessageDefinitions.definitions = Array();
    goTransport.MessageDefinitions = MessageDefinitions;
})(goTransport || (goTransport = {}));
var goTransport;
(function (goTransport) {
    class MessageManager {
        constructor(client) {
            this.client = client;
        }
        connect(url) {
            if (this.socket == null) {
                this.socket = Socket.Adapter.getSocket("SockJSClient", url, this);
            }
            return this.client.connected.promise;
        }
        message(data) {
            console.log('send', data);
            this.socket.send(data);
        }
        connected() {
            this.client.connected.resolve();
        }
        messaged(data) {
            console.log('received', data);
        }
        disconnected(code, reason, wasClean) {
            console.log('Disconnected');
        }
    }
    goTransport.MessageManager = MessageManager;
})(goTransport || (goTransport = {}));
var goTransport;
(function (goTransport) {
    class MessageMethod extends goTransport.Message {
        constructor() {
            super(MessageMethod.type);
        }
        validate() {
            return null;
        }
        run() {
            console.log('ran');
            return null;
        }
    }
    MessageMethod.type = goTransport.MessageType.MessageTypeMethod;
    goTransport.MessageMethod = MessageMethod;
    goTransport.MessageDefinitions.set(MessageMethod.type, MessageMethod);
    var message = goTransport.MessageDefinitions.get(goTransport.MessageType.MessageTypeMethod, '{"name": "test"}');
    console.log(message);
    message = goTransport.MessageDefinitions.get(goTransport.MessageType.MessageTypeMethod, '{"name": "test"}');
    console.log(message);
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
            this.delegate.messaged(e.data);
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
var goTransport;
(function (goTransport) {
    "use strict";
    class Angular1 extends goTransport.Client {
        constructor($q, $timeout) {
            super($q, $timeout);
            this.$q = $q;
            this.$timeout = $timeout;
        }
        static getInstance($q, $timeout) {
            if (!Angular1.instance)
                Angular1.instance = new Angular1($q, $timeout);
            return Angular1.instance;
        }
    }
    goTransport.Angular1 = Angular1;
    "use strict";
    angular
        .module("goTransport", ['bd.sockjs']);
    angular
        .module("goTransport")
        .factory("goTransport", ["$q", "$timeout", Angular1.getInstance]);
})(goTransport || (goTransport = {}));
//# sourceMappingURL=goTransport-angular1.js.map