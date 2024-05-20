import { Timestamp } from "@bufbuild/protobuf";
import { what2DoClient } from ".";
import { SearchRecommendationsRequest } from "../protos-gen/what2do/v1/recommendations_pb";

type SearchRecommendationsParams = {
    searchParam: string;
    location: string;
    eventTime: Timestamp
}

export const searchRecommendations = ({searchParam, location, eventTime}: SearchRecommendationsParams) => {
    const req = new SearchRecommendationsRequest({searchParam, location, eventTime});
    return what2DoClient.searchRecommendations(req);
}