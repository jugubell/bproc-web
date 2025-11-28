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
	import { editorDataStore } from '$lib/state.svelte.js';
	import LineNumber from '$lib/components/LineNumber.svelte';
	import { onMount, tick } from 'svelte';
	import { apiGet } from '$lib/index.js';

	let asmProgram = $state("");
	let programLine = $state(0);
	let cursorLineNumber = $state(0);

	function update() {
		window.localStorage.setItem("asmProgram", asmProgram);
		programLine = getLineNumber();
		cursorLineNumber = getCursorLineNumber();
		updateCursor();
		updateTextareaHeight();
	}

	function getLineNumber() {
		let textareaVal = document.getElementById("progArea").value;
		return (textareaVal === null || textareaVal === "") ? 0 : textareaVal.split(/\n/).length;
	}

	function getCursorLineNumber() {
		let textarea = document.getElementById("progArea");
		const textToCusrsor = textarea.value.slice(0, textarea.selectionStart);
		return textToCusrsor.split(/\n/).length;
	}

	function updateCursor() {
		cursorLineNumber = getCursorLineNumber();
	}

	async function readStoredProgram() {
		let storedProgram = window.localStorage.getItem("asmProgram");
		if (storedProgram) {
			asmProgram = storedProgram;
		} else {
			const progEx = await getProgramExample();
			if(progEx && progEx !== null) {
				asmProgram = progEx;
				update();
			}
		}
	}

	function updateTextareaHeight() {
		let textarea = document.getElementById("progArea");
		textarea.style.height = textarea.scrollHeight + "px";
	}

	function loadCompileType() {
		const compileType = window.localStorage.getItem("compileType");
		const inputGroup = document.getElementById("compileTypeRadio");
		for(const inputDiv of inputGroup.querySelectorAll("div")) {
			const input = inputDiv.querySelector("input");
			input.checked = input && (input.id.toLowerCase() === `cmp${compileType.toLowerCase()}`);
		}
	}

	async function getProgramExample() {
		const res = await apiGet('example');
		if(res.status === 200) {
			return res.data.message;
		} else {
			editorDataStore.update(state => ({
				...state,
				consoleValue: `Unable to get program example.\n${res.message}\n`,
			}))
			return null;
		}
	}

	onMount(async () => {
		await readStoredProgram();
		await tick();
		programLine = getLineNumber();
		cursorLineNumber = getCursorLineNumber();
		updateTextareaHeight();
		loadCompileType();
	})

</script>

<div class="flex w-full h-screen gap-x-2">
	<div class="programInput w-1/2 flex flex-row gap-1 px-1">
		<LineNumber lineNumber={programLine} highlight={cursorLineNumber} />
		<textarea id="progArea" oninput={update} bind:value={asmProgram} onmouseup={updateCursor} onkeyup={updateCursor}></textarea>
	</div>
	<pre class="programOutput w-1/2 px-5">{@html $editorDataStore.consoleValue}></pre>
</div>

<style>
	@import "tailwindcss";
	.programInput, .programOutput {
			@apply
      border-2
			rounded-xl
			py-5
			border-green-500
      focus:outline
      focus:outline-green-200
      focus:shadow-green-800
      focus:shadow-2xl
      focus:inset-shadow-sm
      focus:inset-shadow-gray-200
			min-w-0
			overflow-auto
			;
	}
	.programInput::-webkit-scrollbar {
			display: none;
			width: 0;
			height: 0;
	}

	textarea {
			@apply
			border
			rounded-md
      w-full
			h-full
			p-2
		border-green-500
		focus:outline-green-200
			font-mono
			overflow-y-hidden
			resize-none;
	}
</style>