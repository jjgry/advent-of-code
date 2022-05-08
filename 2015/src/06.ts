import { getFileLines } from "./utils/input";

const input = getFileLines("06");

export const day06part1 = () => {
  const grid = newGrid(false);
  for (const { command, start, end } of instructions) {
    for (const { x, y } of getIntermediateCoords(start, end)) {
      grid[x][y] = getNewValuePart1(grid[x][y], command);
    }
  }
  const numTurnedOn = grid.reduce((accumulator, gridRow) => {
    return accumulator + gridRow.filter((x) => x).length;
  }, 0);
  return numTurnedOn;
};

export const day06part2 = () => {
  const grid = newGrid(0);
  for (const { command, start, end } of instructions) {
    for (const { x, y } of getIntermediateCoords(start, end)) {
      grid[x][y] = getNewValuePart2(grid[x][y], command);
    }
  }
  const totalBrightness = grid.reduce((accumulator, gridRow) => {
    return accumulator + gridRow.reduce((a, b) => a + b, 0);
  }, 0);
  return totalBrightness;
};

type Coordinate = { x: number; y: number };

type Instruction = {
  command: "turnOn" | "turnOff" | "toggle";
  start: Coordinate;
  end: Coordinate;
};

const newGrid = <T>(fillValue: T): T[][] =>
  new Array(1000).fill(fillValue).map(() => new Array(1000).fill(fillValue));

const parseCoord = (stringCoord: string): Coordinate => {
  const coord = stringCoord.split(",").map((numStr) => parseInt(numStr));
  return {
    x: coord[0],
    y: coord[1],
  };
};

const instructions: Instruction[] = input.map((instructionString) => {
  const words = instructionString.split(" ");
  const isToggle = words.length === 4;
  return {
    command: isToggle ? "toggle" : words[1] === "on" ? "turnOn" : "turnOff",
    start: parseCoord(words[words.length - 3]),
    end: parseCoord(words[words.length - 1]),
  };
});

const getIntermediateCoords = (start: Coordinate, end: Coordinate) => {
  const toReturn: Coordinate[] = [];
  for (let x = start.x; x <= end.x; x++) {
    for (let y = start.y; y <= end.y; y++) {
      toReturn.push({ x, y });
    }
  }
  return toReturn;
};

const getNewValuePart1 = (
  oldValue: boolean,
  command: Instruction["command"]
): boolean => {
  switch (command) {
    case "turnOn":
      return true;
    case "turnOff":
      return false;
    case "toggle":
      return !oldValue;
  }
};

const getNewValuePart2 = (
  oldValue: number,
  command: Instruction["command"]
): number => {
  switch (command) {
    case "turnOn":
      return oldValue + 1;
    case "turnOff":
      return oldValue === 0 ? 0 : oldValue - 1;
    case "toggle":
      return oldValue + 2;
  }
};
