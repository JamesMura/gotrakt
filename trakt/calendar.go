package trakt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Ids struct {
	Trakt  int    `json:"trakt"`
	Slug   string `json:"slug"`
	Imdb   string `json:"imdb"`
	Tmdb   int    `json:"imdb"`
	Tvdb   int    `json:"tvdb"`
	Tvrage int    `json:"tvrage"`
}
type Show struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	Ids   Ids    `json:"ids"`
}

type Movie struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	Ids   Ids    `json:"ids"`
}

type Episode struct {
	Season int    `json:"season"`
	Number int    `json:"number"`
	Title  string `json:"title"`
	Ids    Ids    `json:"ids"`
}

type Showing struct {
	AirsAt  string  `json:"airs_at"`
	Episode Episode `json:"episode"`
	Show    Show    `json:"show"`
}
type Listing struct {
	Movie Movie `json:"movie"`
}

type CalendarResponse map[string][]Showing
type CalendarMovieResponse map[string][]Listing

func (t Trakt) getCalendarResponse(url string) CalendarResponse {
	resp, err := t.Get(url)
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error : %s", err)
		os.Exit(1)
	}
	var calendarResponse CalendarResponse
	err = json.Unmarshal(contents, &calendarResponse)
	if err != nil {
		fmt.Printf("error : %s", err)
		os.Exit(1)
	}
	return calendarResponse
}

func (t Trakt) getCalendarMovieResponse(url string) CalendarMovieResponse {
	resp, err := t.Get(url)
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error : %s", err)
		os.Exit(1)
	}
	var calendarResponse CalendarMovieResponse
	err = json.Unmarshal(contents, &calendarResponse)
	if err != nil {
		fmt.Printf("error : %s", err)
		os.Exit(1)
	}
	return calendarResponse
}

func (t Trakt) GetShows(startDate string, days int) CalendarResponse {
	path := fmt.Sprintf("calendars/shows/%s/%d", startDate, days)
	return t.getCalendarResponse(path)
}

func (t Trakt) GetNewShows(startDate string, days int) CalendarResponse {
	path := fmt.Sprintf("calendars/shows/new/%s/%d", startDate, days)
	return t.getCalendarResponse(path)
}

func (t Trakt) GetPremieres(startDate string, days int) CalendarMovieResponse {
	path := fmt.Sprintf("calendars/shows/premieres/%s/%d", startDate, days)
	return t.getCalendarMovieResponse(path)
}

func (t Trakt) GetMovies(startDate string, days int) CalendarMovieResponse {
	path := fmt.Sprintf("calendars/movies/%s/%d", startDate, days)
	return t.getCalendarMovieResponse(path)
}
