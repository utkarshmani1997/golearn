package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	url := "http://localhost:8080/read"
	body, err := getResponse(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func getResponse(url string) ([]byte, error) {

	if len(url) == 0 {
		return nil, errors.New("Invalid URL")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	c := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := c.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	code := resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil && code != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	if code != http.StatusOK {
		return nil, fmt.Errorf("Server status error: %v", http.StatusText(code))
	}

	return body, nil
}
