package http

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	closeError = errors.New("error closing response body")
)

type Status struct {
	Code []int
	Text []string
}

func (s *Status) get(page, auth string) error {
	var request *http.Request
	var response *http.Response
	var err error

	if request, err = http.NewRequest("GET", page, nil); err != nil {
		return fmt.Errorf("error creating request: %s", err)
	}
	request.Header.Set("Authorization", "Basic "+auth)

	if response, err = http.DefaultClient.Do(request); err != nil {
		return fmt.Errorf("error getting response: %s", err)
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
	var (
		wg            sync.WaitGroup
		mu            sync.Mutex
		combinedError error
	)

	for _, url := range urls {
		wg.Add(1)
		status := new(Status)
		go func(url string) {
			defer wg.Done()
			if err := status.get(url, auth); err != nil {
				mu.Lock()
				defer mu.Unlock()
				if combinedError == nil {
					combinedError = err
				}
			} else {
				combinedError = fmt.Errorf("%w: %v", combinedError, err)
			}
		}(url)
	}

	wg.Wait()
	return combinedError
}

func Store(statusCode int, stat string) *Status {
	h := new(Status)
	h.Code = append(h.Code, statusCode)
	h.Text = append(h.Text, stat)
	return h
}
