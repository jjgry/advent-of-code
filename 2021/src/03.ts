import { getFileLines } from "./utils/input";

const lines = getFileLines("03");

export const day03part1 = () => {
  const linelength = lines[0].length;
  let gammaString = "";
  let epsillonString = "";

  for (let i = 0; i < linelength; i++) {
    let numOnes = 0;
    let numZeroes = 0;
    lines.forEach((line) => {
      const bit = line[i];
      if (bit === "0") {
        numZeroes++;
      } else if (bit === "1") {
        numOnes++;
      } else {
        console.log("not meant to happen");
      }
    });

    if (numOnes > numZeroes) {
      gammaString += "1";
      epsillonString += "0";
    } else if (numOnes < numZeroes) {
      gammaString += "0";
      epsillonString += "1";
    } else {
      console.log("shouldn't happen either");
    }
  }

  const gamma = parseInt(gammaString, 2);
  const epsillon = parseInt(epsillonString, 2);
  return gamma * epsillon;
};

export const day03part2 = () => {
  const linelength = lines[0].length;
  const getNum = (more: boolean) => {
    let newLines = lines.slice();

    for (let i = 0; i < linelength; i++) {
      if (newLines.length <= 1) {
        break;
      }

      let numOnes = 0;
      let numZeroes = 0;
      newLines.forEach((line) => {
        const bit = line[i];
        if (bit === "0") {
          numZeroes++;
        } else if (bit === "1") {
          numOnes++;
        } else {
          console.log("not meant to happen");
        }
      });

      if (numOnes > numZeroes) {
        newLines = newLines.filter((line) => line[i] === (more ? "1" : "0"));
      } else if (numOnes < numZeroes) {
        newLines = newLines.filter((line) => line[i] === (more ? "0" : "1"));
      } else {
        newLines = newLines.filter((line) => line[i] === (more ? "1" : "0"));
      }
    }

    return parseInt(newLines[0], 2);
  };

  const oxy = getNum(true);
  const co2 = getNum(false);

  return oxy * co2;
};
