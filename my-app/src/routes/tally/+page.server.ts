import fs from "fs"
import path from "path"

// calculating previous week's records
function calculateScores(input: string) {
  let res = []
  for (let line of input.split("\n")) {
    let sum = 0
    for (let num of line.split(" ")) {
      sum += Number(num);
    }
    res.push(sum.toFixed(2))
  }

  return res;
}

const filePath = path.resolve("src/lib/alexScores.txt")
let input = fs.readFileSync(filePath, "utf-8")
let res = calculateScores(input)
console.log(res)
