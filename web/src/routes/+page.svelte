<script>
	import Header from '$lib/components/Header.svelte';
	import Filters from '$lib/components/Filters.svelte';

	let theme = 'dark'
	document.documentElement.setAttribute('data-theme', theme);

	function toggleTheme() {
		theme = theme === 'dark' ? 'light' : 'dark';
		document.documentElement.setAttribute('data-theme', theme);
	}
</script>

<div class="min-h-screen {theme === 'dark' ? 'dark bg-gray-900' : 'bg-gray-100'}">
	<Header {theme} onToggle={toggleTheme} />
	<Filters />
	<div class="p-4">
		<h1 class="dark:text-gray-500">Welcome to SvelteKit</h1>
		<p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>
	</div>
</div>
