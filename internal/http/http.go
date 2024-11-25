package http

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	closeError = errors.New("error closing response body")
	throttle   = time.Tick(time.Second)
	empty      error
)

type Status struct {
	Code []int
	Text []string
}

func (s *Status) get(page, auth string) error {
	var request *http.Request
	var response *http.Response
	var err error

	<-throttle // wait for a tick
	if request, err = http.NewRequest("GET", page, nil); err != nil {
		return errors.New(fmt.Sprintf("error creating request: %s", err))
	}
	request.Header.Set("Authorization", "Basic "+auth)

	if response, err = http.DefaultClient.Do(request); err != nil {
		return errors.New(fmt.Sprintf("error getting response: %s", err))
	}

	defer func() {
		if err = response.Body.Close(); err != nil {
			log.Print(closeError.Error())
		}
	}()

	Store(response.StatusCode, response.Status)
	return nil
}

func Get(urls []string, auth string) error {
	for _, url := range urls {
		status := new(Status)
		go func() {
			err := status.get(url, auth)
			if err != nil {
				empty = errors.New(fmt.Sprintf("failed to get the page: %s", err))
			}
		}()
	}
	return empty
}

func Store(statusCode int, stat string) *Status {
	h := new(Status)
	h.Code = append(h.Code, statusCode)
	h.Text = append(h.Text, stat)
	return h
}
