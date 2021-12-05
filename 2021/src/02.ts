import { getFileLines } from "./utils/input";

const lines = getFileLines("02");

const navigateSumbmarine = (useAim: boolean) => {
  let height = 0;
  let distance = 0;
  let aim = 0;

  lines.forEach((line) => {
    if (line === "") {
      return;
    }
    const [direction, magnitude] = line.split(" ");
    const magNum = parseInt(magnitude);

    if (useAim) {
      switch (direction) {
        case "forward":
          distance += magNum;
          height += aim! * magNum;
          break;
        case "up":
          aim! -= magNum;
          break;
        case "down":
          aim! += magNum;
          break;
        default:
          throw new Error(
            `Cannot process unrecognised direction '${direction}'`
          );
      }
    } else {
      switch (direction) {
        case "forward":
          distance += magNum;
          break;
        case "up":
          height -= magNum;
          break;
        case "down":
          height += magNum;
          break;
        default:
          throw new Error(
            `Cannot process unrecognised direction '${direction}'`
          );
      }
    }
  });

  return height * distance;
};

export const day02part1 = () => navigateSumbmarine(false);
export const day02part2 = () => navigateSumbmarine(true);
