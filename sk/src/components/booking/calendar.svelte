<script lang="ts">
	import { createStepController } from '@efstajas/svelte-stepper';
	import { DatePicker } from '@svelte-plugins/datepicker';
	import { onMount } from 'svelte';

	const stepController = createStepController();

	const DAYS_IN_RANGE = 60;

	let isDaySelected = false,
		selectedDate = new Date(),
		enabledDates: string[] = [];

	function onDayClick(date: Date) {
		isDaySelected = true;
		selectedDate = date;
		stepController.nextStep();
	}

	function generateDateRange(): string[] {
		const startDate = new Date();
		const endDate = new Date();
		endDate.setDate(startDate.getDate() + DAYS_IN_RANGE);

		const dateRange: string[] = [];

		while (startDate <= endDate) {
			const startFormatted = `${
				startDate.getMonth() + 1
			}/${startDate.getDate()}/${startDate.getFullYear()}`;
			const endFormatted = `${
				endDate.getMonth() + 1
			}/${endDate.getDate()}/${endDate.getFullYear()}`;
			dateRange.push(`${startFormatted}:${endFormatted}`);
			startDate.setDate(startDate.getDate() + 1);
		}

		return dateRange;
	}

	onMount(() => {
		enabledDates = generateDateRange();
	});
	// TODO
	// when click day navigate to next step
</script>

<div class="flex flex-col">
	<div class="showcase relative">
		<div>
			<h1 class="text-lg leading-6 text-gray-900 font-bold text-center">Select a Day</h1>
		</div>

		<DatePicker
			isOpen={true}
			alwaysShow={true}
			enableFutureDates={true}
			{selectedDate}
			{onDayClick}
			{enabledDates}
		/>
	</div>
</div>

<!-- 
   {#if errorMessage}
    <div class="text-red-500">{errorMessage}</div>
  {/if}
  {#if isDaySelected}
   
  {/if}
 -->

<style>
	.showcase {
		display: table-caption;
		height: 440px;
		position: relative;
	}
</style>
