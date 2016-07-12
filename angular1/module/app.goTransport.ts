/// <reference path="typings/angularjs/angular.d.ts" />
/// <reference path="typings/angularjs/angular-route.d.ts" />

module goTransport {
    "use strict";

    function factory(socketFactory: any, $q : ng.IQService){
        var transport = {
            socket: null,
            connected: $q.defer(),
            id: 0,
            messageTypes: {
                MessageTypeMethod: 0,
                MessageTypePub: 1
            },
            message: function(message) {
                message.data = JSON.parse(message.data);
                console.log('receiving', message);
            },
            send: function(type, data) {
                this.connected.promise.then(function() {
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
            call: function() {

            }
        };

        return {
            connect: function(url) {
                if(transport.socket == null) {
                    transport.socket = socketFactory({
                        url: url
                    });
                    transport.socket.setHandler('open', transport.connected.resolve);
                    transport.socket.setHandler('message', transport.message);
                    // transport.socket.setHandler('close', connected.reject);
                }
                return transport.connected.promise;
            },
            //Can be called anywhere. Tells us when
            onConnect: function() {
                return transport.connected.promise;
            },
            method: function(methodName, parameters) {
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
}