module goTransport {

    export class MessageMethod extends Message{
        static type = MessageType.MessageTypeMethod;
        private name: string;
        private parameters: Array<any>;

        constructor() {
            super(MessageMethod.type);
        }

        validate(): Error {
            return null;
        }

        run(): Error {
            console.log('ran');
            return null;
        }

    }

    MessageDefinitions.set(MessageMethod.type, MessageMethod);
    //
    var message = MessageDefinitions.get(MessageType.MessageTypeMethod, '{"name": "test"}');
    console.log(message);

    message = MessageDefinitions.get(MessageType.MessageTypeMethod, '{"name": "test"}');
    console.log(message);
    // console.log('id', message.id);
    //
    //
    // class MessageBuilder<T extends Message> {
    //     constructor(private testType:any) {
    //         console.log('breakin my balls')
    //     }
    //
    //     getNew() : T {
    //         return new this.testType();
    //     }
    // }
    //
    // var test = new MessageBuilder<MessageMethod>(MessageMethod);
    //
    // var example = test.getNew();
    // example.run();
    // Object.assign(example, {name: "bever", parameters: ["heel smerig"]}, {//JSON.parse(data)
    //     // convert fields that need converting
    // });
    // console.log(example);

}