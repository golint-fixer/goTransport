module goTransport {

    export class MessageMethodResult extends Message{
        static type = MessageType.MessageTypeMethodResult;

        constructor(private result: boolean = false, private parameters: Array<any> = null) {
            super(MessageMethodResult.type);
        }

        validate(): Error {
            if(!(this.reply instanceof MessageMethod)) {
                return new Error("Invalid reply. Not messageMethod.");
            }
        }

        run(): Error {
            console.log('Result came back!', this.parameters);
            let promise = (this.reply as MessageMethod).getPromise();
            if(promise) {
                promise.resolve.apply(promise, this.parameters);
            }

            return null;
        }

    }

    MessageDefinition.set(MessageMethodResult.type, MessageMethodResult);
}