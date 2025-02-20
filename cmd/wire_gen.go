// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package cmd

import (
	"github.com/deluan/navidrome/core"
	"github.com/deluan/navidrome/core/transcoder"
	"github.com/deluan/navidrome/engine"
	"github.com/deluan/navidrome/persistence"
	"github.com/deluan/navidrome/scanner"
	"github.com/deluan/navidrome/server"
	"github.com/deluan/navidrome/server/app"
	"github.com/deluan/navidrome/server/subsonic"
	"github.com/google/wire"
)

// Injectors from wire_injectors.go:

func CreateServer(musicFolder string) *server.Server {
	dataStore := persistence.New()
	scannerScanner := scanner.New(dataStore)
	serverServer := server.New(scannerScanner, dataStore)
	return serverServer
}

func CreateScanner(musicFolder string) *scanner.Scanner {
	dataStore := persistence.New()
	scannerScanner := scanner.New(dataStore)
	return scannerScanner
}

func CreateAppRouter() *app.Router {
	dataStore := persistence.New()
	router := app.New(dataStore)
	return router
}

func CreateSubsonicAPIRouter() (*subsonic.Router, error) {
	dataStore := persistence.New()
	browser := engine.NewBrowser(dataStore)
	imageCache, err := core.NewImageCache()
	if err != nil {
		return nil, err
	}
	cover := core.NewCover(dataStore, imageCache)
	nowPlayingRepository := engine.NewNowPlayingRepository()
	listGenerator := engine.NewListGenerator(dataStore, nowPlayingRepository)
	users := engine.NewUsers(dataStore)
	playlists := engine.NewPlaylists(dataStore)
	ratings := engine.NewRatings(dataStore)
	scrobbler := engine.NewScrobbler(dataStore, nowPlayingRepository)
	search := engine.NewSearch(dataStore)
	transcoderTranscoder := transcoder.New()
	transcodingCache, err := core.NewTranscodingCache()
	if err != nil {
		return nil, err
	}
	mediaStreamer := core.NewMediaStreamer(dataStore, transcoderTranscoder, transcodingCache)
	players := engine.NewPlayers(dataStore)
	router := subsonic.New(browser, cover, listGenerator, users, playlists, ratings, scrobbler, search, mediaStreamer, players)
	return router, nil
}

// wire_injectors.go:

var allProviders = wire.NewSet(engine.Set, core.Set, scanner.New, subsonic.New, app.New, persistence.New)
