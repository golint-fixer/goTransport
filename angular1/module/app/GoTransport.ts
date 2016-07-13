module goTransport {
    "use strict";
    import Adapter = Socket.Adapter;

    export class GoTransport implements Socket.SocketDelegate{
        public socket: Socket.Socket;
        public connectedPromise: ng.IDeferred<{}>;
        public static instance: GoTransport;
        public callback : Callback;

        constructor(public $q : ng.IQService, public $timeout : ng.ITimeoutService) {
            this.connectedPromise = $q.defer();
            this.callback = new Callback();
        }

        public static GetInstance($q : ng.IQService, $timeout : ng.ITimeoutService): GoTransport {
            if(!GoTransport.instance)
                GoTransport.instance = new GoTransport($q, $timeout);
            return GoTransport.instance;
        }

        public connect(url : string): ng.IPromise<{}> {
            if(this.socket == null) {
                this.socket = Socket.Adapter.getSocket("SockJSClient", url, this);
            }
            return this.connectedPromise.promise;
        }

        public connected() {
            this.connectedPromise.resolve();
        }

        public message(data:string) {
            var message = Message.fromJSON(data);
            console.log('receiving', message);
        }

        public disconnected(code:number, reason:string, wasClean:boolean) {
            this.connectedPromise.reject(reason);
        }

        public method(methodName: string, parameters: any[]): ng.IPromise<{}> {
            var message = new Message(
                new MethodHandler(methodName, parameters)
            );
            return message.send();
        }

        public onConnect(): ng.IPromise<{}> {
            return this.connectedPromise.promise;
        }

    }

    //Attach the above to angular
    angular
        .module("goTransport")
        .factory("goTransport", ["$q", "$timeout", GoTransport.GetInstance]);
}