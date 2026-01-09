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

<select bind:value={selected}>
	<option value="" disabled>Select a service</option>
	{#each services as service (service)}
		<option value={service}>{service}</option>
	{/each}
</select>