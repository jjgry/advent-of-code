const fs = require("fs");

let data;
try {
  data = fs.readFileSync("input.txt", "utf8");
} catch (err) {
  console.error(err);
}

const lines = data.split("\n");

let numberOfIncreases = 0;
let prevValue = parseInt(lines[0]);

for (let i = 1; i < lines.length; i++) {
  const newVal = parseInt(lines[i]);
  if (newVal > prevValue) {
    numberOfIncreases++;
  }
  prevValue = newVal;
}

console.log(numberOfIncreases);
