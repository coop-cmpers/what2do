import { PromiseClient, Transport, createPromiseClient } from "@bufbuild/connect";
import { createGrpcWebTransport } from "@bufbuild/connect-web";
import { ServiceType } from "@bufbuild/protobuf";
import { ENVOY_URL } from "../constants";

const options = {
  baseUrl: ENVOY_URL,
};

export function setupPromiseConnectClient<T extends ServiceType>(
  service: T
): PromiseClient<T> {
  const transport: Transport = createGrpcWebTransport(options);
  return createPromiseClient(service, transport);
}