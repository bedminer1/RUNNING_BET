<script lang="ts">
	import { superForm } from 'sveltekit-superforms'

	export let data

	const { form, errors, message } = superForm(data.form)
	$: $form.neededScore = $form.herScore === undefined ? undefined : Number(calculateCutoff($form.herScore).toFixed(2))
	$: $form.winForMe = $form.myScore! > $form.neededScore!

	// calculateCutoff takes in Yoona's score and returns the score needed to beat her
	function calculateCutoff(input: number|undefined): number {
		let herScore: number = Number(input) ?? 0
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
<!-- FORM -->
	<form class="w-1/4 flex flex-col mb-5 justify-center items-center gap-3" action="?/saveRecord" method="POST">
		<div class="flex flex-col w-full">
			{#if $errors.myScore}<span class="invalid italic m-0 p-0 text-xs text-error-400 ml-2 w-full">{$errors.myScore}</span>{/if}
			<input name="myScore" type="text" bind:value={$form.myScore} class="input mt-0" placeholder="Alex's score">
		</div>
		
		<div class="flex flex-col w-full">
			{#if $errors.herScore}<span class="invalid italic m-0 p-0 text-xs text-error-400 ml-2 w-full">{$errors.herScore}</span>{/if}
			<input name="herScore" type="text" bind:value={$form.herScore} class="input mt-0" placeholder="윤아's score">
		</div>
		<input name="neededScore" type="text" bind:value={$form.neededScore} class="input" placeholder="Distance cutoff">
		<input name="winForMe" type="hidden" bind:value={$form.winForMe}>

		{#if $message} 
			<p class="text-success-500">{$message}</p>
		{/if}
		<button class="btn border-dashed border-2 w-1/2 rounded-lg mt-4 border-wheat-500 text-wheat-500">Save Result</button>
	</form>

<!-- TABLE -->
<div class="table-container w-1/2">
	<table class="table table-hover">
		<thead>
			<tr>
				<th>week</th>
				<th>Alex's Score</th>
				<th>Yoona's Score</th>
				<th>Needed Score</th>
				<th>Who won?</th>
				<th>Score</th>
			</tr>
		</thead>
		<tbody>
			{#each [...data.records].reverse() as record, i}
				<tr>
					<td>{record.weekID}</td>
					<td>{record.myScore}</td>
					<td>{record.herScore}</td>
					<td>{record.neededScore}</td>
					<td>{record.winForMe ? "Alex" : "윤아"}</td>
					<td>{record.score}</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>

</div>
