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