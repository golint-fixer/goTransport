angular.module('goTransport-example', ['goTransport'])
	.controller('mainController', function($scope, goTransport) {
		goTransport.connect('http://localhost:8081/ws');

		goTransport.onConnect().then(function() {
			console.log('Connected!');
		});

		$scope.pong = '';
		$scope.ping = function() {
			goTransport.method('ping', ['hai']).then(function(result, err) {
				if(err) {
					console.error(err);
					return;
				}
				console.log(result);
			});
		};
	});