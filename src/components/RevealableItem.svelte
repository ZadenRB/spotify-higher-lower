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
		width: 50vw;
		height: 100vh;
		overflow: hidden;

		display: flex;
		align-items: center;
		justify-content: center;

		padding: 0 calc(min(10vw, 120px) / 2);
		box-sizing: border-box;

		border-left: 1px solid white;
	}

	.item-content {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: space-evenly;
	}

	.keyword {
		display: flex;
		align-items: center;
		justify-content: flex-end;
		flex-direction: column;
	}

	.name,
	.value {
		font-size: 6vh;
	}

	.name {
		text-align: center;
		line-clamp: 2;
		overflow: hidden;
		display: -moz-box;
		display: -webkit-box;
		-moz-box-orient: vertical;
		-webkit-box-orient: vertical;
		-webkit-line-clamp: 2;
		max-width: 40vw;
	}

	.image {
		width: min(38vw, 40vh);
		height: min(38vw, 40vh);
		max-width: 800px;
		max-height: 800px;
		border: 1px solid white;
	}

	.value-group,
	.button-group {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: space-evenly;
	}

	.value {
		color: #1db964;
	}

	.button {
		width: 75px;

		display: flex;
		justify-content: space-between;
	}


	@media (min-width: 895px) {
		.button-group,
		.value-group {
			height: 140px;
		}
		.keyword {
			height: calc(min(38vw, 40vh) + 15vh + 20px);
		}
	}

	@media (max-width: 894px) {
		.item {
			width: 100vw;
			height: 50vh;
			overflow: hidden;

			padding: calc(min(6vh, 120px) / 2) 0;

			border-top: 1px solid white;
			border-left: 0;
		}
		.image {
			width: 25vh;
			height: 25vh;
		}
		.name,
		.value {
			font-size: 5vw;
		}
		.buttons {
			height: 42px;
			width: 40vw;
			min-width: 264px;

			display: flex;
			flex-direction: row;
			align-items: center;
			justify-content: space-evenly;
		}
	}
</style>
