const fs = require("fs");

let data;
try {
  data = fs.readFileSync("input.txt", "utf8");
} catch (err) {
  console.error(err);
}

const lines = data.split("\n");

let numberOfIncreases = 0;
let prevValue = parseInt(lines[0]) + parseInt(lines[1]) + parseInt(lines[2]);

for (let i = 1; i < lines.length - 2; i++) {
  const newVal =
    parseInt(lines[i]) + parseInt(lines[i + 1]) + parseInt(lines[i + 2]);
  if (newVal > prevValue) {
    numberOfIncreases++;
  }
  prevValue = newVal;
}

console.log(numberOfIncreases);
