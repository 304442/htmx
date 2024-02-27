<script>
	// @ts-nocheck

	import { slide } from 'svelte/transition';
	import PocketBase from 'pocketbase';
	import countries from '../data/countries.json';
	import companies_size from '../data/company_size.json';
	import jobs_role from '../data/jobs_role.json';
	import industries from '../data/industries.json';

	let isOpen = false;
	let msgSentSuccessfully = false;
	let dataValid = false;
	let data = {
		name: {
			value: '',
			RegX: /\w{3,}/
		},
		email: {
			value: '',
			RegX: /^\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,3}$/
		},
		phone_number: {
			value: '',
			RegX: /^\d{10,}$/
		},
		message: {
			value: '',
			RegX: /\w{5,}/
		},
		business_type: {
			value: '',
			RegX: /\w{3,}/
		},
		job_title: {
			value: '',
			RegX: /\w{3,}/
		},
		company_name: {
			value: '',
			RegX: /\w{3,}/
		},
		job_role: {
			value: 'job role',
			RegX: /\w{3,}/
		},
		industry: {
			value: 'industry',
			RegX: /\w{3,}/
		},
		company_size: {
			value: 'company size',
			RegX: /\w{3,}/
			// RegX: /^\d+$/,
		},
		company_type: {
			value: '',
			RegX: /\w{3,}/
		},
		country: {
			value: 'country',
			RegX: /\w{3,}/
		},
		use_case: {
			value: '',
			RegX: /\w{3,}/
		}
	};
	const toggle = () => (isOpen = !isOpen);
	const isDataValid = () => {
		for (const key in data) {
			_data[key] = data[key].value;
			if (!data[key].RegX.test(data[key].value.trim())) {
				console.log(data[key]);
				return false;
			}
		}
		return true;
	};
	const checkInput = (target) => {
		const container = target.parentElement;
		const val = data[target.name].value.trim();
		const reg = data[target.name].RegX;
		const isValid = reg.test(val);
		container.querySelector('.msg').style.display = isValid ? 'none' : 'block';
		container.querySelector(!isValid ? '.danger' : '.success').style.display = 'block';
		container.querySelector(isValid ? '.danger' : '.success').style.display = 'none';
		target.style.borderColor = !isValid ? '#ef4444' : '#000000';

		if (isDataValid) dataValid = true;
	};
	const sendMessage = async (e) => {
		e.preventDefault();

		let _data = {};
		for (const key in data) {
			_data[key] = data[key].value;
			if (!data[key].RegX.test(data[key].value.trim())) {
				console.log(data[key]);
				return;
			}
		}

		const pb = new PocketBase('http://127.0.0.1:8090');
		const authData = await pb.admins.authWithPassword('admin@gmail.com', '1234512345');
		if (!!authData) {
			const record = await pb.collection('messages').create(_data);
			for (const key in data) {
				data[key].value = '';
			}
			msgSentSuccessfully = true;
			document.querySelectorAll('.success')?.forEach((i) => {
				i.style.display = 'none';
				i.parentElement.querySelector('input, textarea').style.borderColor = 'transparent';
			});
		}
	};
</script>

<button
	class="flex item-center w-full p-4 bg-gray text-black hover:text-white hover:bg-light-black"
	on:click={toggle}
	aria-expanded={isOpen}
