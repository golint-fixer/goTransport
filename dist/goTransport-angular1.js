var goTransport;
(function (goTransport) {
    "use strict";
    angular
        .module("goTransport", ['bd.sockjs']);
})(goTransport || (goTransport = {}));
var goTransport;
(function (goTransport) {
    "use strict";
    console.log('okay');
    function factory(socketFactory, $q) {
        var transport = {
            socket: null,
            connected: $q.defer(),
            id: 0,
            messageTypes: {
                MessageTypeMethod: 0,
                MessageTypePub: 1
            },
            message: function (message) {
                message.data = JSON.parse(message.data);
                console.log('receiving', message);
            },
            send: function (type, data) {
                this.connected.promise.then(function () {
                    this.socket.send(JSON.stringify({
                        id: this.id++,
                        type: type,
                        data: data
                    }));
                }.bind(this));
            }
        };
        var handlers = {};
        handlers[transport.messageTypes.MessageTypeMethod] = {
            call: function () {
            }
        };
        return {
            connect: function (url) {
                if (transport.socket == null) {
                    transport.socket = socketFactory({
                        url: url
                    });
                    transport.socket.setHandler('open', transport.connected.resolve);
                    transport.socket.setHandler('message', transport.message);
                }
                return transport.connected.promise;
            },
            onConnect: function () {
                return transport.connected.promise;
            },
            method: function (methodName, parameters) {
                transport.send(transport.messageTypes.MessageTypeMethod, {
                    name: methodName,
                    parameters: parameters
                });
            }
        };
    }
    factory.$inject = ["socketFactory", "$q"];
    angular
        .module("goTransport")
        .factory('goTransport', factory);
})(goTransport || (goTransport = {}));
//# sourceMappingURL=goTransport-angular1.js.map