import md5 from "md5";
import { getFileLines } from "./utils/input";

const input = getFileLines("04")[0];

export const day04part1 = () => {
  for (let i = 1; ; i++) {
    if (md5(`${input}${i}`).substring(0, 5) === "00000") return i;
  }
};

export const day04part2 = () => {
  for (let i = 1; ; i++) {
    if (md5(`${input}${i}`).substring(0, 6) === "000000") return i;
  }
};
