import "./style.css";
import "./app.css";

import { GetRandomQuote, Greet } from "../wailsjs/go/main/App";
// import { EventsOn } from "../wailsjs/runtime/runtime";

const quoteEl = document.getElementById("qoute");

window.getQuote = function () {
  import("../wailsjs/go/main/App").then(({ GetRandomQuote }) => {
    GetRandomQuote()
      .then((quote) => {
        quoteEl.innerText = quote.text || quote;
      })
      .catch((err) => {
        quoteEl.innerText = "Could not fetch quote.";
        console.error(err);
      });
  });
};

setInterval(() => {
    GetRandomQuote()
      .then((quote) => {
        console.log(quote);
        quoteEl.innerText = quote.text || quote;
      })
      .catch((err) => {
        quoteEl.innerText = "Could not fetch quote.";
        console.error(err);
      });
}, 300000);

async function displayQuote(){
  try {
    const quote = await GetRandomQuote();
    quoteEl.innerText = quote.text || quote;
  } catch (error) {
    quoteEl.innerText = "Could not fetch quote.";
    console.error(error);
  }
}

// Setup the greet function
window.greet = function () {
  let name = nameElement.value.trim();
  if (name === "") {
    quoteEl.innerText = "Please enter your name!";
    return;
  }
  Greet(name)
    .then((result) => {
      quoteEl.innerText = result;
    })
    .catch((err) => {
      quoteEl.innerText = "Could not greet.";
      console.error(err);
    });
};
