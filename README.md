
lfu
=======

[![GoDoc](https://godoc.org/github.com/Sigumaa/lfu?status.svg)](http://godoc.org/github.com/Sigumaa/lfu)

This is a Go wrapper for working with the [Last.fm API](https://www.last.fm/api).

The library is designed to support user API methods in the API Reference. Methods other than user methods are not supported.

Please note that by using this library, you are agreeing to the [Last.fm API Terms of Service](https://www.last.fm/api/tos).

## Installation

To install the library, simply 

```bash
go get github.com/Sigumaa/lfu
```

## Usage

To begin, please register your application at the following page: https://www.last.fm/api/account/create

Upon completion of registration, you will receive an API Key for your application.

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/Sigumaa/lfu"
)

func main() {
	name := os.Getenv("USER_NAME")
	key := os.Getenv("API_KEY")
	
	client := lfu.New(name, key)
	topTracks, err := client.TopTracks(context.TODO(), lfu.Limit(1))
	if err != nil {
		log.Fatal(err)
	}
	
	// This is planned to become more user-friendly in a future update.
	log.Printf("The track you have played the most is %s.", topTracks.TopTracks.Track[0].Name)
}

```


- [x] user.getFriends
- [x] user.getInfo
- [x] user.getLovedTracks
- [x] user.getPersonalTags
- [x] user.getRecentTracks
- [x] user.getTopAlbums
- [x] user.getTopArtists
- [x] user.getTopTags
- [x] user.getTopTracks
- [x] user.getWeeklyAlbumChart
- [x] user.getWeeklyArtistChart
- [x] user.getWeeklyChartList
- [x] user.getWeeklyTrackChart

## TODO

more methods

- [x] user.getFriends
- [ ] user.getInfo
- [ ] user.getLovedTracks
- [ ] user.getPersonalTags
- [ ] user.getRecentTracks
- [ ] user.getTopAlbums
- [ ] user.getTopArtists
- [ ] user.getTopTags
- [ ] user.getTopTracks
- [ ] user.getWeeklyAlbumChart
- [ ] user.getWeeklyArtistChart
- [ ] user.getWeeklyChartList
- [ ] user.getWeeklyTrackChart
