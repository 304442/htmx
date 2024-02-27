<script>
	import { createStepController } from '@efstajas/svelte-stepper';
	import BackArrow from '../icons/arrow-left.svelte';
	const stepController = createStepController();

	function saveEvent() {
		stepController.nextStep();
	}
</script>

<div class="max-w-80 mx-auto">
	<div class="flex justify-between items-start mb-6">
		<button
			class="text-2xl font-bold text-primary-500"
			type="button"
			on:click={() => stepController.previousStep()}
		>
			<BackArrow />
		</button>

		<h1 class="text-lg mx-auto leading-6 text-gray-900 font-bold text-center block pr-6">
			Details
		</h1>
	</div>

	<form class="">
		<div class="mb-4">
			<input
				type="text"
				class="w-full p-2 shadow rounded-lg bg-gray-100/70 outline-none focus:border-2 focus:border-black text-black border border-transparent"
				placeholder="Name"
				name="name"
			/>
		</div>

		<div class="mb-4">
			<input
				type="text"
				class="p-2 shadow rounded-lg bg-gray-100/70 outline-none focus:border-2 focus:border-black text-black border border-transparent w-full"
				placeholder="Email"
				name="email"
			/>
		</div>

		<div class="mb-4">
			<textarea
				class="w-full h-32 mt-2 p-3 rounded-lg bg-gray-100/70 outline-none focus:border-2 focus:border-black text-black border border-transparent"
				placeholder="Please share anything that will help prepare for our meeting"
				name="message"
			/>
		</div>

		<div>
			<button
				class="p-3 bg-black text-white font-bold w-full rounded-lg uppercase hover:bg-light-black disabled:cursor-not-allowed"
				on:click={() => saveEvent()}
				type="button"
			>
				Schedule Event
			</button>
		</div>
	</form>
</div>
