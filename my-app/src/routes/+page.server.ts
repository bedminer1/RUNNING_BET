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

export const actions = {
    saveRecord: async ({ request }) => {
        try {
            const data = await request.formData()
            console.log(data)
        }
        catch (err) {
            console.error("Error saving record:", err)
        }
    }
}

function parseRecord(input: string): WeeklyRecord[] {
    let res: WeeklyRecord[] = []
    let lines = input.split("\n")
    for (let line of lines) {

        let data = line.split(" ")
        let record: WeeklyRecord = {
            weekID: lines.indexOf(line) + 1,
            myScore: Number(data[0]),
            herScore: Number(data[1]),
            neededScore: Number(data[2]),
            winForMe: data[3] === "true"
        }
        res.push(record)
    }

    return res
}