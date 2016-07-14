module goTransport {
    export class MessageBuilder<T extends Message> {
        constructor(private testType:any) {
            
        }

        build() : T {
            return new this.testType();
        }
    }
}