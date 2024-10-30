<script lang="ts">
    import LineChart from '$lib/components/LineChart.svelte';

    export let data
    let records = data.records

    // graph comparing points overtime
    let alexPoints: number[] = []
    let yoonaPoints: number[] = []
    for (let r of records) {
        alexPoints.push(r.myPoints!)
        yoonaPoints.push(r.herPoints!)
    }
    
    // graph comparing distance ran per week
    let alexScores: number[] = []
    let yoonaScores: number[] = []
    for (let r of records) {
        alexScores.push(r.myScore!)
        yoonaScores.push(r.herScore!)
    }

    // graph showing total distance ran
    let alexAccmScores: number[] = []
    let alexAccmScore = 0
    let yoonaAccmScores: number[] = []
    let yoonaAccmScore = 0

    for (let i = 0; i < alexScores.length; i++) {
    alexAccmScore += alexScores[i];
    yoonaAccmScore += yoonaScores[i];

    // Format to 2 decimal places and convert back to float
    alexAccmScores[i] = parseFloat(alexAccmScore.toFixed(2));
    yoonaAccmScores[i] = parseFloat(yoonaAccmScore.toFixed(2));
}
    // streak system showing number of weeks where person ran more than X (their goal)

    let ChartData = [];
    let accmData = {
        alexAccmScores: alexAccmScores,
        yoonaAccmScores: yoonaAccmScores
    }
    for (let i = 0; i < alexScores.length; i++) {
        ChartData.push({
            Week: i + 1,
            alexScores: alexScores[i],
            yoonaScores: yoonaScores[i],
            alexPoints: alexPoints[i],
            yoonaPoints: yoonaPoints[i],
            alexAccmScores: alexAccmScores[i],
            yoonaAccmScores: yoonaAccmScores[i],
        });
    }


    // console.table(ChartData)
</script>

<h1>hi there</h1>
<div class="w-screen border-2 h-full">
    <LineChart {...accmData}></LineChart>
</div>