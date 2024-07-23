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

<div>
	<input type="text" bind:value={myInput} class="input">
	<input type="text" bind:value={herInput} class="input">
	{winForMe}
</div>
