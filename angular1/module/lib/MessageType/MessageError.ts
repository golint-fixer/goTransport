module goTransport {

    export class MessageError extends Message{
        static type = MessageType.MessageTypeError;
        private promise : Promise;

        constructor(public reason: any) {
            super(MessageError.type);
        }

        validate(): Error {
            return null;
        }

        run(): Error {
            console.error(this.reason);

            //On error
            if((this.reply instanceof MessageMethod)) {
                let promise = (this.reply as MessageMethod).getPromise();
                if(promise) {
                    console.debug(this);
                    promise.reject(this.reason);
                }
            }
            return null;
        }
    }

    MessageDefinition.set(MessageError.type, MessageError);
}