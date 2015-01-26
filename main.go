package main

import (
	"fmt"
	"os"

	"github.com/jamesmura/gotrakt/trakt"
)

func main() {
	api := trakt.Trakt{Url: "https://api.trakt.tv/", ApiKey: os.Getenv("TRAKT_API_KEY"), AccessToken: os.Getenv("TRAKT_ACCESS_TOKEN")}
	shows := api.GetShows("2015-01-26", 5)
	fmt.Println(shows)
}
