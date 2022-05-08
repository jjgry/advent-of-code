import { getFileLines } from "./utils/input";

const input = getFileLines("03");

const directions = input[0].split("") as ("^" | ">" | "<" | "v")[];

type XYGrid = {
  [x: number]: {
    [y: number]: number;
  };
};

export const day03part1 = () => {
  const counts: XYGrid = { "0": { "0": 1 } };
  const position = { x: 0, y: 0 };

  for (const direction of directions) {
    if (direction === "^") position.y++;
    if (direction === "v") position.y--;
    if (direction === ">") position.x++;
    if (direction === "<") position.x--;

    if (counts[position.x] === undefined) counts[position.x] = {};
    const existingCount = counts[position.x]?.[position.y];
    counts[position.x][position.y] = existingCount ? existingCount + 1 : 1;
  }

  return Object.values(counts).reduce((subtotal, row) => {
    return subtotal + Object.values(row).length;
  }, 0);
};

export const day03part2 = () => {
  const counts: XYGrid = { "0": { "0": 1 } };
  const santaPosition = { x: 0, y: 0 };
  const robotPosition = { x: 0, y: 0 };
  let isRoboSanta = false;

  for (const direction of directions) {
    const position = isRoboSanta ? robotPosition : santaPosition;

    if (direction === "^") position.y++;
    if (direction === "v") position.y--;
    if (direction === ">") position.x++;
    if (direction === "<") position.x--;

    if (counts[position.x] === undefined) counts[position.x] = {};
    const existingCount = counts[position.x]?.[position.y];
    counts[position.x][position.y] = existingCount ? existingCount + 1 : 1;

    isRoboSanta = !isRoboSanta;
  }

  return Object.values(counts).reduce((subtotal, row) => {
    return subtotal + Object.values(row).length;
  }, 0);
};
