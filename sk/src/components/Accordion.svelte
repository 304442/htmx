<script>
	// @ts-nocheck

	import { slide } from 'svelte/transition';
	let isOpen = false;
	export let Icon; // icon component
	export let text;
	const toggle = () => (isOpen = !isOpen);
</script>

<button
	class="flex item-center w-full p-4 bg-gray text-black hover:text-white hover:bg-light-black"
	on:click={toggle}
	aria-expanded={isOpen}
>
	<svelte:component this={Icon} />

	<span class="text-sm font-bold uppercase ml-4">{text}</span>

	<svg
		style="tran"
		width="20"
		height="20"
		fill="none"
		class="duration-200 ml-auto {isOpen && 'rotate-[0.25turn]'}"
		stroke-linecap="round"
		stroke-linejoin="round"
		stroke-width="2"
		viewBox="0 0 24 24"
		stroke="currentColor"><path d="M9 5l7 7-7 7" /></svg
	>
</button>
{#if isOpen}
	<div transition:slide={{ duration: 300 }}>
		<div class="text-gray-500 text-md p-4 body bg-white">
			<slot />
		</div>
	</div>
{/if}
