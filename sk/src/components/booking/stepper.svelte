<script context="module" lang="ts">
	export interface ExampleStepperContext {
		defaultTransitionDuration: number;
	}
</script>

<script lang="ts">
	import { makeStep } from '@efstajas/svelte-stepper';
	import { Stepper } from '@efstajas/svelte-stepper';
	import { createStepController } from '@efstajas/svelte-stepper';
	import { writable, type Writable } from 'svelte/store';
	import MeetForm from './meet-form.svelte';
	import Calendar from './calendar.svelte';
	import Hours from './hours.svelte';
	import BookingDetails from './booking-details.svelte';

	let context: Writable<ExampleStepperContext>;
	const stepperContext = () => {
		const w = writable<ExampleStepperContext>({
			defaultTransitionDuration: 1000
		});

		context = w;

		return w;
	};

	const exampleSteps = [
		// makeStep({
		// 	component: Hours,
		// 	props: undefined
		// }),
		makeStep({
			component: Calendar,
			props: undefined
		}),
		makeStep({
			component: Hours,
			props: undefined
		}),
		makeStep({
			component: MeetForm,
			props: undefined
		}),
		makeStep({
			component: BookingDetails,
			props: undefined
		})
	];

	let stepsMaxIndex = exampleSteps.length - 1;
	let currIndex = 0;

	let stepperWrapperElem: HTMLDivElement;

	// import { scale, fly, slide, fade } from 'svelte/transition';
</script>

<div class="h-[400px]">
	<div class="" bind:this={stepperWrapperElem}>
		<Stepper
			context={stepperContext}
			steps={exampleSteps}
			on:stepChange={(e) => {
				const { detail } = e;
				currIndex = detail.newIndex;
				stepsMaxIndex = detail.of;

				stepperWrapperElem.scrollTo({
					top: 0,
					behavior: 'smooth'
				});
			}}
			defaultTransitionDuration={$context?.defaultTransitionDuration}
		/>
	</div>
</div>

<!-- 
    stepIntroTransition={{
					transitionFn: slide,
					params: (direction) => ({
						duration: 5000,
						delay: 0,
						y: direction === 'forward' ? 100 : -100
					})
				}}
				stepOutroTransition={{
					transitionFn: fade,
					params: () => ({
						duration: 2000,
						delay: 0,
						scale: 1
					})
				}}
 -->
