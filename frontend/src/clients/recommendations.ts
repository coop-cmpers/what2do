import { what2DoClient } from ".";
import { SearchRecommendationsRequest } from "../protos-gen/what2do/v1/recommendations_pb";


export const searchRecommendations = () => {
    const req = new SearchRecommendationsRequest();
    return what2DoClient.searchRecommendations(req);
}