export function parseRecord(input: string): WeeklyRecord[] {
    let res: WeeklyRecord[] = []
    let lines = input.split("\n")
    for (let line of lines) {

        let data = line.split(" ")
        let record: WeeklyRecord = {
            weekID:  Number(data[0]),
            myScore: Number(data[1]),
            herScore: Number(data[2]),
            neededScore: Number(data[3]),
            winForMe: data[4] === "true",
            myPoints: Number(data[5]),
            herPoints: Number(data[6])
        }
        res.push(record)
    }

    return res
}