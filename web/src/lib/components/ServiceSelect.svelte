<script>
	import { onMount } from 'svelte';

	let services = [];
	let selected = '';

	onMount(async () => {
		const res = await fetch('http://192.168.1.160:8091/api/services');
		if (!res.ok) throw new Error('Failed to load services');
		services = await res.json();
	})
</script>

<select
	bind:value={selected}
	class="px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 min-w-[200px] max-w-[250px]"
>
	<option value="" disabled>Select a service</option>
	{#each services as service (service)}
		<option value={service}>{service}</option>
	{/each}
</select>