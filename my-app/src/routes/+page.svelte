<script lang="ts">
	import { superForm } from 'sveltekit-superforms'
	import { Paginator, getToastStore } from '@skeletonlabs/skeleton'

	export let data

	const { form, errors, message, enhance } = superForm(data.form)
	$: $form.neededScore = $form.herScore === undefined ? undefined : Number(calculateCutoff($form.herScore).toFixed(2))
	$: $form.winForMe = $form.myScore! > $form.neededScore!
	// toasts logic
	const toastScore = getToastStore()
	const successToast = {
		message: "Saved 👍🏻",
		hideDismiss: true,
		timeout: 1000,
		hoverable: true,
		background: "variant-filled-success",
		classes: "border-md"
	}
	const errorToast = {
		message: "Failed to Save :(",
		hideDismiss: true,
		timeout: 1000,
		hoverable: true,
		background: "variant-filled-error",
		classes: "border-md"
	}

	$: if ($errors.myScore && $errors.herScore) {
		toastScore.trigger(errorToast)
	} else if ($message === "Success") {
		toastScore.trigger(successToast)
	}

	// paginator logic
	const source = [...data.records].reverse()
	let paginationSettings = {
		page: 0,
		limit: 5,
		size: source.length,
		amounts: [5,10,30],
	}
	$: paginatedSource = source.slice(
		paginationSettings.page * paginationSettings.limit,
		paginationSettings.page * paginationSettings.limit + paginationSettings.limit
	)

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
	<form class="w-1/4 flex flex-col mb-10 justify-center items-center gap-3" action="?/saveRecord" method="POST" use:enhance>
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
		<button class="btn variant-ghost-primary w-1/2 rounded-md mt-4 text-wheat-500">Save Result</button>
	</form>

<!-- TABLE -->
<div class="table-container w-1/2">
	<table class="table table-hover mb-4">
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
			{#each paginatedSource as record, i}
				<tr>
					<td class="text-center">{record.weekID}</td>
					<td class="text-center">{record.myScore}</td>
					<td class="text-center">{record.herScore}</td>
					<td class="text-center">{record.neededScore}</td>
					<td class="text-center">{record.winForMe ? "Alex" : "윤아"}</td>
					<td class="text-center">{record.score}</td>
				</tr>
			{/each}
		</tbody>
	</table>
	<Paginator
		bind:settings={paginationSettings}
		showFirstLastButtons="{false}"
		showPreviousNextButtons="{true}"
		amountText="Records"
	/>
</div>

</div>
