<script>
	import bproc_logo from '$lib/assets/bproc_logo.svg';
</script>

<header class="bg-gray-100 px-40">
	<nav class="flex items-center justify-between py-10">
		<a href="/">
			<div>
				<img src={bproc_logo} alt="bproc-logo" class="h-10">
			</div>
		</a>
		<div class="flex flex-wrap gap-10 font-semibold text-lg">
			<a href="/" class="hover:text-green-700 hover:scale-110">Home</a>
			<a href="/settings" class="hover:text-green-700 hover:scale-110">Settings</a>
		</div>
	</nav>
</header>