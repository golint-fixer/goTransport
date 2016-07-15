module goTransport {
    import SocketDelegate = Socket.SocketDelegate;

    export class MessageManager implements SocketDelegate{
        public socket: Socket.Socket;
        private messages : Array<Message>;

        constructor(private client: Client) {
           this.messages = [];
        }

        public connect(url : string): IPromise<{}> {
            if(this.socket == null) {
                this.socket = Socket.Adapter.getSocket("SockJSClient", url, this);
            }
            return this.client.connected.promise;
        }

        private set(message: Message) {
            this.messages[message.id] = message;
        }

        private get(message: Message): Message {
            return this.messages[message.id];
        }

        connected() {
            this.client.connected.resolve()
        }

        //Send
        public send(message : Message) {
            //Set it
            this.set(message);

            message.start();

            //Send it
            this.socket.send(message.serialize());
            console.log('sent', message.serialize());
        }

        //Receive
        messaged(data:string) {
            let message = Message.unSerialize(data);
            if(!message) {
                console.warn("Invalid message received.");
                return;
            }

            message.setReply(this.get(message));
            message.validate();
            message.run();
        }

        disconnected(code:number, reason:string, wasClean:boolean) {
            console.warn('Disconnected', code)
        }

    }

}