// @generated by protoc-gen-es v1.9.0 with parameter "target=js+dts"
// @generated from file what2do/v1/recommendations.proto (package what2do.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from enum what2do.v1.EventType
 */
export const EventType = /*@__PURE__*/ proto3.makeEnum(
  "what2do.v1.EventType",
  [
    {no: 0, name: "EVENT_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "EVENT_TYPE_RESTAURANT", localName: "RESTAURANT"},
    {no: 2, name: "EVENT_TYPE_ACTIVITY", localName: "ACTIVITY"},
  ],
);

/**
 * @generated from message what2do.v1.Recommendation
 */
export const Recommendation = /*@__PURE__*/ proto3.makeMessageType(
  "what2do.v1.Recommendation",
  () => [
    { no: 1, name: "rank", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "address", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "PriceLevel", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "Open", kind: "message", T: Timestamp },
    { no: 6, name: "Close", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message what2do.v1.SearchRecommendationsRequest
 */
export const SearchRecommendationsRequest = /*@__PURE__*/ proto3.makeMessageType(
  "what2do.v1.SearchRecommendationsRequest",
  () => [
    { no: 1, name: "event_time", kind: "message", T: Timestamp },
    { no: 2, name: "location", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "event_type", kind: "enum", T: proto3.getEnumType(EventType) },
    { no: 4, name: "search_param", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message what2do.v1.SearchRecommendationsResponse
 */
export const SearchRecommendationsResponse = /*@__PURE__*/ proto3.makeMessageType(
  "what2do.v1.SearchRecommendationsResponse",
  () => [
    { no: 1, name: "recommendations", kind: "message", T: Recommendation, repeated: true },
  ],
);

