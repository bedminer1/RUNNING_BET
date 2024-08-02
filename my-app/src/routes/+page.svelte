<script lang="ts">
	export let data

	let myInput = ""
	let herInput = ""
	$: scoreNeeded = herInput === "" ? undefined : Number(calculateCutoff(herInput).toFixed(2))
	$: winForMe = Number(myInput) > scoreNeeded!

	function calculateCutoff(input: string): number {
		let herScore: number = Number(input)
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
</script>

<div class="flex flex-col h-screen items-center justify-center">
	<form class="w-1/4 flex flex-col gap-3 mb-5" action="?/saveRecord" method="POST">
		<input name="myScore" type="text" bind:value={myInput} class="input" placeholder="Alex's score">
		<input name="herScore" type="text" bind:value={herInput} class="input" placeholder="윤아's score">
		<input name="scoreNeeded" type="text" bind:value={scoreNeeded} class="input" placeholder="Distance cutoff">
		<input name="winForMe" type="hidden" bind:value={winForMe}>
		<button class="btn">Save Result</button>
	</form>

	<div class="mb-10">
		{#if !myInput.length || !herInput.length}
			<p>Please input scores</p>
		{:else if winForMe}
			<p>Alex wins</p>
		{:else}
			<p>윤아 wins</p>
		{/if}
	</div>

<div class="table-container w-1/2">
	<table class="table table-hover">
		<thead>
			<tr>
				<th>week</th>
				<th>Alex's Score</th>
				<th>Yoona's Score</th>
				<th>Needed Score</th>
				<th>Who won?</th>
			</tr>
		</thead>
		<tbody>
			{#each data.records as record, i}
				<tr>
					<td>{record.weekID}</td>
					<td>{record.myScore}</td>
					<td>{record.herScore}</td>
					<td>{record.neededScore}</td>
					<td>{record.winForMe}</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>

</div>
