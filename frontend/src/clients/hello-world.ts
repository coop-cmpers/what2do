import { helloWorldClient } from ".";
import { HelloWorldRequest } from "../protos-gen/helloworld/v1/helloworld_pb"

export const helloWorld = () => {
    const req = new HelloWorldRequest();
    return helloWorldClient.helloWorld(req);
}