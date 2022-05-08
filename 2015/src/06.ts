import { getFileLines } from "./utils/input";

const input = getFileLines("06");

type Coordinate = { x: number; y: number };

type Instruction = {
  command: "turnOn" | "turnOff" | "toggle";
  start: Coordinate;
  end: Coordinate;
};

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

export const day06part1 = () => {
  const grid: boolean[][] = new Array(1000)
    .fill(false)
    .map(() => new Array(1000).fill(false));

  for (const { command, start, end } of instructions) {
    const coords = getAllIntermediateCoords(start, end);
    for (const { x, y } of coords) {
      grid[x][y] = getNewValue(grid[x][y], command);
    }
  }

  const numTurnedOn = grid.reduce((accumulator, gridRow) => {
    return accumulator + gridRow.filter((x) => x).length;
  }, 0);

  return numTurnedOn;
};

export const day06part2 = () => {
  return 0;
};

const getAllIntermediateCoords = (start: Coordinate, end: Coordinate) => {
  const toReturn: Coordinate[] = [];
  for (let x = start.x; x <= end.x; x++) {
    for (let y = start.y; y <= end.y; y++) {
      toReturn.push({ x, y });
    }
  }
  return toReturn;
};

const getNewValue = (
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
