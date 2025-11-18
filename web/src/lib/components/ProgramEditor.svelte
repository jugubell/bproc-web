<script>
	import { onMount } from 'svelte';

	let asmProgram = $state("");

	function update() {
		window.localStorage.setItem("asmProgram", asmProgram);
	}

	function readStoredProgram() {
		let storedProgram = window.localStorage.getItem("asmProgram");
		if (storedProgram) {
			asmProgram = storedProgram;
		}
	}

	onMount(() => {
		readStoredProgram();
	})

</script>

<div class="flex w-full h-screen gap-x-2">
	<textarea name="programInput" class="w-2/3" oninput={update} bind:value={asmProgram}></textarea>
	<textarea name="programOutput" class="w-1/3" disabled></textarea>
</div>

<style>
	@import "tailwindcss";
	textarea {
			@apply
      border-2
			rounded-xl
			p-5
			border-green-500
      focus:outline
      focus:outline-green-200
      focus:shadow-green-800
      focus:shadow-2xl
      focus:inset-shadow-sm
      focus:inset-shadow-gray-200
			min-w-0
			resize-y
			overflow-auto
			;
	}
</style>