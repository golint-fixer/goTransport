module goTransport {

    // interface MessageMethodJson {
    //     name: string;
    //     parameters:Array<any>;
    // }

    export class MessageMethod extends Message implements MessageMethodJson{
        static type = MessageType.MessageTypeMethod;
        private promise : Promise;

        constructor(public name: string = null, public parameters: Array<any> = null) {
            super(MessageMethod.type);
        }

        validate(): Error {
            return null;
        }

        run(): Error {
            console.log('ran');
            this.promise = new Promise();
            return null;
        }

        public getPromise(): Promise {
            return this.promise;
        }

        // public encode():MessageMethodJson {
        //     let value = Object.assign({}, this, {
        //         // convert fields that need converting
        //     });
        //     console.log('value', value);
        //     return value;
        // }

    }

    MessageDefinition.set(MessageMethod.type, MessageMethod);
}