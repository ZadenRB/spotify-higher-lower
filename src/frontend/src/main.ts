import App from './App.svelte';

const app = new App({
	target: document.body,
	props: {
		popularityDescription: "popularity",
		requestURI: "https://spotify-higher-lower-api.herokuapp.com/random",
	},
});

export default app;