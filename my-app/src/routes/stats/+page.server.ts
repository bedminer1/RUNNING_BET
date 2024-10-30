import fs from "fs"
import path from "path"
import { parseRecord } from "$lib/utils/parseRecords"

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