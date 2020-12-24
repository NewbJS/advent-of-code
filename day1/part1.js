const fs = require("fs");

const data = fs.readFileSync("c:\\Users\\solar\\Downloads\\input.txt", "utf8");
const dArray = data.split("\n");
const dElementsAsInt = dArray.map((x) => {
  return parseInt(x, 10);
});

let answer;

for (let i of dElementsAsInt) {
    const z = dElementsAsInt.find(n => n + i === 2020);
    if (z != undefined){
        answer = z * i;
    }
}

console.log("answer:", answer);
