# goTransport
[GoTransport](https://github.com/iain17/goTransport) GoLang-SockJS-Angular 1, RPC socket.
A RPC SockJS GoLang library. Starting with an Angular 1 client this library tries to fulfill the dire need I have for a proper javascript to Go RPC socket connection.
Eventually it should also pub and sub on events and collections of data.

Building on top off the work of: [Ben Drucker's angular-sockjs](https://github.com/bendrucker/angular-sockjs) and [Igor Mihalik's sockjs-go](https://github.com/igm/sockjs-go).

##GoTransport - Angular 1 module
###Install
1. `bower install goTransport-angular1`
2. Made sure the [SockJS client library](https://github.com/sockjs/sockjs-client) is loaded. (It should automatically get loaded in as a bower dependency)
3. Add `goTransport` as a module dependency to your app.
4. Open your Go application and `import "github.com/iain17/goTransport/transport"`

###How to get the example running
You must have Bower and GoLang installed before you can continue.
- [Documentation about installing bower](https://bower.io/#install-bower)
- [Documentation about installing GoLang](https://golang.org/doc/install)
1. Git clone this repo and open terminal and navigate to this directory.
2. `bower install`
3. `cd transport/example`
4. `go get .`
5. `go run main.go`
6. Open a browser and navigate to localhost:8081/angular1/example.
This should fetch all the dependencies using bower and run a basic go http server serving the static files and GoTransport socket instance for the example.

## Usage
### Making a Remote procedure call
```javascript
angular.module('goTransport-example', ['goTransport'])
	.controller('mainController', function($scope, goTransport) {
		goTransport.connect('http://localhost:8081/ws');

		goTransport.onConnect().then(function() {
			console.log('Connected!');
		});

		$scope.pong = '';
		$scope.ping = function() {
			goTransport.method('ping', ['hai']);
		};
	});
```
### Publish–subscribe.

### Data-collection sync.

## Develop
The angular 1 module is written in Typescript. In order to get started my preference is to do as follows:
1. Make sure both npm and typescript are installed.
2. `npm install -g tsd`
3. `cd angular1/module/`
4. `tsd install angular`
5. Open the project using JetBrains WebStorm.