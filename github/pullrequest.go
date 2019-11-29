package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	helpText = `
Example: listPRs <orgname> <repo> <state>
supported state: all, open, closed
	`
)

type PullRequest struct {
	ID        json.Number `json:"id"`
	Number    json.Number `json:"number"`
	State     string      `json:"state"`
	URL       string      `json:"url"`
	Title     string      `json:"title"`
	User      User        `json:"user"`
	CreatedAt string      `json:"created_at"`
}

type User struct {
	Login string `json:"login"`
}

func main() {
	pulls := map[json.Number]PullRequest{}
	flag.Parse()
	if len(os.Args) < 2 {
		panic(helpText)
	}
	org := flag.Arg(0)
	repo := flag.Arg(1)
	state := flag.Arg(2)
	for i := 0; i < 4; i++ {
		resp, err := http.Get("https://api.github.com/repos/" + org + "/" + repo + "/pulls?state=" + state + "&page=" + strconv.Itoa(i))
		if err != nil {
			panic(err)
		}

		prs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			panic(string(prs))
		}

		pull := []PullRequest{}
		if err := json.Unmarshal(prs, &pull); err != nil {
			panic(err)
		}

		for _, p := range pull {
			t, err := time.Parse(time.RFC3339, p.CreatedAt)
			if err != nil {
				panic(err)
			}
			startDate := "2019-10-01T00:00:00Z"
			tt, err := time.Parse(time.RFC3339, startDate)
			if err != nil {
				panic(err)
			}

			endDate := "2019-10-30T23:59:00Z"
			ttt, err := time.Parse(time.RFC3339, endDate)
			if err != nil {
				panic(err)
			}

			if t.After(tt) && t.Before(ttt) {
				_, ok := pulls[p.ID]
				if !ok {
					pulls[p.ID] = p
				}
			}
			resp.Body.Close()
		}
	}

	data, err := json.MarshalIndent(&pulls, "", "\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
