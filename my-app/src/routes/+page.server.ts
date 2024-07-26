import fs from "fs"
import path from "path"

export const load = async () => {
    const filePath = path.resolve("src/lib/record.txt")
    let rawInput = ""
    let records: WeeklyRecord[] = []

    try {
        rawInput = fs.readFileSync(filePath, "utf-8")
        records = parseRecord(rawInput)
    } catch (err) {
        console.error("Failed to read summary file:", err)
    }

    return {
        records
    }
}

function parseRecord(input: string): WeeklyRecord[] {
    let res: WeeklyRecord[] = []
    for (let line of input.split("\n")) {
        let data = line.split(" ")
        let record: WeeklyRecord = {
            weekID: 1,
            myScore: Number(data[0]),
            herScore: Number(data[1]),
            neededScore: Number(data[2]),
            winForMe: data[3] === "true"
        }
        res.push(record)
    }

    return res
}