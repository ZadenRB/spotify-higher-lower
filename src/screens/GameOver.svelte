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
            <div class="buttons">
                <div class="button" on:click={back}>Back to Menu</div>
                <div class="button" on:click={restart}>Play Again</div>
            </div>
        {/if}
    </div>
</main>

<style>
</style>
