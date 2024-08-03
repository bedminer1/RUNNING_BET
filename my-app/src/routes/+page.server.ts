import fs from "fs"
import path from "path"
import { z } from 'zod'
import { superValidate } from 'sveltekit-superforms'
import { zod } from 'sveltekit-superforms/adapters'
import { fail } from '@sveltejs/kit'

const schema = z.object({
    myScore: z.number().optional(),
    herScore: z.number().optional(),
    neededScore: z.number().optional(),
    winForMe: z.boolean()
})

export const load = async () => {
    const form = await superValidate(zod(schema))
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
        form,
        records
    }
}

export const actions = {
    saveRecord: async ({ request }) => {
        try {
            const form = await superValidate(request, zod(schema))
            // custom check for undefined since values must be intialized as undefined for placeholder to work
            if (form.data.myScore === undefined) {
                form.valid = false
                form.errors.myScore = ["*Please fill in field*"]
            }
            if (form.data.herScore === undefined) {
                form.valid = false
                form.errors.herScore = ["*Please fill in field*"]
            }
            console.log(form)

            if (!form.valid) {
                return fail(400, { form });
            }
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