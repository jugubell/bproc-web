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