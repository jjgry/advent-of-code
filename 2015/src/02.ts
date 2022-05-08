import { getFileLines } from "./utils/input";

const lines = getFileLines("02");

const boxes = lines.map((line) => {
  const [l, w, h] = line.split("x");
  return {
    l: parseInt(l),
    w: parseInt(w),
    h: parseInt(h),
  };
});

export const day02part1 = () => {
  let total = 0;
  for (const box of boxes) {
    const face1 = box.l * box.w;
    const face2 = box.w * box.h;
    const face3 = box.h * box.l;

    const faces = 2 * (face1 + face2 + face3);
    const extra = Math.min(face1, face2, face3);
    total += faces + extra;
  }
  return total;
};

export const day02part2 = () => {
  let total = 0;
  for (const box of boxes) {
    const perimeter1 = 2 * (box.l + box.w);
    const perimeter2 = 2 * (box.w + box.h);
    const perimeter3 = 2 * (box.h + box.l);

    const perimeter = Math.min(perimeter1, perimeter2, perimeter3);
    const bow = box.l * box.w * box.h;
    total += perimeter + bow;
  }
  return total;
};
