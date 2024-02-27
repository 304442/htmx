<script>
	import Quote from '../components/Quote.svelte';
	import Accordion from '../components/Accordion.svelte';
	import * as Icons from '../components/icons/index.js';

	import Booking from '../components/booking';

	import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query';

	const queryClient = new QueryClient();
</script>

<QueryClientProvider client={queryClient}>
	<div class="flex-block max-w-xl p-4 md:pd-8 mx-auto mt-4">
		<div class="bg-cool-gray-0.5 text-light-black shadow overflow-hidden rounded-lg">
			<div class="px-4 py-5 sm:px-6 flex items-center justify-between bg-white flex-wrap">
				<div><img class="rounded-lg w-14" src="./hcc.svg" alt="48D1CC" /></div>
				<div class="ml-4 hidden">
					<h1 class="text-l leading-6 font-bold text-gray-900">48D1CC</h1>
					<h2 class="text-sm leading-6 text-gray-500 -mt-1">AWS, AZURE, GCP | IaaS, PaaS, SaaS</h2>
				</div>
				<div class="flex space-x-4">
					<a href="auth" class=" text-light-black px-4 py-2 rounded uppercase font-semibold"
						>Log in</a
					>
					<a
						href="auth"
						class="bg-light-black text-white px-4 py-2 rounded-full hover:shadow-xl duration-200 uppercase font-semibold"
						>Sign up</a
					>
				</div>
			</div>
			<div class="border-t border-gray-200 bg-cool-gray-0.5">
				<div class="bg-blue px-4 py-5 w-full flex space-x-0.5 bg-light-wash">
					<div class="mt-1 text-sm text-gray">
						<b class="text-green">Services:</b> Big Data, Analytics, Machine Learning, DevOps, IoT,
						Data Science <br />
						<b class="text-green">Solutions:</b> Migrations, Application Modernization, Digital Transformation,
						Cost Optimization, Training and Certification
					</div>
				</div>
				<Accordion Icon={Icons.BlogIcon} text="Blog"><!-- Blog Content Goes Here --></Accordion>
				<Accordion Icon={Icons.CloudIcon} text="Solutions Marketplace"></Accordion>
				<Accordion Icon={Icons.CalendarIcon} text="Book a Free Consultation">
					<Booking />
				</Accordion>
				<Quote />
			</div>
		</div>
	</div>
</QueryClientProvider>
