module goTransport {
    export enum MessageType {
        MessageTypeMethod,
        MessageTypeMethodResult,
        MessageTypePub
    }

    export class MessageDefinitions {

        private static definitions = Array<any>();

        public static set(type : MessageType, definition : any) {
            if(!definition || !definition.prototype) {
                console.warn("Invalid message definition set for type", type);
                return
            }
            MessageDefinitions.definitions[type] = definition;
        }

        public static get(type : MessageType, data : string) : Message {
            var definition = MessageDefinitions.definitions[type];
            if(definition === undefined) {
                console.warn("Invalid messageType requested", type);
                return null;
            }

            // // create an instance of the Message class
            // let message = Object.create(definition.prototype);
            // // copy all the fields from the json object
            // return Object.assign(message, JSON.parse(data), {
            //     // convert fields that need converting
            // });

            let messageBuilder = new MessageBuilder<definition>(definition);
            var message = messageBuilder.build();
            Object.assign(message, JSON.parse(data), {
                // convert fields that need converting
            });
            return message;

        }

    }

}