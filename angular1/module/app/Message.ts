module goTransport {

    export abstract class Message{
        id:number;
        static current_id = 0;
        protected reply: Message;
        private static headerDelimiter = "\f";

        constructor(private type : MessageType) {
            this.id = Message.current_id++;
        }

        public getType():MessageType {
            return this.type;
        }

        public setReply(message: Message) {
            this.reply = message;
        }

        abstract validate(): Error

        abstract run(): Error

        public start(): boolean {
            var error = this.validate();
            if(error) {
                console.error(error);
                return false;
            }
            error = this.run();
            if(error) {
                console.error(error);
                return false;
            }
            return true;
        }

        // toJSON is automatically used by JSON.stringify
        serialize(): string {
            // copy all fields from `this` to an empty object and return in
            return this.type + Message.headerDelimiter + JSON.stringify(this.encode());
        }

        static unSerialize(data : string):Message {
            var parts = data.split(Message.headerDelimiter);
            if(parts[1] === undefined) {
                console.warn("Invalid message", data);
                return null;
            }
            return MessageDefinition.get(parseInt(parts[0]), JSON.parse(parts[1]));
        }

        public encode():any {
            return Object.assign({}, this, {
                // convert fields that need converting
            });
        }

    }
}