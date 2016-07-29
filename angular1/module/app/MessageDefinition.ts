module goTransport {
    export enum MessageType {
        MessageTypeMethod,
        MessageTypeMethodResult,
        MessageTypeError,
        MessageTypePub
    }
    
    export class MessageDefinition{
        private static definitions = Array<any>();

        public static set(type : MessageType, definition : any) {
            if(!definition || !definition.prototype) {
                console.warn("Invalid message definition set for type", type);
                return
            }
            MessageDefinition.definitions[type] = definition;
        }

        public static get(type : MessageType, data : string) : Message {
            var definition = MessageDefinition.definitions[type];
            if(definition === undefined) {
                console.warn("Invalid messageType requested", type);
                return null;
            }

            let messageBuilder = new MessageBuilder<definition>(definition);
            var message = messageBuilder.build();
            Object.assign(message, JSON.parse(data), {
                // convert fields that need converting
            });
            return message;

        }

    }

}