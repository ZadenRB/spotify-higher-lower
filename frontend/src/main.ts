import App from './App.svelte';

const app = new App({
	target: document.body,
	props: {
		valueDescription: "popularity rating",
		shortValueDescription: "rating",
		requestURI: "http://localhost:8000/random",
	},
});

export default app;