module goTransport {

    export abstract class Message{
        id:number;
        static current_id = 0;

        constructor(private type : MessageType) {
            this.id = Message.current_id++;
        }

        abstract validate(): Error

        abstract run(): Error

        //
        // //Sends a to the SockJS server
        // send(timeout: number=3000): ng.IPromise<{}> {
        //     Message.promises[this.id] = Client.instance.$q.defer();
        //
        //     // Log_Print("Dispatching RPC message with ID %d and type %d.\n", g_rpc.messageID, type);
        //     Client.instance.connectedPromise.promise.then(function() {
        //         Client.instance.socket.send(JSON.stringify(this.toJSON()));
        //
        //         Client.instance.$timeout(function () {
        //
        //             if(Message.promises[this.id].promise.$$state.status == 0) {//Pending
        //                 console.log("Timed out");
        //                 Message.promises[this.id].reject("Timed out"); //reject the service in case of timeout
        //             }
        //
        //         }.bind(this), timeout);
        //
        //     }.bind(this));
        //     return Message.promises[this.id].promise;
        // }
        //
        // // receive(): {
        // //
        // // }
        //
        // toJSON is automatically used by JSON.stringify
        // toJSON(): Message {
        //     // copy all fields from `this` to an empty object and return in
        //     return Object.assign({}, this, {
        //         // convert fields that need converting
        //     });
        // }
        //
        // // fromJSON is used to convert an serialized version
        // // of the Message to an instance of the class
        // fromJSON(json: Message|string): Message {
        //     if (typeof json === 'string') {
        //         // if it's a string, parse it first
        //         return JSON.parse(json, Message.reviver);
        //     } else {
        //         // create an instance of the Message class
        //         let message = Object.create(Message.prototype);
        //         // copy all the fields from the json object
        //         return Object.assign(message, json, {
        //             // convert fields that need converting
        //
        //         });
        //     }
        // }
        //
        // // reviver can be passed as the second parameter to JSON.parse
        // // to automatically call Message.fromJSON on the resulting value.
        // static reviver(key: string, value: any): any {
        //     return key === "" ? Message.fromJSON(value) : value;
        // }
    }
    
}