<!--
  - File: ProgramEditor.svelte
  - Project: bproc-web
  - Last modified: 2025-11-18 23:11
  -
  - This file: ProgramEditor.svelte is part of BProC-WEB project.
  -
  - BProC-WEB is free software: you can redistribute it and/or modify it
  - under the terms of the GNU General Public License as published
  - by the Free Software Foundation, either version 2 of the License,
  - or (at your option) any later version.
  -
  - BProC-WEB is distributed in the hope that it will be useful,
  - but WITHOUT ANY WARRANTY; without even the implied warranty
  - of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
  - See the GNU General Public License for more details.
  -
  - You should have received a copy of the GNU General Public License
  - along with BProC-WEB. If not, see <https://www.gnu.org/licenses/>.
  -
  - Copyright (C) 2025 Jugurtha Bellagh
  -->

<script>
	import { onMount } from 'svelte';
	import { editorDataStore } from '$lib/state.svelte.js';
	import { derived } from 'svelte/store';

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

	console.log($editorDataStore.consoleValue);

</script>

<div class="flex w-full h-screen gap-x-2">
	<textarea id="progArea" class="programInput w-1/2" oninput={update} bind:value={asmProgram}></textarea>
	<pre class="programOutput w-1/2">{@html $editorDataStore.consoleValue}></pre>
</div>

<style>
	@import "tailwindcss";
	.programInput, .programOutput {
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