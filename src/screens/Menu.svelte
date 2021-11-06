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
            <div id="play" on:click={() => onPlay()}>Play</div>
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

    #play {
        width: 75px;
        padding: 10px 20px;

        border: 1px solid white;
        border-radius: 20px;

        cursor: pointer;

        transition: background-color 150ms ease, color 150ms ease;

        color: #1db964;
        text-align: center;
    }

    #play:hover {
        background-color: white;
        color: #191414;
    }
</style>
