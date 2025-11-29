<!--
  - File: Toolbar.svelte
  - Project: bproc-web
  - Last modified: 2025-11-18 23:11
  -
  - This file: Toolbar.svelte is part of BProC-WEB project.
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
	import { FileCheck, Binary, FileText, CircleQuestionMark, List } from '@lucide/svelte';
	import { apiGet, apiPost, getCompileType, notImplemented, getAsmProgram } from '$lib/index.js';
	import { editorDataStore } from '$lib/state.svelte.js';
	import { onMount } from 'svelte';

	async function getAndStore(str) {
		let res = await apiGet(`${str}`)
		if (res.status === 200) {
			let apiRes = res.data.message;
			console.log(`Data received ${apiRes}`);
			editorDataStore.update(state => ({
				...state,
				consoleValue: `${apiRes}\n`,
			}));
		} else {
			editorDataStore.update(state => ({
				...state,
				consoleValue: `${res.message}\n`,
			}))
		}
	}

	function storeCompileType() {
		window.localStorage.setItem("compileType", getCompileType());
	}

	async function post(url) {
		let data = {
			program: getAsmProgram(),
			type: url === "verify" ? null : getCompileType()
		}

		console.log(data.program.value)

		if(data.program && data.program !== "") {
			let res = await apiPost(url, data);
			if (res.status === 200) {
				let apiRes = res.data.message;
				console.log(`Data received ${apiRes}`);
				editorDataStore.update(state => ({
					...state,
					consoleValue: `${preprocess(apiRes)}\n`,
				}));
			} else {
				editorDataStore.update(state => ({
					...state,
					consoleValue: `${preprocess(res.message)}\n`,
				}))
			}
		} else {
			editorDataStore.update(state => ({
				...state,
				consoleValue: "<span style=\"color: #fb8456;\">Program empty, write some code, then try again</span>\n",
			}))
		}
	}



	function preprocess(msg) {
		msg = msg.replace(/\u001B\[31m/g, '<span style="color: #da4141;">');
		msg = msg.replace(/\u001B\[32m/g, '<span style="color: #2ca12c;">');
		msg = msg.replace(/\u001B\[33m/g, '<span style="color: #fb8456">;">');
		msg = msg.replace(/\u001B\[0m/g, '</span>');
		return msg;
	}

	onMount(() => {
		console.log(getCompileType());
	})

</script>

<div class="flex flex-row justify-between items-center w-full my-2 p-2 gap-x-3 border rounded-xl border-green-500 bg-green-50">
	<div class="flex flex-row justify-start items-center">
		<button onclick="{() => post('verify')}"><FileCheck />Verify</button>
		<button onclick="{() => post('compile')}"><Binary />Compile</button>
		<button onclick="{() => notImplemented()}"><FileText />Generate File</button>
	</div>
	<div id="compileTypeRadio" class="flex flex-row justify-start items-center gap-x-3">
		<div class="flex items-center">
			<input type="radio"	name="compileType" id="cmpBin" checked onchange={storeCompileType}><label for="cmpBin">Binary</label>
		</div>
		<div class="flex items-center">
			<input type="radio"	name="compileType" id="cmpHex" onchange={storeCompileType}><label for="cmpHex">Hex</label>
		</div>
		<div class="flex items-center">
			<input type="radio"	name="compileType" id="cmpHexV3" onchange={storeCompileType}><label for="cmpHexV3">Hex V3</label>
		</div>
		<div class="flex items-center">
			<input type="radio"	name="compileType" id="cmpVhdl" onchange={storeCompileType}><label for="cmpVhdl">VHDL</label>
		</div>
		<div class="flex items-center">
			<input type="radio"	name="compileType" id="cmpVrlg" onchange={storeCompileType}><label for="cmpVrlg">Verilog</label>
		</div>
	</div>
	<div class="flex flex-row justify-start items-center">
		<button onclick={() => getAndStore('is')}><List />Instructions</button>
		<button onclick={() => getAndStore('help')}><CircleQuestionMark />Help</button>
	</div>
</div>

<style>
	@import 'tailwindcss';
	button {
		@apply
      hover:outline-green-400
      hover:outline
      hover:bg-green-100
      hover:text-green-950
      hover:shadow-xl
      hover:inset-shadow-sm
      hover:inset-shadow-gray-100
      rounded-md
      py-1
      px-2
      flex
      flex-row
      gap-1
      items-center
      font-semibold
      text-gray-600;
	}
	label {
			@apply
			text-gray-600
			font-semibold
			pl-2;
	}
	input[type="radio"] {
			@apply
			appearance-none
			bg-transparent
			checked:bg-green-600
			rounded-sm
			outline-offset-2
			outline-1
			outline-green-500
			h-4
			w-4;
	}
</style>