>
	<svg
		xmlns="http://www.w3.org/2000/svg"
		fill="none"
		viewBox="0 0 24 24"
		stroke-width="1.5"
		stroke="currentColor"
		class="w-6 h-6"
	>
		<path
			stroke-linecap="round"
			stroke-linejoin="round"
			d="M9.879 7.519c1.171-1.025 3.071-1.025 4.242 0 1.172 1.025 1.172 2.687 0 3.712-.203.179-.43.326-.67.442-.745.361-1.45.999-1.45 1.827v.75M21 12a9 9 0 11-18 0 9 9 0 0118 0zm-9 5.25h.008v.008H12v-.008z"
		/>
	</svg>

	<span class="text-sm font-bold uppercase ml-4">request a free quote</span>

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
			<div class="text-center">
				<h1 class="text-lg leading-6 text-gray-900 font-bold">Get in touch, we’re happy to help</h1>
				<h2 class="text-sm leading-6 text-gray-500 -mt-1">
					Fill out the form and we'll review your company's systems and needs.
				</h2>
				<h3 class="text-sm leading-6 text-gray-500 -mt-1">
					You'll receive a free consultation and a quote.
				</h3>
			</div>
			<form class="grid grid-cols-2 gap-3">
				<!-- name -->
				<div class="flex-col flex py-3 relative">
					<input
						type="text"
						class="p-2 shadow rounded-lg bg-gray-100/70 outline-none focus:border-2 focus:border-black text-black border border-transparent"
						placeholder="Name"
						name="name"
						bind:value={data.name.value}
						on:input={(e) => checkInput(e.target)}
					/>
					<div
						class="danger text-red-500 text-xl font-bold absolute right-3 top-8 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-exclamation-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M7.002 11a1 1 0 1 1 2 0 1 1 0 0 1-2 0zM7.1 4.995a.905.905 0 1 1 1.8 0l-.35 3.507a.552.552 0 0 1-1.1 0L7.1 4.995z"
							/>
						</svg>
					</div>
					<div
						class="success text-black-500 text-xl font-bold absolute right-3 top-1/2 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-check-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M10.97 4.97a.235.235 0 0 0-.02.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-1.071-1.05z"
							/>
						</svg>
					</div>
					<p class="text-sm text-red-500 msg hidden">Please Enter a valid name.</p>
				</div>

				<!-- email -->
				<div class="flex-col flex py-3 relative">
					<input
						type="text"
						class="p-2 shadow rounded-lg bg-gray-100/70 outline-none focus:border-2 focus:border-black text-black border border-transparent"
						placeholder="Email"
						name="email"
						bind:value={data.email.value}
						on:input={(e) => checkInput(e.target)}
					/>
					<div
						class="danger text-red-500 text-xl font-bold absolute right-3 top-8 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-exclamation-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M7.002 11a1 1 0 1 1 2 0 1 1 0 0 1-2 0zM7.1 4.995a.905.905 0 1 1 1.8 0l-.35 3.507a.552.552 0 0 1-1.1 0L7.1 4.995z"
							/>
						</svg>
					</div>
					<div
						class="success text-black-500 text-xl font-bold absolute right-3 top-1/2 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-check-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M10.97 4.97a.235.235 0 0 0-.02.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-1.071-1.05z"
							/>
						</svg>
					</div>
					<p class="text-sm text-red-500 msg hidden">Please Enter a valid email.</p>
				</div>

				<!-- Phone number -->
				<div class="flex-col flex py-3 relative">
					<input
						type="text"
						class="p-2 shadow rounded-lg bg-gray-100/70 outline-none focus:border-2 focus:border-black text-black border border-transparent"
						placeholder="Phone Number"
						name="phone_number"
						bind:value={data.phone_number.value}
						on:input={(e) => checkInput(e.target)}
					/>
					<div
						class="danger text-red-500 text-xl font-bold absolute right-3 top-8 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-exclamation-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M7.002 11a1 1 0 1 1 2 0 1 1 0 0 1-2 0zM7.1 4.995a.905.905 0 1 1 1.8 0l-.35 3.507a.552.552 0 0 1-1.1 0L7.1 4.995z"
							/>
						</svg>
					</div>
					<div
						class="success text-black-500 text-xl font-bold absolute right-3 top-1/2 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-check-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M10.97 4.97a.235.235 0 0 0-.02.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-1.071-1.05z"
							/>
						</svg>
					</div>
					<p class="text-sm text-red-500 msg hidden">Please Enter a valid phone number.</p>
				</div>

				<!-- job title -->
				<div class="flex-col flex py-3 relative">
					<input
						type="text"
						class="p-2 shadow rounded-lg bg-gray-100/70 outline-none focus:border-2 focus:border-black text-black border border-transparent"
						placeholder="Job Title"
						name="job_title"
						bind:value={data.job_title.value}
						on:input={(e) => checkInput(e.target)}
					/>
					<div
						class="danger text-red-500 text-xl font-bold absolute right-3 top-8 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-exclamation-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M7.002 11a1 1 0 1 1 2 0 1 1 0 0 1-2 0zM7.1 4.995a.905.905 0 1 1 1.8 0l-.35 3.507a.552.552 0 0 1-1.1 0L7.1 4.995z"
							/>
						</svg>
					</div>
					<div
						class="success text-black-500 text-xl font-bold absolute right-3 top-1/2 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-check-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M10.97 4.97a.235.235 0 0 0-.02.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-1.071-1.05z"
							/>
						</svg>
					</div>
					<p class="text-sm text-red-500 msg hidden">Please Enter a valid job title.</p>
				</div>

				<!-- job role -->
				<div class="flex-col flex py-3 relative">
					<!-- <input
            type="text"
            class="p-2 shadow rounded-lg bg-gray-100/70 outline-none focus:border-2 focus:border-black text-black border border-transparent"
            placeholder="job role"
            name="job_role"
            bind:value={data.job_role.value}
            on:input={(e) => checkInput(e.target)}
          /> -->
					<select
						bind:value={data.job_role.value}
						class="p-2 shadow rounded-lg bg-gray-100/70 outline-none focus:border-2 focus:border-black text-black border border-transparent"
					>
						<option disabled>job role</option>
						{#each jobs_role as c}
							<option value={c}>{c}</option>
						{/each}
					</select>
					<div
						class="danger text-red-500 text-xl font-bold absolute right-3 top-8 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-exclamation-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M7.002 11a1 1 0 1 1 2 0 1 1 0 0 1-2 0zM7.1 4.995a.905.905 0 1 1 1.8 0l-.35 3.507a.552.552 0 0 1-1.1 0L7.1 4.995z"
							/>
						</svg>
					</div>
					<div
						class="success text-black-500 text-xl font-bold absolute right-3 top-1/2 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-check-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M10.97 4.97a.235.235 0 0 0-.02.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-1.071-1.05z"
							/>
						</svg>
					</div>
					<p class="text-sm text-red-500 msg hidden">Please Enter a valid job role.</p>
				</div>

				<!-- business type -->
				<div class="flex-col flex py-3 relative">
					<input
						type="text"
						class="p-2 shadow rounded-lg bg-gray-100/70 outline-none focus:border-2 focus:border-black text-black border border-transparent"
						placeholder="Business Type"
						name="business_type"
						bind:value={data.business_type.value}
						on:input={(e) => checkInput(e.target)}
					/>
					<div
						class="danger text-red-500 text-xl font-bold absolute right-3 top-8 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-exclamation-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M7.002 11a1 1 0 1 1 2 0 1 1 0 0 1-2 0zM7.1 4.995a.905.905 0 1 1 1.8 0l-.35 3.507a.552.552 0 0 1-1.1 0L7.1 4.995z"
							/>
						</svg>
					</div>
					<div
						class="success text-black-500 text-xl font-bold absolute right-3 top-1/2 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-check-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M10.97 4.97a.235.235 0 0 0-.02.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-1.071-1.05z"
							/>
						</svg>
					</div>
					<p class="text-sm text-red-500 msg hidden">Please Enter a valid business type.</p>
				</div>

				<!-- Company name -->
				<div class="flex-col flex py-3 relative">
					<input
						type="text"
						class="p-2 shadow rounded-lg bg-gray-100/70 outline-none focus:border-2 focus:border-black text-black border border-transparent"
						placeholder="Company Name"
						name="company_name"
						bind:value={data.company_name.value}
						on:input={(e) => checkInput(e.target)}
					/>
					<div
						class="danger text-red-500 text-xl font-bold absolute right-3 top-8 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-exclamation-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M7.002 11a1 1 0 1 1 2 0 1 1 0 0 1-2 0zM7.1 4.995a.905.905 0 1 1 1.8 0l-.35 3.507a.552.552 0 0 1-1.1 0L7.1 4.995z"
							/>
						</svg>
					</div>
					<div
						class="success text-black-500 text-xl font-bold absolute right-3 top-1/2 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-check-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M10.97 4.97a.235.235 0 0 0-.02.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-1.071-1.05z"
							/>
						</svg>
					</div>
					<p class="text-sm text-red-500 msg hidden">Please Enter a valid company_name.</p>
				</div>

				<!-- Company type -->
				<div class="flex-col flex py-3 relative">
					<input
						type="text"
						class="p-2 shadow rounded-lg bg-gray-100/70 outline-none focus:border-2 focus:border-black text-black border border-transparent"
						placeholder="Company Type"
						name="company_type"
						bind:value={data.company_type.value}
						on:input={(e) => checkInput(e.target)}
					/>
					<div
						class="danger text-red-500 text-xl font-bold absolute right-3 top-8 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-exclamation-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M7.002 11a1 1 0 1 1 2 0 1 1 0 0 1-2 0zM7.1 4.995a.905.905 0 1 1 1.8 0l-.35 3.507a.552.552 0 0 1-1.1 0L7.1 4.995z"
							/>
						</svg>
					</div>
					<div
						class="success text-black-500 text-xl font-bold absolute right-3 top-1/2 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-check-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M10.97 4.97a.235.235 0 0 0-.02.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-1.071-1.05z"
							/>
						</svg>
					</div>
					<p class="text-sm text-red-500 msg hidden">Please Enter a valid company type.</p>
				</div>

				<!-- Company size -->
				<div class="flex-col flex py-3 relative">
					<!-- <input
            type="text"
            class="p-2 shadow rounded-lg bg-gray-100/70 outline-none focus:border-2 focus:border-black text-black border border-transparent"
            placeholder="company size"
            name="company_size"
            bind:value={data.company_size.value}
            on:input={(e) => checkInput(e.target)}
          /> -->
					<select
						bind:value={data.company_size.value}
						class="p-2 shadow rounded-lg bg-gray-100/70 outline-none focus:border-2 focus:border-black text-black border border-transparent"
					>
						<option disabled>company size</option>
						{#each companies_size as c}
							<option value={c}>{c}</option>
						{/each}
					</select>
					<div
						class="danger text-red-500 text-xl font-bold absolute right-3 top-8 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-exclamation-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M7.002 11a1 1 0 1 1 2 0 1 1 0 0 1-2 0zM7.1 4.995a.905.905 0 1 1 1.8 0l-.35 3.507a.552.552 0 0 1-1.1 0L7.1 4.995z"
							/>
						</svg>
					</div>
					<div
						class="success text-black-500 text-xl font-bold absolute right-3 top-1/2 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-check-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M10.97 4.97a.235.235 0 0 0-.02.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-1.071-1.05z"
							/>
						</svg>
					</div>
					<p class="text-sm text-red-500 msg hidden">Please Enter a valid company size.</p>
				</div>

				<!-- industry -->
				<div class="flex-col flex py-3 relative">
					<!-- <input
            type="text"
            class="p-2 shadow rounded-lg bg-gray-100/70 outline-none focus:border-2 focus:border-black text-black border border-transparent"
            placeholder="industry"
            name="industry"
            bind:value={data.industry.value}
            on:input={(e) => checkInput(e.target)}
          /> -->
					<select
						bind:value={data.industry.value}
						class="p-2 shadow rounded-lg bg-gray-100/70 outline-none focus:border-2 focus:border-black text-black border border-transparent"
					>
						<option disabled>industry</option>
						{#each industries as c}
							<option value={c}>{c}</option>
						{/each}
					</select>
					<div
						class="danger text-red-500 text-xl font-bold absolute right-3 top-8 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-exclamation-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M7.002 11a1 1 0 1 1 2 0 1 1 0 0 1-2 0zM7.1 4.995a.905.905 0 1 1 1.8 0l-.35 3.507a.552.552 0 0 1-1.1 0L7.1 4.995z"
							/>
						</svg>
					</div>
					<div
						class="success text-black-500 text-xl font-bold absolute right-3 top-1/2 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-check-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M10.97 4.97a.235.235 0 0 0-.02.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-1.071-1.05z"
							/>
						</svg>
					</div>
					<p class="text-sm text-red-500 msg hidden">Please Enter a valid industry.</p>
				</div>

				<!-- country/region -->
				<div class="flex-col flex py-3 relative">
					<select
						bind:value={data.country.value}
						class="p-2 shadow rounded-lg bg-gray-100/70 outline-none focus:border-2 focus:border-black text-black border border-transparent"
					>
						<option disabled>country</option>
						{#each countries as c}
							<option value={c.name}>{c.name}</option>
						{/each}
					</select>
					<div
						class="danger text-red-500 text-xl font-bold absolute right-3 top-8 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-exclamation-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M7.002 11a1 1 0 1 1 2 0 1 1 0 0 1-2 0zM7.1 4.995a.905.905 0 1 1 1.8 0l-.35 3.507a.552.552 0 0 1-1.1 0L7.1 4.995z"
							/>
						</svg>
					</div>
					<div
						class="success text-black-500 text-xl font-bold absolute right-3 top-1/2 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-check-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M10.97 4.97a.235.235 0 0 0-.02.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-1.071-1.05z"
							/>
						</svg>
					</div>
					<p class="text-sm text-red-500 msg hidden">Please Enter a valid country.</p>
				</div>

				<!-- use case -->
				<div class="flex-col flex py-3 relative">
					<input
						type="text"
						class="p-2 shadow rounded-lg bg-gray-100/70 outline-none focus:border-2 focus:border-black text-black border border-transparent"
						placeholder="Use Case"
						name="use_case"
						bind:value={data.use_case.value}
						on:input={(e) => checkInput(e.target)}
					/>
					<div
						class="danger text-red-500 text-xl font-bold absolute right-3 top-8 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-exclamation-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M7.002 11a1 1 0 1 1 2 0 1 1 0 0 1-2 0zM7.1 4.995a.905.905 0 1 1 1.8 0l-.35 3.507a.552.552 0 0 1-1.1 0L7.1 4.995z"
							/>
						</svg>
					</div>
					<div
						class="success text-black-500 text-xl font-bold absolute right-3 top-1/2 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-check-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M10.97 4.97a.235.235 0 0 0-.02.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-1.071-1.05z"
							/>
						</svg>
					</div>
					<p class="text-sm text-red-500 msg hidden">Please Enter a valid country.</p>
				</div>

				<!-- Message -->
				<div class="py-3 relative col-start-1 col-end-3">
					<textarea
						class="w-full h-32 mt-2 p-3 rounded-lg bg-gray-100/70 outline-none focus:border-2 focus:border-black text-black border border-transparent"
						placeholder="Project Description"
						name="message"
						bind:value={data.message.value}
						on:input={(e) => checkInput(e.target)}
					/>
					<div
						class="danger text-red-500 text-xl font-bold absolute right-3 top-10 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-exclamation-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M7.002 11a1 1 0 1 1 2 0 1 1 0 0 1-2 0zM7.1 4.995a.905.905 0 1 1 1.8 0l-.35 3.507a.552.552 0 0 1-1.1 0L7.1 4.995z"
							/>
						</svg>
					</div>
					<div
						class="success text-black-500 text-xl font-bold absolute right-3 top-10 -translate-y-1/2 hidden"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-check-circle"
							viewBox="0 0 16 16"
						>
							<path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
							<path
								d="M10.97 4.97a.235.235 0 0 0-.02.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-1.071-1.05z"
							/>
						</svg>
					</div>
					<p class="text-sm text-red-500 msg hidden">Please Enter a valid Message.</p>
				</div>

				<div class="py-3 col-start-1 col-end-3">
					{#if msgSentSuccessfully}
						<div class="bg-black text-white p-4 text-center font-bold rounded-lg uppercase">
							<p>request sent</p>
						</div>
					{:else}
						<button
							on:click={(e) => sendMessage(e)}
							class="p-3 bg-black text-white font-bold w-full rounded-lg uppercase hover:bg-light-black disabled:cursor-not-allowed"
							disabled={!dataValid}>send</button
						>
					{/if}
				</div>
			</form>
		</div>
	</div>
{/if}
