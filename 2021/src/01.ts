import { getFileLines } from "./utils/input";

const lines = getFileLines("01");

const getWindowSum = (base: number, windowSize: number) => {
  let sum = 0;
  for (let offset = 0; offset < windowSize; offset++) {
    sum += parseInt(lines[base + offset]);
  }
  return sum;
};

const getSlidingWindowIncreases = (windowSize: number) => {
  let numberOfIncreases = 0;
  let prevValue = getWindowSum(0, windowSize);

  const numWindows = lines.length - (windowSize - 1);
  for (let base = 1; base < numWindows; base++) {
    const windowSum = getWindowSum(base, windowSize);
    if (windowSum > prevValue) {
      numberOfIncreases++;
    }
    prevValue = windowSum;
  }

  return numberOfIncreases;
};

export const day01part1 = () => getSlidingWindowIncreases(1);
export const day01part2 = () => getSlidingWindowIncreases(3);
