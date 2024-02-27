<script>
	import { Card } from 'flowbite-svelte';
	import {
		ClockOutline,
		CalendarMonthOutline,
		MapLocationOutline,
		UserOutline,
		CheckCircleSolid
	} from 'flowbite-svelte-icons';

	// import { createStepController } from '@efstajas/svelte-stepper';
	// const stepController = createStepController();
</script>

<div class="">
	<div>
		<h1 class="text-lg leading-6 text-gray-900 font-bold text-center mb-14">Confirmation</h1>
	</div>

	<div class="w-80 mx-auto">
		<Card class="w-80 mx-auto">
			<h1 class="text-center font-bold text-xl mb-4">
				<span>Booked successfully</span>
				<CheckCircleSolid class="w-6 h-6 text-green-500 dark:text-gray-400 inline" />
			</h1>

			<dl class="grid grid-cols-6 gap-x-4 gap-y-2">
				<dt class="col-span-1 flex items-center"><UserOutline class="w-6 h-6" /></dt>
				<dd class="col-span-5 flex items-center">Ahmed Zougari</dd>

				<dt class="col-span-1"><ClockOutline class="w-6 h-6" /></dt>
				<dd class="col-span-5">30 min</dd>

				<dt class="col-span-1"><CalendarMonthOutline class="w-6 h-6" /></dt>
				<dd class="col-span-5">10:30 - 11:00, Thursday, February 22, 2024</dd>

				<dt class="col-span-1"><MapLocationOutline class="w-6 h-6" /></dt>
				<dd class="col-span-5">Central European Time</dd>
			</dl>

		</Card>
		<div class="mt-4">
			<button
				class="p-3 bg-black text-white font-bold w-full rounded-lg uppercase hover:bg-light-black disabled:cursor-not-allowed"
				type="button"
			>
				Confirm
			</button>
		</div>
	</div>
</div>
