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
	import { FileCheck, Binary, FileText } from '@lucide/svelte';

</script>

<div class="flex flex-row justify-start items-center w-full my-2 p-2 gap-x-3 border rounded-xl border-green-500 bg-green-50">
		<button><FileCheck />Verify</button>
		<button><Binary />Compile</button>
		<button><FileText />Generate File</button>
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
</style>