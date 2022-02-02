import App from './App.svelte';

const app = new App({
	target: document.body,
	props: {
		popularityDescription: "popularity",
		requestURI: "http://10.110.117.218:8000/random",
		// requestURI: "https://spotify-higher-lower-api.herokuapp.com/random",
	},
});

export default app;