import { readFileSync } from "fs";
import { resolve } from "path";

export const getFile = (filename: string) => {
  try {
    const path = resolve(__dirname, `../input/${filename}.txt`);
    return readFileSync(path, "utf8");
  } catch (err) {
    console.log(`Could not find file ${__dirname}/../input/${filename}.txt`);
    throw err;
  }
};

export const getFileLines = (filename: string) => {
  const data = getFile(filename);
  return data.split("\n");
};
