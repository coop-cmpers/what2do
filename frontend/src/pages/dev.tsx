import React, { useState } from "react";
import { searchRecommendations } from "../clients/recommendations";
import { Recommendation } from "../protos-gen/what2do/v1/recommendations_pb";
import { Timestamp } from "@bufbuild/protobuf";

const DevPage = () => {
  const [recommendations, setRecommendations] = useState<Recommendation[]>([]);
  const [searchParam, setSearchParam] = useState<string>("");
  const [location, setLocation] = useState<string>("");
  const handleSearchButton = () => {
    searchRecommendations({searchParam: searchParam, location: location, eventTime: Timestamp.now()}).then((resp) => setRecommendations(resp.recommendations));
  }

  return (
    <>
      <input type="text" onChange={(e) => setSearchParam(e.target.value)}></input>
      <input type="text" onChange={(e) => setLocation(e.target.value)}></input>
      <button onClick={() => handleSearchButton()}>Search</button>
      {recommendations.map((recommendation) => (
        <div>
        <h3>{recommendation.name}</h3>
        <p>{recommendation.address}</p>
        <p>{recommendation.PriceLevel}</p>
        <p>{recommendation.Open?.toJsonString()}</p>
        </div>
      ))}
    </>
  );
}

export default DevPage;