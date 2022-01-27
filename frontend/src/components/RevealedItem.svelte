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
		border-right: 1px solid white;
	}

	@media (max-width: 894px) {
		.item {
			border-bottom: 1px solid white;
			border-right: 0;
		}
	}
</style>
