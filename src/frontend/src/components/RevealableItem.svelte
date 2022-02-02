<script lang="ts">
	import { fly } from "svelte/transition";

	export let item: Song;
	export let reveal: () => void;
	export let afterReveal: (b: boolean) => void;
	export let popularityDescription: string;
	export let showValue: boolean;

	let higher: boolean;

	const clicked = (b: boolean) => {
		higher = b;
		reveal();
	};

	const transitionProperties = { y: 50, duration: 500 };
</script>

<div class="item">
	<div class="item-content">
		<img class="image" src={item.image} alt={item.name} />
        <div class="name-wrapper">
            <span class="name">{item.name}</span>
        </div>
        <div class="artists-wrapper">
            <span class="artists">{item.artists}</span>
        </div>
		{#if !showValue}
			<div class="buttons">
				<div
					class="button"
					tabindex="0"
					on:click={() => clicked(true)}
					on:keypress={(event) => {
						if (event.key == "Enter") {
							clicked(true);
						}
					}}
				>
					<span class="button-text">Higher</span>
					<span class="button-arrow">&#9650;</span>
				</div>
				<div
					class="button"
					tabindex="0"
					on:click={() => clicked(false)}
					on:keypress={(event) => {
						if (event.key == "Enter") {
							clicked(false);
						}
					}}
				>
					<span class="button-text">Lower</span>
					<span class="button-arrow">&#9660;</span>
				</div>
			</div>
		{:else}
			<div
				class="popularity"
				in:fly={transitionProperties}
				on:introend={() => afterReveal(higher)}
			>
				{item.popularity}
				{popularityDescription}
			</div>
		{/if}
	</div>
</div>

<style>
	.item {
		border-left: 1px solid white;
	}

	@media (max-width: 894px) {
		.item {
			border-top: 1px solid white;
			border-left: 0;
		}
	}
</style>
