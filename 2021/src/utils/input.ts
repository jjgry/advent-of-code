import { readFileSync } from "fs";
import { resolve } from "path";

export const getFileLines = (filename: string) => {
  try {
    const path = resolve(__dirname, `../input/${filename}.txt`);
    const data = readFileSync(path, "utf8");
    return data.split("\n");
  } catch (err) {
    console.log(`Could not find file ${__dirname}/../input/${filename}.txt`);
    throw err;
  }
};
