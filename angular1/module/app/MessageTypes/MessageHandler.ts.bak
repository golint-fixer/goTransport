module goTransport {

    export abstract class MessageHandler {

        constructor(private messageType: MessageType) {

        }

        getMessageType():MessageType {
            return this.messageType;
        }

        abstract validate(message : Message) : Error;
        abstract run(message : Message) : Error;

        // toJSON is automatically used by JSON.stringify
        abstract toJSON(): any

    }

}