import { HelloWorldService } from "../protos-gen/helloworld/v1/helloworld_connect";
import { setupPromiseConnectClient } from "./connect-web-setup";

export const helloWorldClient = setupPromiseConnectClient(HelloWorldService);