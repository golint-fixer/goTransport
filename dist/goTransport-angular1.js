var Socket;
(function (Socket) {
    var Adapter = (function () {
        function Adapter() {
        }
        Adapter.getSocket = function (type, url, delegate) {
            switch (type) {
                case "SockJSClient":
                    return Socket.SockJSClient.getInstance(url, delegate);
                default:
                    throw ("Invalid socket type:" + type);
            }
        };
        return Adapter;
    }());
    Socket.Adapter = Adapter;
})(Socket || (Socket = {}));
var Socket;
(function (Socket) {
    var SockJSClient = (function () {
        function SockJSClient(url, delegate) {
            this.delegate = delegate;
            this.connection = new SockJS(url);
            this.connection.onopen = function (e) { this.open(e); }.bind(this);
            this.connection.onclose = function (e) { this.disconnect(e); }.bind(this);
            this.connection.onmessage = function (e) { this.message(e); }.bind(this);
        }
        SockJSClient.getInstance = function (url, delegate) {
            console.log('Ik ben bezig');
            if (!SockJSClient.instance) {
                console.log('new instance');
                SockJSClient.instance = new SockJSClient(url, delegate);
            }
            return SockJSClient.instance;
        };
        SockJSClient.prototype.open = function (e) {
            console.log('deletegate', this.delegate);
            this.delegate.connected();
        };
        SockJSClient.prototype.disconnect = function (e) {
            this.delegate.disconnected(e.code, e.reason, e.wasClean);
        };
        SockJSClient.prototype.message = function (e) {
            this.delegate.message(e.data);
        };
        SockJSClient.prototype.send = function (data) {
            this.connection.send(data);
        };
        SockJSClient.prototype.close = function () {
            this.connection.close();
        };
        return SockJSClient;
    }());
    Socket.SockJSClient = SockJSClient;
})(Socket || (Socket = {}));
var goTransport;
(function (goTransport) {
    "use strict";
    angular
        .module("goTransport", ['bd.sockjs']);
})(goTransport || (goTransport = {}));
var goTransport;
(function (goTransport) {
    "use strict";
    var GoTransport = (function () {
        function GoTransport($q) {
            this.$q = $q;
            this.$q = $q;
            GoTransport.connected = $q.defer();
        }
        GoTransport.Main = function ($q) {
            return new GoTransport($q);
        };
        GoTransport.prototype.connect = function (url) {
            if (GoTransport.socket == null) {
                GoTransport.socket = Socket.Adapter.getSocket("SockJSClient", url, this);
            }
            return GoTransport.connected.promise;
        };
        GoTransport.prototype.connected = function () {
            console.log('connected');
            GoTransport.connected.resolve();
        };
        GoTransport.prototype.message = function (data) {
            var message = goTransport.Message.fromJSON(data);
            console.log('receiving', message);
        };
        GoTransport.prototype.disconnected = function (code, reason, wasClean) {
            console.log(code);
        };
        GoTransport.prototype.send = function (type, data) {
            var message = new goTransport.Message(type, data);
            message.send();
        };
        GoTransport.prototype.method = function (methodName, parameters) {
            var q = this.$q.defer();
            this.send(goTransport.MessageType.MessageTypeMethod, {
                name: methodName,
                parameters: parameters
            });
            return q.promise;
        };
        GoTransport.prototype.onConnect = function () {
            return GoTransport.connected.promise;
        };
        return GoTransport;
    }());
    goTransport.GoTransport = GoTransport;
    angular
        .module("goTransport")
        .factory("goTransport", ["$q", GoTransport.Main]);
})(goTransport || (goTransport = {}));
var goTransport;
(function (goTransport) {
    (function (MessageType) {
        MessageType[MessageType["MessageTypeMethod"] = 0] = "MessageTypeMethod";
        MessageType[MessageType["MessageTypeMethodResult"] = 1] = "MessageTypeMethodResult";
        MessageType[MessageType["MessageTypePub"] = 2] = "MessageTypePub";
    })(goTransport.MessageType || (goTransport.MessageType = {}));
    var MessageType = goTransport.MessageType;
    var Message = (function () {
        function Message(type, data) {
            this.type = type;
            this.id = Message.current_id++;
            this.data = data;
        }
        Message.prototype.send = function () {
            goTransport.GoTransport.connected.promise.then(function () {
                goTransport.GoTransport.socket.send(JSON.stringify(this.toJSON()));
                console.log('sent');
            }.bind(this));
        };
        Message.prototype.toJSON = function () {
            return Object.assign({}, this, {});
        };
        Message.fromJSON = function (json) {
            if (typeof json === 'string') {
                return JSON.parse(json, Message.reviver);
            }
            else {
                var message = Object.create(Message.prototype);
                return Object.assign(message, json, {});
            }
        };
        Message.reviver = function (key, value) {
            return key === "" ? Message.fromJSON(value) : value;
        };
        Message.current_id = 0;
        return Message;
    }());
    goTransport.Message = Message;
})(goTransport || (goTransport = {}));
//# sourceMappingURL=goTransport-angular1.js.map