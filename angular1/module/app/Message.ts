module goTransport {

    export abstract class Message{
        id:number;
        static current_id = 0;
        protected reply: Message;
        private static headerDelimiter = "\f";

        constructor(private type : MessageType) {
            this.id = null;
        }

        public getType():MessageType {
            return this.type;
        }

        public setReply(message: Message) {
            this.reply = message;
        }

        abstract validate(): Error

        abstract run(): Error

        public start(): Error {
            var error = this.validate();
            if(error) {
                return error;
            }
            error = this.run();
            if(error) {
                return error;
            }
            return null;
        }

        // toJSON is automatically used by JSON.stringify
        serialize(): string {
            // copy all fields from `this` to an empty object and return in
            return this.type + Message.headerDelimiter + JSON.stringify(this.encode());
        }

        static unSerialize(data : string):Message {
            var parts = data.split(Message.headerDelimiter);
            if(parts[1] === undefined) {
                console.warn("Invalid message. Invalid amount of parts", data);
                return null;
            }
            console.log(parts);
            return MessageDefinition.get(parseInt(parts[0]), parts[1]);
        }

        public encode():any {
            return Object.assign({}, this, {
                // convert fields that need converting
            });
        }

    }
}