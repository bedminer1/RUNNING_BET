<script lang="ts">
    import LineChart from '$lib/components/LineChart.svelte';

    export let data
    let records = data.records

    // graph comparing points overtime
    let alexPoints: number[] = []
    let yoonaPoints: number[] = []
    
    // graph comparing distance ran per week
    let alexScores: number[] = []
    let yoonaScores: number[] = []
    let neededScores: number[] = []

    for (let r of records) {
        alexPoints.push(r.myPoints!)
        yoonaPoints.push(r.herPoints!)
        alexScores.push(r.myScore!)
        yoonaScores.push(r.herScore!)
        neededScores.push(r.neededScore!)
    }

    // graph showing total distance ran
    let alexAccmScores: number[] = []
    let alexAccmScore = 0
    let yoonaAccmScores: number[] = []
    let yoonaAccmScore = 0

    for (let i = 0; i < alexScores.length; i++) {
        alexAccmScore += alexScores[i]
        yoonaAccmScore += yoonaScores[i]

        // Format to 2 decimal places and convert back to float
        alexAccmScores[i] = parseFloat(alexAccmScore.toFixed(2))
        yoonaAccmScores[i] = parseFloat(yoonaAccmScore.toFixed(2))
    }

    // streak system showing number of weeks where person ran more than X (their goal)
    let alexStreakGoal: number = 5
    let alexStreakDays: number[] = new Array(alexScores.length).fill(0)
    let yoonaStreakGoal: number = 7
    let yoonaStreakDays: number[] = new Array(yoonaScores.length).fill(0)

    $: {
        for (let i = 0; i < alexScores.length; i++) {
            if (alexScores[i] > alexStreakGoal) {
                if (i === 0) {
                    alexStreakDays[i] = 1
                } else {
                    alexStreakDays[i] = alexStreakDays[i-1] + 1
                }
            }
    
            if (yoonaScores[i] > yoonaStreakGoal) {
                if (i === 0) {
                    yoonaStreakDays[i] = 1
                } else {
                    yoonaStreakDays[i] = yoonaStreakDays[i-1] + 1
                }
            }
        }
    }




    let ChartData = []

    for (let i = 0; i < alexScores.length; i++) {
        ChartData.push({
            Week: i + 1,
            alexScores: alexScores[i],
            yoonaScores: yoonaScores[i],
            alexPoints: alexPoints[i],
            yoonaPoints: yoonaPoints[i],
            alexAccmScores: alexAccmScores[i],
            yoonaAccmScores: yoonaAccmScores[i],
        })
    }


    // console.table(ChartData)
</script>

<div class="flex flex-col w-full justify-center items-center p-10">
    <div class="w-full"> 
        <a href="/" class="text-3xl">&#9756;</a>
    </div>
    <h1 class="text-4xl mb-4">RUNNING BET STATS</h1>
    <div class="text-xl mb-10">
        as of {new Date().toLocaleDateString()}
    </div>
    <div class="w-3/4 flex flex-col">
        <div class="mb-4">
            <h1 class="text-center mb-2 text-xl">Accumulated Score</h1>
            <div class="flex justify-center gap-10">
                <p>Alex: {alexAccmScores[alexAccmScores.length-1].toFixed(2)} km</p>
                <p>|</p>
                <p>Yoona: {yoonaAccmScores[yoonaAccmScores.length-1].toFixed(2)} km</p>
            </div>
        </div>
        <div class="mb-10 border-2 border-dotted rounded-lg px-4 py-2">
            <LineChart {...{stats: [
                { label: 'Alex', data: alexAccmScores, borderColor: 'rgba(75, 192, 192, 1)', backgroundColor: 'rgba(75, 192, 192, 0.2)' },
                { label: 'Yoona', data: yoonaAccmScores, borderColor: 'rgba(255, 99, 132, 1)', backgroundColor: 'rgba(255, 99, 132, 0.2)' }
            ], label: "Accumulated Score"}}></LineChart>
        </div>

        <div class="mb-4">
            <h1 class="text-center mb-2 text-xl">Latest Weekly Score</h1>
            <div class="flex justify-center gap-10">
                <p>Alex: {alexScores[alexScores.length-1].toFixed(2)} km</p>
                <p>|</p>
                <p>Yoona: {yoonaScores[yoonaScores.length-1].toFixed(2)} km</p>
            </div>
        </div>
        <div class="mb-10 border-2 border-dotted rounded-lg px-4 py-2">
            <LineChart {...{stats: [
                { label: 'Alex', data: alexScores, borderColor: 'rgba(75, 192, 192, 1)', backgroundColor: 'rgba(75, 192, 192, 0.2)' },
                { label: 'Yoona', data: yoonaScores, borderColor: 'rgba(255, 99, 132, 1)', backgroundColor: 'rgba(255, 99, 132, 0.2)' },
                { label: 'Needed', data: neededScores, borderColor: 'rgba(255, 165, 0, 1)', backgroundColor: 'rgba(255, 165, 0, 0.2)' },
            ], label: "Weekly Scores"}}></LineChart>
        </div>

        <div class="mb-4">
            <h1 class="text-center mb-2 text-xl">Current Points</h1>
            <div class="flex justify-center gap-10">
                <p>Alex: {alexPoints[alexPoints.length-1]}</p>
                <p>|</p>
                <p>Yoona: {yoonaPoints[yoonaScores.length-1]}</p>
            </div>
        </div>
        <div class="mb-10 border-2 border-dotted rounded-lg px-4 py-2">
            <LineChart {...{stats: [
                { label: 'Alex', data: alexPoints, borderColor: 'rgba(75, 192, 192, 1)', backgroundColor: 'rgba(75, 192, 192, 0.2)' },
                { label: 'Yoona', data: yoonaPoints, borderColor: 'rgba(255, 99, 132, 1)', backgroundColor: 'rgba(255, 99, 132, 0.2)' }
            ], label: "Points"}}></LineChart>
        </div>

        <div class="w-1/3">
            <input type="number" class="input" bind:value={alexStreakGoal}>
            <input type="number" class="input" bind:value={yoonaStreakGoal}>
        </div>
        <div class="mb-10 border-2 border-dotted rounded-lg px-4 py-2">
            <LineChart {...{stats: [
                { label: 'Alex', data: alexStreakDays, borderColor: 'rgba(75, 192, 192, 1)', backgroundColor: 'rgba(75, 192, 192, 0.2)' },
                { label: 'Yoona', data: yoonaStreakDays, borderColor: 'rgba(255, 99, 132, 1)', backgroundColor: 'rgba(255, 99, 132, 0.2)' }
            ], label: "Streaks"}}></LineChart>
        </div>
    </div>
</div>