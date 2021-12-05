const { log } = require("console");
const fs = require("fs");

let data;
try {
  data = fs.readFileSync("input3.txt", "utf8");
} catch (err) {
  console.error(err);
}

const lines = data.split("\n");

const linelength = lines[0].length - 1;

let gammaArray = "";
let epsillonArray = "";

for (let i = 0; i < linelength; i++) {
  let numOnes = 0;
  let numZeroes = 0;
  lines.forEach((line) => {
    const bit = line[i];
    if (bit === "0") {
      numZeroes++;
    } else if (bit === "1") {
      numOnes++;
    } else {
      console.log("not meant to happen");
    }
  });

  if (numOnes > numZeroes) {
    gammaArray += "1";
    epsillonArray += "0";
  } else if (numOnes < numZeroes) {
    gammaArray += "0";
    epsillonArray += "1";
  } else {
    console.log("shouldn't happen either");
  }
}

const gamma = parseInt(gammaArray, 2);
const epsillon = parseInt(epsillonArray, 2);

console.log(gamma);
console.log(epsillon);
console.log(gamma * epsillon);
