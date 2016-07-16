module goTransport {

    export class MessageMethodResult extends Message{
        static type = MessageType.MessageTypeMethodResult;

        constructor(private result: boolean = false, private parameters: Array<any> = null) {
            super(MessageMethod.type);
        }

        validate(): Error {
            console.log('validating method result', this.reply);
            return null;
        }

        run(): Error {
            console.log('Running method result', this.reply);
            return null;
        }

    }

    MessageDefinition.set(MessageMethod.type, MessageMethod);
}