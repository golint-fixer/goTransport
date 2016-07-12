'use strict';

angular.module('goTransport', [
		'bd.sockjs'
	])
	.factory('goTransport', function (socketFactory, $q) {
		var transport = {
			socket: null,
			connected: $q.defer(),
			messageTypes: {
				MessageTypeMethod: 0,
				MessageTypePub: 1
			},
			message: function(message) {
				message.data = JSON.parse(message.data);
				console.log('receiving', message);
			},
			send: function(type, data) {
				if(this.socket == null)
					return;
				console.log('sending', JSON.stringify({
					type: type,
					data: data
				}));

				this.socket.send(JSON.stringify({
					type: type,
					data: data
				}));
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

	});