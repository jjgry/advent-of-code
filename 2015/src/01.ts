import { getFileLines } from "./utils/input";

const lines = getFileLines("01");

export const day01part1 = () => {
  const input = lines[0];

  let floor = 0;
  for (const char of input) {
    if (char === "(") {
      floor++;
    }
    if (char === ")") {
      floor--;
    }
  }
  return floor;
};
