<script lang="ts">
    import { scale, fade } from "svelte/transition";
    import { tweened } from "svelte/motion";
    import { writable } from "svelte/store";

    import RevealableItem from "../components/RevealableItem.svelte";
    import UnrevealedItem from "../components/UnrevealedItem.svelte";
    import RevealedItem from "../components/RevealedItem.svelte";

    export let popularityDescription: string;
    export let requestURI: string;
    export let current: Song;
    export let atBat: Song;
    export let onDeck: Song;

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
            (higher && atBat.popularity >= current.popularity) ||
            (!higher && atBat.popularity <= current.popularity)
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
            : "top: " + dist + "%;"}
    >
        <RevealedItem item={current} {popularityDescription} />
        <RevealableItem
            item={atBat}
            {reveal}
            {afterReveal}
            {popularityDescription}
            {showValue}
        />
        <UnrevealedItem item={onDeck} />
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
        height: 100%;

        position: absolute;

        display: flex;
    }

    #correct,
    #incorrect,
    #neutral {
        width: 10vw;
        height: 10vw;
        max-width: 120px;
        max-height: 120px;
        position: absolute;
        top: calc(50% - calc(min(10vw, 120px) / 2));
        left: calc(50vw - calc(min(10vw, 120px) / 2));

        display: flex;
        align-items: center;
        justify-content: center;

        border-radius: 50%;
    }

    #correct,
    #incorrect {
        font-size: 52.5px;
    }

    #correct {
        background-color: #00bb00;
    }

    #incorrect {
        background-color: #bb0000;
    }

    #neutral {
        font-size: 35px;
        background-color: #555;
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

        #correct,
        #incorrect,
        #neutral {
            width: 6vh;
            height: 6vh;
            top: calc(50% - calc(min(6vh, 120px) / 2));
            left: calc(50vw - calc(min(6vh, 120px) / 2));
        }

        #correct,
        #incorrect {
            font-size: 30px;
        }

        #neutral {
            font-size: 20px;
        }

        #high-score,
        #score {
            font-size: 16px;
        }
    }
</style>
