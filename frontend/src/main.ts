import App from './App.svelte';

const app = new App({
	target: document.body,
	props: {
		valueDescription: "popularity rating",
		shortValueDescription: "rating",
		requestURI: "https://spotify-higher-lower-api.herokuapp.com/random",
	},
});

export default app;