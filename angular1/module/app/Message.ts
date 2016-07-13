module goTransport {

    export enum MessageType {
        MessageTypeMethod,
        MessageTypeMethodResult,
        MessageTypePub
    }

    export class Message implements IMessage{
        type:MessageType;
        id:number;
        static current_id = 0;
        data:any;
        private static promises: ng.IDeferred<{}>[];

        constructor(messageHandler : MessageHandler) {
            this.type = messageHandler.getMessageType();
            this.id = Message.current_id++;
            this.data = messageHandler.toJSON();

            if(!Message.promises) {
                Message.promises = [];
            }
        }

        //Sends a to the SockJS server
        send(timeout: number=3000): ng.IPromise<{}> {
            Message.promises[this.id] = GoTransport.instance.$q.defer();

            // Log_Print("Dispatching RPC message with ID %d and type %d.\n", g_rpc.messageID, type);
            GoTransport.instance.connectedPromise.promise.then(function() {
                GoTransport.instance.socket.send(JSON.stringify(this.toJSON()));

                GoTransport.instance.$timeout(function () {

                    if(Message.promises[this.id].promise.$$state.status == 0) {//Pending
                        console.log("Timed out");
                        Message.promises[this.id].reject("Timed out"); //reject the service in case of timeout
                    }

                }.bind(this), timeout);

            }.bind(this));
            return Message.promises[this.id].promise;
        }

        // receive(): {
        //
        // }

        // toJSON is automatically used by JSON.stringify
        toJSON(): IMessage {
            // copy all fields from `this` to an empty object and return in
            return Object.assign({}, this, {
                // convert fields that need converting
            });
        }

        // fromJSON is used to convert an serialized version
        // of the Message to an instance of the class
        static fromJSON(json: IMessage|string): Message {
            if (typeof json === 'string') {
                // if it's a string, parse it first
                return JSON.parse(json, Message.reviver);
            } else {
                // create an instance of the Message class
                let message = Object.create(Message.prototype);
                // copy all the fields from the json object
                return Object.assign(message, json, {
                    // convert fields that need converting

                });
            }
        }

        // reviver can be passed as the second parameter to JSON.parse
        // to automatically call Message.fromJSON on the resulting value.
        static reviver(key: string, value: any): any {
            return key === "" ? Message.fromJSON(value) : value;
        }
    }

    // A representation of Message's data that can be converted to
    // and from JSON without being altered.
    interface IMessage {
        type:   MessageType;//TODO MessageType enum?
        id:     number;
        data:   any;
    }

}