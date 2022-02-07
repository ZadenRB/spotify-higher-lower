<script lang="ts">
	export let item: Song;
	export let popularityDescription: string;
</script>

<div class="item">
	<div class="item-content">
		<img class="image" src={item.image} alt={item.name} />
        <div class="name-wrapper">
            <span class="name scrolling">{item.name}</span>
        </div>
        <div class="artists-wrapper">
            <span class="artists">{item.artists}</span>
        </div>
		<div class="popularity scrolling">{item.popularity} {popularityDescription}</div>
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
