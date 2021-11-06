<script lang="ts">
	import { fly } from "svelte/transition";

	export let item: ValuedObject;
	export let comparedName: string;
	export let reveal: () => void;
	export let afterReveal: (b: boolean) => void;
	export let valueDescription: string;
	export let shortValueDescription: string;
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
		<div class="keyword">
			<span class="name">{item.name}</span>
			<img class="image" src={item.image} alt={item.name} />
			<span>has a</span>
		</div>
		{#if !showValue}
			<div class="button-group">
				<div class="buttons">
					<div class="button" on:click={() => clicked(true)}>
						<span class="button-text">Higher</span>
						<span class="button-arrow">&#9650;</span>
					</div>
					<div class="button" on:click={() => clicked(false)}>
						<span class="button-text">Lower</span>
						<span class="button-arrow">&#9660;</span>
					</div>
				</div>
				<span>{shortValueDescription} than {comparedName}</span>
			</div>
		{:else}
			<div
				class="value-group"
				in:fly={transitionProperties}
				on:introend={() => afterReveal(higher)}
			>
				<span class="value">{item.value}</span>
				<span>{valueDescription}</span>
			</div>
		{/if}
	</div>
</div>

<style>
	.item {
		border-left: 1px solid white;
	}

	.button {
		width: 75px;

		display: flex;
		justify-content: space-between;
	}

	@media (max-width: 894px) {
		.item {
			border-top: 1px solid white;
			border-left: 0;
		}
	}
</style>
