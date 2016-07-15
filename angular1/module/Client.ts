module goTransport {

    export abstract class Client {

        public connected: ng.IDeferred<{}>;
        protected static instance: Client;
        private messageManager: MessageManager;

        constructor() {
            Client.instance = this;
            this.connected = new Promise();
            this.messageManager = new MessageManager(this);
        }

        public connect(url : string): IPromise<{}> {
            return this.messageManager.connect(url);
        }

        public method(name: string, parameters: any[], timeout: number = 3000): IPromise<{}> {
            let message = new MessageMethod(name, parameters);
            this.messageManager.send(message);
            var promise = message.getPromise();
            promise.setTimeOut(timeout);
            return promise.promise;
        }

        public onConnect(): IPromise<{}> {
            return this.connected.promise;
        }

    }

}