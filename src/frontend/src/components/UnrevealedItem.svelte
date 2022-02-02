<script lang="ts">
    export let item: Song;
</script>

<div class="item">
    <div class="item-content">
        <img class="image" src={item.image} alt={item.name} />
        <div class="name">{item.name}</div>
        <div class="artists">{item.artists}</div>
        <div class="buttons">
            <div class="button">
                <span class="button-text">Higher</span>
                <span class="button-arrow">&#9650;</span>
            </div>
            <div class="button">
                <span class="button-text">Lower</span>
                <span class="button-arrow">&#9660;</span>
            </div>
        </div>
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
