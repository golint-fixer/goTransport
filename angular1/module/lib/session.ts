module goTransport {
    import SocketDelegate = Socket.SocketDelegate;

    export class Session implements SocketDelegate{
        public socket: Socket.Socket;
        private messages : Array<Message>;
        private connectedPromise : Promise;

        constructor() {
            this.messages = [];
        }

        public connect(url : string): IPromise<{}> {
            this.connectedPromise = new Promise();
            if(this.socket == null) {
                this.socket = Socket.Adapter.getSocket("SockJSClient", url, this);
            }
            return this.getConnectedPromise();
        }

        public getConnectedPromise():IPromise<{}> {
            return this.connectedPromise.promise;
        }

        private set(message: Message) {
            console.log('messageManager', 'set', message.id);
            this.messages[message.id] = message;
        }

        private get(message: Message): Message {
            console.log('messageManager', 'get', message.id, this.messages[message.id]);
            return this.messages[message.id];
        }

        connected() {
            console.log('connected');
            this.connectedPromise.resolve();
        }

        //Send
        public send(message : Message) {
            Message.current_id++;
            message.id = Message.current_id;

            message.start();
            this.set(message);

            //Send it
            this.getConnectedPromise().then(function() {
                this.socket.send(message.serialize());
                console.log('sent', message.serialize());
            }.bind(this));
        }

        //Receive
        messaged(data:string) {
            console.debug('received', data);

            let message = Message.unSerialize(data);
            if(!message) {
                console.warn("Invalid message received.");
                return;
            }

            message.setReply(this.get(message));
            var error = message.start();
            if (error != null) {
                console.error(error);
            }
            
        }

        disconnected(code:number, reason:string, wasClean:boolean) {
            console.warn('Disconnected', code)
        }

    }

}