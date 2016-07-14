module goTransport {
    import SocketDelegate = Socket.SocketDelegate;

    export class MessageManager implements SocketDelegate{
        public socket: Socket.Socket;

        constructor(private client: Client) {

        }
        
        public connect(url : string): ng.IPromise<{}> {
            if(this.socket == null) {
                this.socket = Socket.Adapter.getSocket("SockJSClient", url, this);
            }
            return this.client.connected.promise;
        }

        public message(data : string) {
            console.log('send', data);
            this.socket.send(data);
        }

        connected() {
            this.client.connected.resolve()
        }

        messaged(data:string) {
            console.log('received', data);
        }

        disconnected(code:number, reason:string, wasClean:boolean) {
            console.log('Disconnected')
        }

    }

}