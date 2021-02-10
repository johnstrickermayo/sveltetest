<script lang="ts">
	import functionTest from './functionTest'
	export let name: string;

	let x = functionTest()
	console.log("dfsdfdsf:")
	console.log(x)
</script>

<main>
	<h1>blurb {name}!</h1>
	<p> the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn how to build Svelte apps.</p>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>