# GoTransport
[![GoDoc](https://godoc.org/github.com/iain17/goTransport?status.svg)](https://godoc.org/github.com/iain17/goTransport)
[![Build Status: Linux](https://travis-ci.org/iain17/goTransport.svg?branch=master)](https://travis-ci.org/iain17/goTransport)
[![Build Status: Windows](https://ci.appveyor.com/api/projects/status/vcvontw1aus6ixln/branch/master?svg=true)](https://ci.appveyor.com/project/zg/service/branch/master)
[![Coverage Status](https://codecov.io/gh/iain17/goTransport/branch/master/graph/badge.svg)](https://codecov.io/gh/iain17/goTransport)
[![Go Report Card](https://goreportcard.com/badge/github.com/iain17/goTransport)](https://goreportcard.com/report/github.com/iain17/goTransport)
[![Gitter](https://badges.gitter.im/join_chat.svg)](https://gitter.im/iain17/goTransport)

[GoTransport](https://github.com/iain17/goTransport) GoLang-SockJS-Angular 1, RPC++ socket.
A RPC SockJS GoLang library. Starting with an Angular 1 client this library tries to fulfill the dire need I have for a proper javascript to Go RPC socket connection.
Eventually it should also pub and sub on events and sync collections of data.

Building on top off the work of: [Igor Mihalik's sockjs-go](https://github.com/igm/sockjs-go).

##GoTransport - Angular 1 module
You must have Bower and GoLang installed before you can continue.
- [Documentation about installing bower](https://bower.io/#install-bower)
- [Documentation about installing GoLang](https://golang.org/doc/install)
- Basic knowledge of both Angular and GoLang.

###Install
1. `bower install goTransport-angular1`
2. Made sure the [SockJS client library](https://github.com/sockjs/sockjs-client) is loaded. (It should automatically get loaded in as a bower dependency)
3. Add `goTransport` as a module dependency to your app.
4. Open your Go application and `import "github.com/iain17/goTransport/transport"`

###How to get the example running
1. `go get github.com/iain17/goTransport/transport`
2. `cd "$GOPATH/src/github.com/iain17/goTransport"`
2. `bower install`
3. `cd transport/example`
4. `go run main.go`
5. Open a browser and navigate to localhost:8081/angular1/example.
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

		$scope.pong = '';
		$scope.ping = function() {
			goTransport.method('ping', ['hai']).then(function(message, err) {
			    if(err) {
			        console.error(err);
			        return;
			    }
			    console.log(result);
            });
        };
    });
```
**Example Server:**
```go
package main

import (
	"net/http"
	"log"
	"github.com/iain17/goTransport/transport"
)

func main() {
	transporter := transport.NewTransport("/ws")
	transporter.SetRPCMethod("ping", ping)
	http.Handle("/ws/", transporter.HttpHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
//Free parameters and return values. No array of interfaces.
func ping(message string) (string, error) {
	log.Print("called", message)
	return "bar", nil
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
