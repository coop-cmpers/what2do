import { HelloWorldService } from "../protos-gen/helloworld/v1/service_connect";
import { What2DoService } from "../protos-gen/what2do/v1/service_connect";
import { setupPromiseConnectClient } from "./connect-web-setup";

export const helloWorldClient = setupPromiseConnectClient(HelloWorldService);

export const what2DoClient = setupPromiseConnectClient(What2DoService);