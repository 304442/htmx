<script lang="ts">
	import { createStepController } from '@efstajas/svelte-stepper';
	import { createQuery } from '@tanstack/svelte-query';
	import { Button, Label } from 'flowbite-svelte';
	import moment from 'moment-timezone';
	import Select from 'svelte-select';
	import { onMount } from 'svelte';
	import BackArrow from '../icons/arrow-left.svelte';

	const START_HOUR = 10,
		END_HOUR = 18,
		DEFAULT_DURATION = 30,
		guessedTimeZone = moment.tz.guess(),
		timeZoneList = moment.tz.names();

	let timeZone = {
		label: guessedTimeZone,
		value: guessedTimeZone
	};
	// const offsetFromGMT = moment.tz(guessedTimeZone).utcOffset();

	const stepController = createStepController();
	let selected: string = '';

	function generateAvailableHours(minutes: number): string[] {
		const availableHours: string[] = [];
		for (let hour = START_HOUR; hour < END_HOUR; hour++) {
			for (let minute = 0; minute < 60; minute += minutes) {
				const time = `${hour.toString().padStart(2, '0')}:${minute.toString().padStart(2, '0')}`;
				availableHours.push(time);
			}
		}
		return availableHours;
	}

	function onHourClick(hour: string) {
		stepController.nextStep();
	}

	let hours = generateAvailableHours(DEFAULT_DURATION);

	console.log(hours);

	function onTimeZoneChange(event: any) {
		// const newTimeZone = event.target.value;
		// TODO here we convert hours to gmt so wee need last timezone sleceted

		const newTimeZone = timeZone.value;
		console.log(timeZone);
		const offsetFromGMT = moment.tz(newTimeZone).utcOffset(); // in minute (Number)

		const newHours = hours.map((time) => {
			const [hour, minute] = time.split(':').map(Number);

			let newMinute = (minute + offsetFromGMT) % 60;
			let carryHour = Math.floor((minute + offsetFromGMT) / 60);
			let newHour = (hour + carryHour) % 24;

			const formattedHour = newHour.toString().padStart(2, '0');
			const formattedMinute = newMinute.toString().padStart(2, '0');

			return `${formattedHour}:${formattedMinute}`;
		});

		console.log('🚀 ~ newHours ~ newHours:', newHours);
	}

	// TODO on select change call or map on hours

	// onMount(async () => {});

	// const query = createQuery({
	// 	queryKey: ['available-hours'],
	// 	queryFn: async () => await fetch('').then((r) => r.json()),
	// 	refetchInterval: 50000
	// });

	// collection events
	// client_name, email, description, date, start, end

	import PocketBase from 'pocketbase';
	const pb = new PocketBase('https://app.48d1.cc');

	let takenHours: string[] = [];
	async function test() {
		// const authData = await pb.admins.authWithPassword('admin@gmail.com', '1234512345');
		// const resultList = await pb.collection('events').getFullList({
		// 	// filter: 'created >= "2022-01-01 00:00:00" && someField1 != someField2',
		// 	filter: 'start >= "2023-01-01 00:00:00" && start < "2023-01-02 00:00:00"',
		// 	fields: 'start'
		// });

		// takenHours = resultList.map((item) => item.start);

		// console.log(takenHours);

		stepController.nextStep();
	}
</script>

<div class="max-w-80 mx-auto">
	<div class="flex justify-between items-start mb-4">
		<button
			class="text-2xl font-bold text-primary-500"
			type="button"
			on:click={() => stepController.previousStep()}
		>
			<BackArrow />
		</button>

		<div class="text-center mx-auto pr-6">
			<span class="text-lg leading-6 text-gray-900 font-bold text-center block">Wednesday</span>
			<span class="block text-sm">july 26, 2023</span>
		</div>
	</div>

	<div>
		<Label class="mb-1">Time zone</Label>
		<Select
			items={timeZoneList}
			showChevron
			bind:value={timeZone}
			on:select={onTimeZoneChange}
			clearable={false}
		/>
	</div>

	<hr class="divide border-b-0 px-2 my-2" />

	<div class="max-h-56 overflow-y-scroll">
		{#each hours as hour}
			<Button
				color="alternative"
				size="sm"
				type="button"
				class="w-full block mb-2"
				on:click={() => test()}>{hour}</Button
			>
		{/each}
	</div>
</div>
