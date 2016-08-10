# GoTransport
[![GoDoc](https://godoc.org/github.com/iain17/goTransport?status.svg)](https://godoc.org/github.com/iain17/goTransport)
[![Build Status: Linux](https://travis-ci.org/iain17/goTransport.svg?branch=master)](https://travis-ci.org/iain17/goTransport)
[![Coverage Status](https://codecov.io/gh/iain17/goTransport/branch/master/graph/badge.svg)](https://codecov.io/gh/iain17/goTransport)
[![Go Report Card](https://goreportcard.com/badge/github.com/iain17/goTransport?v=1)](https://goreportcard.com/report/github.com/iain17/goTransport)
[![Gitter](https://badges.gitter.im/join_chat.svg)](https://gitter.im/iain17/goTransport)
After years of working with MeteorJS in combination with Angular, I wanted to recreate and improve to some extent some of that magic but using Go as a backend.

This project is designed to do the following:
- Easy Remote Procedure Calls (RPC) on both server (GoLang) and client (Angular)
- Sync collections of data.

This repository is a work in progress. Anything is Master is considered Stable. Its not recommended to use this yet in release.

Building on top off the work of: [Igor Mihalik's sockjs-go](https://github.com/igm/sockjs-go).

## Omissions
You must have Bower and GoLang installed before you can continue.
- [Documentation about installing bower](https://bower.io/#install-bower)
- [Documentation about installing GoLang](https://golang.org/doc/install)
- Basic knowledge of both Angular and GoLang.

###Install
1. `bower install goTransport-client`
2. Made sure the [SockJS client library](https://github.com/sockjs/sockjs-client) is loaded. (It should automatically get loaded in as a bower dependency)
3. Add `goTransport` as a module dependency to your app.
4. Open your Go application and `import "github.com/iain17/goTransport"`

###How to get the example running
1. `go get github.com/iain17/goTransport`
2. `cd "$GOPATH/src/github.com/iain17/goTransport/goTransport-client"`
2. `bower install`
3. `cd ../example`
4. `go run main.go`
5. Open a browser and navigate to localhost:8081/src/example/.
This should fetch all the dependencies using bower and run a basic go http server serving the static files and GoTransport socket instance for the example.

## Usage
### Making a Remote procedure call
**Example Client:**
```javascript
angular.module('goTransport-example', ['goTransport'])
	.controller('mainController', function($scope, goTransport) {
		goTransport.connect('http://localhost:8081/ws');

		goTransport.onConnect().then(function() {
			console.log('Connected!');
		});

		//Bidirectional method calling. With dynamic parameters

		//Server calling the client and sending back a optional response.
		goTransport.method('example', function(message, number) {
			console.log(message, number);
			return "Hello there server :-)";
		});

		//Client calling the server and getting a response.
		$scope.pong = '';
		$scope.ping = function() {
			goTransport.call('ping', ['hai']).then(function(result, err) {
				if(err) {
					console.error(err);
					return;
				}
				console.log(result);
				$scope.pong = result;
			}, function(err) {
				console.error(err);
			});
		};

	});
```
**Example Server:**
```go
package main

import (
	"errors"
	"github.com/iain17/goTransport"
	"github.com/iain17/goTransport/lib/interfaces"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("../goTransport-client/")))
	transporter := goTransport.New("/ws", connected)
	transporter.Method("ping", ping)

	log.Print("goTransport server spawning at port: 8081")
	log.Print("Angular 1 example available at: localhost:8081/src/example/")
	http.Handle("/ws/", transporter.GetHttpHandler())
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func connected(session interfaces.ICallableSession) {
	log.Printf("New client connected: %s", session.GetId())

	log.Print("Calling example method client side.")
	promise := session.Call("example", []interface{}{
		"A test",
		1337,
	}, 0)
	promise.OnSuccess(func(v interface{}) {
		log.Print("Success: ", v)
	}).OnFailure(func(v interface{}) {
		log.Print("Failure: ", v)
	})
}

//Free parameters and return values. No array of interfaces.
func ping(session interfaces.ICallableSession, message string) (string, error) {
	log.Print("called with parameter: ", message)
	return "bar", errors.New("test")
}

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
