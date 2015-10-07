package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpGet(url string) (string, error) {

	respChan := make(chan string, 1)
	errChan := make(chan error, 1)

	go func() {
		resp, err := http.Get("http://" + url)
		if err != nil {
			fmt.Println("ERROR::::", err.Error())
			errChan <- err
		}

		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("ERROR::::", err.Error())
			errChan <- err
		}
		respChan <- string(bodyBytes)
	}()

	select {
	case body := <-respChan:
		return body, nil
	case err := <-errChan:
		return "", err
	case <-time.After(3 * time.Second):
		return "Too Busy", nil
	}
}
