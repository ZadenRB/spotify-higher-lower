package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

type Song struct {
	Name       string `json:"name"`
	Artists    string `json:"artists"`
	Popularity int    `json:"popularity"`
	Image      string `json:"image"`
}

var client spotify.Client
var config *clientcredentials.Config

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Get("/random", getRandomSong)

	return router
}

// Generate random query
func randomSearch() string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyz")
	randomRune := letterRunes[rand.Intn(len(letterRunes))]

	var randomSearch string

	switch rand.Intn(2) {
	case 0:
		randomSearch = string(randomRune) + string('%')
		break
	case 1:
		randomSearch = string('%') + string(randomRune)
		break
	}

	return randomSearch
}

// Fetch a random song
func getRandomSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("SPOTIFY_HIGHER_LOWER_FRONTEND"))

	var searchResult *spotify.SearchResult
	var err error
	for ok := true; ok; ok = searchResult == nil || len(searchResult.Tracks.Tracks) == 0 || len(searchResult.Tracks.Tracks[0].Album.Images) == 0 {
		offset := rand.Intn(1001)
		opts := spotify.Options{Offset: &offset}
		search := randomSearch()
		searchResult, err = client.SearchOpt(search, spotify.SearchTypeTrack, &opts)
		if err != nil {
			fmt.Println(err)
		}
	}

	song := searchResult.Tracks.Tracks[0]

	var artistsBuilder strings.Builder
	for i, artist := range song.Artists {
		if i != 0 {
			artistsBuilder.WriteRune(',')
			artistsBuilder.WriteRune(' ')
		}
		artistsBuilder.WriteString(artist.Name)
	}

	randomSong := Song{
		Name:       song.Name,
		Artists:    artistsBuilder.String(),
		Popularity: song.Popularity,
		Image:      song.Album.Images[0].URL,
	}

	render.JSON(w, r, randomSong)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	router := Routes()

	config = &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_HIGHER_LOWER_ID"),
		ClientSecret: os.Getenv("SPOTIFY_HIGHER_LOWER_SECRET"),
		TokenURL:     spotify.TokenURL,
	}

	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	client = spotify.Authenticator{}.NewClient(token)

	quit := make(chan bool)
	defer close(quit)

	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				if token, err = client.Token(); err != nil {
					token, err = config.Token(context.Background())
					client = spotify.Authenticator{}.NewClient(token)
				}
			}
		}
	}()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal(http.ListenAndServe(":8000", router))
	} else {
		log.Fatal(http.ListenAndServe(":"+port, router))
	}
}
