module goTransport {

    export class Client {

        public connected: ng.IDeferred<{}>;
        public static instance: Client;
        private messageManager: MessageManager;

        constructor(public $q : ng.IQService, public $timeout : ng.ITimeoutService) {
            this.connected = $q.defer();
            this.messageManager = new MessageManager(this);
        }

        public connect(url : string): ng.IPromise<{}> {
            return this.messageManager.connect(url);
        }

        public method(methodName: string, parameters: any[]): ng.IPromise<{}> {
            // var message = new Message(
            //     new MethodHandler(methodName, parameters)
            // );
            // return message.send();
            return null;
        }

        public onConnect(): ng.IPromise<{}> {
            return this.connected.promise;
        }

    }

}