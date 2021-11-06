<script lang="ts">
	export let item: ValuedObject;
	export let valueDescription: string;
</script>

<div class="item">
	<div class="item-content">
		<div class="keyword">
			<span class="name">{item.name}</span>
			<img class="image" src={item.image} alt={item.name} />
			<span>has a</span>
		</div>
		<div class="value-group">
			<span class="value">{item.value}</span>
			<span>{valueDescription}</span>
		</div>
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

		border-right: 1px solid white;
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

	.value-group {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: space-evenly;
	}

	.value {
		color: #1db964;
	}

	@media (min-width: 895px) {
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

			border-bottom: 1px solid white;
			border-right: 0;
		}
		.image {
			width: 25vh;
			height: 25vh;
		}
		.name,
		.value {
			font-size: 5vw;
		}
	}
</style>
