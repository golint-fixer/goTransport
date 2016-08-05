var goTransport;
(function (goTransport) {
    class Client {
        constructor() {
            this.messageManager = new goTransport.Session();
            Client.instance = this;
        }
        connect(url) {
            return this.messageManager.connect(url);
        }
        call(name, parameters, timeout = 3000) {
            let message = new goTransport.MessageMethod(name, parameters);
            this.messageManager.send(message);
            var promise = message.getPromise();
            promise.setTimeOut(timeout);
            return promise.promise;
        }
        method(name) {
        }
        onConnect() {
            return this.messageManager.getConnectedPromise();
        }
    }
    goTransport.Client = Client;
})(goTransport || (goTransport = {}));
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
var goTransport;
(function (goTransport) {
    class Message {
        constructor(type) {
            this.type = type;
            this.id = null;
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
                return error;
            }
            error = this.run();
            if (error) {
                return error;
            }
            return null;
        }
        serialize() {
            return this.type + Message.headerDelimiter + JSON.stringify(this.encode());
        }
        static unSerialize(data) {
            var parts = data.split(Message.headerDelimiter);
            if (parts[1] === undefined) {
                console.warn("Invalid message. Invalid amount of parts", data);
                return null;
            }
            console.log(parts);
            return goTransport.MessageDefinition.get(parseInt(parts[0]), parts[1]);
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
        MessageType[MessageType["MessageTypeTest"] = 0] = "MessageTypeTest";
        MessageType[MessageType["MessageTypeMethod"] = 1] = "MessageTypeMethod";
        MessageType[MessageType["MessageTypeMethodResult"] = 2] = "MessageTypeMethodResult";
        MessageType[MessageType["MessageTypeError"] = 3] = "MessageTypeError";
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
    class Session {
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
            console.log('messageManager', 'set', message.id);
            this.messages[message.id] = message;
        }
        get(message) {
            console.log('messageManager', 'get', message.id, this.messages[message.id]);
            return this.messages[message.id];
        }
        connected() {
            console.log('connected');
            this.connectedPromise.resolve();
        }
        send(message) {
            goTransport.Message.current_id++;
            message.id = goTransport.Message.current_id;
            message.start();
            this.set(message);
            this.getConnectedPromise().then(function () {
                this.socket.send(message.serialize());
                console.log('sent', message.serialize());
            }.bind(this));
        }
        messaged(data) {
            console.debug('received', data);
            let message = goTransport.Message.unSerialize(data);
            if (!message) {
                console.warn("Invalid message received.");
                return;
            }
            message.setReply(this.get(message));
            var error = message.start();
            if (error != null) {
                console.error(error);
            }
        }
        disconnected(code, reason, wasClean) {
            console.warn('Disconnected', code);
        }
    }
    goTransport.Session = Session;
})(goTransport || (goTransport = {}));
var goTransport;
(function (goTransport) {
    class MessageError extends goTransport.Message {
        constructor(reason) {
            super(MessageError.type);
            this.reason = reason;
        }
        validate() {
            return null;
        }
        run() {
            console.error(this.reason);
            if ((this.reply instanceof goTransport.MessageMethod)) {
                let promise = this.reply.getPromise();
                if (promise) {
                    console.debug(this);
                    promise.reject(this.reason);
                }
            }
            return null;
        }
    }
    MessageError.type = goTransport.MessageType.MessageTypeError;
    goTransport.MessageError = MessageError;
    goTransport.MessageDefinition.set(MessageError.type, MessageError);
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
            super(MessageMethodResult.type);
            this.result = result;
            this.parameters = parameters;
        }
        validate() {
            if (!(this.reply instanceof goTransport.MessageMethod)) {
                console.debug(this.reply);
                return new Error("Invalid reply. Not messageMethod.");
            }
        }
        run() {
            console.log('Result came back!', this.parameters);
            let promise = this.reply.getPromise();
            if (promise) {
                promise.resolve.apply(promise, this.parameters);
            }
            return null;
        }
    }
    MessageMethodResult.type = goTransport.MessageType.MessageTypeMethodResult;
    goTransport.MessageMethodResult = MessageMethodResult;
    goTransport.MessageDefinition.set(MessageMethodResult.type, MessageMethodResult);
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
//# sourceMappingURL=goTransport-angular1.js.map