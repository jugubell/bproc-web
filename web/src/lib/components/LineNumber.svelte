<script>
	let { lineNumber, highlight } = $props()
	let i;

	function highlightLine(line) {
		return line === highlight ? "highlight" : "";
	}
</script>

<div class="flex flex-col gap-0 pt-2 p-1 items-center justify-start font-mono">
	{#each { length: 10 } as _, i}
		<div class="lineNumber {highlightLine(i+1)}">{i+1}</div>
	{/each}
	{#if lineNumber > 10}
		{#each { length: lineNumber-10 } as _, i}
			<div class="lineNumber {highlightLine(i+11)}">{i+11}</div>
		{/each}
	{/if}

</div>

<style>
	@import 'tailwindcss';
	.highlight {
			background-color: rgb(154, 200, 154);
			color: white;
	}
	.lineNumber {
			@apply
      w-full
			text-right
			pr-1
			pl-5;
	}
</style>