<!--
  - File: Header.svelte
  - Project: bproc-web
  - Last modified: 2025-11-18 23:11
  -
  - This file: Header.svelte is part of BProC-WEB project.
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
	import bproc_logo from '$lib/assets/bproc_logo.svg';
	import { Settings, Home } from '@lucide/svelte';
</script>

<header class="bg-gray-100 px-40">
	<nav class="flex items-center justify-between py-10">
		<a href="/">
			<div>
				<img src={bproc_logo} alt="bproc-logo" class="h-10">
			</div>
		</a>
		<div class="flex flex-wrap gap-10 font-semibold text-lg">
			<a href="/" class="hover:text-green-700 hover:scale-110"><span class="flex flex-row gap-1 items-center"><Home />Home</span></a>
			<a href="/settings" class="hover:text-green-700 hover:scale-110"><span class="flex flex-row gap-1 items-center"><Settings />Settings</span></a>
		</div>
	</nav>
</header>