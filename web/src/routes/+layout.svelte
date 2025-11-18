<script>
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import Header from '$lib/components/Header.svelte';
	import Footer from '$lib/components/Footer.svelte';

</script>

<svelte:head>
	<link rel="icon" href={favicon} />
	<title>BProC Web</title>
</svelte:head>

<Header />

<main class="mx-40 mt-10">
	<slot />
</main>

<Footer />


