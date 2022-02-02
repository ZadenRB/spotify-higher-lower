<script lang="ts">
	import Game from "./screens/Game.svelte";
	import GameOver from "./screens/GameOver.svelte";
	import Menu from "./screens/Menu.svelte";

	export let popularityDescription: string;
	export let requestURI: string;

	let current: Song;
	let atBat: Song;
	let onDeck: Song;

	let gameOver = false;
	let loading = false;
	let playing = false;

	const toMenu = () => {
		playing = false;
		gameOver = false;
	}

	const onPlay = async () => {
		loading = true;
		const currentRes = await fetch(requestURI);
		const atBatRes = await fetch(requestURI);
		const onDeckRes = await fetch(requestURI);
		current = await currentRes.json();
		atBat = await atBatRes.json();
		onDeck = await onDeckRes.json();
		playing = true;
		gameOver = false;
		loading = false;
	};

	const onEnd = () => {
		gameOver = true;
		current = null;
		atBat = null;
		onDeck = null;
	}
</script>

{#if gameOver}
	<GameOver {loading} restart={onPlay} back={toMenu}/>
{:else if playing && current && atBat && onDeck}
	<Game
		{popularityDescription}
		{requestURI}
		{current}
		{atBat}
		{onDeck}
		{onEnd}
	/>
{:else if !playing}
	<Menu {loading} {onPlay}/>
{/if}
