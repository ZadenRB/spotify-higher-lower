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
        font-size: 6vh;
    }

    .image {
        width: min(38vw, 40vh);
        height: min(38vw, 40vh);
        max-width: 800px;
        max-height: 800px;
        border: 1px solid white;
    }

    .button-group {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: space-evenly;
    }

    .buttons {
        height: 100px;

        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: space-evenly;
    }

    .button {
        width: 75px;
        padding: 10px 20px;

        display: flex;
        justify-content: space-between;

        border: 1px solid white;
        border-radius: 20px;

        cursor: pointer;

        transition: background-color 150ms ease;
    }

    .button-text {
        color: #1db964;
    }

    .button-text,
    .button-arrow {
        transition: color 150ms ease;
    }

    .button:hover {
        background-color: white;
    }

    .button:hover > .button-text {
        color: #191414;
    }

    .button:hover > .button-arrow {
        color: #191414;
    }

    @media (min-width: 895px) {
        .button-group {
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
        .name {
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
