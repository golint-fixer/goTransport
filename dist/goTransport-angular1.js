var goTransport;
(function (goTransport) {
    class Client {
        constructor() {
            this.messageManager = new goTransport.MessageManager();
            Client.instance = this;
        }
        connect(url) {
            return this.messageManager.connect(url);
        }
        method(name, parameters, timeout = 3000) {
            let message = new goTransport.MessageMethod(name, parameters);
            this.messageManager.send(message);
            var promise = message.getPromise();
            promise.setTimeOut(timeout);
            return promise.promise;
        }
        onConnect() {
            return this.messageManager.getConnectedPromise();
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
        getType() {
            return this.type;
        }
        setReply(message) {
            this.reply = message;
        }
        start() {
            var error = this.validate();
            if (error) {
                console.error(error);
                return false;
            }
            error = this.run();
            if (error) {
                console.error(error);
                return false;
            }
            return true;
        }
        serialize() {
            return this.type + Message.headerDelimiter + JSON.stringify(this.encode());
        }
        static unSerialize(data) {
            var parts = data.split(Message.headerDelimiter);
            if (parts[1] === undefined) {
                console.warn("Invalid message", data);
                return null;
            }
            return goTransport.MessageDefinition.get(parseInt(parts[0]), JSON.parse(parts[1]));
        }
        encode() {
            return Object.assign({}, this, {});
        }
    }
    Message.current_id = 0;
    Message.headerDelimiter = "\f";
    goTransport.Message = Message;
})(goTransport || (goTransport = {}));
var goTransport;
(function (goTransport) {
    class MessageBuilder {
        constructor(messageType) {
            this.messageType = messageType;
        }
        build() {
            return new this.messageType();
        }
    }
    goTransport.MessageBuilder = MessageBuilder;
})(goTransport || (goTransport = {}));
var goTransport;
(function (goTransport) {
    (function (MessageType) {
        MessageType[MessageType["MessageTypeMethod"] = 0] = "MessageTypeMethod";
        MessageType[MessageType["MessageTypeMethodResult"] = 1] = "MessageTypeMethodResult";
        MessageType[MessageType["MessageTypeError"] = 2] = "MessageTypeError";
        MessageType[MessageType["MessageTypePub"] = 3] = "MessageTypePub";
    })(goTransport.MessageType || (goTransport.MessageType = {}));
    var MessageType = goTransport.MessageType;
    class MessageDefinition {
        static set(type, definition) {
            if (!definition || !definition.prototype) {
                console.warn("Invalid message definition set for type", type);
                return;
            }
            MessageDefinition.definitions[type] = definition;
        }
        static get(type, data) {
            var definition = MessageDefinition.definitions[type];
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
    MessageDefinition.definitions = Array();
    goTransport.MessageDefinition = MessageDefinition;
})(goTransport || (goTransport = {}));
var goTransport;
(function (goTransport) {
    class MessageManager {
        constructor() {
            this.messages = [];
        }
        connect(url) {
            this.connectedPromise = new goTransport.Promise();
            if (this.socket == null) {
                this.socket = Socket.Adapter.getSocket("SockJSClient", url, this);
            }
            return this.getConnectedPromise();
        }
        getConnectedPromise() {
            return this.connectedPromise.promise;
        }
        set(message) {
            this.messages[message.id] = message;
        }
        get(message) {
            return this.messages[message.id];
        }
        connected() {
            console.log('connected');
            this.connectedPromise.resolve();
        }
        send(message) {
            message.start();
            this.set(message);
            this.getConnectedPromise().then(function () {
                this.socket.send(message.serialize());
                console.log('sent', message.serialize());
            }.bind(this));
        }
        messaged(data) {
            let message = goTransport.Message.unSerialize(data);
            if (!message) {
                console.warn("Invalid message received.");
                return;
            }
            message.setReply(this.get(message));
            message.start();
        }
        disconnected(code, reason, wasClean) {
            console.warn('Disconnected', code);
        }
    }
    goTransport.MessageManager = MessageManager;
})(goTransport || (goTransport = {}));
var goTransport;
(function (goTransport) {
    class MessageMethod extends goTransport.Message {
        constructor(name = null, parameters = null) {
            super(MessageMethod.type);
            this.name = name;
            this.parameters = parameters;
        }
        validate() {
            return null;
        }
        run() {
            console.log('ran');
            this.promise = new goTransport.Promise();
            return null;
        }
        getPromise() {
            return this.promise;
        }
    }
    MessageMethod.type = goTransport.MessageType.MessageTypeMethod;
    goTransport.MessageMethod = MessageMethod;
    goTransport.MessageDefinition.set(MessageMethod.type, MessageMethod);
})(goTransport || (goTransport = {}));
var goTransport;
(function (goTransport) {
    class MessageMethodResult extends goTransport.Message {
        constructor(result = false, parameters = null) {
            super(goTransport.MessageMethod.type);
            this.result = result;
            this.parameters = parameters;
        }
        validate() {
            console.log('validating method result', this.reply);
            return null;
        }
        run() {
            console.log('Running method result', this.reply);
            return null;
        }
    }
    MessageMethodResult.type = goTransport.MessageType.MessageTypeMethodResult;
    goTransport.MessageMethodResult = MessageMethodResult;
    goTransport.MessageDefinition.set(goTransport.MessageMethod.type, goTransport.MessageMethod);
})(goTransport || (goTransport = {}));
var Socket;
(function (Socket) {
    class SockJSClient {
        constructor(url, delegate) {
            this.delegate = delegate;
            this.connection = new SockJS(url);
            this.connection.onopen = this.open.bind(this);
            this.connection.onclose = this.disconnect.bind(this);
            this.connection.onmessage = this.message.bind(this);
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
        constructor($q) {
            goTransport.Promise.$q = $q;
            super();
        }
        static getInstance($q) {
            if (!Angular1.instance)
                Angular1.instance = new Angular1($q);
            return Angular1.instance;
        }
    }
    goTransport.Angular1 = Angular1;
    "use strict";
    angular
        .module("goTransport", ['bd.sockjs']);
    angular
        .module("goTransport")
        .factory("goTransport", ["$q", Angular1.getInstance]);
})(goTransport || (goTransport = {}));
var goTransport;
(function (goTransport) {
    class Promise {
        constructor(timeout = 0) {
            this.timeout = timeout;
            this.defer = Promise.$q.defer();
            this.promise = this.defer.promise;
            this.setTimeOut(timeout);
        }
        resolve(value) {
            this.defer.resolve(value);
        }
        reject(reason) {
            this.defer.reject(reason);
        }
        notify(state) {
            this.defer.notify(state);
        }
        setTimeOut(timeout = 3000) {
            this.timeout = timeout;
            if (this.timer)
                clearTimeout(this.timer);
            if (timeout > 0) {
                this.timer = setTimeout(this.timedOut.bind(this), timeout);
            }
        }
        timedOut() {
            if (this.promise.$$state.status == 0) {
                this.defer.reject("Timed out. Exceeded:" + this.timeout);
            }
        }
    }
    goTransport.Promise = Promise;
})(goTransport || (goTransport = {}));
//# sourceMappingURL=goTransport-angular1.js.map