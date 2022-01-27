<script lang="ts">
    export let item: ValuedObject;
    export let comparedName: string;
    export let shortValueDescription: string;
</script>

<div class="item">
    <div class="item-content">
        <div class="keyword">
            <span class="name">{item.name}</span>
            <img class="image" src={item.image} alt={item.name} />
            <span>has a</span>
        </div>
        <div class="button-group">
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
            <span>{shortValueDescription} than {comparedName}</span>
        </div>
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
