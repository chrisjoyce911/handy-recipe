import React from "react";

import { useParams } from "react-router-dom";
import data from "../recipe.json";

export default function RecipePage() {
  let { id } = useParams();
  var r = data.Recipes[id];
  return (
    <>
      <h1>{r.name}</h1>

      <p>{r.description}</p>
    </>
  );
}