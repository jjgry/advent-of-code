import { getFileLines } from "./utils/input";

const input = getFileLines("05");

export const day05part1 = () => {
  return input.filter(isNicePart1).length;
};

export const day05part2 = () => {
  return input.filter(isNicePart2).length;
};

const isNicePart1 = (string: string): boolean => {
  const hasThreeVowels = /([aeiou]).*([aeiou]).*([aeiou])/.test(string);
  const hasDoubleLetter = /(.)\1/.test(string);
  const hasForbiddenPairs = /(ab)|(cd)|(pq)|(xy)/.test(string);
  return hasThreeVowels && hasDoubleLetter && !hasForbiddenPairs;
};

const isNicePart2 = (string: string): boolean => {
  const hasPairThatRepeats = /(..).*\1/.test(string);
  const hasOneLetterThatRepeatsAroundAnother = /(.).\1/.test(string);
  return hasPairThatRepeats && hasOneLetterThatRepeatsAroundAnother;
};
