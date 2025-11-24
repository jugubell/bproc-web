<!--
  - File: +page.svelte
  - Project: bproc-web
  - Last modified: 2025-11-18 23:11
  -
  - This file: +page.svelte is part of BProC-WEB project.
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
	import Toolbar from '$lib/components/Toolbar.svelte';
	import PageTitle from '$lib/components/PageTitle.svelte';
	import ProgramEditor from '$lib/components/ProgramEditor.svelte';
	import { setContext } from 'svelte';

	let cliHelp = $state({
		message: ''
	});

	setContext('consoleDisplay', cliHelp.message);
</script>

<PageTitle title="BProC Compiler" />
<Toolbar />
<ProgramEditor />
