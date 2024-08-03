<script lang="ts">
	import { superForm } from 'sveltekit-superforms'

	export let data

	const { form, message } = superForm(data.form)
	$: $form.neededScore = $form.herScore === undefined ? undefined : Number(calculateCutoff($form.herScore).toFixed(2))
	$: $form.winForMe = $form.myScore! > $form.neededScore!

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
	<form class="w-1/4 flex flex-col gap-3 mb-5" action="?/saveRecord" method="POST">
		<input name="myScore" type="text" bind:value={$form.myScore} class="input" placeholder="Alex's score">
		<input name="herScore" type="text" bind:value={$form.herScore} class="input" placeholder="윤아's score">
		<input name="neededScore" type="text" bind:value={$form.neededScore} class="input" placeholder="Distance cutoff">
		<input name="winForMe" type="hidden" bind:value={$form.winForMe}>
		{#if $message} 
			<p class="text-error-500">{$message}</p>
		{/if}
		<button class="btn">Save Result</button>
	</form>

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
