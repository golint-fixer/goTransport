module goTransport {

    export class MethodHandler extends MessageHandler implements IMethodHandler{

        constructor(private name: string, private parameters: any[]) {
            super(MessageType.MessageTypeMethod)
        }

        validate(): Error {
            return null;
        }

        run(): Error {
            return null;
        }

        toJSON(): IMethodHandler {
            // copy all fields from `this` to an empty object and return in
            return Object.assign({}, this, {
                // convert fields that need converting
            });
        }
        
    }

    interface IMethodHandler {
        name:   string;
        parameters: any[];
    }

}