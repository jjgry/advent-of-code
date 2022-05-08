import { day01part1, day01part2 } from "./01";
import { day02part1, day02part2 } from "./02";
import { day03part1, day03part2 } from "./03";
// import { day04part1, day04part2 } from "./04";
import { day05part1, day05part2 } from "./05";

const main = () => {
  console.log(`1.1: ${day01part1()}`);
  console.log(`1.2: ${day01part2()}`);
  console.log(`2.1: ${day02part1()}`);
  console.log(`2.2: ${day02part2()}`);
  console.log(`3.1: ${day03part1()}`);
  console.log(`3.2: ${day03part2()}`);
  console.log(`4.1: Skipping`); //${day04part1()}`);
  console.log(`4.2: Skipping`); //${day04part2()}`);
  console.log(`5.1: ${day05part1()}`);
  console.log(`5.2: ${day05part2()}`);
};

main();
