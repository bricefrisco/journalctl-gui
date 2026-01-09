<script>
	import TerminalIcon from './icons/TerminalIcon.svelte';
	import CalendarIcon from './icons/CalendarIcon.svelte';
	import SunIcon from './icons/SunIcon.svelte';
	import MoonIcon from './icons/MoonIcon.svelte';

	let { theme = 'dark', onToggle } = $props();
</script>

<div class="bg-white dark:bg-gray-800 border-b border-gray-200 dark:border-gray-700 shadow-sm">
	<div class="px-6 py-4">
		<div class="flex items-center justify-between">
			<div class="flex items-center gap-3">
				<TerminalIcon class="size-8 text-blue-600 dark:text-blue-400" />
				<div>
					<h1 class="text-2xl font-bold text-gray-900 dark:text-white">journalctl-gui</h1>
					<p class="text-sm text-gray-600 dark:text-gray-400">Graphical user interface for journalctl</p>
				</div>
			</div>
			<div class="flex items-center gap-4">
				<div class="flex items-center gap-2 text-sm text-gray-600 dark:text-gray-400">
					<CalendarIcon class="size-4" />
					<span>Last updated: {new Date().toLocaleTimeString()}</span>
				</div>
				<button
					onclick={() => onToggle(theme === 'dark' ? 'light' : 'dark')}
					class="p-2 rounded-lg bg-gray-100 dark:bg-gray-700 hover:bg-gray-200 dark:hover:bg-gray-600 transition-colors"
					title={theme === 'dark' ? 'Switch to light mode' : 'Switch to dark mode'}
				>
					{#if theme === 'dark'}
						<SunIcon class="size-5 text-yellow-500" />
					{:else}
						<MoonIcon class="size-5 text-gray-700" />
					{/if}
				</button>
			</div>
		</div>
	</div>
</div>
