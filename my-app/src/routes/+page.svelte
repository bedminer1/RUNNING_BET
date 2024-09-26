<script lang="ts">
	import { superForm } from 'sveltekit-superforms'
	import { Paginator, getToastStore } from '@skeletonlabs/skeleton'
	import { CalculateCutoff } from '$lib/calculate.js';

	export let data

	const { form, errors, message } = superForm(data.form)

	$form.weekID = data.records.length + 1
	$: $form.neededScore = $form.herScore === undefined ? undefined : Number(CalculateCutoff($form.herScore, [[5,2], [10,1.5]]).toFixed(2))
	$: $form.winForMe = $form.myScore! > $form.neededScore!
	$: if ($form.winForMe) {
		$form.myPoints = data.records.at(-1)!.myPoints + 1
		$form.herPoints = data.records.at(-1)!.herPoints
	} else {
		$form.myPoints = data.records.at(-1)!.myPoints
		$form.herPoints = data.records.at(-1)!.herPoints + 1
	}
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

	// Configure scheme
	let scheme: number[][] = [[5,2], [10,1.5]]
	$: displayScheme = JSON.stringify(scheme)
	let distance: string
	let multiplier: string

	function addInterval(distance: string, multiplier: string) {
		scheme = [...scheme, [Number(distance), Number(multiplier)]]
		scheme.sort((a,b) => a[0] - b[0])
	}

	function deleteInterval() {
		scheme.pop()
		scheme = scheme
	}

</script>

<div class="flex flex-col h-screen items-center justify-center">
<!-- FORM -->
 <div class="w-full flex gap-10 justify-center">
	 <form class="w-1/4 flex flex-col mb-10 justify-center items-center gap-3" action="?/saveRecord" method="POST">
		 <div class="flex flex-col w-full">
			 {#if $errors.myScore}<span class="invalid italic m-0 p-0 text-xs text-error-400 ml-2 w-full">{$errors.myScore}</span>{/if}
			 <input name="myScore" type="text" bind:value={$form.myScore} class="input mt-0" placeholder="Alex's score">
		 </div>
		 
		 <div class="flex flex-col w-full">
			 {#if $errors.herScore}<span class="invalid italic m-0 p-0 text-xs text-error-400 ml-2 w-full">{$errors.herScore}</span>{/if}
			 <input name="herScore" type="text" bind:value={$form.herScore} class="input mt-0" placeholder="윤아's score">
		 </div>
		 <input name="neededScore" type="text" bind:value={$form.neededScore} class="input" placeholder="Distance cutoff">
		 <input name="weekID" type="hidden" bind:value={$form.weekID}>
		 <input name="myPoints" type="hidden" bind:value={$form.myPoints}>
		 <input name="herPoints" type="hidden" bind:value={$form.herPoints}>
		 <input name="winForMe" type="hidden" bind:value={$form.winForMe}>
		 <button class="btn variant-ghost-primary w-1/2 rounded-md mt-4 text-wheat-500">Save Result</button>
	 </form>
	 <form class="w-1/4 flex flex-col mb-10 justify-center items-center gap-3">
		<input name="neededScore" type="text" bind:value={distance} class="input" placeholder="Distance Upper Bound">
		<input name="neededScore" type="text" bind:value={multiplier} class="input" placeholder="Multiplier">
		<input class="input h-12" disabled placeholder="Scheme" bind:value={displayScheme}>
		<div class="flex gap-2">
			<button on:click={() => addInterval(distance, multiplier)} class="btn variant-ghost-primary w-1/2 rounded-md mt-4 text-wheat-500">Add</button>
			<button on:click={() => deleteInterval()} class="btn variant-ghost-primary w-1/2 rounded-md mt-4 text-wheat-500">Remove</button>
		</div>
	 </form>
 </div>

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
					<td class="text-center">{record.myScore?.toFixed(2)}</td>
					<td class="text-center">{record.herScore?.toFixed(2)}</td>
					<td class="text-center">{record.neededScore?.toFixed(2)}</td>
					<td class="text-center">{record.winForMe ? "Alex" : "윤아"}</td>
					<td class="text-center">{record.myPoints}-{record.herPoints}</td>
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
