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
Command: listPRs <orgname> <repo> <state> <start_time> <end_time> <no_of_pages>

supported args:
	state: all, open, closed
	orgname: organization name
	repo: repository name
	start time: 2019-10-01T00:00:00Z
	end time: 2019-10-01T00:00:00Z
	no of pages: default is 4

Example: ./github openebs maya all 2019-10-01T00:00:00Z 2019-10-01T00:00:00Z 5
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
	var (
		startDate, endDate time.Time
		pages              int = 4
	)

	sd, ed := "2019-10-01T00:00:00Z", "2019-10-30T23:59:00Z"
	pulls := map[json.Number]PullRequest{}
	var err error
	flag.Parse()
	if len(os.Args) < 4 {
		panic(helpText)
	}
	org := flag.Arg(0)
	repo := flag.Arg(1)
	state := flag.Arg(2)

	if len(flag.Arg(3)) != 0 {
		sd = flag.Arg(3)
	}
	startDate, err = time.Parse(time.RFC3339, sd)
	if err != nil {
		panic(err)
	}

	if len(flag.Arg(4)) != 0 {
		ed = flag.Arg(4)
	}
	endDate, err = time.Parse(time.RFC3339, ed)
	if err != nil {
		panic(err)
	}

	if len(flag.Arg(5)) != 0 {
		pages, err = strconv.Atoi(flag.Arg(5))
		if err != nil {
			panic(err)
		}
	}

	for i := 0; i < pages; i++ {
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
			if t.After(startDate) && t.Before(endDate) {
				_, ok := pulls[p.ID]
				if !ok {
					pulls[p.ID] = p
				}
			}
		}
		resp.Body.Close()
	}

	data, err := json.MarshalIndent(&pulls, "", "\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
