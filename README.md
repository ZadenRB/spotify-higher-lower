# spotify-higher-lower

A guessing game where the player must guess if Spotify songs are more or less popular than one another.

![](https://github.com/ZadenRB/spotify-higher-lower/blob/master/Higher-Lower.gif)

## Setup & Prerequisites

### API

1. Install [Go](https://go.dev/doc/install)
2. Get packages `go get -v`
3. Create a [Spotify app](https://developer.spotify.com/dashboard/applications)
4. Set environment variables (add to `rc` file if desired)
	1. `export SPOTIFY_HIGHER_LOWER_FRONTEND=<URL of Frontend>` - For local development, use `http://localhost:5000`
	2. `export SPOTIFY_HIGHER_LOWER_ID=<Spotify Client ID>` - The client ID of your Spotify app
	3. `export SPOTIFY_HIGHER_LOWER_SECRET=<Spotify Client Secret>` - The client secret of your Spotify app

### Frontend

1. Install [npm](https://docs.npmjs.com/cli/v7/configuring-npm/install)
2. Get node modules `npm install`
3. In src/main.ts, set `requestUri: <API URI>` - For local development, use `http://localhost:8000/random`
	
## Running

### Running the API

	cd api
	go run main.go

### Running the Frontend

	cd frontend
	npm run dev

## Building

### Building the API

	cd api
	go build main.go

### Building the Frontend

	cd frontend
	npm run build
