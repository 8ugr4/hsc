package http

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	closeError = errors.New("error closing response body")
)

type Status struct {
	mu   sync.Mutex
	Code []int
	Text []string
}

func (s *Status) Add(statusCode int, statusText string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Code = append(s.Code, statusCode)
	s.Text = append(s.Text, statusText)
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

	s.Add(response.StatusCode, http.StatusText(response.StatusCode))
	return nil
}

func Get(urls []string, auth string) (*Status, error) {
	var (
		wg     sync.WaitGroup
		mu     sync.Mutex
		status = &Status{}
		errVar error
	)

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			if err := status.get(url, auth); err != nil {
				errVar = err
				mu.Lock()
				defer mu.Unlock()
			}
		}(url)
		time.Sleep(1 * time.Millisecond)
	}

	wg.Wait()
	return status, errVar
}
