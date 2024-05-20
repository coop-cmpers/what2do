import { helloWorldClient } from ".";
import { HelloBackendRequest } from "../protos-gen/helloworld/v1/hello-backend_pb";

export const helloBackend = () => {
    const req = new HelloBackendRequest();
    return helloWorldClient.helloBackend(req);
}