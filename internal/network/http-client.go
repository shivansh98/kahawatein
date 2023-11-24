package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-errors/errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

var client *http.Client

func GetClient() *http.Client {
	return client
}

// this will run before main function
func init() {
	log.Println("HEHEHEHEHEHE")
	client = &http.Client{
		Timeout:   time.Second * 100,
		Transport: http.DefaultTransport,
	}
}

// Get used for sending http.Get method to any service
// headers such as request-id and app-env will be added in this method only
func Get(uri string, headers, params map[string]string) (map[string]interface{}, error) {
	var body io.Reader
	var query string

	if len(params) != 0 {
		queryParams := url.Values{}
		for key, val := range params {
			queryParams.Add(key, val)
		}
		uri += "?"
		query = queryParams.Encode()
	}

	url := fmt.Sprintf("%s%s", uri, query)

	req, err := http.NewRequest(http.MethodGet, url, body)
	if err != nil {
		return nil, errors.WrapPrefix(err, "FAILED_TO_SEND_GET_REQUEST", 0)
	}
	if headers == nil {
		headers = make(map[string]string)
	}

	for key, val := range headers {
		req.Header.Add(key, val)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.WrapPrefix(err, "FAILED_TO_SEND_GET_REQUEST", 0)
	}

	if resp == nil || resp.Body == nil {
		return nil, errors.WrapPrefix("resp or body is nil", "FAILED_TO_SEND_GET_REQUEST", 0)
	}

	log.Println("Http Response details received for uri", map[string]interface{}{
		"statusCode": resp.StatusCode,
		"status":     resp.Status,
	})

	defer func() {
		resp.Body.Close()
	}()

	respObj := make(map[string]interface{})
	if err = json.NewDecoder(resp.Body).Decode(&respObj); err != nil {
		return nil, errors.WrapPrefix(err, "FAILED_TO_SEND_GET_REQUEST_RESP_MARSHAL_ERROR", 0)
	}
	return respObj, err
}

// Post used for sending http.Post method to any service
// headers such as request-id and app-env will be added in this method only
func Post(uri string, headers, params map[string]string, payload []byte) (map[string]interface{}, error) {
	var body io.Reader
	var query string
	if len(params) != 0 {
		queryParams := url.Values{}
		for key, val := range params {
			queryParams.Add(key, val)
		}
		uri += "?"
		query = queryParams.Encode()
	}

	if len(payload) != 0 {
		body = bytes.NewBuffer(payload)
	}

	url := fmt.Sprintf("%s%s", uri, query)

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, errors.WrapPrefix(err, "FAILED_TO_SEND_POST_REQUEST", 0)
	}

	for key, val := range headers {
		req.Header.Add(key, val)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.WrapPrefix(err, "FAILED_TO_SEND_POST_REQUEST", 0)
	}

	log.Println("Http Response details received for uri", map[string]interface{}{
		"statusCode": resp.StatusCode,
		"status":     resp.Status,
	})

	if resp == nil || resp.Body == nil {
		return nil, errors.WrapPrefix("resp or body is nil", "FAILED_TO_SEND_POST_REQUEST", 0)
	}

	defer func() {
		resp.Body.Close()
	}()

	respObj := make(map[string]interface{})
	if err = json.NewDecoder(resp.Body).Decode(&respObj); err != nil {
		return nil, errors.WrapPrefix(err, "FAILED_TO_SEND_POST_REQUEST_RESP_MARSHAL_ERROR", 0)
	}

	if resp.StatusCode > 206 || resp.StatusCode < 200 {
		respObj["error_response_code"] = resp.StatusCode
	}

	return respObj, err
}
