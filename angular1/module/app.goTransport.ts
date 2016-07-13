/// <reference path="typings/angularjs/angular.d.ts" />

module goTransport {
    "use strict";
    import Adapter = Socket.Adapter;

    export class GoTransport implements Socket.SocketDelegate{
        public static socket: Socket.Socket;
        public static connected: ng.IDeferred<{}>;

        constructor(private $q : ng.IQService) {
            this.$q = $q;
            GoTransport.connected = $q.defer();
        }

        public static Main($q : ng.IQService): GoTransport {
            return new GoTransport($q);
        }

        public connect(url : string): ng.IPromise<{}> {
            if(GoTransport.socket == null) {
                GoTransport.socket = Socket.Adapter.getSocket("SockJSClient", url, this);
            }
            return GoTransport.connected.promise;
        }

        public connected() {
            console.log('connected');
            GoTransport.connected.resolve();
        }

        public message(data:string) {
            var message = Message.fromJSON(data);
            console.log('receiving', message);
        }

        public disconnected(code:number, reason:string, wasClean:boolean) {
            console.log(code);
        }

        private send(type : MessageType, data : any) {
            var message = new Message(type, data);
            message.send();
        }

        public method(methodName: string, parameters: any[]): ng.IPromise<{}> {
            //TODO: proper promise
            var q = this.$q.defer();
            this.send(MessageType.MessageTypeMethod, {
                name: methodName,
                parameters: parameters
            });//TODO: Method class
            return q.promise;
        }

        public onConnect(): ng.IPromise<{}> {
            return GoTransport.connected.promise;
        }

    }

    //Attach the above to angular
    angular
        .module("goTransport")
        .factory("goTransport", ["$q", GoTransport.Main]);
}