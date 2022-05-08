import { getFileLines } from "./utils/input";

const input = getFileLines("01")[0];

export const day01part1 = () => {
  let floor = 0;
  for (const char of input) {
    if (char === "(") floor++;
    if (char === ")") floor--;
  }
  return floor;
};

export const day01part2 = () => {
  let floor = 0;
  for (let i = 0; i < input.length; i++) {
    if (input[i] === "(") floor++;
    if (input[i] === ")") floor--;
    if (floor < 0) return i + 1;
  }
  return -1;
};
