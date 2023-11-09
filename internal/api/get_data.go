package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func getData[T any](api *Api, url string, parsedData T) (responseData T, err error) {
	if url == "" {
		return responseData, errors.New("Url is missing")
	}

	var data []byte
	cachedResponse, hasCacheHit := api.cache.Get(url)

	if hasCacheHit {
		data = cachedResponse
	} else {
		response, getError := http.Get(url)
		if getError != nil {
			return responseData, errors.New(fmt.Sprint("Could not fetch data:", getError))
		}

		body, readBodyError := io.ReadAll(response.Body)
		response.Body.Close()

		if response.StatusCode > 299 {
			return responseData, errors.New(fmt.Sprintf(
				"Response failed with status code: %d and\nbody: %s\n",
				response.StatusCode, body,
			))
		}
		if readBodyError != nil {
			return responseData, errors.New(fmt.Sprint("Could not read body:", readBodyError))
		}

		data = body
		api.cache.Add(url, data)
	}

	parsingJsonError := json.Unmarshal(data, &parsedData)

	if parsingJsonError != nil {
		return responseData, errors.New(fmt.Sprint("Failed unmarshaling data", parsingJsonError))
	}

	return parsedData, nil
}
