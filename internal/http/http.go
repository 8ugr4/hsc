package http

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

var closeError = errors.New("error closing response body")

type Status struct {
	Code int
	Text string
}

func (s *Status) get(page, auth string) error {
	var request *http.Request
	var response *http.Response
	var err error

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

	s.Code, s.Text = response.StatusCode, response.Status

	return nil
}

func Get(page, auth string) (*Status, error) {
	status := new(Status)
	if err := status.get(page, auth); err != nil {
		return nil, err
	}
	return status, nil
}
