import { getFile } from "./utils/input";

const file = getFile("04");

type BingoBoard = (number | "X")[][];

const parseBoard = (
  file: string
): { numbers: number[]; boards: BingoBoard[] } => {
  const input = file.split("\n\n");
  console.log(input);

  const numbers = input[0].split(",").map((val) => parseInt(val));

  const boards = input.slice(1).map((stringBoard) =>
    stringBoard.split("\n").map((stringLine) =>
      stringLine
        .split(" ")
        .filter((val) => val)
        .map((val) => parseInt(val))
    )
  );

  return {
    numbers,
    boards,
  };
};

const markBoards = (boards: BingoBoard[], number: number) => {
  boards.forEach((board) => {
    board.forEach((line) => {
      line.forEach((val, index3) => {
        if (val === number) {
          line[index3] = "X";
        }
      });
    });
  });
  return boards;
};

const isBingo = (board: BingoBoard) => {
  // Rows
  board.forEach((line) => line.every((val) => val === "X"));

  // Columns
  for (let col = 0; col < board[0].length; col++) {
    let allMarked = true;
    for (let row = 0; row < board.length; row++) {
      if (board[row][col] !== "X") {
        allMarked = false;
      }
    }
    if (allMarked) {
      return true;
    }
  }
  return false;
};

const checkForWinningBoards = (boards: BingoBoard[]) => {
  return boards.filter((board) => isBingo(board));
};

export const day04part1 = () => {
  let { numbers, boards } = parseBoard(file);
  console.log(numbers, boards);

  numbers.forEach((drawnNumber) => {
    boards = markBoards(boards, drawnNumber);
    const winningBoards = checkForWinningBoards(boards);

    if (winningBoards.length >= 1) {
      const winningBoard = winningBoards[0];
      return 1;
      // return sumOfUnmarkedNums(winningBoard) * drawnNumber;
    }
  });
  return -1;
};
