import { getFileLines } from "./utils/input";

const input = getFileLines("05");

export const day05part1 = () => {
  return input.filter(isNicePart1).length;
};

export const day05part2 = () => {
  return input.filter(isNicePart2).length;
};

const isNicePart1 = (string: string): boolean => {
  const hasThreeVowels = (string.match(/[aeiou]/g)?.length || 0) >= 3;
  const hasDoubleLetter = /([a-z])\1/.test(string);
  const hasForbiddenPairs = /(ab)|(cd)|(pq)|(xy)/.test(string);
  return hasThreeVowels && hasDoubleLetter && !hasForbiddenPairs;
};

const isNicePart2 = (string: string): boolean => {
  const hasPairThatRepeats = /([a-z][a-z])[a-z]*\1/.test(string);
  const hasOneLetterThatRepeatsAroundAnother = /([a-z])[a-z]\1/.test(string);
  return hasPairThatRepeats && hasOneLetterThatRepeatsAroundAnother;
};
