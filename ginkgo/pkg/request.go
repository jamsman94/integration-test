package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CallAPI(method, host, api string, request interface{}) (int, *Response, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return 0, nil, err
	}
	url := fmt.Sprintf("%s/%s", host, api)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	response := new(Response)
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	err = json.Unmarshal(respBody, response)
	if err != nil {
		return 0, nil, err
	}
	return resp.StatusCode, response, nil
}
