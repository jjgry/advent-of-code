const { log } = require("console");
const fs = require("fs");

let data;
try {
  data = fs.readFileSync("input2.txt", "utf8");
} catch (err) {
  console.error(err);
}

const lines = data.split("\n");

let height = 0;
let distance = 0;

console.log(lines.slice(-10));

lines.forEach((line) => {
  if (line === "") {
    return;
  }
  const [direction, magnitude] = line.split(" ");
  const magNum = parseInt(magnitude);
  console.log(direction, magNum);

  switch (direction) {
    case "forward":
      console.log("forward");
      distance += magNum;
      break;
    case "up":
      console.log("up");
      height -= magNum;
      break;
    case "down":
      console.log("down");
      height += magNum;
      break;
    default:
      console.log("probably shouldn't be here");
  }
});

console.log("height", height);
console.log("distance", distance);
console.log("height * distance", height * distance);
