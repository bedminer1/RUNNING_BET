import fs from "fs"
import path from "path"
import { parseRecord } from "$lib/utils/parseRecords.js"
import { z } from 'zod'
import { superValidate, message } from 'sveltekit-superforms'
import { zod } from 'sveltekit-superforms/adapters'
import { fail } from '@sveltejs/kit'
import PocketBase from 'pocketbase'
import { SECRET_PASSWORD, SECRET_EMAIL, SECRET_URL } from '$env/static/private'

const schema = z.object({
    weekID: z.number(),
    myScore: z.number().optional(),
    herScore: z.number().optional(),
    neededScore: z.number().optional(),
    winForMe: z.boolean(),
    myPoints: z.number(),
    herPoints: z.number()
})

export const load = async () => {
    const form = await superValidate(zod(schema))
    const filePath = path.resolve("src/lib/local_storage/record.txt")
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
        const filePath = path.resolve("src/lib/local_storage/record.txt")

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

            if (!form.valid) {
                return fail(400, { form });
            }

            let newRecord: WeeklyRecord = {
                weekID: form.data.weekID,
                myScore: form.data.myScore,
                herScore: form.data.herScore,
                neededScore: form.data.neededScore,
                winForMe: form.data.winForMe,
                myPoints: form.data.myPoints,
                herPoints: form.data.herPoints
            }

            // saving to PB
            const pb = new PocketBase(SECRET_URL)
		    await pb.admins.authWithPassword(SECRET_EMAIL, SECRET_PASSWORD)
            const record = await pb.collection('RUNNING').create(newRecord)

            // saving to file
            let oldTxt = fs.readFileSync(filePath, "utf-8")
            let newTxt = javascriptToTxt(oldTxt, newRecord)
            fs.writeFileSync(filePath, newTxt, 'utf-8')

            return message(form, "Success")

        }
        catch (err) {
            console.error("Error saving record:", err)
        }
    }
}

function javascriptToTxt(oldTxt: string, record: WeeklyRecord): string {
    return oldTxt + "\n" + `${record.weekID} ${record.myScore?.toFixed(2)} ${record.herScore?.toFixed(2)} ${record.neededScore} ${record.winForMe} ${record.myPoints} ${record.herPoints}` 
}
