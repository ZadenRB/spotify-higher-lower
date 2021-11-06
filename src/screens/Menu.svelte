<script lang="ts">
    import { fade } from "svelte/transition";
    import Spinner from "../components/Spinner.svelte";

    export let onPlay: () => void;
    export let loading: boolean;
</script>

<main transition:fade>
    <div id="title">
        <span class="big-text">Higher-Lower</span>
        <br />
        <span class="small-text">for</span>
        <img id="spotify" src="/img/spotify_logo_white.png" alt="Spotify" />
    </div>
    <div class="button-loading-group">
        {#if loading}
            <Spinner />
        {:else}
            <div class="button" on:click={() => onPlay()}>Play</div>
        {/if}
    </div>
</main>

<style>
    #title {
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    #spotify {
        display: block;
        width: auto;
        height: 10vw;
        min-width: auto;
        min-height: 21px;
        margin: 4.5vw 4.5vw;
    }

    .button {
        width: 100px;
    }
</style>
