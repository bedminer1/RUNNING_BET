<script lang="ts">
	let myInput = ""
	let herInput = ""
	$: scoreNeeded = calculateCutoff(herInput)
	$: winForMe = Number(myInput) > scoreNeeded

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
	<div class="w-1/4 flex flex-col gap-3 mb-5">
		<input type="text" bind:value={myInput} class="input" placeholder="Alex's score">
		<input type="text" bind:value={herInput} class="input" placeholder="Yunna's score">
		<p class="pl-2">Distance needed for Alex: {scoreNeeded.toFixed(2)}</p>
	</div>
	{#if !myInput.length || !herInput.length}
		<p>Please input scores</p>
	{:else if winForMe}
		<p>Alex wins</p>
	{:else}
		<p>Yuuna wins</p>
	{/if}
</div>
