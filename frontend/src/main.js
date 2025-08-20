import "./style.css";
import "./app.css";

import { Greet } from "../wailsjs/go/main/App";

// Setup the getQuote function
window.getQuote = function () {
  // Call App.GetRandomQuote()
  import("../wailsjs/go/main/App").then(({ GetRandomQuote }) => {
    GetRandomQuote()
      .then((quote) => {
        resultElement.innerText = quote.text || quote;
      })
      .catch((err) => {
        resultElement.innerText = "Could not fetch quote.";
        console.error(err);
      });
  });
};

// Setup the greet function
window.greet = function () {
  let name = nameElement.value.trim();
  if (name === "") {
    resultElement.innerText = "Please enter your name!";
    return;
  }
  Greet(name)
    .then((result) => {
      resultElement.innerText = result;
    })
    .catch((err) => {
      resultElement.innerText = "Could not greet.";
      console.error(err);
    });
};
