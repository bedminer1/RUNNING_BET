import { fail } from "assert";
import fs from "fs"
import path from "path"
import PocketBase from 'pocketbase'
import { SECRET_PASSWORD, SECRET_EMAIL, SECRET_URL } from '$env/static/private'

// takes in raw scores from txt and return arr of numbers containing scores for each week
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

// variables to access the txt files containing the scores
const filePathAlex = path.resolve("src/lib/alexScores.txt")
const filePathYoona = path.resolve("src/lib/yoonaScores.txt")
const inputAlex = fs.readFileSync(filePathAlex, "utf-8")
const inputYoona = fs.readFileSync(filePathYoona, "utf-8")
const resAlex = calculateScores(inputAlex)
const resYoona = calculateScores(inputYoona)

const saveToPB = async (newRecord: WeeklyRecord) => {
  const pb = new PocketBase(SECRET_URL)
  await pb.admins.authWithPassword(SECRET_EMAIL, SECRET_PASSWORD)
  const record = await pb.collection('RUNNING').create(newRecord)
}
// returns an array of WeeklyRecord with the score at the end
function parseRecords(alexScores: string[], yoonaScores: string[]) {
  try {
    if (alexScores.length !== yoonaScores.length) {
      throw new Error("Array lengths not equal")
    }
    let res = []
    let alexScore = 0
    let yoonaScore = 0

    for (let i = 0; i < alexScores.length; i++) {
      if (Number(alexScores[i]) > calculateDifferentCutoffs(Number(yoonaScores[i]), i + 1)) alexScore++
      else yoonaScore++
      let newRecord: WeeklyRecord = {
          weekID: i + 1,
          myScore: Number(alexScores[i]),
          herScore: Number(yoonaScores[i]),
          neededScore: calculateDifferentCutoffs(Number(yoonaScores[i]), i + 1),
          winForMe: Number(alexScores[i]) > calculateDifferentCutoffs(Number(yoonaScores[i]), i + 1),
          myPoints: alexScore,
          herPoints: yoonaScore
      }
      res.push(newRecord)
      // uncomment to save to PB
      // saveToPB(newRecord)
    }
    
    return res
  } catch (err) {
    return []
  }
}

// returns different cutoffs depending on the multiplier scheme at the time
function calculateDifferentCutoffs(herScore: number, weekId: number) {
  let multiplier

  // before 1APR
  if (weekId <= 10) multiplier = (ele: number)=> 2 * ele
  // between 1APR and 16JUN
  else if (weekId <= 21) multiplier = (ele: number)=> 1.5 * ele
  // before 22JUL
  else if (weekId <= 26) multiplier = (ele: number)=> 1 * ele
  // after 22JUL
  else multiplier = calculateCurrentCutoff

  return Number(multiplier(herScore).toFixed(2))
}

// calculateCutoff takes in Yoona's score and 
// returns the score needed to beat her for the current tiered scheme
function calculateCurrentCutoff(input: number|undefined): number {
  let herScore: number = Number(input) ?? 0
  if (herScore < 5) {
    return herScore * 2
  } else if (herScore < 10) {
    // 10 + (herScore - 5) * 1.5
    return 2.5 + (1.5 * herScore)
  } else {
    // 17.5 + (herScore - 10)
    return 7.5 + herScore
  }
}

let records = parseRecords(resAlex, resYoona)
console.log(records)

function arrayToString(records: WeeklyRecord[]) {
  let res = ""
  for (let record of records) {
    res += `${record.weekID} ${record.myScore} ${record.herScore} ${record.neededScore} ${record.winForMe} ${record.myPoints} ${record.herPoints}\n`
  }
  return res
}

let outputTxt = arrayToString(records)
const outputFilePath = path.resolve("src/lib/autoTallyOutput.txt")
fs.writeFileSync(outputFilePath, outputTxt, 'utf-8')