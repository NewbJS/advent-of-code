const fs = require("fs");

const data = fs.readFileSync("c:\\Users\\solar\\Downloads\\input.txt", "utf8");
const dArray = data.split("\n");
const dElementsAsInt = dArray.map((x) => {
  return parseInt(x, 10);
});

let answer;

for (let i of dElementsAsInt) {
  for (let n of dElementsAsInt) {
    for (let x of dElementsAsInt) {
      if (i + n + x === 2020) {
        answer =  i * n * x;
      }
    }
  }
}
console.log("answer:", answer);
