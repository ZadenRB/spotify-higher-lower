<script lang="ts">
    import { fade } from "svelte/transition";
    import Spinner from "../components/Spinner.svelte";

    export let loading: boolean;
    export let restart: () => void;
    export let back: () => void;
</script>

<main transition:fade>
    <span class="big-text">Game Over</span>
    <div class="button-loading-group">
        {#if loading}
            <Spinner />
        {:else}
            <div class="over-buttons">
                <div
                    class="button"
                    tabindex="0"
                    on:click={restart}
                    on:keypress={(event) => {
                        if (event.key == "Enter") {
                            restart();
                        }
                    }}
                >
                    Play Again
                </div>
                <div
                    class="button"
                    tabindex="0"
                    on:click={back}
                    on:keypress={(event) => {
                        if (event.key == "Enter") {
                            back();
                        }
                    }}
                >
                    Back to Menu
                </div>
            </div>
        {/if}
    </div>
</main>

<style>
    .button {
        width: 140px;
    }
</style>
