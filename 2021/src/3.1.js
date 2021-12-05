const fs = require("fs");
const { listenerCount } = require("process");

let data;
try {
  data = fs.readFileSync("input3.txt", "utf8");
} catch (err) {
  console.error(err);
}

const lines = data.split("\r\n");

const linelength = lines[0].length;

const getNum = (more) => {
  let newLines = lines.slice();

  for (let i = 0; i < linelength; i++) {
    if (newLines.length <= 1) {
      break;
    }

    let numOnes = 0;
    let numZeroes = 0;
    newLines.forEach((line) => {
      const bit = line[i];
      if (bit === "0") {
        numZeroes++;
      } else if (bit === "1") {
        numOnes++;
      } else {
        console.log("not meant to happen");
      }
    });

    console.log(numZeroes, numOnes);

    if (numOnes > numZeroes) {
      if (more) {
        newLines = newLines.filter((line) => line[i] === "1");
      } else {
        newLines = newLines.filter((line) => line[i] === "0");
      }
    } else if (numOnes < numZeroes) {
      if (more) {
        newLines = newLines.filter((line) => line[i] === "0");
      } else {
        newLines = newLines.filter((line) => line[i] === "1");
      }
    } else {
      if (more) {
        newLines = newLines.filter((line) => line[i] === "1");
      } else {
        newLines = newLines.filter((line) => line[i] === "0");
      }
    }

    console.log(newLines);
  }

  // console.log(newLines);

  const target = newLines[0];

  return parseInt(target, 2);
};

const oxy = getNum(true);
const co2 = getNum(false);

console.log(oxy);
console.log(co2);
console.log(oxy * co2);
