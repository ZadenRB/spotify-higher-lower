<script lang="ts">
    import { scale, fade } from "svelte/transition";
    import { tweened } from "svelte/motion";
    import { writable } from "svelte/store";

    import RevealableItem from "../components/RevealableItem.svelte";
    import UnrevealedItem from "../components/UnrevealedItem.svelte";
    import RevealedItem from "../components/RevealedItem.svelte";

    export let valueDescription: string;
    export let shortValueDescription: string;
    export let requestURI: string;
    export let current: ValuedObject;
    export let atBat: ValuedObject;
    export let onDeck: ValuedObject;

    export let onEnd: () => void;

    const tween = {
        duration: 250,
    };
    const noTween = {
        duration: 0,
    };

    let score = 0;

    let showStatus = true;
    let status: number = 0;

    let highScore = parseInt(localStorage.getItem("highScore")) || 0;
    const highScoreStore = writable(highScore);
    highScoreStore.subscribe((value) => {
        highScore = value;
        localStorage.setItem("highScore", value.toString());
    });

    let dist = 0;
    const distStore = tweened(dist, tween);
    distStore.subscribe((value) => {
        dist = value;
    });

    let showValue = false;

    const reveal = () => {
        showValue = true;
    };

    const afterReveal = (higher: boolean) => {
        if (
            (higher && atBat.value >= current.value) ||
            (!higher && atBat.value <= current.value)
        ) {
            score += 1;
            if (score > highScore) {
                highScoreStore.set(score);
            }
            status = 1;
        } else {
            status = 2;
        }
    };

    const nextTurn = async () => {
        const res = await fetch(requestURI);
        distStore.set(-50).then(async () => {
            let tmp = await res.json();
            showValue = false;
            current = atBat;
            atBat = onDeck;
            onDeck = tmp;
            distStore.set(0, noTween);
            status = 0;
            showStatus = true;
        });
    };
</script>

<main transition:fade>
    <div
        id="items"
        style={window.innerWidth > 894
            ? "left: " + dist + "vw;"
            : "top: " + dist + "vh;"}
    >
        <RevealedItem item={current} {valueDescription} />
        <RevealableItem
            item={atBat}
            comparedName={current.name}
            {reveal}
            {afterReveal}
            {valueDescription}
            {shortValueDescription}
            {showValue}
        />
        <UnrevealedItem
            item={onDeck}
            comparedName={atBat.name}
            {shortValueDescription}
        />
    </div>
    {#if showStatus}
        <div id="neutral" transition:scale>VS</div>
        {#if status === 1}
            <div
                id="correct"
                class="material-icons"
                in:scale
                out:scale={{ delay: 750 }}
                on:introend={() => (showStatus = false)}
                on:outroend={nextTurn}
            >
                check
            </div>
        {:else if status === 2}
            <div
                id="incorrect"
                class="material-icons"
                in:scale
                out:scale={{ delay: 750 }}
                on:introend={() => (showStatus = false)}
                on:outroend={onEnd}
            >
                clear
            </div>
        {/if}
    {/if}
    <span id="high-score">High Score: {highScore}</span>
    <span id="score">Score: {score}</span>
</main>

<style>
    main {
        display: block;
    }

    #items {
        position: absolute;

        display: flex;
    }

    #neutral,
    #correct,
    #incorrect {
        width: 10vw;
        height: 10vw;
        max-width: 120px;
        max-height: 120px;
        position: absolute;
        top: calc(50vh - calc(min(10vw, 120px) / 2));
        left: calc(50vw - calc(min(10vw, 120px) / 2));

        display: flex;
        align-items: center;
        justify-content: center;

        border-radius: 50%;
    }

    #neutral {
        font-size: 35px;
        background-color: #999;
    }

    #correct {
        font-size: 52.5px;
        background-color: #00bb00;
    }

    #incorrect {
        font-size: 52.5px;
        background-color: #bb0000;
    }

    #high-score {
        position: absolute;
        top: 5px;
        left: 1vw;

        font-size: 20px;
    }

    #score {
        position: absolute;
        top: 5px;
        right: 1vw;

        font-size: 20px;
    }

    @media (max-width: 894px) {
        #items {
            flex-direction: column;
        }

        #neutral,
        #correct,
        #incorrect {
            width: 6vh;
            height: 6vh;
            top: calc(50vh - calc(min(6vh, 120px) / 2));
            left: calc(50vw - calc(min(6vh, 120px) / 2));
        }

        #neutral {
            font-size: 20px;
        }

        #correct,
        #incorrect {
            font-size: 30px;
        }

        #high-score,
        #score {
            font-size: 16px;
        }
    }
</style>
