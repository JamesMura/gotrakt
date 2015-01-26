package trakt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Show struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
}

type Episode struct {
	Season int    `json:"season"`
	Number int    `json:"number"`
	Title  string `json:"title"`
}

type Showing struct {
	AirsAt  string  `json:"airs_at"`
	Episode Episode `json:"episode"`
	Show    Show    `json:"show"`
}

type CalendarResponse map[string][]Showing

func (t Trakt) GetShows(startDate string, days int) CalendarResponse {
	path := fmt.Sprintf("calendars/shows/%s/%d", startDate, days)
	resp, err := t.Get(path)
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
