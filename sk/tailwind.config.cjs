/** @type {import('tailwindcss').Config} */
const config = {
	content: [
		'./src/**/*.{html,js,svelte,ts}',
		'./node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}'
	],

	plugins: [require('flowbite/plugin')],

	darkMode: 'class',

	theme: {
		extend: {
			colors: {
				white: '#FFFFFF',
				ghost: '#F7F7F7',
				ivory: '#F6F6F4',
				'light-wash': '#F7FBFB',
				'medium-wash': '#EFF7F8',
				'dark-wash': '#E6F3F3',
				'light-warm-gray': '#EEEDEE',
				'dark-warm-gray': '#A3ACB0',
				'cool-gray-0.5': '#EDEEF1',
				'cool-gray-1': '#C5C5D2',
				'cool-gray-1.5': '#AEAEBB',
				'cool-gray-2': '#8E8EA0',
				'cool-gray-3': '#6E6E80',
				'cool-gray-4': '#404452',
				'light-black': '#191927',
				'medium-black': '#0E0E1A',
				black: '#000000',
				// flowbite-svelte
				primary: {
					50: '#eff6ff',
					100: '#dbeafe',
					200: '#bfdbfe',
					300: '#93c5fd',
					400: '#60a5fa',
					500: '#3b82f6',
					600: '#2563eb',
					700: '#1d4ed8',
					800: '#1e40af',
					900: '#1e3a8a'
				}
			}
		}
	}
};

module.exports = config;
// export default {
// 	content: ['./src/**/*.{html,js,svelte,ts}'],
// 	theme: {
// 		extend: {
// 			colors: {
// 				white: '#FFFFFF',
// 				ghost: '#F7F7F7',
// 				ivory: '#F6F6F4',
// 				'light-wash': '#F7FBFB',
// 				'medium-wash': '#EFF7F8',
// 				'dark-wash': '#E6F3F3',
// 				'light-warm-gray': '#EEEDEE',
// 				'dark-warm-gray': '#A3ACB0',
// 				'cool-gray-0.5': '#EDEEF1',
// 				'cool-gray-1': '#C5C5D2',
// 				'cool-gray-1.5': '#AEAEBB',
// 				'cool-gray-2': '#8E8EA0',
// 				'cool-gray-3': '#6E6E80',
// 				'cool-gray-4': '#404452',
// 				'light-black': '#191927',
// 				'medium-black': '#0E0E1A',
// 				black: '#000000'
// 			}
// 		}
// 	},
// 	plugins: []
// };
