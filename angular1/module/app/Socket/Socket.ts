module Socket {

    export interface SocketDelegate {
        connected(): any;
        message(data: string): any;
        disconnected(code: number, reason: string, wasClean: boolean): any;
    }

    export interface Socket {
        send(data:string):any;
        close():any;
    }

    export class Adapter {
        static getSocket(type: string, url: string, delegate: Socket.SocketDelegate):Socket.Socket {
            switch(type) {

                case "SockJSClient":
                    return SockJSClient.getInstance(url, delegate);
                default:
                    throw("Invalid socket type:"+type)
            }
        }
    }

}