<div id="loading" />

<style>
    @keyframes spinner {
        0% {
            transform: rotate(0deg);
        }
        100% {
            transform: rotate(360deg);
        }
    }

    #loading {
        animation: 1.5s linear infinite spinner;
        animation-play-state: running;
        border: solid 5px #fff;
        border-bottom-color: #1db964;
        border-radius: 50%;
        height: 110px;
        width: 110px;
        transform: rotate(0deg);
        will-change: transform;
    }
</style>
